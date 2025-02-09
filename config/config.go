package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
	"gopkg.in/yaml.v2"
)

type Account struct {
	User  string `yaml:"user"`
	Token string `yaml:"token"`
}

type GitHub struct {
	Accounts []Account `yaml:"accounts"`
}

type Config struct {
	GitHub GitHub `yaml:"github"`
}

var GlobalConfig *Config

func InitConfig(exe executor.Executor) {
	globalConfig, err := loadGlobalConfig()
	if err != nil {
		fmt.Println("globalConfig", err)
		globalConfig = &Config{}
	}
	localConfig, err := loadLocalConfig(exe)
	if err != nil {
		fmt.Println("localConfig", err)
		localConfig = &Config{}
	}
	GlobalConfig = mergeConfig(*globalConfig, *localConfig)
}

func mergeConfig(c1, c2 Config) *Config {
	merged := c1
	vMerged := reflect.ValueOf(&merged).Elem()
	vc2 := reflect.ValueOf(c2)

	for i := 0; i < vMerged.NumField(); i++ {
		fieldValue := vc2.Field(i)
		defaultValue := reflect.Zero(fieldValue.Type()).Interface()

		if !reflect.DeepEqual(fieldValue.Interface(), defaultValue) {
			vMerged.Field(i).Set(fieldValue)
		}
	}

	return &merged
}

func loadGlobalConfig() (*Config, error) {
	// git rev-parse --show-toplevel
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".gtconfig.yaml")

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return cfg, nil
}

func loadLocalConfig(exe executor.Executor) (*Config, error) {
	exeArgs := []string{"rev-parse", "--show-toplevel"}
	output, err := exe.Execute("git", exeArgs...)
	if err != nil {
		return nil, err
	}
	localConfig := strings.TrimSpace(output.String())

	configPath := filepath.Join(localConfig, ".gtconfig.yaml")

	file, err := os.Open(configPath)
	cfg := &Config{}
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return cfg, nil
}
