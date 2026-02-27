package config

import (
	"context"
	"testing"
)

func TestConfigContextDirtyFlags(t *testing.T) {
	cfg := NewConfigContext(&GlobalConfigStruct{}, &LocalConfigStruct{})
	if cfg.IsDirty() || cfg.IsLocalDirty() {
		t.Fatal("new context should start clean")
	}

	cfg.MarkDirty()
	cfg.MarkLocalDirty()
	if !cfg.IsDirty() || !cfg.IsLocalDirty() {
		t.Fatal("expected dirty flags to be set")
	}
}

func TestWithAndGetConfig(t *testing.T) {
	base := context.Background()
	cfg := NewConfigContext(&GlobalConfigStruct{}, nil)
	ctx := WithConfig(base, cfg)

	got, ok := GetConfig(ctx)
	if !ok || got != cfg {
		t.Fatal("expected config in context")
	}

	if _, ok := GetConfig(context.TODO()); ok {
		t.Fatal("expected context without config to return not found")
	}
}

func TestRequireGlobal(t *testing.T) {
	ctx := context.Background()
	if _, err := RequireGlobal(ctx); err == nil {
		t.Fatal("expected error without config")
	}

	cfg := NewConfigContext(&GlobalConfigStruct{}, nil)
	ctx = WithConfig(ctx, cfg)
	if _, err := RequireGlobal(ctx); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
