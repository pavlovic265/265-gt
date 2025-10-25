package helpers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestGetCurrentBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	expectedBranch := "main"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(expectedBranch + "\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetCurrentBranch()

	// Assertions
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

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	expectedError := errors.New("git error")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rev-parse", "--abbrev-ref", "HEAD"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(bytes.Buffer{}, expectedError).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetCurrentBranch()

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
	if result != "" {
		t.Error("Expected empty string result on error")
	}
}

func TestGetBranches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	// Mock git branch output
	branchOutput := `* main
  feature1
  feature2
  develop`

	expectedBranches := []string{"main", "feature1", "feature2", "develop"}

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "--list"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(branchOutput)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetBranches()

	// Assertions
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

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	// Mock empty git branch output
	branchOutput := ""

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "--list"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(branchOutput)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetBranches()

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 branches, got %d", len(result))
	}
}

func TestGetBranches_WithEmptyLines(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	// Mock git branch output with empty lines
	branchOutput := `* main

  feature1
  
  feature2
`

	expectedBranches := []string{"main", "feature1", "feature2"}

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "--list"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(branchOutput)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetBranches()

	// Assertions
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

func TestGetBranches_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	expectedError := errors.New("git error")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"branch", "--list"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(bytes.Buffer{}, expectedError).
		Times(1)

	// Execute the function
	result, err := gitHelper.GetBranches()

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
	if result != nil {
		t.Error("Expected nil result on error")
	}
}

func TestRebaseBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"
	parent := "main"

	// Set up expectations for checkout
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"checkout", branch}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for setting pending parent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.pending.parent", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for setting pending child
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.pending.child", branch}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for rebase
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rebase", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for clearPendingMove - unsetting pending parent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "--unset", "gt.pending.parent"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for clearPendingMove - unsetting pending child
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "--unset", "gt.pending.child"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for SetParent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.branch." + branch + ".parent", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.RebaseBranch(branch, parent)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRebaseBranch_CheckoutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"
	parent := "main"
	expectedError := errors.New("checkout failed")

	// Set up expectations for checkout error
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"checkout", branch}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.RebaseBranch(branch, parent)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestRebaseBranch_RebaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"
	parent := "main"
	expectedError := errors.New("rebase failed")

	// Set up expectations for checkout (success)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"checkout", branch}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for setting pending parent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.pending.parent", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for setting pending child
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.pending.child", branch}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for rebase (error)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rebase", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.RebaseBranch(branch, parent)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestRebaseBranch_SetParentError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"
	parent := "main"
	expectedError := errors.New("set parent failed")

	// Set up expectations for checkout (success)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"checkout", branch}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for setting pending parent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.pending.parent", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for setting pending child
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.pending.child", branch}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for rebase (success)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"rebase", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for clearPendingMove - unsetting pending parent
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "--unset", "gt.pending.parent"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for clearPendingMove - unsetting pending child
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "--unset", "gt.pending.child"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Set up expectations for SetParent (error)
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--local", "gt.branch." + branch + ".parent", parent}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.RebaseBranch(branch, parent)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
