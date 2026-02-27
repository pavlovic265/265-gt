package config

import (
	"path/filepath"
	"testing"
)

func TestWriteAndReadConfig(t *testing.T) {
	type sample struct {
		Name string `yaml:"name"`
	}

	file := filepath.Join(t.TempDir(), "cfg.yml")
	in := &sample{Name: "gt"}

	if err := writeConfig(file, in); err != nil {
		t.Fatalf("writeConfig failed: %v", err)
	}

	out, err := readConfig[sample](file)
	if err != nil {
		t.Fatalf("readConfig failed: %v", err)
	}

	if out.Name != in.Name {
		t.Fatalf("expected %q, got %q", in.Name, out.Name)
	}
}
