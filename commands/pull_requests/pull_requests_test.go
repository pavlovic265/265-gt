package pullrequests_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	pullrequests "github.com/pavlovic265/265-gt/commands/pull_requests"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPullRequestCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	prCmd := pullrequests.NewPullRequestCommand(mockExecutor, mockConfigManager, mockGitHelper)
	cmd := prCmd.Command()

	assert.Equal(t, "pull_request", cmd.Use)
	assert.Equal(t, []string{"pr"}, cmd.Aliases)
	assert.Equal(t, "commands for pull request", cmd.Short)
	assert.NotNil(t, cmd.PersistentPreRun)
}

func TestNewPullRequestCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	prCmd := pullrequests.NewPullRequestCommand(mockExecutor, mockConfigManager, mockGitHelper)
	cmd := prCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "pull_request", cmd.Use)
}
