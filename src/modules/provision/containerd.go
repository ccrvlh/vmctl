package provision

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
	"vmctl/src/config"
	"vmctl/src/utils"
)

func initContainerd(config ProvisionOptions, cfg *config.AppConfig) {
	// do_all_containerd "$ctrd_version" "$set_thinpool"
	installContainerd(config, cfg)
	createContainerdDirectories(config, cfg)
	renderConfigTemplate(cfg)
	startContainerdService(config, cfg)
}

// Doers
//
func installContainerd(opts ProvisionOptions, config *config.AppConfig) {
	fmt.Printf("Installing containerd version %s to %s", config.Containerd.Version, config.InstallPath)

	// if [[ "$version" == "$DEFAULT_VERSION" ]]; then
	// 	tag=$(latest_release_tag "$CONTAINERD_REPO")
	// fi
	var tag = config.Containerd.Version
	if config.Containerd.Version == config.DefaultVersion {
		tag = utils.GetLatestReleaseTag()
	}

	// bin=$(build_containerd_release_bin_name "$tag" "$ARCH")
	// url=$(build_download_url "$CONTAINERD_REPO" "$tag" "$bin")
	// install_release_tar "$url" "$(dirname $INSTALL_PATH)" || die "could not install containerd"
	var binFile = _buildReleaseName(tag, config.Arch)
	var url = _buildDownloadURL(config, tag, binFile)
	utils.DownloadToPath(url, config.InstallPath)

	// "$CONTAINERD_BIN" --version &>/dev/null
	// ok_or_die "Containerd version $tag not installed"
	// say "Containerd version $tag successfully installed"
	var installSucceded = _checkCointainerdInstallation(config.Containerd.Bin)
	if !installSucceded {
		log.Fatal("Couldn't install containerd.")
	}

	// End
	fmt.Printf("Containerd installed successfully!")
}

func createContainerdDirectories(opts ProvisionOptions, config *config.AppConfig) {
	// CONTAINERD_CONFIG_PATH="/etc/containerd/config$tag.toml"
	// CONTAINERD_ROOT_DIR="/var/lib/containerd$tag"
	// CONTAINERD_STATE_DIR="/run/containerd$tag"
	// CONTAINERD_SERVICE_FILE="/etc/systemd/system/containerd$tag.service"
	// CONTAINERD_SYSTEMD_SVC="containerd$tag.service"
	// DEVMAPPER_DIR="$CONTAINERD_ROOT_DIR/snapshotter/devmapper"
	// DEVPOOL_METADATA="$DEVMAPPER_DIR/metadata"
	// DEVPOOL_DATA="$DEVMAPPER_DIR/data"
	var allDirs = []string{}
	allDirs = append(allDirs, config.Containerd.ConfigPath)
	allDirs = append(allDirs, config.Containerd.RootDir)
	allDirs = append(allDirs, config.Containerd.StateDir)
	allDirs = append(allDirs, config.Containerd.ServiceFile)
	allDirs = append(allDirs, config.Containerd.SystemdSvc)

	for _, value := range allDirs {
		os.Mkdir(value, 0600)
	}
}

// Starts the service using Systemctl
// Maybe account for other init systems in the future
func startContainerdService(opts ProvisionOptions, config *config.AppConfig) {
	// say "Starting containerd service with $CONTAINERD_SERVICE_FILE"
	// service="$CONTAINERD_BIN.service"
	// fetch_service_file "$CONTAINERD_REPO" "$service" "$CONTAINERD_SERVICE_FILE"
	// sed -i "s|ExecStart=.*|& --config $CONTAINERD_CONFIG_PATH|" "$CONTAINERD_SERVICE_FILE"
	// start_service "$CONTAINERD_SYSTEMD_SVC"
	// say "Containerd running"
	fmt.Println("Starting Containerd service")
}

// Render's containerd's configuration template
// and saves the resulting file to the path
// specified on the App's configuration
func renderConfigTemplate(cfg *config.AppConfig) {
	vars := make(map[string]interface{})
	vars["RootDir"] = cfg.Containerd.RootDir
	vars["StateDir"] = cfg.Containerd.StateDir
	vars["MetricsEndpoint"] = cfg.Containerd.MetricsEndpoint
	vars["Thinpool"] = cfg.Thinpool.Default
	vars["DevmapperDir"] = cfg.Containerd.DevmapperDir
	vars["BaseImageSize"] = cfg.Containerd.BaseImageSize
	vars["LogLevel"] = cfg.Containerd.LogLevel
	tmpl, _ := template.ParseFiles("templates/containerd.config.tmpl")
	fullPath := buildContainerdConfigPath(cfg)
	file, _ := os.Create(fullPath)
	defer file.Close()
	tmpl.Execute(file, vars)
}

// Builds the Network File Path according
// to the chosen Installation Path
func buildContainerdConfigPath(cfg *config.AppConfig) string {
	var fullPath = fmt.Sprintf("%s/containerd.config", cfg.Containerd.ConfigPath)
	return fullPath
}

// Build the Release Name
func _buildReleaseName(tag string, arch string) string {
	return "tagarch"
}

// Builds a fully qualified Download URL
// for a given `tag` and desired `binFile`
// eg: `var downloadURL = _buildDownloadURL(config, "latest", "containerd_amd64")`
func _buildDownloadURL(config *config.AppConfig, tag string, binFile string) string {
	return "tagarch"
}

// Runs `--version` on the containerd binary
// as a way to ensure the installation succeded
func _checkCointainerdInstallation(containerdBinary string) bool {
	var startCmd = fmt.Sprintf("%s --version", containerdBinary)
	var _, startErr = exec.Command("bash", "-c", startCmd).Output()
	return startErr == nil
}
