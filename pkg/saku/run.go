package saku

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/simonleung8/flags"
)

func separateExtraArgs(args []string) ([]string, []string) {
	for i, arg := range args {
		if arg == "--" {
			return args[:i], args[i+1:]
		}
	}

	return args, []string{}
}

// Run saku command in the given cwd and arguments
func Run(cwd string, args ...string) ExitCode {
	fc := flags.New()

	fc.NewBoolFlag("help", "h", "Show the help message and exits.")
	fc.NewBoolFlag("version", "v", "")
	fc.NewBoolFlag("parallel", "p", "Runs tasks in parallel.")
	fc.NewBoolFlag("race", "r", "")
	fc.NewBoolFlag("serial", "s", "")
	fc.NewBoolFlag("info", "i", "")
	fc.NewStringFlagWithDefault("config", "c", "", defaultConfigFile)

	mainArgs, extraArgs := separateExtraArgs(args)

	err := fc.Parse(mainArgs...)

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
			fmt.Println("  First you need to set up", color.CyanString("saku.md"))
			fmt.Println("  See", color.MagentaString("https://github.com/kt3k/saku"), "for details")
		}

		return ExitCodeError
	}

	tasks := ParseConfig(&config)

	titles := fc.Args()

	if len(titles) == 0 || fc.Bool("info") {
		return actionInfo(tasks)
	}

	runOpts := &runOptions{cwd: "", fc: fc, extraArgs: extraArgs}

	return actionRun(titles, tasks, runOpts)
}
