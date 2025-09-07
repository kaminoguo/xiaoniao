//go:build darwin
// +build darwin

package config

import (
	"os"
	"path/filepath"
)

func GetConfigDir() string {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, "Library", "Application Support", "xiaoniao")
	
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}
	
	return configDir
}

func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.json")
}

func GetPromptsPath() string {
	return filepath.Join(GetConfigDir(), "prompts.json")
}