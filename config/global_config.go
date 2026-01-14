package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pavlovic265/265-gt/constants"
)

func (d *DefaultConfigManager) getGlobalConfigPath() (string, error) {
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

func (d *DefaultConfigManager) LoadGlobalConfig() (*GlobalConfigStruct, error) {
	configPath, err := d.getGlobalConfigPath()
	if err != nil {
		return nil, err
	}
	return readConfig[GlobalConfigStruct](configPath)
}

func (d *DefaultConfigManager) SaveGlobalConfig(configToSave GlobalConfigStruct) error {
	configPath, err := d.getGlobalConfigPath()
	if err != nil {
		return err
	}
	return writeConfig(configPath, &configToSave)
}
