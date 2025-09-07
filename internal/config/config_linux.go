//go:build linux
// +build linux

package config

import (
	"os"
	"path/filepath"
)

// GetConfigDir returns the config directory path for Linux
func GetConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".", ".config", "xiaoniao")
	}
	return filepath.Join(home, ".config", "xiaoniao")
}

// GetConfigPath returns the config file path for Linux
func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.json")
}

// GetPromptsPath returns the prompts file path for Linux
func GetPromptsPath() string {
	return filepath.Join(GetConfigDir(), "prompts.json")
}