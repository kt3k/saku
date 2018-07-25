package saku

// ExitCode is the exit code of saku's Run function
type ExitCode int

// The exit codes
const (
	ExitCodeOk ExitCode = iota
	ExitCodeError
)

// RunMode represents the mode of running multiple tasks
type RunMode string

// The run modes
const (
	RunModeSequence     RunMode = "sequence"
	RunModeParallel     RunMode = "parallel"
	RunModeParallelRace RunMode = "parallel-race"
)

// Version is the version number of saku.
const Version = "1.1.0"
