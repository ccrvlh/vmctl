package utils

func GetLatestReleaseTag() string {
	// # Returns the tag associated with a "latest" release
	// latest_release_tag() {
	// 	# shellcheck disable=SC2155
	// 	local latest_url=$(build_release_url "$1")
	// 	# shellcheck disable=SC2155
	// 	local url=$(curl -sL "$latest_url" | awk -F'"' '/tag_name/ {printf $4}')
	// 	echo "$url"
	// }
	return "latest"
}
