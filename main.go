package main

import (
	"fmt"
	"os"

	"github.com/simonleung8/flags"
)

func main() {
	fc := flags.New()

	fc.NewBoolFlag("help", "h", "Show the help message and exits.")
	fc.NewBoolFlag("version", "v", "")
	fc.NewBoolFlag("parallel", "p", "Runs tasks in parallel.")
	fc.NewBoolFlag("race", "r", "")
	fc.NewBoolFlag("serial", "s", "")
	fc.NewStringFlagWithDefault("config", "c", "", "saku.md")

	err := fc.Parse(os.Args...)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(ExitCodeError)
	}

	if fc.Bool("help") {
		Usage()
		os.Exit(ExitCodeOk)
	}

	if fc.Bool("version") {
		fmt.Printf("saku@%s\n", Version)
		os.Exit(ExitCodeOk)
	}

	tasks := ParseTasks()
	tasks.Run()
}
