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
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all todo tasks for today",
				Action: func(c *cli.Context) error {
					router.Route("list", c.Args())
					return nil
				},
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Mark a todo as started",
				Action: func(c *cli.Context) error {
					router.Route("start", c.Args())
					return nil
				},
			},
			{
				Name:    "pause",
				Aliases: []string{"p"},
				Usage:   "Mark a todo as paused",
				Action: func(c *cli.Context) error {
					router.Route("pause", c.Args())
					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "Mark a todo as done",
				Action: func(c *cli.Context) error {
					router.Route("done", c.Args())
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
