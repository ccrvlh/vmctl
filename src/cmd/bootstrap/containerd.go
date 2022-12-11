package bootstrap

import (
	"vmctl/src/config"
	cfg "vmctl/src/config"
)

func initContainerd(config BootstrapOptions, cfg *config.AppConfig) {
	// do_all_containerd "$ctrd_version" "$set_thinpool"
	installContainerd(config, cfg)
	createContainerdDirectories()
	writeContainerdConfig()
	startContainerdService()
}

func installContainerd(opts BootstrapOptions, config *cfg.AppConfig) {
	// fmt.Printf("Installing containerd")
	// var containerdRepo = "https://github.com"
	// local tag="$1"
	// say "Installing containerd version $tag to $INSTALL_PATH"

	// if [[ "$version" == "$DEFAULT_VERSION" ]]; then
	// 	tag=$(latest_release_tag "$CONTAINERD_REPO")
	// fi

	// bin=$(build_containerd_release_bin_name "$tag" "$ARCH")
	// url=$(build_download_url "$CONTAINERD_REPO" "$tag" "$bin")
	// install_release_tar "$url" "$(dirname $INSTALL_PATH)" || die "could not install containerd"
	// utils.DownloadToPath()

	// "$CONTAINERD_BIN" --version &>/dev/null
	// ok_or_die "Containerd version $tag not installed"

	// say "Containerd version $tag successfully installed"
}

func _buildReleaseBinName() {}

func createContainerdDirectories() {

	// CONTAINERD_CONFIG_PATH="/etc/containerd/config$tag.toml"
	// CONTAINERD_ROOT_DIR="/var/lib/containerd$tag"
	// CONTAINERD_STATE_DIR="/run/containerd$tag"
	// CONTAINERD_SERVICE_FILE="/etc/systemd/system/containerd$tag.service"

	// CONTAINERD_SYSTEMD_SVC="containerd$tag.service"

	// DEVMAPPER_DIR="$CONTAINERD_ROOT_DIR/snapshotter/devmapper"
	// DEVPOOL_METADATA="$DEVMAPPER_DIR/metadata"
	// DEVPOOL_DATA="$DEVMAPPER_DIR/data"
}

func writeContainerdConfig() {
	// 	local thinpool="$1"

	// 	say "Writing containerd config to $CONTAINERD_CONFIG_PATH"

	// 	cat <<EOF >"$CONTAINERD_CONFIG_PATH"
	// version = 2

	// root = "$CONTAINERD_ROOT_DIR"
	// state = "$CONTAINERD_STATE_DIR"

	// [grpc]
	//   address = "$CONTAINERD_STATE_DIR/containerd.sock"

	// [metrics]
	//   address = "127.0.0.1:1338"

	// [plugins]
	//   [plugins."io.containerd.snapshotter.v1.devmapper"]
	//     pool_name = "$thinpool-thinpool"
	//     root_path = "$DEVMAPPER_DIR"
	//     base_image_size = "10GB"
	//     discard_blocks = true

	// [debug]
	//   level = "trace"
	// EOF

	//		say "Containerd config saved"
	//	}
}
func startContainerdService() {
	// say "Starting containerd service with $CONTAINERD_SERVICE_FILE"

	// service="$CONTAINERD_BIN.service"
	// fetch_service_file "$CONTAINERD_REPO" "$service" "$CONTAINERD_SERVICE_FILE"

	// sed -i "s|ExecStart=.*|& --config $CONTAINERD_CONFIG_PATH|" "$CONTAINERD_SERVICE_FILE"

	// start_service "$CONTAINERD_SYSTEMD_SVC"

	// say "Containerd running"
}
