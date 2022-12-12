package app

import (
	"fmt"
	"time"
	"vmctl/src/config"
	"vmctl/src/modules/bootstrap"

	"github.com/urfave/cli/v2"
)

type MicroApp struct {
	Cli    cli.App
	Config config.AppConfig
}

// Builds a new MicroVM CTL Application
// It takes a previously lodade configuration (AppConfig)
// Returns a VMCtlApp struct with the CLI Interface and the Config
func NewApp(cfg *config.AppConfig) MicroApp {
	var cliApp = &cli.App{
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
				Action: bootstrap.BootstrapAction,
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

	var app = MicroApp{
		Cli:    *cliApp,
		Config: *cfg,
	}

	return app
}
