package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func actionRun(titles []string, tasks *TaskCollection, runOpts *runOptions) ExitCode {
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

	fmt.Print(color.CyanString("[saku]"), " Run ", color.MagentaString(strings.Join(titles, ", ")))

	if len(titles) > 1 {
		fmt.Print(" ", runOpts.runLabel())
	}

	fmt.Println()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		sig := <-sigs

		runTasks.abort()

		done <- fmt.Errorf("aborted: signal=%s", sig)
	}()

	go func() {
		done <- runTasks.Run(runOpts)
	}()

	err0 := <-done

	if err0 != nil {
		fmt.Println(color.RedString("Error:"), err0)

		return ExitCodeError
	}

	fmt.Print(color.CyanString("[saku]"), " ", prependEmoji("✨", "Finish ", emojiEnabled() && !invokedInSaku()))

	fmt.Print(color.MagentaString(strings.Join(titles, ", ")))

	if len(titles) > 1 {
		fmt.Print(" ", runOpts.runLabel())
	}

	fmt.Println()

	return ExitCodeOk
}
