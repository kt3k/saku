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

	err := fc.Parse()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(StatusCodeError)
	}

	fmt.Println("hello from saku")
	fmt.Println("Args", fc.Args())
	fmt.Println("help", fc.Bool("help))
}
