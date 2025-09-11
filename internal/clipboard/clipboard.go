//go:build !windows
// +build !windows

// Package clipboard provides clipboard monitoring functionality
// This is a stub implementation for non-Windows platforms
package clipboard

import (
	"runtime"
	"time"
)

// Monitor represents a clipboard monitor
type Monitor struct {
	running        bool
	onChange       func(string)
	lastText       string
	lastTranslation string
	stopChan       chan bool
}

// NewMonitor creates a new clipboard monitor
func NewMonitor() *Monitor {
	return &Monitor{
		stopChan: make(chan bool, 1),
	}
}

// SetOnChange sets the callback function for clipboard changes
func (m *Monitor) SetOnChange(callback func(string)) {
	m.onChange = callback
}

// Start starts monitoring the clipboard
func (m *Monitor) Start() {
	if m.running {
		return
	}
	m.running = true
	
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-m.stopChan:
				return
			case <-ticker.C:
				// On non-Windows platforms, just continue without actual monitoring
				// This is a stub implementation
				continue
			}
		}
	}()
}

// Stop stops monitoring the clipboard
func (m *Monitor) Stop() {
	if !m.running {
		return
	}
	m.running = false
	select {
	case m.stopChan <- true:
	default:
	}
}

// SetLastTranslation sets the last translation to avoid retranslating
func (m *Monitor) SetLastTranslation(text string) {
	m.lastTranslation = text
}

// GetClipboard returns the current clipboard content (stub)
func GetClipboard() (string, error) {
	return "", nil
}

// SetClipboard sets the clipboard content (stub)
func SetClipboard(text string) error {
	// Print to console since we can't actually set clipboard on non-Windows
	if runtime.GOOS != "windows" {
		// This is a stub - actual implementation would depend on the platform
		return nil
	}
	return nil
}