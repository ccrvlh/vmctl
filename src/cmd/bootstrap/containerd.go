package bootstrap

import (
	"fmt"
	"os"
	"text/template"
	"vmctl/src/config"
	"vmctl/src/utils"
)

func initContainerd(config BootstrapOptions, cfg *config.AppConfig) {
	// do_all_containerd "$ctrd_version" "$set_thinpool"
	installContainerd(config, cfg)
	createContainerdDirectories(config, cfg)
	writeContainerdConfig(config, cfg)
	startContainerdService(config, cfg)
}

// Doers
//
func installContainerd(opts BootstrapOptions, config *config.AppConfig) {
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
	fmt.Printf("Containerd installed successfully!")
}

func createContainerdDirectories(opts BootstrapOptions, config *config.AppConfig) {
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

func writeContainerdConfig(opts BootstrapOptions, config *config.AppConfig) {
	// 	local thinpool="$1"
	// 	say "Writing containerd config to $CONTAINERD_CONFIG_PATH"
	//		say "Containerd config saved"
	//	}
	var ctrConfig = _loadConfigTemplate(config)
	os.WriteFile("containerd.conf", []byte(ctrConfig), 0600)
}

func startContainerdService(opts BootstrapOptions, config *config.AppConfig) {
	// say "Starting containerd service with $CONTAINERD_SERVICE_FILE"
	// service="$CONTAINERD_BIN.service"
	// fetch_service_file "$CONTAINERD_REPO" "$service" "$CONTAINERD_SERVICE_FILE"
	// sed -i "s|ExecStart=.*|& --config $CONTAINERD_CONFIG_PATH|" "$CONTAINERD_SERVICE_FILE"
	// start_service "$CONTAINERD_SYSTEMD_SVC"
	// say "Containerd running"
	fmt.Println("Starting Containerd service")
}

// Helpers
//
func _loadConfigTemplate(config *config.AppConfig) string {
	var template = template.New("containerd.config")
	var _, err = template.Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	return "ok"
}

func _buildReleaseName(tag string, arch string) string {
	return "tagarch"
}

func _buildDownloadURL(config *config.AppConfig, tag string, binFile string) string {
	return "tagarch"
}
