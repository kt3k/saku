package saku

import (
	"os/exec"
	"testing"
)

func TestTerminateWithNilt(t *testing.T) {
	err := terminateCommand(nil)

	if err != nil {
		t.Error("terminateCommand with nil should return nil")
	}
}

func TestExecWithNoProcessCommand(t *testing.T) {
	cmd := exec.Command("foo")

	err := terminateCommand(cmd)

	if err != nil {
		t.Error("terminateCommand with empty process cmd should return nil")
	}
}
