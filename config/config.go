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

type Account struct {
	User     string             `yaml:"user"`
	Token    string             `yaml:"token"`
	Platform constants.Platform `yaml:"platform"`
}

type Version struct {
	LastChecked string `yaml:"last_checked"`
	LastVersion string `yaml:"last_version"`
}

type GlobalConfigStruct struct {
	Accounts      []Account       `yaml:"accounts"`
	ActiveAccount *Account        `yaml:"active_account,omitempty"`
	Version       Version         `yaml:"version"`
	Theme         constants.Theme `yaml:"theme"`
}

type LocalConfigStruct struct {
	Protected []string `yaml:"protected"`
}

var (
	LocalConfig  LocalConfigStruct
	GlobalConfig GlobalConfigStruct
)

func InitConfig(exe executor.Executor, loadLocal bool) {
	globalConfig, err := loadGlobalConfig()
	if err != nil {
		fmt.Printf("Warning: Failed to load global config: %v\n", err)
		// Use empty config as fallback, but don't overwrite existing config
		globalConfig = GlobalConfigStruct{}
	}

	var localConfig LocalConfigStruct
	if loadLocal {
		var err error
		localConfig, err = loadLocalConfig(exe)
		if err != nil {
			fmt.Printf("Warning: Failed to load local config: %v\n", err)
			localConfig = LocalConfigStruct{}
		}
	}

	GlobalConfig = globalConfig
	LocalConfig = localConfig
}

func loadGlobalConfig() (GlobalConfigStruct, error) {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return GlobalConfigStruct{}, err
	}
	gconf, err := readConfig[GlobalConfigStruct](configPath)
	if err != nil {
		return GlobalConfigStruct{}, err
	}

	return pointer.Deref(gconf), nil
}

func loadLocalConfig(exe executor.Executor) (LocalConfigStruct, error) {

	configPath, err := getLocalConfigPath(exe)
	if err != nil {
		return LocalConfigStruct{}, err
	}

	lconf, err := readConfig[LocalConfigStruct](configPath)
	if err != nil {
		return LocalConfigStruct{}, err
	}

	return pointer.Deref(lconf), nil
}

func SaveConfig[T any](configPath string, configToSave T) error {
	err := writeConfig(configPath, &configToSave)
	if err != nil {
		return err
	}

	return nil
}

func UpdateLastChecked() error {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return err
	}

	GlobalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)

	return SaveConfig(configPath, GlobalConfig)
}

func UpdateVersion(version string) error {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return err
	}

	GlobalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)
	GlobalConfig.Version.LastVersion = version

	return SaveConfig(configPath, GlobalConfig)
}

func UpdateActiveAccount(account Account) error {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return err
	}

	GlobalConfig.ActiveAccount = &account

	return SaveConfig(configPath, GlobalConfig)
}

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

func GetGlobalConfigPath() (string, error) {
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

func getLocalConfigPath(exe executor.Executor) (string, error) {
	exeArgs := []string{"rev-parse", "--show-toplevel"}
	output, err := exe.WithGit().WithArgs(exeArgs).RunWithOutput()
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

func SetActiveAccount(account Account) error {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return err
	}

	GlobalConfig.ActiveAccount = &account
	return SaveConfig(configPath, GlobalConfig)
}

func GetActiveAccount() *Account {
	return GlobalConfig.ActiveAccount
}

func ClearActiveAccount() error {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return err
	}

	GlobalConfig.ActiveAccount = nil
	return SaveConfig(configPath, GlobalConfig)
}

func HasActiveAccount() bool {
	return GlobalConfig.ActiveAccount != nil
}
