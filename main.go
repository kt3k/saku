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

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "i, info",
			Usage: "show the task information and exits",
		},

		cli.BoolFlag{
			Name:  "p, parallel",
			Usage: "run tasks in parallel",
		},

		cli.BoolFlag{
			Name:  "r, race",
			Usage: "set the flag to kill all tasks when a task finished with zero. This option is valid only with 'parallel' option",
		},

		cli.BoolFlag{
			Name:  "s, serial",
			Usage: "runs tasks sequentially",
		},

		cli.BoolFlag{
			Name:  "q, quiet",
			Usage: "stop logging",
		},

		cli.StringFlag{
			Name: "c, config",
			Usage: "Specifies the config file. Default is 'saku.md'.",
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
