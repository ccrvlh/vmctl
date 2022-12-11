
## CONTAINERD
#
#
do_all_containerd() {
	local version="$1"
	local thinpool="$2"

	install_containerd "$version"
	write_containerd_config "$thinpool"
	start_containerd_service
}

# Fetch and install the containerd binary
install_containerd() {
	local tag="$1"

	say "Installing containerd version $tag to $INSTALL_PATH"

	if [[ "$version" == "$DEFAULT_VERSION" ]]; then
		tag=$(latest_release_tag "$CONTAINERD_REPO")
	fi

	bin=$(build_containerd_release_bin_name "$tag" "$ARCH")
	url=$(build_download_url "$CONTAINERD_REPO" "$tag" "$bin")
	install_release_tar "$url" "$(dirname $INSTALL_PATH)" || die "could not install containerd"

	"$CONTAINERD_BIN" --version &>/dev/null
	ok_or_die "Containerd version $tag not installed"

	say "Containerd version $tag successfully installed"
}

# Prepare the containerd state dirs
make_containerd_dirs() {
	local dirs=("$DEVMAPPER_DIR" "$CONTAINERD_STATE_DIR" "$(dirname $CONTAINERD_CONFIG_PATH)")
	for d in "${dirs[@]}"; do
		say "Creating containerd directory $d"

		mkdir -p "$d" || die "Failed to make containerd directory $d"
	done

	say "All containerd directories created"
}

# Write out the containerd config file
write_containerd_config() {
	local thinpool="$1"

	say "Writing containerd config to $CONTAINERD_CONFIG_PATH"

	cat <<EOF >"$CONTAINERD_CONFIG_PATH"
version = 2

root = "$CONTAINERD_ROOT_DIR"
state = "$CONTAINERD_STATE_DIR"

[grpc]
  address = "$CONTAINERD_STATE_DIR/containerd.sock"

[metrics]
  address = "127.0.0.1:1338"

[plugins]
  [plugins."io.containerd.snapshotter.v1.devmapper"]
    pool_name = "$thinpool-thinpool"
    root_path = "$DEVMAPPER_DIR"
    base_image_size = "10GB"
    discard_blocks = true

[debug]
  level = "trace"
EOF

	say "Containerd config saved"
}

# Start the containerd systemd service
start_containerd_service() {
	say "Starting containerd service with $CONTAINERD_SERVICE_FILE"

	service="$CONTAINERD_BIN.service"
	fetch_service_file "$CONTAINERD_REPO" "$service" "$CONTAINERD_SERVICE_FILE"

	sed -i "s|ExecStart=.*|& --config $CONTAINERD_CONFIG_PATH|" "$CONTAINERD_SERVICE_FILE"

	start_service "$CONTAINERD_SYSTEMD_SVC"

	say "Containerd running"
}
