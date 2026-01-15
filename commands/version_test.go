package commands_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createVersionCommandWithMock(t *testing.T) (
	*mocks.MockRunner, *mocks.MockConfigManager, *gomock.Controller, *cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	versionCmd := commands.NewVersionCommand(mockRunner, mockConfigManager)
	cmd := versionCmd.Command()
	return mockRunner, mockConfigManager, ctrl, cmd
}

// setVersionCommandContext sets up the context with config for version command tests
func setVersionCommandContext(cmd *cobra.Command, version string) {
	cfg := config.NewConfigContext(&config.GlobalConfigStruct{
		Version: &config.Version{
			CurrentVersion: version,
		},
	}, nil)
	ctx := config.WithConfig(context.Background(), cfg)
	cmd.SetContext(ctx)
}

func TestVersionCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	assert.Equal(t, "version", cmd.Use)
	assert.Equal(t, "version of current build", cmd.Short)
}

func TestVersionCommand_RunE_NoFlags(t *testing.T) {
	_, _, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	setVersionCommandContext(cmd, "v1.0.0")

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestVersionCommand_RunE_WithLatestFlag(t *testing.T) {
	_, _, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	setVersionCommandContext(cmd, "v1.0.0")

	err := cmd.Flags().Set("latest", "true")
	require.NoError(t, err)

	_ = cmd.RunE(cmd, []string{})
}

func TestVersionCommand_RunE_ExecutorError(t *testing.T) {
	_, _, ctrl, cmd := createVersionCommandWithMock(t)
	defer ctrl.Finish()

	setVersionCommandContext(cmd, "v1.0.0")

	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestNewVersionCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)

	versionCmd := commands.NewVersionCommand(mockRunner, mockConfigManager)

	cmd := versionCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "version", cmd.Use)
}
