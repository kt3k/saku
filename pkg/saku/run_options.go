package saku

import (
	"github.com/simonleung8/flags"
)

type runOptions struct {
	cwd       string
	fc        flags.FlagContext
	extraArgs []string
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
