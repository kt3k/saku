package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

type task struct {
	title        string
	descriptions []string
	commands     []string
	options      taskOptions
	aborted      bool
}

func newTask() task {
	return task{
		title:        "",
		descriptions: []string{},
		commands:     []string{},
		options:      taskOptions{},
		aborted:      false,
	}
}

type taskOptions struct {
}

// Runs a task.
func (t *task) run(opts *runOptions) error {
	for _, command := range t.commands {
		if t.aborted {
			return nil
		}

		fmt.Println("+" + command)
		err := execCommand(command)

		if err != nil {
			return errors.New("Task " + color.MagentaString(t.title) + " failed")
		}
	}

	return nil
}

// Aborts a task.
func (t *task) abort() {
	if t.aborted {
		return
	}
}

// Adds the description.
func (t *task) addDescription(description string) {
	t.descriptions = append(t.descriptions, description)
}

// Sets the title.
func (t *task) setTitle(title string) {
	t.title = title
}

// Adds the code.
func (t *task) addCommands(commands []string) {
	t.commands = append(t.commands, commands...)
}
