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

func createContCommandWithMock(
	t *testing.T,
) (*mocks.MockExecutor, *mocks.MockGitHelper, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	contCmd := commands.NewContCommand(mockExecutor, mockGitHelper)
	cmd := contCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestContCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createContCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "cont", cmd.Use)
	assert.Equal(t, "short for rebase --continue", cmd.Short)
}

func TestContCommand_RunE_Success(t *testing.T) {
	mockExecutor, mockGitHelper, ctrl, cmd := createContCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git rebase --continue
	expectedArgs := []string{"rebase", "--continue"}

	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs(expectedArgs).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		Run().
		Return(nil)

	mockGitHelper.EXPECT().
		IsRebaseInProgress().
		Return(false)

	mockGitHelper.EXPECT().
		GetPending(gomock.Any()).
		Return("", errors.New("no pending branch")).
		Times(2)

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestContCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createContCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git rebase failed")

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
	assert.Contains(t, err.Error(), "Failed to continue rebase")
	assert.Contains(t, err.Error(), "git rebase failed")
}

func TestNewContCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewContCommand creates a command with the correct executor
	contCmd := commands.NewContCommand(mockExecutor, mockGitHelper)

	// Verify the command can be created
	cmd := contCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "cont", cmd.Use)
}
