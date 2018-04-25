package main

type TaskCollection struct {
	tasks []Task
}

func (tasks *TaskCollection) Run() error {
	println("tasks running")

	return &TaskError{message: "Method not implemented"}
}
