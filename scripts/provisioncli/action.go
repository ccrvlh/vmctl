package provision

// import (
// 	"fmt"
// 	"time"
// 	"vmctl/src/config"
// 	svc "vmctl/src/modules/provision"

// 	"github.com/briandowns/spinner"
// 	"github.com/urfave/cli/v2"
// )

// // Works as a `dry-run` method for the provisioning command
// // checking the status of each dependency befor any action
// // is taken
// func CheckAction(cCtx *cli.Context) error {
// 	var _ = buildProvisionOptions(cCtx)

// 	fmt.Println("VmCTL | Checking host status...")

// 	checkSystemArch()
// 	checkKVM()
// 	checkRequiredPackages()

// 	return nil
// }

// // Function that actually bootstraps a Flintlock-enabled server
// // It will run all checks & steps necessary to get a Flintlock server running
// // This includes Containerd, Firecracker, etc.
// func ProvisionAction(cCtx *cli.Context) error {
// 	var opts = buildProvisionOptions(cCtx)

// 	fmt.Println("VmCTL | Provisioning host...")

// 	checkSystemArch()
// 	checkKVM()
// 	checkRequiredPackages()

// 	// Setting up Networking
// 	fmt.Println("Setting up network...")
// 	svc.SetupNetwork(opts, &config.Cfg)

// 	// Setting Up Containerd
// 	fmt.Println("Creating Containerd directories...")
// 	svc.InstallContainerd(opts, &config.Cfg)

// 	// // Setup Thin Pool
// 	// fmt.Println("Setup Disks...")
// 	// initDisks(opts, &config.Cfg)

// 	// // Setup Firecracker
// 	// // install_firecracker "$fc_version"
// 	// fmt.Println("Installing Firecracker...")
// 	// initFirecracker(opts, &config.Cfg)

// 	// // Installing Flintlock
// 	// fmt.Println("Installing Flintlock...")
// 	// initFlintlock(opts, &config.Cfg)

// 	return nil
// }

// func checkContainerd() error {
// 	var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
// 	s.Color("white")
// 	s.Suffix = " Checking if containerd is installed..."
// 	s.Start()
// 	// var _, packs = svc.CheckContainerdInstalled()
// 	// if len(packs) != 0 {
// 	// 	s.Stop()
// 	// 	s.FinalMSG = utils.ErrorMsg("Some required packages were not found")
// 	// 	return fmt.Errorf("KVM is needed to proceed")
// 	// } else {
// 	// 	s.Stop()
// 	// 	s.FinalMSG = utils.SuccessMsg("KVM check")
// 	// }
// 	return nil
// }
