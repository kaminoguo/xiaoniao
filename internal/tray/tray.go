package tray

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
)

// Status represents the tray icon status
type Status string

const (
	StatusIdle        Status = "idle"        // Blue bird (normal)
	StatusTranslating Status = "translating" // Green bird (processing)
	StatusError       Status = "error"       // Red bird (error)
)

// Manager manages the system tray
type Manager struct {
	status            Status
	visible           bool
	isMonitoring      bool
	translationCount  int
	currentPromptName string
	onQuit            func()
	onShow            func()
	onSettings        func()
	onToggleMonitor   func(bool)
	onRefresh         func()
	onSwitchPrompt    func()
	onSelectPrompt    func(string) // Callback for selecting a specific prompt
	onToggleTerminal  func()
	// Menu items
	mToggle     *systray.MenuItem
	mPromptInfo *systray.MenuItem
	mStatus     *systray.MenuItem
	mRefresh    *systray.MenuItem
	mPromptMenu *systray.MenuItem // Main prompt menu
	promptItems []*systray.MenuItem // Individual prompt menu items
}

// NewManager creates a new tray manager
func NewManager() *Manager {
	return &Manager{
		status:       StatusIdle,
		visible:      true,
		isMonitoring: false,
	}
}

// SetStatus sets the tray icon status (changes color)
func (m *Manager) SetStatus(status Status) {
	m.status = status
	
	// Load appropriate icon based on status
	configDir, _ := os.UserConfigDir()
	var iconPath string
	
	// 获取当前风格名称
	t := i18n.T()
	promptName := m.currentPromptName
	if promptName == "" {
		promptName = t.NotSet
	}
	
	// 如果监控已关闭，显示红色图标
	if !m.isMonitoring && status != StatusError {
		iconPath = filepath.Join(configDir, "xiaoniao", "icon_red.png")
		systray.SetTooltip(fmt.Sprintf("xiaoniao - %s | %s: %s", t.MonitorStopped, t.TranslateStyle, promptName))
		systray.SetTitle("")  // 不显示额外标记
	} else {
		switch status {
		case StatusTranslating:
			iconPath = filepath.Join(configDir, "xiaoniao", "icon_green.png")
			systray.SetTooltip(fmt.Sprintf("xiaoniao - %s | %s: %s", t.Translating, t.TranslateStyle, promptName))
			systray.SetTitle("")
		case StatusError:
			iconPath = filepath.Join(configDir, "xiaoniao", "icon_red.png")
			systray.SetTooltip(fmt.Sprintf("xiaoniao - %s | %s: %s", t.Failed, t.TranslateStyle, promptName))
			systray.SetTitle("")
		default: // StatusIdle
			iconPath = filepath.Join(configDir, "xiaoniao", "icon_blue.png")
			systray.SetTooltip(fmt.Sprintf("xiaoniao - %s (%s %d %s) | %s: %s", t.Monitoring, t.TotalCount, m.translationCount, t.TranslateCount, t.TranslateStyle, promptName))
			systray.SetTitle("")  // 不显示额外标记
		}
	}
	
	// Update icon if file exists
	if iconData, err := os.ReadFile(iconPath); err == nil {
		systray.SetIcon(iconData)
	}
}

// GetStatus returns the current status
func (m *Manager) GetStatus() Status {
	return m.status
}

// SetVisible shows or hides the tray icon
func (m *Manager) SetVisible(visible bool) {
	m.visible = visible
}

// IsVisible returns whether the tray is visible
func (m *Manager) IsVisible() bool {
	return m.visible
}

// SetOnQuit sets the quit callback
func (m *Manager) SetOnQuit(callback func()) {
	m.onQuit = callback
}

// SetOnShow sets the show window callback
func (m *Manager) SetOnShow(callback func()) {
	m.onShow = callback
}

// SetOnSettings sets the settings callback
func (m *Manager) SetOnSettings(callback func()) {
	m.onSettings = callback
}

// SetOnToggleMonitor sets the monitor toggle callback
func (m *Manager) SetOnToggleMonitor(callback func(bool)) {
	m.onToggleMonitor = callback
}


// SetOnRefresh sets the refresh callback
func (m *Manager) SetOnRefresh(callback func()) {
	m.onRefresh = callback
}

// SetOnSwitchPrompt sets the prompt switch callback
func (m *Manager) SetOnSwitchPrompt(callback func()) {
	m.onSwitchPrompt = callback
}

// SetOnSelectPrompt sets the prompt selection callback
func (m *Manager) SetOnSelectPrompt(callback func(string)) {
	m.onSelectPrompt = callback
}

// SetOnToggleTerminal sets the terminal toggle callback
func (m *Manager) SetOnToggleTerminal(callback func()) {
	m.onToggleTerminal = callback
}

// Initialize initializes the system tray
func (m *Manager) Initialize() error {
	go systray.Run(m.onReady, m.onExit)
	return nil
}


