package helpers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestSetChildren(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	parent := "main"
	children := "feature1 feature2"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.main.children", "feature1 feature2"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.SetChildren(parent, children)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSetChildren_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	parent := "main"
	children := "feature1 feature2"
	expectedError := errors.New("git config error")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.main.children", "feature1 feature2"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.SetChildren(parent, children)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
}

func TestGetChildren(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "main"
	expectedChildren := "feature1 feature2"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(expectedChildren + "\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result := gitHelper.GetChildren(branch)

	// Assertions
	if result != expectedChildren {
		t.Errorf("Expected '%s', got '%s'", expectedChildren, result)
	}
}

func TestGetChildren_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "main"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output with empty result
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString("\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result := gitHelper.GetChildren(branch)

	// Assertions
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestDeleteChildren(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "main"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--unset", "branch.main.children"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.DeleteChildren(branch)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteChildren_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "main"
	expectedError := errors.New("git config error")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--unset", "branch.main.children"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.DeleteChildren(branch)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
}
