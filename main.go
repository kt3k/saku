package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "saku"
	app.Usage = "Markdown-based task runner"

	app.Version = Version

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
