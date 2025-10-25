package config

import (
	"fmt"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
)

// ============================================================================
// STRUCTS
// ============================================================================

type Account struct {
	User     string             `yaml:"user"`
	Token    string             `yaml:"token"`
	Platform constants.Platform `yaml:"platform"`
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
	exe executor.Executor
}

// ============================================================================
// INTERFACE
// ============================================================================

// ConfigManager interface defines the contract for config operations
type ConfigManager interface {
	// Initialization
	InitConfig(loadLocal bool)

	// Global config operations
	GetGlobalConfigPath() (string, error)
	LoadGlobalConfig() (*GlobalConfigStruct, error)
	SaveGlobalConfig(configToSave GlobalConfigStruct) error

	// Version operations
	SaveLastChecked() error
	SaveVersion(version string) error

	// Local config operations
	LoadLocalConfig() (*LocalConfigStruct, error)
	SaveLocalConfig(configToSave LocalConfigStruct) error

	// Protected branches config operations
	SaveProtectedBranches(branches []string) error
	GetProtectedBranches() []string

	// Account operations
	SaveActiveAccount(account Account) error
	SetActiveAccount(account Account) error
	GetActiveAccount() Account
	GetAccounts() []Account
	ClearActiveAccount() error
	HasActiveAccount() bool

	// Version operations
	GetVersion() Version
	GetCurrentVersion() string
}

var (
	localConfig  LocalConfigStruct
	globalConfig GlobalConfigStruct
)

// NewDefaultConfigManager creates a new DefaultConfigManager instance
func NewDefaultConfigManager(exe executor.Executor) *DefaultConfigManager {
	return &DefaultConfigManager{exe: exe}
}

func (d *DefaultConfigManager) InitConfig(loadLocal bool) {
	var err error
	globalConfig, err = d.loadGlobalConfig()
	if err != nil {
		fmt.Printf("Warning: Failed to load global config: %v\n", err)
		// Use empty config as fallback, but don't overwrite existing config
		globalConfig = GlobalConfigStruct{}
	}

	if loadLocal {
		localConfig, err = d.loadLocalConfig()
		if err != nil {
			fmt.Printf("Warning: Failed to load local config: %v\n", err)
			localConfig = LocalConfigStruct{}
		}
	}
}
