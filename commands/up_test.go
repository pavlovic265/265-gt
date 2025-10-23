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

// Test helper to create an up command with mock executor and git helper
func createUpCommandWithMock(t *testing.T) (
	*mocks.MockGitHelper, *gomock.Controller, *cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	upCmd := commands.NewUpCommand(mockExecutor, mockGitHelper)
	cmd := upCmd.Command()
	return mockGitHelper, ctrl, cmd
}

func TestUpCommand_Command(t *testing.T) {
	_, ctrl, cmd := createUpCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "up", cmd.Use)
	assert.Equal(t, "move to brunch up in stack", cmd.Short)
}

func TestUpCommand_RunE_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	upCmd := commands.NewUpCommand(mockExecutor, mockGitHelper)
	cmd := upCmd.Command()

	// Set up expectations for GetCurrentBranch
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranch().
		Return(branchName, nil)

	// Set up expectations for GetChildren
	mockGitHelper.EXPECT().
		GetChildren("main").
		Return([]string{"feature-branch"})

	// Set up expectations for git checkout
	expectedArgs := []string{"checkout", "feature-branch"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestUpCommand_RunE_NoChildren(t *testing.T) {
	mockGitHelper, ctrl, cmd := createUpCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranch
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranch().
		Return(branchName, nil)

	// Set up expectations for GetChildren (no children)
	mockGitHelper.EXPECT().
		GetChildren("main").
		Return([]string{})

	// Execute the command - should return error about no child branches
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Cannot move up - no child branches available")
}

func TestUpCommand_RunE_ExecutorError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	upCmd := commands.NewUpCommand(mockExecutor, mockGitHelper)
	cmd := upCmd.Command()

	expectedError := errors.New("git checkout failed")

	// Set up expectations for GetCurrentBranch
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranch().
		Return(branchName, nil)

	// Set up expectations for GetChildren
	mockGitHelper.EXPECT().
		GetChildren("main").
		Return([]string{"feature-branch"})

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to checkout branch")
	assert.Contains(t, err.Error(), "git checkout failed")
}

func TestNewUpCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewUpCommand creates a command with the correct executor and git helper
	upCmd := commands.NewUpCommand(mockExecutor, mockGitHelper)

	// Verify the command can be created
	cmd := upCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "up", cmd.Use)
}
