package main

import (
	"gopkg.in/russross/blackfriday.v2"
	"strings"
)

// ParseConfig parses the given config markdown and returns tasks.
func ParseConfig(config *[]byte) *TaskCollection {
	tasks := newTaskCollection()

	node := blackfriday.New().Parse(*config).FirstChild

	for node != nil {
		if node.Type == blackfriday.Heading {
			/* Heading > Text */
			title := string(node.FirstChild.Literal)

			tasks.newTask()
			tasks.setCurrentTaskTitle(title)
		} else if node.Type == blackfriday.BlockQuote {
			/* BlockQuote > Paragraph */
			p := node.FirstChild

			for p != nil {
				/* Paragraph > Text */
				description := string(p.FirstChild.Literal)

				for _, desc := range strings.Split(description, "\n") {
					tasks.addCurrentTaskDescription(desc)
				}

				p = p.Next
			}
		} else if node.Type == blackfriday.CodeBlock {
			/* CodeBlock > Text */
			code := string(node.Literal)
			commands := strings.Split(code, "\n")

			for _, command := range commands {
				if strings.Trim(command, " \t\r") != "" {
					tasks.addCurrentTaskCommands([]string{command})
				}
			}
		}

		node = node.Next
	}

	return tasks
}
