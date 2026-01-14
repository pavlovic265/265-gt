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

func createStatusCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	statusCmd := commands.NewStatusCommand(mockExecutor, mockGitHelper)
	cmd := statusCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestStatusCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createStatusCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "status", cmd.Use)
	assert.Equal(t, []string{"st"}, cmd.Aliases)
	assert.Equal(t, "git status", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing, "Expected DisableFlagParsing to be true")
	assert.True(t, cmd.SilenceUsage, "Expected SilenceUsage to be true")
}

func TestStatusCommand_RunE_NoArgs(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createStatusCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git status with no arguments
	expectedArgs := []string{"status"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("On branch main\nnothing to commit, working tree clean"), nil)

	// Execute the command with no arguments
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestStatusCommand_RunE_WithArgs(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createStatusCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git status with arguments
	expectedArgs := []string{"status", "--porcelain", "-b"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("## main...origin/main\n"), nil)

	// Execute the command with arguments
	err := cmd.RunE(cmd, []string{"--porcelain", "-b"})
	assert.NoError(t, err)
}

func TestStatusCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createStatusCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git status failed")

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(bytes.Buffer{}, expectedError)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestNewStatusCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewStatusCommand creates a command with the correct executor
	statusCmd := commands.NewStatusCommand(mockExecutor, mockGitHelper)

	// Verify the command can be created
	cmd := statusCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "status", cmd.Use)
}
