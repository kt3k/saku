package saku

import "testing"

func TestRunSingleCommandOfAbortedTask(t *testing.T) {
	task := newTask(0)
	l := &logger{enabled: true}
	ctx := &runContext{
		l:         l,
		extraArgs: []string{},
		mode:      RunModeSequence,
	}
	task.aborted = true

	err := task.runSingleCommand("non-existing-command", ctx)

	if err != nil {
		t.Error("aborted task's runSingleCommand should always pass")
	}
}
