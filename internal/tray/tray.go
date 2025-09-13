//go:build windows
// +build windows

package tray

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

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
	mu                sync.RWMutex // Protect shared state
	status            Status
	visible           bool
	isMonitoring      bool
	isReady           bool  // Whether tray is initialized
	translationCount  int
	currentPromptName string
	businessLogic     func() // Business logic to run after tray is ready
	onQuit            func()
	onShow            func()
	onSettings        func()
	onToggleMonitor   func(bool)
	onRefresh         func()
	onSwitchPrompt    func()
	onSelectPrompt    func(string) // Callback for selecting a specific prompt
	onToggleDebugConsole func()
	// Menu items
	mToggle       *systray.MenuItem
	mPromptInfo   *systray.MenuItem
	mStatus       *systray.MenuItem
	mRefresh      *systray.MenuItem
	mPromptMenu   *systray.MenuItem // Main prompt menu
	mDebugConsole *systray.MenuItem // Debug console menu item
	promptItems   []*systray.MenuItem // Individual prompt menu items
}

// NewManager creates a new tray manager
func NewManager() (*Manager, error) {
	return &Manager{
		status:       StatusIdle,
		visible:      true,
		isMonitoring: false,
	}, nil
}

// SetStatus sets the tray icon status (changes color)
func (m *Manager) SetStatus(status Status) {
	m.mu.Lock()
	m.status = status
	isReady := m.isReady
	isMonitoring := m.isMonitoring
	promptName := m.currentPromptName
	translationCount := m.translationCount
	m.mu.Unlock()

	// Only update if tray is ready
	if !isReady {
		return
	}

	// Load appropriate icon based on status
	configDir, _ := os.UserConfigDir()
	var iconPath string

	// 获取当前风格名称
	t := i18n.T()
	if promptName == "" {
		promptName = t.NotSet
	}

	// Determine which color icon to use
	var iconColor string

	// 如果监控已关闭，显示红色图标
	if !isMonitoring && status != StatusError {
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
			systray.SetTooltip(fmt.Sprintf("xiaoniao - %s (%s %d %s) | %s: %s", t.Monitoring, t.TotalCount, translationCount, t.TranslateCount, t.TranslateStyle, promptName))
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
	m.mu.RLock()
	defer m.mu.RUnlock()
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

// SetOnToggleDebugConsole sets the debug console toggle callback
func (m *Manager) SetOnToggleDebugConsole(callback func()) {
	m.onToggleDebugConsole = callback
}

// SetBusinessLogic sets the business logic callback to run after tray initialization
func (m *Manager) SetBusinessLogic(callback func()) {
	m.businessLogic = callback
}

// Initialize initializes the system tray
// On Windows, this must be called from the main thread
func (m *Manager) Initialize() error {
	// Windows: Run directly in main thread (blocking call)
	systray.Run(m.onReady, m.onExit)
	return nil
}


func (m *Manager) onReady() {
	// Mark as ready before any systray operations
	m.mu.Lock()
	m.isReady = true
	isMonitoring := m.isMonitoring
	currentPromptName := m.currentPromptName
	m.mu.Unlock()

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
	// fmt.Printf("🏗️ DEBUG: 创建托盘菜单，TrayToggle文本: '%s'，isMonitoring: %v\n", t.TrayToggle, isMonitoring)
	m.mToggle = systray.AddMenuItemCheckbox(t.TrayToggle, t.TrayToggle, isMonitoring)
	if m.mToggle != nil {
		// fmt.Println("🏗️ DEBUG: mToggle菜单项创建成功")
		if m.mToggle.ClickedCh != nil {
			// fmt.Println("🏗️ DEBUG: mToggle.ClickedCh通道创建成功")
		} else {
			// fmt.Println("❌ DEBUG: 警告：mToggle.ClickedCh通道为nil！")
		}
	} else {
		// fmt.Println("❌ DEBUG: 错误：mToggle菜单项创建失败！")
	}
	
	// 显示当前 prompt
	promptLabel := fmt.Sprintf("%s: %s", t.TranslateStyle, currentPromptName)
	if currentPromptName == "" {
		promptLabel = fmt.Sprintf("%s: %s", t.TranslateStyle, t.NotSet)
	}
	m.mPromptInfo = systray.AddMenuItem(promptLabel, t.TranslateStyle)
	m.mPromptInfo.Disable() // 这个只是显示，不能点击
	
	systray.AddSeparator()
	
	m.mRefresh = systray.AddMenuItem(t.TrayRefresh, t.TrayRefresh)
	mConfig := systray.AddMenuItem(t.TraySettings, t.TraySettings)
	m.mPromptMenu = systray.AddMenuItem(t.TranslateStyle, t.TranslateStyle)
	m.mDebugConsole = systray.AddMenuItem(t.ExportLogs, t.ExportLogs)
	systray.AddSeparator()

	mTutorial := systray.AddMenuItem(t.Tutorial, t.Tutorial)
	mAbout := systray.AddMenuItem(t.TrayAbout, t.TrayAbout)
	mQuit := systray.AddMenuItem(t.TrayQuit, t.TrayQuit)
	
	// Handle menu events
	go func() {
		// fmt.Println("✅ DEBUG: 菜单事件监听goroutine已启动")
		for {
			select {
			case <-m.mToggle.ClickedCh:
				// fmt.Println("🔥 DEBUG: 检测到停止/启动监控菜单点击事件")
				m.toggleMonitor()
			case <-m.mRefresh.ClickedCh:
				// fmt.Println("🔥 DEBUG: 检测到刷新菜单点击事件")
				m.refreshConfig()
			case <-mConfig.ClickedCh:
				// fmt.Println("🔥 DEBUG: 检测到设置菜单点击事件")
				m.openSettings()
			case <-m.mDebugConsole.ClickedCh:
				// fmt.Println("🔥 DEBUG: 检测到调试控制台菜单点击事件")
				m.toggleDebugConsole()
			case <-mTutorial.ClickedCh:
				m.showTutorial()
			case <-mAbout.ClickedCh:
				// fmt.Println("🔥 DEBUG: 检测到关于菜单点击事件")
				m.showAbout()
			case <-mQuit.ClickedCh:
				// fmt.Println("🔥 DEBUG: 检测到退出菜单点击事件")
				m.quit()
				return
			}
		}
	}()
	
	// 确保菜单事件监听启动后，再启动业务逻辑，避免时序问题
	go func() {
		// 给菜单事件监听一点时间启动
		time.Sleep(100 * time.Millisecond)
		// fmt.Println("🚀 DEBUG: 菜单事件监听应该已经启动，现在启动业务逻辑")
		
		// After tray is initialized, run the business logic
		if m.businessLogic != nil {
			// fmt.Println("🚀 DEBUG: 准备启动业务逻辑goroutine")
			go m.businessLogic() // 在新的goroutine中运行，避免阻塞
		} else {
			// fmt.Println("❌ DEBUG: businessLogic为nil，跳过业务逻辑启动")
		}
		
		// 测试机制已移除 - 仅在需要时启用
	}()
}

func (m *Manager) onExit() {
	// Cleanup
}

func (m *Manager) toggleMonitor() {
	// fmt.Printf("🔧 DEBUG: toggleMonitor() 开始执行，当前监控状态: %v\n", m.isMonitoring)

	// 切换监控状态
	m.mu.Lock()
	m.isMonitoring = !m.isMonitoring
	newMonitoringState := m.isMonitoring
	m.mu.Unlock()
	// fmt.Printf("🔧 DEBUG: 监控状态已切换为: %v\n", newMonitoringState)

	if newMonitoringState {
		m.mToggle.Check()
		// fmt.Println("🔧 DEBUG: 菜单项已设为选中状态")
	} else {
		m.mToggle.Uncheck()
		// fmt.Println("🔧 DEBUG: 菜单项已设为未选中状态")
	}

	// 更新图标状态
	m.SetStatus(StatusIdle)
	// fmt.Println("🔧 DEBUG: 图标状态已更新为StatusIdle")

	// 检查回调函数是否存在
	if m.onToggleMonitor != nil {
		// fmt.Printf("🔧 DEBUG: 准备调用onToggleMonitor回调，参数: %v\n", newMonitoringState)
		m.onToggleMonitor(newMonitoringState)
		// fmt.Println("🔧 DEBUG: onToggleMonitor回调执行完毕")
	} else {
		// fmt.Println("❌ DEBUG: onToggleMonitor回调函数为nil！")
	}
}


func (m *Manager) openSettings() {
	if m.onSettings != nil {
		m.onSettings()
	} else {
		// Windows: 获取当前程序路径并在新终端窗口中打开配置界面
		exePath, err := os.Executable()
		if err != nil {
			// 如果获取不到程序路径，使用默认的xiaoniao.exe
			exePath = "xiaoniao.exe"
		} else {
			// 确保Windows下的可执行文件有.exe扩展名
			if filepath.Ext(exePath) == "" {
				exePath = exePath + ".exe"
			}
		}
		
		// 创建命令并启动
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", exePath, "config")
		err = cmd.Start()
		if err != nil {
			// 如果启动失败，尝试使用绝对路径
			if absPath, absErr := filepath.Abs(exePath); absErr == nil {
				cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", absPath, "config")
				cmd.Start()
			}
		}
	}
}

func (m *Manager) showTutorial() {
	// Windows: 获取当前程序路径并直接打开教程页面
	exePath, err := os.Executable()
	if err != nil {
		// 如果获取不到程序路径，使用默认的xiaoniao.exe
		exePath = "xiaoniao.exe"
	} else {
		// 确保Windows下的可执行文件有.exe扩展名
		if filepath.Ext(exePath) == "" {
			exePath = exePath + ".exe"
		}
	}

	// 创建命令并启动
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", exePath, "tutorial")
	err = cmd.Start()
	if err != nil {
		// 如果启动失败，尝试使用绝对路径
		if absPath, absErr := filepath.Abs(exePath); absErr == nil {
			cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", absPath, "tutorial")
			cmd.Start()
		}
	}
}

func (m *Manager) showAbout() {
	// Windows: 获取当前程序路径并直接打开关于页面
	exePath, err := os.Executable()
	if err != nil {
		// 如果获取不到程序路径，使用默认的xiaoniao.exe
		exePath = "xiaoniao.exe"
	} else {
		// 确保Windows下的可执行文件有.exe扩展名
		if filepath.Ext(exePath) == "" {
			exePath = exePath + ".exe"
		}
	}

	// 创建命令并启动
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", exePath, "about")
	err = cmd.Start()
	if err != nil {
		// 如果启动失败，尝试使用绝对路径
		if absPath, absErr := filepath.Abs(exePath); absErr == nil {
			cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", absPath, "about")
			cmd.Start()
		}
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

func (m *Manager) toggleDebugConsole() {
	if m.onToggleDebugConsole != nil {
		m.onToggleDebugConsole()
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
	m.mu.Lock()
	m.translationCount++
	count := m.translationCount
	status := m.status
	m.mu.Unlock()

	if m.mStatus != nil {
		m.mStatus.SetTitle(fmt.Sprintf("状态: 已翻译 %d 次", count))
	}
	m.SetStatus(status) // Update tooltip
}

// UpdateMonitorStatus updates the monitor status in UI
func (m *Manager) UpdateMonitorStatus(running bool) {
	t := i18n.T()
	m.mu.Lock()
	m.isMonitoring = running
	m.mu.Unlock()
	
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
	m.mu.Lock()
	m.currentPromptName = promptName
	isMonitoring := m.isMonitoring
	m.mu.Unlock()
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
		if isMonitoring {
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

// UpdateDebugConsoleMenu 保留函数签名以兼容，但不再更新菜单文字
func (m *Manager) UpdateDebugConsoleMenu(isVisible bool) {
	// 不再需要更新菜单文字，因为现在是"导出日志"功能
}

