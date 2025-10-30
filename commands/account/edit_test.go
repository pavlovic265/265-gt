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

func TestEditCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	editCmd := account.NewEditCommand(mockExecutor, mockConfigManager)
	cmd := editCmd.Command()

	assert.Equal(t, "edit", cmd.Use)
	assert.Equal(t, "Edit an existing account", cmd.Short)
	assert.NotNil(t, cmd.RunE)
}

func TestEditCommand_LoadConfigError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(nil, errors.New("config not found"))

	editCmd := account.NewEditCommand(mockExecutor, mockConfigManager)
	cmd := editCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Global config not found")
}

func TestEditCommand_NoAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	globalConfig := &config.GlobalConfigStruct{
		Accounts: []config.Account{},
	}

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(globalConfig, nil)

	editCmd := account.NewEditCommand(mockExecutor, mockConfigManager)
	cmd := editCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestNewEditCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	editCmd := account.NewEditCommand(mockExecutor, mockConfigManager)
	cmd := editCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "edit", cmd.Use)
}
