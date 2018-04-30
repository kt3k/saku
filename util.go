package main

import (
	"github.com/mattn/go-isatty"
	"os"
)

// Returns the string prepended by the given emoji when the terminal is tty, otherwise drops emoji and returns the string.
func prependEmoji(e string, str string) string {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return str
	}

	return e + "  " + str
}

// Returns true if the process is invoked in saku.
func invokedInSaku() bool {
	return os.Getenv("IN_SAKU") == "true"
}
