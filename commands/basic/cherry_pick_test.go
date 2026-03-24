package basic_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/basic"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCherryPickCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cherryPickCmd := basic.NewCherryPickCommand(mockRunner, mockGitHelper)
	cmd := cherryPickCmd.Command()

	assert.Equal(t, "cherry-pick", cmd.Use)
	assert.Equal(t, []string{"cp"}, cmd.Aliases)
	assert.Equal(t, "git cherry-pick", cmd.Short)
	assert.True(t, cmd.DisableFlagParsing)
}

func TestNewCherryPickCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	cherryPickCmd := basic.NewCherryPickCommand(mockRunner, mockGitHelper)
	cmd := cherryPickCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "cherry-pick", cmd.Use)
}

func TestCherryPickCommand_RunE(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	mockGitHelper.EXPECT().EnsureGitRepository().Return(nil)
	mockRunner.EXPECT().Git("cherry-pick", "abc123").Return(nil)

	cmd := basic.NewCherryPickCommand(mockRunner, mockGitHelper).Command()
	cmd.SetContext(context.Background())

	err := cmd.RunE(cmd, []string{"abc123"})
	assert.NoError(t, err)
}
