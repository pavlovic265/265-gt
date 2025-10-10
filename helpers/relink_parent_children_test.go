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
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test data
	parent := "main"
	parentChildren := "feature1 feature2"
	branch := "feature1"
	branchChildren := "feature1.1 feature1.2"

	// Set up expectations
	mockGitHelper.EXPECT().
		SetParent(mockExecutor, parent, "feature1.1").
		Return(nil).
		Times(1)

	mockGitHelper.EXPECT().
		SetParent(mockExecutor, parent, "feature1.2").
		Return(nil).
		Times(1)

	mockGitHelper.EXPECT().
		SetChildren(mockExecutor, parent, "feature2 feature1.1 feature1.2").
		Return(nil).
		Times(1)

	// Execute the function
	err := RelinkParentChildren(
		mockGitHelper,
		mockExecutor,
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
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test data with empty parent
	parent := ""
	parentChildren := ""
	branch := "feature1"
	branchChildren := "feature1.1"

	// No expectations needed since function should return early

	// Execute the function
	err := RelinkParentChildren(
		mockGitHelper,
		mockExecutor,
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
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	mockExecutor := mocks.NewMockExecutor(ctrl)

	// Test data
	parent := "main"
	parentChildren := "feature1"
	branch := "feature1"
	branchChildren := "feature1.1"

	// Set up expectation for error
	expectedError := errors.New("git config error")
	mockGitHelper.EXPECT().
		SetParent(mockExecutor, parent, "feature1.1").
		Return(expectedError).
		Times(1)

	// Execute the function
	err := RelinkParentChildren(
		mockGitHelper,
		mockExecutor,
		parent,
		parentChildren,
		branch,
		branchChildren,
	)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
