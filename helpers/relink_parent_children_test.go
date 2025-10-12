package helpers

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestRelinkParentChildren(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	// Test data
	parent := "main"
	parentChildren := "feature1 feature2"
	branch := "feature1"
	branchChildren := "feature1.1 feature1.2"

	// Set up expectations for SetParent calls
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(2)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.feature1.1.parent", "main"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.feature1.2.parent", "main"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(2)

	// Set up expectations for SetChildren call
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.main.children", "feature2 feature1.1 feature1.2"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.RelinkParentChildren(
		parent,
		parentChildren,
		branch,
		branchChildren,
	)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRelinkParentChildren_EmptyParent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	gitHelper := &GitHelperImpl{}

	// Test data with empty parent
	parent := ""
	parentChildren := ""
	branch := "feature1"
	branchChildren := "feature1.1"

	// No expectations needed since function should return early

	// Execute the function
	err := gitHelper.RelinkParentChildren(
		parent,
		parentChildren,
		branch,
		branchChildren,
	)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRelinkParentChildren_SetParentError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	// Test data
	parent := "main"
	parentChildren := "feature1"
	branch := "feature1"
	branchChildren := "feature1.1"

	// Set up expectation for error
	expectedError := errors.New("git config error")
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.feature1.1.parent", "main"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.RelinkParentChildren(
		parent,
		parentChildren,
		branch,
		branchChildren,
	)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
}
