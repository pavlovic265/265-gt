package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDownCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	downCmd := commands.NewDownCommand(mockRunner, mockGitHelper)
	cmd := downCmd.Command()

	assert.Equal(t, "down", cmd.Use)
	assert.Equal(t, "move to branch down in stack", cmd.Short)
}

func TestNewDownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	downCmd := commands.NewDownCommand(mockRunner, mockGitHelper)
	cmd := downCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "down", cmd.Use)
}
