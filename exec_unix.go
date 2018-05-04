// +build !windows

package main

import (
	"os"
	"syscall"
	"os/exec"
)

func execCommand(command string) error {
	cmd := exec.Command("/bin/sh", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Setsid = true
	cmd.Env = append(os.Environ(), "IN_SAKU=true")

	err := cmd.Run()

	return err
}

func terminateCommand(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return nil
	}

	group, err := os.FindProcess(-1 * cmd.Process.Pid)

	if err == nil {
		group.Signal(syscall.SIGKILL)
	}

	return err
}
