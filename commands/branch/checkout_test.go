package branch_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/commands/branch"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	checkoutCmd := branch.NewCheckoutCommand(mockRunner, mockGitHelper)
	cmd := checkoutCmd.Command()

	assert.Equal(t, "checkout", cmd.Use)
	assert.Equal(t, []string{"co"}, cmd.Aliases)
	assert.Equal(t, "checkout branch", cmd.Short)
}

func TestNewCheckoutCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)

	checkoutCmd := branch.NewCheckoutCommand(mockRunner, mockGitHelper)
	cmd := checkoutCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "checkout", cmd.Use)
}
