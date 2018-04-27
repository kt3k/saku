package main

type Task struct {
	title       string
	description string
	commands    []string
	options     TaskOptions
	aborted     bool
}

func NewTask() *Task {
	return &Task{
		title:       "",
		description: "",
		commands:    []string{},
		options:     TaskOptions{},
		aborted:     false,
	}
}

type TaskOptions struct {
}

// Runs a single command
func (*Task) RunSingle(command string) {
}

// Runs a task
func (t *Task) Run(opts *RunOptions) {
	// TODO: runs all commands
}

// Aborts a task
func (t *Task) Abort() {
	if t.aborted {
		return
	}
}

// Adds the description.
func (t *Task) AddDescription(description string) {
	t.description = t.description + description
}

// Sets the title.
func (t *Task) SetTitle(title string) {
	t.title = title
}

// Adds the code.
func (t *Task) AddCommands(commands []string) {
	t.commands = append(t.commands, commands...)
}
