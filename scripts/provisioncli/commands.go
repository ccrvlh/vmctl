package provision

// import (
// 	"github.com/urfave/cli/v2"
// )

// func NewProvisionCommand() *cli.Command {
// 	var command = &cli.Command{
// 		Name:   "provision",
// 		Usage:  "provisions a new server and dependencies",
// 		Flags:  ProvisionFlags(),
// 		Action: ProvisionAction,
// 		Subcommands: []*cli.Command{
// 			{
// 				Name:  "requirements",
// 				Usage: "checks for the system's specific requirements",
// 			},
// 			{
// 				Name:  "packages",
// 				Usage: "installs required packages",
// 			},
// 			{
// 				Name:  "firecracker",
// 				Usage: "provisions Firecracker and it's specific configurations",
// 			},
// 			{
// 				Name:  "devmapper",
// 				Usage: "provisions containerd's devmapper and it's specific configurations",
// 			},
// 			{
// 				Name:  "containerd",
// 				Usage: "provisions Containerd and it's specific configurations",
// 			},
// 			{
// 				Name:  "thinpool",
// 				Usage: "provisions the Thinpool and it's specific configurations",
// 			},
// 			{
// 				Name:  "network",
// 				Usage: "provisions the Network and it's specific configurations",
// 			},
// 		},
// 	}
// 	return command
// }

// func NewCheckCommand() *cli.Command {
// 	var command = &cli.Command{
// 		Name:  "check",
// 		Usage: "check if an existing setup is functional",
// 		// Flags:  ProvisionFlags(),
// 		Action: CheckAction,
// 	}
// 	return command
// }
