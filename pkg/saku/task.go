package saku

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

type task struct {
	title        string
	level        int
	descriptions []string
	commands     []string
	aborted      bool
	cmd          *exec.Cmd
	children     *TaskCollection
}

func newTask(level int) *task {
	return &task{
		title:        "",
		level:        level,
		descriptions: []string{},
		commands:     []string{},
		aborted:      false,
		cmd:          nil,
		children:     nil,
	}
}

// Runs a task.
func (t *task) run(ctx *runContext, c chan error, stack *taskStack) {
	for _, command := range t.commands {
		err := t.runSingleCommand(command, ctx)

		if err != nil {
			c <- err
			return
		}
	}

	if t.children == nil {
		c <- nil
		return
	}

	c <- t.children.Run(ctx, stack.appended(t))
}

// Runs a single command
func (t *task) runSingleCommand(command string, ctx *runContext) error {
	if t.aborted {
		return nil
	}

	if len(ctx.extraArgs) > 0 {
		command = command + " " + strings.Join(ctx.extraArgs, " ")
	}

	ctx.l.println("+" + command)
	t.cmd = execCommand(command)

	if t.cmd.Run() != nil {
		return errors.New("Task " + color.MagentaString(t.title) + " failed")
	}

	return nil
}

// Aborts a task.
func (t *task) abort() {
	if !t.aborted {
		terminateCommand(t.cmd)

		if t.children != nil {
			t.children.abort()
		}

		t.aborted = true
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

// findByTitle finds the task by the given title
func (t *task) findByTitle(title string) *task {
	if t.title == title {
		return t
	}

	if t.children == nil {
		return nil
	}

	return t.children.findByTitle(title)
}

// gotNewTask is the handler when it gets the new task of the given level and title. This is used during the parsing markdown.
func (t *task) gotNewTask(level int, title string) *task {
	if t.children == nil {
		t.children = newTaskCollection()
	}

	return t.children.gotNewTask(level, title)
}

func (t *task) taskCount() int {
	if t.children == nil {
		return 1
	}

	return t.children.taskCount() + 1
}

// setChildrenRunMode sets the run mode of the childre by the give mode.
func (t *task) setChildrenRunMode(mode RunMode) {
	if t.children == nil {
		t.children = newTaskCollection()
	}

	t.children.SetRunMode(mode)
}
