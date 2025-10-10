package commands_test

import (
	"errors"
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

	// Set environment variable for the test
	t.Setenv("GT_REPOSITORY", "test/repo")

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestUpgradeCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, ctrl, cmd := createUpgradeCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("upgrade failed")

	mockExecutor.EXPECT().
		WithName("bash").
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError)

	// Set environment variable for the test
	t.Setenv("GT_REPOSITORY", "test/repo")

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
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
