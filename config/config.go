package config

import (
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/runner"
)

type Account struct {
	User       string             `yaml:"user"`
	Token      string             `yaml:"token"`
	Platform   constants.Platform `yaml:"platform"`
	Email      string             `yaml:"email,omitempty"`
	Name       string             `yaml:"name,omitempty"`
	SigningKey string             `yaml:"signingkey,omitempty"`
}

type Version struct {
	LastChecked    string `yaml:"last_checked"`
	CurrentVersion string `yaml:"current_version"`
}

type ThemeConfig struct {
	Type constants.Theme `yaml:"type"`
}

type GlobalConfigStruct struct {
	Accounts      []Account    `yaml:"accounts,omitempty"`
	ActiveAccount *Account     `yaml:"active_account,omitempty"`
	Version       *Version     `yaml:"version,omitempty"`
	Theme         *ThemeConfig `yaml:"theme,omitempty"`
}

type LocalConfigStruct struct {
	Protected []string `yaml:"protected,omitempty"`
}

type DefaultConfigManager struct {
	runner runner.Runner
}

type ConfigManager interface {
	LoadGlobalConfig() (*GlobalConfigStruct, error)
	SaveGlobalConfig(configToSave GlobalConfigStruct) error
	LoadLocalConfig() (*LocalConfigStruct, error)
	SaveLocalConfig(configToSave LocalConfigStruct) error
}

func NewDefaultConfigManager(runner runner.Runner) *DefaultConfigManager {
	return &DefaultConfigManager{runner: runner}
}
