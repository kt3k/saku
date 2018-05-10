package main

import (
	"github.com/fatih/color"
	"github.com/simonleung8/flags"
)

type runOptions struct {
	cwd       string
	fc        flags.FlagContext
	extraArgs []string
}

func (r *runOptions) runLabel() string {
	if r.isParallel() {
		return "in " + color.CyanString("parallel")
	} else if r.isRace() {
		return "in " + color.CyanString("parallel-race")
	} else {
		return "in " + color.CyanString("sequence")
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
