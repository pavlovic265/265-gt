package config

// GetProtectedBranches returns the list of protected branches from local config
func (d *DefaultConfigManager) GetProtectedBranches() []string {
	return localConfig.Protected
}

// SaveProtectedBranches saves protected branches to the local config
func (d *DefaultConfigManager) SaveProtectedBranches(branches []string) error {
	localConfig.Protected = append(localConfig.Protected, branches...)

	return d.SaveLocalConfig(localConfig)
}
