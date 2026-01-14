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

func createSwitchCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	switchCmd := commands.NewSwitchCommand(mockExecutor, mockGitHelper)
	cmd := switchCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestSwitchCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createSwitchCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "switch", cmd.Use)
	assert.Equal(t, []string{"sw"}, cmd.Aliases)
	assert.Equal(t, "switch back to previous branch", cmd.Short)
}

func TestSwitchCommand_RunE_Success(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createSwitchCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for git checkout -
	expectedArgs := []string{"checkout", "-"}

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

func TestSwitchCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, _, ctrl, cmd := createSwitchCommandWithMock(t)
	defer ctrl.Finish()

	expectedError := errors.New("git checkout failed")

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
	assert.Contains(t, err.Error(), "Failed to switch to previous branch")
	assert.Contains(t, err.Error(), "git checkout failed")
}

func TestNewSwitchCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	// Test that NewSwitchCommand creates a command with the correct executor
	switchCmd := commands.NewSwitchCommand(mockExecutor, mockGitHelper)

	// Verify the command can be created
	cmd := switchCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "switch", cmd.Use)
}
