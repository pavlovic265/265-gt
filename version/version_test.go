package version

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pavlovic265/265-gt/config"
)

func TestShouldCheckVersion(t *testing.T) {
	if !shouldCheckVersion(config.Version{}) {
		t.Fatal("expected true when never checked")
	}

	recent := config.Version{LastChecked: time.Now().Add(-1 * time.Hour).Format(time.RFC3339)}
	if shouldCheckVersion(recent) {
		t.Fatal("expected false for recent check")
	}

	old := config.Version{LastChecked: time.Now().Add(-25 * time.Hour).Format(time.RFC3339)}
	if !shouldCheckVersion(old) {
		t.Fatal("expected true for old check")
	}

	invalid := config.Version{LastChecked: "not-a-date"}
	if !shouldCheckVersion(invalid) {
		t.Fatal("expected true for invalid timestamp")
	}
}

func TestShowVersionNotification(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	showVersionNotification("v1.0.0", "v1.1.0", "https://example.com")

	_ = w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	out := buf.String()

	if !strings.Contains(out, "1.0.0") || !strings.Contains(out, "1.1.0") {
		t.Fatalf("expected versions in output, got: %s", out)
	}
	if !strings.Contains(out, "gt upgrade") {
		t.Fatalf("expected upgrade instruction, got: %s", out)
	}
}
