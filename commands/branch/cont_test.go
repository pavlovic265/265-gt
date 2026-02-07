package branch_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/branch"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestContCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	contCmd := branch.NewContCommand(mockRunner, mockGitHelper)
	cmd := contCmd.Command()

	assert.Equal(t, "cont", cmd.Use)
	assert.Equal(t, "short for rebase --continue", cmd.Short)
}

func TestNewContCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	contCmd := branch.NewContCommand(mockRunner, mockGitHelper)
	cmd := contCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "cont", cmd.Use)
}
