package main

type taskError struct {
	message string
}

func (e *taskError) Error() string {
	return e.message
}
