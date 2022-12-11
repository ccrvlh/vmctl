
## BUILDER FUNCS
#
#
# Returns URL to latest release
build_release_url() {
	local repo_name="$1"
	echo "https://api.github.com/repos/$repo_name/releases/latest"
}

# Returns containerd release binary name for linux-amd64
# If/when we need to support others, we can ammend
build_containerd_release_bin_name() {
	local tag=${1//v/} # remove the 'v' from arg $1
	local arch="$2"

	echo "containerd-$tag-linux-$arch.tar.gz"
}

# Returns the desired binary download url for a repo, tag and binary
build_download_url() {
	local repo_name="$1"
	local tag="$2"
	local bin="$3"
	echo "https://github.com/$repo_name/releases/download/$tag/$bin"
}

# Returns the URL to a raw github file
build_raw_url() {
	local repo_name="$1"
	local file_name="$2"
	echo "https://raw.githubusercontent.com/$repo_name/$DEFAULT_BRANCH/$file_name"
}

# Sets various global variables for state paths
build_containerd_paths() {
	local tag=""

	if [[ "$DEVELOPMENT" == "true" ]]; then
		tag="-dev"
	fi

	CONTAINERD_CONFIG_PATH="/etc/containerd/config$tag.toml"
	CONTAINERD_ROOT_DIR="/var/lib/containerd$tag"
	CONTAINERD_STATE_DIR="/run/containerd$tag"
	CONTAINERD_SERVICE_FILE="/etc/systemd/system/containerd$tag.service"
	CONTAINERD_SYSTEMD_SVC="containerd$tag.service"
	DEVMAPPER_DIR="$CONTAINERD_ROOT_DIR/snapshotter/devmapper"
	DEVPOOL_METADATA="$DEVMAPPER_DIR/metadata"
	DEVPOOL_DATA="$DEVMAPPER_DIR/data"
}
