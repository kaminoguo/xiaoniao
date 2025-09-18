package tray

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/getlantern/systray"
)

// Status represents the tray icon status
type Status string

const (
	StatusIdle        Status = "idle"        // Blue bird (normal)
	StatusTranslating Status = "translating" // Green bird (processing)
	StatusError       Status = "error"       // Red bird (error)
)

// Manager manages the system tray for macOS
type Manager struct {
	mu                sync.Mutex
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
	businessLogic     func()

	// Menu items
	statusItem      *systray.MenuItem
	monitorItem     *systray.MenuItem
	promptItem      *systray.MenuItem
	promptSubItems  []*systray.MenuItem
	refreshItem     *systray.MenuItem
	settingsItem    *systray.MenuItem
	terminalItem    *systray.MenuItem
	quitItem        *systray.MenuItem
}

// NewManager creates a new tray manager for macOS
func NewManager() (*Manager, error) {
	// Ensure we're on the main thread for macOS
	runtime.LockOSThread()

	return &Manager{
		status:  StatusIdle,
		visible: true,
		isReady: false,
	}, nil
}

// Initialize initializes the tray
func (m *Manager) Initialize() error {
	// Start systray - will be called from mainthread.Init in main
	return nil
}

// Run starts the systray (must be called from main thread)
func (m *Manager) Run() {
	systray.Run(m.onReady, m.onExit)
}

// onReady is called when systray is ready
func (m *Manager) onReady() {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Set icon based on current status
	m.updateIcon()

	// Set tooltip
	systray.SetTooltip("小鸟翻译 - 剪贴板翻译工具")

	// Create menu items
	m.statusItem = systray.AddMenuItem("状态: 就绪", "")
	m.statusItem.Disable()

	systray.AddSeparator()

	m.monitorItem = systray.AddMenuItem("▶ 开始监控", "切换剪贴板监控")

	m.promptItem = systray.AddMenuItem("提示词: 默认", "切换提示词")

	m.refreshItem = systray.AddMenuItem("🔄 刷新配置", "重新加载配置")

	systray.AddSeparator()

	m.settingsItem = systray.AddMenuItem("⚙️ 设置", "打开设置")

	m.terminalItem = systray.AddMenuItem("📟 显示终端", "显示/隐藏终端")

	systray.AddSeparator()

	m.quitItem = systray.AddMenuItem("退出", "退出程序")

	m.isReady = true

	// Start business logic if set
	if m.businessLogic != nil {
		go m.businessLogic()
	}

	// Handle menu item clicks
	go m.handleMenuClicks()
}

// onExit is called when systray exits
func (m *Manager) onExit() {
	// Cleanup
}

// handleMenuClicks handles menu item click events
func (m *Manager) handleMenuClicks() {
	for {
		select {
		case <-m.monitorItem.ClickedCh:
			m.mu.Lock()
			m.isMonitoring = !m.isMonitoring
			m.UpdateMonitorStatus(m.isMonitoring)
			if m.onToggleMonitor != nil {
				go m.onToggleMonitor(m.isMonitoring)
			}
			m.mu.Unlock()

		case <-m.refreshItem.ClickedCh:
			if m.onRefresh != nil {
				go m.onRefresh()
			}

		case <-m.settingsItem.ClickedCh:
			if m.onSettings != nil {
				go m.onSettings()
			}

		case <-m.terminalItem.ClickedCh:
			if m.onToggleTerminal != nil {
				go m.onToggleTerminal()
			}

		case <-m.quitItem.ClickedCh:
			if m.onQuit != nil {
				m.onQuit()
			}
			systray.Quit()
		}
	}
}

// updateIcon updates the tray icon based on status
func (m *Manager) updateIcon() {
	// For now, use a simple default icon
	// TODO: Add proper icon resources
	systray.SetIcon(iconDefault)
}

// SetStatus sets the tray icon status
func (m *Manager) SetStatus(status Status) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.status = status
	if m.isReady {
		m.updateIcon()

		// Update status text
		statusText := "状态: 就绪"
		switch status {
		case StatusTranslating:
			statusText = "状态: 翻译中..."
		case StatusError:
			statusText = "状态: 错误"
		}
		m.statusItem.SetTitle(statusText)
	}
}

// SetCurrentPrompt sets the current prompt name
func (m *Manager) SetCurrentPrompt(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.currentPromptName = name
	if m.isReady && m.promptItem != nil {
		m.promptItem.SetTitle(fmt.Sprintf("提示词: %s", name))
	}
}

// IncrementTranslationCount increments the translation counter
func (m *Manager) IncrementTranslationCount() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.translationCount++
	if m.isReady && m.statusItem != nil {
		m.statusItem.SetTitle(fmt.Sprintf("已翻译: %d 次", m.translationCount))
	}
}

// UpdateMonitorStatus updates monitoring status
func (m *Manager) UpdateMonitorStatus(isMonitoring bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.isMonitoring = isMonitoring
	if m.isReady && m.monitorItem != nil {
		if isMonitoring {
			m.monitorItem.SetTitle("⏸ 暂停监控")
			systray.SetTooltip("小鸟翻译 - 监控中")
		} else {
			m.monitorItem.SetTitle("▶ 开始监控")
			systray.SetTooltip("小鸟翻译 - 已暂停")
		}
	}
}

// SetOnQuit sets the quit callback
func (m *Manager) SetOnQuit(callback func()) {
	m.onQuit = callback
}

// SetOnSettings sets the settings callback
func (m *Manager) SetOnSettings(callback func()) {
	m.onSettings = callback
}

// SetOnToggleMonitor sets the toggle monitor callback
func (m *Manager) SetOnToggleMonitor(callback func(bool)) {
	m.onToggleMonitor = callback
}

// SetOnRefresh sets the refresh callback
func (m *Manager) SetOnRefresh(callback func()) {
	m.onRefresh = callback
}

// SetOnSelectPrompt sets the select prompt callback
func (m *Manager) SetOnSelectPrompt(callback func(string)) {
	m.onSelectPrompt = callback
}

// SetOnToggleTerminal sets the toggle terminal callback
func (m *Manager) SetOnToggleTerminal(callback func()) {
	m.onToggleTerminal = callback
}

// SetBusinessLogic sets the business logic callback
func (m *Manager) SetBusinessLogic(callback func()) {
	m.businessLogic = callback
}

// UpdatePromptList updates the prompt list in tray menu
func (m *Manager) UpdatePromptList(prompts []struct{ ID, Name string }) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Clear existing submenu items
	for _, item := range m.promptSubItems {
		item.Hide()
	}
	m.promptSubItems = nil

	// Add new submenu items
	for _, prompt := range prompts {
		subItem := m.promptItem.AddSubMenuItem(prompt.Name, "")
		m.promptSubItems = append(m.promptSubItems, subItem)

		// Handle click in a goroutine
		go func(id string) {
			for range subItem.ClickedCh {
				if m.onSelectPrompt != nil {
					m.onSelectPrompt(id)
				}
			}
		}(prompt.ID)
	}
}

// Quit quits the tray
func (m *Manager) Quit() {
	systray.Quit()
}

// IsReady returns whether the tray is ready
func (m *Manager) IsReady() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.isReady
}