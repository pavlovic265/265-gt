package commit

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test helper to create a commit command with mock executor
func createCommitCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, commitCommand) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	commitCmd := NewCommitCommand(mockExecutor)
	return mockExecutor, ctrl, commitCmd
}

func TestCommitCommand_Command(t *testing.T) {
	_, ctrl, commitCmd := createCommitCommandWithMock(t)
	defer ctrl.Finish()

	cmd := commitCmd.Command()

	// Test that the command is properly configured
	assert.Equal(t, "commit", cmd.Use)
	assert.Equal(t, []string{"cm"}, cmd.Aliases)
	assert.Equal(t, "create commit", cmd.Short)
}

func TestCommitCommand_RunE_WithMessage(t *testing.T) {
	mockExecutor, ctrl, commitCmd := createCommitCommandWithMock(t)
	defer ctrl.Finish()

	cmd := commitCmd.Command()

	// Set up expectations
	expectedArgs := []string{"commit", "-m", "test message"}

	// Create a chain of mock calls
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
	err := cmd.RunE(cmd, []string{"test message"})
	assert.NoError(t, err)
}

func TestCommitCommand_RunE_WithEmptyFlag(t *testing.T) {
	mockExecutor, ctrl, commitCmd := createCommitCommandWithMock(t)
	defer ctrl.Finish()

	cmd := commitCmd.Command()

	// Set the empty flag
	err := cmd.Flags().Set("empty", "true")
	require.NoError(t, err)

	// Set up expectations for empty commit
	// Note: The exeArgs contains the original message, not the modified one
	expectedArgs := []string{"commit", "--allow-empty", "-m", "test message"}

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
	err = cmd.RunE(cmd, []string{"test message"})
	assert.NoError(t, err)
}

func TestCommitCommand_RunE_WithEmptyFlag_NoMessage(t *testing.T) {
	mockExecutor, ctrl, commitCmd := createCommitCommandWithMock(t)
	defer ctrl.Finish()

	cmd := commitCmd.Command()

	// Set the empty flag
	err := cmd.Flags().Set("empty", "true")
	require.NoError(t, err)

	// Set up expectations for empty commit with timestamp message
	// The message will be a timestamp, so we'll use gomock.Any() for the args
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		DoAndReturn(func(args []string) executor.Executor {
			// Verify the args structure
			assert.Equal(t, 4, len(args), "Expected 4 args for empty commit")
			assert.Equal(t, "commit", args[0])
			assert.Equal(t, "--allow-empty", args[1])
			assert.Equal(t, "-m", args[2])
			// The fourth arg should be a timestamp, so we just verify it's not empty
			assert.NotEmpty(t, args[3], "Expected timestamp message")
			return mockExecutor
		})

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command without arguments (will use timestamp)
	err = cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestCommitCommand_RunE_NoMessage(t *testing.T) {
	mockExecutor, ctrl, commitCmd := createCommitCommandWithMock(t)
	defer ctrl.Finish()

	cmd := commitCmd.Command()

	// Set up expectations for commit with timestamp message
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(gomock.Any()).
		DoAndReturn(func(args []string) executor.Executor {
			// Verify the args structure
			assert.Equal(t, 3, len(args), "Expected 3 args for regular commit")
			assert.Equal(t, "commit", args[0])
			assert.Equal(t, "-m", args[1])
			// The third arg should be a timestamp, so we just verify it's not empty
			assert.NotEmpty(t, args[2], "Expected timestamp message")
			return mockExecutor
		})

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command without arguments (will use timestamp)
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestCommitCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, ctrl, commitCmd := createCommitCommandWithMock(t)
	defer ctrl.Finish()

	cmd := commitCmd.Command()

	expectedError := errors.New("git commit failed")

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
	err := cmd.RunE(cmd, []string{"test message"})
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestCommitCommand_EmptyFlag(t *testing.T) {
	_, ctrl, commitCmd := createCommitCommandWithMock(t)
	defer ctrl.Finish()

	cmd := commitCmd.Command()

	// Test that the empty flag exists and has correct properties
	emptyFlag := cmd.Flags().Lookup("empty")
	require.NotNil(t, emptyFlag, "Expected 'empty' flag to exist")
	assert.Equal(t, "empty", emptyFlag.Name)
	assert.Equal(t, "e", emptyFlag.Shorthand)
	assert.Equal(t, "false", emptyFlag.DefValue)
}
