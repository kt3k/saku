package saku

import "fmt"

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
