package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
	"github.com/kaminoguo/xiaoniao/internal/translator"
)

// 使用主文件中的screen常量，这些只是作为参考
// modelSelectScreen定义在主文件中的screen枚举里

type apiConfigModel struct {
	screen         int
	apiKeyInput    textinput.Model
	modelSearch    textinput.Model
	
	// API相关
	provider       string
	apiKey         string
	baseURL        string
	
	// 模型相关
	allModels      []string
	filteredModels []string
	selectedModel  int
	currentModel   string
	
	// 状态
	testing        bool
	testResult     string
	testSuccess    bool
	
	// UI
	width          int
	height         int
}

// 扩展configModel以包含API配置
func (m *configModel) initAPIConfig() {
	// 初始化API Key输入框
	apiInput := textinput.New()
	apiInput.Placeholder = "sk-..."
	apiInput.CharLimit = 200
	apiInput.Width = 60
	apiInput.EchoMode = textinput.EchoPassword
	apiInput.EchoCharacter = '•'
	
	// 只有在没有API key时才让输入框获得焦点
	if m.config.APIKey == "" {
		apiInput.Focus()
	} else {
		apiInput.Blur()  // 确保输入框失焦
	}
	
	// 初始化模型搜索框
	searchInput := textinput.New()
	searchInput.Placeholder = i18n.T().SearchModel
	searchInput.CharLimit = 50
	searchInput.Width = 40
	
	m.apiKeyInput = apiInput
}

