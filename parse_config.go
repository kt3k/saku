package main

import (
	"gopkg.in/russross/blackfriday.v2"
)

// Parses config markdown and returns tasks.
func ParseConfig(config *[]byte) TaskCollection {
	tasks := TaskCollection{tasks: []Task{}}

	node := blackfriday.New().Parse(*config).FirstChild

	for node != nil {
		if node.Type == blackfriday.Heading {
			/* Heading > Text */
			println("Heading=" + string(node.FirstChild.Literal))
		} else if node.Type == blackfriday.BlockQuote {
			/* BlockQuote > Paragraph > Text */
			println("BlockQuote=" + string(node.FirstChild.FirstChild.Literal))
		} else if node.Type == blackfriday.CodeBlock {
			/* CodeBlock > Text */
			println("CodeBlock=" + string(node.Literal))
		}

		node = node.Next
	}

	return tasks
}
