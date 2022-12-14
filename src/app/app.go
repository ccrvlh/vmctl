package app

import (
	"time"
	"vmctl/src/cli/provision"
	config "vmctl/src/config"

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
			provision.NewProvisionCommand(),
			provision.NewCheckCommand(),
		},
	}

	var app = MicroApp{
		Cli:    *cliApp,
		Config: *cfg,
	}

	return app
}
