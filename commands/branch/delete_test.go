package branch_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/branch"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	deleteCmd := branch.NewDeleteCommand(mockRunner, mockGitHelper)
	cmd := deleteCmd.Command()

	assert.Equal(t, "delete", cmd.Use)
	assert.Equal(t, []string{"dl"}, cmd.Aliases)
	assert.Equal(t, "delete branch", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing)
}

func TestNewDeleteCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	deleteCmd := branch.NewDeleteCommand(mockRunner, mockGitHelper)
	cmd := deleteCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "delete", cmd.Use)
}
