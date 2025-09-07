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
		s += t.EnterNewAPIKey + ":\n"
		s += m.apiKeyInput.View() + "\n\n"
		
		// 支持的Provider提示
		s += mutedStyle.Render(t.SupportedProviders + ":\n")
		providers := translator.GetSupportedProviders()
		// 显示所有支持的providers，按列排列
		for i, p := range providers {
			s += mutedStyle.Render("  • " + p + "\n")
			if i >= 15 { // 限制显示数量以免界面过长
				remaining := len(providers) - i - 1
				if remaining > 0 {
					s += mutedStyle.Render(fmt.Sprintf("  ...还有%d个\n", remaining))
				}
				break
			}
		}
	} else if m.config.APIKey != "" {
		// 已配置
		provider := m.config.Provider
		if provider == "" {
			provider = translator.DetectProviderByKey(m.config.APIKey)
		}
		
		s += successStyle.Render("✓ " + t.APIKeySet) + "\n"
		s += normalStyle.Render(fmt.Sprintf("Provider: %s", provider)) + "\n"
		s += normalStyle.Render(fmt.Sprintf("%s: %s", t.MainModel, m.config.Model)) + "\n"
		if m.config.FallbackModel != "" {
			s += normalStyle.Render(fmt.Sprintf("%s: %s", t.FallbackModel, m.config.FallbackModel)) + "\n"
		} else {
			s += mutedStyle.Render(fmt.Sprintf("%s: %s", t.FallbackModel, t.NotSet)) + "\n"
		}
		s += "\n"
		
		// 选项菜单
		options := []string{
			"1. " + t.TestConnection,
			"2. " + t.SelectMainModel,
			"3. " + t.SelectFallback,
			"4. " + t.ChangeAPIKey,
		}
		
		for i, option := range options {
			if i == m.cursor {
				s += selectedStyle.Render("▶ " + option) + "\n"
			} else {
				s += normalStyle.Render("  " + option) + "\n"
			}
		}
	} else {
		// 未配置
		s += errorStyle.Render("✗ " + t.APIKeyNotSet) + "\n\n"
		s += t.EnterAPIKey + ":\n"
		s += m.apiKeyInput.View() + "\n\n"
		
		// 支持的Provider提示
		s += mutedStyle.Render(t.SupportedProviders + ":\n")
		providers := translator.GetSupportedProviders()
		// 显示所有支持的providers，按列排列
		for i, p := range providers {
			s += mutedStyle.Render("  • " + p + "\n")
			if i >= 15 { // 限制显示数量以免界面过长
				remaining := len(providers) - i - 1
				if remaining > 0 {
					s += mutedStyle.Render(fmt.Sprintf("  ...还有%d个\n", remaining))
				}
				break
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
	const VERSION = "v1.4"
	title := "🤖 " + t.SelectMainModel
	currentModel := m.config.Model
	if m.selectingFallback {
		title = "🔧 " + t.SelectFallback
		currentModel = m.config.FallbackModel
		if currentModel == "" {
			currentModel = t.NotSet
		}
	}
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
	
	// 显示模型列表（滚动）
	const HEIGHT = 12
	total := len(models)
	
	if total == 0 {
		s += mutedStyle.Render("没有找到模型\n")
	} else {
		// 计算滚动窗口
		viewStart := 0
		if total > HEIGHT {
			if m.selectedPrompt < HEIGHT/2 {
				viewStart = 0
			} else if m.selectedPrompt > total - HEIGHT/2 - 1 {
				viewStart = total - HEIGHT
			} else {
				viewStart = m.selectedPrompt - HEIGHT/2
			}
			
			if viewStart < 0 {
				viewStart = 0
			}
			if viewStart > total - HEIGHT {
				viewStart = total - HEIGHT
			}
		}
		
		// 显示模型
		for i := 0; i < HEIGHT && viewStart+i < total; i++ {
			idx := viewStart + i
			model := models[idx]
			
			// 标记当前使用的模型
			if model == m.config.Model {
				model += " (当前)"
			}
			
			if idx == m.selectedPrompt {
				s += selectedStyle.Render("▶ " + model) + "\n"
			} else {
				prefix := "  "
				if total > HEIGHT {
					if i == 0 && viewStart > 0 {
						prefix = "↑ "
					} else if i == HEIGHT-1 && viewStart+HEIGHT < total {
						prefix = "↓ "
					}
				}
				s += normalStyle.Render(prefix + model) + "\n"
			}
		}
		
		// 显示总数
		if total > HEIGHT {
			s += "\n" + mutedStyle.Render(fmt.Sprintf(t.TotalModelsCount, total))
		}
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
	// 如果provider为空，先尝试检测
	if provider == "" {
		provider = translator.DetectProviderByKey(apiKey)
		if provider == "" {
			// 尝试通过API调用检测
			detectedProvider, models, err := translator.DetectProvider(apiKey)
			if err != nil {
				return false, fmt.Sprintf("无法识别API Key: %v", err), nil
			}
			t := i18n.T()
			return true, fmt.Sprintf("%s: %s, %s", t.DetectingProvider, detectedProvider, fmt.Sprintf(t.ModelsCount, len(models))), models
		}
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
	case "DeepSeek":
		p = translator.NewDeepSeekProvider(apiKey, "")
	case "Moonshot":
		p = translator.NewMoonshotProvider(apiKey, "")
	case "OpenRouter":
		// 使用专门的OpenRouter provider
		p = translator.NewOpenRouterProvider(apiKey, "")
	case "Groq":
		// 使用专门的Groq provider
		p = translator.NewGroqProvider(apiKey, "")
	case "Together", "TogetherAI":
		// 使用专门的Together provider
		p = translator.NewTogetherProvider(apiKey, "")
	default:
		// 对于其他provider，使用通用的OpenAI兼容provider
		// 这包括：Perplexity, Replicate, HuggingFace, Cohere, Mistral, 
		// Google, Zhipu, Baidu, Alibaba, Azure, AWS等
		p = translator.NewOpenAICompatibleProvider(provider, apiKey, "", "")
	}
	
	// 测试连接
	fmt.Fprintf(os.Stderr, "[DEBUG] Testing connection...\n")
	if err := p.TestConnection(); err != nil {
		fmt.Fprintf(os.Stderr, "[DEBUG] Connection test failed: %v\n", err)
		return false, fmt.Sprintf("连接失败 (%s): %v", provider, err), nil
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

// 处理API配置的更新
func (m configModel) updateAPIConfig(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// 添加调试信息
	// m.testResult = fmt.Sprintf("键: %s, cursor: %d", msg.String(), m.cursor)
	
	// 处理不同的API配置状态
	if m.config.APIKey == "" || m.changingAPIKey {
		// 没有API Key或正在更改，显示输入界面
		switch msg.String() {
		case "enter":
			apiKey := m.apiKeyInput.Value()
			if apiKey != "" {
				// 保存并检测
				m.config.APIKey = apiKey
				m.testing = true
				m.changingAPIKey = false  // 重置标志
				
				// 异步检测Provider和测试连接
				return m, m.detectAndTestAPI(apiKey)
			}
			
		case "esc":
			// 如果是更改API密钥，取消更改
			if m.changingAPIKey {
				m.changingAPIKey = false
				// API密钥保持不变，因为我们没有清空它
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
		switch msg.String() {
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
				// 选择副模型
				return m.showFallbackModelSelector()
			case 3:
				// 更改API密钥
				m.changingAPIKey = true  // 设置标志
				m.apiKeyInput.SetValue(m.config.APIKey)
				m.apiKeyInput.Focus()
				return m, nil
			default:
				// 处理意外的cursor值
				m.testResult = fmt.Sprintf("调试: cursor=%d", m.cursor)
			}
			
		case "1":
			// 测试连接
			m.cursor = 0  // 确保cursor正确
			m.testing = true
			m.testResult = "正在测试..."  // 添加即时反馈
			return m, m.testAPIConnection()
			
		case "2":
			// 选择主模型
			m.cursor = 1  // 确保cursor正确
			return m.showModelSelector()
			
		case "3":
			// 选择副模型
			m.cursor = 2  // 确保cursor正确
			return m.showFallbackModelSelector()
			
		case "4":
			// 更改API密钥
			m.cursor = 3  // 确保cursor正确
			m.changingAPIKey = true  // 设置标志
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

// 检测并测试API
func (m *configModel) detectAndTestAPI(apiKey string) tea.Cmd {
	return func() tea.Msg {
		// 检测Provider
		provider := translator.DetectProviderByKey(apiKey)
		if provider == "" {
			// 尝试通过API调用检测
			detectedProvider, models, err := translator.DetectProvider(apiKey)
			if err != nil {
				// 返回错误消息而不是nil
				return fmt.Sprintf("❌ 无法识别Provider: %v", err)
			}
			provider = detectedProvider
			m.config.Provider = provider
			
			// 设置默认模型
			if len(models) > 0 {
				m.config.Model = models[0]
			}
		} else {
			m.config.Provider = provider
			// 获取默认模型
			if models, exists := translator.ProviderModels[provider]; exists && len(models) > 0 {
				m.config.Model = models[0]
			}
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
	m.selectingFallback = false // 选择主模型
	
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

// 显示副模型选择器
func (m configModel) showFallbackModelSelector() (tea.Model, tea.Cmd) {
	// 切换到模型选择界面（使用同一个界面，通过selectingFallback标志区分）
	m.screen = modelSelectScreen  
	m.selectedPrompt = 0 // 重置选择索引
	m.cursor = 0 // 重置光标
	m.selectingFallback = true // 选择副模型
	
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