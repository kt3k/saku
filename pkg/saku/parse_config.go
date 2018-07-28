package saku

import (
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

// ParseConfig parses the given config markdown and returns tasks.
func ParseConfig(config *[]byte) *TaskCollection {
	currentTask := newTask(0)
	tasks := newTaskCollection()

	node := blackfriday.New().Parse(*config).FirstChild

	for node != nil {
		if node.Type == blackfriday.Heading {
			/* Heading > Text */
			title := string(node.FirstChild.Literal)

			currentTask = tasks.gotNewTask(node.Level, title)
		} else if node.Type == blackfriday.BlockQuote {
			/* BlockQuote > Paragraph */
			p := node.FirstChild

			for p != nil {
				/* Paragraph > Text */
				description := string(p.FirstChild.Literal)

				for _, desc := range strings.Split(description, "\n") {
					currentTask.addDescription(desc)
				}

				p = p.Next
			}
		} else if node.Type == blackfriday.CodeBlock {
			/* CodeBlock > Text */
			code := string(node.Literal)
			commands := strings.Split(code, "\n")

			for _, command := range commands {
				if strings.Trim(command, " \t\r") != "" {
					currentTask.addCommands([]string{command})
				}
			}
		}

		node = node.Next
	}

	return tasks
}
