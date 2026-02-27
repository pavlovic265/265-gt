package githelper

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestGetCurrentBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	expectedBranch := "main"

	mockRunner.EXPECT().
		GitOutput("rev-parse", "--abbrev-ref", "HEAD").
		Return(expectedBranch, nil).
		Times(1)

	result, err := gitHelper.GetCurrentBranch()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expectedBranch {
		t.Errorf("Expected '%s', got '%s'", expectedBranch, result)
	}
}

func TestGetCurrentBranch_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	expectedError := errors.New("git error")

	mockRunner.EXPECT().
		GitOutput("rev-parse", "--abbrev-ref", "HEAD").
		Return("", expectedError).
		Times(1)

	result, err := gitHelper.GetCurrentBranch()

	if err == nil {
		t.Error("Expected error, got nil")
	}
	if result != "" {
		t.Error("Expected empty string result on error")
	}
}

func TestGetBranches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	branchOutput := `* main
  feature1
  feature2
  develop`

	expectedBranches := []string{"main", "feature1", "feature2", "develop"}

	mockRunner.EXPECT().
		GitOutput("branch", "--list").
		Return(branchOutput, nil).
		Times(1)

	result, err := gitHelper.GetBranches()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(result) != len(expectedBranches) {
		t.Errorf("Expected %d branches, got %d", len(expectedBranches), len(result))
	}
	for i, expected := range expectedBranches {
		if result[i] != expected {
			t.Errorf("Expected branch '%s' at index %d, got '%s'", expected, i, result[i])
		}
	}
}

func TestGetBranches_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	mockRunner.EXPECT().
		GitOutput("branch", "--list").
		Return("", nil).
		Times(1)

	result, err := gitHelper.GetBranches()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 branches, got %d", len(result))
	}
}

func TestGetBranches_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	expectedError := errors.New("git error")

	mockRunner.EXPECT().
		GitOutput("branch", "--list").
		Return("", expectedError).
		Times(1)

	result, err := gitHelper.GetBranches()

	if err == nil {
		t.Error("Expected error, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on error")
	}
}

func TestRebaseBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	branch := "feature1"
	parent := "main"

	// checkout
	mockRunner.EXPECT().
		Git("checkout", branch).
		Return(nil).
		Times(1)

	// set pending parent
	mockRunner.EXPECT().
		Git("config", "--local", "gt.pending.parent", parent).
		Return(nil).
		Times(1)

	// set pending child
	mockRunner.EXPECT().
		Git("config", "--local", "gt.pending.child", branch).
		Return(nil).
		Times(1)

	// rebase
	mockRunner.EXPECT().
		Git("rebase", parent).
		Return(nil).
		Times(1)

	// clear pending parent
	mockRunner.EXPECT().
		Git("config", "--local", "--unset", "gt.pending.parent").
		Return(nil).
		Times(1)

	// clear pending child
	mockRunner.EXPECT().
		Git("config", "--local", "--unset", "gt.pending.child").
		Return(nil).
		Times(1)

	// SetParent
	mockRunner.EXPECT().
		Git("config", "--local", "gt.branch."+branch+".parent", parent).
		Return(nil).
		Times(1)

	err := gitHelper.RebaseBranch(branch, parent)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRebaseBranch_CheckoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	branch := "feature1"
	parent := "main"
	expectedError := errors.New("checkout failed")

	mockRunner.EXPECT().
		Git("checkout", branch).
		Return(expectedError).
		Times(1)

	err := gitHelper.RebaseBranch(branch, parent)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestRebaseBranch_RebaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	gitHelper := &GitHelperImpl{runner: mockRunner}

	branch := "feature1"
	parent := "main"
	expectedError := errors.New("rebase failed")

	// checkout success
	mockRunner.EXPECT().
		Git("checkout", branch).
		Return(nil).
		Times(1)

	// set pending parent
	mockRunner.EXPECT().
		Git("config", "--local", "gt.pending.parent", parent).
		Return(nil).
		Times(1)

	// set pending child
	mockRunner.EXPECT().
		Git("config", "--local", "gt.pending.child", branch).
		Return(nil).
		Times(1)

	// rebase fails
	mockRunner.EXPECT().
		Git("rebase", parent).
		Return(expectedError).
		Times(1)

	err := gitHelper.RebaseBranch(branch, parent)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
