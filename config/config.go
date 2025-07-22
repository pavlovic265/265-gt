package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
	"gopkg.in/yaml.v3"
)

type Account struct {
	User  string `yaml:"user"`
	Token string `yaml:"token"`
}

type GitHub struct {
	Accounts []Account `yaml:"accounts"`
}

type GlobalConfigStruct struct {
	GitHub GitHub `yaml:"github"`
}

type LocalConfigStruct struct {
	Protected []string `yaml:"protected"`
}

type ConfigStruct struct {
	LocalConfig  LocalConfigStruct
	GlobalConfig GlobalConfigStruct
}

var Config ConfigStruct

func InitConfig(exe executor.Executor) {
	globalConfig, err := loadGlobalConfig()
	if err != nil {
		fmt.Println("globalConfig", err)
		globalConfig = GlobalConfigStruct{}
	}
	localConfig, err := loadLocalConfig(exe)
	if err != nil {
		fmt.Println("localConfig", err)
		localConfig = LocalConfigStruct{}
	}
	Config = ConfigStruct{
		GlobalConfig: globalConfig,
		LocalConfig:  localConfig,
	}
}

func loadGlobalConfig() (GlobalConfigStruct, error) {
	gconf := GlobalConfigStruct{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return gconf, err
	}

	configPath := filepath.Join(homeDir, FileName)
	_, err = os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return gconf, nil
	}
	file, err := os.Open(configPath)
	if err != nil {
		return gconf, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&gconf); err != nil {
		return gconf, fmt.Errorf("failed to decode config file: %w", err)
	}

	return gconf, nil
}

func loadLocalConfig(exe executor.Executor) (LocalConfigStruct, error) {
	lconf := LocalConfigStruct{}
	exeArgs := []string{"rev-parse", "--show-toplevel"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		// This should not happen since we already checked for git repository in main.go
		// But handle gracefully just in case
		return lconf, nil
	}

	localConfig := strings.TrimSpace(output.String())

	configPath := filepath.Join(localConfig, ".gtconfig.yaml")

	_, err = os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return lconf, nil
	}

	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return lconf, nil
		}
		return lconf, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&lconf); err != nil {
		return lconf, fmt.Errorf("failed to decode config file: %w", err)
	}

	return lconf, nil
}
