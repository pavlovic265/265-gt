package config

import (
	"path/filepath"
)

func (d *DefaultConfigManager) LoadLocalConfig() (*LocalConfigStruct, error) {
	configPath, err := d.getLocalConfigPath()
	if err != nil {
		return nil, err
	}
	return readConfig[LocalConfigStruct](configPath)
}

func (d *DefaultConfigManager) SaveLocalConfig(configToSave LocalConfigStruct) error {
	configPath, err := d.getLocalConfigPath()
	if err != nil {
		return err
	}
	return writeConfig(configPath, &configToSave)
}

func (d *DefaultConfigManager) getLocalConfigPath() (string, error) {
	output, err := d.runner.GitOutput("rev-parse", "--show-toplevel")
	if err != nil {
		return "", nil // Local config is optional - not being in a git repo is fine
	}

	configPath := filepath.Join(output, ".gtconfig.yaml")
	return configPath, nil
}
