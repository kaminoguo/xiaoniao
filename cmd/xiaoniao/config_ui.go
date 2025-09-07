package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
	"github.com/kaminoguo/xiaoniao/internal/translator"
)

// 版本号定义
const APP_VERSION = "v1.4"

var (
	// 修复颜色问题 - 使用高对比度配色
	primaryColor   = lipgloss.Color("#00FFFF")  // 青色文字（默认）
	bgColor        = lipgloss.Color("#1a1a1a")  // 深灰背景
	accentColor    = lipgloss.Color("#00FFFF")  // 青色强调
	mutedColor     = lipgloss.Color("#888888")  // 灰色次要文字
	successColor   = lipgloss.Color("#00FF00")  // 绿色成功
	errorColor     = lipgloss.Color("#FF0000")  // 红色错误
	warningColor   = lipgloss.Color("#FFA500")  // 橙色警告
	selectBgColor  = lipgloss.Color("#333333")  // 选中背景

	// 样式定义
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(accentColor).
			Padding(1, 2).
			MarginBottom(1)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(accentColor).
			Foreground(primaryColor).
			Padding(1).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Background(selectBgColor).
			Bold(true).
			Padding(0, 1)

	normalStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			MarginTop(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			MarginTop(1)
			
	inputStyle = lipgloss.NewStyle().
			Foreground(primaryColor)
			
	previewStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(mutedColor).
			Padding(0, 1).
			MarginTop(1)
			
	successStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)
	
	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)
		
	mutedStyle = lipgloss.NewStyle().
			Foreground(mutedColor)
			
	warningStyle = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)
			
	dimStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true)
)

type screen int

const (
	mainScreen screen = iota
	apiKeyScreen
	promptScreen
	promptEditScreen
	testScreen
	languageScreen
	modelSelectScreen      // 主模型选择界面
	themeScreen           // 主题选择界面
	fallbackModelScreen   // 副模型选择界面
	hotkeyScreen          // 快捷键设置界面
	aboutScreen           // 关于界面
)

type CustomPrompt struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type configModel struct {
	screen           screen
	cursor           int
	apiKeyInput      textinput.Model
	promptNameInput  textinput.Model
	promptContentInput textarea.Model
	prompts          []Prompt
	customPrompts    []CustomPrompt
	selectedPrompt   int
	editingPromptIdx int
	width            int
	height           int
	testResult       string
	testInput        string  // 新增：测试输入的文字
	testing          bool
	quitting         bool
	config           *Config
	confirmDelete    bool
	promptMode       string // "select", "manage"
	promptsModified  bool  // 标记prompts是否被修改
	cachedModels     []string  // 缓存的模型列表
	selectedTheme    int      // 选中的主题索引
	modelsLoaded     bool      // 模型是否已加载
	selectingFallback bool     // 是否正在选择副模型
	recordingHotkey  bool      // 是否正在录制快捷键
	hotkeyBuffer     string    // 快捷键缓冲区
	detectedKeys     []string  // 检测到的按键组合
	modifierKeys     map[string]bool  // 当前按下的修饰键
	changingAPIKey   bool      // 是否正在更改API密钥
}

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Back   key.Binding
	Quit   key.Binding
	Tab    key.Binding
	Edit   key.Binding
	Delete key.Binding
	Add    key.Binding
	Test   key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc", "q"),
		key.WithHelp("esc/q", "return"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Add: key.NewBinding(
		key.WithKeys("n", "a"),
		key.WithHelp("n/a", "new"),
	),
	Test: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "test"),
	),
}

func initialModel() configModel {
	// 加载配置
	loadConfig()
	
	// 设置语言
	if config.Language != "" {
		i18n.SetLanguage(i18n.Language(config.Language))
	}
	
	// 检查是否要显示关于页面
	showAbout := os.Getenv("SHOW_ABOUT") == "1"
	
	// 初始化API输入框
	ti := textinput.New()
	ti.Placeholder = "sk-..."
	ti.CharLimit = 200
	ti.Width = 50
	ti.TextStyle = inputStyle
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(mutedColor)
	if config.APIKey != "" {
		ti.SetValue(config.APIKey)
	}

	// 初始化Prompt输入框
	nameInput := textinput.New()
	nameInput.Placeholder = i18n.T().PromptName
	nameInput.CharLimit = 50
	nameInput.Width = 50
	nameInput.TextStyle = inputStyle
	nameInput.PlaceholderStyle = lipgloss.NewStyle().Foreground(mutedColor)
	
	contentInput := textarea.New()
	contentInput.Placeholder = i18n.T().PromptContent
	contentInput.CharLimit = 2000
	contentInput.SetWidth(70)
	contentInput.SetHeight(12)  // 显示12行
	contentInput.ShowLineNumbers = false

	// 加载所有prompts（包括已修改的）
	prompts := loadAllPrompts()
	customPrompts := loadCustomPrompts()

	// 设置初始屏幕
	initialScreen := mainScreen
	if showAbout {
		initialScreen = aboutScreen
	}

	return configModel{
		screen:             initialScreen,
		apiKeyInput:        ti,
		promptNameInput:    nameInput,
		promptContentInput: contentInput,
		prompts:            prompts,
		customPrompts:      customPrompts,
		selectedPrompt:     getPromptIndex(config.PromptID),
		config:             &config,
		promptMode:         "select",
	}
}

func (m configModel) Init() tea.Cmd {
	return nil
}

func (m configModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Store original message before type assertion
	originalMsg := msg
	
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	
	case tea.KeyMsg:
		// For promptEditScreen, we need to handle it specially
		if m.screen == promptEditScreen {
			// Pass the original message to textarea
			return m.updatePromptEditScreenWithMsg(originalMsg)
		}
		
		switch m.screen {
		case mainScreen:
			return m.updateMainScreen(msg)
		case apiKeyScreen:
			return m.updateAPIKeyScreen(msg)
		case promptScreen:
			return m.updatePromptScreen(msg)
		case testScreen:
			return m.updateTestScreen(msg)
		case languageScreen:
			return m.updateLanguageScreen(msg)
		case modelSelectScreen:
			return m.updateModelSelectScreen(msg)
		case themeScreen:
			return m.updateThemeScreen(msg)
		case hotkeyScreen:
			return m.updateHotkeyScreen(msg)
		case aboutScreen:
			return m.updateAboutScreen(msg)
		}
		
	case string:
		// 处理自定义消息
		if msg == "show_model_selector" {
			// 显示模型选择器
			return m.showModelSelector()
		}
		// 处理清除快捷键结果消息
		if msg == "clear_hotkey_result" {
			m.testResult = ""
			return m, nil
		}
		// 处理测试结果消息
		if strings.Contains(msg, "✅") || strings.Contains(msg, "❌") || strings.Contains(msg, "翻译结果") {
			m.testResult = msg
			m.testing = false
			return m, nil
		}
	}

	return m, nil
}

