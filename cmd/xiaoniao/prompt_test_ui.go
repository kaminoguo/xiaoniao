package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
	"github.com/kaminoguo/xiaoniao/internal/translator"
)

// Prompt测试界面扩展
type promptTestModel struct {
	// 输入
	testInput    textarea.Model
	resultOutput textarea.Model
	
	// 当前prompt
	promptName    string
	promptContent string
	
	// 测试状态
	testing       bool
	testResult    string
	testError     string
	
	// 焦点
	focusIndex    int // 0: testInput, 1: resultOutput
	
	// 配置
	config        *Config
}

// 初始化测试界面
func initPromptTestModel(promptName, promptContent string, config *Config) promptTestModel {
	// 测试输入框
	testInput := textarea.New()
	testInput.Placeholder = "输入要测试的文本..."
	testInput.SetWidth(60)
	testInput.SetHeight(4)
	testInput.Focus()
	testInput.CharLimit = 1000
	
	// 结果输出框
	resultOutput := textarea.New()
	resultOutput.Placeholder = "翻译结果将显示在这里..."
	resultOutput.SetWidth(60)
	resultOutput.SetHeight(8)
	
	return promptTestModel{
		testInput:     testInput,
		resultOutput:  resultOutput,
		promptName:    promptName,
		promptContent: promptContent,
		focusIndex:    0,
		config:        config,
	}
}

// 渲染测试界面
func (m promptTestModel) View() string {
	s := titleStyle.Render("测试Prompt")
	s += "\n\n"
	
	// 显示当前prompt信息
	s += lipgloss.NewStyle().Bold(true).Render("当前Prompt: ") + 
		lipgloss.NewStyle().Foreground(accentColor).Render(m.promptName) + "\n"
	s += mutedStyle.Render("内容: " + m.promptContent) + "\n\n"
	
	// 测试输入区域
	inputTitle := "测试文本"
	if m.focusIndex == 0 {
		inputTitle = selectedStyle.Render("▶ " + inputTitle)
	} else {
		inputTitle = normalStyle.Render("  " + inputTitle)
	}
	s += inputTitle + "\n"
	s += m.testInput.View() + "\n\n"
	
	// 显示测试状态
	if m.testing {
		s += mutedStyle.Render("⏳ 正在调用AI翻译...\n\n")
	} else if m.testError != "" {
		s += errorStyle.Render("❌ 错误: " + m.testError) + "\n\n"
	}
	
	// 结果输出区域
	outputTitle := "翻译结果"
	if m.focusIndex == 1 {
		outputTitle = selectedStyle.Render("▶ " + outputTitle)
	} else {
		outputTitle = normalStyle.Render("  " + outputTitle)
	}
	s += outputTitle + "\n"
	s += m.resultOutput.View() + "\n"
	
	// 显示API状态
	if m.config != nil && m.config.APIKey != "" {
		s += "\n" + successStyle.Render(fmt.Sprintf("✓ 使用: %s / %s", 
			m.config.Provider, m.config.Model))
	} else {
		s += "\n" + errorStyle.Render("✗ 未配置API")
	}
	
	// 帮助信息
	s += "\n\n" + helpStyle.Render("Tab 切换焦点 | Ctrl+Enter 测试 | Esc 返回")
	
	return boxStyle.Render(s)
}

// 更新测试界面
func (m promptTestModel) Update(msg tea.Msg) (promptTestModel, tea.Cmd) {
	var cmds []tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			// 切换焦点
			m.focusIndex = (m.focusIndex + 1) % 2
			if m.focusIndex == 0 {
				m.testInput.Focus()
				m.resultOutput.Blur()
			} else {
				m.testInput.Blur()
				m.resultOutput.Focus()
			}
			
		case "ctrl+enter":
			// 执行测试
			if !m.testing && m.testInput.Value() != "" && m.config.APIKey != "" {
				m.testing = true
				m.testError = ""
				m.resultOutput.SetValue("正在翻译...")
				
				// 异步调用翻译API
				return m, m.performTranslation()
			}
			
		case "esc":
			// 返回上一界面
			return m, func() tea.Msg {
				return "back_to_edit"
			}
		}
		
	case translationResult:
		// 处理翻译结果
		m.testing = false
		if msg.err != nil {
			m.testError = msg.err.Error()
			m.resultOutput.SetValue("")
		} else {
			m.testResult = msg.result
			m.resultOutput.SetValue(msg.result)
		}
		
	case tea.WindowSizeMsg:
		// 调整大小
		width := msg.Width - 10
		if width > 80 {
			width = 80
		}
		m.testInput.SetWidth(width)
		m.resultOutput.SetWidth(width)
	}
	
	// 更新当前焦点的组件
	if m.focusIndex == 0 {
		var cmd tea.Cmd
		m.testInput, cmd = m.testInput.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		var cmd tea.Cmd
		m.resultOutput, cmd = m.resultOutput.Update(msg)
		cmds = append(cmds, cmd)
	}
	
	return m, tea.Batch(cmds...)
}

