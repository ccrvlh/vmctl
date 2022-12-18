package provision

import (
	"fmt"
	"vmctl/src/config"
)

// # Returns URL to latest release
//
//	build_release_url() {
//		local repo_name="$1"
//		echo "https://api.github.com/repos/$repo_name/releases/latest"
//	}
func BuildLatestReleaseURL(repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
}

func BuildVersionReleaseURL(repo string, version string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/releases/%s", repo, version)
}

// # Returns the desired binary download url for a repo, tag and binary
//
//	build_download_url() {
//		local repo_name="$1"
//		local tag="$2"
//		local bin="$3"
//		echo "https://github.com/$repo_name/releases/download/$tag/$bin"
//	}
func BuildBinaryDownloadURL(repo, tag, bin string) string {
	return fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", repo, tag, bin)
}

// # Returns the URL to a raw github file
//
//	build_raw_url() {
//		local repo_name="$1"
//		local file_name="$2"
//		echo "https://raw.githubusercontent.com/$repo_name/$DEFAULT_BRANCH/$file_name"
//	}
func BuildRawGithubAssetURL(repo, file string, cfg *config.AppConfig) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, cfg.DefaultBranch, file)

}

// # Returns the tag associated with a "latest" release
// latest_release_tag() {
// 	# shellcheck disable=SC2155
// 	local latest_url=$(build_release_url "$1")
// 	# shellcheck disable=SC2155
// 	local url=$(curl -sL "$latest_url" | awk -F'"' '/tag_name/ {printf $4}')
// 	echo "$url"
// }
func GetLatestReleaseTag(releaseURL string) {

}

// # Install the untarred binary attached to a release to /usr/local/bin
// install_release_bin() {
// 	local download_url="$1"
// 	local dest_file="$2"
// 	wget -q "$download_url" -O "$INSTALL_PATH/$dest_file" || die "failed to download release for $dest_file"
// 	chmod +x "$INSTALL_PATH/$dest_file"
// }
func DownloadReleaseBinary() {}

// # Install and untar the tarred binary attached to a release to /usr/local/bin
// install_release_tar() {
// 	local download_url="$1"
// 	local dest_path="$2"
// 	curl -sL "$download_url" | tar xz -C "$dest_path"
// }
func InstallReleaseBinary() {}

// # Download the given service file from the given repo
// fetch_service_file() {
// 	local repo="$1"
// 	local service="$2"
// 	local dest="$3"
// 	# shellcheck disable=SC2155
// 	local url=$(build_raw_url "$repo" "$service")
// 	curl -o "$dest" -sL "$url" || die "failed to download $service"
// 	chmod 0664 "$dest"
// 	systemctl daemon-reload
// }
func DownloadServiceFile() {}

// # Enable and start the given systemd service
// start_service() {
// 	local service="$1"
// 	systemctl enable "$service" &>/dev/null || die "failed to enable $service service"
// 	systemctl start "$service" || die "failed to start $service service"
// }
func StartService()  {}
func EnableService() {}
