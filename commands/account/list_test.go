package account_test

import (
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

func TestListCommand_NoContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no config found")
}

func TestListCommand_NoAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	setAccountCommandContext(cmd, []config.Account{}, nil)

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestListCommand_WithAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	activeAccount := &config.Account{
		User:     "user1",
		Email:    "user1@example.com",
		Platform: constants.GitHubPlatform,
		Name:     "User One",
	}

	accounts := []config.Account{
		*activeAccount,
		{
			User:     "user2",
			Email:    "user2@example.com",
			Platform: constants.GitLabPlatform,
			Name:     "User Two",
		},
	}

	listCmd := account.NewListCommand(mockConfigManager)
	cmd := listCmd.Command()

	setAccountCommandContext(cmd, accounts, activeAccount)

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
