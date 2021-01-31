package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "todo",
		Usage: "keep track of your todo items on the cli",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello world!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
