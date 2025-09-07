//go:build windows
// +build windows

package hotkey

import (
	"fmt"
	"strings"
	"sync"

	"golang.design/x/hotkey"
)

// ParseHotkeyString parses a hotkey string and returns modifiers and key
// Windows version - same as Linux for now since we use the same library
func ParseHotkeyString(hotkeyStr string) ([]hotkey.Modifier, hotkey.Key, error) {
	if hotkeyStr == "" {
		return nil, 0, fmt.Errorf("empty hotkey string")
	}

	parts := strings.Split(hotkeyStr, "+")
	if len(parts) < 2 {
		return nil, 0, fmt.Errorf("invalid hotkey format")
	}

	var mods []hotkey.Modifier
	var key hotkey.Key

	// Parse modifiers
	for i := 0; i < len(parts)-1; i++ {
		switch strings.ToLower(strings.TrimSpace(parts[i])) {
		case "ctrl", "control":
			mods = append(mods, hotkey.ModCtrl)
		case "alt":
			mods = append(mods, hotkey.ModAlt)
		case "shift":
			mods = append(mods, hotkey.ModShift)
		case "win", "super", "cmd":
			mods = append(mods, hotkey.ModWin)
		default:
			return nil, 0, fmt.Errorf("unknown modifier: %s", parts[i])
		}
	}

	// Parse key
	keyStr := strings.ToUpper(strings.TrimSpace(parts[len(parts)-1]))
	if len(keyStr) == 1 && keyStr[0] >= 'A' && keyStr[0] <= 'Z' {
		key = hotkey.Key(keyStr[0])
	} else {
		// Handle special keys
		switch keyStr {
		case "SPACE":
			key = hotkey.KeySpace
		case "ENTER", "RETURN":
			key = hotkey.KeyReturn
		case "TAB":
			key = hotkey.KeyTab
		case "ESC", "ESCAPE":
			key = hotkey.KeyEscape
		case "F1":
			key = hotkey.KeyF1
		case "F2":
			key = hotkey.KeyF2
		case "F3":
			key = hotkey.KeyF3
		case "F4":
			key = hotkey.KeyF4
		case "F5":
			key = hotkey.KeyF5
		case "F6":
			key = hotkey.KeyF6
		case "F7":
			key = hotkey.KeyF7
		case "F8":
			key = hotkey.KeyF8
		case "F9":
			key = hotkey.KeyF9
		case "F10":
			key = hotkey.KeyF10
		case "F11":
			key = hotkey.KeyF11
		case "F12":
			key = hotkey.KeyF12
		default:
			return nil, 0, fmt.Errorf("unknown key: %s", keyStr)
		}
	}

	return mods, key, nil
}

// Hotkey represents a global hotkey
type Hotkey struct {
	ID       string
	Key      string
	Modifier string
	Callback func()
	hk       *hotkey.Hotkey
}

// Manager manages global hotkeys for Windows
type Manager struct {
	mu      sync.RWMutex
	hotkeys map[string]*Hotkey
	enabled bool
}

// NewManager creates a new hotkey manager for Windows
func NewManager() *Manager {
	return &Manager{
		hotkeys: make(map[string]*Hotkey),
		enabled: false,
	}
}

// Register registers a new hotkey
func (m *Manager) Register(id, hotkeyStr string, callback func()) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Parse hotkey string
	mods, key, err := ParseHotkeyString(hotkeyStr)
	if err != nil {
		return err
	}

	// Create hotkey
	hk := hotkey.New(mods, key)
	
	newHotkey := &Hotkey{
		ID:       id,
		Key:      hotkeyStr,
		Modifier: "",
		Callback: callback,
		hk:       hk,
	}

	// Register callback
	go func() {
		for range hk.Keydown() {
			if callback != nil {
				callback()
			}
		}
	}()

	// Register hotkey
	if err := hk.Register(); err != nil {
		return fmt.Errorf("failed to register hotkey: %v", err)
	}

	m.hotkeys[id] = newHotkey
	return nil
}

// Unregister unregisters a hotkey
func (m *Manager) Unregister(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	hk, exists := m.hotkeys[id]
	if !exists {
		return fmt.Errorf("hotkey %s not found", id)
	}

	if err := hk.hk.Unregister(); err != nil {
		return err
	}

	delete(m.hotkeys, id)
	return nil
}

// Enable enables all hotkeys
func (m *Manager) Enable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = true
}

// Disable disables all hotkeys
func (m *Manager) Disable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.enabled = false
}

// IsEnabled returns whether hotkeys are enabled
func (m *Manager) IsEnabled() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.enabled
}

// GetHotkey returns a hotkey by ID
func (m *Manager) GetHotkey(id string) (*Hotkey, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	hk, exists := m.hotkeys[id]
	return hk, exists
}

// RegisterFromString registers a hotkey from a string representation
func (m *Manager) RegisterFromString(id, hotkeyStr string, callback func()) error {
	return m.Register(id, hotkeyStr, callback)
}