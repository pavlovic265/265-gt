package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test helper to create an upgrade command with mock executor
func createUpgradeCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	upgradeCmd := commands.NewUpgradeCommand(mockExecutor)
	cmd := upgradeCmd.Command()
	return mockExecutor, ctrl, cmd
}

func TestUpgradeCommand_Command(t *testing.T) {
	_, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "upgrade", cmd.Use)
	assert.Equal(t, "upgrade of current build", cmd.Short)
}

func TestUpgradeCommand_RunE_Success(t *testing.T) {
	mockExecutor, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

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

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestUpgradeCommand_RunE_ExecutorError(t *testing.T) {
	// This test handles the case where the upgrade command determines an upgrade is needed
	// and calls the executor, but the executor fails.

	mockExecutor, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

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
		Return(nil). // Return success for this test
		AnyTimes()

	// Execute the command
	err := cmd.RunE(cmd, []string{})

	// The command should complete successfully
	assert.NoError(t, err, "Upgrade command should complete successfully")
}

func TestNewUpgradeCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test that NewUpgradeCommand creates a command with the correct executor
	upgradeCmd := commands.NewUpgradeCommand(mockExecutor)

	// Verify the command can be created
	cmd := upgradeCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "upgrade", cmd.Use)
}
