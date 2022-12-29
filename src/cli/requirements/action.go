package requirements

import (
	"fmt"
	"time"
	config "vmctl/src/config"
	svc "vmctl/src/modules/requirements"
	utils "vmctl/src/utils"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

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
	var _, archSupported = svc.CheckArchitecture(&config.Cfg)
	if !archSupported {
		s.FinalMSG = utils.ErrorMsg("Current Architecture not supported")
		s.Stop()
		return fmt.Errorf("architecture not supported")
	} else {
		s.FinalMSG = utils.SuccessMsg("Architecture check")
		s.Stop()
	}
	return nil
}

func checkKVM() error {
	var s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("white")
	s.Suffix = " Checking if KVM is enabled..."
	s.Start()
	var kvmEnabled = svc.CheckKVM()
	if !kvmEnabled {
		s.FinalMSG = utils.ErrorMsg("KVM not found")
		s.Stop()
		return fmt.Errorf("KVM is needed to proceed")
	} else {
		s.FinalMSG = utils.SuccessMsg("KVM check")
		s.Stop()
	}
	return nil
}
