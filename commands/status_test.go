package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStatusCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	statusCmd := commands.NewStatusCommand(mockRunner, mockGitHelper)
	cmd := statusCmd.Command()

	assert.Equal(t, "status", cmd.Use)
	assert.Equal(t, []string{"st"}, cmd.Aliases)
	assert.Equal(t, "git status", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing)
	assert.True(t, cmd.SilenceUsage)
}

func TestNewStatusCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	statusCmd := commands.NewStatusCommand(mockRunner, mockGitHelper)
	cmd := statusCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "status", cmd.Use)
}
