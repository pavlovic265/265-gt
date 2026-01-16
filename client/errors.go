package client

import "errors"

// Common client errors.
var (
	ErrConfigNotLoaded = errors.New("config not loaded")
	ErrNoActiveAccount = errors.New("no active account")
)
