package lvm

import (
	"github.com/urfave/cli/v2"
)

func NewLVMCommand() *cli.Command {
	var command = &cli.Command{
		Name:  "lvm",
		Usage: "Set up loop device thinpool (development environments)",
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
		Action: LVMAction,
	}
	return command
}

func LVMAction(cCtx *cli.Context) error {
	return nil
}
