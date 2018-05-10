package main

import "os"

func main() {
	if len(os.Args) < 2 {
		panic("need extra arguments")
	}

	println("ok")
}
