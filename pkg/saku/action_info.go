package saku

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func actionInfo(tasks *TaskCollection) {
	fmt.Println("There are", color.MagentaString(strconv.Itoa(tasks.taskCount())), "task(s)")

	printTasks(tasks)
}

func printTasks(tasks *TaskCollection) {
	for _, t := range tasks.tasks {
		indent := strings.Repeat("  ", t.level)
		fmt.Println(indent + color.CyanString("["+t.title+"]"))
		if len(t.descriptions) == 0 {
			fmt.Println(indent + "  " + color.New(color.Italic).Sprint("No description"))
		}

		for _, desc := range t.descriptions {
			fmt.Println(indent + "  " + desc)
		}

		if t.children != nil {
			printTasks(t.children)
		}
	}
}
