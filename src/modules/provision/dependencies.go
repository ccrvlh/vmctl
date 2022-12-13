package provision

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"vmctl/src/config"
)

// Init
func initDependencies(options ProvisionOptions, cfg *config.AppConfig) {
	var requiredPackages = []string{
		"thin-provisioning-tools",
		"lvm2",
		"git",
		"curl",
		"wget",
		"dmsetup",
		"bc",
		"qemu",
		"qemu-kvm",
		"libvirt-clients",
		"libvirt-daemon-system",
		"virtinst",
		"bridge-utils",
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
		log.Fatal("KVM is not installed, virtualization is needed to bootstrap.")
	}

	// Check whether all dependencies are installed.
	// It does *not* install anything
	fmt.Println("Checking for dependencies...")
	var hasMissingPackages = checkPackages(requiredPackages)
	if len(hasMissingPackages) != 0 {
		var msg = fmt.Sprintf("Packages %s are required, install it with `apt-get` before proceding", hasMissingPackages)
		log.Fatal(msg)
	}

	// Check whether Libvirt service is enabled and running.
	fmt.Println("Checking for libvirt...")
	var libvirtEnabled = checkLibvirt()
	if !libvirtEnabled {
		var msg = "Couldn't initialize Libvirt"
		log.Fatal(msg)
	}

}

// Helpers
func checkArch(config *config.AppConfig) bool {
	switch arch := runtime.GOARCH; arch {
	case "x86_64", "amd64":
		config.Arch = "amd64"
		return true
	case "aarch64", "arm64":
		config.Arch = "arm64"
		return true
	default:
		fmt.Printf("Architecture not supported.")
		return false
	}
}

func checkKVM() bool {
	if _, err := os.Stat("/dev/kvm"); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkPackages(requiredPackages []string) []string {
	var missingPackages []string
	for _, value := range requiredPackages {
		var arg = fmt.Sprintf("apt -qq list %s", value)
		var out, err = exec.Command("bash", "-c", arg).Output()
		if err != nil {
			log.Fatal(err)
		}
		if out == nil {
			missingPackages = append(missingPackages, value)
		}
	}
	return missingPackages
}

func checkLibvirt() bool {
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
