package config

import (
	"path/filepath"
	"strings"
)

// LoadLocalConfig loads the local (repository-specific) configuration from disk
func (d *DefaultConfigManager) LoadLocalConfig() (*LocalConfigStruct, error) {
	configPath, err := d.getLocalConfigPath()
	if err != nil {
		return nil, err
	}
	return readConfig[LocalConfigStruct](configPath)
}

// SaveLocalConfig saves the local (repository-specific) configuration to disk
func (d *DefaultConfigManager) SaveLocalConfig(configToSave LocalConfigStruct) error {
	configPath, err := d.getLocalConfigPath()
	if err != nil {
		return err
	}
	err = writeConfig(configPath, &configToSave)
	if err != nil {
		return err
	}
	return nil
}

// loadLocalConfig loads the local configuration (private helper)
func (d *DefaultConfigManager) loadLocalConfig() (*LocalConfigStruct, error) {
	configPath, err := d.getLocalConfigPath()
	if err != nil {
		return nil, err
	}

	lconf, err := readConfig[LocalConfigStruct](configPath)
	if err != nil {
		return nil, err
	}

	return lconf, nil
}

// getLocalConfigPath returns the path to the local configuration file
func (d *DefaultConfigManager) getLocalConfigPath() (string, error) {
	exeArgs := []string{"rev-parse", "--show-toplevel"}
	output, err := d.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		// This should not happen since we already checked for git repository in main.go
		// But handle gracefully just in case
		return "", nil
	}

	localConfig := strings.TrimSpace(output.String())

	configPath := filepath.Join(localConfig, ".gtconfig.yaml")

	return configPath, nil
}
