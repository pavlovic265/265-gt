package config

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

type Account struct {
	User  string `yaml:"user"`
	Token string `yaml:"token"`
}

type GitHub struct {
	Assignee string    `yaml:"assignee"`
	Accounts []Account `yaml:"accounts"`
}

type Config struct {
	GitHub GitHub `yaml:"github"`
}

var GlobalConfig *Config

func InitConfig() {
	globalConfig, err := loadGlobalConfig()
	if err != nil {
		fmt.Println("globalConfig", err)
		globalConfig = &Config{}
	}
	localConfig, err := loadLoaclConfig()
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

func loadLoaclConfig() (*Config, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // Capture standard error as well
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running git command: %w - %s", err, strings.TrimSuffix(out.String(), "\n"))
	}
	localConfig := strings.TrimSpace(out.String())

	configPath := filepath.Join(localConfig, ".gtconfig.yaml")

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