func (m configModel) updateMainScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Quit):
		m.quitting = true
		saveConfig()
		if m.promptsModified {
			saveAllPrompts(m.prompts)
		}
		return m, tea.Quit

	case key.Matches(msg, keys.Up):
		if m.cursor > 0 {
			m.cursor--
		}

	case key.Matches(msg, keys.Down):
		if m.cursor < 8 { // 9个选项(移除测试连接)
			m.cursor++
		}

	case key.Matches(msg, keys.Enter):
		switch m.cursor {
		case 0: // API配置
			m.screen = apiKeyScreen
			m.cursor = 0  // Reset cursor for API config menu
			m.initAPIConfig()
			m.apiKeyInput.SetValue(m.config.APIKey)
			// 如果已有API key，不要让输入框获得焦点
			if m.config.APIKey == "" {
				m.apiKeyInput.Focus()
				return m, textinput.Blink
			}
			return m, nil
		case 1: // 翻译风格
			m.screen = promptScreen
			m.promptMode = "select"
			m.confirmDelete = false
		case 2: // 界面语言
			m.screen = languageScreen
			// 初始化cursor到当前语言位置
			languages := i18n.GetAvailableLanguages()
			for i, lang := range languages {
				if lang == i18n.GetLanguage() {
					m.cursor = i
					break
				}
			}
		case 3: // 界面主题
			m.screen = themeScreen
			m.cursor = 0
		case 4: // 快捷键设置
			m.screen = hotkeyScreen
			m.cursor = 0
		case 5: // 刷新配置
			// 重新加载配置
			loadConfig()
			// 重新加载 prompts
			m.prompts = loadAllPrompts()
			m.config = &config
			// 创建刷新信号文件通知运行中的守护进程
			homeDir, _ := os.UserHomeDir()
			signalPath := filepath.Join(homeDir, ".config", "xiaoniao", ".refresh_signal")
			os.WriteFile(signalPath, []byte(time.Now().Format(time.RFC3339)), 0644)
			m.testResult = "✅ 配置已刷新，翻译器将重新初始化"
			return m, nil
		case 6: // 测试翻译
			m.screen = testScreen
			m.testInput = ""
			m.testResult = ""
			m.testing = false
			m.promptNameInput.SetValue("")
			m.promptNameInput.Focus()
			return m, textinput.Blink
		case 7: // 关于
			m.screen = aboutScreen
			m.cursor = 0
		case 8: // 保存退出
			m.quitting = true
			saveConfig()
			if m.promptsModified {
				saveAllPrompts(m.prompts)
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m configModel) updateLanguageScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	languages := i18n.GetAvailableLanguages()
	
	switch {
	case key.Matches(msg, keys.Back):
		m.screen = mainScreen
		m.cursor = 2 // 返回主菜单的语言选项
		return m, nil

	case key.Matches(msg, keys.Up):
		if m.cursor > 0 {
			m.cursor--
		}

	case key.Matches(msg, keys.Down):
		if m.cursor < len(languages)-1 {
			m.cursor++
		}

	case key.Matches(msg, keys.Enter):
		// 应用选中的语言
		i18n.SetLanguage(languages[m.cursor])
		config.Language = string(languages[m.cursor])
		saveConfig()
		m.screen = mainScreen
		m.cursor = 2
		return m, nil
	}

	return m, nil
}

func (m configModel) updateModelSelectScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// 动态获取当前provider的模型列表
	var models []string
	if m.promptNameInput.Value() != "" {
		// 搜索模型
		allModels := m.getAvailableModels()
		searchTerm := strings.ToLower(m.promptNameInput.Value())
		for _, model := range allModels {
			if strings.Contains(strings.ToLower(model), searchTerm) {
				models = append(models, model)
			}
		}
	} else {
		// 获取所有模型
		models = m.getAvailableModels()
	}
	
	totalModels := len(models)
	
	switch msg.String() {
	case "esc":
		m.screen = apiKeyScreen
		return m, nil
		
	case "enter":
		// 选择模型
		if totalModels > 0 && m.selectedPrompt < totalModels {
			if m.selectingFallback {
				// 选择副模型
				m.config.FallbackModel = models[m.selectedPrompt]
				config = *m.config
				saveConfig()
				m.screen = apiKeyScreen
				m.testResult = fmt.Sprintf("✅ 副模型已设置为: %s", m.config.FallbackModel)
			} else {
				// 选择主模型
				m.config.Model = models[m.selectedPrompt]
				config = *m.config
				saveConfig()
				m.screen = apiKeyScreen
				m.testResult = fmt.Sprintf("✅ 主模型已更改为: %s", m.config.Model)
			}
		}
		return m, nil
	
	case "t":
		// 测试当前选中的模型
		if totalModels > 0 && m.selectedPrompt < totalModels {
			selectedModel := models[m.selectedPrompt]
			m.testing = true
			m.testResult = fmt.Sprintf("🔄 " + i18n.T().TestingModel, selectedModel)
			
			// 创建测试命令
			return m, func() tea.Msg {
				// 临时设置模型进行测试
				testConfig := Config{
					APIKey:   m.config.APIKey,
					Provider: m.config.Provider,
					Model:    selectedModel,
					PromptID: "direct",
				}
				
				// 测试翻译
				transConfig := &translator.Config{
					Provider: testConfig.Provider,
					APIKey:   testConfig.APIKey,
					Model:    testConfig.Model,
				}
				trans, err := translator.NewTranslator(transConfig)
				if err != nil {
					return fmt.Sprintf("❌ 模型 %s 初始化失败: %v", selectedModel, err)
				}
				result, err := trans.Translate("Hello world", "请仅翻译以下内容成中文，不要回答或解释，只输出译文：")
				
				if err != nil {
					return fmt.Sprintf("❌ " + i18n.T().ModelTestFailed, selectedModel, err)
				}
				
				if result.Success && result.Translation != "" {
					return fmt.Sprintf("✅ 模型 %s 可用！译文: %s", selectedModel, result.Translation)
				}
				
				return fmt.Sprintf("❌ 模型 %s 无响应", selectedModel)
			}
		}
		return m, nil
		
	case "up", "k":
		if m.selectedPrompt > 0 {
			m.selectedPrompt--
		} else if totalModels > 0 {
			m.selectedPrompt = totalModels - 1 // 循环到底部
		}
		
	case "down", "j":
		if m.selectedPrompt < totalModels-1 {
			m.selectedPrompt++
		} else {
			m.selectedPrompt = 0 // 循环到顶部
		}
		
	case "/":
		// 开始搜索
		m.promptNameInput.SetValue("")
		m.promptNameInput.Focus()
		return m, textinput.Blink
		
	default:
		// 处理搜索输入
		if m.promptNameInput.Focused() {
			var cmd tea.Cmd
			m.promptNameInput, cmd = m.promptNameInput.Update(msg)
			// 重置选择索引
			m.selectedPrompt = 0
			return m, cmd
		}
	}
	
	return m, nil
}