// API配置主界面
func (m configModel) viewAPIConfigScreen() string {
	t := i18n.T()
	s := titleStyle.Render(t.APIConfig)
	s += "\n"

	// 显示当前配置状态
	if m.changingAPIKey {
		// 正在更改API密钥
		s += warningStyle.Render(t.ChangeAPIKey) + "\n\n"

		// 先选择Provider - 使用固定列表显示
		if m.selectingProvider || m.config.Provider == "" {
			s += t.SelectProvider + ":\n\n"
			providers := translator.GetSupportedProviders()

			// 固定列表显示所有Provider - 两列布局
			const colWidth = 25
			halfLen := (len(providers) + 1) / 2

			for i := 0; i < halfLen; i++ {
				// 左列
				if i < len(providers) {
					if i == m.providerCursor {
						s += selectedStyle.Render(fmt.Sprintf("▶ %-*s", colWidth, providers[i]))
					} else {
						s += normalStyle.Render(fmt.Sprintf("  %-*s", colWidth, providers[i]))
					}
				}

				// 右列
				rightIdx := i + halfLen
				if rightIdx < len(providers) {
					if rightIdx == m.providerCursor {
						s += selectedStyle.Render(fmt.Sprintf("▶ %s", providers[rightIdx]))
					} else {
						s += normalStyle.Render(fmt.Sprintf("  %s", providers[rightIdx]))
					}
				}
				s += "\n"
			}
			s += "\n" + helpStyle.Render(fmt.Sprintf("%s | %s | %s", t.HelpUpDownSelect, t.HelpEnterConfirm, t.HelpEscReturn))
		} else {
			// 已选择Provider，输入API Key
			if m.config.Provider != "" {
				s += successStyle.Render("Provider: " + m.config.Provider) + "\n\n"
			}
			s += t.EnterNewAPIKey + ":\n"
			s += m.apiKeyInput.View() + "\n\n"

			// 支持的Provider提示
			s += mutedStyle.Render(t.SupportedProviders + ":") + "\n"
			providers := translator.GetSupportedProviders()
			// 显示所有支持的providers，以3列整齐排列
			cols := 3
			for i := 0; i < len(providers); i += cols {
				line := ""
				for j := 0; j < cols && i+j < len(providers); j++ {
					provider := providers[i+j]
					// 固定宽度20个字符，确保对齐
					line += fmt.Sprintf("• %-22s", provider)
				}
				s += mutedStyle.Render(line) + "\n"
			}
		}
	} else if m.config.APIKey != "" {
		// 已配置
		provider := m.config.Provider
		if provider == "" {
			provider = "Unknown"
		}
		
		s += successStyle.Render("✓ " + t.APIKeySet) + "\n"
		s += normalStyle.Render(fmt.Sprintf("Provider: %s", provider)) + "\n"
		s += normalStyle.Render(fmt.Sprintf("%s: %s", t.MainModel, m.config.Model)) + "\n"
		s += "\n"
		
		// 选项菜单
		options := []string{
			"1. " + t.TestConnection,
			"2. " + t.SelectMainModel,
			"3. " + t.ChangeAPIKey,
		}
		
		for i, option := range options {
			if i == m.cursor {
				s += selectedStyle.Render("▶ " + option) + "\n"
			} else {
				s += normalStyle.Render("  " + option) + "\n"
			}
		}
	} else {
		// 未配置 - 先选择Provider
		s += errorStyle.Render("✗ " + t.APIKeyNotSet) + "\n\n"

		if m.selectingProvider {
			s += t.SelectProvider + ":\n\n"
			providers := translator.GetSupportedProviders()

			// 固定列表显示所有Provider - 两列布局
			const colWidth = 25
			halfLen := (len(providers) + 1) / 2

			for i := 0; i < halfLen; i++ {
				// 左列
				if i < len(providers) {
					if i == m.providerCursor {
						s += selectedStyle.Render(fmt.Sprintf("▶ %-*s", colWidth, providers[i]))
					} else {
						s += normalStyle.Render(fmt.Sprintf("  %-*s", colWidth, providers[i]))
					}
				}

				// 右列
				rightIdx := i + halfLen
				if rightIdx < len(providers) {
					if rightIdx == m.providerCursor {
						s += selectedStyle.Render(fmt.Sprintf("▶ %s", providers[rightIdx]))
					} else {
						s += normalStyle.Render(fmt.Sprintf("  %s", providers[rightIdx]))
					}
				}
				s += "\n"
			}
			s += "\n" + helpStyle.Render(fmt.Sprintf("%s | %s | %s", t.HelpUpDownSelect, t.HelpEnterConfirm, t.HelpEscReturn))
		} else {
			// 已选择Provider，输入API Key
			if m.config.Provider != "" {
				s += successStyle.Render("Provider: " + m.config.Provider) + "\n\n"
			}
			s += t.EnterAPIKey + ":\n"
			s += m.apiKeyInput.View() + "\n\n"

			// 支持的Provider提示
			s += mutedStyle.Render(t.SupportedProviders + ":") + "\n"
			providers := translator.GetSupportedProviders()
			// 显示所有支持的providers，以3列整齐排列
			cols := 3
			for i := 0; i < len(providers); i += cols {
				line := ""
				for j := 0; j < cols && i+j < len(providers); j++ {
					provider := providers[i+j]
					// 固定宽度20个字符，确保对齐
					line += fmt.Sprintf("• %-22s", provider)
				}
				s += mutedStyle.Render(line) + "\n"
			}
		}
	}
	
	// 测试结果
	if m.testResult != "" {
		s += "\n"
		if m.testing {
			s += mutedStyle.Render("⏳ " + t.TestingConnection + "...\n")
		} else if strings.Contains(m.testResult, "成功") || strings.Contains(m.testResult, "Success") {
			s += successStyle.Render(m.testResult) + "\n"
		} else {
			s += errorStyle.Render(m.testResult) + "\n"
		}
	}
	
	// 帮助信息
	s += "\n" + helpStyle.Render(fmt.Sprintf("%s | %s | %s", t.HelpEnterConfirm, t.HelpTabSwitch, t.HelpEscReturn))
	
	return boxStyle.Render(s)
}

