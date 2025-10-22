package commands_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Test helper to create a delete command with mock executor and git helper
func createDeleteCommandWithMock(t *testing.T) (
	*mocks.MockGitHelper, *gomock.Controller, *cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	deleteCmd := commands.NewDeleteCommand(mockExecutor, mockGitHelper)
	cmd := deleteCmd.Command()
	return mockGitHelper, ctrl, cmd
}

func TestDeleteCommand_Command(t *testing.T) {
	_, ctrl, cmd := createDeleteCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "delete", cmd.Use)
	assert.Equal(t, []string{"dl"}, cmd.Aliases)
	assert.Equal(t, "delete branch", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing, "Expected DisableFlagParsing to be true")
}

func TestDeleteCommand_RunE_WithBranchName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	deleteCmd := commands.NewDeleteCommand(mockExecutor, mockGitHelper)
	cmd := deleteCmd.Command()

	// Set up expectations for GetCurrentBranchName
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranchName().
		Return(&branchName, nil)

	// Set up expectations for IsProtectedBranch
	mockGitHelper.EXPECT().
		IsProtectedBranch("test-branch").
		Return(false)

	// Set up expectations for GetParent
	mockGitHelper.EXPECT().
		GetParent("test-branch").
		Return("main")

	// Set up expectations for GetChildren (branch)
	mockGitHelper.EXPECT().
		GetChildren("test-branch").
		Return([]string{})

	// Set up expectations for git branch -D command
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "-D", "test-branch"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Set up expectations for RelinkParentChildren
	mockGitHelper.EXPECT().
		RelinkParentChildren("main", []string{}).
		Return(nil)

	// Execute the command
	err := cmd.RunE(cmd, []string{"test-branch"})
	assert.NoError(t, err)
}

func TestDeleteCommand_RunE_GetCurrentBranchNameError(t *testing.T) {
	mockGitHelper, ctrl, cmd := createDeleteCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName to return error
	mockGitHelper.EXPECT().
		GetCurrentBranchName().
		Return(nil, errors.New("git error"))

	// Execute the command
	err := cmd.RunE(cmd, []string{"test-branch"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "git error")
}

func TestDeleteCommand_RunE_WithoutBranchName(t *testing.T) {
	mockGitHelper, ctrl, cmd := createDeleteCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranchName().
		Return(&branchName, nil)

	// Set up expectations for GetBranches
	mockGitHelper.EXPECT().
		GetBranches().
		Return([]string{"main", "feature-branch", "test-branch"}, nil)

	// Set up expectations for IsProtectedBranch
	mockGitHelper.EXPECT().
		IsProtectedBranch("feature-branch").
		Return(false)

	mockGitHelper.EXPECT().
		IsProtectedBranch("test-branch").
		Return(false)

	// Execute the command (this will trigger the interactive selection)
	err := cmd.RunE(cmd, []string{})
	// Note: This test expects a TTY error since we can't run interactive commands in tests
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not open a new TTY")
}

func TestNewDeleteCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewDeleteCommand creates a command with the correct executor and git helper
	deleteCmd := commands.NewDeleteCommand(mockExecutor, mockGitHelper)
	assert.NotNil(t, deleteCmd)
}
