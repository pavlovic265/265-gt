package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func createMoveCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	moveCmd := commands.NewMoveCommand(mockExecutor, mockGitHelper)
	cmd := moveCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestMoveCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createMoveCommandWithMock(t)
	defer ctrl.Finish()

	assert.Equal(t, "move", cmd.Use)
	assert.Equal(t, []string{"mo"}, cmd.Aliases)
	assert.Equal(t, "rebase branch onto other branch", cmd.Short)
}

func TestNewMoveCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	moveCmd := commands.NewMoveCommand(mockExecutor, mockGitHelper)
	cmd := moveCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "move", cmd.Use)
}
