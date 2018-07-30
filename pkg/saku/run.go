package saku

import (
	"fmt"

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
	l := &logger{enabled: true}

	err := selectActions(cwd, l, args...)

	if err != nil {
		l.printlnError(err)
		return ExitCodeError
	}

	return ExitCodeOk
}

// Run saku command in the given cwd and arguments
func selectActions(cwd string, l *logger, args ...string) error {
	runOpts, err := parseRunOptions(args)

	if err != nil {
		return err
	}

	l.enabled = !runOpts.isQuiet()

	if runOpts.isHelpAction() {
		return actionHelp()
	}

	if runOpts.isVersionAction() {
		return actionVersion()
	}

	configFile := runOpts.config()

	config, err1 := readConfig(cwd, configFile, l)

	if err1 != nil {
		if configFile != defaultConfigFile {
			return fmt.Errorf("File not found: " + configFile)
		}

		return fmt.Errorf("File not found: " + configFile + "\n  First you need to set up " + color.CyanString("saku.md") + "\n  See " + color.MagentaString("https://github.com/kt3k/saku") + " for details")
	}

	tasks := ParseConfig(&config)

	titles := runOpts.titles()
	runCtx := &runContext{
		l:         l,
		extraArgs: runOpts.extraArgs,
		mode:      runOpts.runMode(),
	}

	if runOpts.isInfoAction() {
		return actionInfo(tasks)
	}

	if runOpts.isSerialAndParallel() {
		return fmt.Errorf("both --serial and --parallel options are specified")
	}

	return actionRun(titles, tasks, runCtx)
}
