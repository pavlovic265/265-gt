package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCleanCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cleanCmd := commands.NewCleanCommand(mockRunner, mockGitHelper)
	cmd := cleanCmd.Command()

	assert.Equal(t, "clean", cmd.Use)
	assert.Equal(t, []string{"cl"}, cmd.Aliases)
	assert.Equal(t, "Clean up branches interactively", cmd.Short)
}

func TestNewCleanCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cleanCmd := commands.NewCleanCommand(mockRunner, mockGitHelper)
	cmd := cleanCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "clean", cmd.Use)
}
