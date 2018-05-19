package saku

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
)

const defaultConfigFile = "saku.md"

var patternEmbedDirective = regexp.MustCompile(`(?ism)<!--\s*saku\s+start\s*-->(.*)<!--\s*saku\s+end\s*-->`)

// Reads task config from markdown files
func readConfig(cwd string, configFile string, l *logger) ([]byte, error) {
	absPath := filepath.Join(cwd, configFile)

	data, err := ioutil.ReadFile(absPath)

	if err == nil {
		printReading(absPath, l)
		return data, nil
	}

	if configFile != defaultConfigFile {
		return []byte{}, err
	}

	absPath = filepath.Join(cwd, "README.md")
	data, err = ioutil.ReadFile(absPath)

	if err != nil {
		return []byte{}, err
	}

	if !patternEmbedDirective.Match(data) {
		return []byte{}, errors.New("No <!-- saku start --><!-- saku end --> directive found")
	}

	printReading(absPath, l)
	return patternEmbedDirective.FindSubmatch(data)[1], nil
}

func printReading(path string, l *logger) {
	if !invokedInSaku() {
		l.println("Read", prependEmoji("🔎", color.MagentaString(path), emojiEnabled()))
	}
}
