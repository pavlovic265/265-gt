package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func createCheckoutCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	checkoutCmd := commands.NewCheckoutCommand(mockExecutor, mockGitHelper)
	cmd := checkoutCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestCheckoutCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createCheckoutCommandWithMock(t)
	defer ctrl.Finish()

	assert.Equal(t, "checkout", cmd.Use)
	assert.Equal(t, []string{"co"}, cmd.Aliases)
	assert.Equal(t, "checkout branch", cmd.Short)
}

func TestNewCheckoutCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	checkoutCmd := commands.NewCheckoutCommand(mockExecutor, mockGitHelper)
	cmd := checkoutCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "checkout", cmd.Use)
}
