package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
)

// Reads task config from markdown files
func readConfig() ([]byte, error) {
	data, err := ioutil.ReadFile("saku.md")

	if err != nil {
		// TODO: Tries reading from README.md, readme.md, README.markdown, readme.markdown
		return []byte{}, err
	}

	fmt.Println("Read", color.MagentaString("saku.md"))

	return data, nil
}
