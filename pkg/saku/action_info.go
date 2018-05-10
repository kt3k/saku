package saku

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

func actionInfo(tasks *TaskCollection) ExitCode {
	fmt.Println("There are", color.MagentaString(strconv.Itoa(len(tasks.tasks))), "task(s)")

	for _, t := range tasks.tasks {
		fmt.Println("  " + color.CyanString("["+t.title+"]"))
		if len(t.descriptions) == 0 {
			fmt.Println("    " + color.New(color.Italic).Sprint("No description"))
		}

		for _, desc := range t.descriptions {
			fmt.Println("    " + desc)
		}
	}

	return ExitCodeOk
}
