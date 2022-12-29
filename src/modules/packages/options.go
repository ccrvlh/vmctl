package packages

// The struct with all possible packaging options
type PackagesOptions struct {
	DryRun       bool // Set up development environment (uses Loop thinpools)
	ShouldUpdate bool // Whether to update packages
	CheckOnly    bool // Whether to only check for current status
}
