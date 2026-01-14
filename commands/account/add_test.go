package account_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/account"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// setAccountCommandContext sets up the context with config for account command tests
func setAccountCommandContext(cmd *cobra.Command, accounts []config.Account, activeAccount *config.Account) {
	cfg := config.NewConfigContext(&config.GlobalConfigStruct{
		Accounts:      accounts,
		ActiveAccount: activeAccount,
	}, nil)
	ctx := config.WithConfig(context.Background(), cfg)
	cmd.SetContext(ctx)
}

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

func TestAddCommand_NoContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	addCmd := account.NewAddCommand(mockExecutor, mockConfigManager)
	cmd := addCmd.Command()

	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no config found")
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
