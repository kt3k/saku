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

	l.print(color.CyanString("[saku]"), " Run ", color.MagentaString(strings.Join(titles, ", ")))

	if len(titles) > 1 {
		l.print(" ", runOpts.runLabel())
	}

	l.println()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		sig := <-sigs

		runTasks.abort()

		done <- fmt.Errorf("aborted: signal=%s", sig)
	}()

	go func () {
		for {
			l.println("+", <-runTasks.onCommand)
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

	l.print(color.CyanString("[saku]"), " ", prependEmoji("✨", "Finish ", emojiEnabled() && !invokedInSaku()))

	l.print(color.MagentaString(strings.Join(titles, ", ")))

	if len(titles) > 1 {
		l.print(" ", runOpts.runLabel())
	}

	l.println()

	return ExitCodeOk
}
