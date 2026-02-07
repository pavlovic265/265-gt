package remote_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/remote"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func setCloneCommandContext(cmd *cobra.Command, activeAccount *config.Account) {
	cfg := config.NewConfigContext(&config.GlobalConfigStruct{
		ActiveAccount: activeAccount,
	}, nil)
	ctx := config.WithConfig(context.Background(), cfg)
	cmd.SetContext(ctx)
}

func TestCloneCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()

	assert.Equal(t, "clone <repo>", cmd.Use)
	assert.Equal(t, "Clone a repository using active account's SSH config", cmd.Short)
	assert.NotNil(t, cmd.RunE)
}

func TestNewCloneCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "clone <repo>", cmd.Use)
}

func TestCloneCommand_NoContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()

	err := cmd.RunE(cmd, []string{"owner/repo"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "config not loaded")
}

func TestCloneCommand_NoActiveAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()
	setCloneCommandContext(cmd, nil)

	err := cmd.RunE(cmd, []string{"owner/repo"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no active account")
}

func TestCloneCommand_SimpleRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	account := &config.Account{
		User:     "testuser",
		Platform: constants.GitHubPlatform,
		SSHHost:  "github.com-testuser",
	}

	mockRunner.EXPECT().
		Git("clone", "git@github.com-testuser:owner/repo.git").
		Return(nil).
		Times(1)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()
	setCloneCommandContext(cmd, account)

	err := cmd.RunE(cmd, []string{"owner/repo"})
	assert.NoError(t, err)
}

func TestCloneCommand_NoSSHHost_GitHub(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	account := &config.Account{
		User:     "testuser",
		Platform: constants.GitHubPlatform,
	}

	mockRunner.EXPECT().
		Git("clone", "git@github.com:owner/repo.git").
		Return(nil).
		Times(1)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()
	setCloneCommandContext(cmd, account)

	err := cmd.RunE(cmd, []string{"owner/repo"})
	assert.NoError(t, err)
}

func TestCloneCommand_NoSSHHost_GitLab(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	account := &config.Account{
		User:     "testuser",
		Platform: constants.GitLabPlatform,
	}

	mockRunner.EXPECT().
		Git("clone", "git@gitlab.com:owner/repo.git").
		Return(nil).
		Times(1)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()
	setCloneCommandContext(cmd, account)

	err := cmd.RunE(cmd, []string{"owner/repo"})
	assert.NoError(t, err)
}

func TestCloneCommand_FullSSHURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	account := &config.Account{
		User:     "testuser",
		Platform: constants.GitHubPlatform,
		SSHHost:  "github.com-testuser",
	}

	mockRunner.EXPECT().
		Git("clone", "git@github.com:other/repo.git").
		Return(nil).
		Times(1)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()
	setCloneCommandContext(cmd, account)

	err := cmd.RunE(cmd, []string{"git@github.com:other/repo.git"})
	assert.NoError(t, err)
}

func TestCloneCommand_HTTPSUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	account := &config.Account{
		User:     "testuser",
		Platform: constants.GitHubPlatform,
		SSHHost:  "github.com-testuser",
	}

	mockRunner.EXPECT().
		Git("clone", "git@github.com-testuser:owner/repo.git").
		Return(nil).
		Times(1)

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()
	setCloneCommandContext(cmd, account)

	err := cmd.RunE(cmd, []string{"https://github.com/owner/repo"})
	assert.NoError(t, err)
}

func TestCloneCommand_InvalidFormat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	account := &config.Account{
		User:     "testuser",
		Platform: constants.GitHubPlatform,
	}

	cloneCmd := remote.NewCloneCommand(mockRunner, mockGitHelper)
	cmd := cloneCmd.Command()
	setCloneCommandContext(cmd, account)

	err := cmd.RunE(cmd, []string{"invalid-format"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid repository format")
}
