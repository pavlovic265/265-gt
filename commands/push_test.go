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

// Test helper to create a push command with mock executor
func createPushCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	pushCmd := commands.NewPushCommand(mockExecutor)
	cmd := pushCmd.Command()
	return mockExecutor, ctrl, cmd
}

func TestPushCommand_Command(t *testing.T) {
	_, ctrl, cmd := createPushCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "push", cmd.Use)
	assert.Equal(t, []string{"pu"}, cmd.Aliases)
	assert.Equal(t, "push branch always froce", cmd.Short)
	assert.False(t, cmd.DisableFlagParsing, "Expected DisableFlagParsing to be false")
}

func TestPushCommand_RunE_NoArgs(t *testing.T) {
	mockExecutor, ctrl, cmd := createPushCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName call
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Set up expectations for git push (ignores arguments, always pushes --force origin <current-branch>)
	expectedArgs := []string{"push", "--force", "origin", "main"}

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

func TestPushCommand_RunE_WithArgs(t *testing.T) {
	mockExecutor, ctrl, cmd := createPushCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName call
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Set up expectations for git push (ignores arguments, always pushes --force origin <current-branch>)
	expectedArgs := []string{"push", "--force", "origin", "main"}

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
	err := cmd.RunE(cmd, []string{"origin", "main", "--force"})
	assert.NoError(t, err)
}

func TestPushCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, ctrl, cmd := createPushCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git push failed")

	// Set up expectations for GetCurrentBranchName call
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Set up expectations for git push that will fail
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"push", "--force", "origin", "main"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestNewPushCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test that NewPushCommand creates a command with the correct executor
	pushCmd := commands.NewPushCommand(mockExecutor)

	// Verify the command can be created
	cmd := pushCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "push", cmd.Use)
}
