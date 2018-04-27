package main

type TaskCollection struct {
	currentTask Task
	tasks       []Task
}

func NewTaskCollection() *TaskCollection {
	task := NewTask()
	tasks := []Task{task}

	return &TaskCollection{
		currentTask: task,
		tasks:       tasks,
	}
}

func (tc *TaskCollection) Run() error {
	println("tasks running")

	return &TaskError{message: "Method not implemented"}
}

func (tc *TaskCollection) NewTask() {
	tc.currentTask = NewTask()
	tc.tasks = append(tc.tasks, tc.currentTask)
}

func (tc *TaskCollection) SetCurrentTaskTitle(title string) {
	tc.currentTask.SetTitle(title)
}

func (tc *TaskCollection) AddCurrentTaskDescription(description string) {
	tc.currentTask.AddDescription(description)
}

func (tc *TaskCollection) AddCurrentTaskCommands(commands []string) {
	tc.currentTask.AddCommands(commands)
}
