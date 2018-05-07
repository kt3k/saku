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
	os.Exit(int(Run(cwd, os.Args[1:]...)))
}

// Run saku command in the given cwd and arguments
func Run(cwd string, args ...string) ExitCode {
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
		return ExitCodeError
	}

	if fc.Bool("help") {
		usage()
		return ExitCodeOk
	}

	if fc.Bool("version") {
		fmt.Printf("saku@%s\n", Version)
		return ExitCodeOk
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

		return ExitCodeError
	}

	tasks := ParseConfig(&config)

	titles := fc.Args()

	if len(titles) == 0 {
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

	for _, title := range titles {
		_, ok := tasks.getByTitle(title)

		if !ok {
			fmt.Println(color.RedString("Error:"), "Task not defined:", title)
			return ExitCodeError
		}
	}

	runOpts := &runOptions{cwd: "", fc: fc}

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
