package branch_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/branch"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTrackCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	trackCmd := branch.NewTrackCommand(mockRunner, mockGitHelper)
	cmd := trackCmd.Command()

	assert.Equal(t, "track", cmd.Use)
	assert.Equal(t, []string{"tr"}, cmd.Aliases)
	assert.Equal(t, "track existing branch", cmd.Short)
}

func TestNewTrackCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	trackCmd := branch.NewTrackCommand(mockRunner, mockGitHelper)
	cmd := trackCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "track", cmd.Use)
}
