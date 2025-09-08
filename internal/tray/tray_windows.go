//go:build windows
// +build windows

package tray

import (
	"os"
	"path/filepath"
	"github.com/getlantern/systray"
)

// getConfigPath returns the config file path for Windows
func getConfigPath() string {
	appData := os.Getenv("APPDATA")
	if appData == "" {
		appData = os.Getenv("USERPROFILE")
	}
	configDir := filepath.Join(appData, "xiaoniao")
	os.MkdirAll(configDir, 0755)
	return filepath.Join(configDir, "config.json")
}

// openConfig opens the config UI on Windows
func openConfig() error {
	// On Windows, we use the same TUI interface
	// The user needs to run this in Windows Terminal or similar
	return nil
}

// InitializeForWindows initializes the system tray for Windows
// This must be called from the main thread
func (m *Manager) InitializeForWindows() {
	// Windows: Run directly without goroutine
	systray.Run(m.onReady, m.onExit)
}