func (m configModel) updatePromptScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back):
		m.screen = mainScreen
		m.confirmDelete = false
		return m, nil

	case key.Matches(msg, keys.Add):
		m.screen = promptEditScreen
		m.editingPromptIdx = -1 // 新建
		m.promptNameInput.SetValue("")
		m.promptContentInput.SetValue("")
		m.promptNameInput.Focus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Up):
		if m.selectedPrompt > 0 {
			m.selectedPrompt--
			m.confirmDelete = false
		} else {
			// 在顶部时循环到底部
			m.selectedPrompt = len(m.prompts) - 1
			m.confirmDelete = false
		}

	case key.Matches(msg, keys.Down):
		if m.selectedPrompt < len(m.prompts)-1 {
			m.selectedPrompt++
			m.confirmDelete = false
		} else {
			// 在底部时循环到顶部
			m.selectedPrompt = 0
			m.confirmDelete = false
		}

	case key.Matches(msg, keys.Edit):
		// 可以编辑任何prompt
		currentPrompt := m.prompts[m.selectedPrompt]
		m.screen = promptEditScreen
		m.editingPromptIdx = m.selectedPrompt
		m.promptNameInput.SetValue(strings.TrimSuffix(currentPrompt.Name, " (自定义)"))
		m.promptContentInput.SetValue(currentPrompt.Content)
		m.promptNameInput.Focus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Delete):
		// 可以删除任何prompt
		if m.confirmDelete {
			// 删除选中的prompt - 立即保存到文件
			if m.selectedPrompt < len(m.prompts) {
				promptToDelete := m.prompts[m.selectedPrompt]
				err := DeletePrompt(promptToDelete.ID)
				if err != nil {
					m.testResult = fmt.Sprintf("删除失败: %v", err)
				} else {
					// 重新加载prompts以确保同步
					m.prompts = loadAllPrompts()
					if m.selectedPrompt >= len(m.prompts) && m.selectedPrompt > 0 {
						m.selectedPrompt--
					}
				}
			}
			m.confirmDelete = false
		} else {
			m.confirmDelete = true
		}

	case key.Matches(msg, keys.Enter):
		if !m.confirmDelete {
			m.config.PromptID = m.prompts[m.selectedPrompt].ID
			// 立即保存配置，避免刷新时丢失
			config = *m.config
			saveConfig()
			m.screen = mainScreen
			return m, nil
		} else {
			m.confirmDelete = false
		}
	}

	return m, nil
}

func (m configModel) updatePromptEditScreenWithMsg(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle key messages
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		// Handle special keys first
		switch {
		case key.Matches(keyMsg, keys.Back):
			m.screen = promptScreen
			return m, nil

		case key.Matches(keyMsg, keys.Tab):
			if m.promptNameInput.Focused() {
				m.promptNameInput.Blur()
				m.promptContentInput.Focus()
				return m, nil
			} else {
				m.promptContentInput.Blur()
				m.promptNameInput.Focus()
				return m, textinput.Blink
			}

		case keyMsg.String() == "ctrl+s":
			// Save with Ctrl+S
			name := m.promptNameInput.Value()
			content := m.promptContentInput.Value()
			
			if name != "" && content != "" {
				if m.editingPromptIdx == -1 {
					// 新建 - 立即保存到文件
					// 找到下一个可用的ID
					maxID := -1
					for _, p := range m.prompts {
						if strings.HasPrefix(p.ID, "custom_") {
							idStr := strings.TrimPrefix(p.ID, "custom_")
							if id, err := strconv.Atoi(idStr); err == nil && id > maxID {
								maxID = id
							}
						}
					}
					id := fmt.Sprintf("custom_%d", maxID+1)
					err := AddPrompt(id, name, content)
					if err != nil {
						m.testResult = fmt.Sprintf("保存失败: %v", err)
					} else {
						// 重新加载prompts以确保同步
						m.prompts = loadAllPrompts()
					}
				} else if m.editingPromptIdx < len(m.prompts) {
					// 编辑现有prompt - 立即保存到文件
					prompt := m.prompts[m.editingPromptIdx]
					err := UpdatePrompt(prompt.ID, name, content)
					if err != nil {
						m.testResult = fmt.Sprintf("更新失败: %v", err)
					} else {
						// 重新加载prompts以确保同步
						m.prompts = loadAllPrompts()
					}
				}
				
				m.screen = promptScreen
			}
			return m, nil
		}
	}
	
	// Update the focused input with the full message
	if m.promptNameInput.Focused() {
		var cmd tea.Cmd
		m.promptNameInput, cmd = m.promptNameInput.Update(msg)
		return m, cmd
	} else if m.promptContentInput.Focused() {
		var cmd tea.Cmd
		m.promptContentInput, cmd = m.promptContentInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

// Keep the old function for compatibility
func (m configModel) updatePromptEditScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m.updatePromptEditScreenWithMsg(msg)
}

func (m configModel) updateAPIKeyScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// 添加调试信息到testResult中
	keyPressed := msg.String()
	
	// 直接在这里处理API配置逻辑
	if m.config.APIKey == "" || m.changingAPIKey {
		// 没有API Key或正在更改，显示输入界面
		switch msg.String() {
		case "enter":
			apiKey := m.apiKeyInput.Value()
			if apiKey != "" {
				m.config.APIKey = apiKey
				m.testing = true
				m.changingAPIKey = false  // 重置标志
				return m, m.detectAndTestAPI(apiKey)
			}
		case "esc":
			if m.changingAPIKey {
				m.changingAPIKey = false
				// API密钥保持不变
			} else {
				m.screen = mainScreen
			}
			return m, nil
		default:
			var cmd tea.Cmd
			m.apiKeyInput, cmd = m.apiKeyInput.Update(msg)
			return m, cmd
		}
	} else {
		// 已有API Key，显示配置菜单
		// 确保输入框失焦
		if m.apiKeyInput.Focused() {
			m.apiKeyInput.Blur()
		}
		
		// 显示按键调试信息
		if keyPressed != "up" && keyPressed != "down" && keyPressed != "k" && keyPressed != "j" {
			m.testResult = fmt.Sprintf("%s: [%s], %s: %d, %s: %v", i18n.T().KeyPressed, keyPressed, i18n.T().CursorPosition, m.cursor, i18n.T().InputFocus, m.apiKeyInput.Focused())
		}
		
		switch msg.String() {
		case "enter":
			switch m.cursor {
			case 0:
				// 测试连接
				m.testing = true
				m.testResult = i18n.T().TestingConnection + "..."
				return m, func() tea.Msg {
					success, result, _ := testAPIConnectionStandalone(m.config.APIKey, m.config.Provider)
					if success {
						return fmt.Sprintf("✅ %s", result)
					}
					return fmt.Sprintf("❌ %s", result)
				}
			case 1:
				// 选择主模型
				return m.showModelSelector()
			case 2:
				// 选择副模型
				return m.showFallbackModelSelector()
			case 3:
				// 更改API密钥
				m.changingAPIKey = true
				m.apiKeyInput.SetValue(m.config.APIKey)
				m.apiKeyInput.Focus()
				return m, nil
			}
			
		case "1":
			// 测试连接
			m.cursor = 0
			m.testing = true
			m.testResult = "正在测试..."
			return m, func() tea.Msg {
				success, result, _ := testAPIConnectionStandalone(m.config.APIKey, m.config.Provider)
				if success {
					return fmt.Sprintf("✅ %s", result)
				}
				return fmt.Sprintf("❌ %s", result)
			}
			
		case "2":
			// 选择主模型
			m.cursor = 1
			return m.showModelSelector()
			
		case "3":
			// 选择副模型
			m.cursor = 2
			return m.showFallbackModelSelector()
			
		case "4":
			// 更改API密钥
			m.cursor = 3
			m.changingAPIKey = true
			m.apiKeyInput.SetValue(m.config.APIKey)
			m.apiKeyInput.Focus()
			return m, nil
			
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			
		case "down", "j":
			if m.cursor < 3 {  // 现在有4个选项
				m.cursor++
			}
			
		case "esc":
			m.screen = mainScreen
			m.cursor = 0
			return m, nil
		}
	}
	
	return m, nil
}

