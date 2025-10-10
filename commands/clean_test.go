package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Test helper to create a clean command with mock executor and git helper
func createCleanCommandWithMock(t *testing.T) (*mocks.MockExecutor, *mocks.MockGitHelper, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cleanCmd := commands.NewCleanCommand(mockExecutor, mockGitHelper)
	cmd := cleanCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestCleanCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createCleanCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "clean", cmd.Use)
	assert.Equal(t, []string{"cl"}, cmd.Aliases)
	assert.Equal(t, "Clean up branches interactively", cmd.Short)
	assert.Equal(t, "Clean up branches one by one with confirmation. "+
		"Protected branches and current branch are skipped.", cmd.Long)
}

func TestCleanCommand_RunE_NoBranchesToClean(t *testing.T) {
	mockExecutor, mockGitHelper, ctrl, cmd := createCleanCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranchName(mockExecutor).
		Return(&branchName, nil)

	// Set up expectations for GetBranches
	mockGitHelper.EXPECT().
		GetBranches(mockExecutor).
		Return([]string{"main"}, nil)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestCleanCommand_RunE_WithBranchesToClean(t *testing.T) {
	mockExecutor, mockGitHelper, ctrl, cmd := createCleanCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranchName(mockExecutor).
		Return(&branchName, nil)

	// Set up expectations for GetBranches
	mockGitHelper.EXPECT().
		GetBranches(mockExecutor).
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
	// Note: This test doesn't fully test the interactive part, but shows the setup
	assert.NoError(t, err)
}

func TestNewCleanCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewCleanCommand creates a command with the correct executor and git helper
	cleanCmd := commands.NewCleanCommand(mockExecutor, mockGitHelper)
	assert.NotNil(t, cleanCmd)
}
