//go:build !windows
// +build !windows

// Package hotkey provides hotkey registration functionality
// This is a stub implementation for non-Windows platforms
package hotkey

import (
	"errors"
	"fmt"
)

// Manager represents a hotkey manager
type Manager struct {
	registeredHotkeys map[string]func()
}

// NewManager creates a new hotkey manager
func NewManager() *Manager {
	return &Manager{
		registeredHotkeys: make(map[string]func()),
	}
}

// RegisterFromString registers a hotkey from a string representation
func (m *Manager) RegisterFromString(name, hotkeyStr string, callback func()) error {
	if hotkeyStr == "" {
		return errors.New("hotkey string cannot be empty")
	}
	
	// For non-Windows platforms, we just store the callback but don't actually register
	// This is a stub implementation
	m.registeredHotkeys[name] = callback
	fmt.Printf("Hotkey registration (stub): %s -> %s\n", name, hotkeyStr)
	return nil
}

// Unregister removes a registered hotkey
func (m *Manager) Unregister(name string) error {
	delete(m.registeredHotkeys, name)
	return nil
}

// Stop stops the hotkey manager
func (m *Manager) Stop() {
	m.registeredHotkeys = make(map[string]func())
}