package stack_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/stack"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRestackCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	restackCmd := stack.NewRestackCommand(mockRunner, mockGitHelper)
	cmd := restackCmd.Command()

	assert.Equal(t, "restack", cmd.Use)
	assert.Equal(t, []string{"rs"}, cmd.Aliases)
	assert.Equal(t, "Restack branches", cmd.Short)
}

func TestNewRestackCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	restackCmd := stack.NewRestackCommand(mockRunner, mockGitHelper)
	cmd := restackCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "restack", cmd.Use)
}
