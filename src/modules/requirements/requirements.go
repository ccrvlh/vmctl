package provision

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"vmctl/src/config"
)

// func CheckDependencies(options ProvisionOptions, cfg *config.AppConfig) {

// 	var results = []utils.ActionResult{}

// 	// Check whether KVM is Installed

// 	// Check whether KVM is Installed
// 	fmt.Println("Checking if KVM...")
// 	var kvmResult = utils.ActionResult{Name: "KVM", Status: true, Details: ""}
// 	var isKVMInstalled = CheckKVM()
// 	if !isKVMInstalled {
// 		kvmResult.Status = false
// 		kvmResult.Details = "KVM not installed, virtualization is needed to bootstrap."
// 		results = append(results, kvmResult)
// 	} else {
// 		kvmResult.Status = false
// 		archOk := utils.ActionResult{Name: "KVM", Status: true, Details: ""}
// 		results = append(results, archOk)
// 	}

// 	// Check whether Libvirt service is enabled and running.
// 	fmt.Println("Checking for libvirt...")
// 	var libvirtEnabled = checkLibvirt()
// 	if !libvirtEnabled {
// 		var msg = "Couldn't initialize Libvirt"
// 		log.Fatal(msg)
// 	}
// }

// Helpers
func IsArchSupoprted() error {
	var _, archSupported = CheckArchitecture(&config.Cfg)
	if !archSupported {
		return fmt.Errorf("current architecture not supported")
	}
	return nil
}

func CheckArchitecture(config *config.AppConfig) (string, bool) {
	switch arch := runtime.GOARCH; arch {
	case "x86_64", "amd64":
		config.Arch = "amd64"
		return arch, true
	case "aarch64", "arm64":
		config.Arch = "arm64"
		return arch, true
	default:
		fmt.Printf("Architecture not supported.")
		return arch, false
	}
}

func CheckKVM() bool {
	if _, err := os.Stat("/dev/kvm"); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckLibvirt() bool {
	var cmd = "systemctl check libvirtd"
	var _, err = exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			log.Default().Printf("Libvirtd not running. Startit with: %v\n", exitErr)
			return false
		} else {
			log.Default().Printf("failed to check libvirtd: %v", err)
			os.Exit(1)
			return false
		}
	}
	return true
}
