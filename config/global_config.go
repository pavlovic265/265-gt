package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pavlovic265/265-gt/constants"
)

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

// loadGlobalConfig loads the global configuration (private helper)
func (d *DefaultConfigManager) loadGlobalConfig() (*GlobalConfigStruct, error) {
	configPath, err := d.GetGlobalConfigPath()
	if err != nil {
		return nil, err
	}
	gconf, err := readConfig[GlobalConfigStruct](configPath)
	if err != nil {
		return nil, err
	}

	return gconf, nil
}
