package stack_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/stack"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStackCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	stackCmd := stack.NewStackCommand(mockRunner, mockGitHelper)
	cmd := stackCmd.Command()

	assert.Equal(t, "stack", cmd.Use)
	assert.Equal(t, []string{"s"}, cmd.Aliases)
	assert.Equal(t, "stack management commands", cmd.Short)
}

func TestNewStackCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	stackCmd := stack.NewStackCommand(mockRunner, mockGitHelper)
	cmd := stackCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "stack", cmd.Use)
}

func TestStackCommand_HasSubcommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	stackCmd := stack.NewStackCommand(mockRunner, mockGitHelper)
	cmd := stackCmd.Command()

	assert.True(t, cmd.HasSubCommands())
	assert.Len(t, cmd.Commands(), 1)
	assert.Equal(t, "restack", cmd.Commands()[0].Use)
}
