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
		os.Exit(exitCodeError)
	}

	if fc.Bool("help") {
		usage()
		os.Exit(exitCodeOk)
	}

	if fc.Bool("version") {
		fmt.Printf("saku@%s\n", Version)
		os.Exit(exitCodeOk)
	}

	config, err1 := readConfig()

	if err1 != nil {
		fmt.Println("Error: File not found: saku.md")
		os.Exit(exitCodeOk)
	}

	tasks := parseConfig(&config)

	err0 := tasks.run()

	if err0 != nil {
		fmt.Println("Error:", err0)
		os.Exit(exitCodeOk)
	}
}
