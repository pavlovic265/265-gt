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

	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(homeDir, ".config")
	}

	configDir := filepath.Join(configHome, constants.GlobalConfigDir)
	configPath := filepath.Join(configDir, constants.GlobalConfigFileName)

	// Auto-migrate from old path (~/.gtconfig.yaml) to new XDG path
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		oldPath := filepath.Join(homeDir, constants.LocalConfigFileName)
		if _, err := os.Stat(oldPath); err == nil {
			if err := os.MkdirAll(configDir, 0o755); err != nil {
				return "", fmt.Errorf("failed to create config directory: %w", err)
			}
			if err := os.Rename(oldPath, configPath); err != nil {
				return "", fmt.Errorf("failed to migrate config file: %w", err)
			}
		}
	}

	return configPath, nil
}

func (d *DefaultConfigManager) LoadGlobalConfig() (*GlobalConfigStruct, error) {
	configPath, err := d.getGlobalConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("config file does not exist: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed to stat config file: %w", err)
	}

	return readConfig[GlobalConfigStruct](configPath)
}

func (d *DefaultConfigManager) SaveGlobalConfig(configToSave GlobalConfigStruct) error {
	configPath, err := d.getGlobalConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	return writeConfig(configPath, &configToSave)
}