// 模型选择界面
func (m configModel) viewModelSelectScreen() string {
	t := i18n.T()
	// 导入版本号
	const VERSION = "v1.1.0"
	title := t.SelectMainModel
	currentModel := m.config.Model
	s := titleStyle.Render(title + " " + VERSION)
	s += "\n"
	
	// 显示provider
	s += normalStyle.Render(fmt.Sprintf("Provider: %s", m.config.Provider)) + "\n"
	s += normalStyle.Render(fmt.Sprintf("%s: %s", t.Model, currentModel)) + "\n\n"
	
	// 搜索框
	s += t.SearchModels2 + ": " + m.promptNameInput.View() + "\n"
	s += strings.Repeat("─", 50) + "\n"
	
	// 动态获取模型列表
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
	
	// 显示模型列表（优化的滚动）
	const HEIGHT = 15  // 增加显示高度
	total := len(models)

	if total == 0 {
		s += mutedStyle.Render("没有找到模型\n")
	} else {
		// 简化滚动计算
		viewStart := 0
		if m.selectedPrompt >= HEIGHT {
			viewStart = m.selectedPrompt - HEIGHT + 1
			if viewStart + HEIGHT > total {
				viewStart = total - HEIGHT
			}
		}
		if viewStart < 0 {
			viewStart = 0
		}

		// 显示模型
		viewEnd := viewStart + HEIGHT
		if viewEnd > total {
			viewEnd = total
		}

		for i := viewStart; i < viewEnd; i++ {
			model := models[i]

			// 标记当前使用的模型
			if model == m.config.Model {
				model += " *"
			}

			if i == m.selectedPrompt {
				s += selectedStyle.Render("▶ " + model) + "\n"
			} else {
				s += normalStyle.Render("  " + model) + "\n"
			}
		}

		// 显示总数和当前位置
		s += "\n" + mutedStyle.Render(fmt.Sprintf("[%d/%d]", m.selectedPrompt+1, total))
	}
	
	// 显示测试结果
	if m.testResult != "" {
		s += "\n"
		if m.testing {
			s += mutedStyle.Render(m.testResult) + "\n"
		} else if strings.Contains(m.testResult, "✅") {
			s += successStyle.Render(m.testResult) + "\n"
		} else if strings.Contains(m.testResult, "❌") {
			s += errorStyle.Render(m.testResult) + "\n"
		} else {
			s += normalStyle.Render(m.testResult) + "\n"
		}
	}
	
	// 帮助信息
	s += "\n" + helpStyle.Render(fmt.Sprintf("%s | %s | %s | %s | %s", t.HelpUpDownSelect, t.HelpEnterConfirm, t.HelpTTest, t.HelpSearchSlash, t.HelpEscReturn))
	
	return boxStyle.Render(s)
}

// 测试API连接（独立函数）
func testAPIConnectionStandalone(apiKey, provider string) (bool, string, []string) {
	t := i18n.T()
	// Provider必须已经设置
	if provider == "" {
		return false, "请先选择服务商", nil
	}
	
	// 记录正在测试的provider
	fmt.Fprintf(os.Stderr, "[DEBUG] Testing connection for provider: %s\n", provider)
	
	// 根据provider创建相应的客户端
	var p translator.Provider

	switch provider {
	case "OpenAI":
		p = translator.NewOpenAIProvider(apiKey, "")
	case "Anthropic":
		p = translator.NewAnthropicProvider(apiKey, "")
	case "Google":
		p = translator.NewGoogleProvider(apiKey, "")
	case "DeepSeek":
		p = translator.NewDeepSeekProvider(apiKey, "")
	case "Moonshot":
		p = translator.NewMoonshotProvider(apiKey, "")
	case "Alibaba":
		p = translator.NewAlibabaProvider(apiKey, "")
	case "Baidu":
		p = translator.NewBaiduProvider(apiKey, "")
	case "ByteDance":
		p = translator.NewByteDanceProvider(apiKey, "")
	case "Zhipu":
		p = translator.NewZhipuProvider(apiKey, "")
	case "01AI":
		p = translator.New01AIProvider(apiKey, "")
	case "Mistral":
		p = translator.NewMistralProvider(apiKey, "")
	case "Cohere":
		p = translator.NewCohereProvider(apiKey, "")
	case "Perplexity":
		p = translator.NewPerplexityProvider(apiKey, "")
	case "xAI":
		p = translator.NewXAIProvider(apiKey, "")
	case "Meta":
		p = translator.NewMetaProvider(apiKey, "")
	case "OpenRouter":
		p = translator.NewOpenRouterProvider(apiKey, "")
	case "Groq":
		p = translator.NewGroqProvider(apiKey, "")
	case "Together":
		p = translator.NewTogetherProvider(apiKey, "")
	case "Replicate":
		p = translator.NewReplicateProvider(apiKey, "")
	case "HuggingFace":
		p = translator.NewHuggingFaceProvider(apiKey, "")
	case "AWS":
		p = translator.NewAWSProvider(apiKey, "")
	case "Azure":
		p = translator.NewAzureProvider(apiKey, "")
	default:
		// 未知的provider，使用通用的OpenAI兼容接口
		p = translator.NewOpenAICompatibleProvider(provider, apiKey, "", "")
	}
	
	// 测试连接
	fmt.Fprintf(os.Stderr, "[DEBUG] Testing connection...\n")
	if err := p.TestConnection(); err != nil {
		fmt.Fprintf(os.Stderr, "[DEBUG] Connection test failed: %v\n", err)
		return false, fmt.Sprintf(t.ConnectionFailed+" (%s): %v", provider, err), nil
	}
	
	// 获取模型列表（可选）
	fmt.Fprintf(os.Stderr, "[DEBUG] Getting model list...\n")
	models, err := p.ListModels()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DEBUG] Failed to get models: %v\n", err)
		// 即使获取模型失败，连接可能是成功的
		return true, fmt.Sprintf("连接成功 (%s) - 无法获取模型列表: %v", provider, err), nil
	}
	
	fmt.Fprintf(os.Stderr, "[DEBUG] Found %d models\n", len(models))
	return true, fmt.Sprintf("连接成功 (%s) - %d个模型", provider, len(models)), models
}

