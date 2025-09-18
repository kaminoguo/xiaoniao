//go:build darwin
// +build darwin

// Package clipboard provides clipboard monitoring functionality for macOS
package clipboard

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Monitor represents a clipboard monitor for macOS
type Monitor struct {
	running         bool
	onChange        func(string)
	lastText        string
	lastTranslation string
	stopChan        chan bool
	lastChangeCount int
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
				text, err := GetClipboard()
				if err != nil {
					continue
				}

				// Check if clipboard content has changed
				if text != "" && text != m.lastText && text != m.lastTranslation {
					m.lastText = text
					if m.onChange != nil {
						m.onChange(text)
					}
				}
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

// GetClipboard returns the current clipboard content using pbpaste
func GetClipboard() (string, error) {
	cmd := exec.Command("pbpaste")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get clipboard: %v", err)
	}
	return string(output), nil
}

// SetClipboard sets the clipboard content using pbcopy
func SetClipboard(text string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set clipboard: %v", err)
	}
	return nil
}