package provision

import (
	svc "vmctl/src/modules/packages"

	"github.com/urfave/cli/v2"
)

// Takes the CLI Context and builds the Options object
// This then can be used by every `init` function for the
// different methods that will be ran on the bootstrap
func buildPackageOptions(cCtx *cli.Context) svc.PackagesOptions {
	var dryRun = cCtx.Bool("dry-run")
	var shouldUpdate = cCtx.Bool("update")
	var checkOnly = cCtx.Bool("check")

	var newOptions = svc.PackagesOptions{
		DryRun:       dryRun,
		ShouldUpdate: shouldUpdate,
		CheckOnly:    checkOnly,
	}

	return newOptions
}
