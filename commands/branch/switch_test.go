package branch_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/branch"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSwitchCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	switchCmd := branch.NewSwitchCommand(mockRunner, mockGitHelper)
	cmd := switchCmd.Command()

	assert.Equal(t, "switch", cmd.Use)
	assert.Equal(t, []string{"sw"}, cmd.Aliases)
	assert.Equal(t, "switch back to previous branch", cmd.Short)
}

func TestNewSwitchCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	switchCmd := branch.NewSwitchCommand(mockRunner, mockGitHelper)
	cmd := switchCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "switch", cmd.Use)
}
