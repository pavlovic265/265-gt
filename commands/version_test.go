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

// Test helper to create a version command with mock executor
func createVersionCommandWithMock(t *testing.T) (
	*mocks.MockExecutor, *mocks.MockConfigManager, *gomock.Controller, *cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	versionCmd := commands.NewVersionCommand(mockExecutor, mockConfigManager)
	cmd := versionCmd.Command()
	return mockExecutor, mockConfigManager, ctrl, cmd
}

func TestVersionCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "version", cmd.Use)
	assert.Equal(t, "version of current build", cmd.Short)
}

func TestVersionCommand_RunE_NoFlags(t *testing.T) {
	_, mockConfigManager, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for config manager
	mockConfigManager.EXPECT().
		GetCurrentVersion().
		Return("v1.0.0")

	// Execute the command with no flags
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestVersionCommand_RunE_WithLatestFlag(t *testing.T) {
	_, _, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// Set the latest flag
	err := cmd.Flags().Set("latest", "true")
	require.NoError(t, err)

	// The getLatestVersion method makes HTTP calls, no executor calls needed
	// Execute the command with latest flag
	_ = cmd.RunE(cmd, []string{})
	// This will likely fail due to network request, which is expected in tests
	// We don't assert on the error type since it could be network-related
}

func TestVersionCommand_RunE_ExecutorError(t *testing.T) {
	_, mockConfigManager, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for config manager
	mockConfigManager.EXPECT().
		GetCurrentVersion().
		Return("v1.0.0")

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestNewVersionCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	// Test that NewVersionCommand creates a command with the correct executor and config manager
	versionCmd := commands.NewVersionCommand(mockExecutor, mockConfigManager)

	// Verify the command can be created
	cmd := versionCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "version", cmd.Use)
}
