package saku

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func actionInfo(tasks *TaskCollection) error {
	colorablePrintln("There are", color.MagentaString(strconv.Itoa(tasks.taskCount())), "task(s)")

	printTasks(tasks)

	return nil
}

func printTasks(tasks *TaskCollection) {
	for _, t := range tasks.tasks {
		indent := strings.Repeat("  ", t.level)
		colorablePrintln(indent + color.CyanString("["+t.title+"]"))
		if len(t.descriptions) == 0 {
			colorablePrintln(indent + "  " + color.New(color.Italic).Sprint("No description"))
		}

		for _, desc := range t.descriptions {
			colorablePrintln(indent + "  " + desc)
		}

		if t.children != nil {
			printTasks(t.children)
		}
	}
}
