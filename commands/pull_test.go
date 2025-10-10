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

// Test helper to create a pull command with mock executor and git helper
func createPullCommandWithMock(t *testing.T) (*mocks.MockExecutor, *mocks.MockGitHelper, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	pullCmd := commands.NewPullCommand(mockExecutor, mockGitHelper)
	cmd := pullCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestPullCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createPullCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "pull", cmd.Use)
	assert.Equal(t, []string{"pl"}, cmd.Aliases)
	assert.Equal(t, "pull branch", cmd.Short)
	assert.False(t, cmd.DisableFlagParsing, "Expected DisableFlagParsing to be false")
}

func TestPullCommand_RunE_NoArgs(t *testing.T) {
	mockExecutor, mockGitHelper, ctrl, cmd := createPullCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName call
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranchName(mockExecutor).
		Return(&branchName, nil)

	// Set up expectations for git pull
	expectedArgs := []string{"pull", "origin", "main"}

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

func TestPullCommand_RunE_WithArgs(t *testing.T) {
	mockExecutor, mockGitHelper, ctrl, cmd := createPullCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName call
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranchName(mockExecutor).
		Return(&branchName, nil)

	// Set up expectations for git pull (ignores arguments, always pulls origin <current-branch>)
	expectedArgs := []string{"pull", "origin", "main"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with arguments (but they are ignored)
	err := cmd.RunE(cmd, []string{"origin", "main", "--rebase"})
	assert.NoError(t, err)
}

func TestPullCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, mockGitHelper, ctrl, cmd := createPullCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git pull failed")

	// Set up expectations for GetCurrentBranchName call
	branchName := "main"
	mockGitHelper.EXPECT().
		GetCurrentBranchName(mockExecutor).
		Return(&branchName, nil)

	// Set up expectations for git pull that will fail
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"pull", "origin", "main"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestNewPullCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewPullCommand creates a command with the correct executor and git helper
	pullCmd := commands.NewPullCommand(mockExecutor, mockGitHelper)

	// Verify the command can be created
	cmd := pullCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "pull", cmd.Use)
}
