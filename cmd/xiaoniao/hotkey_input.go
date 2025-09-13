package main

import (
	"fmt"
	"strings"
)

// ValidateHotkeyString validates a hotkey string format
func ValidateHotkeyString(hotkeyStr string) error {
	if hotkeyStr == "" {
		return nil // Empty is allowed (no hotkey)
	}

	// Special case: single modifier keys
	singleMods := []string{"ctrl", "alt", "shift", "win", "cmd", "meta"}
	lowerStr := strings.ToLower(strings.TrimSpace(hotkeyStr))
	for _, mod := range singleMods {
		if lowerStr == mod {
			return nil // Valid single modifier
		}
	}

	// Parse combination format: modifier+modifier+key
	parts := strings.Split(hotkeyStr, "+")
	if len(parts) < 2 {
		// Check if it's a single key (F1-F12, A-Z, 0-9)
		if isValidSingleKey(hotkeyStr) {
			return nil
		}
		return fmt.Errorf("格式错误：请使用 '修饰键+主键' 格式，如 'Ctrl+C'")
	}

	// Validate modifiers (all parts except the last one)
	validModifiers := map[string]bool{
		"ctrl": true, "control": true,
		"alt": true,
		"shift": true,
		"win": true, "cmd": true, "meta": true, "super": true,
	}

	for i := 0; i < len(parts)-1; i++ {
		mod := strings.ToLower(strings.TrimSpace(parts[i]))
		if !validModifiers[mod] {
			return fmt.Errorf("无效的修饰键: %s", parts[i])
		}
	}

	// Validate the main key (last part)
	mainKey := strings.TrimSpace(parts[len(parts)-1])
	if !isValidMainKey(mainKey) {
		return fmt.Errorf("无效的主键: %s", mainKey)
	}

	return nil
}

// isValidSingleKey checks if a string is a valid single key
func isValidSingleKey(key string) bool {
	key = strings.ToUpper(strings.TrimSpace(key))

	// F1-F12
	if len(key) >= 2 && key[0] == 'F' {
		if len(key) == 2 && key[1] >= '1' && key[1] <= '9' {
			return true
		}
		if key == "F10" || key == "F11" || key == "F12" {
			return true
		}
	}

	// A-Z
	if len(key) == 1 && key[0] >= 'A' && key[0] <= 'Z' {
		return true
	}

	// 0-9
	if len(key) == 1 && key[0] >= '0' && key[0] <= '9' {
		return true
	}

	return false
}

// isValidMainKey checks if a string is a valid main key
func isValidMainKey(key string) bool {
	key = strings.ToUpper(strings.TrimSpace(key))

	// Single character keys
	if isValidSingleKey(key) {
		return true
	}

	// Special keys
	specialKeys := []string{
		"SPACE", "ENTER", "RETURN", "TAB", "ESCAPE", "ESC",
		"BACKSPACE", "DELETE", "INSERT", "HOME", "END",
		"PAGEUP", "PAGEDOWN", "UP", "DOWN", "LEFT", "RIGHT",
		"PLUS", "MINUS", "COMMA", "PERIOD", "SLASH",
	}

	for _, sk := range specialKeys {
		if key == sk {
			return true
		}
	}

	return false
}

// NormalizeHotkeyString normalizes a hotkey string to a standard format
func NormalizeHotkeyString(hotkeyStr string) string {
	if hotkeyStr == "" {
		return ""
	}

	// Handle single modifier keys
	singleMods := map[string]string{
		"ctrl": "Ctrl", "control": "Ctrl",
		"alt": "Alt",
		"shift": "Shift",
		"win": "Win", "cmd": "Win", "meta": "Win", "super": "Win",
	}

	lowerStr := strings.ToLower(strings.TrimSpace(hotkeyStr))
	if normalized, ok := singleMods[lowerStr]; ok {
		return normalized
	}

	// Handle combinations
	parts := strings.Split(hotkeyStr, "+")
	if len(parts) == 1 {
		// Single key, just uppercase it
		return strings.ToUpper(strings.TrimSpace(hotkeyStr))
	}

	// Normalize modifiers
	var normalizedParts []string
	modOrder := []string{"Ctrl", "Alt", "Shift", "Win"} // Standard order
	modPresent := make(map[string]bool)

	// Mark which modifiers are present
	for i := 0; i < len(parts)-1; i++ {
		mod := strings.ToLower(strings.TrimSpace(parts[i]))
		switch mod {
		case "ctrl", "control":
			modPresent["Ctrl"] = true
		case "alt":
			modPresent["Alt"] = true
		case "shift":
			modPresent["Shift"] = true
		case "win", "cmd", "meta", "super":
			modPresent["Win"] = true
		}
	}

	// Add modifiers in standard order
	for _, mod := range modOrder {
		if modPresent[mod] {
			normalizedParts = append(normalizedParts, mod)
		}
	}

	// Add the main key (uppercase)
	mainKey := strings.ToUpper(strings.TrimSpace(parts[len(parts)-1]))
	normalizedParts = append(normalizedParts, mainKey)

	return strings.Join(normalizedParts, "+")
}

// GetHotkeyExamples returns example hotkey combinations
func GetHotkeyExamples() []string {
	return []string{
		"Ctrl+C", "Alt+Tab", "Shift+F1",
		"Ctrl", "Alt", "Shift",
		"F1", "F2", "F12",
		"Ctrl+Alt+S", "Win+D", "Ctrl+Shift+N",
	}
}

// GetHotkeyHelp returns help text for hotkey input
func GetHotkeyHelp() string {
	return `输入格式说明：
• 修饰键+主键：Ctrl+C, Alt+Tab, Shift+F1
• 单个修饰键：Ctrl, Alt, Shift, Win
• 功能键：F1-F12
• 字母数字：A-Z, 0-9
• 组合键：Ctrl+Alt+S, Ctrl+Shift+N`
}