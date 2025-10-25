package config

import "github.com/pavlovic265/265-gt/utils/pointer"

// GetCurrentVersion returns the current version string
func (d *DefaultConfigManager) GetCurrentVersion() string {
	if globalConfig.Version == nil {
		return ""
	}
	return globalConfig.Version.CurrentVersion
}

// GetVersion returns the version information
func (d *DefaultConfigManager) GetVersion() Version {
	return pointer.Deref(globalConfig.Version)
}
