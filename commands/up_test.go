package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	upCmd := commands.NewUpCommand(mockRunner, mockGitHelper)
	cmd := upCmd.Command()

	assert.Equal(t, "up", cmd.Use)
	assert.Equal(t, "move to branch up in stack", cmd.Short)
}

func TestNewUpCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	upCmd := commands.NewUpCommand(mockRunner, mockGitHelper)
	cmd := upCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "up", cmd.Use)
}
