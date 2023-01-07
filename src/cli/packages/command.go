package packages

import (
	"fmt"
	"time"
	req "vmctl/src/cli/requirements"
	utils "vmctl/src/utils"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func NewPackageCommand() *cli.Command {
	var command = &cli.Command{
		Name:  "packages",
		Usage: "checks & installs for required system packages",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "dry-run",
				Value: false,
				Usage: "Dry run mode",
			},
			&cli.BoolFlag{
				Name:  "check",
				Value: true,
				Usage: "Checks if the needed packages are installed",
			},
			&cli.BoolFlag{
				Name:    "update",
				Aliases: []string{"d", "dev"},
				Usage:   "Update current installed packages",
			},
		},
		Action: PackageAction,
	}
	return command
}

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

func checkRequiredPackages(opts PackagesOptions) ([]string, error) {
	var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("white")
	s.Suffix = " Checking if all required packages are installed..."
	s.Start()
	var _, packs = CheckRequiredPackages()
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
	var _, err = InstallMissingPackages(packages)
	if err != nil {
		s.Stop()
		s.FinalMSG = utils.ErrorMsg("There was an error installing some packages")
	}
	return nil
}