func (m *Manager) onReady() {
	// 只显示图标，不显示标题
	systray.SetTitle("")
	systray.SetTooltip("xiaoniao")
	
	// Load blue icon initially
	configDir, _ := os.UserConfigDir()
	iconPath := filepath.Join(configDir, "xiaoniao", "icon_blue.png")
	if iconData, err := os.ReadFile(iconPath); err == nil {
		systray.SetIcon(iconData)
	} else {
		// Fallback to default icon if blue not available
		m.loadIcon()
	}
	
	// Create menu items
	t := i18n.T()
	m.mToggle = systray.AddMenuItemCheckbox(t.TrayToggle, t.TrayToggle, m.isMonitoring)
	
	// 显示当前 prompt
	promptLabel := fmt.Sprintf("%s: %s", t.TranslateStyle, m.currentPromptName)
	if m.currentPromptName == "" {
		promptLabel = fmt.Sprintf("%s: %s", t.TranslateStyle, t.NotSet)
	}
	m.mPromptInfo = systray.AddMenuItem(promptLabel, t.TranslateStyle)
	m.mPromptInfo.Disable() // 这个只是显示，不能点击
	
	systray.AddSeparator()
	
	m.mRefresh = systray.AddMenuItem(t.TrayRefresh, t.TrayRefresh)
	mConfig := systray.AddMenuItem(t.TraySettings, t.TraySettings)
	m.mPromptMenu = systray.AddMenuItem(t.TranslateStyle, t.TranslateStyle)
	mTerminal := systray.AddMenuItem(t.TrayShow+"/"+t.TrayHide, t.TrayShow+"/"+t.TrayHide)
	systray.AddSeparator()
	
	mAbout := systray.AddMenuItem(t.TrayAbout, t.TrayAbout)
	mQuit := systray.AddMenuItem(t.TrayQuit, t.TrayQuit)
	
	// Handle menu events
	go func() {
		for {
			select {
			case <-m.mToggle.ClickedCh:
				m.toggleMonitor()
			case <-m.mRefresh.ClickedCh:
				m.refreshConfig()
			case <-mConfig.ClickedCh:
				m.openSettings()
			case <-mTerminal.ClickedCh:
				m.toggleTerminal()
			case <-mAbout.ClickedCh:
				m.showAbout()
			case <-mQuit.ClickedCh:
				m.quit()
				return
			}
		}
	}()
}

func (m *Manager) onExit() {
	// Cleanup
}

func (m *Manager) toggleMonitor() {
	m.isMonitoring = !m.isMonitoring
	
	if m.isMonitoring {
		m.mToggle.Check()
	} else {
		m.mToggle.Uncheck()
	}
	// 更新图标状态
	m.SetStatus(StatusIdle)
	
	if m.onToggleMonitor != nil {
		m.onToggleMonitor(m.isMonitoring)
	}
}


func (m *Manager) openSettings() {
	if m.onSettings != nil {
		m.onSettings()
	} else {
		// 在新终端窗口中打开配置界面
		switch runtime.GOOS {
		case "linux":
			// 尝试不同的终端模拟器
			terminals := [][]string{
				{"gnome-terminal", "--", "xiaoniao", "config"},
				{"konsole", "-e", "xiaoniao", "config"},
				{"xfce4-terminal", "-e", "xiaoniao config"},
				{"xterm", "-e", "xiaoniao", "config"},
				{"kitty", "xiaoniao", "config"},
				{"alacritty", "-e", "xiaoniao", "config"},
			}
			
			for _, term := range terminals {
				cmd := exec.Command(term[0], term[1:]...)
				if err := cmd.Start(); err == nil {
					return
				}
			}
			
			// 如果都失败了，尝试通过 x-terminal-emulator（Debian/Ubuntu）
			exec.Command("x-terminal-emulator", "-e", "xiaoniao", "config").Start()
			
		case "darwin":
			// macOS: 使用 Terminal.app
			script := `tell application "Terminal" to do script "xiaoniao config"`
			exec.Command("osascript", "-e", script).Start()
			
		case "windows":
			// Windows: 使用 cmd
			exec.Command("cmd", "/c", "start", "cmd", "/k", "xiaoniao", "config").Start()
		}
	}
}

