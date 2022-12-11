
## COMMANDS
#
#
cmd_all() {
	local skip_apt=false
	local disk=""
	local fl_address=""
	local fl_iface=""
	local bridge_name=""
	local insecure=false
	local thinpool="$DEFAULT_THINPOOL"
	local fc_version="$FIRECRACKER_VERSION"
	local fl_version="$FLINTLOCK_VERSION"
	local ctrd_version="$CONTAINERD_VERSION"
	local flintlock_config_file=""

	while [ $# -gt 0 ]; do
		case "$1" in
		"-h" | "--help")
			cmd_all_help
			exit 1
			;;
		"-y" | "--unattended")
			OPT_UNATTENDED=true
			;;
		"-d" | "--disk")
			shift
			disk="$1"
			;;
		"-t" | "--thinpool")
			shift
			thinpool="$1"
			;;
		"-a" | "--grpc-address")
			shift
			fl_address="$1"
			;;
		"-i" | "--parent-iface")
			shift
			fl_iface="$1"
			;;
		"-b" | "--bridge")
			shift
			bridge_name="$1"
			;;
		"-s" | "--skip-apt")
			skip_apt=true
			;;
		"-k" | "--insecure")
			insecure=true
			;;
		"--dev")
			DEVELOPMENT=true
			;;
		"-f" | "--flintlock-config-file")
			shift
			flintlock_config_file="$1"
			;;
		*)
			die "Unknown argument: $1. Please use --help for help."
			;;
		esac
		shift
	done

	say "$(date -u +'%F %H:%M:%S %Z'): Provisioning host $(hostname)"
	say "The following subcommands will be performed:" \
		"apt, firecracker, containerd, flintlock, direct_lvm|devpool"

	set_arch
	say "Will install binaries for architecture: $ARCH"

	ensure_kvm

	if [[ "$skip_apt" == false ]]; then
		cmd_apt
	fi

	prepare_dirs

	# if the env is a dev one, then we don't want to use a real disk
	# and we want to tag all state dirs with 'dev'
	if [[ "$DEVELOPMENT" == false ]]; then
		set_thinpool="${DEFAULT_THINPOOL:=$thinpool}"
		do_all_direct_lvm "$disk" "$set_thinpool"
	else
		set_thinpool="${DEFAULT_DEV_THINPOOL:=$thinpool}"
		do_all_devpool "$set_thinpool"
	fi

	install_firecracker "$fc_version"
	do_all_containerd "$ctrd_version" "$set_thinpool"
	do_all_flintlock "$fl_version" "$fl_address" "$fl_iface" "$bridge_name" "$insecure" "$flintlock_config_file"

	say "$(date -u +'%F %H:%M:%S %Z'): Host $(hostname) provisioned"
}

cmd_apt() {
	say "Installing required apt packages"
	apt update
	apt install -qq -y \
		thin-provisioning-tools \
		lvm2 \
		git \
		curl \
		wget || die "failed to install apt packages"
	say "Packages installed"
}

cmd_firecracker() {
	local version="$FIRECRACKER_VERSION"

	while [ $# -gt 0 ]; do
		case "$1" in
		"-h" | "--help")
			cmd_firecracker_help
			exit 1
			;;
		"-v" | "--version")
			shift
			version="$1"
			;;
		*)
			die "Unknown argument: $1. Please use --help for help."
			;;
		esac
		shift
	done

	set_arch
	install_firecracker "$version"
}

cmd_containerd() {
	local version="$CONTAINERD_VERSION"
	local thinpool="$DEFAULT_THINPOOL"

	while [ $# -gt 0 ]; do
		case "$1" in
		"-h" | "--help")
			cmd_containerd_help
			exit 1
			;;
		"-v" | "--version")
			shift
			version="$1"
			;;
		"-t" | "--thinpool")
			shift
			thinpool="$1"
			;;
		"--dev")
			DEVELOPMENT=true
			thinpool="$DEFAULT_DEV_THINPOOL"
			;;
		*)
			die "Unknown argument: $1. Please use --help for help."
			;;
		esac
		shift
	done

	set_arch
	prepare_dirs
	do_all_containerd "$version" "$thinpool"
}

cmd_flintlock() {
	local version="$FLINTLOCK_VERSION"
	local address=""
	local parent_iface=""
	local bridge_name=""
	local insecure=false
	local config_file=""

	while [ $# -gt 0 ]; do
		case "$1" in
		"-h" | "--help")
			cmd_flintlock_help
			exit 1
			;;
		"-v" | "--version")
			shift
			version="$1"
			;;
		"-a" | "--grpc-address")
			shift
			address="$1"
			;;
		"-i" | "--parent-iface")
			shift
			parent_iface="$1"
			;;
		"-b" | "--bridge")
			shift
			bridge_name="$1"
			;;
		"-k" | "--insecure")
			insecure=true
			;;
		"-f" | "--config-file")
			shift
			config_file="$1"
			;;
		"--dev")
			DEVELOPMENT=true
			;;
		*)
			die "Unknown argument: $1. Please use --help for help."
			;;
		esac
		shift
	done

	set_arch
	prepare_dirs
	do_all_flintlock "$version" "$address" "$parent_iface" "$bridge_name" "$insecure" "$config_file"
}

