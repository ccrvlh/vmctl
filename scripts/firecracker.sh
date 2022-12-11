
## FIRECRACKER
#
#
# Fetch and install the firecracker binary
install_firecracker() {
	local tag="$1"

	say "Installing firecracker version $tag to $INSTALL_PATH"

	if [[ "$tag" == "$DEFAULT_VERSION" ]]; then
		tag=$(latest_release_tag "$FIRECRACKER_REPO")
	fi

	bin_name="${FIRECRACKER_BIN}_${ARCH}"

	# older firecracker macvtap releases had binaries with different names
	if is_older_version "$tag"; then
		bin_name="$FIRECRACKER_BIN"
	fi

	url=$(build_download_url "$FIRECRACKER_REPO" "$tag" "$bin_name")
	install_release_bin "$url" "$FIRECRACKER_BIN" || die "could not install firecracker"

	"$FIRECRACKER_BIN" --version &>/dev/null
	ok_or_die "firecracker version $tag not installed"

	say "Firecracker version $tag successfully installed"
}

is_older_version() {
	local tag="$1"
	local cutoff="v1.1.1-macvtap"

	[[ "$tag" == $(printf "%s\n%s" "$tag" "$cutoff" | sort -V | head -n 1) && "$tag" != "$cutoff" ]]
}
