package utility_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/utility"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpgradeCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	upgradeCmd := utility.NewUpgradeCommand(mockRunner, mockConfigManager)
	cmd := upgradeCmd.Command()

	assert.Equal(t, "upgrade", cmd.Use)
	assert.Equal(t, "upgrade of current build", cmd.Short)
}

func TestNewUpgradeCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	upgradeCmd := utility.NewUpgradeCommand(mockRunner, mockConfigManager)
	cmd := upgradeCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "upgrade", cmd.Use)
}