cmd_direct_lvm() {
	local thinpool="$DEFAULT_THINPOOL"
	local skip_apt=false

	local disk=""
	while [ $# -gt 0 ]; do
		case "$1" in
		"-h" | "--help")
			cmd_direct_lvm_help
			exit 1
			;;
		"-y" | "--unattended")
			OPT_UNATTENDED=true
			;;
		"-d" | "--disk")
			shift
			disk="$1"
			;;
		"-t" | "--thinpool")
			shift
			thinpool="$1"
			;;
		"-s" | "--skip-apt")
			skip_apt=true
			;;
		*)
			die "Unknown argument: $1. Please use --help for help."
			;;
		esac
		shift
	done

	if [[ "$skip_apt" == false ]]; then
		cmd_apt
	fi

	do_all_direct_lvm "$disk" "$thinpool"
	say_warn "remember to set pool_name to $thinpool-thinpool in your containerd config"
}

cmd_devpool() {
	local thinpool="$DEFAULT_DEV_THINPOOL"

	local disk=""
	while [ $# -gt 0 ]; do
		case "$1" in
		"-h" | "--help")
			cmd_devpool_help
			exit 1
			;;
		"-t" | "--thinpool")
			shift
			thinpool="$1"
			;;
		*)
			die "Unknown argument: $1. Please use --help for help."
			;;
		esac
		shift
	done

	DEVELOPMENT=true
	prepare_dirs
	do_all_devpool "$thinpool"
	say_warn "remember to set pool_name to $thinpool-thinpool in your containerd config"
}



## COMMAND HELP FUNCS
#
#
cmd_apt_help() {
	cat <<EOF
  apt                    Install all apt packages required by flintlock

EOF
}

cmd_all_help() {
	cat <<EOF
  all                    Complete setup for production ready host. Component versions
                         can be configured by setting the FLINTLOCK, CONTAINERD and FIRECRACKER
			 environment variables.
    OPTIONS:
      -y                          Autoapprove all prompts (danger)
      --skip-apt, -s              Skip installation of apt packages
      --thinpool, -t              Name of thinpool to create (default: flintlock or flintlock-dev)
      --disk, -d                  Name blank unpartioned disk to use for direct lvm thinpool (ignored if --dev set)
      --grpc-address, -a          Address on which to start the Flintlock GRPC server (default: local ipv4 address)
      --parent-iface, -i          Interface of the default route of the host
      --bridge, -b                Bridge to use instead of an interface (will override --parent-iface)
      --insecure, -k              Start flintlockd without basic auth or certs
      --dev                       Set up development environment. Loop thinpools will be created.
      --flintlock-config-file, -f Path to a valid flintlockd configuration file with overriding config

EOF
}

cmd_firecracker_help() {
	cat <<EOF
  firecracker            Install firecracker from feature branch
    OPTIONS:
      --version, -v      Version to install (default: latest)

EOF
}

cmd_containerd_help() {
	cat <<EOF
  containerd             Install, configure and start containerd service
    OPTIONS:
      --version, -v      Version to install (default: latest)
      --thinpool, -t     Name of thinpool to include in config toml (default: flintlock-thinpool)
      --dev              Set up development environment. Containerd will keep state under 'dev' tagged paths.

EOF
}

cmd_flintlock_help() {
	cat <<EOF
  flintlock              Install and start flintlockd service (note: will not succeed without containerd)
    OPTIONS:
      --version, -v      Version to install (default: latest)
      --grpc-address, -a Address on which to start the GRPC server (default: local ipv4 address)
      --parent-iface, -i Interface of the default route of the host
      --bridge, -b       Bridge to use instead of an interface (will override --parent-iface)
      --insecure, -k     Start flintlockd without basic auth or certs
      --dev              Assumes containerd has been provisioned in a dev environment
      --config-file, -f  Path to a valid flintlockd configuration file with overriding config

EOF
}

cmd_direct_lvm_help() {
	cat <<EOF
  direct_lvm             Set up direct_lvm thinpool
    OPTIONS:
      -y                 Autoapprove all prompts (danger)
      --thinpool, -t     Name of thinpool to create (default: flintlock)
      --disk, -d         Name blank unpartioned disk to use for direct lvm thinpool
      --skip-apt, -s     Skip installation of apt packages

EOF
}

cmd_devpool_help() {
	cat <<EOF
  devpool                Set up loop device thinpool (development environments)
    OPTIONS:
      --thinpool, -t     Name of thinpool to create (default: flintlock-dev)

EOF
}

cmd_help() {
	cat <<EOF
usage: $0 <COMMAND> <OPTIONS>

Script to provision hosts for running flintlock microvms

COMMANDS:

EOF

	cmd_all_help
	cmd_apt_help
	cmd_firecracker_help
	cmd_containerd_help
	cmd_flintlock_help
	cmd_direct_lvm_help
	cmd_devpool_help
}

## LET'S DO THIS THING
#
#
main() {
	if [ $# = 0 ]; then
		die "No command provided. Please use \`$0 help\` for help."
	fi

	# Parse main command line args.
	#
	while [ $# -gt 0 ]; do
		case "$1" in
		-h | --help)
			cmd_help
			exit 1
			;;
		-*)
			die "Unknown arg: $1. Please use \`$0 help\` for help."
			;;
		*)
			break
			;;
		esac
		shift
	done

	if [[ $(id -u) != 0 ]]; then
		die "Run this script as root..." >&2
	fi

	# $1 is now a command name. Check if it is a valid command and, if so,
	# run it.
	#
	declare -f "cmd_$1" >/dev/null
	ok_or_die "Unknown command: $1. Please use \`$0 help\` for help."

	cmd=cmd_$1
	shift

	# $@ is now a list of command-specific args
	#
	$cmd "$@"
}

main "$@"