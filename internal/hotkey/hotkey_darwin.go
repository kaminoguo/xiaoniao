//go:build darwin
// +build darwin

// Package hotkey provides hotkey registration functionality for macOS
package hotkey

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

// Manager represents a hotkey manager for macOS
type Manager struct {
	mu                sync.Mutex
	registeredHotkeys map[string]*hotkey.Hotkey
	callbacks         map[string]func()
}

// NewManager creates a new hotkey manager
func NewManager() *Manager {
	return &Manager{
		registeredHotkeys: make(map[string]*hotkey.Hotkey),
		callbacks:         make(map[string]func()),
	}
}

// RegisterFromString registers a hotkey from a string representation
func (m *Manager) RegisterFromString(name, hotkeyStr string, callback func()) error {
	if hotkeyStr == "" {
		return errors.New("hotkey string cannot be empty")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Parse the hotkey string
	mods, key, err := parseHotkeyString(hotkeyStr)
	if err != nil {
		return fmt.Errorf("failed to parse hotkey: %v", err)
	}

	// Unregister existing hotkey if any
	if existing, ok := m.registeredHotkeys[name]; ok {
		if err := existing.Unregister(); err != nil {
			return fmt.Errorf("failed to unregister existing hotkey: %v", err)
		}
		delete(m.registeredHotkeys, name)
		delete(m.callbacks, name)
	}

	// Create and register new hotkey
	hk := hotkey.New(mods, key)

	// Register the hotkey on the main thread
	mainthread.Call(func() {
		err = hk.Register()
	})

	if err != nil {
		return fmt.Errorf("failed to register hotkey: %v", err)
	}

	m.registeredHotkeys[name] = hk
	m.callbacks[name] = callback

	// Start listening for the hotkey
	go func() {
		for range hk.Keydown() {
			if cb, ok := m.callbacks[name]; ok && cb != nil {
				cb()
			}
		}
	}()

	return nil
}

// Unregister removes a registered hotkey
func (m *Manager) Unregister(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	hk, ok := m.registeredHotkeys[name]
	if !ok {
		return fmt.Errorf("hotkey %s not found", name)
	}

	var err error
	mainthread.Call(func() {
		err = hk.Unregister()
	})

	if err != nil {
		return fmt.Errorf("failed to unregister hotkey: %v", err)
	}

	delete(m.registeredHotkeys, name)
	delete(m.callbacks, name)

	return nil
}

// Stop stops the hotkey manager
func (m *Manager) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for name, hk := range m.registeredHotkeys {
		mainthread.Call(func() {
			hk.Unregister()
		})
		delete(m.registeredHotkeys, name)
		delete(m.callbacks, name)
	}
}

// parseHotkeyString parses a hotkey string like "Ctrl+Alt+C" into modifiers and key
func parseHotkeyString(hotkeyStr string) ([]hotkey.Modifier, hotkey.Key, error) {
	parts := strings.Split(strings.ToUpper(hotkeyStr), "+")
	if len(parts) == 0 {
		return nil, 0, errors.New("invalid hotkey string")
	}

	var mods []hotkey.Modifier
	var key hotkey.Key

	for i, part := range parts {
		part = strings.TrimSpace(part)

		// Last part should be the key
		if i == len(parts)-1 {
			k, err := parseKey(part)
			if err != nil {
				return nil, 0, err
			}
			key = k
		} else {
			// Other parts are modifiers
			mod, err := parseModifier(part)
			if err != nil {
				return nil, 0, err
			}
			mods = append(mods, mod)
		}
	}

	return mods, key, nil
}

// parseModifier converts a string to a hotkey.Modifier
func parseModifier(s string) (hotkey.Modifier, error) {
	switch strings.ToUpper(s) {
	case "CTRL", "CONTROL":
		return hotkey.ModCtrl, nil
	case "ALT", "OPTION":
		return hotkey.ModOption, nil // Alt/Option on macOS
	case "SHIFT":
		return hotkey.ModShift, nil
	case "CMD", "COMMAND", "SUPER", "WIN", "WINDOWS":
		return hotkey.ModCmd, nil
	default:
		return 0, fmt.Errorf("unknown modifier: %s", s)
	}
}

// parseKey converts a string to a hotkey.Key
func parseKey(s string) (hotkey.Key, error) {
	// Handle single letter keys
	if len(s) == 1 {
		r := rune(s[0])
		if r >= 'A' && r <= 'Z' {
			return hotkey.Key(r), nil
		}
		if r >= '0' && r <= '9' {
			return hotkey.Key(r), nil
		}
	}

	// Handle special keys
	switch strings.ToUpper(s) {
	case "SPACE":
		return hotkey.KeySpace, nil
	case "RETURN", "ENTER":
		return hotkey.KeyReturn, nil
	case "TAB":
		return hotkey.KeyTab, nil
	case "ESC", "ESCAPE":
		return hotkey.KeyEscape, nil
	case "UP":
		return hotkey.KeyUp, nil
	case "DOWN":
		return hotkey.KeyDown, nil
	case "LEFT":
		return hotkey.KeyLeft, nil
	case "RIGHT":
		return hotkey.KeyRight, nil
	case "F1":
		return hotkey.KeyF1, nil
	case "F2":
		return hotkey.KeyF2, nil
	case "F3":
		return hotkey.KeyF3, nil
	case "F4":
		return hotkey.KeyF4, nil
	case "F5":
		return hotkey.KeyF5, nil
	case "F6":
		return hotkey.KeyF6, nil
	case "F7":
		return hotkey.KeyF7, nil
	case "F8":
		return hotkey.KeyF8, nil
	case "F9":
		return hotkey.KeyF9, nil
	case "F10":
		return hotkey.KeyF10, nil
	case "F11":
		return hotkey.KeyF11, nil
	case "F12":
		return hotkey.KeyF12, nil
	default:
		return 0, fmt.Errorf("unknown key: %s", s)
	}
}