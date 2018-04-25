package main

type TaskError struct {
	message string
}

func (e *TaskError) Error() string {
	return e.message
}
