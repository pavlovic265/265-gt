package account_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAttachCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	attachCmd := account.NewAttachCommand(mockConfigManager)
	cmd := attachCmd.Command()

	assert.Equal(t, "attach [directory]", cmd.Use)
	assert.Equal(t, []string{"at"}, cmd.Aliases)
	assert.Equal(t, "Attach active account to a directory", cmd.Short)
	assert.NotNil(t, cmd.RunE)
}

func TestAttachCommand_NoActiveAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	attachCmd := account.NewAttachCommand(mockConfigManager)
	cmd := attachCmd.Command()

	setAccountCommandContext(cmd, []config.Account{}, nil)

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No active account")
}

func TestAttachCommand_NonExistentDirectory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	attachCmd := account.NewAttachCommand(mockConfigManager)
	cmd := attachCmd.Command()

	activeAccount := &config.Account{
		User:  "testuser",
		Email: "test@example.com",
		Name:  "Test User",
	}
	setAccountCommandContext(cmd, []config.Account{*activeAccount}, activeAccount)

	err := cmd.RunE(cmd, []string{"/path/that/does/not/exist"})
	assert.Error(t, err)
}

func TestNewAttachCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	attachCmd := account.NewAttachCommand(mockConfigManager)
	cmd := attachCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "attach [directory]", cmd.Use)
}