// 翻译结果消息
type translationResult struct {
	result string
	err    error
}

// 执行翻译
func (m *promptTestModel) performTranslation() tea.Cmd {
	return func() tea.Msg {
		// 获取输入文本
		inputText := m.testInput.Value()
		if inputText == "" {
			return translationResult{err: fmt.Errorf("输入文本为空")}
		}
		
		// 检查配置
		if m.config == nil || m.config.APIKey == "" {
			return translationResult{err: fmt.Errorf("未配置API Key")}
		}
		
		// 创建translator - 使用统一的方法支持所有Provider
		translatorConfig := &translator.Config{
			APIKey:       m.config.APIKey,
			Provider:     m.config.Provider,
			Model:        m.config.Model,
			MaxRetries:   1,
			Timeout:      30,
		}
		
		trans, err := translator.NewTranslator(translatorConfig)
		if err != nil {
			return translationResult{err: fmt.Errorf("创建翻译器失败: %v", err)}
		}
		
		// 执行翻译
		result, err := trans.Translate(inputText, m.promptContent)
		if err != nil {
			return translationResult{err: err}
		}
		
		return translationResult{result: result.Translation}
	}
}

// 在prompt编辑界面添加测试按钮
func (m configModel) viewPromptEditScreenWithTest() string {
	t := i18n.T()
	s := titleStyle.Render("✏️ " + t.EditPrompt)
	s += "\n\n"
	
	// 编辑模式标题
	if m.editingPromptIdx >= 0 {
		s += normalStyle.Render("编辑: " + m.prompts[m.editingPromptIdx].Name) + "\n\n"
	} else {
		s += normalStyle.Render("新建Prompt") + "\n\n"
	}
	
	// 名称输入
	s += "名称:\n"
	s += m.promptNameInput.View() + "\n\n"
	
	// 内容输入
	s += "内容:\n"
	s += m.promptContentInput.View() + "\n\n"
	
	// 操作按钮
	buttons := []string{
		"[Enter] 保存",
		"[T] 测试",
		"[Esc] 取消",
	}
	
	buttonStyle := lipgloss.NewStyle().
		Foreground(accentColor).
		Bold(true).
		Padding(0, 1)
	
	var buttonRow string
	for i, btn := range buttons {
		if i > 0 {
			buttonRow += "  "
		}
		buttonRow += buttonStyle.Render(btn)
	}
	s += buttonRow + "\n"
	
	// 帮助信息
	s += "\n" + helpStyle.Render("Tab 切换输入框 | T 测试prompt")
	
	return boxStyle.Render(s)
}

// 集成到主配置模型
func (m configModel) handlePromptTest() (tea.Model, tea.Cmd) {
	// 获取当前编辑的prompt内容
	promptName := m.promptNameInput.Value()
	promptContent := m.promptContentInput.Value()
	
	if promptName == "" {
		promptName = "未命名Prompt"
	}
	if promptContent == "" {
		promptContent = "将以下内容翻译成中文："
	}
	
	// 切换到测试界面
	m.screen = testScreen
	
	return m, nil
}

// 快速测试函数 - 用于在编辑界面直接显示结果
func quickTestPrompt(prompt, testText string, config *Config) (string, error) {
	if config == nil || config.APIKey == "" {
		return "", fmt.Errorf("未配置API")
	}
	
	// 创建provider - 使用统一的Provider创建方法
	translatorConfig := &translator.Config{
		APIKey:       config.APIKey,
		Provider:     config.Provider,
		Model:        config.Model,
		MaxRetries:   1,
		Timeout:      30,
	}
	
	trans, err := translator.NewTranslator(translatorConfig)
	if err != nil {
		return "", fmt.Errorf("创建翻译器失败: %v", err)
	}
	
	// 执行翻译
	result, err := trans.Translate(testText, prompt)
	if err != nil {
		return "", err
	}
	
	return result.Translation, nil
}

// 预设测试文本
var testTexts = []string{
	"Hello, world!",
	"The quick brown fox jumps over the lazy dog.",
	"To be or not to be, that is the question.",
	"人工智能正在改变我们的生活方式。",
	"Yesterday once more, when I was young.",
}