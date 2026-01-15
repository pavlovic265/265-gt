package commands_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	createCmd := commands.NewCreateCommand(mockRunner, mockGitHelper)
	cmd := createCmd.Command()

	assert.Equal(t, "create", cmd.Use)
	assert.Equal(t, []string{"c"}, cmd.Aliases)
	assert.Equal(t, "create branch", cmd.Short)
}

func TestNewCreateCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	createCmd := commands.NewCreateCommand(mockRunner, mockGitHelper)
	cmd := createCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "create", cmd.Use)
}
