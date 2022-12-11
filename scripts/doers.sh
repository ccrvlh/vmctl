
## DOER FUNCS
#
#
# Checks user input for valid architecture and sets the global value for pulling
# correct binaries.
set_arch() {
	# shellcheck disable=SC2155
	local arch=$(uname -m)

	case $arch in
	x86_64 | amd64)
		ARCH=amd64
		;;
	aarch64 | arm64)
		ARCH=arm64
		;;
	*)
		die "Unknown arch or arch not supported: $arch."
		;;
	esac
}

# Returns the tag associated with a "latest" release
latest_release_tag() {
	# shellcheck disable=SC2155
	local latest_url=$(build_release_url "$1")
	# shellcheck disable=SC2155
	local url=$(curl -sL "$latest_url" | awk -F'"' '/tag_name/ {printf $4}')
	echo "$url"
}

# Install the untarred binary attached to a release to /usr/local/bin
install_release_bin() {
	local download_url="$1"
	local dest_file="$2"
	wget -q "$download_url" -O "$INSTALL_PATH/$dest_file" || die "failed to download release for $dest_file"
	chmod +x "$INSTALL_PATH/$dest_file"
}

# Install and untar the tarred binary attached to a release to /usr/local/bin
install_release_tar() {
	local download_url="$1"
	local dest_path="$2"
	curl -sL "$download_url" | tar xz -C "$dest_path"
}

# Set and create the correct state dirs
prepare_dirs() {
	build_containerd_paths
	make_containerd_dirs
}

# Download the given service file from the given repo
fetch_service_file() {
	local repo="$1"
	local service="$2"
	local dest="$3"
	# shellcheck disable=SC2155
	local url=$(build_raw_url "$repo" "$service")
	curl -o "$dest" -sL "$url" || die "failed to download $service"
	chmod 0664 "$dest"
	systemctl daemon-reload
}

# Enable and start the given systemd service
start_service() {
	local service="$1"
	systemctl enable "$service" &>/dev/null || die "failed to enable $service service"
	systemctl start "$service" || die "failed to start $service service"
}
