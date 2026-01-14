package config

import (
	"context"
	"fmt"
)

type configContextKey struct{}

type ConfigContext struct {
	Global     *GlobalConfigStruct
	Local      *LocalConfigStruct
	dirty      bool
	localDirty bool
}

func NewConfigContext(global *GlobalConfigStruct, local *LocalConfigStruct) *ConfigContext {
	return &ConfigContext{Global: global, Local: local, dirty: false, localDirty: false}
}

func (c *ConfigContext) MarkDirty() {
	c.dirty = true
}

func (c *ConfigContext) MarkLocalDirty() {
	c.localDirty = true
}

func (c *ConfigContext) IsDirty() bool {
	return c.dirty
}

func (c *ConfigContext) IsLocalDirty() bool {
	return c.localDirty
}

func WithConfig(ctx context.Context, cfg *ConfigContext) context.Context {
	return context.WithValue(ctx, configContextKey{}, cfg)
}

func GetConfig(ctx context.Context) (*ConfigContext, bool) {
	if ctx == nil {
		return nil, false
	}
	cfg, ok := ctx.Value(configContextKey{}).(*ConfigContext)
	return cfg, ok
}

func RequireGlobal(ctx context.Context) (*ConfigContext, error) {
	cfg, ok := GetConfig(ctx)
	if !ok || cfg.Global == nil {
		return nil, fmt.Errorf("no config found - run 'gt config global' to set up")
	}
	return cfg, nil
}
