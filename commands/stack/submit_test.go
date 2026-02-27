package stack_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands/stack"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSubmitCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	assert.Equal(t, "submit-stack", cmd.Use)
	assert.Equal(t, []string{"ss"}, cmd.Aliases)
	assert.Equal(t, "Push and create PRs for the entire stack", cmd.Short)
}

func TestNewSubmitCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "submit-stack", cmd.Use)
}

func TestSubmitCommand_HasDraftFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	draftFlag := cmd.Flags().Lookup("draft")
	assert.NotNil(t, draftFlag)
	assert.Equal(t, "d", draftFlag.Shorthand)
	assert.Equal(t, "false", draftFlag.DefValue)
}

func TestSubmitCommand_HasInteractiveFlag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	submitCmd := stack.NewSubmitCommand(mockRunner, mockGitHelper, cliClient)
	cmd := submitCmd.Command()

	interactiveFlag := cmd.Flags().Lookup("interactive")
	assert.NotNil(t, interactiveFlag)
	assert.Equal(t, "i", interactiveFlag.Shorthand)
	assert.Equal(t, "false", interactiveFlag.DefValue)
}
