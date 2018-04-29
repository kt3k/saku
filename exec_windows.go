// +build windows

package main

import (
	"bytes"
	"os/exec"
)

func execCommand(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("cmd.exe", "/s", "/c", "\""+command+"\"")

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
