package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func readConfig[T any](filename string) (*T, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var cfg T
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &cfg, nil
}

func writeConfig[T any](filename string, cfg *T) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Secure file writing — ensure permissions are restrictive (rw for owner only)
	return os.WriteFile(filename, data, 0600)
}
