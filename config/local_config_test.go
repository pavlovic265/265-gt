package config

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/pavlovic265/265-gt/constants"
)

func TestGetLocalConfigPath_NotRepoReturnsEmpty(t *testing.T) {
	r := stubRunner{gitOutputErr: errors.New("boom")}

	mgr := NewDefaultConfigManager(r)
	path, err := mgr.getLocalConfigPath()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if path != "" {
		t.Fatalf("expected empty path, got %q", path)
	}
}

func TestGetLocalConfigPath_Repo(t *testing.T) {
	r := stubRunner{gitOutputValue: "/tmp/repo"}

	mgr := NewDefaultConfigManager(r)
	path, err := mgr.getLocalConfigPath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := filepath.Join("/tmp/repo", constants.LocalConfigFileName)
	if path != expected {
		t.Fatalf("expected %q, got %q", expected, path)
	}
}

type stubRunner struct {
	gitOutputValue string
	gitOutputErr   error
}

func (s stubRunner) Git(args ...string) error               { return nil }
func (s stubRunner) Exec(name string, args ...string) error { return nil }
func (s stubRunner) ExecOutput(name string, args ...string) (string, error) {
	return "", nil
}
func (s stubRunner) GitOutput(args ...string) (string, error) {
	if len(args) == 2 && args[0] == "rev-parse" && args[1] == "--show-toplevel" {
		return s.gitOutputValue, s.gitOutputErr
	}
	return "", nil
}
