package auth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAuthSwitchCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	switchCmd := auth.NewSwitchCommand(mockConfigManager)
	cmd := switchCmd.Command()

	assert.Equal(t, "switch", cmd.Use)
	assert.Equal(t, []string{"sw"}, cmd.Aliases)
	assert.Equal(t, "switch accounts", cmd.Short)
}

func TestNewAuthSwitchCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	switchCmd := auth.NewSwitchCommand(mockConfigManager)
	cmd := switchCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "switch", cmd.Use)
}
