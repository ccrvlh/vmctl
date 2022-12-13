package config

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var Cfg AppConfig

// Loads the configuration
// First it tries to read from the a `.yml` file
// After it loaded config values from a config file
// Overwrites values with variables that were set
// with EnvironmentVariables
func LoadConfig(cfg *AppConfig) *AppConfig {
	readFile(cfg)
	setMainDefaults(cfg)
	readEnv(cfg)
	if cfg.Shell == "" {
		cfg.Shell = os.Getenv("$SHELL")
	}
	return cfg
}

// Set main Defaults (Home/Shell)
// This included
func setMainDefaults(cfg *AppConfig) {
	var shellEcho, shellErr = exec.Command("sh", "-c", "echo", "$SHELL").Output()
	if shellErr != nil {
		fmt.Printf("Couldn't get Shell, defaulting to Bash")
	} else {
		cfg.Shell = string(shellEcho)
	}

	var homeEcho, homeErr = exec.Command(cfg.Shell, "-c", "echo", "$HOME").Output()
	if homeErr != nil {
		fmt.Printf("Couldn't get Home directory, defaulting to /usr/local")
		cfg.HomeDir = "/usr/local/.flintlock"
	} else {
		cfg.HomeDir = string(homeEcho)
	}

}

// The file path comes from VMCTL_CONFIG_PATH
// If not provided, default values are used.
func readFile(cfg *AppConfig) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var decoder = yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

// Reads environment variables with `MICROVM_` prefix
// It will overwrite current configuration that was set in YAML
func readEnv(cfg *AppConfig) {
	var err = envconfig.Process("MICROVM", cfg)
	if err != nil {
		processError(err)
	}
}

// Utils
func processError(err error) {
	fmt.Println("Couldn't load config")
	fmt.Println(err)
	os.Exit(2)
}

// The Application's full configuration
// This includes everything related to Flintlock, Firecracker, Containerd
// And others.
type AppConfig struct {
	// General vars
	Port           string `yaml:"port" envconfig:"PORT"`
	InstallPath    string `yaml:"install_path" envconfig:"INSTALL_PATH"`
	OptUnattended  bool   `yaml:"opt_unattended" envconfig:"OPT_UNATTENDED"`
	Development    bool   `yaml:"dev" envconfig:"DEVELOPMENT"`
	DefaultVersion string `yaml:"default_version" envconfig:"DEFAULT_VERSION"`
	DefaultBranch  string `yaml:"default_branch" envconfig:"DEFAULT_BRANCH"`
	Arch           string `yaml:"arch" envconfig:"ARCH"`
	HomeDir        string `yaml:"home" envconfig:"HOME"`
	Shell          string `yaml:"shell" envconfig:"SHELL"`
	LogLevel       string `yaml:"log_levle" envconfig:"LOG_LEVEL"`

	// Magic calculations, for more information go and read
	// https://www.kernel.org/doc/Documentation/device-mapper/thin-provisioning.txt
	Sectorsize    int `yaml:"sector_size" envconfig:"SECTOR_SIZE"`
	DataBlockSize int `yaml:"datablock_size" envconfig:"DATABLOCK_SIZE"`

	// If free space on the data device drops below this level then a dm event will
	// be triggered which a userspace daemon should catch allowing it to extend the
	// pool device (picked arbitrarily)
	LowWaterMark int `yaml:"low_watermark" envconfig:"low_watermark"`

	// Firecracker Info
	Firecracker struct {
		Version string `yaml:"version" envconfig:"FIRECRACKER_VERSION"`
		Bin     string `yaml:"bin" envconfig:"FIRECRACKER_BIN"`
		Repo    string `yaml:"repo" envconfig:"FIRECRACKER_REPO"`
	} `yaml:"firecracker"`

	// Containerd Info
	Containerd struct {
		Version               string `yaml:"version" envconfig:"CONTAINERD_VERSION"`
		Bin                   string `yaml:"bin" envconfig:"CONTAINERD_BIN"`
		Repo                  string `yaml:"repo" envconfig:"CONTAINERD_REPO"`
		ConfigPath            string `yaml:"config_path" envconfig:"CONTAINERD_CONFIG_PATH"`
		RootDir               string `yaml:"root_dir" envconfig:"CONTAINERD_ROOT_DIR"`
		StateDir              string `yaml:"state_dir" envconfig:"CONTAINERD_STATE_DIR"`
		ServiceFile           string `yaml:"systemd_unit" envconfig:"CONTAINERD_SYSTEMD_UNIT"`
		SystemdSvc            string `yaml:"systemd_svc" envconfig:"CONTAINERD_SYSTEMD_SVC"`
		MetricsEndpoint       string `yaml:"metrics_endpoint" envconfig:"CONTAINERD_METRICS_ENDPOINT"`
		BaseImageSize         string `yaml:"base_image_size" envconfig:"CONTAINERD_BASE_IMAGE_SIZE"`
		LogLevel              string `yaml:"log_level" envconfig:"CONTAINERD_LOG_LEVEL"`
		DevmapperDir          string `yaml:"devmapper_pool" envconfig:"CONTAINERD_LOG_LEVEL"`
		DevmapperPoolData     string `yaml:"devmapper_pool_data" envconfig:"CONTAINERD_DEVMAPPER_POOL_DATA"`
		DevmapperPoolMetadata string `yaml:"devmapper_pool_metadaa" envconfig:"CONTAINERD_DEVMAPPER_POOL_METADATA"`
	} `yaml:"containerd"`

	// Thinpool
	Thinpool struct {
		ProfilePath        string `yaml:"profile_path" envconfig:"THINPOOL_PROFILE_PATH"`
		Default            string `yaml:"default" envconfig:"THINPOOL_DEFAULT"`
		DefaultDev         string `yaml:"default_deev" envconfig:"THINPOOL_DEFAULT_DEEV"`
		DataSparseSize     string `yaml:"data_sparse_size" envconfig:"THINPOOL_DATA_SPARSE_SIZE"`
		MetadataSparseSize string `yaml:"metadata_sparse_size" envconfig:"THINPOOL_METADATA_SPARSE_SIZE"`
	} `yaml:"thinpool"`

	// Flintlock INfo
	Flintlock struct {
		Version           string `yaml:"version" envconfig:"FLINTLOCK_VERSION"`
		Bin               string `yaml:"bin" envconfig:"FLINTLOCK_BIN"`
		Repo              string `yaml:"repo" envconfig:"FLINTLOCK_REPO"`
		DaemonServiceFile string `yaml:"daemon_svc_file" envconfig:"FLINTLOCK_DAEMON_SVC_FILE"`
		DaemonConfigPath  string `yaml:"daemon_config_path" envconfig:"FLINTLOCK_DAEMON_CONFIG_PATH"`
		NetworkName       string `yaml:"network_name" envconfig:"FLINTLOCK_NETWORK_NAME"`
		BridgeName        string `yaml:"bridge_name" envconfig:"FLINTLOCK_BRIDGE_NAME"`
		IpAddress         string `yaml:"ip_address" envconfig:"FLINTLOCK_IP_ADDRESS"`
		IpAddressStart    string `yaml:"ip_address_start" envconfig:"FLINTLOCK_IP_ADDRESS_START"`
		IpAddressEnd      string `yaml:"ip_address_end" envconfig:"FLINTLOCK_IP_ADDRESS_END"`
	} `yaml:"flintlock"`
}
