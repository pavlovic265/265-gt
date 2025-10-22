package helpers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/mocks"
)

func TestSetParent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	parent := "main"
	child := "feature1"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.feature1.parent", "main"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.SetParent(parent, child)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSetParent_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	parent := "main"
	child := "feature1"
	expectedError := errors.New("git config error")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.feature1.parent", "main"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.SetParent(parent, child)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
}

func TestGetParent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"
	expectedParent := "main"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.feature1.parent"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString(expectedParent + "\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Execute the function
	result := gitHelper.GetParent(branch)

	// Assertions
	if result != expectedParent {
		t.Errorf("Expected '%s', got '%s'", expectedParent, result)
	}
}

func TestGetParent_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.feature1.parent"}).
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
	result := gitHelper.GetParent(branch)

	// Assertions
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestDeleteParent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--unset", "branch.feature1.parent"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.DeleteParent(branch)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteParent_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{exe: mockExecutor}

	branch := "feature1"
	expectedError := errors.New("git config error")

	// Set up expectations
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--unset", "branch.feature1.parent"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(expectedError).
		Times(1)

	// Execute the function
	err := gitHelper.DeleteParent(branch)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
}
