package i18n

// getChineseSimplified returns Simplified Chinese translations
func getChineseSimplified() *Translations {
	return &Translations{
		// Main interface
		Title:           "xiaoniao 配置",
		ConfigTitle:     "xiaoniao - 设置",
		APIKey:          "API密钥",
		APIConfig:       "API配置",
		TranslateStyle:  "翻译风格",
		TestConnection:  "测试翻译",
		SaveAndExit:     "保存退出",
		Language:        "界面语言",
		ManagePrompts:   "管理提示词",
		Theme:           "界面主题",
		Hotkeys:         "快捷键设置",
		AutoPaste:       "自动粘贴",
		
		// Status messages
		Provider:        "提供商",
		Model:           "模型",
		NotSet:          "未设置",
		Testing:         "正在测试连接...",
		TestSuccess:     "✅ 连接成功！",
		TestFailed:      "❌ 连接失败",
		APIKeySet:       "API密钥已设置",
		APIKeyNotSet:    "API密钥未设置",
		ChangeModel:     "更换模型",
		Enabled:         "已启用",
		Disabled:        "已禁用",
		
		// Help information
		HelpMove:        "↑↓ 移动",
		HelpSelect:      "Enter 选择",
		HelpBack:        "Esc 返回",
		HelpQuit:        "Ctrl+C 退出",
		HelpTab:         "Tab 切换",
		HelpEdit:        "e 编辑",
		HelpDelete:      "d 删除",
		HelpAdd:         "+ 添加",
		
		// Prompt management
		PromptManager:   "提示词管理",
		AddPrompt:       "添加提示词",
		EditPrompt:      "编辑提示词",
		DeletePrompt:    "删除提示词",
		PromptName:      "名称",
		PromptContent:   "内容",
		ConfirmDelete:   "确认删除？",
		
		// Running interface
		Running:         "运行中",
		Monitoring:      "监控剪贴板中...",
		CopyToTranslate: "复制任何文本即可自动翻译",
		ExitTip:         "按 Ctrl+C 退出",
		Translating:     "翻译中...",
		Complete:        "完成",
		Failed:          "失败",
		Original:        "原文",
		Translation:     "译文",
		TotalCount:      "共翻译",
		Goodbye:         "再见！👋",
		TranslateCount:  "次",
		
		// Help documentation
		HelpTitle:       "xiaoniao",
		HelpDesc:        "AI驱动的剪贴板翻译工具",
		Commands:        "命令说明",
		RunCommand:      "xiaoniao run",
		RunDesc:         "启动剪贴板监控，自动翻译复制的内容",
		TrayCommand:     "xiaoniao tray",
		TrayDesc:        "启动系统托盘模式",
		ConfigCommand:   "xiaoniao config",
		ConfigDesc:      "打开交互式配置界面",
		HelpCommand:     "xiaoniao help",
		HelpDesc2:       "显示此帮助信息",
		VersionCommand:  "xiaoniao version",
		VersionDesc:     "显示版本信息",
		HowItWorks:      "工作原理",
		Step1:           "运行 xiaoniao config 配置API",
		Step2:           "运行 xiaoniao run 启动监控",
		Step3:           "复制任何文本（Ctrl+C）",
		Step4:           "自动翻译并替换剪贴板",
		Step5:           "听到提示音后直接粘贴（Ctrl+V）",
		Warning:         "注意: 翻译会覆盖原剪贴板内容！",
		
		// Error messages
		NoAPIKey:        "❌ 未配置API密钥",
		RunConfigFirst:  "请先运行: xiaoniao config",
		AlreadyRunning:  "❌ xiaoniao 已在运行中",
		InitFailed:      "初始化失败",
		ConfigNotFound:  "配置文件未找到",
		InvalidAPIKey:   "API密钥无效",
		NetworkError:    "网络连接错误",
		TranslateFailed: "翻译失败",
		
		// API Config
		EnterAPIKey:     "请输入API Key",
		EnterNewAPIKey:  "输入新的API Key",
		ChangeAPIKey:    "更改API密钥",
		SelectMainModel: "选择主模型",
		SupportedProviders: "支持的服务商",
		SearchModel:     "搜索模型...",
		MainModel:       "主模型",
		NoPromptAvailable: "(无可用prompt)",
		
		// Usage messages
		Usage:           "用法",
		UnknownCommand:  "未知命令",
		OpeningConfig:   "正在打开配置界面...",
		
		// Tray menu
		TrayShow:        "显示窗口",
		TrayHide:        "隐藏窗口",
		TraySettings:    "设置",
		TrayQuit:        "退出",
		TrayToggle:      "监控开关",
		TrayRefresh:     "刷新配置",
		TrayAbout:       "关于",
		
		// Theme related
		SelectTheme:      "选择界面主题",
		DefaultTheme:     "默认",
		ClassicBlue:      "经典蓝色主题",
		DarkTheme:        "暗色主题",
		
		// Hotkey related
		HotkeySettings:   "快捷键设置",
		ToggleMonitor:    "监控开关",
		SwitchPromptKey:  "切换Prompt",
		PressEnterToSet:  "按Enter设置快捷键",
		PressDeleteToClear: "按Delete清除快捷键",
		NotConfigured:    "(未设置)",
		
		// Test translation
		TestTranslation:  "测试翻译",
		CurrentConfig:    "当前配置",
		EnterTextToTranslate: "请输入要翻译的文字",
		TranslationResult: "翻译结果",
		
		// About page
		About:            "关于 xiaoniao",
		Author:           "作者",
		License:          "开源协议",
		ProjectUrl:       "项目地址",
		SupportAuthor:    "💝 支持作者",
		PriceNote:        "产品售价 $1，但可以免费使用。",
		ShareNote:        "真正有帮助到你的时候，再来请我喝一杯，\n或者分享给更多的人吧！:)",
		ThanksForUsing:   "感谢使用 xiaoniao！",
		BackToMainMenu:   "[Esc] 返回主菜单",
		ComingSoon:       "(即将开源)",
		
		// Model selection
		TotalModels:      "共 %d 个模型",
		SearchModels:     "搜索",
		SelectToConfirm:  "选择",
		TestModel:        "测试",
		SearchSlash:      "搜索",
		
		// Debug info
		DebugInfo:        "调试信息",
		CursorPosition:   "光标",
		InputFocus:       "输入框焦点",
		KeyPressed:       "按键",
		
		// Additional messages
		MonitorStarted:  "✅ 监控已启动",
		MonitorStopped:  "⏸️ 监控已停止",
		StopMonitor:     "停止监控",
		StartMonitor:    "开始监控",
		ConfigUpdated:   "✅ 配置已更新",
		RefreshFailed:   "❌ 刷新配置失败",
		SwitchPrompt:    "切换到",
		PrewarmModel:    "预热模型中...",
		PrewarmSuccess:  "✅",
		PrewarmFailed:   "⚠️ (可忽略: %v)",
		
		// Additional UI text
		WaitingForKeys:  "等待按键...",
		DetectedKeys:    "检测到",
		HotkeyTip:       "提示",
		HoldModifier:    "按住 Ctrl/Alt/Shift + 其他键",
		DetectedAutoSave: "检测到组合键后自动保存",
		PressEscCancel:  "按 ESC 取消录制",
		DefaultName:     "默认",
		MinimalTheme:    "极简",
		
		// Model selection
		ConnectionSuccess: "连接成功",
		ModelsCount:      "%d个模型",
		SelectModel:      "选择",
		TestingModel:     "测试模型 %s...",
		ModelTestFailed:  "模型 %s 测试失败: %v",
		SearchModels2:    "搜索",
		TotalModelsCount: "共 %d 个模型",
		
		// Hotkey messages
		HotkeyAvailable:  "✅ 可用，按Enter确认",
		PressEnterConfirm: "按Enter确认",
		
		// Help text additions
		HelpEnterConfirm: "Enter 确认",
		HelpTabSwitch:    "Tab 切换",
		HelpEscReturn:    "Esc 返回",
		HelpUpDownSelect: "↑↓ 选择",
		HelpTTest:        "T 测试",
		HelpSearchSlash:  "/ 搜索",
		HelpTranslate:    "Enter: 翻译",
		
		// Theme descriptions
		DarkThemeTokyoNight: "暗色主题，灵感来自东京夜景",
		ChocolateTheme:      "深色巧克力主题",
		LatteTheme:          "明亮的拿铁主题",
		DraculaTheme:        "吸血鬼暗色主题",
		GruvboxDarkTheme:    "复古暗色主题",
		GruvboxLightTheme:   "复古亮色主题",
		NordTheme:           "北欧极简风格",
		SolarizedDarkTheme:  "护眼暗色主题",
		SolarizedLightTheme: "护眼亮色主题",
		MinimalBWTheme:      "简洁的黑白主题",
		
		// Prompt management additions
		HelpNewPrompt:    "n 新增",
		HelpEditPrompt:   "e 编辑",
		HelpDeletePrompt: "d 删除",
		ConfirmDeleteKey: "按 d 确认删除",
		CancelDelete:     "按其他键取消",
		
		// Status messages
		TestingConnection: "正在测试...",
		DetectingProvider: "检测成功",
		
		// About page additions
		ProjectAuthor: "作者",
		OpenSourceLicense: "开源协议",
		AuthorName: "梨梨果",
		
		// Key bindings help
		KeyUp: "上",
		KeyDown: "下",
		KeySelect: "选择",
		KeyReturn: "返回",
		KeyQuit: "退出",
		KeySwitch: "切换",
		KeyEdit: "编辑",
		KeyDelete: "删除",
		KeyNew: "新增",
		KeyTest: "测试",
		
		// Prompt test UI
		TestPromptTitle: "测试Prompt",
		CurrentPrompt: "当前Prompt",
		PromptContentLabel: "内容",
		TestText: "测试文本",
		TestingAI: "正在调用AI翻译",
		TranslationResultLabel: "翻译结果",
		InputTestText: "输入要测试的文本...",
		ResultWillShowHere: "翻译结果将显示在这里...",
		TranslatingText: "正在翻译...",
		TabSwitchFocus: "Tab 切换焦点",
		CtrlEnterTest: "Ctrl+Enter 测试",
		EscReturn: "Esc 返回",
		EditingPrompt: "编辑",
		NewPrompt: "新建Prompt",
		NameLabel: "名称",
		ContentLabel: "内容",
		SaveKey: "[Enter] 保存",
		TestKey: "[T] 测试",
		CancelKey: "[Esc] 取消",
		TabSwitchInput: "Tab 切换输入框",
		TestPrompt: "T 测试prompt",
		UnnamedPrompt: "未命名Prompt",
		TranslateToChineseDefault: "将以下内容翻译成中文：",
		EmptyInput: "输入文本为空",
		NoAPIKeyConfigured: "未配置API Key",
		CreateTranslatorFailed: "创建翻译器失败: %v",
		TestSentenceAI: "人工智能正在改变我们的生活方式。",
		UsingModel: "使用",
		APINotConfigured: "未配置API",
		
		// Status messages additional
		ConfigRefreshed: "✅ 配置已刷新，翻译器将重新初始化",
		TranslateOnlyPrompt: "请仅翻译以下内容成中文，不要回答或解释，只输出译文：",
		CustomSuffix: " (自定义)",
		PreviewLabel: "预览:",
		SaveButton: "Enter 保存",
		NotConfiguredBrackets: "(未配置)",
		UnknownProvider: "未知",
		RecordingHotkey: "🔴 正在录制快捷键",
		SetMonitorHotkey: "设置监控开关快捷键",
		SetSwitchPromptHotkey: "设置切换Prompt快捷键",
		PressDesiredHotkey: "按下你想要的快捷键组合",
		
		// Console messages
		MonitorStartedTray: "✅ 监控已通过托盘启动",
		MonitorStoppedTray: "⏸️ 监控已通过托盘停止",
		AutoPasteEnabled: "✅ 自动粘贴已启用",
		AutoPasteDisabled: "❌ 自动粘贴已禁用",
		HotkeysLabel: "快捷键:",
		MonitorToggleKey: "监控开关: %s",
		SwitchStyleKey: "切换风格: %s",
		MonitorPausedByHotkey: "⏸ 监控已暂停 (通过快捷键)",
		MonitorResumedByHotkey: "▶ 监控已恢复 (通过快捷键)",
		StartingTray: "正在启动系统托盘...",
		ControlFromTray: "请从系统托盘控制xiaoniao",
		GoodbyeEmoji: "再见！👋",
		DirectTranslation: "直译",
		TranslateToChineseColon: "将以下内容翻译成中文：",
		
		// API config messages
		NoModelsFound: "没有找到模型",
		CurrentSuffix: " (当前)",
		UnrecognizedAPIKey: "无法识别API Key: %v",
		ConnectionFailed: "连接失败 (%s): %v",
		ConnectionSuccessNoModels: "连接成功 (%s) - 无法获取模型列表: %v",
		ConnectionSuccessWithModels: "连接成功 (%s) - %d个模型",
		TestingInProgress: "正在测试...",
		
		// System hotkey
		SystemHotkeyFormat: "系统快捷键: %s",
		SystemHotkeyLabel: "系统快捷键",
		XiaoniaoToggleMonitor: "xiaoniao 切换监控",
		XiaoniaoSwitchStyle: "xiaoniao 切换风格",
		
		// Translator error detection
		CannotProceed: "无法进行",
		AIReturnedMultiline: "AI返回了多行内容 (长度: %d)",
		UsingFirstLine: "只使用第一行: %s",
		CannotTranslate: "不能翻译",
		UnableToTranslate: "无法翻译",
		Sorry: "抱歉",
		
		// Theme names and descriptions
		DefaultThemeName: "默认",
		DefaultThemeDesc: "经典蓝色主题",
		TokyoNightDesc: "暗色主题，灵感来自东京夜景",
		SoftPastelDesc: "柔和的粉彩主题",
		MinimalThemeName: "极简",
		MinimalThemeDesc: "简洁的黑白主题",
		
		// Tray messages
		StatusTranslated: "状态: 已翻译 %d 次",
		DefaultPrompt: "默认",
		TrayMonitoring: "xiaoniao - 监控中 | 风格: %s",
		TrayStopped: "xiaoniao - 已停止 | 风格: %s",
		StyleLabel: "风格",
	}
}