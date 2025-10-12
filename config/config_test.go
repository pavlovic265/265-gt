package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/pavlovic265/265-gt/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestUpdateLastChecked_PreservesGitHubAccounts(t *testing.T) {
	// Create a temporary config file with GitHub accounts
	tempDir := t.TempDir()
	originalHomeDir := os.Getenv("HOME")
	defer func() {
		os.Setenv("HOME", originalHomeDir)
	}()

	os.Setenv("HOME", tempDir)
	configPath := filepath.Join(tempDir, constants.FileName)

	// Create initial config with GitHub accounts
	initialConfig := GlobalConfigStruct{
		Accounts: []Account{
			{User: "user1", Token: "token1", Platform: "GitHub"},
			{User: "user2", Token: "token2", Platform: "GitHub"},
		},
		Version: Version{
			LastChecked: "2023-01-01T00:00:00.000000Z",
			LastVersion: "1.0.0",
		},
		Theme: constants.DarkTheme,
	}

	// Write initial config to file
	file, err := os.Create(configPath)
	require.NoError(t, err)
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(initialConfig)
	require.NoError(t, err)
	encoder.Close()

	// Initialize config (this loads from file)
	GlobalConfig = initialConfig

	// Call UpdateLastChecked
	err = UpdateLastChecked()
	require.NoError(t, err)

	// Reload config from file to verify accounts are preserved
	reloadedConfig, err := loadGlobalConfig()
	require.NoError(t, err)

	// Verify accounts are preserved
	assert.Equal(t, 2, len(reloadedConfig.Accounts))
	assert.Equal(t, "user1", reloadedConfig.Accounts[0].User)
	assert.Equal(t, "token1", reloadedConfig.Accounts[0].Token)
	assert.Equal(t, "user2", reloadedConfig.Accounts[1].User)
	assert.Equal(t, "token2", reloadedConfig.Accounts[1].Token)

	// Verify timestamp was updated
	assert.NotEqual(t, "2023-01-01T00:00:00.000000Z", reloadedConfig.Version.LastChecked)

	// Verify the timestamp is recent (within last minute)
	parsedTime, err := time.Parse("2006-01-02T15:04:05.000000Z", reloadedConfig.Version.LastChecked)
	require.NoError(t, err)
	assert.WithinDuration(t, time.Now(), parsedTime, time.Minute)
}

func TestUpdateVersion_PreservesGitHubAccounts(t *testing.T) {
	// Create a temporary config file with GitHub accounts
	tempDir := t.TempDir()
	originalHomeDir := os.Getenv("HOME")
	defer func() {
		os.Setenv("HOME", originalHomeDir)
	}()

	os.Setenv("HOME", tempDir)
	configPath := filepath.Join(tempDir, constants.FileName)

	// Create initial config with GitHub accounts
	initialConfig := GlobalConfigStruct{
		Accounts: []Account{
			{User: "user1", Token: "token1", Platform: "GitHub"},
			{User: "user2", Token: "token2", Platform: "GitHub"},
		},
		Version: Version{
			LastChecked: "2023-01-01T00:00:00.000000Z",
			LastVersion: "1.0.0",
		},
		Theme: constants.DarkTheme,
	}

	// Write initial config to file
	file, err := os.Create(configPath)
	require.NoError(t, err)
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(initialConfig)
	require.NoError(t, err)
	encoder.Close()

	// Initialize config (this loads from file)
	GlobalConfig = initialConfig

	// Call UpdateVersion
	err = UpdateVersion("2.0.0")
	require.NoError(t, err)

	// Reload config from file to verify accounts are preserved
	reloadedConfig, err := loadGlobalConfig()
	require.NoError(t, err)

	// Verify accounts are preserved
	assert.Equal(t, 2, len(reloadedConfig.Accounts))
	assert.Equal(t, "user1", reloadedConfig.Accounts[0].User)
	assert.Equal(t, "token1", reloadedConfig.Accounts[0].Token)
	assert.Equal(t, "user2", reloadedConfig.Accounts[1].User)
	assert.Equal(t, "token2", reloadedConfig.Accounts[1].Token)

	// Verify version was updated
	assert.Equal(t, "2.0.0", reloadedConfig.Version.LastVersion)

	// Verify timestamp was updated
	assert.NotEqual(t, "2023-01-01T00:00:00.000000Z", reloadedConfig.Version.LastChecked)
}
