package saku

import (
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
	fc.NewBoolFlag("quiet", "q", "")
	fc.NewStringFlagWithDefault("config", "c", "", defaultConfigFile)

	mainArgs, extraArgs := separateExtraArgs(args)

	l := &logger{enabled: !fc.Bool("quiet")}

	err := fc.Parse(mainArgs...)

	if err != nil {
		l.printlnError(err)
		return ExitCodeError
	}

	if fc.Bool("help") {
		actionHelp()
		return ExitCodeOk
	}

	if fc.Bool("version") {
		actionVersion()
		return ExitCodeOk
	}

	configFile := fc.String("config")

	config, err1 := readConfig(cwd, configFile, l)

	if err1 != nil {
		if configFile != defaultConfigFile {
			l.printlnError("File not found:", configFile)
		} else {
			l.printlnError("File not found:", configFile)
			l.println("  First you need to set up", color.CyanString("saku.md"))
			l.println("  See", color.MagentaString("https://github.com/kt3k/saku"), "for details")
		}

		return ExitCodeError
	}

	tasks := ParseConfig(&config)

	titles := fc.Args()

	runOpts := &runOptions{fc: fc}
	runCtx := &runContext{
		l:         l,
		extraArgs: extraArgs,
		mode:      runOpts.runMode(),
	}

	if len(titles) == 0 || fc.Bool("info") {
		actionInfo(tasks)
		return ExitCodeOk
	}

	if runOpts.isSerialAndParallel() {
		l.printlnError("both --serial and --parallel options are specified")
		return ExitCodeError
	}

	err2 := actionRun(titles, tasks, runCtx)

	if err2 != nil {
		l.printlnError(err2)
		return ExitCodeError
	}

	return ExitCodeOk
}
