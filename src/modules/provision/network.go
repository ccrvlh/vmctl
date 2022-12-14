package provision

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
	"vmctl/src/config"
)

// Initializes dedicated Network to Flintlock
func SetupNetwork(opts ProvisionOptions, cfg *config.AppConfig) {
	fmt.Println("Setting up network")
	renderNetworkTemplate(opts, cfg)
	defineAndStartNetwork(cfg)
}

// Renders the Network Template
// Saves the file to disk
func renderNetworkTemplate(opts ProvisionOptions, cfg *config.AppConfig) {
	vars := make(map[string]interface{})
	vars["NetworkName"] = cfg.Flintlock.NetworkName
	vars["BridgeName"] = cfg.Flintlock.BridgeName
	vars["IpAddress"] = cfg.Flintlock.IpAddress
	vars["IpAddressStart"] = cfg.Flintlock.IpAddressStart
	vars["IpAddressEnd"] = cfg.Flintlock.IpAddressEnd
	tmpl, _ := template.ParseFiles("templates/flintlock-net.xml.tmpl")
	fullPath := buildNetworkFilePath(cfg)
	file, _ := os.Create(fullPath)
	defer file.Close()
	tmpl.Execute(file, vars)
}

// Defines & Starts Flintlock's Network
// using Virsh. This can be changes to use
// go-libvirtd package in the future instead of raw Shell commands
func defineAndStartNetwork(cfg *config.AppConfig) {
	var fullFilePath = buildNetworkFilePath(cfg)
	var defineCmd = fmt.Sprintf("sudo virsh net-define %s", fullFilePath)
	var _, defineErr = exec.Command("bash", "-c", defineCmd).Output()
	if defineErr != nil {
		log.Fatal("Couldn't define network. Check `virsh` status.")
	}
	var startCmd = fmt.Sprintf("sudo virsh net-start %s", cfg.Flintlock.NetworkName)
	var _, startErr = exec.Command("bash", "-c", startCmd).Output()
	if startErr != nil {
		log.Fatal("Couldn't start network. Check `virsh` status.")
	}
}

// Builds the Network File Path according
// to the chosen Installation Path
func buildNetworkFilePath(cfg *config.AppConfig) string {
	var fullPath = fmt.Sprintf("%s/flintlock-net.xml", cfg.InstallPath)
	return fullPath
}
