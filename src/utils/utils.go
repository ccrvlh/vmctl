package utils

import (
	"fmt"
	"os/exec"
	"vmctl/src/config"
)

func StartService(serviceName string, config *config.AppConfig) bool {
	var startCmd = fmt.Sprintf("systemctl start %s", serviceName)
	var _, startErr = exec.Command(config.Shell, "-c", startCmd).Output()
	return startErr == nil
}

type ActionResult struct {
	Name    string
	Status  bool
	Details string
}
