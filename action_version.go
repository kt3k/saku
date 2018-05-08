package main

import "fmt"

func actionVersion() ExitCode {
	fmt.Printf("saku@%s\n", Version)

	return ExitCodeOk
}
