// +build !windows

package saku

import (
	"os"
	"os/exec"
	"syscall"
)

// Executes the given string as command
func execCommand(command string) *exec.Cmd {
	cmd := exec.Command("/bin/sh", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Setsid = true
	cmd.Env = append(os.Environ(), "IN_SAKU=true")

	return cmd
}

// Terminates the given command
func terminateCommand(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}

	group, err := os.FindProcess(-1 * cmd.Process.Pid)

	if err == nil {
		group.Signal(syscall.SIGTERM)
	}

	return err
}
