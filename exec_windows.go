// +build windows

package main

import (
	"os"
	"os/exec"
)

func execCommand(command string) error {
	cmd := exec.Command("cmd.exe", "/s", "/c", "\""+command+"\"")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "IN_SAKU=true")

	err := cmd.Run()

	return err
}
