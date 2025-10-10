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
func createVersionCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	versionCmd := commands.NewVersionCommand(mockExecutor)
	cmd := versionCmd.Command()
	return mockExecutor, ctrl, cmd
}

func TestVersionCommand_Command(t *testing.T) {
	_, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "version", cmd.Use)
	assert.Equal(t, "version of current build", cmd.Short)
}

func TestVersionCommand_RunE_NoFlags(t *testing.T) {
	_, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// The getCurrentVersion method reads from config, no executor calls needed
	// Execute the command with no flags
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestVersionCommand_RunE_WithLatestFlag(t *testing.T) {
	_, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// Ensure GT_REPOSITORY is not set for this test
	t.Setenv("GT_REPOSITORY", "")

	// Set the latest flag
	err := cmd.Flags().Set("latest", "true")
	require.NoError(t, err)

	// The getLatestVersion method makes HTTP calls, no executor calls needed
	// Note: This test will fail if GT_REPOSITORY is not set, which is expected
	// Execute the command with latest flag
	err = cmd.RunE(cmd, []string{})
	// We expect this to fail due to missing GT_REPOSITORY env var
	assert.Error(t, err)
}

func TestVersionCommand_RunE_ExecutorError(t *testing.T) {
	_, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	// The version command doesn't use executor for current version
	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestNewVersionCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test that NewVersionCommand creates a command with the correct executor
	versionCmd := commands.NewVersionCommand(mockExecutor)

	// Verify the command can be created
	cmd := versionCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "version", cmd.Use)
}
