package commands_test

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test helper to create a clean command with mock executor
func createCleanCommandWithMock(t *testing.T) (*mocks.MockExecutor, *gomock.Controller, *cobra.Command) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	cleanCmd := commands.NewCleanCommand(mockExecutor)
	cmd := cleanCmd.Command()
	return mockExecutor, ctrl, cmd
}

func TestCleanCommand_Command(t *testing.T) {
	_, ctrl, cmd := createCleanCommandWithMock(t)
	defer ctrl.Finish()

	// Test that the command is properly configured
	assert.Equal(t, "clean", cmd.Use)
	assert.Equal(t, []string{"cl"}, cmd.Aliases)
	assert.Equal(t, "Clean up branches interactively", cmd.Short)
	assert.Equal(t, "Clean up branches one by one with confirmation. "+
		"Protected branches and current branch are skipped.", cmd.Long)
}

func TestCleanCommand_RunE_NoBranchesToClean(t *testing.T) {
	mockExecutor, ctrl, cmd := createCleanCommandWithMock(t)
	defer ctrl.Finish()

	// Set up expectations for GetCurrentBranchName
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("main\n"), nil)

	// Set up expectations for GetBranches
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "--list"}).
		Return(mockExecutor)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(*bytes.NewBufferString("* main\n"), nil)

	// Execute the command - should show "No branches to clean up!"
	err := cmd.RunE(cmd, []string{})
	assert.NoError(t, err)
}

func TestNewCleanCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test that NewCleanCommand creates a command with the correct executor
	cleanCmd := commands.NewCleanCommand(mockExecutor)

	// Verify the command can be created
	cmd := cleanCmd.Command()
	require.NotNil(t, cmd)
	assert.Equal(t, "clean", cmd.Use)
}
