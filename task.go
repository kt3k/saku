package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

type task struct {
	title        string
	descriptions []string
	commands     []string
	options      taskOptions
	aborted      bool
	cmd          *exec.Cmd
}

func newTask() task {
	return task{
		title:        "",
		descriptions: []string{},
		commands:     []string{},
		options:      taskOptions{},
		aborted:      false,
		cmd:          nil,
	}
}

type taskOptions struct {
}

// Runs a task.
func (t *task) run(opts *runOptions, c chan error) {
	for _, command := range t.commands {
		if t.aborted {
			c <- nil
			return
		}

		fmt.Println("+" + command)
		t.cmd = execCommand(command)

		if t.cmd.Run() != nil {
			c <- errors.New("Task " + color.MagentaString(t.title) + " failed")
			return
		}
	}

	c <- nil
}

// Aborts a task.
func (t *task) abort() {
	if t.aborted {
		return
	}

	terminateCommand(t.cmd)

	t.aborted = true
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
