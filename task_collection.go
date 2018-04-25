package main

type TaskError struct {
	message string
}

func (e *TaskError) Error() string {
	return e.message
}

type TaskCollection struct {
	tasks []Task
}

func (tasks *TaskCollection) Run() error {
	println("tasks running")

	return &TaskError{message: "Method not implemented"}
}
