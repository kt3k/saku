package saku

// ExitCode is the exit code of saku's Run function
type ExitCode int

// The exit codes
const (
	ExitCodeOk ExitCode = iota
	ExitCodeError
)

// Version is the version number of saku.
const Version = "0.1.2"
