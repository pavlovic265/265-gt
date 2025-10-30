package account_test

import (
	"errors"
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

func TestRemoveCommand_LoadConfigError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(nil, errors.New("config not found"))

	removeCmd := account.NewRemoveCommand(mockConfigManager)
	cmd := removeCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Global config not found")
}

func TestRemoveCommand_NoAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	globalConfig := &config.GlobalConfigStruct{
		Accounts: []config.Account{},
	}

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(globalConfig, nil)

	removeCmd := account.NewRemoveCommand(mockConfigManager)
	cmd := removeCmd.Command()

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
