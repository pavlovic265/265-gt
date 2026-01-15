package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMoveCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	moveCmd := commands.NewMoveCommand(mockRunner, mockGitHelper)
	cmd := moveCmd.Command()

	assert.Equal(t, "move", cmd.Use)
	assert.Equal(t, []string{"mo"}, cmd.Aliases)
	assert.Equal(t, "rebase branch onto other branch", cmd.Short)
}

func TestNewMoveCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	moveCmd := commands.NewMoveCommand(mockRunner, mockGitHelper)
	cmd := moveCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "move", cmd.Use)
}
