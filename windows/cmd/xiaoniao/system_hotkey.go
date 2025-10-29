package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 检测桌面环境
func detectDesktop() string {
	if desktop := os.Getenv("XDG_CURRENT_DESKTOP"); desktop != "" {
		return strings.ToLower(desktop)
	}
	if desktop := os.Getenv("DESKTOP_SESSION"); desktop != "" {
		return strings.ToLower(desktop)
	}
	return "unknown"
}

// 检测显示服务器
func detectDisplayServer() string {
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		return "wayland"
	}
	if os.Getenv("DISPLAY") != "" {
		return "x11"
	}
	return "unknown"
}

// 检查快捷键是否已被系统占用
func checkHotkeyConflict(hotkey string) (bool, string) {
	desktop := detectDesktop()
	
	// 转换快捷键格式
	gnomeKey := convertToGnomeFormat(hotkey)
	
	switch {
	case strings.Contains(desktop, "gnome"):
		// GNOME检查
		return checkGnomeHotkey(gnomeKey)
	case strings.Contains(desktop, "kde") || strings.Contains(desktop, "plasma"):
		// KDE检查
		return checkKDEHotkey(hotkey)
	case strings.Contains(desktop, "xfce"):
		// XFCE检查
		return checkXFCEHotkey(hotkey)
	default:
		// 无法检查，假设没有冲突
		return false, ""
	}
}

// 转换快捷键格式为GNOME格式
func convertToGnomeFormat(hotkey string) string {
	// Ctrl+Alt+X -> <Ctrl><Alt>x
	hotkey = strings.ReplaceAll(hotkey, "Ctrl", "<Ctrl>")
	hotkey = strings.ReplaceAll(hotkey, "Alt", "<Alt>")
	hotkey = strings.ReplaceAll(hotkey, "Shift", "<Shift>")
	hotkey = strings.ReplaceAll(hotkey, "+", "")
	
	// 最后一个字符转小写
	parts := strings.Split(hotkey, ">")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		if len(lastPart) == 1 {
			parts[len(parts)-1] = strings.ToLower(lastPart)
		}
		hotkey = strings.Join(parts, ">")
	}
	
	return hotkey
}

// 检查GNOME快捷键冲突
func checkGnomeHotkey(hotkey string) (bool, string) {
	// 检查系统快捷键
	schemas := []string{
		"org.gnome.desktop.wm.keybindings",
		"org.gnome.settings-daemon.plugins.media-keys",
		"org.gnome.shell.keybindings",
	}
	
	for _, schema := range schemas {
		cmd := exec.Command("gsettings", "list-recursively", schema)
		output, err := cmd.Output()
		if err == nil && strings.Contains(string(output), hotkey) {
			// 找到冲突
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.Contains(line, hotkey) {
					// 提取功能名称
					parts := strings.Fields(line)
					if len(parts) >= 2 {
						return true, fmt.Sprintf("系统快捷键: %s", parts[1])
					}
				}
			}
			return true, "系统快捷键"
		}
	}
	
	// 检查自定义快捷键
	cmd := exec.Command("gsettings", "get", "org.gnome.settings-daemon.plugins.media-keys", "custom-keybindings")
	output, _ := cmd.Output()
	if len(output) > 0 {
		// 解析自定义快捷键路径
		customKeys := string(output)
		if customKeys != "@as []" && customKeys != "[]" {
			// 这里可以进一步检查每个自定义快捷键，但为了简化，暂时跳过
		}
	}
	
	return false, ""
}

// 检查KDE快捷键冲突
func checkKDEHotkey(hotkey string) (bool, string) {
	// KDE的快捷键配置较复杂，简化处理
	configFile := filepath.Join(os.Getenv("HOME"), ".config", "kglobalshortcutsrc")
	if content, err := os.ReadFile(configFile); err == nil {
		if strings.Contains(string(content), hotkey) {
			return true, "系统快捷键"
		}
	}
	return false, ""
}

// 检查XFCE快捷键冲突
func checkXFCEHotkey(hotkey string) (bool, string) {
	// 使用xfconf-query检查
	cmd := exec.Command("xfconf-query", "-c", "xfce4-keyboard-shortcuts", "-l")
	output, err := cmd.Output()
	if err == nil {
		// 转换格式：Ctrl+Alt+X -> <Primary><Alt>x
		xfceKey := strings.ReplaceAll(hotkey, "Ctrl", "<Primary>")
		xfceKey = strings.ReplaceAll(xfceKey, "Alt", "<Alt>")
		xfceKey = strings.ReplaceAll(xfceKey, "Shift", "<Shift>")
		xfceKey = strings.ReplaceAll(xfceKey, "+", "")
		
		if strings.Contains(string(output), xfceKey) {
			return true, "系统快捷键"
		}
	}
	return false, ""
}

