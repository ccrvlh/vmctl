package provision

import (
	"fmt"
	"time"
	"vmctl/src/config"
	svc "vmctl/src/modules/provision"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

// Works as a `dry-run` method for the provisioning command
// checking the status of each dependency befor any action
// is taken
func CheckAction(cCtx *cli.Context) error {
	return nil
}

// Function that actually bootstraps a Flintlock-enabled server
// It will run all checks & steps necessary to get a Flintlock server running
// This includes Containerd, Firecracker, etc.
func ProvisionAction(cCtx *cli.Context) error {
	var opts = buildProvisionOptions(cCtx)

	fmt.Println("Provisioning host...")

	var s1 = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s1.Color("blue")
	s1.Suffix = " Checking package dependencies..."
	s1.Start()
	svc.CheckDependencies(opts, &config.Cfg)
	s1.Stop()

	var s = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Color("blue")
	s.Suffix = " Checking package dependencies..."
	s.Start()
	svc.CheckDependencies(opts, &config.Cfg)
	s.Stop()

	// Setting up Networking
	fmt.Println("Setting up network...")
	svc.SetupNetwork(opts, &config.Cfg)

	// Setting Up Containerd
	fmt.Println("Creating Containerd directories...")
	svc.InstallContainerd(opts, &config.Cfg)

	// // Setup Thin Pool
	// fmt.Println("Setup Disks...")
	// initDisks(opts, &config.Cfg)

	// // Setup Firecracker
	// // install_firecracker "$fc_version"
	// fmt.Println("Installing Firecracker...")
	// initFirecracker(opts, &config.Cfg)

	// // Installing Flintlock
	// fmt.Println("Installing Flintlock...")
	// initFlintlock(opts, &config.Cfg)

	return nil
}
