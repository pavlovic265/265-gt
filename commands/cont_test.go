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

// Test helper to create a cont command with mock executor
func createContCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	contCmd := commands.NewContCommand(mockExecutor)
	cmd := contCmd.Command()
	return mockExecutor, ctrl, cmd
}

func TestContCommand_Command(t *testing.T) {
	_, ctrl, cmd := createContCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "cont", cmd.Use)
	assert.Equal(t, "short for rebase --continue", cmd.Short)
}

func TestContCommand_RunE_Success(t *testing.T) {
	mockExecutor, ctrl, cmd := createContCommandWithMock(t)
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

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestContCommand_RunE_ExecutorError(t *testing.T) {
	mockExecutor, ctrl, cmd := createContCommandWithMock(t)
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
	assert.Equal(t, expectedError, err)
}

func TestNewContCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test that NewContCommand creates a command with the correct executor
	contCmd := commands.NewContCommand(mockExecutor)

	// Verify the command can be created
	cmd := contCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "cont", cmd.Use)
}
