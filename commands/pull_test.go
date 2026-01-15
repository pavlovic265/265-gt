package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPullCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	pullCmd := commands.NewPullCommand(mockRunner, mockGitHelper)
	cmd := pullCmd.Command()

	assert.Equal(t, "pull", cmd.Use)
	assert.Equal(t, []string{"pl"}, cmd.Aliases)
	assert.Equal(t, "pull branch", cmd.Short)
}

func TestNewPullCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	pullCmd := commands.NewPullCommand(mockRunner, mockGitHelper)
	cmd := pullCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "pull", cmd.Use)
}
