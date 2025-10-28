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

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	stackCmd := stack.NewStackCommand(mockExecutor, mockGitHelper)
	cmd := stackCmd.Command()

	assert.Equal(t, "stack", cmd.Use)
	assert.Equal(t, []string{"s"}, cmd.Aliases)
	assert.Equal(t, "commands for pull request", cmd.Short)
	assert.NotNil(t, cmd.PersistentPreRun)
}

func TestNewStackCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	stackCmd := stack.NewStackCommand(mockExecutor, mockGitHelper)
	cmd := stackCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "stack", cmd.Use)
}
