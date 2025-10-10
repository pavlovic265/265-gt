package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/timeutils"
	"gopkg.in/yaml.v3"
)

type Account struct {
	User  string `yaml:"user"`
	Token string `yaml:"token"`
}

type GitHub struct {
	Accounts []Account `yaml:"accounts"`
}

type Version struct {
	LastChecked string `yaml:"last_checked"`
	LastVersion string `yaml:"last_version"`
}

type Theme struct {
	Type string `yaml:"type"` // "dark" or "light"
}

type GlobalConfigStruct struct {
	GitHub  GitHub  `yaml:"github"`
	Version Version `yaml:"version"`
	Theme   Theme   `yaml:"theme"`
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

// HasValidConfig returns true if the config was successfully loaded from disk
func HasValidConfig() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	configPath := filepath.Join(homeDir, FileName)
	_, err = os.Stat(configPath)
	return err == nil
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
	defer func() { _ = file.Close() }()
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
	defer func() { _ = file.Close() }()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&lconf); err != nil {
		return lconf, fmt.Errorf("failed to decode config file: %w", err)
	}

	return lconf, nil
}

func saveGlobalConfig(configToSave GlobalConfigStruct) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(homeDir, FileName)
	tempPath := configPath + ".tmp"

	file, err := os.Create(tempPath)
	if err != nil {
		return fmt.Errorf("failed to create temp config file: %w", err)
	}
	defer func() { _ = file.Close() }()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	if err := encoder.Encode(configToSave); err != nil {
		_ = file.Close()
		_ = os.Remove(tempPath)
		return fmt.Errorf("failed to encode config file: %w", err)
	}
	_ = file.Close()

	if err := os.Rename(tempPath, configPath); err != nil {
		_ = os.Remove(tempPath)
		return fmt.Errorf("failed to save config file: %w", err)
	}

	return nil
}

func UpdateLastChecked() error {
	Config.GlobalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)

	return saveGlobalConfig(Config.GlobalConfig)
}

func UpdateVersion(version string) error {
	Config.GlobalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)
	Config.GlobalConfig.Version.LastVersion = version

	return saveGlobalConfig(Config.GlobalConfig)
}
