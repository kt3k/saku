package main

import (
	"gopkg.in/russross/blackfriday.v2"
)

// Parses config markdown and returns tasks.
func parseConfig(config *[]byte) taskCollection {
	tasks := taskCollection{tasks: []task{}}

	node := blackfriday.New().Parse(*config).FirstChild

	for node != nil {
		if node.Type == blackfriday.Heading {
			/* Heading > Text */
			title := string(node.FirstChild.Literal)
			println("Heading=" + title)
		} else if node.Type == blackfriday.BlockQuote {
			/* BlockQuote > Paragraph > Text */
			p := node.FirstChild

			for p != nil {
				description := p.FirstChild.Literal
				println("BlockQuote=" + string(description))

				p = p.Next
			}
		} else if node.Type == blackfriday.CodeBlock {
			/* CodeBlock > Text */
			println("CodeBlock=" + string(node.Literal))
		} else {
			println(string(node.Literal))
			println(string(node.String()))
		}

		node = node.Next
	}

	return tasks
}
