package containerd

import (
	"github.com/urfave/cli/v2"
)

func NewContainerdCommand() *cli.Command {
	var command = &cli.Command{
		Name:  "containerd",
		Usage: "Install, configure and start containerd service",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "version",
				Value: false,
				Usage: "Dry run mode",
			},
			&cli.StringFlag{
				Name:  "config",
				Usage: "Use custom containerd configuration",
			},
		},
		Action: ContainerdAction,
	}
	return command
}

func ContainerdAction(cCtx *cli.Context) error {
	return nil
}