func (m configModel) updateTestScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back):
		m.screen = mainScreen
		m.testResult = ""
		m.testing = false
		m.promptNameInput.Blur()
		return m, nil
		
	case key.Matches(msg, keys.Enter):
		// 获取输入的文字
		testText := m.promptNameInput.Value()
		if testText != "" && !m.testing {
			m.testing = true
			m.testInput = testText
			// 执行翻译测试
			return m, func() tea.Msg {
				// 加载当前配置
				loadConfig()
				
				// 创建translator
				translatorConfig := &translator.Config{
					APIKey:       config.APIKey,
					Provider:     config.Provider,
					Model:        config.Model,
					MaxRetries:   1,
					Timeout:      30,
				}
				
				trans, err := translator.NewTranslator(translatorConfig)
				if err != nil {
					return fmt.Sprintf("❌ 创建翻译器失败: %v", err)
				}
				
				// 获取当前prompt内容
				promptContent := getPromptContent(config.PromptID)
				
				// 执行翻译
				result, err := trans.Translate(testText, promptContent)
				if err != nil {
					return fmt.Sprintf("❌ 翻译失败: %v", err)
				}
				
				// 返回结果
				return fmt.Sprintf("✅ 翻译结果:\n原文: %s\n译文: %s\n模型: %s\nPrompt: %s",
					testText, result.Translation, config.Model, getPromptName(config.PromptID))
			}
		}
		return m, nil
		
	default:
		// 处理输入
		if !m.testing {
			var cmd tea.Cmd
			m.promptNameInput, cmd = m.promptNameInput.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m configModel) View() string {
	switch m.screen {
	case apiKeyScreen:
		return m.viewAPIKeyScreen()
	case promptScreen:
		return m.viewPromptScreen()
	case promptEditScreen:
		return m.viewPromptEditScreen()
	case testScreen:
		return m.viewTestScreen()
	case languageScreen:
		return m.viewLanguageScreen()
	case modelSelectScreen:
		return m.viewModelSelectScreen()
	case themeScreen:
		return m.viewThemeScreen()
	case hotkeyScreen:
		return m.viewHotkeyScreen()
	case aboutScreen:
		return m.viewAboutScreen()
	default:
		return m.viewMainScreen()
	}
}

func (m configModel) viewMainScreen() string {
	t := i18n.T()
	// 主菜单标题
	s := titleStyle.Render(t.Title)
	s += "\n\n"

	// 菜单选项
	options := []struct {
		name  string
		value string
	}{
		{t.APIConfig, m.getAPIStatus()},
		{t.TranslateStyle, m.getPromptName(m.config.PromptID)},
		{t.Language, i18n.GetLanguageName(i18n.GetLanguage())},
		{t.Theme, m.getThemeName()},
		{t.Hotkeys, t.Hotkeys},
		{"[R] " + t.TrayRefresh, ""},
		{"[T] " + t.TestConnection, ""},
		{t.TrayAbout, ""},
		{t.SaveAndExit, ""},
	}

	for i, opt := range options {
		cursor := "  "
		style := normalStyle
		if i == m.cursor {
			cursor = "▶ "
			style = selectedStyle
		}
		
		line := cursor + opt.name
		if opt.value != "" {
			line += ": " + opt.value
		}
		s += style.Render(line) + "\n"
	}

	// 状态信息
	s += "\n" + statusStyle.Render(fmt.Sprintf("%s: %s | %s: %s", 
		t.Provider, m.config.Provider, 
		t.Model, m.config.Model))
	
	// 帮助信息
	s += "\n" + helpStyle.Render(fmt.Sprintf("%s | %s | %s", 
		t.HelpMove, t.HelpSelect, t.HelpQuit))

	return boxStyle.Render(s)
}

func (m configModel) viewLanguageScreen() string {
	t := i18n.T()
	s := titleStyle.Render(t.Language)
	s += "\n\n"

	languages := i18n.GetAvailableLanguages()

	for i, lang := range languages {
		cursor := "  "
		style := normalStyle
		indicator := " "
		
		// 光标位置
		if i == m.cursor {
			cursor = "▶ "
			style = selectedStyle
		}
		
		// 当前选中的语言
		if lang == i18n.GetLanguage() {
			indicator = "●"
		}
		
		s += style.Render(fmt.Sprintf("%s%s %s", cursor, indicator, i18n.GetLanguageName(lang))) + "\n"
	}

	s += "\n" + helpStyle.Render(fmt.Sprintf("%s | %s | %s", 
		t.HelpMove, t.HelpSelect, t.HelpBack))
	
	return boxStyle.Render(s)
}

func (m configModel) viewPromptScreen() string {
	t := i18n.T()
	// 添加版本号到prompt界面
	s := titleStyle.Render(t.TranslateStyle)
	s += "\n"

	// 左侧：Prompt列表
	listWidth := 40
	previewWidth := 50
	
	const HEIGHT = 12
	total := len(m.prompts)
	
	// 固定高度的列表内容
	var lines [HEIGHT]string
	
	if total == 0 {
		lines[0] = normalStyle.Render("  " + t.NoPromptAvailable)
		for i := 1; i < HEIGHT; i++ {
			lines[i] = " "
		}
	} else {
		// 计算视窗起始索引
		viewStart := 0
		
		if total > HEIGHT {
			// 滚动逻辑：保持选中项可见
			if m.selectedPrompt < HEIGHT/2 {
				viewStart = 0
			} else if m.selectedPrompt > total - HEIGHT/2 - 1 {
				viewStart = total - HEIGHT
			} else {
				viewStart = m.selectedPrompt - HEIGHT/2
			}
			
			// 边界检查
			if viewStart < 0 {
				viewStart = 0
			}
			if viewStart > total - HEIGHT {
				viewStart = total - HEIGHT
			}
		}
		
		// 填充固定数组
		for row := 0; row < HEIGHT; row++ {
			itemIndex := viewStart + row
			
			if itemIndex >= 0 && itemIndex < total {
				promptItem := m.prompts[itemIndex]
				displayName := promptItem.Name
				
				// 截断过长名称
				if len(displayName) > listWidth-4 {
					displayName = displayName[:listWidth-7] + "..."
				}
				
				// 构建行内容
				if itemIndex == m.selectedPrompt {
					lines[row] = selectedStyle.Render("▶ " + displayName)
				} else {
					linePrefix := "  "
					if total > HEIGHT {
						if row == 0 && viewStart > 0 {
							linePrefix = "↑ "
						} else if row == HEIGHT-1 && viewStart+HEIGHT < total {
							linePrefix = "↓ "
						}
					}
					lines[row] = normalStyle.Render(linePrefix + displayName)
				}
			} else {
				lines[row] = " "
			}
		}
	}
	
	// 组合成固定高度的字符串
	listContent := lines[0]
	for i := 1; i < HEIGHT; i++ {
		listContent += "\n" + lines[i]
	}

	// 右侧：Prompt预览
	previewContent := ""
	previewTitle := "预览:"
	if m.selectedPrompt < len(m.prompts) {
		prompt := m.prompts[m.selectedPrompt]
		content := prompt.Content
		// 自动换行
		lines := wrapText(content, previewWidth-4)
		for _, line := range lines {
			previewContent += line + "\n"
		}
	}
	
	// 如果是确认删除状态
	if m.confirmDelete {
		previewContent = lipgloss.NewStyle().
			Foreground(errorColor).
			Render(t.ConfirmDelete + "\n\n" + t.ConfirmDeleteKey + "\n" + t.CancelDelete)
	}

	// 拼接左右两栏 - 确保固定高度
	leftBox := lipgloss.NewStyle().
		Width(listWidth).
		Height(HEIGHT).
		MaxHeight(HEIGHT).
		Render(listContent)
		
	rightBox := previewStyle.
		Width(previewWidth).
		Height(HEIGHT).
		MaxHeight(HEIGHT).
		Render(previewTitle + "\n" + previewContent)

	s += lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox)

	// 帮助信息
	helpText := fmt.Sprintf("%s | %s | %s | %s | %s | %s", t.HelpMove, t.HelpSelect, t.HelpNewPrompt, t.HelpEditPrompt, t.HelpDeletePrompt, t.HelpBack)
	s += "\n" + helpStyle.Render(helpText)
	
	return boxStyle.Render(s)
}

