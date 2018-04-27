package main

type task struct {
	title       string
	description string
	commands    []string
	options     taskOptions
	aborted     bool
}

func newTask() task {
	return task{
		title:       "",
		description: "",
		commands:    []string{},
		options:     taskOptions{},
		aborted:     false,
	}
}

type taskOptions struct {
}

// Runs a single command
func (*task) runSingle(command string) {
}

// Runs a task
func (t *task) Run(opts *runOptions) {
	// TODO: runs all commands
}

// Aborts a task
func (t *task) abort() {
	if t.aborted {
		return
	}
}

// Adds the description.
func (t *task) addDescription(description string) {
	t.description = t.description + description
}

// Sets the title.
func (t *task) setTitle(title string) {
	t.title = title
}

// Adds the code.
func (t *task) addCommands(commands []string) {
	t.commands = append(t.commands, commands...)
}
