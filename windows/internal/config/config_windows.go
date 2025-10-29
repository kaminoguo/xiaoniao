//go:build windows
// +build windows

package config

import (
	"os"
	"path/filepath"
)

// GetConfigDir returns the config directory path for Windows
func GetConfigDir() string {
	appData := os.Getenv("APPDATA")
	if appData == "" {
		appData = os.Getenv("USERPROFILE")
		if appData != "" {
			appData = filepath.Join(appData, "AppData", "Roaming")
		}
	}
	if appData == "" {
		// Fallback to current directory
		appData = "."
	}
	return filepath.Join(appData, "xiaoniao")
}

// GetConfigPath returns the config file path for Windows
func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.json")
}

// GetPromptsPath returns the prompts file path for Windows
func GetPromptsPath() string {
	return filepath.Join(GetConfigDir(), "prompts.json")
}