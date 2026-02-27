package pr_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pavlovic265/265-gt/client"
	"github.com/pavlovic265/265-gt/commands/pr"
	"github.com/pavlovic265/265-gt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListCommand_Command(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	listCmd := pr.NewListCommand(mockRunner, mockConfigManager, mockGitHelper, cliClient)
	cmd := listCmd.Command()

	assert.Equal(t, "list", cmd.Use)
	assert.Equal(t, []string{"li"}, cmd.Aliases)
	assert.Equal(t, "show list of pull requests", cmd.Short)
}

func TestNewListCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRunner := mocks.NewMockRunner(ctrl)
	mockConfigManager := mocks.NewMockConfigManager(ctrl)
	mockGitHelper := mocks.NewMockGitHelper(ctrl)
	cliClient := client.NewGitHubClient(mockGitHelper)

	listCmd := pr.NewListCommand(mockRunner, mockConfigManager, mockGitHelper, cliClient)
	cmd := listCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "list", cmd.Use)
}

func TestFormatPullRequest_ReviewApproved(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	listCmd := pr.NewListCommand(
		mocks.NewMockRunner(ctrl),
		mocks.NewMockConfigManager(ctrl),
		mocks.NewMockGitHelper(ctrl),
		client.NewGitHubClient(mocks.NewMockGitHelper(ctrl)),
	)

	result := listCmd.FormatPullRequest(client.PullRequest{
		Number:      42,
		Title:       "My PR",
		URL:         "https://github.com/test/repo/pull/42",
		Branch:      "feature",
		Mergeable:   "MERGEABLE",
		StatusState: client.StatusStateTypeSuccess,
		ReviewState: client.ReviewStateApproved,
	})

	assert.Contains(t, result.Title, "42")
	assert.Contains(t, result.Title, "My PR")
	assert.Contains(t, result.Title, "✓")
	assert.Contains(t, result.Title, "●")
	assert.NotContains(t, result.Title, "⚠")
}

func TestFormatPullRequest_ReviewChangesRequested(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	listCmd := pr.NewListCommand(
		mocks.NewMockRunner(ctrl),
		mocks.NewMockConfigManager(ctrl),
		mocks.NewMockGitHelper(ctrl),
		client.NewGitHubClient(mocks.NewMockGitHelper(ctrl)),
	)

	result := listCmd.FormatPullRequest(client.PullRequest{
		Number:      10,
		Title:       "Needs work",
		URL:         "https://github.com/test/repo/pull/10",
		Branch:      "fix",
		Mergeable:   "MERGEABLE",
		StatusState: client.StatusStateTypeFailure,
		ReviewState: client.ReviewStateChangesRequested,
	})

	assert.Contains(t, result.Title, "10")
	assert.Contains(t, result.Title, "Needs work")
	assert.Contains(t, result.Title, "✗")
	assert.Contains(t, result.Title, "●")
	assert.NotContains(t, result.Title, "⚠")
}

func TestFormatPullRequest_NoReviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	listCmd := pr.NewListCommand(
		mocks.NewMockRunner(ctrl),
		mocks.NewMockConfigManager(ctrl),
		mocks.NewMockGitHelper(ctrl),
		client.NewGitHubClient(mocks.NewMockGitHelper(ctrl)),
	)

	result := listCmd.FormatPullRequest(client.PullRequest{
		Number:      7,
		Title:       "New feature",
		URL:         "https://github.com/test/repo/pull/7",
		Branch:      "feat",
		Mergeable:   "UNKNOWN",
		StatusState: client.StatusStateTypePending,
		ReviewState: "",
	})

	assert.Contains(t, result.Title, "7")
	assert.Contains(t, result.Title, "New feature")
	assert.Contains(t, result.Title, "*")
	assert.Contains(t, result.Title, "●")
	assert.NotContains(t, result.Title, "⚠")
}

func TestFormatPullRequest_ConflictIndicator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	listCmd := pr.NewListCommand(
		mocks.NewMockRunner(ctrl),
		mocks.NewMockConfigManager(ctrl),
		mocks.NewMockGitHelper(ctrl),
		client.NewGitHubClient(mocks.NewMockGitHelper(ctrl)),
	)

	result := listCmd.FormatPullRequest(client.PullRequest{
		Number:      99,
		Title:       "Conflicting PR",
		URL:         "https://github.com/test/repo/pull/99",
		Branch:      "conflict-branch",
		Mergeable:   "CONFLICTING",
		StatusState: client.StatusStateTypeSuccess,
		ReviewState: client.ReviewStateApproved,
	})

	assert.Contains(t, result.Title, "99")
	assert.Contains(t, result.Title, "Conflicting PR")
	assert.Contains(t, result.Title, "●")
	assert.Contains(t, result.Title, "⚠")
}

func TestFormatPullRequest_NoConflictWhenMergeable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	listCmd := pr.NewListCommand(
		mocks.NewMockRunner(ctrl),
		mocks.NewMockConfigManager(ctrl),
		mocks.NewMockGitHelper(ctrl),
		client.NewGitHubClient(mocks.NewMockGitHelper(ctrl)),
	)

	result := listCmd.FormatPullRequest(client.PullRequest{
		Number:      5,
		Title:       "Clean PR",
		URL:         "https://github.com/test/repo/pull/5",
		Branch:      "clean",
		Mergeable:   "MERGEABLE",
		StatusState: client.StatusStateTypeSuccess,
		ReviewState: client.ReviewStateApproved,
	})

	assert.NotContains(t, result.Title, "⚠")
}

func TestFormatPullRequest_MergeQueuedIndicator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	listCmd := pr.NewListCommand(
		mocks.NewMockRunner(ctrl),
		mocks.NewMockConfigManager(ctrl),
		mocks.NewMockGitHelper(ctrl),
		client.NewGitHubClient(mocks.NewMockGitHelper(ctrl)),
	)

	result := listCmd.FormatPullRequest(client.PullRequest{
		Number:      12,
		Title:       "Queued PR",
		URL:         "https://github.com/test/repo/pull/12",
		Branch:      "queued",
		Mergeable:   "MERGEABLE",
		StatusState: client.StatusStateTypeSuccess,
		ReviewState: client.ReviewStateApproved,
		MergeQueued: true,
	})

	assert.Contains(t, result.Title, "⧗")
}
