package config

import (
	"os"
	"path/filepath"
)

// GetConfigDir returns the configuration directory for macOS
// On macOS, we use ~/Library/Application Support/xiaoniao/
func GetConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		// Fallback to current directory
		return "."
	}

	configDir := filepath.Join(home, "Library", "Application Support", "xiaoniao")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		// Fallback to home directory
		return home
	}

	return configDir
}

// GetLogPath returns the log file path for macOS
func GetLogPath() string {
	return filepath.Join(GetConfigDir(), "xiaoniao.log")
}