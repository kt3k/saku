package main

// TaskCollection is the model of the list of tasks.
type TaskCollection struct {
	currentTask *task
	tasks       []*task
	taskMap     map[string]*task
}

// Creates a new task collection.
func newTaskCollection() *TaskCollection {

	// This is a dummy task, and will be discarded when the first task is created
	t := newTask()

	return &TaskCollection{
		currentTask: &t,
		tasks:       []*task{},
		taskMap:     map[string]*task{},
	}
}

// Run runs the tasks.
func (tc *TaskCollection) Run(opts *runOptions) error {
	if opts.isParallel() {
		return tc.runParallel(opts)
	} else if opts.isRace() {
		return tc.runInRace(opts)
	}

	return tc.runSequentially(opts)
}

func (tc *TaskCollection) runSequentially(opts *runOptions) error {
	c := make(chan error)

	for _, t := range tc.tasks {
		go t.run(opts, c)

		err := <-c

		if err != nil {
			return err
		}
	}

	return nil
}

func (tc *TaskCollection) runParallel(opts *runOptions) error {
	c := make(chan error)

	for i := range tc.tasks {
		t := tc.tasks[i]
		go t.run(opts, c)
	}

	for range tc.tasks {
		err := <-c

		if err != nil {
			tc.abort()
			return err
		}
	}

	return nil
}

func (tc *TaskCollection) runInRace(opts *runOptions) error {
	c := make(chan error)

	for i := range tc.tasks {
		go tc.tasks[i].run(opts, c)
	}

	defer tc.abort()

	return <-c
}

func (tc *TaskCollection) abort() {
	for _, t := range tc.tasks {
		t.abort()
	}
}

func (tc *TaskCollection) newTask() {
	t := newTask()
	tc.tasks = append(tc.tasks, &t)
	tc.currentTask = &t
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
	tasks := []*task{}
	taskMap := map[string]*task{}

	for _, title := range titles {
		tasks = append(tasks, tc.taskMap[title])
		taskMap[title] = tc.taskMap[title]
	}

	return &TaskCollection{
		currentTask: tasks[len(tasks)-1],
		tasks:       tasks,
		taskMap:     taskMap,
	}
}

// Gets a task by the given title.
func (tc *TaskCollection) getByTitle(title string) (*task, bool) {
	t, ok := tc.taskMap[title]

	return t, ok
}
