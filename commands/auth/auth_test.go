package auth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAuthCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	authCmd := auth.NewAuthCommand(mockConfigManager, cliClient)
	cmd := authCmd.Command()

	assert.Equal(t, "auth", cmd.Use)
	assert.Equal(t, "auth user", cmd.Short)
}

func TestNewAuthCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	authCmd := auth.NewAuthCommand(mockConfigManager, cliClient)
	cmd := authCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "auth", cmd.Use)
}
