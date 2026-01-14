package commands_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createUpgradeCommandWithMock(t *testing.T) (
	*mocks.MockExecutor, *mocks.MockConfigManager, *gomock.Controller, *cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	upgradeCmd := commands.NewUpgradeCommand(mockExecutor, mockConfigManager)
	cmd := upgradeCmd.Command()
	return mockExecutor, mockConfigManager, ctrl, cmd
}

// setUpgradeCommandContext sets up the context with config for upgrade command tests
func setUpgradeCommandContext(cmd *cobra.Command, version string) {
	cfg := config.NewConfigContext(&config.GlobalConfigStruct{
		Version: &config.Version{
			CurrentVersion: version,
		},
	}, nil)
	ctx := config.WithConfig(context.Background(), cfg)
	cmd.SetContext(ctx)
}

func TestUpgradeCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	assert.Equal(t, "upgrade", cmd.Use)
	assert.Equal(t, "upgrade of current build", cmd.Short)
}

func TestUpgradeCommand_RunE_Success(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	setUpgradeCommandContext(cmd, "v0.1.0")

	mockExecutor.EXPECT().
		WithName("command").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"-v", "gt"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("homebrew"), nil)

	mockExecutor.EXPECT().
		WithName("bash").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestUpgradeCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	setUpgradeCommandContext(cmd, "v0.1.0")

	mockExecutor.EXPECT().
		WithName("command").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"-v", "gt"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("homebrew"), nil)

	mockExecutor.EXPECT().
		WithName("bash").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(fmt.Errorf("executor failed"))

	err := cmd.RunE(cmd, []string{})

	assert.Error(t, err)
}

func TestNewUpgradeCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	upgradeCmd := commands.NewUpgradeCommand(mockExecutor, mockConfigManager)

	cmd := upgradeCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "upgrade", cmd.Use)
}
