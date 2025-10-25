package config

import "github.com/pavlovic265/265-gt/utils/pointer"

// SaveActiveAccount saves the active account to the global config
func (d *DefaultConfigManager) SaveActiveAccount(account Account) error {
	globalConfig.ActiveAccount = pointer.From(account)

	return d.SaveGlobalConfig(globalConfig)
}

// SetActiveAccount sets the active account in the global config
func (d *DefaultConfigManager) SetActiveAccount(account Account) error {
	return d.SaveActiveAccount(account)
}

// GetActiveAccount returns the currently active account
func (d *DefaultConfigManager) GetActiveAccount() Account {
	return pointer.Deref(globalConfig.ActiveAccount)
}

// ClearActiveAccount clears the active account from the global config
func (d *DefaultConfigManager) ClearActiveAccount() error {
	globalConfig.ActiveAccount = nil

	return d.SaveGlobalConfig(globalConfig)
}

// HasActiveAccount checks if there is an active account configured
func (d *DefaultConfigManager) HasActiveAccount() bool {
	return globalConfig.ActiveAccount != nil && globalConfig.ActiveAccount.User != ""
}

// GetAccounts returns all configured accounts
func (d *DefaultConfigManager) GetAccounts() []Account {
	accounts := make([]Account, len(globalConfig.Accounts))
	copy(accounts, globalConfig.Accounts)
	return accounts
}
