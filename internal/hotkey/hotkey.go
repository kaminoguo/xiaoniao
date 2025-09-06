package hotkey

import (
	"fmt"
	"strings"
	"sync"
	
	"golang.design/x/hotkey"
)

// Hotkey represents a global hotkey
type Hotkey struct {
	ID       string
	Key      string
	Modifier string
	Callback func()
	hk       *hotkey.Hotkey // 实际的热键对象
}

// Manager manages global hotkeys
type Manager struct {
	mu      sync.RWMutex
	hotkeys map[string]*Hotkey
	enabled bool
}

// NewManager creates a new hotkey manager
func NewManager() *Manager {
	return &Manager{
		hotkeys: make(map[string]*Hotkey),
		enabled: false,
	}
}

// parseModifiers 解析修饰键字符串
func parseModifiers(modStr string) []hotkey.Modifier {
	var mods []hotkey.Modifier
	parts := strings.Split(modStr, "+")
	
	for _, part := range parts {
		switch strings.ToLower(strings.TrimSpace(part)) {
		case "ctrl", "control":
			mods = append(mods, hotkey.ModCtrl)
		case "alt":
			mods = append(mods, hotkey.Mod1) // Alt key
		case "shift":
			mods = append(mods, hotkey.ModShift)
		case "win", "super", "cmd":
			mods = append(mods, hotkey.Mod4) // Super/Win key
		}
	}
	
	return mods
}

// parseKey 解析按键
func parseKey(keyStr string) hotkey.Key {
	// 移除修饰键部分，获取实际按键
	parts := strings.Split(keyStr, "+")
	actualKey := strings.ToUpper(strings.TrimSpace(parts[len(parts)-1]))
	
	// 字母键
	if len(actualKey) == 1 && actualKey[0] >= 'A' && actualKey[0] <= 'Z' {
		return hotkey.Key(actualKey[0])
	}
	
	// 数字键
	if len(actualKey) == 1 && actualKey[0] >= '0' && actualKey[0] <= '9' {
		return hotkey.Key(actualKey[0])
	}
	
	// 功能键
	switch actualKey {
	case "F1":
		return hotkey.KeyF1
	case "F2":
		return hotkey.KeyF2
	case "F3":
		return hotkey.KeyF3
	case "F4":
		return hotkey.KeyF4
	case "F5":
		return hotkey.KeyF5
	case "F6":
		return hotkey.KeyF6
	case "F7":
		return hotkey.KeyF7
	case "F8":
		return hotkey.KeyF8
	case "F9":
		return hotkey.KeyF9
	case "F10":
		return hotkey.KeyF10
	case "F11":
		return hotkey.KeyF11
	case "F12":
		return hotkey.KeyF12
	case "SPACE":
		return hotkey.KeySpace
	case "RETURN", "ENTER":
		return hotkey.KeyReturn
	case "TAB":
		return hotkey.KeyTab
	case "DELETE":
		return hotkey.KeyDelete
	case "ESCAPE", "ESC":
		return hotkey.KeyEscape
	case "UP":
		return hotkey.KeyUp
	case "DOWN":
		return hotkey.KeyDown
	case "LEFT":
		return hotkey.KeyLeft
	case "RIGHT":
		return hotkey.KeyRight
	default:
		// 默认返回空格键
		return hotkey.KeySpace
	}
}

// RegisterFromString 从字符串注册快捷键（如 "Ctrl+Shift+X"）
func (m *Manager) RegisterFromString(id, hotkeyStr string, callback func()) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.hotkeys[id]; exists {
		// 先注销旧的
		if oldHk := m.hotkeys[id]; oldHk != nil && oldHk.hk != nil {
			oldHk.hk.Unregister()
		}
	}

	// 解析快捷键字符串
	parts := strings.Split(hotkeyStr, "+")
	if len(parts) < 2 {
		return fmt.Errorf("invalid hotkey format: %s", hotkeyStr)
	}
	
	// 获取修饰键（除了最后一个部分）
	modStr := strings.Join(parts[:len(parts)-1], "+")
	mods := parseModifiers(modStr)
	
	// 获取实际按键
	key := parseKey(hotkeyStr)
	
	// 创建热键
	hk := hotkey.New(mods, key)
	
	// 注册热键
	if err := hk.Register(); err != nil {
		return fmt.Errorf("failed to register hotkey %s: %v", hotkeyStr, err)
	}
	
	// 启动监听协程
	go func() {
		for range hk.Keydown() {
			if callback != nil {
				callback()
			}
		}
	}()
	
	// 保存到管理器
	m.hotkeys[id] = &Hotkey{
		ID:       id,
		Key:      hotkeyStr,
		Modifier: modStr,
		Callback: callback,
		hk:       hk,
	}
	
	fmt.Printf("Hotkey registered: %s (ID: %s)\n", hotkeyStr, id)
	return nil
}

// Register registers a new hotkey (保留兼容性)
func (m *Manager) Register(id, key, modifier string, callback func()) error {
	hotkeyStr := modifier + "+" + key
	return m.RegisterFromString(id, hotkeyStr, callback)
}

// Unregister unregisters a hotkey
func (m *Manager) Unregister(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	hk, exists := m.hotkeys[id]
	if !exists {
		return fmt.Errorf("hotkey %s not found", id)
	}

	// 注销实际的热键
	if hk.hk != nil {
		hk.hk.Unregister()
	}
	
	delete(m.hotkeys, id)
	
	fmt.Printf("Hotkey unregistered: %s\n", id)
	
	return nil
}

// Enable enables all hotkeys
func (m *Manager) Enable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.enabled = true
	fmt.Println("Hotkeys enabled")
}

// Disable disables all hotkeys
func (m *Manager) Disable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.enabled = false
	fmt.Println("Hotkeys disabled")
}

// IsEnabled returns whether hotkeys are enabled
func (m *Manager) IsEnabled() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.enabled
}

// RegisterDefaultHotkeys registers the default hotkeys for the application
func (m *Manager) RegisterDefaultHotkeys(onTranslate, onSwitch, onToggleMonitor func()) error {
	// Alt+T for translate
	if err := m.Register("translate", "T", "Alt", onTranslate); err != nil {
		return err
	}

	// Alt+Q for switch (original/translation)
	if err := m.Register("switch", "Q", "Alt", onSwitch); err != nil {
		return err
	}

	// Alt+M for toggle monitor
	if err := m.Register("toggle_monitor", "M", "Alt", onToggleMonitor); err != nil {
		return err
	}

	return nil
}

// GetRegisteredHotkeys returns all registered hotkeys
func (m *Manager) GetRegisteredHotkeys() []Hotkey {
	m.mu.RLock()
	defer m.mu.RUnlock()

	hotkeys := make([]Hotkey, 0, len(m.hotkeys))
	for _, hk := range m.hotkeys {
		hotkeys = append(hotkeys, *hk)
	}

	return hotkeys
}

// Note: In a production implementation, we would use a library like:
// - github.com/go-vgo/robotgo for cross-platform hotkey support
// - Or platform-specific implementations using CGO
// For now, this is a mock implementation that demonstrates the API