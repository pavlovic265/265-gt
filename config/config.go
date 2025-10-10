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
	InitConfigWithLocal(exe, true)
}

func InitConfigWithLocal(exe executor.Executor, loadLocal bool) {
	globalConfig, err := loadGlobalConfig()
	if err != nil {
		fmt.Printf("Warning: Failed to load global config: %v\n", err)
		// Use empty config as fallback, but don't overwrite existing config
		globalConfig = GlobalConfigStruct{}
	}

	var localConfig LocalConfigStruct
	if loadLocal {
		var err error
		localConfig, err = loadLocalConfig(exe)
		if err != nil {
			fmt.Printf("Warning: Failed to load local config: %v\n", err)
			localConfig = LocalConfigStruct{}
		}
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
		return gconf, fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, FileName)
	_, err = os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return gconf, nil
	}
	if err != nil {
		return gconf, fmt.Errorf("failed to stat config file: %w", err)
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
	backupPath := configPath + ".backup"

	// Create backup of existing config if it exists
	if _, err := os.Stat(configPath); err == nil {
		if err := os.Rename(configPath, backupPath); err != nil {
			// If backup fails, continue anyway - not critical
			fmt.Printf("Warning: Failed to create backup: %v\n", err)
		}
	}

	file, err := os.Create(tempPath)
	if err != nil {
		// Restore backup if temp file creation fails
		if _, statErr := os.Stat(backupPath); statErr == nil {
			os.Rename(backupPath, configPath)
		}
		return fmt.Errorf("failed to create temp config file: %w", err)
	}
	defer func() { _ = file.Close() }()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	if err := encoder.Encode(configToSave); err != nil {
		_ = file.Close()
		_ = os.Remove(tempPath)
		// Restore backup if encoding fails
		if _, statErr := os.Stat(backupPath); statErr == nil {
			os.Rename(backupPath, configPath)
		}
		return fmt.Errorf("failed to encode config file: %w", err)
	}
	_ = file.Close()

	if err := os.Rename(tempPath, configPath); err != nil {
		_ = os.Remove(tempPath)
		// Restore backup if rename fails
		if _, statErr := os.Stat(backupPath); statErr == nil {
			os.Rename(backupPath, configPath)
		}
		return fmt.Errorf("failed to save config file: %w", err)
	}

	// Remove backup after successful save
	os.Remove(backupPath)

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
