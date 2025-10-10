package commands_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test helper to create a delete command with mock executor
func createDeleteCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	deleteCmd := commands.NewDeleteCommand(mockExecutor)
	cmd := deleteCmd.Command()
	return mockExecutor, ctrl, cmd
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
	mockExecutor, ctrl, cmd := createDeleteCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Set up expectations for GetParent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.feature-branch.parent"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Set up expectations for GetChildren (parent)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("feature-branch\n"), nil)

	// Set up expectations for GetChildren (branch)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.feature-branch.children"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString(""), nil)

	// Set up expectations for git branch -D
	expectedArgs := []string{"branch", "-D", "feature-branch"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Set up expectations for RelinkParentChildren - SetChildren call
	// Since branchChildren is empty, the final children string will be empty after filtering
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.main.children", ""}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with branch name
	err := cmd.RunE(cmd, []string{"feature-branch"})
	assert.NoError(t, err)
}

func TestDeleteCommand_RunE_NoArgs(t *testing.T) {
	mockExecutor, ctrl, cmd := createDeleteCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Set up expectations for GetBranches
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "--list"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("* main\n  feature-branch\n"), nil)

	// Execute the command with no arguments - should show interactive selection
	// This will fail because we can't easily mock the interactive selection
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not open a new TTY")
}

func TestDeleteCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, ctrl, cmd := createDeleteCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git branch failed")

	// Mock GetCurrentBranchName call
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Mock GetParent call
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.feature-branch.parent"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Mock GetChildren call for parent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("feature-branch\n"), nil)

	// Mock GetChildren call for branch
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.feature-branch.children"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString(""), nil)

	// Mock the actual git branch -D command that should fail
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "-D", "feature-branch"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError)

	// Execute the command
	err := cmd.RunE(cmd, []string{"feature-branch"})
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestNewDeleteCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test that NewDeleteCommand creates a command with the correct executor
	deleteCmd := commands.NewDeleteCommand(mockExecutor)

	// Verify the command can be created
	cmd := deleteCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "delete", cmd.Use)
}
