package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/simonleung8/flags"
)

func main() {
	fc := flags.New()

	fc.NewBoolFlag("help", "h", "Show the help message and exits.")
	fc.NewBoolFlag("version", "v", "")
	fc.NewBoolFlag("parallel", "p", "Runs tasks in parallel.")
	fc.NewBoolFlag("race", "r", "")
	fc.NewBoolFlag("serial", "s", "")
	fc.NewStringFlagWithDefault("config", "c", "", "saku.md")

	err := fc.Parse(os.Args...)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(exitCodeError)
	}

	if fc.Bool("help") {
		usage()
		os.Exit(exitCodeOk)
	}

	if fc.Bool("version") {
		fmt.Printf("saku@%s\n", Version)
		os.Exit(exitCodeOk)
	}

	config, err1 := readConfig()

	if err1 != nil {
		fmt.Println(color.RedString("Error:"), "File not found: saku.md")
		os.Exit(exitCodeError)
	}

	tasks := ParseConfig(&config)

	titles := fc.Args()[1:]

	if len(titles) == 0 {
		fmt.Println("There are", color.MagentaString(strconv.Itoa(len(tasks.tasks))), "task(s)")

		for _, t := range tasks.tasks {
			fmt.Println("  " + color.CyanString("["+t.title+"]"))
			for _, desc := range t.descriptions {
				fmt.Println("    " + desc)
			}
		}

		os.Exit(exitCodeOk)
	}

	for _, title := range titles {
		_, ok := tasks.getByTitle(title)

		if !ok {
			fmt.Println(color.RedString("Error:"), "Task not defined:", title)
			os.Exit(exitCodeOk)
		}
	}

	runTasks := tasks.filterByTitles(titles)

	fmt.Println(color.CyanString("[saku]"), "Run", color.MagentaString(strings.Join(titles, ", ")), "in", color.CyanString("sequence"))

	err0 := runTasks.Run(&runOptions{})


	if err0 != nil {
		fmt.Println(color.RedString("Error:"), err0)
		os.Exit(exitCodeError)
	} else {
	  fmt.Println(color.CyanString("[saku]"), "Finish", color.MagentaString(strings.Join(titles, ", ")), "in", color.CyanString("sequence"))
	}
}
