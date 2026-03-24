package stack_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands/stack"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/mocks"
	clientmocks "github.com/pavlovic265/265-gt/mocks/client"
	"github.com/stretchr/testify/assert"
)

func testCommandContext() context.Context {
	cfg := config.NewConfigContext(&config.GlobalConfigStruct{
		ActiveAccount: &config.Account{
			User:     "alice",
			Token:    "test-token",
			Platform: constants.GitHubPlatform,
		},
	}, nil)
	return config.WithConfig(context.Background(), cfg)
}

func TestSubmitCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	assert.Equal(t, "submit-stack", cmd.Use)
	assert.Equal(t, []string{"ss"}, cmd.Aliases)
	assert.Equal(t, "Push and create PRs for the entire stack", cmd.Short)
}

func TestNewSubmitCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "submit-stack", cmd.Use)
}

func TestSubmitCommand_HasDraftFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	draftFlag := cmd.Flags().Lookup("draft")
	assert.NotNil(t, draftFlag)
	assert.Equal(t, "d", draftFlag.Shorthand)
	assert.Equal(t, "false", draftFlag.DefValue)
}

func TestSubmitCommand_HasInteractiveFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	interactiveFlag := cmd.Flags().Lookup("interactive")
	assert.NotNil(t, interactiveFlag)
	assert.Equal(t, "i", interactiveFlag.Shorthand)
	assert.Equal(t, "false", interactiveFlag.DefValue)
}

func TestSubmitCommand_RunE_UpdatesBaseBranchForExistingPR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockCliClient := clientmocks.NewMockCliClient(ctrl)

	mockGitHelper.EXPECT().EnsureGitRepository().Return(nil)
	mockCliClient.EXPECT().
		ListPullRequests(gomock.Any(), []string{}).
		Return([]client.PullRequest{{Branch: "feature/test"}}, nil)
	mockGitHelper.EXPECT().GetCurrentBranch().Return("feature/test", nil)
	mockRunner.EXPECT().Git("checkout", "feature/test").Return(nil)
	mockRunner.EXPECT().Git("push", "--force", "origin", "feature/test").Return(nil)
	mockCliClient.EXPECT().
		UpdatePullRequestBaseBranch(gomock.Any(), "feature/test").
		Return(nil)
	mockGitHelper.EXPECT().GetChildren("feature/test").Return(nil)
	mockRunner.EXPECT().Git("checkout", "feature/test").Return(nil)

	cmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, mockCliClient).Command()
	cmd.SetContext(testCommandContext())

	if err := cmd.RunE(cmd, nil); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}

func TestSubmitCommand_RunE_CreatesPRWhenMissing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockCliClient := clientmocks.NewMockCliClient(ctrl)

	mockGitHelper.EXPECT().EnsureGitRepository().Return(nil)
	mockCliClient.EXPECT().
		ListPullRequests(gomock.Any(), []string{}).
		Return(nil, nil)
	mockGitHelper.EXPECT().GetCurrentBranch().Return("feature/test", nil)
	mockRunner.EXPECT().Git("checkout", "feature/test").Return(nil)
	mockRunner.EXPECT().Git("push", "--force", "origin", "feature/test").Return(nil)
	mockCliClient.EXPECT().
		CreatePullRequest(gomock.Any(), []string(nil)).
		Return(nil)
	mockGitHelper.EXPECT().GetChildren("feature/test").Return(nil)
	mockRunner.EXPECT().Git("checkout", "feature/test").Return(nil)

	cmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, mockCliClient).Command()
	cmd.SetContext(testCommandContext())

	if err := cmd.RunE(cmd, nil); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}
