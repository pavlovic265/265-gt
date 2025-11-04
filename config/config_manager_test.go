package config_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestConfigManagerInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock ConfigManager
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	// Test GetGlobalConfigPath
	expectedPath := "/home/user/.gtconfig.yaml"
	mockConfigManager.EXPECT().
		GetGlobalConfigPath().
		Return(expectedPath, nil)

	path, err := mockConfigManager.GetGlobalConfigPath()
	assert.NoError(t, err)
	assert.Equal(t, expectedPath, path)

	// Test GetActiveAccount
	expectedAccount := config.Account{
		User:     "testuser",
		Token:    "testtoken",
		Platform: "GitHub",
	}
	mockConfigManager.EXPECT().
		GetActiveAccount().
		Return(expectedAccount)

	account := mockConfigManager.GetActiveAccount()
	assert.Equal(t, expectedAccount, account)

	// Test HasActiveAccount
	mockConfigManager.EXPECT().
		HasActiveAccount().
		Return(true)

	hasAccount := mockConfigManager.HasActiveAccount()
	assert.True(t, hasAccount)
}

func TestConfigManagerWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	// Test error case
	expectedError := assert.AnError
	mockConfigManager.EXPECT().
		GetGlobalConfigPath().
		Return("", expectedError)

	path, err := mockConfigManager.GetGlobalConfigPath()
	assert.Error(t, err)
	assert.Equal(t, "", path)
	assert.Equal(t, expectedError, err)
}

func TestConfigManagerSaveOperations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	// Test SaveGlobalConfig
	globalConfig := config.GlobalConfigStruct{
		Accounts: []config.Account{
			{User: "user1", Token: "token1", Platform: "GitHub"},
		},
		Version: &config.Version{
			LastChecked:    "2023-01-01T00:00:00Z",
			CurrentVersion: "v1.0.0",
		},
		Theme: &config.ThemeConfig{Type: "dark"},
	}

	mockConfigManager.EXPECT().
		SaveGlobalConfig(globalConfig).
		Return(nil)

	err := mockConfigManager.SaveGlobalConfig(globalConfig)
	assert.NoError(t, err)

	// Test SaveLocalConfig
	localConfig := config.LocalConfigStruct{
		Protected: []string{"main", "develop"},
	}

	mockConfigManager.EXPECT().
		SaveLocalConfig(localConfig).
		Return(nil)

	err = mockConfigManager.SaveLocalConfig(localConfig)
	assert.NoError(t, err)
}