// API配置的子页面状态
type apiConfigState int

const (
	apiStateMain apiConfigState = iota
	apiStateKeyInput
	apiStateTesting
	apiStateModelSelect
	apiStateModelSearch
)

// 扩展configModel以包含API配置状态
type apiConfigData struct {
	state          apiConfigState
	testSuccess    bool
	availableModels []string
	selectedModelIdx int
	modelSearchQuery string
}

// 处理API配置的更新 - 新版本接受tea.Msg以便正确传递给textinput
func (m configModel) updateAPIConfigWithMsg(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 首先尝试转换为KeyMsg
	keyMsg, isKeyMsg := msg.(tea.KeyMsg)
	if !isKeyMsg {
		// 如果不是键盘消息，直接返回
		return m, nil
	}

	// 调用原始的updateAPIConfig
	return m.updateAPIConfig(keyMsg, msg)
}

// 处理API配置的更新 - 修改为接受原始消息
func (m configModel) updateAPIConfig(keyMsg tea.KeyMsg, originalMsg tea.Msg) (tea.Model, tea.Cmd) {
	t := i18n.T()

	// 处理不同的API配置状态
	if m.config.APIKey == "" || m.changingAPIKey {
		// 没有API Key或正在更改

		// 首先处理Provider选择
		if m.selectingProvider {
			providers := translator.GetSupportedProviders()
			halfLen := (len(providers) + 1) / 2

			switch keyMsg.String() {
			case "enter":
				// 确保cursor在有效范围内
				if m.providerCursor >= 0 && m.providerCursor < len(providers) {
					m.config.Provider = providers[m.providerCursor]
					m.selectingProvider = false
					// 重新初始化API输入框以确保它能接收输入
					m.apiKeyInput = textinput.New()
					m.apiKeyInput.Placeholder = "sk-..."
					m.apiKeyInput.CharLimit = 200
					m.apiKeyInput.Width = 60
					m.apiKeyInput.EchoMode = textinput.EchoPassword
					m.apiKeyInput.EchoCharacter = '•'
					m.apiKeyInput.Focus()
					return m, textinput.Blink
				}
				return m, nil

			case "up", "k":
				// 简单的上移逻辑
				if m.providerCursor > 0 {
					m.providerCursor--
				}
				return m, nil

			case "down", "j":
				// 简单的下移逻辑
				if m.providerCursor < len(providers)-1 {
					m.providerCursor++
				}
				return m, nil

			case "left", "h":
				// 从右列移动到左列
				if m.providerCursor >= halfLen {
					m.providerCursor -= halfLen
				}
				return m, nil

			case "right", "l":
				// 从左列移动到右列
				if m.providerCursor < halfLen && m.providerCursor+halfLen < len(providers) {
					m.providerCursor += halfLen
				}
				return m, nil

			case "esc":
				if m.changingAPIKey {
					m.changingAPIKey = false
					m.selectingProvider = false
				} else {
					m.screen = mainScreen
				}
				return m, nil

			default:
				// 忽略其他按键
				return m, nil
			}
		}

		// 处理API Key输入 - 只有在Provider已选择且不在选择Provider时
		if !m.selectingProvider && m.config.Provider != "" {
			switch keyMsg.String() {
			case "enter":
				apiKey := m.apiKeyInput.Value()
				if apiKey != "" {
					// 保存配置
					m.config.APIKey = apiKey
					m.testing = true
					m.changingAPIKey = false  // 重置标志

					// 测试连接
					return m, m.testAPIConnectionWithProvider(apiKey, m.config.Provider)
				}
				return m, nil

			case "esc":
				// 如果是更改API密钥，取消更改
				if m.changingAPIKey {
					m.changingAPIKey = false
					m.selectingProvider = false
				} else {
					m.screen = mainScreen
				}
				return m, nil

			case "tab":
				// Tab键切换到Provider选择
				m.selectingProvider = true
				m.providerCursor = 0
				return m, nil

			default:
				// 处理文本输入 - 使用原始消息
				var cmd tea.Cmd
				m.apiKeyInput, cmd = m.apiKeyInput.Update(originalMsg)
				return m, cmd
			}
		} else if m.config.Provider == "" && !m.selectingProvider {
			// 如果Provider未设置且不在选择，自动开始选择
			m.selectingProvider = true
			m.providerCursor = 0
			return m, nil
		}
	} else {
		// 已有API Key，显示配置菜单
		switch keyMsg.String() {
		case "enter":
			// 根据当前光标位置执行操作
			switch m.cursor {
			case 0:
				// 测试连接
				m.testing = true
				m.testResult = i18n.T().TestingConnection + "..."  // 添加即时反馈
				return m, m.testAPIConnection()
			case 1:
				// 选择主模型
				return m.showModelSelector()
			case 2:
				// 更改API密钥
				m.changingAPIKey = true  // 设置标志
				m.selectingProvider = true  // 先选择Provider
				m.providerCursor = 0
				// 清空Provider和API Key，让用户重新选择
				m.config.Provider = ""
				m.apiKeyInput.SetValue("")
				return m, nil
			default:
				// 处理意外的cursor值
				m.testResult = fmt.Sprintf("调试: cursor=%d", m.cursor)
			}
			
		case "1":
			// 测试连接
			m.cursor = 0  // 确保cursor正确
			m.testing = true
			m.testResult = t.TestingConnection  // 添加即时反馈
			return m, m.testAPIConnection()
			
		case "2":
			// 选择主模型
			m.cursor = 1  // 确保cursor正确
			return m.showModelSelector()
			
		case "3":
			// 更改API密钥
			m.cursor = 2  // 确保cursor正确
			m.changingAPIKey = true  // 设置标志
			m.selectingProvider = true  // 先选择Provider
			m.providerCursor = 0
			// 清空Provider和API Key，让用户重新选择
			m.config.Provider = ""
			m.apiKeyInput.SetValue("")
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

// 使用指定Provider测试API连接
func (m *configModel) testAPIConnectionWithProvider(apiKey, provider string) tea.Cmd {
	return func() tea.Msg {
		// 设置Provider
		m.config.Provider = provider

		// 获取默认模型
		if models, exists := translator.ProviderModels[provider]; exists && len(models) > 0 {
			m.config.Model = models[0]
		}

		// 保存配置
		config = *m.config
		saveConfig()

		// 测试连接
		success, result, _ := testAPIConnectionStandalone(apiKey, provider)

		if success {
			// 返回成功消息
			return fmt.Sprintf("✅ %s", result)
		} else {
			// 返回失败消息
			return fmt.Sprintf("❌ %s", result)
		}
	}
}

// 检测并测试API（保留为兼容）
func (m *configModel) detectAndTestAPI(apiKey string) tea.Cmd {
	return func() tea.Msg {
		// 如果已经有Provider，直接使用
		if m.config.Provider != "" {
			return m.testAPIConnectionWithProvider(apiKey, m.config.Provider)()
		}

		// 否则返回错误
		return fmt.Sprintf("❌ 请先选择Provider")
	}
}

// 测试API连接命令
func (m *configModel) testAPIConnection() tea.Cmd {
	return func() tea.Msg {
		// 调用独立的测试函数
		success, result, _ := testAPIConnectionStandalone(m.config.APIKey, m.config.Provider)
		
		if success {
			return fmt.Sprintf("✅ %s", result)
		} else {
			return fmt.Sprintf("❌ %s", result)
		}
	}
}

// 显示模型选择器
func (m configModel) showModelSelector() (tea.Model, tea.Cmd) {
	// 切换到模型选择界面
	m.screen = modelSelectScreen
	m.selectedPrompt = 0 // 重置选择索引
	m.cursor = 0 // 重置光标

	// 清除模型缓存，强制重新加载
	m.cachedModels = nil
	m.modelsLoaded = false
	
	// 初始化搜索框
	if m.promptNameInput.Value() != "" {
		m.promptNameInput.SetValue("")
	}
	m.promptNameInput.Focus()
	
	return m, textinput.Blink
}

