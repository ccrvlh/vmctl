package packages

import (
	"fmt"
	"log"
	"os/exec"
)

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

func CheckRequiredPackages() (bool, []string) {
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
	return true, missingPackages
}

func InstallMissingPackages(missingPackages []string) (output []byte, err error) {
	var args = []string{"install", "-y"}
	for _, pack := range missingPackages {
		if pack == "" {
			return nil, fmt.Errorf("apt.Install: Invalid package with empty Name")
		}
		args = append(args, pack)
	}
	cmd := exec.Command("apt-get", args...)
	return cmd.CombinedOutput()
}
