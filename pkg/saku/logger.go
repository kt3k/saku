package saku

import (
	"fmt"

	"github.com/fatih/color"
)

type logger struct {
	enabled bool
}

func (l *logger) print(a ...interface{}) {
	if l.enabled {
		fmt.Print(a...)
	}
}

func (l *logger) println(a ...interface{}) {
	if l.enabled {
		fmt.Println(a...)
	}
}

func (l *logger) printlnError(a ...interface{}) {
	fmt.Print(color.RedString("Error:"), " ")
	fmt.Println(a...)
}
