package saku

import "testing"

func TestRunSingleCommandOfAbortedTask(t *testing.T) {
	task := newTask(0)
	l := &logger{enabled: true}
	task.aborted = true

	err := task.runSingleCommand("non-existing-command", nil, l)

	if err != nil {
		t.Error("aborted task's runSingleCommand should always pass")
	}
}
