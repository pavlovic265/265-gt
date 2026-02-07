package basic_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/basic"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	addCmd := basic.NewAddCommand(mockRunner, mockGitHelper)
	cmd := addCmd.Command()

	assert.Equal(t, "add", cmd.Use)
	assert.Equal(t, []string{"a"}, cmd.Aliases)
	assert.Equal(t, "git add", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing)
}

func TestNewAddCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	addCmd := basic.NewAddCommand(mockRunner, mockGitHelper)
	cmd := addCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "add", cmd.Use)
}
