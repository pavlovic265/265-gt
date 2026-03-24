package remote_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/remote"
	"github.com/pavlovic265/265-gt/mocks"
	clientmocks "github.com/pavlovic265/265-gt/mocks/client"
	"github.com/stretchr/testify/assert"
)

func TestPushCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockCliClient := clientmocks.NewMockCliClient(ctrl)

	pushCmd := remote.NewPushCommand(mockRunner, mockGitHelper, mockCliClient)
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
	mockCliClient := clientmocks.NewMockCliClient(ctrl)

	pushCmd := remote.NewPushCommand(mockRunner, mockGitHelper, mockCliClient)
	cmd := pushCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "push", cmd.Use)
}

func TestPushCommand_RunE_UpdatesOpenPullRequestBaseBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockCliClient := clientmocks.NewMockCliClient(ctrl)

	mockGitHelper.EXPECT().EnsureGitRepository().Return(nil)
	mockGitHelper.EXPECT().GetCurrentBranch().Return("feature/test", nil)
	mockRunner.EXPECT().Git("push", "--force", "origin", "feature/test").Return(nil)
	mockCliClient.EXPECT().
		HasOpenPullRequestForBranch(gomock.Any(), "feature/test").
		Return(true, nil)
	mockCliClient.EXPECT().
		UpdatePullRequestBaseBranch(gomock.Any(), "feature/test").
		Return(nil)

	cmd := remote.NewPushCommand(mockRunner, mockGitHelper, mockCliClient).Command()
	cmd.SetContext(context.Background())

	if err := cmd.RunE(cmd, nil); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}

func TestPushCommand_RunE_SkipsBaseBranchUpdateWhenNoOpenPullRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockCliClient := clientmocks.NewMockCliClient(ctrl)

	mockGitHelper.EXPECT().EnsureGitRepository().Return(nil)
	mockGitHelper.EXPECT().GetCurrentBranch().Return("feature/test", nil)
	mockRunner.EXPECT().Git("push", "--force", "origin", "feature/test").Return(nil)
	mockCliClient.EXPECT().
		HasOpenPullRequestForBranch(gomock.Any(), "feature/test").
		Return(false, nil)

	cmd := remote.NewPushCommand(mockRunner, mockGitHelper, mockCliClient).Command()
	cmd.SetContext(context.Background())

	if err := cmd.RunE(cmd, nil); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}
