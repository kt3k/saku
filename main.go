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
	cwd, _ := os.Getwd()
	os.Exit(int(run(cwd, os.Args...)))
}

func run(cwd string, args ...string) exitCode {
	fc := flags.New()

	fc.NewBoolFlag("help", "h", "Show the help message and exits.")
	fc.NewBoolFlag("version", "v", "")
	fc.NewBoolFlag("parallel", "p", "Runs tasks in parallel.")
	fc.NewBoolFlag("race", "r", "")
	fc.NewBoolFlag("serial", "s", "")
	fc.NewStringFlagWithDefault("config", "c", "", defaultConfigFile)

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

	config, err1 := readConfig(cwd, configFile)

	if err1 != nil {
		if configFile != defaultConfigFile {
			fmt.Println(color.RedString("Error:"), "File not found:", configFile)
		} else {
			fmt.Println(color.RedString("Error:"), "File not found:", configFile)
			fmt.Println("  And <!-- saku start --><!-- saku end --> directive not found in README.md as well")
		}

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

	runOpts := &runOptions{cwd: "", fc: fc}

	if runOpts.isSerialAndParallel() {
		fmt.Println(color.RedString("Error:"), "both --serial and --parallel options are specified")
		return exitCodeError
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

		return exitCodeError
	} else {
		fmt.Print(color.CyanString("[saku]"))

		if !invokedInSaku() {
			fmt.Print(" ", prependEmoji("✨", "Finish "))
		} else {
			fmt.Print(" Finish ")
		}

		fmt.Print(color.MagentaString(strings.Join(titles, ", ")))

		if len(titles) > 1 {
			fmt.Print(" ", runOpts.runLabel())
		}

		fmt.Println()
	}

	return exitCodeOk
}
