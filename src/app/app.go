package app

import (
	"time"
	"vmctl/src/config"
	"vmctl/src/modules/provision"

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
				Name:   "provision",
				Usage:  "provisions a new server and dependencies",
				Flags:  provision.ProvisionFlags(),
				Action: provision.ProvisionAction,
				Subcommands: []*cli.Command{
					{
						Name:  "firecracker",
						Usage: "provisions Firecracker and it's specific configurations",
					},
					{
						Name:  "devmapper",
						Usage: "provisions containerd's devmapper and it's specific configurations",
					},
					{
						Name:  "containerd",
						Usage: "provisions Containerd and it's specific configurations",
					},
					{
						Name:  "thinpool",
						Usage: "provisions the Thinpool and it's specific configurations",
					},
					{
						Name:  "network",
						Usage: "provisions the Network and it's specific configurations",
					},
				},
			},
			{
				Name:   "check",
				Usage:  "check if an existing setup is functional",
				Flags:  provision.ProvisionFlags(),
				Action: provision.ProvisionAction,
			},
		},
	}

	var app = MicroApp{
		Cli:    *cliApp,
		Config: *cfg,
	}

	return app
}
