
## HELPER FUNCS
#
#
# Send a green message to stdout, followed by a new line
say() {
	[ -t 1 ] && [ -n "$TERM" ] &&
		echo "$(tput setaf 2)[$MY_NAME]$(tput sgr0) $*" ||
		echo "[$MY_NAME] $*"
}

# Send a green message to stdout, without a trailing new line
say_noln() {
	[ -t 1 ] && [ -n "$TERM" ] &&
		echo -n "$(tput setaf 2)[$MY_NAME]$(tput sgr0) $*" ||
		echo "[$MY_NAME] $*"
}

# Send a red message to stdout, followed by a new line
say_err() {
	[ -t 2 ] && [ -n "$TERM" ] &&
		echo -e "$(tput setaf 1)[$MY_NAME] $*$(tput sgr0)" 1>&2 ||
		echo -e "[$MY_NAME] $*" 1>&2
}

# Send a yellow message to stdout, followed by a new line
say_warn() {
	[ -t 1 ] && [ -n "$TERM" ] &&
		echo "$(tput setaf 3)[$MY_NAME] $*$(tput sgr0)" ||
		echo "[$MY_NAME] $*"
}

# Send a yellow message to stdout, without a trailing new line
say_warn_noln() {
	[ -t 1 ] && [ -n "$TERM" ] &&
		echo -n "$(tput setaf 3)[$MY_NAME] $*$(tput sgr0)" ||
		echo "[$MY_NAME] $*"
}

# Exit with an error message and (optional) code
# Usage: die [-c <error code>] <error message>
die() {
	code=1
	[[ "$1" = "-c" ]] && {
		code="$2"
		shift 2
	}
	say_err "$@"
	exit "$code"
}

# Exit with an error message if the last exit code is not 0
ok_or_die() {
	code=$?
	[[ $code -eq 0 ]] || die -c $code "$@"
}

# Check if /dev/kvm exists. Exit if it doesn't.
ensure_kvm() {
	[[ -c /dev/kvm ]] || die "/dev/kvm not found. Required for virtualisation. Aborting."
}

# Pause whatever is going on to ask for user confirmation
# Can pass optional custom message and options
get_user_confirmation() {
	# Skip if running unattended
	[[ "$OPT_UNATTENDED" = true ]] && return 0

	# Fail if STDIN is not a terminal (there's no user to confirm anything)
	[[ -t 0 ]] || return 1

	# Otherwise, ask the user
	msg=$([ -n "$1" ] && echo -n "$1" || echo -n "Continue? (y/n) ")
	yes=$([ -n "$2" ] && echo -n "$2" || echo -n "y")
	say_warn_noln "$msg"
	# shellcheck disable=SC2162
	read c && [ "$c" = "$yes" ] && return 0
	return 1
}