func (m configModel) viewPromptEditScreen() string {
	t := i18n.T()
	title := t.AddPrompt
	if m.editingPromptIdx >= 0 {
		title = t.EditPrompt
	}
	
	s := titleStyle.Render("✏️ " + title)
	s += "\n\n"
	
	s += t.PromptName + ":\n"
	s += m.promptNameInput.View() + "\n\n"
	
	s += t.PromptContent + ":\n"
	s += m.promptContentInput.View() + "\n\n"
	
	s += helpStyle.Render(fmt.Sprintf("%s | Enter 保存 | %s", 
		t.HelpTab, t.HelpBack))
	
	return boxStyle.Render(s)
}

func (m configModel) viewAPIKeyScreen() string {
	return m.viewAPIConfigScreen()
}

func (m configModel) viewTestScreen() string {
	t := i18n.T()
	s := titleStyle.Render(t.TestTranslation)
	s += "\n\n"
	
	// 显示当前配置
	s += fmt.Sprintf("%s:\n", t.CurrentConfig)
	s += fmt.Sprintf("  %s: %s\n", t.Provider, m.config.Provider)
	s += fmt.Sprintf("  %s: %s\n", t.Model, m.config.Model)
	s += fmt.Sprintf("  Prompt: %s\n\n", getPromptName(m.config.PromptID))
	
	// 输入框
	s += t.EnterTextToTranslate + ":\n"
	s += inputStyle.Render(m.promptNameInput.View()) + "\n\n"
	
	// 显示测试结果
	if m.testing {
		s += t.Translating + "...\n"
	} else if m.testResult != "" {
		if strings.Contains(m.testResult, "✅") {
			s += lipgloss.NewStyle().Foreground(successColor).Render(m.testResult) + "\n"
		} else if strings.Contains(m.testResult, "❌") {
			s += lipgloss.NewStyle().Foreground(errorColor).Render(m.testResult) + "\n"
		} else {
			s += m.testResult + "\n"
		}
	}
	
	s += "\n" + helpStyle.Render(fmt.Sprintf("%s | Esc: %s", t.HelpTranslate, t.HelpBack))
	
	return boxStyle.Render(s)
}

// 辅助方法
func wrapText(text string, width int) []string {
	var lines []string
	words := strings.Fields(text)
	currentLine := ""
	
	for _, word := range words {
		if len(currentLine)+len(word)+1 > width {
			if currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = word
			} else {
				// 单词太长，强制分割
				for len(word) > width {
					lines = append(lines, word[:width])
					word = word[width:]
				}
				currentLine = word
			}
		} else {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		}
	}
	
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	
	return lines
}

func (m *configModel) detectProvider(apiKey string) {
	if strings.HasPrefix(apiKey, "sk-ant-") {
		m.config.Provider = "Anthropic"
		m.config.Model = "claude-3-haiku-20240307"
	} else if strings.HasPrefix(apiKey, "sk-") && strings.Contains(apiKey, "-") && len(apiKey) > 50 {
		// OpenAI的key通常较长且包含多个-
		m.config.Provider = "OpenAI"
		m.config.Model = "gpt-4o-mini"
	} else if len(apiKey) == 32 {
		// DeepSeek的API密钥通常是32位
		m.config.Provider = "DeepSeek"
		m.config.Model = "deepseek-chat"
	} else if strings.HasPrefix(apiKey, "sk-") && len(apiKey) > 40 {
		// Moonshot的密钥较长
		m.config.Provider = "Moonshot"
		m.config.Model = "moonshot-v1-8k"
	} else {
		// 默认OpenAI
		m.config.Provider = "OpenAI"
		m.config.Model = "gpt-4o-mini"
	}
}

func (m *configModel) maskAPIKey(key string) string {
	t := i18n.T()
	if key == "" {
		return t.NotSet
	}
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "..." + key[len(key)-4:]
}

// getAvailableModels 动态获取可用模型列表
func (m *configModel) getAvailableModels() []string {
	// 如果已经加载了模型，直接返回缓存
	if m.modelsLoaded && len(m.cachedModels) > 0 {
		return m.cachedModels
	}
	
	// 创建provider实例
	var p translator.Provider
	switch m.config.Provider {
	case "OpenRouter":
		p = translator.NewOpenRouterProvider(m.config.APIKey, "")
	case "Groq":
		p = translator.NewGroqProvider(m.config.APIKey, "")
	case "Together", "TogetherAI":
		p = translator.NewTogetherProvider(m.config.APIKey, "")
	case "OpenAI":
		p = translator.NewOpenAIProvider(m.config.APIKey, "")
	case "Anthropic":
		p = translator.NewAnthropicProvider(m.config.APIKey, "")
	case "DeepSeek":
		p = translator.NewDeepSeekProvider(m.config.APIKey, "")
	case "Moonshot":
		p = translator.NewMoonshotProvider(m.config.APIKey, "")
	default:
		// 对于其他provider，尝试使用OpenAI兼容接口
		p = translator.NewOpenAICompatibleProvider(m.config.Provider, m.config.APIKey, "", "")
	}
	
	// 尝试获取模型列表
	models, err := p.ListModels()
	if err != nil {
		// 如果失败，返回硬编码的列表作为备用
		if fallback, exists := translator.ProviderModels[m.config.Provider]; exists {
			return fallback
		}
		return []string{m.config.Model} // 至少返回当前模型
	}
	
	// 缓存结果
	m.cachedModels = models
	m.modelsLoaded = true
	
	return models
}

func (m *configModel) getPromptName(id string) string {
	for _, p := range m.prompts {
		if p.ID == id {
			return p.Name
		}
	}
	return i18n.T().NotSet
}

func (m configModel) getAPIStatus() string {
	if m.config.APIKey == "" {
		return i18n.T().NotConfiguredBrackets
	}
	provider := m.config.Provider
	if provider == "" {
		provider = translator.DetectProviderByKey(m.config.APIKey)
	}
	if provider == "" {
		provider = i18n.T().UnknownProvider
	}
	model := m.config.Model
	if model == "" {
		model = i18n.T().DefaultName
	}
	return fmt.Sprintf("%s / %s", provider, model)
}

func (m *configModel) rebuildPrompts() {
	// 重新构建prompts列表
	m.prompts = GetAllPrompts()
	for _, cp := range m.customPrompts {
		m.prompts = append(m.prompts, Prompt{
			ID:      cp.ID,
			Name:    cp.Name + i18n.T().CustomSuffix,
			Content: cp.Content,
		})
	}
}

