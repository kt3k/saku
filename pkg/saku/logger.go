package saku

import (
	"fmt"
	"strings"

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

// logLine logs a line of saku's phase message.
func (l *logger) logPhase(phaseLabel string, tc *TaskCollection, stack *taskStack) {
	titles := tc.titles()

	l.print(color.CyanString("[saku]"), " ")

	if !stack.isEmpty() {
		for _, t := range stack.tasks {
			l.print(t.title, " > ")
		}
	}

	l.print(phaseLabel, " ")

	l.print(color.MagentaString(strings.Join(titles, ", ")))

	if len(titles) > 1 {
		l.print(" in ", color.CyanString(string(tc.mode)))
	}

	l.println()
}

func (l *logger) logStart(tc *TaskCollection, stack *taskStack) {
	l.logPhase("Run", tc, stack)
}

func (l *logger) logEnd(tc *TaskCollection, stack *taskStack) {
	l.logPhase(prependEmoji("✨", "Finish", emojiEnabled() && !invokedInSaku() && stack.isEmpty()), tc, stack)
}
