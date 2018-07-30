package saku

import (
	"github.com/simonleung8/flags"
)

type runOptions struct {
	fc        flags.FlagContext
	extraArgs []string
}

func parseRunOptions(args []string) (*runOptions, error) {
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

	err := fc.Parse(mainArgs...)

	if err != nil {
		return nil, err
	}

	return &runOptions{
		fc:        fc,
		extraArgs: extraArgs,
	}, nil

}

func (r *runOptions) runMode() RunMode {
	if r.isParallel() {
		return RunModeParallel
	} else if r.isRace() {
		return RunModeParallelRace
	} else {
		return RunModeSequence
	}
}

func (r *runOptions) isSerialAndParallel() bool {
	return r.fc.Bool("serial") && r.fc.Bool("parallel")
}

func (r *runOptions) isParallel() bool {
	return r.fc.Bool("parallel") && !r.fc.Bool("race")
}

func (r *runOptions) isRace() bool {
	return r.fc.Bool("parallel") && r.fc.Bool("race")
}

func (r *runOptions) titles() []string {
	return r.fc.Args()
}

func (r *runOptions) isQuiet() bool {
	return r.fc.Bool("quiet")
}

func (r *runOptions) isHelpAction() bool {
	return r.fc.Bool("help")
}

func (r *runOptions) isVersionAction() bool {
	return r.fc.Bool("version")
}

func (r *runOptions) config() string {
	return r.fc.String("config")
}

func (r *runOptions) isInfoAction() bool {
	return len(r.titles()) == 0 || r.fc.Bool("info")
}