func (m *Manager) showAbout() {
	// 直接打开关于页面
	switch runtime.GOOS {
	case "linux":
		// 尝试不同的终端模拟器，使用about命令直接打开关于页面
		terminals := [][]string{
			// ptyxis 优先（Fedora的新默认终端）
			{"ptyxis", "-x", "xiaoniao about"},
			{"gnome-terminal", "--", "xiaoniao", "about"},
			{"konsole", "-e", "xiaoniao", "about"},
			{"xfce4-terminal", "-e", "xiaoniao about"},
			{"xterm", "-e", "xiaoniao", "about"},
			{"kitty", "xiaoniao", "about"},
			{"alacritty", "-e", "xiaoniao", "about"},
		}
		
		for _, term := range terminals {
			cmd := exec.Command(term[0], term[1:]...)
			if err := cmd.Start(); err == nil {
				return
			}
		}
		
		// 如果都失败了，尝试通过 x-terminal-emulator（Debian/Ubuntu）
		exec.Command("x-terminal-emulator", "-e", "xiaoniao", "about").Start()
		
	case "darwin":
		// macOS: 使用 Terminal.app
		script := `tell application "Terminal" to do script "xiaoniao about"`
		exec.Command("osascript", "-e", script).Start()
		
	case "windows":
		// Windows: 使用 cmd
		exec.Command("cmd", "/c", "start", "cmd", "/k", "xiaoniao", "about").Start()
	}
}

func (m *Manager) refreshConfig() {
	if m.onRefresh != nil {
		m.onRefresh()
		// 不显示通知
	}
}

func (m *Manager) switchPrompt() {
	if m.onSwitchPrompt != nil {
		m.onSwitchPrompt()
	}
}

func (m *Manager) toggleTerminal() {
	if m.onToggleTerminal != nil {
		m.onToggleTerminal()
	}
}

func (m *Manager) quit() {
	if m.onQuit != nil {
		m.onQuit()
	}
	systray.Quit()
}

func (m *Manager) loadIcon() {
	// Try to find and load icon
	iconPaths := []string{
		filepath.Join(os.Getenv("HOME"), ".config/xiaoniao/icon.png"),
		"/usr/share/icons/xiaoniao.png",
		"/usr/local/share/icons/xiaoniao.png",
		"./assets/icon.png",
	}
	
	for _, path := range iconPaths {
		if data, err := os.ReadFile(path); err == nil {
			systray.SetIcon(data)
			return
		}
	}
	
	// If no icon found, try to use embedded icon data
	// For now, we'll skip this
}

// Quit quits the tray
func (m *Manager) Quit() {
	m.quit()
}

// ShowNotification shows a system notification
func (m *Manager) ShowNotification(title, message string) {
	// 不再显示任何通知
}

// IncrementTranslationCount increments the translation counter
func (m *Manager) IncrementTranslationCount() {
	m.translationCount++
	if m.mStatus != nil {
		m.mStatus.SetTitle(fmt.Sprintf("状态: 已翻译 %d 次", m.translationCount))
	}
	m.SetStatus(m.status) // Update tooltip
}

// UpdateMonitorStatus updates the monitor status in UI
func (m *Manager) UpdateMonitorStatus(running bool) {
	t := i18n.T()
	m.isMonitoring = running
	if running {
		m.mToggle.SetTitle(fmt.Sprintf("[||] %s", t.StopMonitor))
		m.mToggle.Check()
	} else {
		m.mToggle.SetTitle(fmt.Sprintf("[>] %s", t.StartMonitor))
		m.mToggle.Uncheck()
	}
	m.SetStatus(StatusIdle)
}

// SetCurrentPrompt 设置当前 prompt 显示
func (m *Manager) SetCurrentPrompt(promptName string) {
	t := i18n.T()
	m.currentPromptName = promptName
	if m.mPromptInfo != nil {
		promptLabel := fmt.Sprintf("%s: %s", t.TranslateStyle, promptName)
		if promptName == "" {
			promptLabel = fmt.Sprintf("%s: %s", t.TranslateStyle, t.NotSet)
		}
		m.mPromptInfo.SetTitle(promptLabel)
	}
	
	// 同时更新托盘图标的tooltip，这样不用打开菜单也能看到
	if promptName == "" {
		promptName = "默认"
	}
	if m.isMonitoring {
		systray.SetTooltip(fmt.Sprintf("xiaoniao - 监控中 | 风格: %s", promptName))
	} else {
		systray.SetTooltip(fmt.Sprintf("xiaoniao - 已停止 | 风格: %s", promptName))
	}
}

// UpdatePromptList updates the prompt submenu with available prompts
func (m *Manager) UpdatePromptList(prompts []struct{ ID, Name string }) {
	// Clear existing prompt items
	m.promptItems = make([]*systray.MenuItem, 0, len(prompts))
	
	// Add each prompt as a submenu item
	if m.mPromptMenu != nil {
		t := i18n.T()
		for _, prompt := range prompts {
			item := m.mPromptMenu.AddSubMenuItem(prompt.Name, fmt.Sprintf("%s %s", t.SwitchPrompt, prompt.Name))
			m.promptItems = append(m.promptItems, item)
			
			// Handle clicks in a goroutine
			go func(promptID string) {
				for {
					<-item.ClickedCh
					if m.onSelectPrompt != nil {
						m.onSelectPrompt(promptID)
					}
				}
			}(prompt.ID)
		}
	}
}

