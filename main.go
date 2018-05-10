// +build !coverage

package main

import (
	"os"

	saku "github.com/kt3k/saku/pkg/saku"
)

func main() {
	cwd, _ := os.Getwd()
	os.Exit(int(saku.Run(cwd, os.Args[1:]...)))
}
