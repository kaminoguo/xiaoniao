//go:build darwin
// +build darwin

package hotkey

import (
	"fmt"
	"strings"

	"golang.design/x/hotkey"
)

const (
	Mod1 = hotkey.Mod1
	Mod2 = hotkey.Mod2
	Mod3 = hotkey.Mod3
	Mod4 = hotkey.Mod4
	ModCtrl = hotkey.ModCtrl
	ModShift = hotkey.ModShift
	ModOption = hotkey.ModOption
	ModCmd = hotkey.ModCmd
)

type Manager struct {
	hotkeys map[string]*hotkey.Hotkey
}

func NewManager() *Manager {
	return &Manager{
		hotkeys: make(map[string]*hotkey.Hotkey),
	}
}

func (m *Manager) Register(id string, mods []hotkey.Modifier, key hotkey.Key, callback func()) error {
	var mod hotkey.Modifier
	for _, m := range mods {
		mod |= m
	}
	
	hk := hotkey.New(mods, key)
	if err := hk.Register(); err != nil {
		return fmt.Errorf("failed to register hotkey: %w", err)
	}
	
	m.hotkeys[id] = hk
	
	go func() {
		for range hk.Keydown() {
			callback()
		}
	}()
	
	return nil
}

func (m *Manager) RegisterFromString(id, keyStr string, callback func()) error {
	mods, key, err := parseHotkeyString(keyStr)
	if err != nil {
		return err
	}
	return m.Register(id, mods, key, callback)
}

func (m *Manager) Unregister(id string) error {
	if hk, ok := m.hotkeys[id]; ok {
		if err := hk.Unregister(); err != nil {
			return err
		}
		delete(m.hotkeys, id)
	}
	return nil
}

func (m *Manager) UnregisterAll() {
	for id := range m.hotkeys {
		m.Unregister(id)
	}
}

func parseHotkeyString(keyStr string) ([]hotkey.Modifier, hotkey.Key, error) {
	parts := strings.Split(strings.ToLower(keyStr), "+")
	if len(parts) < 2 {
		return nil, 0, fmt.Errorf("invalid hotkey format: %s", keyStr)
	}
	
	var mods []hotkey.Modifier
	keyPart := parts[len(parts)-1]
	
	for i := 0; i < len(parts)-1; i++ {
		switch parts[i] {
		case "ctrl", "control":
			mods = append(mods, hotkey.ModCtrl)
		case "alt", "option":
			mods = append(mods, hotkey.ModOption)
		case "shift":
			mods = append(mods, hotkey.ModShift)
		case "cmd", "command", "super":
			mods = append(mods, hotkey.ModCmd)
		default:
			return nil, 0, fmt.Errorf("unknown modifier: %s", parts[i])
		}
	}
	
	key := getKeyFromString(keyPart)
	if key == 0 {
		return nil, 0, fmt.Errorf("unknown key: %s", keyPart)
	}
	
	return mods, key, nil
}

func getKeyFromString(keyStr string) hotkey.Key {
	keyMap := map[string]hotkey.Key{
		"a": hotkey.KeyA, "b": hotkey.KeyB, "c": hotkey.KeyC, "d": hotkey.KeyD,
		"e": hotkey.KeyE, "f": hotkey.KeyF, "g": hotkey.KeyG, "h": hotkey.KeyH,
		"i": hotkey.KeyI, "j": hotkey.KeyJ, "k": hotkey.KeyK, "l": hotkey.KeyL,
		"m": hotkey.KeyM, "n": hotkey.KeyN, "o": hotkey.KeyO, "p": hotkey.KeyP,
		"q": hotkey.KeyQ, "r": hotkey.KeyR, "s": hotkey.KeyS, "t": hotkey.KeyT,
		"u": hotkey.KeyU, "v": hotkey.KeyV, "w": hotkey.KeyW, "x": hotkey.KeyX,
		"y": hotkey.KeyY, "z": hotkey.KeyZ,
		"0": hotkey.Key0, "1": hotkey.Key1, "2": hotkey.Key2, "3": hotkey.Key3,
		"4": hotkey.Key4, "5": hotkey.Key5, "6": hotkey.Key6, "7": hotkey.Key7,
		"8": hotkey.Key8, "9": hotkey.Key9,
		"f1": hotkey.KeyF1, "f2": hotkey.KeyF2, "f3": hotkey.KeyF3, "f4": hotkey.KeyF4,
		"f5": hotkey.KeyF5, "f6": hotkey.KeyF6, "f7": hotkey.KeyF7, "f8": hotkey.KeyF8,
		"f9": hotkey.KeyF9, "f10": hotkey.KeyF10, "f11": hotkey.KeyF11, "f12": hotkey.KeyF12,
		"return": hotkey.KeyReturn, "enter": hotkey.KeyReturn,
		"tab": hotkey.KeyTab, "space": hotkey.KeySpace,
		"delete": hotkey.KeyDelete, "backspace": hotkey.KeyDelete,
		"escape": hotkey.KeyEscape, "esc": hotkey.KeyEscape,
	}
	
	return keyMap[keyStr]
}