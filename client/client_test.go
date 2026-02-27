package client

import (
	"testing"

	"github.com/pavlovic265/265-gt/constants"
)

func TestParseRemoteURL_SSH(t *testing.T) {
	repo, err := ParseRemoteURL("git@github.com:owner/repo.git")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if repo.Platform != "github.com" || repo.Owner != "owner" || repo.Repo != "repo" {
		t.Fatalf("unexpected parsed repo: %+v", repo)
	}
}

func TestParseRemoteURL_HTTPS(t *testing.T) {
	repo, err := ParseRemoteURL("https://gitlab.com/owner/repo")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if repo.Platform != "gitlab.com" || repo.Owner != "owner" || repo.Repo != "repo" {
		t.Fatalf("unexpected parsed repo: %+v", repo)
	}
}

func TestParseRemoteURL_Invalid(t *testing.T) {
	_, err := ParseRemoteURL("not-a-url")
	if err == nil {
		t.Fatal("expected parse error")
	}
}

func TestNewRestCliClient(t *testing.T) {
	gh, err := NewRestCliClient(constants.GitHubPlatform, nil)
	if err != nil || gh == nil {
		t.Fatalf("expected github client, err=%v", err)
	}

	gl, err := NewRestCliClient(constants.GitLabPlatform, nil)
	if err != nil || gl == nil {
		t.Fatalf("expected gitlab client, err=%v", err)
	}

	_, err = NewRestCliClient(constants.Platform("Unknown"), nil)
	if err == nil {
		t.Fatal("expected unsupported platform error")
	}
}
