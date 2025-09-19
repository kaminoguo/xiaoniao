package config

import (
	"os"
	"path/filepath"
)

// GetConfigDir returns the configuration directory path for Android
func GetConfigDir() string {
	// Try to get from environment first (for development)
	if home := os.Getenv("HOME"); home != "" {
		return filepath.Join(home, ".xiaoniao")
	}

	// Android internal storage path
	// This will be /data/data/com.liliguo.xiaoniao/files
	if dataDir := os.Getenv("ANDROID_DATA"); dataDir != "" {
		return dataDir
	}

	// Fallback to current directory
	return "./data"
}

// GetLogDir returns the log directory path for Android
func GetLogDir() string {
	return filepath.Join(GetConfigDir(), "logs")
}

// GetConfigPath returns the main config file path
func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.json")
}

// GetPromptsPath returns the user prompts file path
func GetPromptsPath() string {
	return filepath.Join(GetConfigDir(), "prompts.json")
}