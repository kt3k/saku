package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Shows the help message.
func actionHelp() ExitCode {
	fmt.Printf(`
  Usage: %s [options] <task, ...> [-- extra-options]

  Options:
    -v, --version   - - - Shows the version number and exits.
    -h, --help  - - - - - Shows the help message and exits.
    -i, --info  - - - - - Shows the task information and exits.
    -p, --parallel  - - - Runs tasks in parallel. Default false.
    -s, --sequential  - - Runs tasks in serial. Default true.
    -c, --config <path> - Specifies the config file. Default is 'saku.md'.
    -r, --race  - - - - - Set the flag to kill all tasks when a task
                          finished with zero. This option is valid only
                          with 'parallel' option.
    -q, --quiet   - - - - Stops the logging.

  The extra options after '--' are passed to each task command.

`, color.CyanString("saku"))

	return ExitCodeOk
}
