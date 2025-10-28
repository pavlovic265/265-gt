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

// Test helper to create an unstage command with mock executor
func createUnstageCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	unstageCmd := commands.NewUnstageCommand(mockExecutor, mockGitHelper)
	cmd := unstageCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestUnstageCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createUnstageCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "unstage", cmd.Use)
	assert.Equal(t, []string{"us"}, cmd.Aliases)
	assert.Equal(t, "unstage ", cmd.Short)
	assert.False(t, cmd.DisableFlagParsing, "Expected DisableFlagParsing to be false")
}

func TestUnstageCommand_RunE_NoArgs(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createUnstageCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git restore --staged with no arguments
	expectedArgs := []string{"restore", "--staged"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with no arguments
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestUnstageCommand_RunE_WithArgs(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createUnstageCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git restore --staged with arguments
	expectedArgs := []string{"restore", "--staged", "file.txt", "file2.txt"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with arguments
	err := cmd.RunE(cmd, []string{"file.txt", "file2.txt"})
	assert.NoError(t, err)
}

func TestUnstageCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createUnstageCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git restore failed")

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"restore", "--staged"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to unstage files")
	assert.Contains(t, err.Error(), "git restore failed")
}

func TestNewUnstageCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewUnstageCommand creates a command with the correct executor
	unstageCmd := commands.NewUnstageCommand(mockExecutor, mockGitHelper)

	// Verify the command can be created
	cmd := unstageCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "unstage", cmd.Use)
}
