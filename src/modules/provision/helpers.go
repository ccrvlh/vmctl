package provision

import (
	"fmt"
	"vmctl/src/config"
)

// # Returns URL to latest release
// build_release_url() {
// 	local repo_name="$1"
// 	echo "https://api.github.com/repos/$repo_name/releases/latest"
// }
func BuildLatestReleaseURL(repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
}

func BuildVersionReleaseURL(repo string, version string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/releases/%s", repo, version)
}

// # Returns the desired binary download url for a repo, tag and binary
// build_download_url() {
// 	local repo_name="$1"
// 	local tag="$2"
// 	local bin="$3"
// 	echo "https://github.com/$repo_name/releases/download/$tag/$bin"
// }
func BuildBinaryDownloadURL(repo, tag, bin string) string {
	return fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", repo, tag, bin)
}

// # Returns the URL to a raw github file
// build_raw_url() {
// 	local repo_name="$1"
// 	local file_name="$2"
// 	echo "https://raw.githubusercontent.com/$repo_name/$DEFAULT_BRANCH/$file_name"
// }
func BuildRawGithubAssetURL(repo, file string, cfg *config.AppConfig) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, cfg.DefaultBranch, file)

}
