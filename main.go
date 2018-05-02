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
	exitCode := run(os.Args...)

	os.Exit(int(exitCode))
}

func run(args ...string) exitCode {
	fc := flags.New()

	fc.NewBoolFlag("help", "h", "Show the help message and exits.")
	fc.NewBoolFlag("version", "v", "")
	fc.NewBoolFlag("parallel", "p", "Runs tasks in parallel.")
	fc.NewBoolFlag("race", "r", "")
	fc.NewBoolFlag("serial", "s", "")
	fc.NewStringFlagWithDefault("config", "c", "", "saku.md")

	err := fc.Parse(args...)

	if err != nil {
		fmt.Println(color.RedString("Error:"), err)
		return exitCodeError
	}

	if fc.Bool("help") {
		usage()
		return exitCodeOk
	}

	if fc.Bool("version") {
		fmt.Printf("saku@%s\n", Version)
		return exitCodeOk
	}

	configFile := fc.String("config")

	config, err1 := readConfig(configFile)

	if err1 != nil {
		fmt.Println(color.RedString("Error:"), "File not found:", configFile)
		return exitCodeError
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

		return exitCodeOk
	}

	for _, title := range titles {
		_, ok := tasks.getByTitle(title)

		if !ok {
			fmt.Println(color.RedString("Error:"), "Task not defined:", title)
			return exitCodeError
		}
	}

	runTasks := tasks.filterByTitles(titles)

	fmt.Println(color.CyanString("[saku]"), "Run", color.MagentaString(strings.Join(titles, ", ")), "in", color.CyanString("sequence"))

	err0 := runTasks.Run(&runOptions{})

	if err0 != nil {
		fmt.Println(color.RedString("Error:"), err0)

		return exitCodeError
	} else {
		fmt.Println(color.CyanString("[saku]"), "Finish", color.MagentaString(strings.Join(titles, ", ")), "in", color.CyanString("sequence"))
	}

	return exitCodeOk
}
