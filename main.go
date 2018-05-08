// +build !coverage

package main

import (
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	os.Exit(int(Run(cwd, os.Args[1:]...)))
}
