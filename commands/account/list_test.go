package account_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	assert.Equal(t, "list", cmd.Use)
	assert.Equal(t, []string{"ls"}, cmd.Aliases)
	assert.Equal(t, "List all accounts", cmd.Short)
	assert.NotNil(t, cmd.RunE)
}

func TestListCommand_LoadConfigError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(nil, errors.New("config error"))

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to load config")
}

func TestListCommand_NoAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	globalConfig := &config.GlobalConfigStruct{
		Accounts: []config.Account{},
	}

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(globalConfig, nil)

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestListCommand_WithAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	globalConfig := &config.GlobalConfigStruct{
		Accounts: []config.Account{
			{
				User:     "user1",
				Email:    "user1@example.com",
				Platform: constants.GitHubPlatform,
				Name:     "User One",
			},
			{
				User:     "user2",
				Email:    "user2@example.com",
				Platform: constants.GitLabPlatform,
				Name:     "User Two",
			},
		},
	}

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(globalConfig, nil)

	mockConfigManager.EXPECT().
		GetActiveAccount().
		Return(config.Account{
			User:     "user1",
			Platform: constants.GitHubPlatform,
		})

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestNewListCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "list", cmd.Use)
}
