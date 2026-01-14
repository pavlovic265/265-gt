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

func createAddCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	addCmd := commands.NewAddCommand(mockExecutor, mockGitHelper)
	cmd := addCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestAddCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "add", cmd.Use)
	assert.Equal(t, []string{"a"}, cmd.Aliases)
	assert.Equal(t, "git add", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing, "Expected DisableFlagParsing to be true")
}

func TestAddCommand_RunE_NoArgs(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git add with no arguments
	expectedArgs := []string{"add"}

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

func TestAddCommand_RunE_WithSingleFile(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git add with single file
	expectedArgs := []string{"add", "file.txt"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with single file
	err := cmd.RunE(cmd, []string{"file.txt"})
	assert.NoError(t, err)
}

func TestAddCommand_RunE_WithMultipleFiles(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git add with multiple files
	expectedArgs := []string{"add", "file1.txt", "file2.txt", "file3.txt"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with multiple files
	err := cmd.RunE(cmd, []string{"file1.txt", "file2.txt", "file3.txt"})
	assert.NoError(t, err)
}

func TestAddCommand_RunE_WithGitOptions(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git add with options
	expectedArgs := []string{"add", "-A", "--verbose"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with git options
	err := cmd.RunE(cmd, []string{"-A", "--verbose"})
	assert.NoError(t, err)
}

func TestAddCommand_RunE_WithMixedArgs(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git add with mixed arguments
	expectedArgs := []string{"add", "-A", "file.txt", "--verbose", "*.go"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	// Execute the command with mixed arguments
	err := cmd.RunE(cmd, []string{"-A", "file.txt", "--verbose", "*.go"})
	assert.NoError(t, err)
}

func TestAddCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git add failed")

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
	err := cmd.RunE(cmd, []string{"file.txt"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed to stage files")
	assert.Contains(t, err.Error(), "git add failed")
}

func TestAddCommand_RunE_ArgsPassing(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Test that all arguments are properly passed through
	testCases := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			name:     "no args",
			args:     []string{},
			expected: []string{"add"},
		},
		{
			name:     "single file",
			args:     []string{"file.txt"},
			expected: []string{"add", "file.txt"},
		},
		{
			name:     "multiple files",
			args:     []string{"file1.txt", "file2.txt"},
			expected: []string{"add", "file1.txt", "file2.txt"},
		},
		{
			name:     "with flags",
			args:     []string{"-A", "--verbose"},
			expected: []string{"add", "-A", "--verbose"},
		},
		{
			name:     "with patterns",
			args:     []string{"*.go", "*.txt"},
			expected: []string{"add", "*.go", "*.txt"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockExecutor.EXPECT().
				WithGit().
				Return(mockExecutor)

			mockExecutor.EXPECT().
				WithArgs(tc.expected).
				Return(mockExecutor)

			mockExecutor.EXPECT().
				Run().
				Return(nil)

			err := cmd.RunE(cmd, tc.args)
			assert.NoError(t, err)
		})
	}
}

func TestAddCommand_DisableFlagParsing(t *testing.T) {
	_, _, ctrl, cmd := createAddCommandWithMock(t)
	defer ctrl.Finish()

	// Test that DisableFlagParsing is set to true
	// This is important because we want to pass all arguments directly to git
	assert.True(t, cmd.DisableFlagParsing, "DisableFlagParsing should be true to pass all args to git")
}

func TestNewAddCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewAddCommand creates a command with the correct executor
	addCmd := commands.NewAddCommand(mockExecutor, mockGitHelper)

	// Verify the command can be created
	cmd := addCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "add", cmd.Use)
}
