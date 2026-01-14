package account_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRemoveCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	removeCmd := account.NewRemoveCommand(mockConfigManager)
	cmd := removeCmd.Command()

	assert.Equal(t, "remove", cmd.Use)
	assert.Equal(t, []string{"rm"}, cmd.Aliases)
	assert.Equal(t, "Remove an account", cmd.Short)
	assert.NotNil(t, cmd.RunE)
}

func TestRemoveCommand_NoContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	removeCmd := account.NewRemoveCommand(mockConfigManager)
	cmd := removeCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no config found")
}

func TestRemoveCommand_NoAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	removeCmd := account.NewRemoveCommand(mockConfigManager)
	cmd := removeCmd.Command()

	setAccountCommandContext(cmd, []config.Account{}, nil)

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestNewRemoveCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	removeCmd := account.NewRemoveCommand(mockConfigManager)
	cmd := removeCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "remove", cmd.Use)
}
