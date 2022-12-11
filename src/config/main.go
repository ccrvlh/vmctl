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
	fmt.Printf("%+v", cfg)
	return cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
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
	var err = envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

type AppConfig struct {
	// General vars
	Port            string `yaml:"port", envconfig:"SERVER_PORT"`
	INSTALL_PATH    string `yaml:"port", envconfig:"SERVER_PORT"`
	OPT_UNATTENDED  bool   `yaml:"port", envconfig:"SERVER_PORT"`
	DEVELOPMENT     bool   `yaml:"port", envconfig:"SERVER_PORT"`
	DEFAULT_VERSION string `yaml:"port", envconfig:"SERVER_PORT"`
	DEFAULT_BRANCH  string `yaml:"port", envconfig:"SERVER_PORT"`
	ARCH            string `yaml:"port", envconfig:"SERVER_PORT"`

	// Magic calculations, for more information go and read
	// https://www.kernel.org/doc/Documentation/device-mapper/thin-provisioning.txt
	SECTORSIZE      int `yaml:"port", envconfig:"SERVER_PORT"`
	DATA_BLOCK_SIZE int `yaml:"port", envconfig:"SERVER_PORT"`

	// If free space on the data device drops below this level then a dm event will
	// be triggered which a userspace daemon should catch allowing it to extend the
	// pool device (picked arbitrarily)
	LOW_WATER_MARK int `yaml:"port", envconfig:"SERVER_PORT"`

	// DevMapper
	Devmapper struct {
		DEVMAPPER_DIR    string `yaml:"port", envconfig:"SERVER_PORT"`
		DEVPOOL_METADATA string `yaml:"port", envconfig:"SERVER_PORT"`
		DEVPOOL_DATA     string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml: "devmapper"`

	// Firecracker Info
	Firecracker struct {
		FIRECRACKER_VERSION string `yaml:"port", envconfig:"SERVER_PORT"`
		FIRECRACKER_BIN     string `yaml:"port", envconfig:"SERVER_PORT"`
		FIRECRACKER_REPO    string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml: "firecracker"`

	// Containerd Info
	Containerd struct {
		CONTAINERD_VERSION      string `yaml:"port", envconfig:"SERVER_PORT"`
		CONTAINERD_BIN          string `yaml:"port", envconfig:"SERVER_PORT"`
		CONTAINERD_REPO         string `yaml:"port", envconfig:"SERVER_PORT"`
		CONTAINERD_SYSTEMD_SVC  string `yaml:"port", envconfig:"SERVER_PORT"`
		CONTAINERD_CONFIG_PATH  string `yaml:"port", envconfig:"SERVER_PORT"`
		CONTAINERD_ROOT_DIR     string `yaml:"port", envconfig:"SERVER_PORT"`
		CONTAINERD_STATE_DIR    string `yaml:"port", envconfig:"SERVER_PORT"`
		CONTAINERD_SERVICE_FILE string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml: "containerd"`

	// Flintlock INfo
	Flintlock struct {
		FLINTLOCK_VERSION       string `yaml:"port", envconfig:"SERVER_PORT"`
		FLINTLOCK_BIN           string `yaml:"port", envconfig:"SERVER_PORT"`
		FLINTLOCK_REPO          string `yaml:"port", envconfig:"SERVER_PORT"`
		FLINTLOCKD_SERVICE_FILE string `yaml:"port", envconfig:"SERVER_PORT"`
		FLINTLOCKD_CONFIG_PATH  string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml: "flintlock"`

	// Thinpool
	Thinpool struct {
		THINPOOL_PROFILE_PATH string `yaml:"port", envconfig:"SERVER_PORT"`
		DEFAULT_THINPOOL      string `yaml:"port", envconfig:"SERVER_PORT"`
		DEFAULT_DEV_THINPOOL  string `yaml:"port", envconfig:"SERVER_PORT"`
		DATA_SPARSE_SIZE      string `yaml:"port", envconfig:"SERVER_PORT"`
		METADATA_SPARSE_SIZE  string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml: "thinpool"`
}
