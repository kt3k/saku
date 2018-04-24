package main

type Task struct {
	name        string
	description string
	commands    []string
	options     TaskOptions
	aborted     bool
}

type TaskOptions struct {
}

// Runs a single command
func (t *Task) RunSingle() {
}

// Runs a task
func (t *Task) Run(opts *RunOptions) {
}

// Aborts a task
func (t *Task) Abort() {
	if t.aborted {
		return
	}
}
