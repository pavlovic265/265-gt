package helpers

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestValidateBranchName_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	mockRunner.EXPECT().
		GitOutput("check-ref-format", "--branch", "feature/my-branch").
		Return("feature/my-branch", nil)

	err := gitHelper.ValidateBranchName("feature/my-branch")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestValidateBranchName_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	err := gitHelper.ValidateBranchName("")
	if err == nil {
		t.Error("expected error for empty branch name")
	}
}

func TestValidateBranchName_Invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	mockRunner.EXPECT().
		GitOutput("check-ref-format", "--branch", "branch with spaces").
		Return("", errors.New("invalid ref format"))

	err := gitHelper.ValidateBranchName("branch with spaces")
	if err == nil {
		t.Error("expected error for invalid branch name")
	}
}
