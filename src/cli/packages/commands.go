package provision

import (
	"github.com/urfave/cli/v2"
)

func NewPackageCommand() *cli.Command {
	var command = &cli.Command{
		Name:   "packages",
		Usage:  "checks & installs for required system packages",
		Flags:  PackageFlags(),
		Action: PackageAction,
	}
	return command
}