func (m *configModel) testConnection() tea.Cmd {
	return func() tea.Msg {
		t := i18n.T()
		m.testing = true
		
		// 测试连接
		cfg := &translator.Config{
			APIKey:   m.config.APIKey,
			Provider: m.config.Provider,
			Model:    m.config.Model,
		}
		
		trans, err := translator.NewTranslator(cfg)
		if err != nil {
			m.testResult = fmt.Sprintf("%s: %v", t.TestFailed, err)
			m.testing = false
			return nil
		}
		
		// 简单测试
		result, err := trans.Translate("Hello", "翻译成中文")
		if err != nil {
			m.testResult = fmt.Sprintf("%s: %v", t.TestFailed, err)
		} else if result.Success {
			m.testResult = fmt.Sprintf("%s Provider: %s, Model: %s", 
				t.TestSuccess, m.config.Provider, m.config.Model)
		} else {
			m.testResult = t.TestFailed
		}
		
		m.testing = false
		return nil
	}
}

func getPromptIndex(id string) int {
	prompts := GetAllPrompts()
	customPrompts := loadCustomPrompts()

	// 合并prompts
	for _, cp := range customPrompts {
		prompts = append(prompts, Prompt{
			ID:      cp.ID,
			Name:    cp.Name + i18n.T().CustomSuffix,
			Content: cp.Content,
		})
	}
	
	for i, p := range prompts {
		if p.ID == id {
			return i
		}
	}
	return 0
}

// Prompt管理 - 保存所有prompts（包括修改过的内置prompt）
func loadAllPrompts() []Prompt {
	// 直接从新的prompt系统加载
	return GetAllPrompts()
}

func saveAllPrompts(prompts []Prompt) {
	// 不再保存到all_prompts.json，因为新系统会自动保存到prompts.json
	// 这个函数保留是为了兼容性，但实际不做任何操作
	// 真正的保存通过 AddPrompt/UpdatePrompt/DeletePrompt 完成
}

// 自定义Prompt管理（保留兼容性）
func loadCustomPrompts() []CustomPrompt {
	// 新系统不再区分内置和自定义prompt，全部统一管理
	// 返回空列表，让所有prompts都从统一的系统加载
	return []CustomPrompt{}
}

func saveCustomPrompts(prompts []CustomPrompt) {
	// 新系统不再使用custom_prompts.json
	// 这个函数保留是为了兼容性，但实际不做任何操作
}

// 主题相关方法
func (m configModel) getThemeName() string {
	if m.config.Theme == "" {
		return i18n.T().DefaultName
	}
	themes := map[string]string{
		"default":          i18n.T().DefaultName,
		"tokyo-night":      "Tokyo Night",
		"catppuccin-mocha": "Catppuccin Mocha",
		"catppuccin-latte": "Catppuccin Latte",
		"dracula":          "Dracula",
		"gruvbox-dark":     "Gruvbox Dark",
		"gruvbox-light":    "Gruvbox Light",
		"nord":             "Nord",
		"solarized-dark":   "Solarized Dark",
		"solarized-light":  "Solarized Light",
		"minimal":          i18n.T().MinimalTheme,
	}
	if name, ok := themes[m.config.Theme]; ok {
		return name
	}
	return i18n.T().DefaultName
}

func (m configModel) viewThemeScreen() string {
	t := i18n.T()
	s := titleStyle.Render(t.SelectTheme)
	s += "\n\n"
	
	themes := []struct {
		id   string
		name string
		desc string
	}{
		{"default", t.DefaultTheme, t.ClassicBlue},
		{"tokyo-night", "Tokyo Night", t.DarkThemeTokyoNight},
		{"catppuccin-mocha", "Catppuccin Mocha", t.ChocolateTheme},
		{"catppuccin-latte", "Catppuccin Latte", t.LatteTheme},
		{"dracula", "Dracula", t.DraculaTheme},
		{"gruvbox-dark", "Gruvbox Dark", t.GruvboxDarkTheme},
		{"gruvbox-light", "Gruvbox Light", t.GruvboxLightTheme},
		{"nord", "Nord", t.NordTheme},
		{"solarized-dark", "Solarized Dark", t.SolarizedDarkTheme},
		{"solarized-light", "Solarized Light", t.SolarizedLightTheme},
		{"minimal", t.MinimalTheme, t.MinimalBWTheme},
	}
	
	for i, theme := range themes {
		cursor := "  "
		style := normalStyle
		if i == m.cursor {
			cursor = "▶ "
			style = selectedStyle
		}
		
		line := cursor + theme.name
		if theme.desc != "" {
			line += " - " + theme.desc
		}
		if m.config.Theme == theme.id || (m.config.Theme == "" && theme.id == "default") {
			line += " ✓"
		}
		s += style.Render(line) + "\n"
	}
	
	s += "\n" + helpStyle.Render(fmt.Sprintf("%s | %s | %s", t.HelpMove, t.HelpSelect, t.HelpBack))
	return s
}

func (m configModel) updateThemeScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	themes := []string{
		"default", 
		"tokyo-night", 
		"catppuccin-mocha",
		"catppuccin-latte",
		"dracula",
		"gruvbox-dark",
		"gruvbox-light",
		"nord",
		"solarized-dark",
		"solarized-light",
		"minimal",
	}
	
	switch {
	case key.Matches(msg, keys.Back):
		m.screen = mainScreen
		m.cursor = 3 // 回到主题选项
		
	case key.Matches(msg, keys.Up):
		if m.cursor > 0 {
			m.cursor--
		}
		
	case key.Matches(msg, keys.Down):
		if m.cursor < len(themes)-1 {
			m.cursor++
		}
		
	case key.Matches(msg, keys.Enter):
		m.config.Theme = themes[m.cursor]
		// 应用主题
		applyTheme(m.config.Theme)
		m.screen = mainScreen
		m.cursor = 3
	}
	
	return m, nil
}

// 应用主题
func applyTheme(themeName string) {
	// 这里可以根据主题更新全局样式变量
	// 由于lipgloss样式是不可变的，我们需要重新创建样式
	switch themeName {
	case "tokyo-night":
		primaryColor = lipgloss.Color("#7aa2f7")
		accentColor = lipgloss.Color("#7aa2f7")
		mutedColor = lipgloss.Color("#565f89")
		successColor = lipgloss.Color("#9ece6a")
		errorColor = lipgloss.Color("#f7768e")
		warningColor = lipgloss.Color("#e0af68")
		
	case "catppuccin-mocha":
		primaryColor = lipgloss.Color("#cdd6f4")
		accentColor = lipgloss.Color("#89b4fa")
		mutedColor = lipgloss.Color("#45475a")
		successColor = lipgloss.Color("#a6e3a1")
		errorColor = lipgloss.Color("#f38ba8")
		
	case "catppuccin-latte":
		primaryColor = lipgloss.Color("#4c4f69")
		accentColor = lipgloss.Color("#1e66f5")
		mutedColor = lipgloss.Color("#9ca0b0")
		successColor = lipgloss.Color("#40a02b")
		errorColor = lipgloss.Color("#d20f39")
		
	case "dracula":
		primaryColor = lipgloss.Color("#f8f8f2")
		accentColor = lipgloss.Color("#bd93f9")
		mutedColor = lipgloss.Color("#6272a4")
		successColor = lipgloss.Color("#50fa7b")
		errorColor = lipgloss.Color("#ff5555")
		
	case "gruvbox-dark":
		primaryColor = lipgloss.Color("#ebdbb2")
		accentColor = lipgloss.Color("#fabd2f")
		mutedColor = lipgloss.Color("#928374")
		successColor = lipgloss.Color("#b8bb26")
		errorColor = lipgloss.Color("#fb4934")
		
	case "gruvbox-light":
		primaryColor = lipgloss.Color("#3c3836")
		accentColor = lipgloss.Color("#d79921")
		mutedColor = lipgloss.Color("#7c6f64")
		successColor = lipgloss.Color("#98971a")
		errorColor = lipgloss.Color("#cc241d")
		
	case "nord":
		primaryColor = lipgloss.Color("#d8dee9")
		accentColor = lipgloss.Color("#88c0d0")
		mutedColor = lipgloss.Color("#4c566a")
		successColor = lipgloss.Color("#a3be8c")
		errorColor = lipgloss.Color("#bf616a")
		
	case "solarized-dark":
		primaryColor = lipgloss.Color("#839496")
		accentColor = lipgloss.Color("#268bd2")
		mutedColor = lipgloss.Color("#586e75")
		successColor = lipgloss.Color("#859900")
		errorColor = lipgloss.Color("#dc322f")
		
	case "solarized-light":
		primaryColor = lipgloss.Color("#657b83")
		accentColor = lipgloss.Color("#268bd2")
		mutedColor = lipgloss.Color("#93a1a1")
		successColor = lipgloss.Color("#859900")
		errorColor = lipgloss.Color("#dc322f")
		
	case "minimal":
		primaryColor = lipgloss.Color("#ffffff")
		accentColor = lipgloss.Color("#ffffff")
		mutedColor = lipgloss.Color("#888888")
		successColor = lipgloss.Color("#ffffff")
		errorColor = lipgloss.Color("#ffffff")
		
	default: // default theme
		primaryColor = lipgloss.Color("#00FFFF")
		accentColor = lipgloss.Color("#00FFFF")
		mutedColor = lipgloss.Color("#888888")
		successColor = lipgloss.Color("#00FF00")
		errorColor = lipgloss.Color("#FF0000")
	}
	
	// 重新创建样式
	updateStyles()
}

