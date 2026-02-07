package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pavlovic265/265-gt/constants"
)

func (d *DefaultConfigManager) getPublicConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(homeDir, ".config")
	}

	configDir := filepath.Join(configHome, constants.GlobalConfigDir)
	return filepath.Join(configDir, constants.PublicConfigFileName), nil
}

func (d *DefaultConfigManager) SavePublicConfig(globalConfig GlobalConfigStruct) error {
	configPath, err := d.getPublicConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	publicCfg := PublicConfigStruct{}
	if globalConfig.ActiveAccount != nil {
		publicCfg.User = globalConfig.ActiveAccount.User
		publicCfg.Platform = globalConfig.ActiveAccount.Platform
	}

	return writeConfigWithPerm(configPath, &publicCfg, 0644)
}
