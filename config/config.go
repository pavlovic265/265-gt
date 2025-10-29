package config

import (
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// STRUCTS
// ============================================================================

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
	exe executor.Executor
}

// ============================================================================
// INTERFACE
// ============================================================================

// ConfigManager interface defines the contract for config operations
type ConfigManager interface {
	// Initialization
	InitLocalConfig()
	InitGlobalConfig()

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

func (d *DefaultConfigManager) InitLocalConfig() {
	lc, err := d.loadLocalConfig()
	if err != nil {
		log.Error("Failed to load local config: %v\n", err)
		return
	}
	localConfig = lc
}

func (d *DefaultConfigManager) InitGlobalConfig() {
	gc, err := d.loadGlobalConfig()
	if err != nil {
		log.Error("Failed to load global config: %v\n", err)
		return
	}
	globalConfig = gc
}
