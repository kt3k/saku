package saku

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func actionRun(titles []string, tasks *TaskCollection, l *logger, runOpts *runOptions) ExitCode {
	done := make(chan error, 1)
	sigs := make(chan os.Signal, 1)

	for _, title := range titles {
		_, ok := tasks.getByTitle(title)

		if !ok {
			fmt.Println(color.RedString("Error:"), "Task not defined:", title)
			return ExitCodeError
		}
	}

	if runOpts.isSerialAndParallel() {
		fmt.Println(color.RedString("Error:"), "both --serial and --parallel options are specified")
		return ExitCodeError
	}

	runTasks := tasks.filterByTitles(titles)

	logLine("Run", titles, l, runOpts)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		sig := <-sigs

		runTasks.abort()

		done <- fmt.Errorf("aborted: signal=%s", sig)
	}()

	go func() {
		for {
			l.println("+" + <-runTasks.onCommand)
		}
	}()

	go func() {
		done <- runTasks.Run(runOpts)
	}()

	err0 := <-done

	if err0 != nil {
		fmt.Println(color.RedString("Error:"), err0)

		return ExitCodeError
	}

	logLine(prependEmoji("✨", "Finish", emojiEnabled() && !invokedInSaku()), titles, l, runOpts)

	return ExitCodeOk
}

// logLine logs a line of saku's phase message.
func logLine(phaseLabel string, titles []string, l *logger, runOpts *runOptions) {
	l.print(color.CyanString("[saku]"), " ", phaseLabel, " ")

	l.print(color.MagentaString(strings.Join(titles, ", ")))

	if len(titles) > 1 {
		l.print(" in ", color.CyanString(string(runOpts.runMode())))
	}

	l.println()
}
