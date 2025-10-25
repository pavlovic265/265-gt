package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/pavlovic265/265-gt/utils/timeutils"
)

// ============================================================================
// GLOBAL CONFIG OPERATIONS
// ============================================================================

// GetGlobalConfigPath returns the path to the global configuration file
func (d *DefaultConfigManager) GetGlobalConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, constants.FileName)
	_, err = os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("config file does not exist: %w", err)
	}
	if err != nil {
		return "", fmt.Errorf("failed to stat config file: %w", err)
	}
	return configPath, nil
}

// LoadGlobalConfig loads the global configuration from disk
func (d *DefaultConfigManager) LoadGlobalConfig() (*GlobalConfigStruct, error) {
	configPath, err := d.GetGlobalConfigPath()
	if err != nil {
		return nil, err
	}
	return readConfig[GlobalConfigStruct](configPath)
}

// SaveGlobalConfig saves the global configuration to disk
func (d *DefaultConfigManager) SaveGlobalConfig(configToSave GlobalConfigStruct) error {
	configPath, err := d.GetGlobalConfigPath()
	if err != nil {
		return err
	}
	err = writeConfig(configPath, &configToSave)
	if err != nil {
		return err
	}
	return nil
}

// SaveLastChecked updates the last checked timestamp in the global config
func (d *DefaultConfigManager) SaveLastChecked() error {
	if globalConfig.Version == nil {
		globalConfig.Version = &Version{}
	}
	globalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)

	return d.SaveGlobalConfig(globalConfig)
}

// SaveVersion updates the version information in the global config
func (d *DefaultConfigManager) SaveVersion(version string) error {
	if globalConfig.Version == nil {
		globalConfig.Version = &Version{}
	}
	globalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)
	globalConfig.Version.CurrentVersion = version

	return d.SaveGlobalConfig(globalConfig)
}

// loadGlobalConfig loads the global configuration (private helper)
func (d *DefaultConfigManager) loadGlobalConfig() (GlobalConfigStruct, error) {
	configPath, err := d.GetGlobalConfigPath()
	if err != nil {
		return GlobalConfigStruct{}, err
	}
	gconf, err := readConfig[GlobalConfigStruct](configPath)
	if err != nil {
		return GlobalConfigStruct{}, err
	}

	return pointer.Deref(gconf), nil
}