// 更新样式
func updateStyles() {
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(accentColor).
		Padding(1, 2).
		MarginBottom(1)
		
	boxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(accentColor).
		Foreground(primaryColor).
		Padding(1).
		MarginBottom(1)
		
	selectedStyle = lipgloss.NewStyle().
		Foreground(accentColor).
		Background(selectBgColor).
		Bold(true).
		Padding(0, 1)
		
	normalStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Padding(0, 1)
		
	statusStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		MarginTop(1)
		
	helpStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		MarginTop(1)
		
	inputStyle = lipgloss.NewStyle().
		Foreground(primaryColor)
		
	previewStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		Border(lipgloss.NormalBorder()).
		BorderForeground(mutedColor).
		Padding(0, 1).
		MarginTop(1)
		
	successStyle = lipgloss.NewStyle().
		Foreground(successColor).
		Bold(true)
		
	errorStyle = lipgloss.NewStyle().
		Foreground(errorColor).
		Bold(true)
		
	mutedStyle = lipgloss.NewStyle().
		Foreground(mutedColor)
}

// 启动配置UI
// 快捷键设置界面
func (m configModel) viewHotkeyScreen() string {
	t := i18n.T()
	s := titleStyle.Render(t.HotkeySettings)
	s += "\n\n"
	
	// 如果正在录制快捷键
	if m.recordingHotkey {
		s += warningStyle.Render(i18n.T().RecordingHotkey) + "\n\n"
		
		switch m.cursor {
		case 0:
			s += normalStyle.Render(i18n.T().SetMonitorHotkey) + "\n"
		case 1:
			s += normalStyle.Render(i18n.T().SetSwitchPromptHotkey) + "\n"
		}
		
		s += "\n" + dimStyle.Render(i18n.T().PressDesiredHotkey) + "\n"
		s += normalStyle.Render(i18n.T().DetectedKeys + ": ")
		if m.hotkeyBuffer != "" {
			s += successStyle.Render(m.hotkeyBuffer) + "\n"
		} else {
			s += dimStyle.Render(i18n.T().WaitingForKeys) + "\n"
		}
		
		s += "\n" + helpStyle.Render(i18n.T().HotkeyTip + ":") + "\n"
		s += dimStyle.Render("  • " + i18n.T().HoldModifier) + "\n"
		s += dimStyle.Render("  • " + i18n.T().DetectedAutoSave) + "\n"
		s += dimStyle.Render("  • " + i18n.T().PressEscCancel) + "\n"
		
		return s
	}
	
	// 正常显示模式
	hotkeys := []struct {
		name string
		key  string
		desc string
	}{
		{t.ToggleMonitor, m.config.HotkeyToggle, t.StartMonitor + "/" + t.StopMonitor},
		{t.SwitchPromptKey, m.config.HotkeySwitch, t.TranslateStyle},
	}
	
	for i, hk := range hotkeys {
		cursor := "  "
		style := normalStyle
		if i == m.cursor {
			cursor = "▶ "
			style = selectedStyle
		}
		
		key := hk.key
		if key == "" {
			key = t.NotConfigured
		}
		
		line := fmt.Sprintf("%s%s: %s - %s", cursor, hk.name, key, hk.desc)
		s += style.Render(line) + "\n"
	}
	
	s += "\n" + helpStyle.Render(t.PressEnterToSet + "，" + t.PressDeleteToClear)
	s += "\n\n" + helpStyle.Render(t.HelpBack)
	
	return s
}

// 关于界面
func (m configModel) viewAboutScreen() string {
	t := i18n.T()
	s := titleStyle.Render(t.About)
	s += "\n\n"
	
	s += successStyle.Render("xiaoniao " + APP_VERSION) + "\n\n"
	
	s += normalStyle.Render(t.Author + "：") + mutedStyle.Render("梨梨果") + "\n"
	s += normalStyle.Render(t.License + "：") + mutedStyle.Render("GPL-3.0 License") + "\n"
	s += normalStyle.Render(t.ProjectUrl + "：") + mutedStyle.Render("https://github.com/kaminoguo/xiaoniao") + "\n\n"
	
	s += warningStyle.Render(t.SupportAuthor) + "\n"
	s += mutedStyle.Render(t.PriceNote) + "\n"
	s += mutedStyle.Render(t.ShareNote) + "\n\n"
	
	s += successStyle.Render(t.ThanksForUsing) + "\n\n"
	
	s += helpStyle.Render(t.BackToMainMenu)
	
	return boxStyle.Render(s)
}

// 更新关于界面
func (m configModel) updateAboutScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back), key.Matches(msg, keys.Quit):
		m.screen = mainScreen
		m.cursor = 8 // 返回到关于选项
	}
	return m, nil
}

