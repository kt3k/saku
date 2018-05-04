// +build windows

package main

import (
	"os"
	"os/exec"
	"strconv"
)

func execCommand(command string) error {
	cmd := exec.Command("cmd.exe", "/s", "/c", "\""+command+"\"")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "IN_SAKU=true")

	err := cmd.Run()

	return err
}

func terminateCommand(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return nil
	}

	pid := cmd.Process.Pid
	c := exec.Command("taskkill", "/t", "/f", "/pid", strconv.Itoa(pid))
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
