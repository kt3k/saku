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

	err := cmd.Run()

	return err
}
