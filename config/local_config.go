package config

import (
	"path/filepath"
	"strings"
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
	exeArgs := []string{"rev-parse", "--show-toplevel"}
	output, err := d.exe.WithGit().WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		return "", nil
	}

	localConfig := strings.TrimSpace(output.String())
	configPath := filepath.Join(localConfig, ".gtconfig.yaml")

	return configPath, nil
}
