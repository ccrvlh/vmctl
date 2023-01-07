package provision

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"vmctl/src/config"
	"vmctl/src/utils"
)

// func InitContainerd(config ProvisionOptions, cfg *config.AppConfig) {
// 	// do_all_containerd "$ctrd_version" "$set_thinpool"
// 	InstallContainerd(config, cfg)
// 	createContainerdDirectories(config, cfg)
// 	renderConfigTemplate(cfg)
// 	startContainerdService(config, cfg)
// }

func CheckContainerdInstalled(config *config.AppConfig) {

}

// Doers
func InstallContainerd(opts ProvisionOptions, config *config.AppConfig) {
	fmt.Printf("Installing containerd version %s to %s", config.Containerd.Version, config.InstallPath)

	// if [[ "$version" == "$DEFAULT_VERSION" ]]; then
	// 	tag=$(latest_release_tag "$CONTAINERD_REPO")
	// fi
	var tag = config.Containerd.Version
	if config.Containerd.Version == config.DefaultVersion {
		tag = BuildLatestReleaseURL(config.Containerd.Repo)
	}

	// bin=$(build_containerd_release_bin_name "$tag" "$ARCH")
	// url=$(build_download_url "$CONTAINERD_REPO" "$tag" "$bin")
	// install_release_tar "$url" "$(dirname $INSTALL_PATH)" || die "could not install containerd"
	var binFile = buildContainerdBinaryReleaseName(tag, config)
	var url = BuildBinaryDownloadURL(config.Containerd.Repo, tag, binFile)
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
		os.Mkdir(value, 0666)
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

// Runs `--version` on the containerd binary
// as a way to ensure the installation succeded
func _checkCointainerdInstallation(containerdBinary string) bool {
	var startCmd = fmt.Sprintf("%s --version", containerdBinary)
	var _, startErr = exec.Command("bash", "-c", startCmd).Output()
	return startErr == nil
}

// # Returns containerd release binary name for linux-amd64
// # If/when we need to support others, we can ammend
//
//	build_containerd_release_bin_name() {
//		local tag=${1//v/} # remove the 'v' from arg $1
//		local arch="$2"
//		echo "containerd-$tag-linux-$arch.tar.gz"
//	}
func buildContainerdBinaryReleaseName(tag string, cfg *config.AppConfig) string {
	var trimmedTag = strings.Replace(tag, "v", "", -1)
	return fmt.Sprintf("containerd-%s-tag-linux-%s.tar.gz", trimmedTag, cfg.Arch)
}

// # Sets various global variables for state paths
// build_containerd_paths() {
// 	local tag=""

// 	if [[ "$DEVELOPMENT" == "true" ]]; then
// 		tag="-dev"
// 	fi

// 	CONTAINERD_CONFIG_PATH="/etc/containerd/config$tag.toml"
// 	CONTAINERD_ROOT_DIR="/var/lib/containerd$tag"
// 	CONTAINERD_STATE_DIR="/run/containerd$tag"
// 	CONTAINERD_SERVICE_FILE="/etc/systemd/system/containerd$tag.service"
// 	CONTAINERD_SYSTEMD_SVC="containerd$tag.service"
// 	DEVMAPPER_DIR="$CONTAINERD_ROOT_DIR/snapshotter/devmapper"
// 	DEVPOOL_METADATA="$DEVMAPPER_DIR/metadata"
// 	DEVPOOL_DATA="$DEVMAPPER_DIR/data"
// }

type ContainerdPaths struct {
	CONTAINERD_CONFIG_PATH  string
	CONTAINERD_ROOT_DIR     string
	CONTAINERD_STATE_DIR    string
	CONTAINERD_SERVICE_FILE string
	CONTAINERD_SYSTEMD_SVC  string
	DEVMAPPER_DIR           string
	DEVPOOL_METADATA        string
	DEVPOOL_DATA            string
}

func BuildContainerdPaths(opts ProvisionOptions) ContainerdPaths {
	var tagSuffix = ""
	if opts.Development {
		tagSuffix = "-dev"
	}
	var containerdPaths = ContainerdPaths{
		CONTAINERD_CONFIG_PATH:  fmt.Sprintf("/etc/containerd/config%s", tagSuffix),
		CONTAINERD_ROOT_DIR:     fmt.Sprintf("/var/lib/containerd%stag", tagSuffix),
		CONTAINERD_STATE_DIR:    fmt.Sprintf("/run/containerd%stag", tagSuffix),
		CONTAINERD_SERVICE_FILE: fmt.Sprintf("/etc/systemd/system/containerd%stag.service", tagSuffix),
		CONTAINERD_SYSTEMD_SVC:  fmt.Sprintf("containerd%stag.service", tagSuffix),
		DEVMAPPER_DIR:           fmt.Sprintf("%s/snapshotter/devmapper", "CONTAINERD_ROOT_DIR"),
		DEVPOOL_METADATA:        fmt.Sprintf("%s/metadata", "DEVMAPPER_DIR"),
		DEVPOOL_DATA:            fmt.Sprintf("%s/data", "DEVMAPPER_DIR"),
	}
	return containerdPaths
}

// # Set and create the correct state dirs
// prepare_dirs() {
// 	build_containerd_paths
// 	make_containerd_dirs
// }
func PrepareContainerdDirectories() {}
