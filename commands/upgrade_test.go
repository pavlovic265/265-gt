package commands_test

import (
	"bytes"
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

// Test helper to create an upgrade command with mock executor and config manager
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

func TestUpgradeCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "upgrade", cmd.Use)
	assert.Equal(t, "upgrade of current build", cmd.Short)
}

func TestUpgradeCommand_RunE_Success(t *testing.T) {
	mockExecutor, mockConfigManager, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for config manager

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(&config.GlobalConfigStruct{
			Version: &config.Version{
				CurrentVersion: "v0.1.0", // Different from latest to trigger upgrade
			},
		}, nil)

	// Set up expectations for checkWhichBinary() call
	mockExecutor.EXPECT().
		WithName("command").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"-v", "gt"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("homebrew"), nil)

	// Set up expectations for upgrade process
	mockExecutor.EXPECT().
		WithName("bash").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	mockConfigManager.EXPECT().
		SaveVersion(gomock.Any()).
		Return(nil)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestUpgradeCommand_RunE_ExecutorError(t *testing.T) {
	// This test handles the case where the upgrade command determines an upgrade is needed
	// and calls the executor, but the executor fails.

	mockExecutor, mockConfigManager, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for config manager

	mockConfigManager.EXPECT().
		LoadGlobalConfig().
		Return(&config.GlobalConfigStruct{
			Version: &config.Version{
				CurrentVersion: "v0.1.0", // Different from latest to trigger upgrade
			},
		}, nil)

	// Set up expectations for checkWhichBinary() call
	mockExecutor.EXPECT().
		WithName("command").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"-v", "gt"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("homebrew"), nil)

	// Set up expectations for the executor calls that will happen if an upgrade is needed
	// We use Any() to be flexible about the exact arguments since we can't easily mock the HTTP call
	mockExecutor.EXPECT().
		WithName("bash").
		Return(mockExecutor).
		AnyTimes()

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		Return(mockExecutor).
		AnyTimes()

	mockExecutor.EXPECT().
		Run().
		Return(fmt.Errorf("executor failed")). // Return error for this test
		AnyTimes()

	// Execute the command
	err := cmd.RunE(cmd, []string{})

	// The command should fail due to executor error
	assert.Error(t, err, "Upgrade command should fail when executor fails")
}

func TestNewUpgradeCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	// Test that NewUpgradeCommand creates a command with the correct executor and config manager
	upgradeCmd := commands.NewUpgradeCommand(mockExecutor, mockConfigManager)

	// Verify the command can be created
	cmd := upgradeCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "upgrade", cmd.Use)
}
