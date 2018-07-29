package saku

import (
	"github.com/fatih/color"
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
	runOpts, err := parseRunOptions(args)

	l := &logger{enabled: true}

	if err != nil {
		l.printlnError(err)
		return ExitCodeError
	}

	l.enabled = !runOpts.isQuiet()

	if runOpts.isHelpAction() {
		actionHelp()
		return ExitCodeOk
	}

	if runOpts.isVersionAction() {
		actionVersion()
		return ExitCodeOk
	}

	configFile := runOpts.config()

	config, err1 := readConfig(cwd, configFile, l)

	if err1 != nil {
		if configFile != defaultConfigFile {
			l.printlnError("File not found:", configFile)
		} else {
			l.printlnError("File not found:", configFile, "\n  First you need to set up", color.CyanString("saku.md"))
			l.println("  See", color.MagentaString("https://github.com/kt3k/saku"), "for details")
		}

		return ExitCodeError
	}

	tasks := ParseConfig(&config)

	titles := runOpts.titles()
	runCtx := &runContext{
		l:         l,
		extraArgs: runOpts.extraArgs,
		mode:      runOpts.runMode(),
	}

	if runOpts.isInfoAction() {
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
