package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func actionRun(titles []string, tasks *TaskCollection, runOpts *runOptions) ExitCode {
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

	err0 := runTasks.Run(runOpts)

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
