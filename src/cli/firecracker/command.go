package firecracker

import (
	"github.com/urfave/cli/v2"
)

func NewFirecrackerCommand() *cli.Command {
	var command = &cli.Command{
		Name:  "containerd",
		Usage: "Install, configure and start containerd service",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "dry-run",
				Value: false,
				Usage: "Dry run mode",
			},
			&cli.BoolFlag{
				Name:  "check",
				Value: true,
				Usage: "Checks if the needed packages are installed",
			},
			&cli.BoolFlag{
				Name:    "update",
				Aliases: []string{"d", "dev"},
				Usage:   "Update current installed packages",
			},
		},
		Action: FirecrackerAction,
	}
	return command
}

func FirecrackerAction(cCtx *cli.Context) error {
	return nil
}
