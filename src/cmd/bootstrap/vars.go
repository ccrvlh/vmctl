package bootstrap

type BootstrapVars struct {
	// General vars
	MY_NAME         string
	INSTALL_PATH    string
	OPT_UNATTENDED  bool
	DEVELOPMENT     bool
	DEFAULT_VERSION string
	DEFAULT_BRANCH  string
	ARCH            string

	// Relevant Paths
	CONTAINERD_CONFIG_PATH  string
	CONTAINERD_ROOT_DIR     string
	CONTAINERD_STATE_DIR    string
	CONTAINERD_SERVICE_FILE string
	DEVMAPPER_DIR           string
	DEVPOOL_METADATA        string
	DEVPOOL_DATA            string

	// Firecracker Info
	FIRECRACKER_VERSION string
	FIRECRACKER_BIN     string
	FIRECRACKER_REPO    string

	// Containerd Info
	CONTAINERD_VERSION     string
	CONTAINERD_BIN         string
	CONTAINERD_REPO        string
	CONTAINERD_SYSTEMD_SVC string

	// Flintlock INfo
	FLINTLOCK_VERSION       string
	FLINTLOCK_BIN           string
	FLINTLOCK_REPO          string
	FLINTLOCKD_SERVICE_FILE string
	FLINTLOCKD_CONFIG_PATH  string

	// Thinpool
	THINPOOL_PROFILE_PATH string
	DEFAULT_THINPOOL      string
	DEFAULT_DEV_THINPOOL  string
	DATA_SPARSE_SIZE      string
	METADATA_SPARSE_SIZE  string

	// Magic calculations, for more information go and read
	// https://www.kernel.org/doc/Documentation/device-mapper/thin-provisioning.txt
	SECTORSIZE      int
	DATA_BLOCK_SIZE int

	// If free space on the data device drops below this level then a dm event will
	// be triggered which a userspace daemon should catch allowing it to extend the
	// pool device (picked arbitrarily)
	LOW_WATER_MARK int
}

func NewBootstrapVars() BootstrapVars {
	var mainVars = BootstrapVars{

		// general vars
		MY_NAME:         "flintlock $(basename '$0')",
		INSTALL_PATH:    "/usr/local/bin",
		OPT_UNATTENDED:  false,
		DEVELOPMENT:     false,
		DEFAULT_VERSION: "latest",
		DEFAULT_BRANCH:  "main",
		ARCH:            "amd64",

		// paths to be set later, put here to be explicit
		CONTAINERD_CONFIG_PATH:  "",
		CONTAINERD_ROOT_DIR:     "",
		CONTAINERD_STATE_DIR:    "",
		CONTAINERD_SERVICE_FILE: "",
		DEVMAPPER_DIR:           "",
		DEVPOOL_METADATA:        "",
		DEVPOOL_DATA:            "",

		// firecracker
		FIRECRACKER_VERSION: "${FIRECRACKER:=$DEFAULT_VERSION}",
		FIRECRACKER_BIN:     "firecracker",
		FIRECRACKER_REPO:    "weaveworks/firecracker",

		// containerd
		CONTAINERD_VERSION:     "${CONTAINERD:=$DEFAULT_VERSION}",
		CONTAINERD_BIN:         "containerd",
		CONTAINERD_REPO:        "containerd/containerd",
		CONTAINERD_SYSTEMD_SVC: "",

		// flintlock
		FLINTLOCK_VERSION:       "${FLINTLOCK:=$DEFAULT_VERSION}",
		FLINTLOCK_BIN:           "flintlockd",
		FLINTLOCK_REPO:          "weaveworks/flintlock",
		FLINTLOCKD_SERVICE_FILE: "/etc/systemd/system/flintlockd.service",
		FLINTLOCKD_CONFIG_PATH:  "/etc/opt/flintlockd/config.yaml",

		// thinpool
		THINPOOL_PROFILE_PATH: "/etc/lvm/profile",
		DEFAULT_THINPOOL:      "flintlock",
		DEFAULT_DEV_THINPOOL:  "flintlock-dev",
		DATA_SPARSE_SIZE:      "100G",
		METADATA_SPARSE_SIZE:  "10G",

		// Magic calculations, for more information go and read
		// https://www.kernel.org/doc/Documentation/device-mapper/thin-provisioning.txt
		SECTORSIZE:      512,
		DATA_BLOCK_SIZE: 128,

		// picked arbitrarily
		// If free space on the data device drops below this level then a dm event will
		// be triggered which a userspace daemon should catch allowing it to extend the
		// pool device.
		LOW_WATER_MARK: 32768,
	}
	return mainVars
}
