package bootstrap

func initFlintlock(config BootstrapOptions) {

	// local version="$1"
	// local address="${2%:*}"
	// local parent_iface="$3"
	// local bridge_name="$4"
	// local insecure="$5"
	// local config_file="$6"

	// install_flintlockd "$version"

	// if [[ -z "$parent_iface" ]]; then
	// 	parent_iface=$(lookup_interface)
	// fi
	// if [[ -z "$address" ]]; then
	// 	address=$(lookup_address "$parent_iface")
	// fi
	// write_flintlockd_config "$address" "$parent_iface" "$bridge_name" "$insecure" "$config_file"

	// start_flintlockd_service
	// say "Flintlockd running at $address:9090 via interface $parent_iface"
}

func installFlintlockd() {

	// # Fetch and install the flintlockd binary at the specified version
	// install_flintlockd() {
	// 	local tag="$1"

	// 	say "Installing flintlockd version $tag to $INSTALL_PATH"

	// 	if [[ "$tag" == "$DEFAULT_VERSION" ]]; then
	// 		tag=$(latest_release_tag "$FLINTLOCK_REPO")
	// 	fi

	// 	url=$(build_download_url "$FLINTLOCK_REPO" "$tag" "${FLINTLOCK_BIN}_${ARCH}")
	// 	install_release_bin "$url" "$FLINTLOCK_BIN"

	// 	"$FLINTLOCK_BIN" version &>/dev/null
	// 	ok_or_die "Flintlockd version $tag not installed"

	//		say "Flintlockd version $tag successfully installed"
	//	}
}

func WriteFlintlockConfig() {

	// # Write flintlock config to default location
	// write_flintlockd_config() {
	// 	local address="$1"
	// 	local parent_iface="$2"
	// 	local bridge_name="$3"
	// 	local insecure="$4"
	// 	local config_file="$5"

	// 	mkdir -p "$(dirname "$FLINTLOCKD_CONFIG_PATH")"

	// 	say "Writing flintlockd config to $FLINTLOCKD_CONFIG_PATH."

	// 	declare -A settings
	// 	settings["containerd-socket"]="$CONTAINERD_STATE_DIR/containerd.sock"
	// 	settings["grpc-endpoint"]="$address:9090"
	// 	settings["verbosity"]="9"
	// 	settings["insecure"]="$insecure"

	// 	if [[ -n "$bridge_name" ]]; then
	// 		settings["bridge-name"]="$bridge_name"
	// 	else
	// 		settings["parent-iface"]="$parent_iface"
	// 	fi

	// 	if [[ -n "$config_file" ]]; then
	// 		say "Merging provided flintlockd config file with auto-generated options"
	// 		while IFS= read -r line; do
	// 			if [[ $line != *":"* ]]; then
	// 				continue
	// 			fi
	// 			key=$(echo "$line" | awk 'BEGIN { FS = ":" } ; { print $1 }')
	// 			value=$(echo "$line" | awk 'BEGIN { FS = ":" } ; { print $2 }' | tr -d ' ')
	// 			settings[$key]="$value"
	// 		done <"$config_file"
	// 	fi

	// 	local content=''
	// 	for key in ${!settings[@]}; do
	// 		# note that there is a line-break in this string
	// 		# that is important to keep the settings file valid.
	// 		content+="${key}: ${settings[${key}]}
	// "
	// 	done

	// 	cat <<EOF >"$FLINTLOCKD_CONFIG_PATH"
	// ---
	// $content
	// EOF

	//		say "Flintlockd config saved"
	//	}
}

func StartFlintlockdService() {

	// # Fetch and start the flintlock systemd service
	// start_flintlockd_service() {
	// 	say "Starting flintlockd service with $FLINTLOCKD_SERVICE_FILE"

	// 	service=$(basename "$FLINTLOCKD_SERVICE_FILE")
	// 	fetch_service_file "$FLINTLOCK_REPO" "$service" "$FLINTLOCKD_SERVICE_FILE"

	// 	containerd_service=$(basename "$CONTAINERD_SERVICE_FILE")
	// 	sed -i "s|\(Requires=\)\(.*\)|\1$containerd_service|" "$FLINTLOCKD_SERVICE_FILE"

	// 	start_service "$FLINTLOCK_BIN"
	// }

}
