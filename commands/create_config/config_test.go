package createconfig_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	createconfig "github.com/pavlovic265/265-gt/commands/create_config"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestConfigCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	configCmd := createconfig.NewConfigCommand(mockExecutor, mockConfigManager)
	cmd := configCmd.Command()

	assert.Equal(t, "config", cmd.Use)
	assert.Equal(t, []string{"conf"}, cmd.Aliases)
	assert.Equal(t, "create config", cmd.Short)
}

func TestNewConfigCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	configCmd := createconfig.NewConfigCommand(mockExecutor, mockConfigManager)
	cmd := configCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "config", cmd.Use)
}
