package main

// ExitCode is the exit code of saku's Run function
type ExitCode int

// The exit codes
const (
	ExitCodeOk ExitCode = iota
	ExitCodeError
)
