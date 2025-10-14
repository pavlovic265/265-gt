package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/pavlovic265/265-gt/utils/timeutils"
	"gopkg.in/yaml.v3"
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
	SaveLocalConfig(configToSave LocalConfigStruct) error
	SaveLastChecked() error
	SaveVersion(version string) error

	// Account operations
	SaveActiveAccount(account Account) error
	SetActiveAccount(account Account) error
	GetActiveAccount() Account
	GetAccounts() []Account
	GetVersion() Version
	GetCurrentVersion() string
	GetProtectedBranches() []string
	ClearActiveAccount() error
	HasActiveAccount() bool
}

// ============================================================================
// INITIALIZATION
// ============================================================================

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

// ============================================================================
// FUNCTIONS CONNECTED TO STRUCT (DefaultConfigManager methods)
// ============================================================================

// Global config operations
func (d *DefaultConfigManager) GetGlobalConfigPath() (string, error) {
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
	configPath, err := d.GetGlobalConfigPath()
	if err != nil {
		return nil, err
	}
	return readConfig[GlobalConfigStruct](configPath)
}

func (d *DefaultConfigManager) SaveGlobalConfig(configToSave GlobalConfigStruct) error {
	configPath, err := d.GetGlobalConfigPath()
	if err != nil {
		return err
	}
	err = writeConfig(configPath, &configToSave)
	if err != nil {
		return err
	}
	return nil
}

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

func (d *DefaultConfigManager) SaveLastChecked() error {
	if globalConfig.Version == nil {
		globalConfig.Version = &Version{}
	}
	globalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)

	return d.SaveGlobalConfig(globalConfig)
}

func (d *DefaultConfigManager) SaveVersion(version string) error {
	if globalConfig.Version == nil {
		globalConfig.Version = &Version{}
	}
	globalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)
	globalConfig.Version.CurrentVersion = version

	return d.SaveGlobalConfig(globalConfig)
}

// Account operations
func (d *DefaultConfigManager) SaveActiveAccount(account Account) error {
	globalConfig.ActiveAccount = pointer.From(account)

	return d.SaveGlobalConfig(globalConfig)
}

func (d *DefaultConfigManager) SetActiveAccount(account Account) error {
	globalConfig.ActiveAccount = pointer.From(account)

	return d.SaveGlobalConfig(globalConfig)
}

func (d *DefaultConfigManager) GetActiveAccount() Account {
	return pointer.Deref(globalConfig.ActiveAccount)
}

func (d *DefaultConfigManager) ClearActiveAccount() error {
	globalConfig.ActiveAccount = nil

	return d.SaveGlobalConfig(globalConfig)
}

func (d *DefaultConfigManager) HasActiveAccount() bool {
	return globalConfig.ActiveAccount != nil && globalConfig.ActiveAccount.User != ""
}

func (d *DefaultConfigManager) GetAccounts() []Account {
	accounts := make([]Account, len(globalConfig.Accounts))
	copy(accounts, globalConfig.Accounts)
	return accounts
}

func (d *DefaultConfigManager) GetCurrentVersion() string {
	if globalConfig.Version == nil {
		return ""
	}
	return globalConfig.Version.CurrentVersion
}

func (d *DefaultConfigManager) GetVersion() Version {
	return pointer.Deref(globalConfig.Version)
}

func (d *DefaultConfigManager) GetProtectedBranches() []string {
	return localConfig.Protected
}

// ============================================================================
// FUNCTIONS NOT CONNECTED TO STRUCT (standalone functions)
// ============================================================================

// Generic config operations
func readConfig[T any](filename string) (*T, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var cfg T
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &cfg, nil
}

func writeConfig[T any](filename string, cfg *T) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Secure file writing — ensure permissions are restrictive (rw for owner only)
	return os.WriteFile(filename, data, 0600)
}

// Config loading functions
func (d *DefaultConfigManager) loadGlobalConfig() (GlobalConfigStruct, error) {
	configPath, err := d.GetGlobalConfigPath()
	if err != nil {
		return GlobalConfigStruct{}, err
	}
	gconf, err := readConfig[GlobalConfigStruct](configPath)
	if err != nil {
		return GlobalConfigStruct{}, err
	}

	return pointer.Deref(gconf), nil
}

func (d *DefaultConfigManager) loadLocalConfig() (LocalConfigStruct, error) {
	configPath, err := d.getLocalConfigPath()
	if err != nil {
		return LocalConfigStruct{}, err
	}

	lconf, err := readConfig[LocalConfigStruct](configPath)
	if err != nil {
		return LocalConfigStruct{}, err
	}

	return pointer.Deref(lconf), nil
}

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

	_, err = os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	}

	return configPath, nil
}
