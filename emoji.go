package main

import (
	"os"
	"github.com/mattn/go-isatty"
)

// Returns the string prepended by the given emoji when the terminal is tty, otherwise drops emoji and returns the string.
func prependEmoji(e string, str string) string {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return str
	}

	return e + "  " + str
}