func (m configModel) updateHotkeyScreen(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// 如果正在录制快捷键
	if m.recordingHotkey {
		// ESC 取消录制
		if msg.Type == tea.KeyEsc {
			m.recordingHotkey = false
			m.hotkeyBuffer = ""
			m.testResult = ""
			return m, nil
		}
		
		// Enter 确认保存
		if msg.Type == tea.KeyEnter && m.hotkeyBuffer != "" {
			// 检查快捷键冲突
			if hasConflict, conflictInfo := checkHotkeyConflict(m.hotkeyBuffer); hasConflict {
				m.testResult = errorStyle.Render(fmt.Sprintf("❌ 快捷键已被占用: %s", conflictInfo))
				return m, nil
			}
			
			// 配置系统快捷键
			var function string
			switch m.cursor {
			case 0:
				function = "toggle"
				m.config.HotkeyToggle = m.hotkeyBuffer
			case 1:
				function = "switch"
				m.config.HotkeySwitch = m.hotkeyBuffer
			}
			
			if err := configureSystemHotkey(function, m.hotkeyBuffer); err != nil {
				m.testResult = warningStyle.Render(fmt.Sprintf("⚠️ 无法自动配置系统快捷键: %v\n请手动在系统设置中配置", err))
			} else {
				m.testResult = successStyle.Render("✅ 快捷键已配置成功")
			}
			
			// 保存配置
			config = *m.config
			saveConfig()
			
			m.recordingHotkey = false
			m.hotkeyBuffer = ""
			
			// 2秒后清除提示
			return m, tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
				return "clear_hotkey_result"
			})
		}
		
		// 构建快捷键组合 - 支持多修饰键组合
		var parts []string
		
		// 检测修饰键组合（按固定顺序：Ctrl、Alt、Shift）
		hasCtrl := false
		hasAlt := false  
		hasShift := false
		var mainKey string
		
		// 先检测Ctrl键
		if msg.Type >= tea.KeyCtrlA && msg.Type <= tea.KeyCtrlZ {
			hasCtrl = true
			mainKey = string('A' + (msg.Type - tea.KeyCtrlA))
		} else if msg.Type == tea.KeyCtrlBackslash {
			hasCtrl = true
			mainKey = "\\"
		}
		
		// 检测Alt键
		if msg.Alt {
			hasAlt = true
			// 如果还没有主键（没有Ctrl组合），从消息中提取
			if mainKey == "" {
				key := msg.String()
				if strings.Contains(key, "+") {
					// 处理 "alt+x" 或 "ctrl+alt+x" 格式
					parts := strings.Split(key, "+")
					if len(parts) > 0 {
						// 取最后一个部分作为主键
						lastPart := parts[len(parts)-1]
						if lastPart != "" && lastPart != "alt" && lastPart != "ctrl" {
							mainKey = strings.ToUpper(lastPart)
						}
					}
				} else if len(key) == 1 {
					// 单字符Alt组合
					mainKey = strings.ToUpper(key)
				}
			}
		}
		
		// 如果没有检测到Ctrl或Alt，检查其他情况
		if !hasCtrl && !hasAlt {
			if len(msg.String()) == 1 {
				r := []rune(msg.String())[0]
				if r >= 'A' && r <= 'Z' {
					// 大写字母，说明按了Shift
					hasShift = true
					mainKey = string(r)
				} else if r >= 'a' && r <= 'z' {
					// 小写字母，单独的按键
					mainKey = strings.ToUpper(string(r))
				} else {
					// 其他字符
					switch string(r) {
					case "!":
						hasShift = true
						mainKey = "1"
					case "@":
						hasShift = true
						mainKey = "2"
					case "#":
						hasShift = true
						mainKey = "3"
					case "$":
						hasShift = true
						mainKey = "4"
					case "%":
						hasShift = true
						mainKey = "5"
					case "^":
						hasShift = true
						mainKey = "6"
					case "&":
						hasShift = true
						mainKey = "7"
					case "*":
						hasShift = true
						mainKey = "8"
					case "(":
						hasShift = true
						mainKey = "9"
					case ")":
						hasShift = true
						mainKey = "0"
					default:
						// 数字或其他单键
						mainKey = strings.ToUpper(string(r))
					}
				}
			} else {
				// 功能键
				switch msg.Type {
				case tea.KeyF1:
					mainKey = "F1"
				case tea.KeyF2:
					mainKey = "F2"
				case tea.KeyF3:
					mainKey = "F3"
				case tea.KeyF4:
					mainKey = "F4"
				case tea.KeyF5:
					mainKey = "F5"
				case tea.KeyF6:
					mainKey = "F6"
				case tea.KeyF7:
					mainKey = "F7"
				case tea.KeyF8:
					mainKey = "F8"
				case tea.KeyF9:
					mainKey = "F9"
				case tea.KeyF10:
					mainKey = "F10"
				case tea.KeyF11:
					mainKey = "F11"
				case tea.KeyF12:
					mainKey = "F12"
				case tea.KeyTab:
					mainKey = "Tab"
				case tea.KeySpace:
					mainKey = "Space"
				case tea.KeyLeft:
					mainKey = "Left"
				case tea.KeyRight:
					mainKey = "Right"
				case tea.KeyUp:
					mainKey = "Up"
				case tea.KeyDown:
					mainKey = "Down"
				case tea.KeyHome:
					mainKey = "Home"
				case tea.KeyEnd:
					mainKey = "End"
				case tea.KeyPgUp:
					mainKey = "PageUp"
				case tea.KeyPgDown:
					mainKey = "PageDown"
				case tea.KeyDelete:
					mainKey = "Delete"
				case tea.KeyInsert:
					mainKey = "Insert"
				}
			}
		}
		
		// 构建快捷键字符串（按标准顺序：Ctrl、Alt、Shift、主键）
		if hasCtrl {
			parts = append(parts, "Ctrl")
		}
		if hasAlt {
			parts = append(parts, "Alt")
		}
		if hasShift {
			parts = append(parts, "Shift")
		}
		if mainKey != "" {
			parts = append(parts, mainKey)
		}
		
		// 更灵敏的检测 - 只要有修饰键和主键就立即响应
		if len(parts) > 0 {
			m.hotkeyBuffer = strings.Join(parts, "+")
			
			// 放宽条件：只要有修饰键（Ctrl或Alt）+ 主键就显示
			if (hasCtrl || hasAlt) && mainKey != "" {
				// 实时检查冲突
				if hasConflict, conflictInfo := checkHotkeyConflict(m.hotkeyBuffer); hasConflict {
					m.testResult = warningStyle.Render(fmt.Sprintf("⚠️ 冲突: %s", conflictInfo))
				} else {
					m.testResult = successStyle.Render(fmt.Sprintf("✅ 检测到: %s (按Enter确认)", m.hotkeyBuffer))
				}
			}
		}
		
		return m, nil
	}
	
	// 正常的菜单导航
	switch {
	case key.Matches(msg, keys.Back):
		m.screen = mainScreen
		m.cursor = 4 // 返回到快捷键设置选项
		
	case key.Matches(msg, keys.Up):
		if m.cursor > 0 {
			m.cursor--
		}
		
	case key.Matches(msg, keys.Down):
		if m.cursor < 1 { // 只有2个快捷键选项
			m.cursor++
		}
		
	case key.Matches(msg, keys.Enter):
		// 开始录制快捷键
		m.recordingHotkey = true
		m.hotkeyBuffer = ""
		m.testResult = ""
		if m.modifierKeys == nil {
			m.modifierKeys = make(map[string]bool)
		}
		
	case msg.Type == tea.KeyDelete || msg.Type == tea.KeyBackspace:
		// 清除快捷键
		var function string
		switch m.cursor {
		case 0:
			m.config.HotkeyToggle = ""
			function = "toggle"
		case 1:
			m.config.HotkeySwitch = ""
			function = "switch"
		}
		// 移除系统快捷键配置
		removeSystemHotkey(function)
		config = *m.config
		saveConfig()
		
		// 显示提示信息
		m.testResult = "⚠️ 快捷键已清除。如果守护进程正在运行，请重启以生效。"
		
		// 注意：如果守护进程正在运行，需要重启守护进程才能使快捷键更改生效
		// 这是由于当前架构限制，未来版本会改进自动重载机制
	}
	
	return m, nil
}

func showConfigUI() {
	// 应用保存的主题
	if config.Theme != "" {
		applyTheme(config.Theme)
	}
	
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("配置界面错误: %v\n", err)
	}
}