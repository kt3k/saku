package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
)

const defaultConfigFile = "saku.md"

var patternEmbedDirective = regexp.MustCompile(`(?ism)<!--\s*saku\s+start\s*-->(.*)<!--\s*saku\s+end\s*-->`)

// Reads task config from markdown files
func readConfig(configFile string) ([]byte, error) {
	absPath, _ := filepath.Abs(configFile)

	data, err := ioutil.ReadFile(absPath)

	if err == nil {
		if !invokedInSaku() {
			fmt.Println("Read", prependEmoji("🔎", color.MagentaString(absPath)))
		}
		return data, nil
	}

	if configFile != defaultConfigFile {
		return []byte{}, err
	}

	absPath, _ = filepath.Abs("README.md")
	data, err = ioutil.ReadFile(absPath)

	if err != nil {
		return []byte{}, err
	}

	if !patternEmbedDirective.Match(data) {
		return []byte{}, errors.New("No <!-- saku start --><!-- saku end --> directive found")
	}

	return patternEmbedDirective.FindSubmatch(data)[1], nil
}
