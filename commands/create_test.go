package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func createCreateCommandWithMock(t *testing.T) (
	*mocks.MockExecutor,
	*mocks.MockGitHelper,
	*gomock.Controller,
	*cobra.Command,
) {
	ctrl := gomock.NewController(t)
	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	createCmd := commands.NewCreateCommand(mockExecutor, mockGitHelper)
	cmd := createCmd.Command()
	return mockExecutor, mockGitHelper, ctrl, cmd
}

func TestCreateCommand_Command(t *testing.T) {
	_, _, ctrl, cmd := createCreateCommandWithMock(t)
	defer ctrl.Finish()

	assert.Equal(t, "create", cmd.Use)
	assert.Equal(t, []string{"c"}, cmd.Aliases)
	assert.Equal(t, "create branch", cmd.Short)
}

func TestNewCreateCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	createCmd := commands.NewCreateCommand(mockExecutor, mockGitHelper)
	cmd := createCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "create", cmd.Use)
}
