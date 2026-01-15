package auth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLoginCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	loginCmd := auth.NewLoginCommand(mockConfigManager)
	cmd := loginCmd.Command()

	assert.Equal(t, "login", cmd.Use)
	assert.Equal(t, []string{"lg"}, cmd.Aliases)
	assert.Equal(t, "login user with token", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing)
}

func TestNewLoginCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	loginCmd := auth.NewLoginCommand(mockConfigManager)
	cmd := loginCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "login", cmd.Use)
}
