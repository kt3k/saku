package main

import (
	"gopkg.in/russross/blackfriday.v2"
)

// Parses config markdown and returns tasks.
func ParseConfig(config *[]byte) TaskCollection {
	tasks := TaskCollection{tasks: []Task{}}

	md := blackfriday.New()
	ast := md.Parse(*config)

	//ast.Walk(visitor)
	ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		println(node.String())

		return blackfriday.GoToNext
	})

	return tasks
}

func visitor(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	println(node.String())
	return blackfriday.GoToNext
}
