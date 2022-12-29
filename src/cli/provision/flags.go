package provision

// import "github.com/urfave/cli/v2"

// // Setup all flags for the bootstrap commands
// func ProvisionFlags() []cli.Flag {
// 	var bootFlags = []cli.Flag{
// 		&cli.BoolFlag{
// 			Name:  "dry-run",
// 			Value: false,
// 			Usage: "Dry run runs check to assess what actions should be taken",
// 		},
// 		&cli.BoolFlag{
// 			Name:  "skip-all",
// 			Value: true,
// 			Usage: "Non interactive bootstrap",
// 		},
// 		&cli.BoolFlag{
// 			Name:    "development",
// 			Aliases: []string{"d", "dev"},
// 			Value:   true,
// 			Usage:   "Set up development environment (uses Loop thinpools)",
// 		},
// 		&cli.StringFlag{
// 			Name:  "thinpool",
// 			Value: "flintlock",
// 			Usage: "Name of thinpool (default: flintlock[-dev])",
// 		},
// 		&cli.StringFlag{
// 			Name:  "disk",
// 			Usage: "Disk to use for DirectLVM thinpool (ignored for dev)",
// 		},
// 		&cli.StringFlag{
// 			Name:    "grpc-endpoint",
// 			Aliases: []string{"b", "grpc"},
// 			Value:   "0.0.0.0:9090",
// 			Usage:   "Disk to use for DirectLVM thinpool (ignored for dev)",
// 		},
// 		&cli.StringFlag{
// 			Name:  "fc-version",
// 			Value: "latest",
// 			Usage: "Firecracker Version",
// 		},
// 		&cli.StringFlag{
// 			Name:  "fl-version",
// 			Value: "latest",
// 			Usage: "Flintlock Version",
// 		},
// 		&cli.StringFlag{
// 			Name:  "cd-version",
// 			Value: "latest",
// 			Usage: "Containerd Version",
// 		},
// 	}
// 	return bootFlags
// }
