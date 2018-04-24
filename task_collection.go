package main

type TaskCollection struct {
	tasks []Task
}

func (tasks *TaskCollection) Run() {
	println("tasks running")
}
