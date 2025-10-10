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
	gitHelper := &GitHelperImpl{}

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
	err := gitHelper.SetParent(mockExecutor, parent, child)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSetParent_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

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
	err := gitHelper.SetParent(mockExecutor, parent, child)

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
	gitHelper := &GitHelperImpl{}

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
	result := gitHelper.GetParent(mockExecutor, branch)

	// Assertions
	if result != expectedParent {
		t.Errorf("Expected '%s', got '%s'", expectedParent, result)
	}
}

func TestGetParent_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

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
	result := gitHelper.GetParent(mockExecutor, branch)

	// Assertions
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestDeleteParent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

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
	err := gitHelper.DeleteParent(mockExecutor, branch)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteParent_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

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
	err := gitHelper.DeleteParent(mockExecutor, branch)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
}

func TestDeleteFromParentChildren(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	parent := "main"
	branch := "feature1"
	children := "feature1 feature2 feature3"

	// Set up expectations for GetChildren
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
	mockOutput.WriteString(children + "\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Set up expectations for SetChildren
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.main.children", "feature2 feature3"}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.DeleteFromParentChildren(mockExecutor, parent, branch)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteFromParentChildren_EmptyChildren(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	parent := "main"
	branch := "feature1"

	// Set up expectations for GetChildren
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "--get", "branch.main.children"}).
		Return(mockExecutor).
		Times(1)

	// Create a mock output with empty result (no content, just newline)
	mockOutput := bytes.Buffer{}
	mockOutput.WriteString("\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Since GetChildren returns empty string, strings.Split("", " ") returns [""]
	// which has length 1, so the function continues and processes the empty string
	// Since "" != "feature1", it adds "" to newChildren, so newChildren = [""]
	// Since len(newChildren) > 0, it calls SetChildren with empty string
	// Set up expectations for SetChildren
	mockExecutor.EXPECT().
		WithGit().
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		WithArgs([]string{"config", "branch.main.children", ""}).
		Return(mockExecutor).
		Times(1)

	mockExecutor.EXPECT().
		Run().
		Return(nil).
		Times(1)

	// Execute the function
	err := gitHelper.DeleteFromParentChildren(mockExecutor, parent, branch)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteFromParentChildren_LastChild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExecutor := mocks.NewMockExecutor(ctrl)
	gitHelper := &GitHelperImpl{}

	parent := "main"
	branch := "feature1"
	children := "feature1"

	// Set up expectations for GetChildren
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
	mockOutput.WriteString(children + "\n")

	mockExecutor.EXPECT().
		RunWithOutput().
		Return(mockOutput, nil).
		Times(1)

	// Set up expectations for DeleteChildren
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
	err := gitHelper.DeleteFromParentChildren(mockExecutor, parent, branch)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
