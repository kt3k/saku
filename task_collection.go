package main

// Task collection model.
type taskCollection struct {
	currentTask task
	tasks       []task
}

// Creates a new task collection.
func newTaskCollection() *taskCollection {

	return &taskCollection{
		currentTask: newTask(), // This is a dummy task, and will be discarded when the first task is created
		tasks:       []task{},
	}
}

func (tc *taskCollection) run() error {
	println("tasks running")

	return &taskError{message: "Method not implemented"}
}

func (tc *taskCollection) newTask() {
	tc.currentTask = newTask()
	tc.tasks = append(tc.tasks, tc.currentTask)
}

func (tc *taskCollection) setCurrentTaskTitle(title string) {
	tc.currentTask.setTitle(title)
}

func (tc *taskCollection) addCurrentTaskDescription(description string) {
	tc.currentTask.addDescription(description)
}

func (tc *taskCollection) addCurrentTaskCommands(commands []string) {
	tc.currentTask.addCommands(commands)
}
