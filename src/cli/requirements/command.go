package requirements

import (
	"fmt"
	"time"
	config "vmctl/src/config"
	utils "vmctl/src/utils"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func NewRequirementsCommand() *cli.Command {
	var command = &cli.Command{
		Name:   "requirements",
		Usage:  "checks for the system requirements",
		Action: RequirementsAction,
	}
	return command
}

func RequirementsAction(cCtx *cli.Context) error {
	fmt.Println("VmCTL | Checking for system requirements...")
	var sysErr = checkSystemArch()
	if sysErr != nil {
		return sysErr
	}
	var kvmErr = checkKVM()
	if kvmErr != nil {
		return kvmErr
	}
	return nil
}

func checkSystemArch() error {
	var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("white")
	s.Suffix = " Checking System Architecture..."
	s.Start()
	var _, archSupported = CheckArchitecture(&config.Cfg)
	if !archSupported {
		s.FinalMSG = utils.ErrorMsg("Current Architecture not supported")
		s.Stop()
		return fmt.Errorf("architecture not supported")
	}
	s.FinalMSG = utils.SuccessMsg("Architecture check")
	s.Stop()
	return nil
}

func checkKVM() error {
	var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("white")
	s.Suffix = " Checking if KVM is enabled..."
	s.Start()
	var kvmEnabled = CheckKVM()
	if !kvmEnabled {
		s.FinalMSG = utils.ErrorMsg("KVM not found")
		s.Stop()
		return fmt.Errorf("KVM is needed to proceed")
	}
	s.FinalMSG = utils.SuccessMsg("KVM check")
	s.Stop()
	return nil
}
