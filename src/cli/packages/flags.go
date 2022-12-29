package provision

import "github.com/urfave/cli/v2"

// Setup all flags for the bootstrap commands
func PackageFlags() []cli.Flag {
	var bootFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "dry-run",
			Value: false,
			Usage: "Dry run runs check to assess what actions should be taken",
		},
		&cli.BoolFlag{
			Name:  "check",
			Value: true,
			Usage: "Non interactive bootstrap",
		},
		&cli.BoolFlag{
			Name:    "update",
			Aliases: []string{"d", "dev"},
			Usage:   "Update current installed packages",
		},
	}
	return bootFlags
}
