package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"vmctl/src/cmd/bootstrap"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "vmctl",
		Usage:     "Interface to interact with MicroVMs through Fintlock",
		Version:   "v0.0.1",
		Compiled:  time.Now(),
		Copyright: "MIT",
		HelpName:  "MicroVM Control",
		Commands: []*cli.Command{
			{
				Name:   "bootstrap",
				Usage:  "bootstrap a new server and dependencies",
				Flags:  bootstrap.BootstrapFlags(),
				Action: bootstrap.BootstrapAll,
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
