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

// Test helper to create an up command with mock executor
func createUpCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	upCmd := commands.NewUpCommand(mockExecutor)
	cmd := upCmd.Command()
	return mockExecutor, ctrl, cmd
}

func TestUpCommand_Command(t *testing.T) {
	_, ctrl, cmd := createUpCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "up", cmd.Use)
	assert.Equal(t, "move to brunch up in stack", cmd.Short)
}

func TestUpCommand_RunE_Success(t *testing.T) {
	mockExecutor, ctrl, cmd := createUpCommandWithMock(t)
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

	// Set up expectations for GetChildren
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("feature-branch\n"), nil)

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
	mockExecutor, ctrl, cmd := createUpCommandWithMock(t)
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

	// Set up expectations for GetChildren (no children)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString(""), nil)

	// Execute the command - should return error about TTY (interactive interface)
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not open a new TTY")
}

func TestUpCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, ctrl, cmd := createUpCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git checkout failed")

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

	// Set up expectations for GetChildren
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("feature-branch\n"), nil)

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
	assert.Equal(t, expectedError, err)
}

func TestNewUpCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test that NewUpCommand creates a command with the correct executor
	upCmd := commands.NewUpCommand(mockExecutor)

	// Verify the command can be created
	cmd := upCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "up", cmd.Use)
}