// 创建控制脚本
func createControlScripts() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	
	binDir := filepath.Join(homeDir, ".local", "bin")
	os.MkdirAll(binDir, 0755)
	
	// 切换监控脚本
	toggleScript := `#!/bin/bash
# 切换 xiaoniao 监控状态
if pgrep -f "xiaoniao run" > /dev/null; then
    pkill -USR1 -f "xiaoniao run"
fi
`
	togglePath := filepath.Join(binDir, "xiaoniao-toggle")
	if err := os.WriteFile(togglePath, []byte(toggleScript), 0755); err != nil {
		return err
	}
	
	// 切换Prompt脚本
	switchScript := `#!/bin/bash
# 切换 xiaoniao Prompt
if pgrep -f "xiaoniao run" > /dev/null; then
    pkill -USR2 -f "xiaoniao run"
fi
`
	switchPath := filepath.Join(binDir, "xiaoniao-switch")
	if err := os.WriteFile(switchPath, []byte(switchScript), 0755); err != nil {
		return err
	}
	
	return nil
}

// 配置系统快捷键
func configureSystemHotkey(function string, hotkey string) error {
	// 首先创建控制脚本
	if err := createControlScripts(); err != nil {
		return fmt.Errorf("创建控制脚本失败: %v", err)
	}
	
	desktop := detectDesktop()
	homeDir, _ := os.UserHomeDir()
	
	var command string
	switch function {
	case "toggle":
		command = filepath.Join(homeDir, ".local", "bin", "xiaoniao-toggle")
	case "switch":
		command = filepath.Join(homeDir, ".local", "bin", "xiaoniao-switch")
	default:
		return fmt.Errorf("未知功能: %s", function)
	}
	
	switch {
	case strings.Contains(desktop, "gnome"):
		return configureGnomeHotkey(function, hotkey, command)
	case strings.Contains(desktop, "kde") || strings.Contains(desktop, "plasma"):
		return configureKDEHotkey(function, hotkey, command)
	case strings.Contains(desktop, "xfce"):
		return configureXFCEHotkey(function, hotkey, command)
	default:
		// 尝试使用xbindkeys作为后备
		if detectDisplayServer() == "x11" {
			return configureXbindkeys(function, hotkey, command)
		}
		return fmt.Errorf("不支持的桌面环境: %s", desktop)
	}
}

// 配置GNOME快捷键
func configureGnomeHotkey(function, hotkey, command string) error {
	// 转换格式
	gnomeKey := convertToGnomeFormat(hotkey)
	
	// 获取现有的自定义快捷键
	cmd := exec.Command("gsettings", "get", "org.gnome.settings-daemon.plugins.media-keys", "custom-keybindings")
	output, _ := cmd.Output()
	customKeys := strings.TrimSpace(string(output))
	
	// 快捷键ID
	var id string
	var name string
	switch function {
	case "toggle":
		id = "xiaoniao1"
		name = "xiaoniao 切换监控"
	case "switch":
		id = "xiaoniao2"
		name = "xiaoniao 切换风格"
	}
	
	path := fmt.Sprintf("/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/%s/", id)
	
	// 更新自定义快捷键列表
	if customKeys == "@as []" || customKeys == "[]" {
		customKeys = fmt.Sprintf("['%s']", path)
	} else if !strings.Contains(customKeys, path) {
		// 添加到列表
		customKeys = strings.TrimSuffix(customKeys, "]")
		if customKeys != "[" {
			customKeys += ", "
		}
		customKeys += fmt.Sprintf("'%s']", path)
	}
	
	// 设置自定义快捷键列表
	exec.Command("gsettings", "set", "org.gnome.settings-daemon.plugins.media-keys", "custom-keybindings", customKeys).Run()
	
	// 设置具体的快捷键
	basePath := fmt.Sprintf("org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/%s/", id)
	exec.Command("gsettings", "set", basePath, "name", name).Run()
	exec.Command("gsettings", "set", basePath, "command", command).Run()
	exec.Command("gsettings", "set", basePath, "binding", gnomeKey).Run()
	
	return nil
}

