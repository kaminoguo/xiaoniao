//go:build !windows
// +build !windows

package hotkey

// Manager provides a stub implementation for non-Windows platforms
type Manager struct{}

// NewManager creates a new hotkey manager stub
func NewManager() *Manager {
	return &Manager{}
}

// RegisterFromString registers a hotkey from string (stub)
func (m *Manager) RegisterFromString(id, hotkeyStr string, callback func()) error {
	// Stub implementation - does nothing
	return nil
}

// Unregister unregisters a hotkey (stub)
func (m *Manager) Unregister(id string) error {
	// Stub implementation - does nothing
	return nil
}

// Stop stops the hotkey manager (stub)
func (m *Manager) Stop() {
	// Stub implementation - does nothing
}