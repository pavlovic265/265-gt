package pr_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/pr"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	listCmd := pr.NewListCommand(mockRunner, mockConfigManager, mockGitHelper)
	cmd := listCmd.Command()

	assert.Equal(t, "list", cmd.Use)
	assert.Equal(t, []string{"li"}, cmd.Aliases)
	assert.Equal(t, "show list of pull request", cmd.Short)
}

func TestNewListCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	listCmd := pr.NewListCommand(mockRunner, mockConfigManager, mockGitHelper)
	cmd := listCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "list", cmd.Use)
}
