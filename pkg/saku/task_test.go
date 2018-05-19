package saku

import "testing"

func TestRunSingleCommandOfAbortedTask(t *testing.T) {
	task := newTask()
	task.aborted = true

	err := task.runSingleCommand("non-existing-command", nil, make(chan string))

	if err != nil {
		t.Error("aborted task's runSingleCommand should always pass")
	}
}
