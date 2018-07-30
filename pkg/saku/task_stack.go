package saku

// taskStack represents the stack of tasks.
// This is used in task execution.
type taskStack struct {
	tasks []*task
}

// newTaskStack returns a new empty stack.
func newTaskStack() *taskStack {
	return &taskStack{
		tasks: []*task{},
	}
}

func (ts *taskStack) appended(t *task) *taskStack {
	return &taskStack{
		tasks: append(ts.tasks, t),
	}
}

func (ts *taskStack) isEmpty() bool {
	return len(ts.tasks) == 0
}
