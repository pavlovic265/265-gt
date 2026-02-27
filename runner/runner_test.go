package runner

import (
	"strings"
	"testing"
)

func TestExecOutput_Success(t *testing.T) {
	r := NewRunner()
	out, err := r.ExecOutput("sh", "-c", "echo hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "hello" {
		t.Fatalf("expected hello, got %q", out)
	}
}

func TestExecOutput_ErrorIncludesStderr(t *testing.T) {
	r := NewRunner()
	_, err := r.ExecOutput("sh", "-c", "echo failure >&2; exit 1")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "failure") {
		t.Fatalf("expected stderr in error, got %v", err)
	}
}
