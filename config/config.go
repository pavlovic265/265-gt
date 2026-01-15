// Package config provides configuration management for gt CLI tool.
package config

import (
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/runner"
)

// Account represents a user account for GitHub or GitLab.
type Account struct {
	User       string             `yaml:"user"`
	Token      string             `yaml:"token"`
	Platform   constants.Platform `yaml:"platform"`
	Email      string             `yaml:"email,omitempty"`
	Name       string             `yaml:"name,omitempty"`
	SigningKey string             `yaml:"signingkey,omitempty"`
}

// Version tracks version information for update checking.
type Version struct {
	LastChecked    string `yaml:"last_checked"`
	CurrentVersion string `yaml:"current_version"`
}

// ThemeConfig stores theme preferences.
type ThemeConfig struct {
	Type constants.Theme `yaml:"type"`
}

// GlobalConfigStruct represents the global configuration file (~/.gtconfig.yaml).
type GlobalConfigStruct struct {
	Accounts      []Account    `yaml:"accounts,omitempty"`
	ActiveAccount *Account     `yaml:"active_account,omitempty"`
	Version       *Version     `yaml:"version,omitempty"`
	Theme         *ThemeConfig `yaml:"theme,omitempty"`
}

// LocalConfigStruct represents repository-local configuration (.gtconfig.yaml).
type LocalConfigStruct struct {
	Protected []string `yaml:"protected,omitempty"`
}

// DefaultConfigManager implements ConfigManager interface.
type DefaultConfigManager struct {
	runner runner.Runner
}

// ConfigManager defines the interface for loading and saving configuration.
type ConfigManager interface {
	LoadGlobalConfig() (*GlobalConfigStruct, error)
	SaveGlobalConfig(configToSave GlobalConfigStruct) error
	LoadLocalConfig() (*LocalConfigStruct, error)
	SaveLocalConfig(configToSave LocalConfigStruct) error
}

// NewDefaultConfigManager creates a new DefaultConfigManager instance.
func NewDefaultConfigManager(runner runner.Runner) *DefaultConfigManager {
	return &DefaultConfigManager{runner: runner}
}
