package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func createDownCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	downCmd := commands.NewDownCommand(mockExecutor, mockGitHelper)
	cmd := downCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestDownCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createDownCommandWithMock(t)
	defer ctrl.Finish()

	assert.Equal(t, "down", cmd.Use)
	assert.Equal(t, "move to brunch down in stack", cmd.Short)
}

func TestNewDownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	downCmd := commands.NewDownCommand(mockExecutor, mockGitHelper)
	cmd := downCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "down", cmd.Use)
}
