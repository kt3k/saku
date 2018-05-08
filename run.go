package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/simonleung8/flags"
)

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
		return actionHelp()
	}

	if fc.Bool("version") {
		return actionVersion()
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
		return actionInfo(tasks)
	}

	runOpts := &runOptions{cwd: "", fc: fc}

	return actionRun(titles, tasks, runOpts)
}
