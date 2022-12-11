package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var Cfg AppConfig

func LoadConfig(cfg *AppConfig) *AppConfig {
	readFile(cfg)
	readEnv(cfg)
	return cfg
}

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

func readEnv(cfg *AppConfig) {
	// Reads environment variables with `MICROVM_` prefix
	// It will overwrite current configuration that was set in YAML
	var err = envconfig.Process("MICROVM", cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println("Couldn't load config")
	fmt.Println(err)
	os.Exit(2)
}

type AppConfig struct {
	// General vars
	Port           string `yaml:"port" envconfig:"PORT"`
	InstallPath    string `yaml:"install_path" envconfig:"INSTALL_PATH"`
	OptUnattended  bool   `yaml:"opt_unattended" envconfig:"OPT_UNATTENDED"`
	Development    bool   `yaml:"dev" envconfig:"DEVELOPMENT"`
	DefaultVersion string `yaml:"default_version" envconfig:"DEFAULT_VERSION"`
	DefaultBranch  string `yaml:"default_branch" envconfig:"DEFAULT_BRANCH"`
	Arch           string `yaml:"arch" envconfig:"ARCH"`

	// Magic calculations, for more information go and read
	// https://www.kernel.org/doc/Documentation/device-mapper/thin-provisioning.txt
	Sectorsize    int `yaml:"sector_size" envconfig:"SECTOR_SIZE"`
	DataBlockSize int `yaml:"datablock_size" envconfig:"DATABLOCK_SIZE"`

	// If free space on the data device drops below this level then a dm event will
	// be triggered which a userspace daemon should catch allowing it to extend the
	// pool device (picked arbitrarily)
	LowWaterMark int `yaml:"low_watermark" envconfig:"low_watermark"`

	// DevMapper
	Devmapper struct {
		Dir      string `yaml:"dir" envconfig:"DEVMAPPER_DIR"`
		Data     string `yaml:"data" envconfig:"DEVMAPPER_DATA"`
		Metadata string `yaml:"metadata" envconfig:"DEVMAPPER_METADATA"`
	} `yaml:"devmapper"`

	// Firecracker Info
	Firecracker struct {
		Version string `yaml:"version" envconfig:"FIRECRACKER_VERSION"`
		Bin     string `yaml:"bin" envconfig:"FIRECRACKER_BIN"`
		Repo    string `yaml:"repo" envconfig:"FIRECRACKER_REPO"`
	} `yaml:"firecracker"`

	// Containerd Info
	Containerd struct {
		Version     string `yaml:"version" envconfig:"CONTAINERD_VERSION"`
		Bin         string `yaml:"bin" envconfig:"CONTAINERD_BIN"`
		Repo        string `yaml:"repo" envconfig:"CONTAINERD_REPO"`
		ConfigPath  string `yaml:"config_path" envconfig:"CONTAINERD_CONFIG_PATH"`
		RootDir     string `yaml:"root_dir" envconfig:"CONTAINERD_ROOT_DIR"`
		StateDir    string `yaml:"state_dir" envconfig:"CONTAINERD_STATE_DIR"`
		ServiceFile string `yaml:"systemd_unit" envconfig:"CONTAINERD_SYSTEMD_UNIT"`
		SystemdSvc  string `yaml:"systemd_svc" envconfig:"CONTAINERD_SYSTEMD_SVC"`
	} `yaml:"containerd"`

	// Flintlock INfo
	Flintlock struct {
		Version           string `yaml:"version" envconfig:"FLINTLOCK_VERSION"`
		Bin               string `yaml:"bin" envconfig:"FLINTLOCK_BIN"`
		Repo              string `yaml:"repo" envconfig:"FLINTLOCK_REPO"`
		DaemonServiceFile string `yaml:"daemon_svc_file" envconfig:"FLINTLOCK_DAEMON_SVC_FILE"`
		DaemonConfigPath  string `yaml:"daemon_config_path" envconfig:"FLINTLOCK_DAEMON_CONFIG_PATH"`
	} `yaml:"flintlock"`

	// Thinpool
	Thinpool struct {
		ProfilePath        string `yaml:"profile_path" envconfig:"THINPOOL_PROFILE_PATH"`
		Default            string `yaml:"default" envconfig:"THINPOOL_DEFAULT"`
		DefaultDev         string `yaml:"default_deev" envconfig:"THINPOOL_DEFAULT_DEEV"`
		DataSparseSize     string `yaml:"data_sparse_size" envconfig:"THINPOOL_DATA_SPARSE_SIZE"`
		MetadataSparseSize string `yaml:"metadata_sparse_size" envconfig:"THINPOOL_METADATA_SPARSE_SIZE"`
	} `yaml:"thinpool"`
}
