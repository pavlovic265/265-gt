package basic_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/basic"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUnstageCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	unstageCmd := basic.NewUnstageCommand(mockRunner, mockGitHelper)
	cmd := unstageCmd.Command()

	assert.Equal(t, "unstage", cmd.Use)
	assert.Equal(t, []string{"us"}, cmd.Aliases)
	assert.Equal(t, "unstage files", cmd.Short)
}

func TestNewUnstageCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	unstageCmd := basic.NewUnstageCommand(mockRunner, mockGitHelper)
	cmd := unstageCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "unstage", cmd.Use)
}
