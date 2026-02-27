package auth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogoutCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	logoutCmd := auth.NewLogoutCommand(mockConfigManager, cliClient)
	cmd := logoutCmd.Command()

	assert.Equal(t, "logout", cmd.Use)
	assert.Equal(t, []string{"lo"}, cmd.Aliases)
	assert.Equal(t, "logout user with token", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing)
}

func TestNewLogoutCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	logoutCmd := auth.NewLogoutCommand(mockConfigManager, cliClient)
	cmd := logoutCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "logout", cmd.Use)
}
