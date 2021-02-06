package main

import (
	"log"
	"os"

	"github.com/ahsan/todo/router"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "todo",
		Usage:                "keep track of your todo items on the cli",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to today's list",
				Action: func(c *cli.Context) error {
					router.Route("add", c.Args())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
