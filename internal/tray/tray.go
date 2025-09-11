package tray

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	isReady           bool  // Whether tray is initialized
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
	
	// Only update if tray is ready
	if !m.isReady {
		return
	}
	
	// Load appropriate icon based on status
	configDir, _ := os.UserConfigDir()
	var iconPath string
	
	// 获取当前风格名称
	t := i18n.T()
	promptName := m.currentPromptName
	if promptName == "" {
		promptName = t.NotSet
	}
	
	// Determine which color icon to use
	var iconColor string
	
	// 如果监控已关闭，显示红色图标
	if !m.isMonitoring && status != StatusError {
		iconPath = filepath.Join(configDir, "xiaoniao", "icon_red.png")
		iconColor = "red"
		systray.SetTooltip(fmt.Sprintf("xiaoniao - %s | %s: %s", t.MonitorStopped, t.TranslateStyle, promptName))
		systray.SetTitle("")  // 不显示额外标记
	} else {
		switch status {
		case StatusTranslating:
			iconPath = filepath.Join(configDir, "xiaoniao", "icon_green.png")
			iconColor = "green"
			systray.SetTooltip(fmt.Sprintf("xiaoniao - %s | %s: %s", t.Translating, t.TranslateStyle, promptName))
			systray.SetTitle("")
		case StatusError:
			iconPath = filepath.Join(configDir, "xiaoniao", "icon_red.png")
			iconColor = "red"
			systray.SetTooltip(fmt.Sprintf("xiaoniao - %s | %s: %s", t.Failed, t.TranslateStyle, promptName))
			systray.SetTitle("")
		default: // StatusIdle
			iconPath = filepath.Join(configDir, "xiaoniao", "icon_blue.png")
			iconColor = "blue"
			systray.SetTooltip(fmt.Sprintf("xiaoniao - %s (%s %d %s) | %s: %s", t.Monitoring, t.TotalCount, m.translationCount, t.TranslateCount, t.TranslateStyle, promptName))
			systray.SetTitle("")  // 不显示额外标记
		}
	}
	
	// Try to load icon from file first, fallback to embedded icon
	if iconData, err := os.ReadFile(iconPath); err == nil {
		systray.SetIcon(iconData)
	} else {
		// Use embedded icon with appropriate color
		systray.SetIcon(GetIconForStatus(iconColor))
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
	// Mark as ready before any systray operations
	m.isReady = true
	
	// 只显示图标，不显示标题
	systray.SetTitle("")
	systray.SetTooltip("xiaoniao")
	
	// Load blue icon initially
	configDir, _ := os.UserConfigDir()
	iconPath := filepath.Join(configDir, "xiaoniao", "icon_blue.png")
	if iconData, err := os.ReadFile(iconPath); err == nil {
		systray.SetIcon(iconData)
	} else {
		// Use embedded default icon
		systray.SetIcon(GetDefaultIcon())
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
		// Windows: 使用 cmd 在新终端窗口中打开配置界面
		exec.Command("cmd", "/c", "start", "cmd", "/k", "xiaoniao", "config").Start()
	}
}

func (m *Manager) showAbout() {
	// Windows: 使用 cmd 直接打开关于页面
	exec.Command("cmd", "/c", "start", "cmd", "/k", "xiaoniao", "about").Start()
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
	// Try to find and load icon from file system
	iconPaths := []string{
		filepath.Join(os.Getenv("HOME"), ".config/xiaoniao/icon.png"),
		filepath.Join(os.Getenv("USERPROFILE"), ".config/xiaoniao/icon.png"), // Windows
		filepath.Join(os.Getenv("APPDATA"), "xiaoniao/icon.png"), // Windows AppData
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
	
	// If no icon found, use embedded icon data
	systray.SetIcon(GetDefaultIcon())
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
	
	// Only update menu items if they exist
	if m.mToggle != nil {
		if running {
			m.mToggle.SetTitle(fmt.Sprintf("[||] %s", t.StopMonitor))
			m.mToggle.Check()
		} else {
			m.mToggle.SetTitle(fmt.Sprintf("[>] %s", t.StartMonitor))
			m.mToggle.Uncheck()
		}
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
	// 只有在托盘已初始化后才更新tooltip
	if m.isReady {
		if promptName == "" {
			promptName = "默认"
		}
		if m.isMonitoring {
			systray.SetTooltip(fmt.Sprintf("xiaoniao - 监控中 | 风格: %s", promptName))
		} else {
			systray.SetTooltip(fmt.Sprintf("xiaoniao - 已停止 | 风格: %s", promptName))
		}
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

