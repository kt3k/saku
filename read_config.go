package main

import (
	"fmt"
	"io/ioutil"
)

// Reads task config from markdown files
func ReadConfig () (string, error) {
	data, err := ioutil.ReadFile("saku.md")

	if err != nil {
		// TODO: Tries reading from README.md, readme.md, README.markdown, readme.markdown
		return "", err
	}

	fmt.Println("Read: saku.md")

	return string(data), nil
}
