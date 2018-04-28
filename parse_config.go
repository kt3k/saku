package main

import (
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
	"strings"
)

// Parses config markdown and returns tasks.
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

				tasks.addCurrentTaskDescription(description)

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

	fmt.Printf("tasks=%#v\n", tasks)

	return tasks
}
