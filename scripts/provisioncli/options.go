package provision

// import (
// 	svc "vmctl/src/modules/provision"

// 	"github.com/urfave/cli/v2"
// )

// // Takes the CLI Context and builds the Options object
// // This then can be used by every `init` function for the
// // different methods that will be ran on the bootstrap
// func buildProvisionOptions(cCtx *cli.Context) svc.ProvisionOptions {
// 	var dryRun = cCtx.Bool("dry-run")
// 	var skipAll = cCtx.Bool("skip-all")
// 	var develop = cCtx.Bool("development")
// 	var thinpool = cCtx.String("thinpool")
// 	var disk = cCtx.String("disk")
// 	var gRPCEndpoint = cCtx.String("grpc-endpoint")
// 	var FirecrackerVersion = cCtx.String("fc-version")
// 	var FlintlockVersion = cCtx.String("fl-version")
// 	var ContainerdVersion = cCtx.String("cd-version")

// 	var newOptions = svc.ProvisionOptions{
// 		SkipAll:            skipAll,
// 		Development:        develop,
// 		DryRun:             dryRun,
// 		ThinPool:           thinpool,
// 		Disk:               disk,
// 		GRPCEndpoint:       gRPCEndpoint,
// 		FirecrackerVersion: FirecrackerVersion,
// 		FlintlockVersion:   FlintlockVersion,
// 		ContainerdVersion:  ContainerdVersion,
// 	}

// 	return newOptions
// }
