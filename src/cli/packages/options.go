package packages

import (
	"github.com/urfave/cli/v2"
)

// The struct with all possible packaging options
type PackagesOptions struct {
	DryRun       bool // Set up development environment (uses Loop thinpools)
	ShouldUpdate bool // Whether to update packages
	CheckOnly    bool // Whether to only check for current status
}

// Takes the CLI Context and builds the Options object
// This then can be used by every `init` function for the
// different methods that will be ran on the bootstrap
func buildPackageOptions(cCtx *cli.Context) PackagesOptions {
	var dryRun = cCtx.Bool("dry-run")
	var shouldUpdate = cCtx.Bool("update")
	var checkOnly = cCtx.Bool("check")

	var newOptions = PackagesOptions{
		DryRun:       dryRun,
		ShouldUpdate: shouldUpdate,
		CheckOnly:    checkOnly,
	}

	return newOptions
}
