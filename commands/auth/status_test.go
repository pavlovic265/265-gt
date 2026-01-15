package auth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/auth"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAuthStatusCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	statusCmd := auth.NewStatusCommand(mockConfigManager)
	cmd := statusCmd.Command()

	assert.Equal(t, "status", cmd.Use)
	assert.Equal(t, []string{"st"}, cmd.Aliases)
	assert.Equal(t, "see status of current auth user", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing)
}

func TestNewAuthStatusCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	statusCmd := auth.NewStatusCommand(mockConfigManager)
	cmd := statusCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "status", cmd.Use)
}
