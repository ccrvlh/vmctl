package bootstrap

import (
	"fmt"
	"log"
	"os/exec"
	"vmctl/src/config"
)

func initDependencies(options BootstrapOptions, cfg *config.AppConfig) {
	var requiredPackages = []string{
		"thin-provisioning-tools",
		"lvm2",
		"git",
		"curl",
		"wget",
		"dmsetup",
		"bc",
	}

	// Check whether KVM is Installed
	fmt.Println("Checking Architecture...")
	var archSupported = checkArch()
	if !archSupported {
		log.Fatal("Architecture not supported")
	}

	// Check whether KVM is Installed
	fmt.Println("Checking if KVM...")
	var isKVMInstalled = checkKVM()
	if !isKVMInstalled {
		log.Fatal("KVM is not installed, virtualization is need to bootstrap.")
	}

	// Check whether all dependencies are installed.
	// It does *not* install anything
	fmt.Println("Checking for dependencies...")
	var hasMissingPackages = checkPackages(requiredPackages)
	if len(hasMissingPackages) != 0 {
		var msg = fmt.Sprintf("Packages [%s] are required, install it with `apt-get` before proceding", hasMissingPackages)
		log.Fatal(msg)
	}

}

// ------------------------------------------------------------------

func checkArch() bool {
	var cmd = "uname -m"
	var out, err = exec.Command(cmd).Output()
	if err != nil {
		log.Fatal("Couldn't check for KVM")
	}
	var arch = string(out[:])
	if arch != "x86_64" {
		fmt.Printf("Architecture not supported.")
		return false
	}
	return true
}

func checkKVM() bool {
	var cmd = "[[ ! -d '/dev/kvm' ]] && echo 'KVM Not Found' || echo ''"
	var out, err = exec.Command(cmd).Output()
	if err != nil {
		log.Fatal("Couldn't check for KVM")
	}
	var exists = string(out[:])
	return exists != ""
}

func checkPackages(requiredPackages []string) []string {
	var missingPackages []string
	for _, value := range requiredPackages {
		var arg = fmt.Sprintf("apt -qq list %s", value)
		var out, err = exec.Command(arg).Output()
		if err != nil {
			log.Fatal(err)
		}
		if out == nil {
			missingPackages = append(missingPackages, value)
		}
	}
	return missingPackages
}
