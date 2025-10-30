package account_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	addCmd := account.NewAddCommand(mockExecutor, mockConfigManager)
	cmd := addCmd.Command()

	assert.Equal(t, "add", cmd.Use)
	assert.Equal(t, "Add a new account", cmd.Short)
	assert.NotNil(t, cmd.RunE)
}

func TestAddCommand_LoadConfigError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(nil, errors.New("config not found"))

	addCmd := account.NewAddCommand(mockExecutor, mockConfigManager)
	cmd := addCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Global config not found")
}

func TestNewAddCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	addCmd := account.NewAddCommand(mockExecutor, mockConfigManager)
	cmd := addCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "add", cmd.Use)
}
