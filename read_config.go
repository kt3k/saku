package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/fatih/color"
)

// Reads task config from markdown files
func readConfig() ([]byte, error) {
	absPath, err := filepath.Abs("saku.md")

	if err != nil {
		return []byte{}, err
	}

	data, err := ioutil.ReadFile(absPath)

	if err != nil {
		// TODO: Tries reading from README.md, readme.md, README.markdown, readme.markdown
		return []byte{}, err
	}

	fmt.Println("Read", prependEmoji("🔎", color.MagentaString(absPath)))

	return data, nil
}
