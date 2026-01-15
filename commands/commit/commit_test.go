package commit_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/commit"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCommitCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	commitCmd := commit.NewCommitCommand(mockRunner, mockGitHelper)
	cmd := commitCmd.Command()

	assert.Equal(t, "commit", cmd.Use)
	assert.Equal(t, []string{"cm"}, cmd.Aliases)
	assert.Equal(t, "create commit", cmd.Short)
}

func TestCommitCommand_Flags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	commitCmd := commit.NewCommitCommand(mockRunner, mockGitHelper)
	cmd := commitCmd.Command()

	emptyFlag := cmd.Flags().Lookup("empty")
	assert.NotNil(t, emptyFlag)
	assert.Equal(t, "e", emptyFlag.Shorthand)
	assert.Equal(t, "false", emptyFlag.DefValue)
}

func TestNewCommitCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	commitCmd := commit.NewCommitCommand(mockRunner, mockGitHelper)
	cmd := commitCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "commit", cmd.Use)
}