// 配置KDE快捷键
func configureKDEHotkey(function, hotkey, command string) error {
	// KDE 使用 kwriteconfig5 配置快捷键
	homeDir, _ := os.UserHomeDir()
	
	// 创建自定义快捷键配置
	var name, id string
	switch function {
	case "toggle":
		name = "xiaoniao Toggle Monitor"
		id = "xiaoniao_toggle"
	case "switch":
		name = "xiaoniao Switch Prompt"
		id = "xiaoniao_switch"
	}
	
	// 设置快捷键
	// KDE 使用 kglobalshortcutsrc 文件
	configFile := filepath.Join(homeDir, ".config", "kglobalshortcutsrc")
	
	// 读取现有配置
	content, _ := os.ReadFile(configFile)
	lines := strings.Split(string(content), "\n")
	
	// 查找或创建 [xiaoniao] 节
	sectionFound := false
	newLines := []string{}
	inSection := false
	
	for _, line := range lines {
		if strings.HasPrefix(line, "[xiaoniao]") {
			sectionFound = true
			inSection = true
			newLines = append(newLines, line)
			continue
		}
		if strings.HasPrefix(line, "[") && inSection {
			inSection = false
		}
		if !inSection || !strings.Contains(line, id) {
			newLines = append(newLines, line)
		}
	}
	
	// 如果没有找到节，添加新节
	if !sectionFound {
		newLines = append(newLines, "", "[xiaoniao]")
	}
	
	// 添加快捷键配置
	kdeHotkey := strings.ReplaceAll(hotkey, "+", "\\t")
	entry := fmt.Sprintf("%s=%s,%s,%s", id, command, kdeHotkey, name)
	
	// 插入到 xiaoniao 节
	finalLines := []string{}
	for i, line := range newLines {
		finalLines = append(finalLines, line)
		if strings.HasPrefix(line, "[xiaoniao]") && (i+1 >= len(newLines) || strings.HasPrefix(newLines[i+1], "[")) {
			finalLines = append(finalLines, entry)
		}
	}
	
	// 写回文件
	return os.WriteFile(configFile, []byte(strings.Join(finalLines, "\n")), 0644)
}

// 配置XFCE快捷键
func configureXFCEHotkey(function, hotkey, command string) error {
	// 转换格式：Ctrl+Alt+X -> <Primary><Alt>x
	xfceKey := strings.ReplaceAll(hotkey, "Ctrl", "<Primary>")
	xfceKey = strings.ReplaceAll(xfceKey, "Alt", "<Alt>")
	xfceKey = strings.ReplaceAll(xfceKey, "Shift", "<Shift>")
	xfceKey = strings.ReplaceAll(xfceKey, "+", "")
	
	// 设置快捷键
	path := fmt.Sprintf("/commands/custom/%s", xfceKey)
	exec.Command("xfconf-query", "-c", "xfce4-keyboard-shortcuts", "-p", path, "-n", "-t", "string", "-s", command).Run()
	
	return nil
}

// 配置xbindkeys
func configureXbindkeys(function, hotkey, command string) error {
	homeDir, _ := os.UserHomeDir()
	xbindkeysrc := filepath.Join(homeDir, ".xbindkeysrc")
	
	// 读取现有配置
	content, _ := os.ReadFile(xbindkeysrc)
	lines := strings.Split(string(content), "\n")
	
	// 移除旧的xiaoniao配置
	var newLines []string
	skip := false
	for _, line := range lines {
		if strings.Contains(line, "# xiaoniao") {
			skip = true
			continue
		}
		if skip && (line == "" || strings.HasPrefix(line, "#")) {
			skip = false
		}
		if !skip {
			newLines = append(newLines, line)
		}
	}
	
	// 添加新配置
	var comment string
	switch function {
	case "toggle":
		comment = "# xiaoniao 切换监控"
	case "switch":
		comment = "# xiaoniao 切换风格"
	}
	
	newLines = append(newLines, "", comment)
	newLines = append(newLines, fmt.Sprintf(`"%s"`, command))
	newLines = append(newLines, fmt.Sprintf("    %s", hotkey))
	
	// 写回文件
	return os.WriteFile(xbindkeysrc, []byte(strings.Join(newLines, "\n")), 0644)
}

// 移除系统快捷键配置
func removeSystemHotkey(function string) error {
	desktop := detectDesktop()
	
	switch {
	case strings.Contains(desktop, "gnome"):
		// GNOME：清除特定快捷键
		var id string
		switch function {
		case "toggle":
			id = "xiaoniao1"
		case "switch":
			id = "xiaoniao2"
		}
		basePath := fmt.Sprintf("org.gnome.settings-daemon.plugins.media-keys.custom-keybinding:/org/gnome/settings-daemon/plugins/media-keys/custom-keybindings/%s/", id)
		// 使用 set 命令将 binding 设置为空字符串
		exec.Command("gsettings", "set", basePath, "binding", "''").Run()
	}
	
	return nil
}