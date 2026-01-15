package account_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAccountCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	accountCmd := account.NewAccountCommand(mockRunner, mockConfigManager)
	cmd := accountCmd.Command()

	assert.Equal(t, "account", cmd.Use)
	assert.Equal(t, []string{"acc"}, cmd.Aliases)
	assert.Equal(t, "Manage accounts", cmd.Short)
}

func TestNewAccountCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	accountCmd := account.NewAccountCommand(mockRunner, mockConfigManager)
	cmd := accountCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "account", cmd.Use)
}
