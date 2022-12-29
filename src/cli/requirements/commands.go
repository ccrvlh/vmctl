package requirements

import (
	"github.com/urfave/cli/v2"
)

func NewRequirementsCommand() *cli.Command {
	var command = &cli.Command{
		Name:   "requirements",
		Usage:  "checks for the system requirements",
		Action: RequirementsAction,
	}
	return command
}
