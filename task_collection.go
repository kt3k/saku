package main

// Task collection model.
type TaskCollection struct {
	currentTask *task
	tasks       []task
}

// Creates a new task collection.
func newTaskCollection() *TaskCollection {

	t := newTask()

	return &TaskCollection{
		currentTask: &t, // This is a dummy task, and will be discarded when the first task is created
		tasks:       []task{},
	}
}

func (tc *TaskCollection) Run() error {
	println("tasks running")

	return &taskError{message: "Method not implemented"}
}

func (tc *TaskCollection) newTask() {
	tc.tasks = append(tc.tasks, newTask())
	tc.currentTask = &tc.tasks[len(tc.tasks)-1]
}

func (tc *TaskCollection) setCurrentTaskTitle(title string) {
	tc.currentTask.setTitle(title)
}

func (tc *TaskCollection) addCurrentTaskDescription(description string) {
	tc.currentTask.addDescription(description)
}

func (tc *TaskCollection) addCurrentTaskCommands(commands []string) {
	tc.currentTask.addCommands(commands)
}
