//go:build linux
// +build linux

package tray

import (
	"fmt"
)

// Status represents the tray icon status
type Status string

const (
	StatusIdle        Status = "idle"        // Blue bird (normal)
	StatusTranslating Status = "translating" // Green bird (processing)
	StatusError       Status = "error"       // Red bird (error)
)

// Manager manages the system tray (stub implementation for Linux)
type Manager struct {
	status            Status
	visible           bool
	isMonitoring      bool
	isReady           bool
	translationCount  int
	currentPromptName string
	onQuit            func()
	onShow            func()
	onSettings        func()
	onToggleMonitor   func(bool)
	onRefresh         func()
	onSwitchPrompt    func()
	onSelectPrompt    func(string)
	onToggleTerminal  func()
}

// NewManager creates a new tray manager (stub for Linux)
func NewManager() (*Manager, error) {
	fmt.Println("System tray not available on Linux (stub implementation)")
	return &Manager{
		status:  StatusIdle,
		visible: true,
		isReady: true,
	}, nil
}

// SetStatus sets the tray icon status (stub)
func (m *Manager) SetStatus(status Status) {
	m.status = status
}

// SetCurrentPrompt sets the current prompt name (stub)
func (m *Manager) SetCurrentPrompt(name string) {
	m.currentPromptName = name
}

// IncrementTranslationCount increments the translation counter (stub)
func (m *Manager) IncrementTranslationCount() {
	m.translationCount++
}

// UpdateMonitorStatus updates monitoring status (stub)
func (m *Manager) UpdateMonitorStatus(isMonitoring bool) {
	m.isMonitoring = isMonitoring
}

// SetOnQuit sets the quit callback (stub)
func (m *Manager) SetOnQuit(callback func()) {
	m.onQuit = callback
}

// SetOnSettings sets the settings callback (stub)
func (m *Manager) SetOnSettings(callback func()) {
	m.onSettings = callback
}

// SetOnToggleMonitor sets the toggle monitor callback (stub)
func (m *Manager) SetOnToggleMonitor(callback func(bool)) {
	m.onToggleMonitor = callback
}

// SetOnRefresh sets the refresh callback (stub)
func (m *Manager) SetOnRefresh(callback func()) {
	m.onRefresh = callback
}

// SetOnSelectPrompt sets the select prompt callback (stub)
func (m *Manager) SetOnSelectPrompt(callback func(string)) {
	m.onSelectPrompt = callback
}

// SetOnToggleTerminal sets the toggle terminal callback (stub)
func (m *Manager) SetOnToggleTerminal(callback func()) {
	m.onToggleTerminal = callback
}

// UpdatePromptList updates the prompt list in tray menu (stub)
func (m *Manager) UpdatePromptList(prompts []struct{ ID, Name string }) {
	// Stub implementation - just print the prompt list
	fmt.Printf("Available prompts: %d\n", len(prompts))
}

// Quit quits the tray (stub)
func (m *Manager) Quit() {
	if m.onQuit != nil {
		m.onQuit()
	}
}

// IsReady returns whether the tray is ready (stub)
func (m *Manager) IsReady() bool {
	return m.isReady
}