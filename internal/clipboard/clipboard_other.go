//go:build !windows
// +build !windows

package clipboard

// Monitor provides a stub implementation for non-Windows platforms
type Monitor struct{}

// NewMonitor creates a new clipboard monitor stub
func NewMonitor() *Monitor {
	return &Monitor{}
}

// SetOnChange sets the callback function for clipboard changes (stub)
func (m *Monitor) SetOnChange(callback func(string)) {
	// Stub implementation - does nothing
}

// Start starts monitoring clipboard changes (stub)
func (m *Monitor) Start() {
	// Stub implementation - does nothing
}

// Stop stops monitoring clipboard changes (stub)
func (m *Monitor) Stop() {
	// Stub implementation - does nothing
}

// SetLastTranslation sets the last translation to avoid re-processing (stub)
func (m *Monitor) SetLastTranslation(text string) {
	// Stub implementation - does nothing
}

// SetClipboard sets the clipboard content (stub)
func SetClipboard(text string) error {
	// Stub implementation - does nothing
	return nil
}

// GetClipboard gets the clipboard content (stub)
func GetClipboard() (string, error) {
	// Stub implementation - returns empty string
	return "", nil
}