package main

// Task collection model.
type TaskCollection struct {
	currentTask *task
	tasks       []task
	taskMap     map[string]*task
}

// Creates a new task collection.
func newTaskCollection() *TaskCollection {

	// This is a dummy task, and will be discarded when the first task is created
	t := newTask()

	return &TaskCollection{
		currentTask: &t,
		tasks:       []task{},
		taskMap:     map[string]*task{},
	}
}

func (tc *TaskCollection) Run(opts *runOptions) error {
	if opts.isParallel() {
		return tc.runParallel(opts)
	} else if opts.isRace() {
		return tc.runInRace(opts)
	}

	return tc.runSequentially(opts)
}

func (tc *TaskCollection) runSequentially(opts *runOptions) error {
	for _, t := range tc.tasks {
		err := t.run(opts)

		if err != nil {
			return err
		}
	}

	return nil
}

func (tc *TaskCollection) runParallel(opts *runOptions) error {
	// TODO: run tasks in parallel
	return nil
}

func (tc *TaskCollection) runInRace(opts *runOptions) error {
	// TODO: run tasks in race
	return nil
}

func (tc *TaskCollection) newTask() {
	tc.tasks = append(tc.tasks, newTask())
	tc.currentTask = &tc.tasks[len(tc.tasks)-1]
}

func (tc *TaskCollection) setCurrentTaskTitle(title string) {
	tc.currentTask.setTitle(title)
	tc.taskMap[title] = tc.currentTask
}

func (tc *TaskCollection) addCurrentTaskDescription(description string) {
	tc.currentTask.addDescription(description)
}

func (tc *TaskCollection) addCurrentTaskCommands(commands []string) {
	tc.currentTask.addCommands(commands)
}

func (tc *TaskCollection) filterByTitles(titles []string) *TaskCollection {
	tasks := []task{}
	taskMap := map[string]*task{}
	for _, title := range titles {
		tasks = append(tasks, *tc.taskMap[title])
		taskMap[title] = tc.taskMap[title]
	}
	return &TaskCollection{
		currentTask: &tasks[len(tasks)-1],
		tasks:       tasks,
		taskMap:     taskMap,
	}
}

// Gets a task by the given title.
func (tc *TaskCollection) getByTitle(title string) (*task, bool) {
	task, ok := tc.taskMap[title]

	return task, ok
}
