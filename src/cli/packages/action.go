package provision

import (
	"fmt"
	"time"
	svc "vmctl/src/modules/packages"
	req "vmctl/src/modules/requirements"
	utils "vmctl/src/utils"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func PackageAction(cCtx *cli.Context) error {
	req.IsArchSupoprted()
	var opts = buildPackageOptions(cCtx)
	fmt.Println("VmCTL | Provisioning host...")
	var missingPackages, err = checkRequiredPackages(opts)
	if err != nil {
		fmt.Println("Couldn't check for packages")
		return err
	}

	// All packages are installed
	if len(missingPackages) == 0 {
		fmt.Println("All required packages are installed, ready to proceed")
		return nil
	}

	// Dry Run, won't install anything, but list missing packages
	if opts.DryRun {
		var missing = fmt.Sprintln("Some required packages were missing, this would install:")
		fmt.Println(missing)
		for _, pack := range missingPackages {
			var formattedMsg = fmt.Sprintf("○ %s", pack)
			fmt.Println(formattedMsg)
		}
		return nil
	}

	// Has missing packages, not dry run, install
	var installError = installPackages(missingPackages)
	if installError != nil {
		var missing = fmt.Sprintln("Some required packages were missing, this would install:")
		fmt.Println(missing)
	}
	return nil
}

func checkRequiredPackages(opts svc.PackagesOptions) ([]string, error) {
	var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("white")
	s.Suffix = " Checking if all required packages are installed..."
	s.Start()
	var _, packs = svc.CheckRequiredPackages()
	if len(packs) != 0 {
		s.Stop()
		s.FinalMSG = utils.ErrorMsg("Some required packages were not found")
		return packs, nil
	}
	s.Stop()
	s.FinalMSG = utils.SuccessMsg("All Packages already installed.")
	return packs, nil
}

func installPackages(packages []string) error {
	var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("white")
	s.Suffix = "○ Installing missing packages..."
	s.Start()
	var _, err = svc.InstallMissingPackages(packages)
	if err != nil {
		s.Stop()
		s.FinalMSG = utils.ErrorMsg("There was an error installing some packages")
	}
	return nil
}
