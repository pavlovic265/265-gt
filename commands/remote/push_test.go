package remote_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/remote"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPushCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	pushCmd := remote.NewPushCommand(mockRunner, mockGitHelper)
	cmd := pushCmd.Command()

	assert.Equal(t, "push", cmd.Use)
	assert.Equal(t, []string{"pu"}, cmd.Aliases)
	assert.Equal(t, "push branch always force", cmd.Short)
}

func TestNewPushCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	pushCmd := remote.NewPushCommand(mockRunner, mockGitHelper)
	cmd := pushCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "push", cmd.Use)
}
