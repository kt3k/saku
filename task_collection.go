package main

// Task collection model.
type taskCollection struct {
	currentTask *task
	tasks       []task
}

// Creates a new task collection.
func newTaskCollection() *taskCollection {

	t := newTask()

	return &taskCollection{
		currentTask: &t, // This is a dummy task, and will be discarded when the first task is created
		tasks:       []task{},
	}
}

func (tc *taskCollection) Run() error {
	println("tasks running")

	return &taskError{message: "Method not implemented"}
}

func (tc *taskCollection) newTask() {
	tc.tasks = append(tc.tasks, newTask())
	tc.currentTask = &tc.tasks[len(tc.tasks)-1]
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
