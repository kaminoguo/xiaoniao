package i18n

// getChineseTraditional returns Traditional Chinese translations
func getChineseTraditional() *Translations {
	return &Translations{
		// Main interface
		Title:           "xiaoniao 配置",
		ConfigTitle:     "xiaoniao - 設置",
		APIKey:          "API密鑰",
		APIConfig:       "API配置",
		TranslateStyle:  "翻譯風格",
		TestConnection:  "測試翻譯",
		SaveAndExit:     "保存退出",
		Language:        "界面語言",
		ManagePrompts:   "管理提示詞",
		Theme:           "界面主題",
		Hotkeys:         "快捷鍵設置",
		AutoPaste:       "自動粘貼",
		
		// Status messages
		Provider:        "提供商",
		Model:           "模型",
		NotSet:          "未設置",
		Testing:         "正在測試連接...",
		TestSuccess:     "✅ 連接成功！",
		TestFailed:      "❌ 連接失敗",
		APIKeySet:       "API密鑰已設置",
		APIKeyNotSet:    "API密鑰未設置",
		ChangeModel:     "更換模型",
		Enabled:         "已啟用",
		Disabled:        "已禁用",
		
		// Help information
		HelpMove:        "↑↓ 移動",
		HelpSelect:      "Enter 選擇",
		HelpBack:        "Esc 返回",
		HelpQuit:        "Ctrl+C 退出",
		HelpTab:         "Tab 切換",
		HelpEdit:        "e 編輯",
		HelpDelete:      "d 刪除",
		HelpAdd:         "+ 添加",
		
		// Prompt management
		PromptManager:   "提示詞管理",
		AddPrompt:       "添加提示詞",
		EditPrompt:      "編輯提示詞",
		DeletePrompt:    "刪除提示詞",
		PromptName:      "名稱",
		PromptContent:   "內容",
		ConfirmDelete:   "確認刪除？",
		
		// Running interface
		Running:         "運行中",
		Monitoring:      "監控剪貼板中...",
		CopyToTranslate: "複製任何文本即可自動翻譯",
		ExitTip:         "按 Ctrl+C 退出",
		Translating:     "翻譯中...",
		Complete:        "完成",
		Failed:          "失敗",
		Original:        "原文",
		Translation:     "譯文",
		TotalCount:      "共翻譯",
		Goodbye:         "再見！👋",
		TranslateCount:  "次",
		
		// Help documentation
		HelpTitle:       "xiaoniao",
		HelpDesc:        "AI驅動的剪貼板翻譯工具",
		Commands:        "命令說明",
		RunCommand:      "xiaoniao run",
		RunDesc:         "啟動剪貼板監控，自動翻譯複製的內容",
		TrayCommand:     "xiaoniao tray",
		TrayDesc:        "啟動系統托盤模式",
		ConfigCommand:   "xiaoniao config",
		ConfigDesc:      "打開交互式配置界面",
		HelpCommand:     "xiaoniao help",
		HelpDesc2:       "顯示此幫助信息",
		VersionCommand:  "xiaoniao version",
		VersionDesc:     "顯示版本信息",
		HowItWorks:      "工作原理",
		Step1:           "運行 xiaoniao config 配置API",
		Step2:           "運行 xiaoniao run 啟動監控",
		Step3:           "複製任何文本（Ctrl+C）",
		Step4:           "自動翻譯並替換剪貼板",
		Step5:           "聽到提示音後直接粘貼（Ctrl+V）",
		Warning:         "注意: 翻譯會覆蓋原剪貼板內容！",

		// Tutorial
		Tutorial:        "使用教程",
		TutorialContent: `快速上手指南：

1. 配置API密鑰
   • 在主菜單選擇「API配置」
   • 輸入你的API密鑰（如OpenAI、Anthropic等）
   • 系統會自動識別提供商

2. 選擇模型
   • 設置API後，選擇「選擇模型」
   • 從列表中選擇合適的AI模型

3. 設置快捷鍵（可選）
   • 在主菜單選擇「快捷鍵設置」
   • 設置監控開關和切換prompt的快捷鍵

4. 開始使用
   • Ctrl+C 複製文本觸發翻譯
   • 程序會自動替換剪貼板內容
   • Ctrl+V 粘貼翻譯結果
   • 某些應用可能需要手動粘貼

視頻教程：
• Bilibili: https://www.bilibili.com/video/BV13zpUzhEeK/
• YouTube: https://www.youtube.com/watch?v=iPye0tYkBaY`,

		// Error messages
		NoAPIKey:        "❌ 未配置API密鑰",
		RunConfigFirst:  "請先運行: xiaoniao config",
		AlreadyRunning:  "❌ xiaoniao 已在運行中",
		InitFailed:      "初始化失敗",
		ConfigNotFound:  "配置文件未找到",
		InvalidAPIKey:   "API密鑰無效",
		NetworkError:    "網絡連接錯誤",
		TranslateFailed: "翻譯失敗",
		
		// API Config
		EnterAPIKey:     "請輸入API Key",
		EnterNewAPIKey:  "輸入新的API Key",
		ChangeAPIKey:    "更改API密鑰",
		SelectMainModel: "選擇模型",
		SupportedProviders: "支持的服務商",
		SearchModel:     "搜索模型...",
		MainModel:       "模型",
		NoPromptAvailable: "(無可用prompt)",
		
		// Usage messages
		Usage:           "用法",
		UnknownCommand:  "未知命令",
		OpeningConfig:   "正在打開配置界面...",
		
		// Tray menu
		TrayShow:        "顯示窗口",
		TrayHide:        "隱藏窗口",
		TraySettings:    "設置",
		TrayQuit:        "退出",
		TrayToggle:      "監控開關",
		TrayRefresh:     "刷新配置",
		TrayAbout:       "關於",
		
		// Theme related
		SelectTheme:      "選擇界面主題",
		DefaultTheme:     "默認",
		ClassicBlue:      "經典藍色主題",
		DarkTheme:        "暗色主題",
		
		// Hotkey related
		HotkeySettings:   "快捷鍵設置",
		ToggleMonitor:    "監控開關",
		SwitchPromptKey:  "切換Prompt",
		PressEnterToSet:  "按Enter設置快捷鍵",
		PressDeleteToClear: "按Delete清除快捷鍵",
		NotConfigured:    "(未設置)",
		
		// Test translation
		TestTranslation:  "測試翻譯",
		CurrentConfig:    "當前配置",
		EnterTextToTranslate: "請輸入要翻譯的文字",
		TranslationResult: "翻譯結果",
		
		// About page
		About:            "關於 xiaoniao",
		Author:           "作者",
		License:          "開源協議",
		ProjectUrl:       "項目地址",
		SupportAuthor:    "💝 支持作者",
		PriceNote:        "產品售價 $1，但可以免費使用。",
		ShareNote:        "真正有幫助到你的時候，再來請我喝一杯，\n或者分享給更多的人吧！:)",
		ThanksForUsing:   "感謝使用 xiaoniao！",
		BackToMainMenu:   "[Esc] 返回主菜單",
		ComingSoon:       "(即將開源)",
		
		// Model selection
		TotalModels:      "共 %d 個模型",
		SearchModels:     "搜索",
		SelectToConfirm:  "選擇",
		TestModel:        "測試",
		SearchSlash:      "搜索",
		
		// Debug info
		DebugInfo:        "調試信息",
		CursorPosition:   "光標",
		InputFocus:       "輸入框焦點",
		KeyPressed:       "按鍵",
		
		// Additional messages
		MonitorStarted:  "✅ 監控已啟動",
		MonitorStopped:  "⏸️ 監控已停止",
		StopMonitor:     "停止監控",
		StartMonitor:    "開始監控",
		ConfigUpdated:   "✅ 配置已更新",
		RefreshFailed:   "❌ 刷新配置失敗",
		SwitchPrompt:    "切換到",
		PrewarmModel:    "預熱模型中...",
		PrewarmSuccess:  "✅",
		PrewarmFailed:   "⚠️ (可忽略: %v)",
		
		// Additional UI text
		WaitingForKeys:  "等待按鍵...",
		DetectedKeys:    "檢測到",
		HotkeyTip:       "提示",
		HoldModifier:    "按住 Ctrl/Alt/Shift + 其他鍵",
		DetectedAutoSave: "檢測到組合鍵後自動保存",
		PressEscCancel:  "按 ESC 取消錄製",
		DefaultName:     "默認",
		MinimalTheme:    "極簡",
		
		// Model selection
		ConnectionSuccess: "連接成功",
		ModelsCount:      "%d個模型",
		SelectModel:      "選擇",
		TestingModel:     "測試模型 %s...",
		ModelTestFailed:  "模型 %s 測試失敗: %v",
		SearchModels2:    "搜索",
		TotalModelsCount: "共 %d 個模型",
		
		// Hotkey messages
		HotkeyAvailable:  "✅ 可用，按Enter確認",
		PressEnterConfirm: "按Enter確認",
		
		// Help text additions
		HelpEnterConfirm: "Enter 確認",
		HelpTabSwitch:    "Tab 切換",
		HelpEscReturn:    "Esc 返回",
		HelpUpDownSelect: "↑↓ 選擇",
		HelpTTest:        "T 測試",
		HelpSearchSlash:  "/ 搜索",
		HelpTranslate:    "Enter: 翻譯",
		HelpCtrlSSaveExit: "Ctrl+S 保存並退出",
		HelpCtrlSSave:    "Ctrl+S 保存",
		
		// Theme descriptions
		DarkThemeTokyoNight: "暗色主題，靈感來自東京夜景",
		ChocolateTheme:      "深色巧克力主題",
		LatteTheme:          "明亮的拿鐵主題",
		DraculaTheme:        "吸血鬼暗色主題",
		GruvboxDarkTheme:    "復古暗色主題",
		GruvboxLightTheme:   "復古亮色主題",
		NordTheme:           "北歐極簡風格",
		SolarizedDarkTheme:  "護眼暗色主題",
		SolarizedLightTheme: "護眼亮色主題",
		MinimalBWTheme:      "簡潔的黑白主題",
		
		// Prompt management additions
		HelpNewPrompt:    "n 新增",
		HelpEditPrompt:   "e 編輯",
		HelpDeletePrompt: "d 刪除",
		ConfirmDeleteKey: "按 d 確認刪除",
		CancelDelete:     "按其他鍵取消",
		
		// Status messages
		TestingConnection: "正在測試...",
		DetectingProvider: "檢測成功",
		
		// About page additions
		ProjectAuthor: "作者",
		OpenSourceLicense: "開源協議",
		AuthorName: "梨梨果",
		
		// Key bindings help
		KeyUp: "上",
		KeyDown: "下",
		KeySelect: "選擇",
		KeyReturn: "返回",
		KeyQuit: "退出",
		KeySwitch: "切換",
		KeyEdit: "編輯",
		KeyDelete: "刪除",
		KeyNew: "新增",
		KeyTest: "測試",
		
		// Prompt test UI
		TestPromptTitle: "測試Prompt",
		CurrentPrompt: "當前Prompt",
		PromptContentLabel: "內容",
		TestText: "測試文本",
		TestingAI: "正在調用AI翻譯",
		TranslationResultLabel: "翻譯結果",
		InputTestText: "輸入要測試的文本...",
		ResultWillShowHere: "翻譯結果將顯示在這裡...",
		TranslatingText: "正在翻譯...",
		TabSwitchFocus: "Tab 切換焦點",
		CtrlEnterTest: "Ctrl+Enter 測試",
		EscReturn: "Esc 返回",
		EditingPrompt: "編輯",
		NewPrompt: "新建Prompt",
		NameLabel: "名稱",
		ContentLabel: "內容",
		SaveKey: "[Enter] 保存",
		TestKey: "[T] 測試",
		CancelKey: "[Esc] 取消",
		TabSwitchInput: "Tab 切換輸入框",
		TestPrompt: "T 測試prompt",
		UnnamedPrompt: "未命名Prompt",
		TranslateToChineseDefault: "將以下內容翻譯成中文：",
		EmptyInput: "輸入文本為空",
		NoAPIKeyConfigured: "未配置API Key",
		CreateTranslatorFailed: "創建翻譯器失敗: %v",
		TestSentenceAI: "人工智能正在改變我們的生活方式。",
		UsingModel: "使用",
		APINotConfigured: "未配置API",
		
		// Status messages additional
		ConfigRefreshed: "✅ 配置已刷新，翻譯器將重新初始化",
		TranslateOnlyPrompt: "請僅翻譯以下內容成中文，不要回答或解釋，只輸出譯文：",
		CustomSuffix: " (自定義)",
		PreviewLabel: "預覽:",
		SaveButton: "Enter 保存",
		NotConfiguredBrackets: "(未配置)",
		UnknownProvider: "未知",
		RecordingHotkey: "🔴 正在錄製快捷鍵",
		SetMonitorHotkey: "設置監控開關快捷鍵",
		SetSwitchPromptHotkey: "設置切換Prompt快捷鍵",
		PressDesiredHotkey: "按下你想要的快捷鍵組合",
		
		// Console messages
		MonitorStartedTray: "✅ 監控已通過托盤啟動",
		MonitorStoppedTray: "⏸️ 監控已通過托盤停止",
		AutoPasteEnabled: "✅ 自動粘貼已啟用",
		AutoPasteDisabled: "❌ 自動粘貼已禁用",
		HotkeysLabel: "快捷鍵:",
		MonitorToggleKey: "監控開關: %s",
		SwitchStyleKey: "切換風格: %s",
		MonitorPausedByHotkey: "⏸ 監控已暫停 (通過快捷鍵)",
		MonitorResumedByHotkey: "▶ 監控已恢復 (通過快捷鍵)",
		StartingTray: "正在啟動系統托盤...",
		ControlFromTray: "請從系統托盤控制xiaoniao",
		GoodbyeEmoji: "再見！👋",
		DirectTranslation: "直譯",
		TranslateToChineseColon: "將以下內容翻譯成中文：",
		
		// API config messages
		NoModelsFound: "沒有找到模型",
		CurrentSuffix: " (當前)",
		UnrecognizedAPIKey: "無法識別API Key: %v",
		ConnectionFailed: "連接失敗 (%s): %v",
		ConnectionSuccessNoModels: "連接成功 (%s) - 無法獲取模型列表: %v",
		ConnectionSuccessWithModels: "連接成功 (%s) - %d個模型",
		TestingInProgress: "正在測試...",
		
		// System hotkey
		SystemHotkeyFormat: "系統快捷鍵: %s",
		SystemHotkeyLabel: "系統快捷鍵",
		XiaoniaoToggleMonitor: "xiaoniao 切換監控",
		XiaoniaoSwitchStyle: "xiaoniao 切換風格",
		
		// Translator error detection
		CannotProceed: "無法進行",
		AIReturnedMultiline: "AI返回了多行內容 (長度: %d)",
		UsingFirstLine: "只使用第一行: %s",
		CannotTranslate: "不能翻譯",
		UnableToTranslate: "無法翻譯",
		Sorry: "抱歉",
		
		// Theme names and descriptions
		DefaultThemeName: "默認",
		DefaultThemeDesc: "經典藍色主題",
		TokyoNightDesc: "暗色主題，靈感來自東京夜景",
		SoftPastelDesc: "柔和的粉彩主題",
		MinimalThemeName: "極簡",
		MinimalThemeDesc: "簡潔的黑白主題",
		
		// Tray messages
		StatusTranslated: "狀態: 已翻譯 %d 次",
		DefaultPrompt: "默認",
		TrayMonitoring: "xiaoniao - 監控中 | 風格: %s",
		TrayStopped: "xiaoniao - 已停止 | 風格: %s",
		StyleLabel: "風格",

		// New i18n fields for v1.0
		SingleModifier: "單個修飾鍵",
		SwitchFunction: "切換功能",
		Edit: "編輯",
		Save: "保存",
		FormatError: "格式錯誤：請使用 '修飾鍵+主鍵' 格式，如 'Ctrl+Q'",
		SingleKey: "單個按鍵",
		Back: "返回",
		InvalidModifier: "無效的修飾鍵: %s",
		InvalidMainKey: "無效的主鍵: %s",
		ProviderLabel: "提供商: ",
		CommonExamples: "常用範例",
		InputFormat: "輸入格式",
		ModifierPlusKey: "修飾鍵+主鍵",

		// Critical missing fields from main.go
		ProgramAlreadyRunning: "程式已在運行中。請檢查系統托盤圖標。\n如果看不到托盤圖標，請嘗試結束所有xiaoniao進程後重新啟動。",
		TrayManagerInitFailed: "托盤管理器初始化失敗: %v\n\n請檢查系統是否支持系統托盤功能。",
		SystemTrayStartFailed: "系統托盤啟動失敗: %v\n\n可能的原因:\n1. 系統托盤功能被禁用\n2. 權限不足\n3. 系統資源不足\n\n請檢查系統設置後重試。",
		NotConfiguredStatus: "未配置",
		PleaseConfigureAPIFirst: "請先配置API",
		APIConfigCompleted: "API配置完成，重新初始化翻譯服務...",
		MonitorStartedConsole: "監控已啟動",
		MonitorPausedConsole: "監控已暫停",
		ExportLogsFailed: "導出日誌失敗: %v",
		LogsExportedTo: "日誌已導出到: %s",
		ConfigRefreshedDetail: "配置已刷新: %s | %s | %s",
		RefreshConfigFailed: "刷新配置失敗: %v",
		SwitchedTo: "已切換到: %s",
		ConfigRefreshedAndReinit: "配置已刷新，翻譯器將重新初始化",
		MonitorPausedMsg: "監控已暫停",
		MonitorResumedMsg: "監控已恢復",
		SwitchPromptMsg: "🔄 切換提示詞: %s",
		TranslationStyle: "翻譯風格: %s",
		AutoPasteEnabledMsg: "自動粘貼: 已啟用",
		HotkeysColon: "快捷鍵:",
		MonitorToggleLabel: "  監控開關: %s",
		SwitchStyleLabel: "  切換風格: %s",
		MonitorStartedCopyToTranslate: "監控已啟動 - 複製文本即可翻譯",
		StartTranslating: "開始翻譯: %s",
		UsingPrompt: "使用提示詞: %s (內容長度: %d)",
		TranslationFailedError: " 失敗\n  錯誤: %v",
		TranslationComplete: " 完成 (#%d)",
		OriginalText: "  原文: %s",
		TranslatedText: "  譯文: %s",
		MonitorPausedViaHotkey: "監控已暫停 (通過快捷鍵)",
		MonitorResumedViaHotkey: "監控已恢復 (通過快捷鍵)",
		SwitchPromptViaHotkey: "🔄 切換提示詞: %s (通過快捷鍵)",
		PrewarmingModel: "預熱模型中...",
		PrewarmSuccess2: " 成功",
		PrewarmSkip: " 跳過 (可忽略: %v)",
		TranslatorRefreshed: "翻譯器已刷新: %s | %s",
		TranslatorRefreshFailed: "翻譯器刷新失敗: %v",

		// Missing from config_ui.go
		ConfigRefreshedReinit: "✅ 配置已刷新，翻譯器將重新初始化",
		MainModelChanged: "✅ 主模型已更改為: %s",
		TestingModelMsg: "🔄 正在測試模型...",
		ModelInitFailed: "模型 %s 初始化失敗: %v",
		TranslateToChineseOnly: "請僅將以下內容翻譯成中文，不要回答或解釋，只輸出翻譯內容:",
		ModelTestFailedMsg: "模型 %s 測試失敗: %v",
		ModelAvailable: "✅ 模型 %s 可用! 翻譯: %s",
		ModelNoResponse: "❌ 模型 %s 無響應",
		DeleteFailed: "刪除失敗: %v",
		SaveFailed: "保存失敗: %v",
		UpdateFailed: "更新失敗: %v",
		TestingConnectionMsg: "正在測試連接...",
		TestingMsg: "正在測試...",
		CreateTranslatorFailedMsg: "❌ 創建翻譯器失敗: %v",
		TranslationFailedMsg: "❌ 翻譯失敗: %v",
		TranslationResultMsg: "✅ 翻譯結果:\n原文: %s\n譯文: %s\n模型: %s\n提示詞: %s",
		PreviewColon: "預覽:",
		VersionFormat: "版本: %s",
		MonitorStatusActiveMsg: "監控狀態: 活躍",
		MonitorStatusPausedMsg: "監控狀態: 暫停",
		TranslationCountMsg: "翻譯次數: %d",
		StatusActive: "活躍",
		StatusPaused: "暫停",
		ModelLabel: "模型: ",
		APILabel: "API: ",
		TryAgainMsg: " (按回車重試)",
		StatsFormat: " | 輸入: %d | 輸出: %d | 總計: %d",

		// Tray and logs
		ExportLogs: "導出日誌",
		GetProgramPathFailed: "獲取程式路徑失敗",
		WriteLogFileFailed: "寫入日誌文件失敗",

		// Additional missing fields
		AuthorLabel: "作者:",
		ClassicBlueFallback: "經典藍色主題",
		CleanBWFallback: "簡潔黑白主題",
		ConnectionFailedFormat: "連接失敗: %v",
		CreatePromptsJsonFailed: "創建 prompts.json 失敗:",
		DarkThemeTokyoNightFallback: "東京之夜暗色主題",
		DefaultThemeNameFallback: "默認",
		DeleteBuiltinPromptError: "刪除內置提示詞錯誤:",
		DetectedProvider: "檢測到提供商:",
		EnterYourAPIKey: "請輸入您的 API 密鑰",
		HotkeySettingsTitle: "快捷鍵設置",
		HotkeysSaved: "✅ 快捷鍵已保存",
		LicenseLabel: "許可證:",
		LoadUserPromptsFailed: "加載用戶提示詞失敗:",
		MinimalThemeNameFallback: "極簡",
		ModelAvailableTranslation: "✅ %s 可用！翻譯: %s",
		ModelUnavailable: "❌ %s 不可用: %v",
		MonitorToggleHotkey: "監控開關",
		PleaseSelectModel: "請選擇一個模型",
		ProjectUrlLabel: "項目網址:",
		SelectAIModel: "選擇 AI 模型",
		SelectedBrackets: "[已選擇]",
		SoftPastelFallback: "柔和粉彩主題",
		StatusTranslatedCount: "狀態: 已翻譯 %d 次",
		Success: "成功！",
		SwitchStyleHotkey: "切換風格",
		TestingConnectionDots: "正在測試連接...",
		TestingModelFormat: "正在測試 %s...",
		TranslateToChineseProvider: "翻譯為中文",
		UnknownProviderDefault: "未知提供商（默認為 OpenAI）",
		UnsupportedOS: "不支持的操作系統: %s",
		XiaoniaoMonitoring: "xiaoniao - 監控中 | 風格: %s",
		XiaoniaoStopped: "xiaoniao - 已停止 | 風格: %s",
	}
}

// getJapanese returns Japanese translations
func getJapanese() *Translations {
	return &Translations{
		// Main interface
		Title:           "xiaoniao 設定",
		ConfigTitle:     "xiaoniao - 設定",
		APIKey:          "APIキー",
		APIConfig:       "API設定",
		TranslateStyle:  "翻訳スタイル",
		TestConnection:  "翻訳テスト",
		SaveAndExit:     "保存して終了",
		Language:        "インターフェース言語",
		ManagePrompts:   "プロンプト管理",
		Theme:           "インターフェーステーマ",
		Hotkeys:         "ホットキー設定",
		AutoPaste:       "自動貼り付け",
		
		// Status messages
		Provider:        "プロバイダー",
		Model:           "モデル",
		NotSet:          "未設定",
		Testing:         "接続テスト中...",
		TestSuccess:     "✅ 接続成功！",
		TestFailed:      "❌ 接続失敗",
		APIKeySet:       "APIキーが設定されました",
		APIKeyNotSet:    "APIキーが設定されていません",
		ChangeModel:     "モデル変更",
		Enabled:         "有効",
		Disabled:        "無効",
		
		// Help information
		HelpMove:        "↑↓ 移動",
		HelpSelect:      "Enter 選択",
		HelpBack:        "Esc 戻る",
		HelpQuit:        "Ctrl+C 終了",
		HelpTab:         "Tab 切り替え",
		HelpEdit:        "e 編集",
		HelpDelete:      "d 削除",
		HelpAdd:         "+ 追加",
		
		// Prompt management
		PromptManager:   "プロンプトマネージャー",
		AddPrompt:       "プロンプト追加",
		EditPrompt:      "プロンプト編集",
		DeletePrompt:    "プロンプト削除",
		PromptName:      "名前",
		PromptContent:   "内容",
		ConfirmDelete:   "削除を確認しますか？",
		
		// Running interface
		Running:         "実行中",
		Monitoring:      "クリップボード監視中...",
		CopyToTranslate: "テキストをコピーすると自動翻訳",
		ExitTip:         "Ctrl+C で終了",
		Translating:     "翻訳中...",
		Complete:        "完了",
		Failed:          "失敗",
		Original:        "原文",
		Translation:     "訳文",
		TotalCount:      "合計翻訳",
		Goodbye:         "さようなら！👋",
		TranslateCount:  "回",
		
		// Help documentation
		HelpTitle:       "xiaoniao",
		HelpDesc:        "AI駆動のクリップボード翻訳ツール",
		Commands:        "コマンド説明",
		RunCommand:      "xiaoniao run",
		RunDesc:         "クリップボード監視を開始し、コピーした内容を自動翻訳",
		TrayCommand:     "xiaoniao tray",
		TrayDesc:        "システムトレイモードを起動",
		ConfigCommand:   "xiaoniao config",
		ConfigDesc:      "対話型設定画面を開く",
		HelpCommand:     "xiaoniao help",
		HelpDesc2:       "このヘルプ情報を表示",
		VersionCommand:  "xiaoniao version",
		VersionDesc:     "バージョン情報を表示",
		HowItWorks:      "動作原理",
		Step1:           "xiaoniao config を実行してAPIを設定",
		Step2:           "xiaoniao run を実行して監視を開始",
		Step3:           "任意のテキストをコピー（Ctrl+C）",
		Step4:           "自動翻訳してクリップボードを置換",
		Step5:           "通知音が鳴ったら直接貼り付け（Ctrl+V）",
		Warning:         "注意: 翻訳は元のクリップボード内容を上書きします！",
		
		// Error messages
		NoAPIKey:        "❌ APIキーが設定されていません",
		RunConfigFirst:  "まず実行してください: xiaoniao config",
		AlreadyRunning:  "❌ xiaoniao はすでに実行中です",
		InitFailed:      "初期化失敗",
		ConfigNotFound:  "設定ファイルが見つかりません",
		InvalidAPIKey:   "APIキーが無効です",
		NetworkError:    "ネットワーク接続エラー",
		TranslateFailed: "翻訳失敗",
		
		// API Config
		EnterAPIKey:     "API Keyを入力してください",
		EnterNewAPIKey:  "新しいAPI Keyを入力",
		ChangeAPIKey:    "APIキー変更",
		SelectMainModel: "メインモデルを選択",
		SupportedProviders: "サポートされているプロバイダー",
		SearchModel:     "モデルを検索...",
		MainModel:       "メインモデル",
		NoPromptAvailable: "(利用可能なプロンプトなし)",
		
		// Usage messages
		Usage:           "使用方法",
		UnknownCommand:  "不明なコマンド",
		OpeningConfig:   "設定画面を開いています...",
		
		// Tray menu
		TrayShow:        "ウィンドウを表示",
		TrayHide:        "ウィンドウを隠す",
		TraySettings:    "設定",
		TrayQuit:        "終了",
		TrayToggle:      "監視切り替え",
		TrayRefresh:     "設定を更新",
		TrayAbout:       "情報",
		
		// Theme related
		SelectTheme:      "インターフェーステーマを選択",
		DefaultTheme:     "デフォルト",
		ClassicBlue:      "クラシックブルーテーマ",
		DarkTheme:        "ダークテーマ",
		
		// Hotkey related
		HotkeySettings:   "ホットキー設定",
		ToggleMonitor:    "監視切り替え",
		SwitchPromptKey:  "プロンプト切り替え",
		PressEnterToSet:  "Enterを押してホットキーを設定",
		PressDeleteToClear: "Deleteを押してホットキーをクリア",
		NotConfigured:    "(未設定)",
		
		// Test translation
		TestTranslation:  "翻訳テスト",
		CurrentConfig:    "現在の設定",
		EnterTextToTranslate: "翻訳するテキストを入力してください",
		TranslationResult: "翻訳結果",
		
		// About page
		About:            "xiaoniao について",
		Author:           "作者",
		License:          "オープンソースライセンス",
		ProjectUrl:       "プロジェクトURL",
		SupportAuthor:    "💝 作者を支援",
		PriceNote:        "製品価格は$1ですが、無料で使用できます。",
		ShareNote:        "本当に役立った場合は、コーヒーをおごるか、\nより多くの人と共有してください！:)",
		ThanksForUsing:   "xiaoniaoをご利用いただきありがとうございます！",
		BackToMainMenu:   "[Esc] メインメニューに戻る",
		ComingSoon:       "(近日オープンソース)",
		
		// Model selection
		TotalModels:      "合計 %d モデル",
		SearchModels:     "検索",
		SelectToConfirm:  "選択",
		TestModel:        "テスト",
		SearchSlash:      "検索",
		
		// Debug info
		DebugInfo:        "デバッグ情報",
		CursorPosition:   "カーソル",
		InputFocus:       "入力フォーカス",
		KeyPressed:       "キー押下",
		
		// Additional messages
		MonitorStarted:  "✅ 監視が開始されました",
		MonitorStopped:  "⏸️ 監視が停止されました",
		StopMonitor:     "監視を停止",
		StartMonitor:    "監視を開始",
		ConfigUpdated:   "✅ 設定が更新されました",
		RefreshFailed:   "❌ 設定の更新に失敗しました",
		SwitchPrompt:    "切り替え",
		PrewarmModel:    "モデルを予熱中...",
		PrewarmSuccess:  "✅",
		PrewarmFailed:   "⚠️ (無視可能: %v)",
		
		// Additional UI text
		WaitingForKeys:  "キー入力を待っています...",
		DetectedKeys:    "検出",
		HotkeyTip:       "ヒント",
		HoldModifier:    "Ctrl/Alt/Shift + 他のキーを押してください",
		DetectedAutoSave: "組み合わせキー検出後に自動保存",
		PressEscCancel:  "ESCを押してキャンセル",
		DefaultName:     "デフォルト",
		MinimalTheme:    "ミニマル",
		
		// Model selection
		ConnectionSuccess: "接続成功",
		ModelsCount:      "%d個のモデル",
		SelectModel:      "選択",
		TestingModel:     "モデル %s をテスト中...",
		ModelTestFailed:  "モデル %s のテストに失敗しました: %v",
		SearchModels2:    "検索",
		TotalModelsCount: "合計 %d 個のモデル",
		
		// Hotkey messages
		HotkeyAvailable:  "✅ 利用可能、Enterで確認",
		PressEnterConfirm: "Enterを押して確認",
		
		// Help text additions
		HelpEnterConfirm: "Enter 確認",
		HelpTabSwitch:    "Tab 切り替え",
		HelpEscReturn:    "Esc 戻る",
		HelpUpDownSelect: "↑↓ 選択",
		HelpTTest:        "T テスト",
		HelpSearchSlash:  "/ 検索",
		HelpTranslate:    "Enter: 翻訳",
		HelpCtrlSSaveExit: "Ctrl+S 保存して終了",
		HelpCtrlSSave:    "Ctrl+S 保存",
		
		// Theme descriptions
		DarkThemeTokyoNight: "東京の夜景にインスパイアされたダークテーマ",
		ChocolateTheme:      "ダークチョコレートテーマ",
		LatteTheme:          "明るいラテテーマ",
		DraculaTheme:        "ドラキュラダークテーマ",
		GruvboxDarkTheme:    "レトロダークテーマ",
		GruvboxLightTheme:   "レトロライトテーマ",
		NordTheme:           "北欧ミニマルスタイル",
		SolarizedDarkTheme:  "目に優しいダークテーマ",
		SolarizedLightTheme: "目に優しいライトテーマ",
		MinimalBWTheme:      "シンプルな白黒テーマ",
		
		// Prompt management additions
		HelpNewPrompt:    "n 新規",
		HelpEditPrompt:   "e 編集",
		HelpDeletePrompt: "d 削除",
		ConfirmDeleteKey: "dを押して削除を確認",
		CancelDelete:     "他のキーでキャンセル",
		
		// Status messages
		TestingConnection: "テスト中...",
		DetectingProvider: "検出成功",
		
		// About page additions
		ProjectAuthor: "作者",
		OpenSourceLicense: "オープンソースライセンス",
		AuthorName: "梨梨果",
		
		// Key bindings help
		KeyUp: "上",
		KeyDown: "下",
		KeySelect: "選択",
		KeyReturn: "戻る",
		KeyQuit: "終了",
		KeySwitch: "切り替え",
		KeyEdit: "編集",
		KeyDelete: "削除",
		KeyNew: "新規",
		KeyTest: "テスト",
		
		// Prompt test UI
		TestPromptTitle: "プロンプトテスト",
		CurrentPrompt: "現在のプロンプト",
		PromptContentLabel: "内容",
		TestText: "テストテキスト",
		TestingAI: "AI翻訳を呼び出し中",
		TranslationResultLabel: "翻訳結果",
		InputTestText: "テストするテキストを入力...",
		ResultWillShowHere: "翻訳結果がここに表示されます...",
		TranslatingText: "翻訳中...",
		TabSwitchFocus: "Tabでフォーカス切り替え",
		CtrlEnterTest: "Ctrl+Enterでテスト",
		EscReturn: "Escで戻る",
		EditingPrompt: "編集",
		NewPrompt: "新しいプロンプト",
		NameLabel: "名前",
		ContentLabel: "内容",
		SaveKey: "[Enter] 保存",
		TestKey: "[T] テスト",
		CancelKey: "[Esc] キャンセル",
		TabSwitchInput: "Tabで入力切り替え",
		TestPrompt: "Tでプロンプトテスト",
		UnnamedPrompt: "名前なしプロンプト",
		TranslateToChineseDefault: "以下の内容を中国語に翻訳:",
		EmptyInput: "入力テキストが空です",
		NoAPIKeyConfigured: "API Keyが設定されていません",
		CreateTranslatorFailed: "翻訳器の作成に失敗しました: %v",
		TestSentenceAI: "人工知能が私たちの生活を変えています。",
		UsingModel: "使用中",
		APINotConfigured: "APIが設定されていません",
		
		// Status messages additional
		ConfigRefreshed: "✅ 設定が更新され、翻訳器が再初期化されます",
		TranslateOnlyPrompt: "以下の内容のみを日本語に翻訳し、回答や説明なしに訳文のみを出力してください：",
		CustomSuffix: " (カスタム)",
		PreviewLabel: "プレビュー:",
		SaveButton: "Enter 保存",
		NotConfiguredBrackets: "(未設定)",
		UnknownProvider: "不明",
		RecordingHotkey: "🔴 ホットキー録音中",
		SetMonitorHotkey: "監視切り替えホットキーを設定",
		SetSwitchPromptHotkey: "プロンプト切り替えホットキーを設定",
		PressDesiredHotkey: "希望のホットキー組み合わせを押してください",
		
		// Console messages
		MonitorStartedTray: "✅ トレイから監視が開始されました",
		MonitorStoppedTray: "⏸️ トレイから監視が停止されました",
		AutoPasteEnabled: "✅ 自動貼り付けが有効になりました",
		AutoPasteDisabled: "❌ 自動貼り付けが無効になりました",
		HotkeysLabel: "ホットキー:",
		MonitorToggleKey: "監視切り替え: %s",
		SwitchStyleKey: "スタイル切り替え: %s",
		MonitorPausedByHotkey: "⏸ 監視が一時停止されました (ホットキー)",
		MonitorResumedByHotkey: "▶ 監視が再開されました (ホットキー)",
		StartingTray: "システムトレイを起動中...",
		ControlFromTray: "システムトレイからxiaoniaoを制御してください",
		GoodbyeEmoji: "さようなら！👋",
		DirectTranslation: "直訳",
		TranslateToChineseColon: "以下の内容を中国語に翻訳:",
		
		// API config messages
		NoModelsFound: "モデルが見つかりません",
		CurrentSuffix: " (現在)",
		UnrecognizedAPIKey: "API Keyを認識できません: %v",
		ConnectionFailed: "接続失敗 (%s): %v",
		ConnectionSuccessNoModels: "接続成功 (%s) - モデルリストを取得できません: %v",
		ConnectionSuccessWithModels: "接続成功 (%s) - %d個のモデル",
		TestingInProgress: "テスト中...",
		
		// System hotkey
		SystemHotkeyFormat: "システムホットキー: %s",
		SystemHotkeyLabel: "システムホットキー",
		XiaoniaoToggleMonitor: "xiaoniao 監視切り替え",
		XiaoniaoSwitchStyle: "xiaoniao スタイル切り替え",
		
		// Translator error detection
		CannotProceed: "続行できません",
		AIReturnedMultiline: "AIが複数行を返しました (長さ: %d)",
		UsingFirstLine: "最初の行のみ使用: %s",
		CannotTranslate: "翻訳できません",
		UnableToTranslate: "翻訳不可",
		Sorry: "申し訳ありません",
		
		// Theme names and descriptions
		DefaultThemeName: "デフォルト",
		DefaultThemeDesc: "クラシックブルーテーマ",
		TokyoNightDesc: "東京の夜景にインスパイアされたダークテーマ",
		SoftPastelDesc: "柔らかいパステルテーマ",
		MinimalThemeName: "ミニマル",
		MinimalThemeDesc: "シンプルな白黒テーマ",
		
		// Tray messages
		StatusTranslated: "ステータス: %d回翻訳済み",
		DefaultPrompt: "デフォルト",
		TrayMonitoring: "xiaoniao - 監視中 | スタイル: %s",
		TrayStopped: "xiaoniao - 停止中 | スタイル: %s",
		StyleLabel: "スタイル",

		// New i18n fields for v1.0
		SingleModifier: "単一修飾キー",
		Save: "保存",
		ProviderLabel: "プロバイダー: ",
		InputFormat: "入力形式",
		SingleKey: "単一キー",
		SwitchFunction: "機能切り替え",
		Edit: "編集",
		Back: "戻る",
		FormatError: "フォーマットエラー：'修飾キー+メインキー' 形式を使用してください（例：'Ctrl+Q'）",
		InvalidModifier: "無効な修飾キー: %s",
		InvalidMainKey: "無効なメインキー: %s",
		CommonExamples: "よく使う例",
		ModifierPlusKey: "修飾キー+メインキー",

		// Critical missing fields from main.go
		ProgramAlreadyRunning: "プログラムは既に実行中です。システムトレイアイコンを確認してください。\nトレイアイコンが表示されない場合は、すべてのxiaoniaoプロセスを終了してから再起動してください。",
		TrayManagerInitFailed: "トレイマネージャーの初期化に失敗しました: %v\n\nシステムがシステムトレイ機能をサポートしているか確認してください。",
		SystemTrayStartFailed: "システムトレイの起動に失敗しました: %v\n\n考えられる原因:\n1. システムトレイ機能が無効になっている\n2. 権限が不足している\n3. システムリソースが不足している\n\nシステム設定を確認してから再試行してください。",
		NotConfiguredStatus: "未設定",
		PleaseConfigureAPIFirst: "最初にAPIを設定してください",
		APIConfigCompleted: "API設定が完了しました。翻訳サービスを再初期化しています...",
		MonitorStartedConsole: "監視を開始しました",
		MonitorPausedConsole: "監視を一時停止しました",
		ExportLogsFailed: "ログのエクスポートに失敗しました: %v",
		LogsExportedTo: "ログをエクスポートしました: %s",
		ConfigRefreshedDetail: "設定を更新しました: %s | %s | %s",
		RefreshConfigFailed: "設定の更新に失敗しました: %v",
		SwitchedTo: "切り替えました: %s",
		ConfigRefreshedAndReinit: "設定が更新され、翻訳機能が再初期化されます",
		MonitorPausedMsg: "監視を一時停止しました",
		MonitorResumedMsg: "監視を再開しました",
		SwitchPromptMsg: "🔄 プロンプトを切り替え: %s",
		TranslationStyle: "翻訳スタイル: %s",
		AutoPasteEnabledMsg: "自動貼り付け: 有効",
		HotkeysColon: "ホットキー:",
		MonitorToggleLabel: "  監視の切り替え: %s",
		SwitchStyleLabel: "  スタイルの切り替え: %s",
		MonitorStartedCopyToTranslate: "監視を開始しました - テキストをコピーして翻訳",
		StartTranslating: "翻訳を開始: %s",
		UsingPrompt: "プロンプトを使用: %s (コンテンツの長さ: %d)",
		TranslationFailedError: " 失敗\n  エラー: %v",
		TranslationComplete: " 完了 (#%d)",
		OriginalText: "  原文: %s",
		TranslatedText: "  翻訳: %s",
		MonitorPausedViaHotkey: "監視を一時停止しました (ホットキー経由)",
		MonitorResumedViaHotkey: "監視を再開しました (ホットキー経由)",
		SwitchPromptViaHotkey: "🔄 プロンプトを切り替え: %s (ホットキー経由)",
		PrewarmingModel: "モデルをウォームアップ中...",
		PrewarmSuccess2: " 成功",
		PrewarmSkip: " スキップ (無視可能: %v)",
		TranslatorRefreshed: "翻訳機能を更新しました: %s | %s",
		TranslatorRefreshFailed: "翻訳機能の更新に失敗しました: %v",

		// Missing from config_ui.go
		ConfigRefreshedReinit: "✅ 設定が更新され、翻訳機能が再初期化されます",
		MainModelChanged: "✅ メインモデルを変更しました: %s",
		TestingModelMsg: "🔄 モデルをテスト中...",
		ModelInitFailed: "モデル %s の初期化に失敗しました: %v",
		TranslateToChineseOnly: "以下の内容を中国語に翻訳してください。回答や説明は不要で、翻訳のみを出力してください:",
		ModelTestFailedMsg: "モデル %s のテストに失敗しました: %v",
		ModelAvailable: "✅ モデル %s が利用可能です! 翻訳: %s",
		ModelNoResponse: "❌ モデル %s が応答しません",
		DeleteFailed: "削除に失敗しました: %v",
		SaveFailed: "保存に失敗しました: %v",
		UpdateFailed: "更新に失敗しました: %v",
		TestingConnectionMsg: "接続をテスト中...",
		TestingMsg: "テスト中...",
		CreateTranslatorFailedMsg: "❌ 翻訳機能の作成に失敗しました: %v",
		TranslationFailedMsg: "❌ 翻訳に失敗しました: %v",
		TranslationResultMsg: "✅ 翻訳結果:\n原文: %s\n翻訳: %s\nモデル: %s\nプロンプト: %s",
		PreviewColon: "プレビュー:",
		VersionFormat: "バージョン: %s",
		MonitorStatusActiveMsg: "監視ステータス: アクティブ",
		MonitorStatusPausedMsg: "監視ステータス: 一時停止",
		TranslationCountMsg: "翻訳回数: %d",
		StatusActive: "アクティブ",
		StatusPaused: "一時停止",
		ModelLabel: "モデル: ",
		APILabel: "API: ",
		TryAgainMsg: " (Enterで再試行)",
		StatsFormat: " | 入力: %d | 出力: %d | 合計: %d",

		// Additional missing fields
		AuthorLabel: "作者:",
		ClassicBlueFallback: "クラシックブルーテーマ",
		CleanBWFallback: "クリーンな白黒テーマ",
		ConnectionFailedFormat: "接続失敗: %v",
		CreatePromptsJsonFailed: "prompts.jsonの作成に失敗しました:",
		DarkThemeTokyoNightFallback: "東京の夜にインスパイアされたダークテーマ",
		DefaultThemeNameFallback: "デフォルト",
		DeleteBuiltinPromptError: "組み込みプロンプトの削除エラー:",
		DetectedProvider: "検出されたプロバイダー:",
		EnterYourAPIKey: "APIキーを入力してください",
		ExportLogs: "ログをエクスポート",
		GetProgramPathFailed: "プログラムパスの取得に失敗しました",
		HotkeySettingsTitle: "ホットキー設定",
		HotkeysSaved: "✅ ホットキーが保存されました",
		LicenseLabel: "ライセンス:",
		LoadUserPromptsFailed: "ユーザープロンプトの読み込みに失敗しました:",
		MinimalThemeNameFallback: "ミニマル",
		ModelAvailableTranslation: "✅ %s 利用可能！翻訳: %s",
		ModelUnavailable: "❌ %s 利用不可: %v",
		MonitorToggleHotkey: "監視切り替え",
		PleaseSelectModel: "モデルを選択してください",
		ProjectUrlLabel: "プロジェクトURL:",
		SelectAIModel: "AIモデルを選択",
		SelectedBrackets: "[選択済み]",
		SoftPastelFallback: "ソフトパステルテーマ",
		StatusTranslatedCount: "ステータス: %d 回翻訳済み",
		Success: "成功！",
		SwitchStyleHotkey: "スタイル切り替え",
		TestingConnectionDots: "接続をテスト中...",
		TestingModelFormat: "%s をテスト中...",
		TranslateToChineseProvider: "中国語に翻訳",
		Tutorial: "チュートリアル",
		TutorialContent: `クイックスタートガイド：

1. APIキーの設定
   • メインメニューから「API設定」を選択
   • APIキーを入力（OpenAI、Anthropicなど）
   • システムが自動的にプロバイダーを検出

2. モデルの選択
   • API設定後、「モデルを選択」を選択
   • リストからAIモデルを選択

3. ホットキーの設定（オプション）
   • メインメニューから「ホットキー設定」を選択
   • 監視の切り替えとプロンプト切り替えのホットキーを設定

4. 使用開始
   • Ctrl+X で切り取りまたは Ctrl+C でコピーすると翻訳が開始
   • プログラムが自動的にクリップボードの内容を置き換え
   • Ctrl+V で翻訳結果を貼り付け
   • 一部のアプリでは手動での貼り付けが必要

ビデオチュートリアル：
• Bilibili: （近日公開）
• YouTube: （近日公開）`,
		UnknownProviderDefault: "不明なプロバイダー（デフォルトはOpenAI）",
		UnsupportedOS: "サポートされていないOS: %s",
		WriteLogFileFailed: "ログファイルの書き込みに失敗しました",
		XiaoniaoMonitoring: "xiaoniao - 監視中 | スタイル: %s",
		XiaoniaoStopped: "xiaoniao - 停止中 | スタイル: %s",
	}
}

// getKorean returns Korean translations
func getKorean() *Translations {
	return &Translations{
		// Main interface
		Title:           "xiaoniao 설정",
		ConfigTitle:     "xiaoniao - 설정",
		APIKey:          "API 키",
		APIConfig:       "API 설정",
		TranslateStyle:  "번역 스타일",
		TestConnection:  "번역 테스트",
		SaveAndExit:     "저장하고 종료",
		Language:        "인터페이스 언어",
		ManagePrompts:   "프롬프트 관리",
		Theme:           "인터페이스 테마",
		Hotkeys:         "단축키 설정",
		AutoPaste:       "자동 붙여넣기",
		
		// Status messages
		Provider:        "공급자",
		Model:           "모델",
		NotSet:          "설정 안 됨",
		Testing:         "연결 테스트 중...",
		TestSuccess:     "✅ 연결 성공!",
		TestFailed:      "❌ 연결 실패",
		APIKeySet:       "API 키가 설정됨",
		APIKeyNotSet:    "API 키가 설정되지 않음",
		ChangeModel:     "모델 변경",
		Enabled:         "활성화",
		Disabled:        "비활성화",
		
		// Help information
		HelpMove:        "↑↓ 이동",
		HelpSelect:      "Enter 선택",
		HelpBack:        "Esc 뒤로",
		HelpQuit:        "Ctrl+C 종료",
		HelpTab:         "Tab 전환",
		HelpEdit:        "e 편집",
		HelpDelete:      "d 삭제",
		HelpAdd:         "+ 추가",
		
		// Prompt management
		PromptManager:   "프롬프트 관리자",
		AddPrompt:       "프롬프트 추가",
		EditPrompt:      "프롬프트 편집",
		DeletePrompt:    "프롬프트 삭제",
		PromptName:      "이름",
		PromptContent:   "내용",
		ConfirmDelete:   "삭제하시겠습니까?",
		
		// Running interface
		Running:         "실행 중",
		Monitoring:      "클립보드 모니터링 중...",
		CopyToTranslate: "텍스트를 복사하면 자동 번역됩니다",
		ExitTip:         "Ctrl+C로 종료",
		Translating:     "번역 중...",
		Complete:        "완료",
		Failed:          "실패",
		Original:        "원문",
		Translation:     "번역",
		TotalCount:      "총 번역",
		Goodbye:         "안녕히 가세요! 👋",
		TranslateCount:  "회",
		
		// Help documentation
		HelpTitle:       "xiaoniao",
		HelpDesc:        "AI 기반 클립보드 번역 도구",
		Commands:        "명령어 설명",
		RunCommand:      "xiaoniao run",
		RunDesc:         "클립보드 모니터링을 시작하고 복사한 내용을 자동 번역",
		TrayCommand:     "xiaoniao tray",
		TrayDesc:        "시스템 트레이 모드 시작",
		ConfigCommand:   "xiaoniao config",
		ConfigDesc:      "대화형 설정 화면 열기",
		HelpCommand:     "xiaoniao help",
		HelpDesc2:       "이 도움말 정보 표시",
		VersionCommand:  "xiaoniao version",
		VersionDesc:     "버전 정보 표시",
		HowItWorks:      "작동 방식",
		Step1:           "xiaoniao config 실행하여 API 설정",
		Step2:           "xiaoniao run 실행하여 모니터링 시작",
		Step3:           "아무 텍스트나 복사 (Ctrl+C)",
		Step4:           "자동으로 번역되어 클립보드 교체",
		Step5:           "알림음이 들리면 바로 붙여넣기 (Ctrl+V)",
		Warning:         "주의: 번역이 원본 클립보드 내용을 덮어씁니다!",
		
		// Error messages
		NoAPIKey:        "❌ API 키가 설정되지 않음",
		RunConfigFirst:  "먼저 실행하세요: xiaoniao config",
		AlreadyRunning:  "❌ xiaoniao가 이미 실행 중입니다",
		InitFailed:      "초기화 실패",
		ConfigNotFound:  "설정 파일을 찾을 수 없음",
		InvalidAPIKey:   "잘못된 API 키",
		NetworkError:    "네트워크 연결 오류",
		TranslateFailed: "번역 실패",
		
		// API Config
		EnterAPIKey:     "API Key를 입력하세요",
		EnterNewAPIKey:  "새 API Key 입력",
		ChangeAPIKey:    "API 키 변경",
		SelectMainModel: "메인 모델 선택",
		SupportedProviders: "지원되는 공급자",
		SearchModel:     "모델 검색...",
		MainModel:       "메인 모델",
		NoPromptAvailable: "(사용 가능한 프롬프트 없음)",
		
		// Usage messages
		Usage:           "사용법",
		UnknownCommand:  "알 수 없는 명령",
		OpeningConfig:   "설정 화면 열기 중...",
		
		// Tray menu
		TrayShow:        "창 표시",
		TrayHide:        "창 숨기기",
		TraySettings:    "설정",
		TrayQuit:        "종료",
		TrayToggle:      "모니터링 토글",
		TrayRefresh:     "설정 새로고침",
		TrayAbout:       "정보",
		
		// Theme related
		SelectTheme:      "인터페이스 테마 선택",
		DefaultTheme:     "기본",
		ClassicBlue:      "클래식 블루 테마",
		DarkTheme:        "다크 테마",
		
		// Hotkey related
		HotkeySettings:   "단축키 설정",
		ToggleMonitor:    "모니터링 토글",
		SwitchPromptKey:  "프롬프트 전환",
		PressEnterToSet:  "Enter를 눌러 단축키 설정",
		PressDeleteToClear: "Delete를 눌러 단축키 삭제",
		NotConfigured:    "(설정 안 됨)",
		
		// Test translation
		TestTranslation:  "번역 테스트",
		CurrentConfig:    "현재 설정",
		EnterTextToTranslate: "번역할 텍스트를 입력하세요",
		TranslationResult: "번역 결과",
		
		// About page
		About:            "xiaoniao 정보",
		Author:           "작성자",
		License:          "오픈소스 라이선스",
		ProjectUrl:       "프로젝트 주소",
		SupportAuthor:    "💝 작성자 지원",
		PriceNote:        "제품 가격은 $1이지만 무료로 사용할 수 있습니다.",
		ShareNote:        "정말 도움이 되었다면 커피 한 잔 사주시거나\n더 많은 사람과 공유해 주세요! :)",
		ThanksForUsing:   "xiaoniao를 사용해 주셔서 감사합니다!",
		BackToMainMenu:   "[Esc] 메인 메뉴로 돌아가기",
		ComingSoon:       "(곧 오픈소스)",
		
		// Model selection
		TotalModels:      "총 %d개 모델",
		SearchModels:     "검색",
		SelectToConfirm:  "선택",
		TestModel:        "테스트",
		SearchSlash:      "검색",
		
		// Debug info
		DebugInfo:        "디버그 정보",
		CursorPosition:   "커서",
		InputFocus:       "입력창 포커스",
		KeyPressed:       "키 입력",
		
		// Additional messages
		MonitorStarted:  "✅ 모니터링 시작됨",
		MonitorStopped:  "⏸️ 모니터링 중지됨",
		StopMonitor:     "모니터링 중지",
		StartMonitor:    "모니터링 시작",
		ConfigUpdated:   "✅ 설정이 업데이트됨",
		RefreshFailed:   "❌ 설정 새로고침 실패",
		SwitchPrompt:    "전환됨",
		PrewarmModel:    "모델 예열 중...",
		PrewarmSuccess:  "✅",
		PrewarmFailed:   "⚠️ (무시 가능: %v)",
		
		// Additional UI text
		WaitingForKeys:  "키 입력 대기 중...",
		DetectedKeys:    "감지됨",
		HotkeyTip:       "팁",
		HoldModifier:    "Ctrl/Alt/Shift + 다른 키를 누르세요",
		DetectedAutoSave: "조합키 감지 후 자동 저장",
		PressEscCancel:  "ESC를 눌러 취소",
		DefaultName:     "기본",
		MinimalTheme:    "미니멀",
		
		// Model selection
		ConnectionSuccess: "연결 성공",
		ModelsCount:      "%d개 모델",
		SelectModel:      "선택",
		TestingModel:     "%s 모델 테스트 중...",
		ModelTestFailed:  "%s 모델 테스트 실패: %v",
		SearchModels2:    "검색",
		TotalModelsCount: "총 %d개 모델",
		
		// Hotkey messages
		HotkeyAvailable:  "✅ 사용 가능, Enter로 확인",
		PressEnterConfirm: "Enter를 눌러 확인",
		
		// Help text additions
		HelpEnterConfirm: "Enter 확인",
		HelpTabSwitch:    "Tab 전환",
		HelpEscReturn:    "Esc 돌아가기",
		HelpUpDownSelect: "↑↓ 선택",
		HelpTTest:        "T 테스트",
		HelpSearchSlash:  "/ 검색",
		HelpTranslate:    "Enter: 번역",
		HelpCtrlSSaveExit: "Ctrl+S 저장 및 종료",
		HelpCtrlSSave:    "Ctrl+S 저장",
		
		// Theme descriptions
		DarkThemeTokyoNight: "도쿄 야경에서 영감을 받은 다크 테마",
		ChocolateTheme:      "다크 초콜릿 테마",
		LatteTheme:          "밝은 라떼 테마",
		DraculaTheme:        "드라큘라 다크 테마",
		GruvboxDarkTheme:    "레트로 다크 테마",
		GruvboxLightTheme:   "레트로 라이트 테마",
		NordTheme:           "북유럽 미니멀 스타일",
		SolarizedDarkTheme:  "눈 보호 다크 테마",
		SolarizedLightTheme: "눈 보호 라이트 테마",
		MinimalBWTheme:      "깔끔한 흑백 테마",
		
		// Prompt management additions
		HelpNewPrompt:    "n 새로 만들기",
		HelpEditPrompt:   "e 편집",
		HelpDeletePrompt: "d 삭제",
		ConfirmDeleteKey: "d를 눌러 삭제 확인",
		CancelDelete:     "다른 키를 눌러 취소",
		
		// Status messages
		TestingConnection: "테스트 중...",
		DetectingProvider: "감지 성공",
		
		// About page additions
		ProjectAuthor: "작성자",
		OpenSourceLicense: "오픈소스 라이선스",
		AuthorName: "梨梨果",
		
		// Key bindings help
		KeyUp: "위",
		KeyDown: "아래",
		KeySelect: "선택",
		KeyReturn: "돌아가기",
		KeyQuit: "종료",
		KeySwitch: "전환",
		KeyEdit: "편집",
		KeyDelete: "삭제",
		KeyNew: "새로 만들기",
		KeyTest: "테스트",
		
		// Prompt test UI
		TestPromptTitle: "프롬프트 테스트",
		CurrentPrompt: "현재 프롬프트",
		PromptContentLabel: "내용",
		TestText: "테스트 텍스트",
		TestingAI: "AI 번역 호출 중",
		TranslationResultLabel: "번역 결과",
		InputTestText: "테스트할 텍스트 입력...",
		ResultWillShowHere: "번역 결과가 여기에 표시됩니다...",
		TranslatingText: "번역 중...",
		TabSwitchFocus: "Tab으로 포커스 전환",
		CtrlEnterTest: "Ctrl+Enter로 테스트",
		EscReturn: "Esc로 돌아가기",
		EditingPrompt: "편집",
		NewPrompt: "새 프롬프트",
		NameLabel: "이름",
		ContentLabel: "내용",
		SaveKey: "[Enter] 저장",
		TestKey: "[T] 테스트",
		CancelKey: "[Esc] 취소",
		TabSwitchInput: "Tab으로 입력창 전환",
		TestPrompt: "T로 프롬프트 테스트",
		UnnamedPrompt: "이름 없는 프롬프트",
		TranslateToChineseDefault: "다음 내용을 중국어로 번역:",
		EmptyInput: "입력 텍스트가 비어있음",
		NoAPIKeyConfigured: "API Key가 설정되지 않음",
		CreateTranslatorFailed: "번역기 생성 실패: %v",
		TestSentenceAI: "인공지능이 우리의 생활 방식을 바꾸고 있습니다.",
		UsingModel: "사용 중",
		APINotConfigured: "API가 설정되지 않음",
		
		// Status messages additional
		ConfigRefreshed: "✅ 설정이 새로고침되고 번역기가 재초기화됩니다",
		TranslateOnlyPrompt: "다음 내용만 한국어로 번역하고, 답변이나 설명 없이 번역문만 출력하세요:",
		CustomSuffix: " (사용자 정의)",
		PreviewLabel: "미리보기:",
		SaveButton: "Enter 저장",
		NotConfiguredBrackets: "(설정 안 됨)",
		UnknownProvider: "알 수 없음",
		RecordingHotkey: "🔴 단축키 녹화 중",
		SetMonitorHotkey: "모니터링 토글 단축키 설정",
		SetSwitchPromptHotkey: "프롬프트 전환 단축키 설정",
		PressDesiredHotkey: "원하는 단축키 조합을 누르세요",
		
		// Console messages
		MonitorStartedTray: "✅ 트레이에서 모니터링 시작됨",
		MonitorStoppedTray: "⏸️ 트레이에서 모니터링 중지됨",
		AutoPasteEnabled: "✅ 자동 붙여넣기 활성화됨",
		AutoPasteDisabled: "❌ 자동 붙여넣기 비활성화됨",
		HotkeysLabel: "단축키:",
		MonitorToggleKey: "모니터링 토글: %s",
		SwitchStyleKey: "스타일 전환: %s",
		MonitorPausedByHotkey: "⏸ 모니터링 일시정지됨 (단축키)",
		MonitorResumedByHotkey: "▶ 모니터링 재개됨 (단축키)",
		StartingTray: "시스템 트레이 시작 중...",
		ControlFromTray: "시스템 트레이에서 xiaoniao를 제어하세요",
		GoodbyeEmoji: "안녕히 가세요! 👋",
		DirectTranslation: "직역",
		TranslateToChineseColon: "다음 내용을 중국어로 번역:",
		
		// API config messages
		NoModelsFound: "모델을 찾을 수 없음",
		CurrentSuffix: " (현재)",
		UnrecognizedAPIKey: "API Key를 인식할 수 없음: %v",
		ConnectionFailed: "연결 실패 (%s): %v",
		ConnectionSuccessNoModels: "연결 성공 (%s) - 모델 목록을 가져올 수 없음: %v",
		ConnectionSuccessWithModels: "연결 성공 (%s) - %d개 모델",
		TestingInProgress: "테스트 중...",
		
		// System hotkey
		SystemHotkeyFormat: "시스템 단축키: %s",
		SystemHotkeyLabel: "시스템 단축키",
		XiaoniaoToggleMonitor: "xiaoniao 모니터링 토글",
		XiaoniaoSwitchStyle: "xiaoniao 스타일 전환",
		
		// Translator error detection
		CannotProceed: "진행할 수 없음",
		AIReturnedMultiline: "AI가 여러 줄을 반환함 (길이: %d)",
		UsingFirstLine: "첫 번째 줄만 사용: %s",
		CannotTranslate: "번역할 수 없음",
		UnableToTranslate: "번역 불가",
		Sorry: "죄송합니다",
		
		// Theme names and descriptions
		DefaultThemeName: "기본",
		DefaultThemeDesc: "클래식 블루 테마",
		TokyoNightDesc: "도쿄 야경에서 영감을 받은 다크 테마",
		SoftPastelDesc: "부드러운 파스텔 테마",
		MinimalThemeName: "미니멀",
		MinimalThemeDesc: "깔끔한 흑백 테마",
		
		// Tray messages
		StatusTranslated: "상태: %d회 번역됨",
		DefaultPrompt: "기본",
		TrayMonitoring: "xiaoniao - 모니터링 중 | 스타일: %s",
		TrayStopped: "xiaoniao - 중지됨 | 스타일: %s",
		StyleLabel: "스타일",

		// New i18n fields for v1.0
		Back: "뒤로",
		FormatError: "형식 오류: '수정자+키' 형식을 사용하세요 (예: 'Ctrl+Q')",
		ProviderLabel: "제공자: ",
		CommonExamples: "자주 사용하는 예",
		InputFormat: "입력 형식",
		ModifierPlusKey: "수정자+메인 키",
		SingleModifier: "단일 수정자",
		Edit: "편집",
		Save: "저장",
		InvalidModifier: "잘못된 수정자: %s",
		InvalidMainKey: "잘못된 메인 키: %s",
		SingleKey: "단일 키",
		SwitchFunction: "기능 전환",

		// Critical missing fields from main.go
		ProgramAlreadyRunning: "프로그램이 이미 실행 중입니다. 시스템 트레이 아이콘을 확인하세요.\n트레이 아이콘이 보이지 않으면 모든 xiaoniao 프로세스를 종료한 후 다시 시작하세요.",
		TrayManagerInitFailed: "트레이 관리자 초기화 실패: %v\n\n시스템이 시스템 트레이 기능을 지원하는지 확인하세요.",
		SystemTrayStartFailed: "시스템 트레이 시작 실패: %v\n\n가능한 원인:\n1. 시스템 트레이 기능이 비활성화됨\n2. 권한 부족\n3. 시스템 리소스 부족\n\n시스템 설정을 확인한 후 다시 시도하세요.",
		NotConfiguredStatus: "구성되지 않음",
		PleaseConfigureAPIFirst: "먼저 API를 구성하세요",
		APIConfigCompleted: "API 구성 완료, 번역 서비스 재초기화 중...",
		MonitorStartedConsole: "모니터링 시작됨",
		MonitorPausedConsole: "모니터링 일시 중지됨",
		ExportLogsFailed: "로그 내보내기 실패: %v",
		LogsExportedTo: "로그가 내보내졌습니다: %s",
		ConfigRefreshedDetail: "구성이 새로 고쳐졌습니다: %s | %s | %s",
		RefreshConfigFailed: "구성 새로 고침 실패: %v",
		SwitchedTo: "전환됨: %s",
		ConfigRefreshedAndReinit: "구성이 새로 고쳐졌으며 번역기가 다시 초기화됩니다",
		MonitorPausedMsg: "모니터링 일시 중지됨",
		MonitorResumedMsg: "모니터링 재개됨",
		SwitchPromptMsg: "🔄 프롬프트 전환: %s",
		TranslationStyle: "번역 스타일: %s",
		AutoPasteEnabledMsg: "자동 붙여넣기: 활성화됨",
		HotkeysColon: "단축키:",
		MonitorToggleLabel: "  모니터링 전환: %s",
		SwitchStyleLabel: "  스타일 전환: %s",
		MonitorStartedCopyToTranslate: "모니터링 시작됨 - 텍스트를 복사하여 번역",
		StartTranslating: "번역 시작: %s",
		UsingPrompt: "프롬프트 사용: %s (콘텐츠 길이: %d)",
		TranslationFailedError: " 실패\n  오류: %v",
		TranslationComplete: " 완료 (#%d)",
		OriginalText: "  원문: %s",
		TranslatedText: "  번역: %s",
		MonitorPausedViaHotkey: "모니터링 일시 중지됨 (단축키를 통해)",
		MonitorResumedViaHotkey: "모니터링 재개됨 (단축키를 통해)",
		SwitchPromptViaHotkey: "🔄 프롬프트 전환: %s (단축키를 통해)",
		PrewarmingModel: "모델 예열 중...",
		PrewarmSuccess2: " 성공",
		PrewarmSkip: " 건너뛰기 (무시 가능: %v)",
		TranslatorRefreshed: "번역기가 새로 고쳐졌습니다: %s | %s",
		TranslatorRefreshFailed: "번역기 새로 고침 실패: %v",

		// Missing from config_ui.go
		ConfigRefreshedReinit: "✅ 구성이 새로 고쳐졌으며 번역기가 다시 초기화됩니다",
		MainModelChanged: "✅ 기본 모델이 변경되었습니다: %s",
		TestingModelMsg: "🔄 모델 테스트 중...",
		ModelInitFailed: "모델 %s 초기화 실패: %v",
		TranslateToChineseOnly: "다음 내용을 중국어로만 번역하세요. 답변이나 설명 없이 번역만 출력하세요:",
		ModelTestFailedMsg: "모델 %s 테스트 실패: %v",
		ModelAvailable: "✅ 모델 %s 사용 가능! 번역: %s",
		ModelNoResponse: "❌ 모델 %s 응답 없음",
		DeleteFailed: "삭제 실패: %v",
		SaveFailed: "저장 실패: %v",
		UpdateFailed: "업데이트 실패: %v",
		TestingConnectionMsg: "연결 테스트 중...",
		TestingMsg: "테스트 중...",
		CreateTranslatorFailedMsg: "❌ 번역기 생성 실패: %v",
		TranslationFailedMsg: "❌ 번역 실패: %v",
		TranslationResultMsg: "✅ 번역 결과:\n원문: %s\n번역: %s\n모델: %s\n프롬프트: %s",
		PreviewColon: "미리보기:",
		VersionFormat: "버전: %s",
		MonitorStatusActiveMsg: "모니터링 상태: 활성",
		MonitorStatusPausedMsg: "모니터링 상태: 일시 중지",
		TranslationCountMsg: "번역 횟수: %d",
		StatusActive: "활성",
		StatusPaused: "일시 중지",
		ModelLabel: "모델: ",
		APILabel: "API: ",
		TryAgainMsg: " (Enter로 재시도)",
		StatsFormat: " | 입력: %d | 출력: %d | 총계: %d",

		// Additional missing fields
		AuthorLabel: "작성자:",
		ClassicBlueFallback: "클래식 블루 테마",
		CleanBWFallback: "깔끔한 흑백 테마",
		ConnectionFailedFormat: "연결 실패: %v",
		CreatePromptsJsonFailed: "prompts.json 생성 실패:",
		DarkThemeTokyoNightFallback: "도쿄 나이트 다크 테마",
		DefaultThemeNameFallback: "기본",
		DeleteBuiltinPromptError: "내장 프롬프트 삭제 오류:",
		DetectedProvider: "감지된 제공자:",
		EnterYourAPIKey: "API 키를 입력하세요",
		ExportLogs: "로그 내보내기",
		GetProgramPathFailed: "프로그램 경로 가져오기 실패",
		HotkeySettingsTitle: "단축키 설정",
		HotkeysSaved: "✅ 단축키가 저장되었습니다",
		LicenseLabel: "라이선스:",
		LoadUserPromptsFailed: "사용자 프롬프트 로드 실패:",
		MinimalThemeNameFallback: "미니멀",
		ModelAvailableTranslation: "✅ %s 사용 가능! 번역: %s",
		ModelUnavailable: "❌ %s 사용 불가: %v",
		MonitorToggleHotkey: "모니터 토글",
		PleaseSelectModel: "모델을 선택하세요",
		ProjectUrlLabel: "프로젝트 URL:",
		SelectAIModel: "AI 모델 선택",
		SelectedBrackets: "[선택됨]",
		SoftPastelFallback: "소프트 파스텔 테마",
		StatusTranslatedCount: "상태: %d 번 번역됨",
		Success: "성공!",
		SwitchStyleHotkey: "스타일 전환",
		TestingConnectionDots: "연결 테스트 중...",
		TestingModelFormat: "%s 테스트 중...",
		TranslateToChineseProvider: "중국어로 번역",
		Tutorial: "튜토리얼",
		TutorialContent: `빠른 시작 가이드:

1. API 키 설정
   • 메인 메뉴에서 "API 설정" 선택
   • API 키 입력 (OpenAI, Anthropic 등)
   • 시스템이 자동으로 제공자 감지

2. 모델 선택
   • API 설정 후 "모델 선택" 선택
   • 목록에서 AI 모델 선택

3. 단축키 설정 (선택사항)
   • 메인 메뉴에서 "단축키 설정" 선택
   • 모니터 토글 및 프롬프트 전환 단축키 설정

4. 사용 시작
   • Ctrl+X로 잘라내기 또는 Ctrl+C로 복사 시 번역 시작
   • 프로그램이 자동으로 클립보드 내용 교체
   • Ctrl+V로 번역 결과 붙여넣기
   • 일부 앱에서는 수동 붙여넣기 필요

비디오 튜토리얼:
• Bilibili: (곧 출시)
• YouTube: (곧 출시)`,
		UnknownProviderDefault: "알 수 없는 제공자 (기본값: OpenAI)",
		UnsupportedOS: "지원되지 않는 OS: %s",
		WriteLogFileFailed: "로그 파일 쓰기 실패",
		XiaoniaoMonitoring: "xiaoniao - 모니터링 중 | 스타일: %s",
		XiaoniaoStopped: "xiaoniao - 중지됨 | 스타일: %s",
	}
}

// getSpanish returns Spanish translations
func getSpanish() *Translations {
	return &Translations{
		// Main interface
		Title:           "Configuración de xiaoniao",
		ConfigTitle:     "xiaoniao - Ajustes",
		APIKey:          "Clave API",
		APIConfig:       "Configuración API",
		TranslateStyle:  "Estilo de traducción",
		TestConnection:  "Probar traducción",
		SaveAndExit:     "Guardar y salir",
		Language:        "Idioma de interfaz",
		ManagePrompts:   "Gestionar prompts",
		Theme:           "Tema de interfaz",
		Hotkeys:         "Atajos de teclado",
		AutoPaste:       "Pegado automático",
		
		// Status messages
		Provider:        "Proveedor",
		Model:           "Modelo",
		NotSet:          "No configurado",
		Testing:         "Probando conexión...",
		TestSuccess:     "✅ ¡Conexión exitosa!",
		TestFailed:      "❌ Conexión fallida",
		APIKeySet:       "Clave API configurada",
		APIKeyNotSet:    "Clave API no configurada",
		ChangeModel:     "Cambiar modelo",
		Enabled:         "Habilitado",
		Disabled:        "Deshabilitado",
		
		// Help information
		HelpMove:        "↑↓ Mover",
		HelpSelect:      "Enter Seleccionar",
		HelpBack:        "Esc Volver",
		HelpQuit:        "Ctrl+C Salir",
		HelpTab:         "Tab Cambiar",
		HelpEdit:        "e Editar",
		HelpDelete:      "d Eliminar",
		HelpAdd:         "+ Añadir",
		
		// Prompt management
		PromptManager:   "Gestor de prompts",
		AddPrompt:       "Añadir prompt",
		EditPrompt:      "Editar prompt",
		DeletePrompt:    "Eliminar prompt",
		PromptName:      "Nombre",
		PromptContent:   "Contenido",
		ConfirmDelete:   "¿Confirmar eliminación?",
		
		// Running interface
		Running:         "En ejecución",
		Monitoring:      "Monitoreando portapapeles...",
		CopyToTranslate: "Copia cualquier texto para traducir automáticamente",
		ExitTip:         "Presiona Ctrl+C para salir",
		Translating:     "Traduciendo...",
		Complete:        "Completado",
		Failed:          "Fallido",
		Original:        "Original",
		Translation:     "Traducción",
		TotalCount:      "Total traducido",
		Goodbye:         "¡Adiós! 👋",
		TranslateCount:  "veces",
		
		// Help documentation
		HelpTitle:       "xiaoniao",
		HelpDesc:        "Herramienta de traducción de portapapeles con IA",
		Commands:        "Descripción de comandos",
		RunCommand:      "xiaoniao run",
		RunDesc:         "Iniciar monitoreo del portapapeles y traducir automáticamente el contenido copiado",
		TrayCommand:     "xiaoniao tray",
		TrayDesc:        "Iniciar modo de bandeja del sistema",
		ConfigCommand:   "xiaoniao config",
		ConfigDesc:      "Abrir interfaz de configuración interactiva",
		HelpCommand:     "xiaoniao help",
		HelpDesc2:       "Mostrar esta información de ayuda",
		VersionCommand:  "xiaoniao version",
		VersionDesc:     "Mostrar información de versión",
		HowItWorks:      "Cómo funciona",
		Step1:           "Ejecuta xiaoniao config para configurar API",
		Step2:           "Ejecuta xiaoniao run para iniciar monitoreo",
		Step3:           "Copia cualquier texto (Ctrl+C)",
		Step4:           "Se traduce automáticamente y reemplaza el portapapeles",
		Step5:           "Cuando escuches el sonido, pega directamente (Ctrl+V)",
		Warning:         "Atención: ¡La traducción sobrescribirá el contenido original del portapapeles!",
		
		// Error messages
		NoAPIKey:        "❌ Clave API no configurada",
		RunConfigFirst:  "Por favor ejecuta primero: xiaoniao config",
		AlreadyRunning:  "❌ xiaoniao ya está en ejecución",
		InitFailed:      "Fallo de inicialización",
		ConfigNotFound:  "Archivo de configuración no encontrado",
		InvalidAPIKey:   "Clave API inválida",
		NetworkError:    "Error de conexión de red",
		TranslateFailed: "Traducción fallida",
		
		// API Config
		EnterAPIKey:     "Por favor ingresa la clave API",
		EnterNewAPIKey:  "Ingresa nueva clave API",
		ChangeAPIKey:    "Cambiar clave API",
		SelectMainModel: "Seleccionar modelo principal",
		SupportedProviders: "Proveedores soportados",
		SearchModel:     "Buscar modelo...",
		MainModel:       "Modelo principal",
		NoPromptAvailable: "(Sin prompts disponibles)",
		
		// Usage messages
		Usage:           "Uso",
		UnknownCommand:  "Comando desconocido",
		OpeningConfig:   "Abriendo interfaz de configuración...",
		
		// Tray menu
		TrayShow:        "Mostrar ventana",
		TrayHide:        "Ocultar ventana",
		TraySettings:    "Configuración",
		TrayQuit:        "Salir",
		TrayToggle:      "Alternar monitoreo",
		TrayRefresh:     "Actualizar configuración",
		TrayAbout:       "Acerca de",
		
		// Theme related
		SelectTheme:      "Seleccionar tema de interfaz",
		DefaultTheme:     "Predeterminado",
		ClassicBlue:      "Tema azul clásico",
		DarkTheme:        "Tema oscuro",
		
		// Hotkey related
		HotkeySettings:   "Configuración de atajos",
		ToggleMonitor:    "Alternar monitoreo",
		SwitchPromptKey:  "Cambiar prompt",
		PressEnterToSet:  "Presiona Enter para configurar atajo",
		PressDeleteToClear: "Presiona Delete para borrar atajo",
		NotConfigured:    "(No configurado)",
		
		// Test translation
		TestTranslation:  "Probar traducción",
		CurrentConfig:    "Configuración actual",
		EnterTextToTranslate: "Ingresa el texto a traducir",
		TranslationResult: "Resultado de traducción",
		
		// About page
		About:            "Acerca de xiaoniao",
		Author:           "Autor",
		License:          "Licencia de código abierto",
		ProjectUrl:       "URL del proyecto",
		SupportAuthor:    "💝 Apoyar al autor",
		PriceNote:        "El precio del producto es $1, pero puedes usarlo gratis.",
		ShareNote:        "Si realmente te ayudó, invítame un café\no compártelo con más personas! :)",
		ThanksForUsing:   "¡Gracias por usar xiaoniao!",
		BackToMainMenu:   "[Esc] Volver al menú principal",
		ComingSoon:       "(Próximamente código abierto)",
		
		// Model selection
		TotalModels:      "Total %d modelos",
		SearchModels:     "Buscar",
		SelectToConfirm:  "Seleccionar",
		TestModel:        "Probar",
		SearchSlash:      "Buscar",
		
		// Debug info
		DebugInfo:        "Información de depuración",
		CursorPosition:   "Cursor",
		InputFocus:       "Foco de entrada",
		KeyPressed:       "Tecla presionada",
		
		// Additional messages
		MonitorStarted:  "✅ Monitoreo iniciado",
		MonitorStopped:  "⏸️ Monitoreo detenido",
		StopMonitor:     "Detener monitoreo",
		StartMonitor:    "Iniciar monitoreo",
		ConfigUpdated:   "✅ Configuración actualizada",
		RefreshFailed:   "❌ Fallo al actualizar configuración",
		SwitchPrompt:    "Cambiado a",
		PrewarmModel:    "Precalentando modelo...",
		PrewarmSuccess:  "✅",
		PrewarmFailed:   "⚠️ (ignorable: %v)",
		
		// Additional UI text
		WaitingForKeys:  "Esperando teclas...",
		DetectedKeys:    "Detectado",
		HotkeyTip:       "Consejo",
		HoldModifier:    "Mantén Ctrl/Alt/Shift + otra tecla",
		DetectedAutoSave: "Auto-guardar tras detectar combinación",
		PressEscCancel:  "Presiona ESC para cancelar",
		DefaultName:     "Predeterminado",
		MinimalTheme:    "Minimalista",
		
		// Model selection
		ConnectionSuccess: "Conexión exitosa",
		ModelsCount:      "%d modelos",
		SelectModel:      "Seleccionar",
		TestingModel:     "Probando modelo %s...",
		ModelTestFailed:  "Fallo al probar modelo %s: %v",
		SearchModels2:    "Buscar",
		TotalModelsCount: "Total %d modelos",
		
		// Hotkey messages
		HotkeyAvailable:  "✅ Disponible, presiona Enter para confirmar",
		PressEnterConfirm: "Presiona Enter para confirmar",
		
		// Help text additions
		HelpEnterConfirm: "Enter Confirmar",
		HelpTabSwitch:    "Tab Cambiar",
		HelpEscReturn:    "Esc Volver",
		HelpUpDownSelect: "↑↓ Seleccionar",
		HelpTTest:        "T Probar",
		HelpSearchSlash:  "/ Buscar",
		HelpTranslate:    "Enter: Traducir",
		HelpCtrlSSaveExit: "Ctrl+S Guardar y salir",
		HelpCtrlSSave:    "Ctrl+S Guardar",
		
		// Theme descriptions
		DarkThemeTokyoNight: "Tema oscuro inspirado en el paisaje nocturno de Tokio",
		ChocolateTheme:      "Tema chocolate oscuro",
		LatteTheme:          "Tema latte brillante",
		DraculaTheme:        "Tema Drácula oscuro",
		GruvboxDarkTheme:    "Tema retro oscuro",
		GruvboxLightTheme:   "Tema retro claro",
		NordTheme:           "Estilo minimalista nórdico",
		SolarizedDarkTheme:  "Tema oscuro que protege la vista",
		SolarizedLightTheme: "Tema claro que protege la vista",
		MinimalBWTheme:      "Tema blanco y negro simple",
		
		// Prompt management additions
		HelpNewPrompt:    "n Nuevo",
		HelpEditPrompt:   "e Editar",
		HelpDeletePrompt: "d Eliminar",
		ConfirmDeleteKey: "Presiona d para confirmar eliminación",
		CancelDelete:     "Presiona otra tecla para cancelar",
		
		// Status messages
		TestingConnection: "Probando...",
		DetectingProvider: "Detección exitosa",
		
		// About page additions
		ProjectAuthor: "Autor",
		OpenSourceLicense: "Licencia de código abierto",
		AuthorName: "梨梨果",
		
		// Key bindings help
		KeyUp: "Arriba",
		KeyDown: "Abajo",
		KeySelect: "Seleccionar",
		KeyReturn: "Volver",
		KeyQuit: "Salir",
		KeySwitch: "Cambiar",
		KeyEdit: "Editar",
		KeyDelete: "Eliminar",
		KeyNew: "Nuevo",
		KeyTest: "Probar",
		
		// Prompt test UI
		TestPromptTitle: "Probar Prompt",
		CurrentPrompt: "Prompt Actual",
		PromptContentLabel: "Contenido",
		TestText: "Texto de prueba",
		TestingAI: "Llamando traducción IA",
		TranslationResultLabel: "Resultado de traducción",
		InputTestText: "Ingresa texto para probar...",
		ResultWillShowHere: "El resultado de traducción aparecerá aquí...",
		TranslatingText: "Traduciendo...",
		TabSwitchFocus: "Tab para cambiar foco",
		CtrlEnterTest: "Ctrl+Enter para probar",
		EscReturn: "Esc para volver",
		EditingPrompt: "Editando",
		NewPrompt: "Nuevo Prompt",
		NameLabel: "Nombre",
		ContentLabel: "Contenido",
		SaveKey: "[Enter] Guardar",
		TestKey: "[T] Probar",
		CancelKey: "[Esc] Cancelar",
		TabSwitchInput: "Tab para cambiar entrada",
		TestPrompt: "T para probar prompt",
		UnnamedPrompt: "Prompt sin nombre",
		TranslateToChineseDefault: "Traduce el siguiente contenido al chino:",
		EmptyInput: "Texto de entrada vacío",
		NoAPIKeyConfigured: "Clave API no configurada",
		CreateTranslatorFailed: "Fallo al crear traductor: %v",
		TestSentenceAI: "La inteligencia artificial está cambiando nuestro estilo de vida.",
		UsingModel: "Usando",
		APINotConfigured: "API no configurada",
		
		// Status messages additional
		ConfigRefreshed: "✅ Configuración actualizada, el traductor se reinicializará",
		TranslateOnlyPrompt: "Solo traduce el siguiente contenido al español, sin respuestas ni explicaciones, solo la traducción:",
		CustomSuffix: " (personalizado)",
		PreviewLabel: "Vista previa:",
		SaveButton: "Enter Guardar",
		NotConfiguredBrackets: "(no configurado)",
		UnknownProvider: "Desconocido",
		RecordingHotkey: "🔴 Grabando atajo",
		SetMonitorHotkey: "Configurar atajo de monitoreo",
		SetSwitchPromptHotkey: "Configurar atajo de cambio de prompt",
		PressDesiredHotkey: "Presiona la combinación de teclas deseada",
		
		// Console messages
		MonitorStartedTray: "✅ Monitoreo iniciado desde bandeja",
		MonitorStoppedTray: "⏸️ Monitoreo detenido desde bandeja",
		AutoPasteEnabled: "✅ Pegado automático habilitado",
		AutoPasteDisabled: "❌ Pegado automático deshabilitado",
		HotkeysLabel: "Atajos:",
		MonitorToggleKey: "Alternar monitoreo: %s",
		SwitchStyleKey: "Cambiar estilo: %s",
		MonitorPausedByHotkey: "⏸ Monitoreo pausado (atajo)",
		MonitorResumedByHotkey: "▶ Monitoreo reanudado (atajo)",
		StartingTray: "Iniciando bandeja del sistema...",
		ControlFromTray: "Controla xiaoniao desde la bandeja del sistema",
		GoodbyeEmoji: "¡Adiós! 👋",
		DirectTranslation: "Traducción directa",
		TranslateToChineseColon: "Traduce el siguiente contenido al chino:",
		
		// API config messages
		NoModelsFound: "No se encontraron modelos",
		CurrentSuffix: " (actual)",
		UnrecognizedAPIKey: "No se puede reconocer la clave API: %v",
		ConnectionFailed: "Conexión fallida (%s): %v",
		ConnectionSuccessNoModels: "Conexión exitosa (%s) - No se puede obtener lista de modelos: %v",
		ConnectionSuccessWithModels: "Conexión exitosa (%s) - %d modelos",
		TestingInProgress: "Probando...",
		
		// System hotkey
		SystemHotkeyFormat: "Atajo del sistema: %s",
		SystemHotkeyLabel: "Atajo del sistema",
		XiaoniaoToggleMonitor: "xiaoniao alternar monitoreo",
		XiaoniaoSwitchStyle: "xiaoniao cambiar estilo",
		
		// Translator error detection
		CannotProceed: "No se puede proceder",
		AIReturnedMultiline: "IA devolvió múltiples líneas (longitud: %d)",
		UsingFirstLine: "Usando solo la primera línea: %s",
		CannotTranslate: "No se puede traducir",
		UnableToTranslate: "Imposible traducir",
		Sorry: "Lo siento",
		
		// Theme names and descriptions
		DefaultThemeName: "Predeterminado",
		DefaultThemeDesc: "Tema azul clásico",
		TokyoNightDesc: "Tema oscuro inspirado en el paisaje nocturno de Tokio",
		SoftPastelDesc: "Tema de colores pastel suaves",
		MinimalThemeName: "Minimalista",
		MinimalThemeDesc: "Tema blanco y negro simple",
		
		// Tray messages
		StatusTranslated: "Estado: %d traducciones",
		DefaultPrompt: "Predeterminado",
		TrayMonitoring: "xiaoniao - Monitoreando | Estilo: %s",
		TrayStopped: "xiaoniao - Detenido | Estilo: %s",
		StyleLabel: "Estilo",

		// New i18n fields for v1.0
		Save: "Guardar",
		FormatError: "Error de formato: Use el formato 'Modificador+Tecla', como 'Ctrl+Q'",
		InvalidModifier: "Modificador inválido: %s",
		InvalidMainKey: "Tecla principal inválida: %s",
		ProviderLabel: "Proveedor: ",
		CommonExamples: "Ejemplos comunes",
		InputFormat: "Formato de entrada",
		ModifierPlusKey: "Modificador+Tecla",
		SingleModifier: "Modificador único",
		SingleKey: "Tecla única",
		SwitchFunction: "Cambiar función",
		Edit: "Editar",
		Back: "Atrás",

		// Critical missing fields from main.go
		ProgramAlreadyRunning: "El programa ya se está ejecutando. Por favor, compruebe el icono de la bandeja del sistema.\nSi no ve el icono de la bandeja, intente finalizar todos los procesos de xiaoniao y reinicie.",
		TrayManagerInitFailed: "Error al inicializar el administrador de la bandeja: %v\n\nPor favor, verifique si su sistema admite la función de bandeja del sistema.",
		SystemTrayStartFailed: "Error al iniciar la bandeja del sistema: %v\n\nPosibles razones:\n1. La función de bandeja del sistema está deshabilitada\n2. Permisos insuficientes\n3. Recursos del sistema insuficientes\n\nPor favor, verifique la configuración del sistema e intente nuevamente.",
		NotConfiguredStatus: "No configurado",
		PleaseConfigureAPIFirst: "Por favor, configure la API primero",
		APIConfigCompleted: "Configuración de API completada, reinicializando el servicio de traducción...",
		MonitorStartedConsole: "Monitor iniciado",
		MonitorPausedConsole: "Monitor pausado",
		ExportLogsFailed: "Error al exportar los registros: %v",
		LogsExportedTo: "Registros exportados a: %s",
		ConfigRefreshedDetail: "Configuración actualizada: %s | %s | %s",
		RefreshConfigFailed: "Error al actualizar la configuración: %v",
		SwitchedTo: "Cambiado a: %s",
		ConfigRefreshedAndReinit: "Configuración actualizada, el traductor se reinicializará",
		MonitorPausedMsg: "Monitor pausado",
		MonitorResumedMsg: "Monitor reanudado",
		SwitchPromptMsg: "🔄 Cambiar prompt: %s",
		TranslationStyle: "Estilo de traducción: %s",
		AutoPasteEnabledMsg: "Pegado automático: Habilitado",
		HotkeysColon: "Teclas de acceso rápido:",
		MonitorToggleLabel: "  Alternar monitor: %s",
		SwitchStyleLabel: "  Cambiar estilo: %s",
		MonitorStartedCopyToTranslate: "Monitor iniciado - Copiar texto para traducir",
		StartTranslating: "Iniciando traducción: %s",
		UsingPrompt: "Usando prompt: %s (longitud del contenido: %d)",
		TranslationFailedError: " Error\n  Error: %v",
		TranslationComplete: " Completado (#%d)",
		OriginalText: "  Original: %s",
		TranslatedText: "  Traducción: %s",
		MonitorPausedViaHotkey: "Monitor pausado (mediante tecla de acceso rápido)",
		MonitorResumedViaHotkey: "Monitor reanudado (mediante tecla de acceso rápido)",
		SwitchPromptViaHotkey: "🔄 Cambiar prompt: %s (mediante tecla de acceso rápido)",
		PrewarmingModel: "Precalentando modelo...",
		PrewarmSuccess2: " Éxito",
		PrewarmSkip: " Omitir (se puede ignorar: %v)",
		TranslatorRefreshed: "Traductor actualizado: %s | %s",
		TranslatorRefreshFailed: "Error al actualizar el traductor: %v",

		// Missing from config_ui.go
		ConfigRefreshedReinit: "✅ Configuración actualizada, el traductor se reinicializará",
		MainModelChanged: "✅ Modelo principal cambiado a: %s",
		TestingModelMsg: "🔄 Probando modelo...",
		ModelInitFailed: "Error al inicializar el modelo %s: %v",
		TranslateToChineseOnly: "Por favor, traduzca solo lo siguiente al chino, no responda ni explique, solo muestre la traducción:",
		ModelTestFailedMsg: "Prueba del modelo %s fallida: %v",
		ModelAvailable: "✅ ¡Modelo %s disponible! Traducción: %s",
		ModelNoResponse: "❌ Modelo %s sin respuesta",
		DeleteFailed: "Error al eliminar: %v",
		SaveFailed: "Error al guardar: %v",
		UpdateFailed: "Error al actualizar: %v",
		TestingConnectionMsg: "Probando conexión...",
		TestingMsg: "Probando...",
		CreateTranslatorFailedMsg: "❌ Error al crear el traductor: %v",
		TranslationFailedMsg: "❌ Error en la traducción: %v",
		TranslationResultMsg: "✅ Resultado de la traducción:\nOriginal: %s\nTraducción: %s\nModelo: %s\nPrompt: %s",
		PreviewColon: "Vista previa:",
		VersionFormat: "Versión: %s",
		MonitorStatusActiveMsg: "Estado del monitor: Activo",
		MonitorStatusPausedMsg: "Estado del monitor: Pausado",
		TranslationCountMsg: "Número de traducciones: %d",
		StatusActive: "Activo",
		StatusPaused: "Pausado",
		ModelLabel: "Modelo: ",
		APILabel: "API: ",
		TryAgainMsg: " (Presione Enter para reintentar)",
		StatsFormat: " | Entrada: %d | Salida: %d | Total: %d",

		// Tray and logs
		ExportLogs: "Exportar registros",
		GetProgramPathFailed: "Error al obtener la ruta del programa",
		WriteLogFileFailed: "Error al escribir el archivo de registro",

		// Additional missing fields
		AuthorLabel: "Autor:",
		ClassicBlueFallback: "Tema azul clásico",
		CleanBWFallback: "Tema blanco y negro limpio",
		ConnectionFailedFormat: "Conexión fallida: %v",
		CreatePromptsJsonFailed: "Error al crear prompts.json:",
		DarkThemeTokyoNightFallback: "Tema oscuro inspirado en Tokyo Night",
		DefaultThemeNameFallback: "Predeterminado",
		DeleteBuiltinPromptError: "Error al eliminar prompt integrado:",
		DetectedProvider: "Proveedor detectado:",
		EnterYourAPIKey: "Ingrese su clave API",
		HotkeySettingsTitle: "Configuración de teclas de acceso rápido",
		HotkeysSaved: "✅ Teclas de acceso rápido guardadas",
		LicenseLabel: "Licencia:",
		LoadUserPromptsFailed: "Error al cargar prompts del usuario:",
		MinimalThemeNameFallback: "Mínimo",
		ModelAvailableTranslation: "✅ %s disponible! Traducción: %s",
		ModelUnavailable: "❌ %s no disponible: %v",
		MonitorToggleHotkey: "Alternar monitor",
		PleaseSelectModel: "Por favor seleccione un modelo",
		ProjectUrlLabel: "URL del proyecto:",
		SelectAIModel: "Seleccionar modelo de IA",
		SelectedBrackets: "[Seleccionado]",
		SoftPastelFallback: "Tema pastel suave",
		StatusTranslatedCount: "Estado: Traducido %d veces",
		Success: "¡Éxito!",
		SwitchStyleHotkey: "Cambiar estilo",
		TestingConnectionDots: "Probando conexión...",
		TestingModelFormat: "Probando %s...",
		TranslateToChineseProvider: "Traducir al chino",
		Tutorial: "Tutorial",
		TutorialContent: `Guía de inicio rápido:

1. Configurar clave API
   • Seleccione "Configuración API" del menú principal
   • Ingrese su clave API (OpenAI, Anthropic, etc.)
   • El sistema detectará automáticamente el proveedor

2. Seleccionar modelo
   • Después de configurar la API, seleccione "Seleccionar modelo"
   • Elija un modelo de IA de la lista

3. Configurar teclas de acceso rápido (Opcional)
   • Seleccione "Configuración de teclas de acceso rápido" del menú principal
   • Configure las teclas para alternar monitor y cambiar prompt

4. Comenzar a usar
   • Ctrl+X para cortar o Ctrl+C para copiar texto activa la traducción
   • El programa reemplaza automáticamente el contenido del portapapeles
   • Ctrl+V para pegar el resultado traducido
   • Algunas aplicaciones pueden requerir pegado manual

Tutoriales en video:
• Bilibili: (Próximamente)
• YouTube: (Próximamente)`,
		UnknownProviderDefault: "Proveedor desconocido (predeterminado: OpenAI)",
		UnsupportedOS: "Sistema operativo no compatible: %s",
		XiaoniaoMonitoring: "xiaoniao - Monitoreando | Estilo: %s",
		XiaoniaoStopped: "xiaoniao - Detenido | Estilo: %s",
	}
}
// getFrench returns French translations
func getFrench() *Translations {
	return &Translations{
		// Critical system messages
		ProgramAlreadyRunning: "Le programme est déjà en cours d'exécution. Veuillez vérifier l'icône de la barre d'état système.",
		TrayManagerInitFailed: "Échec de l'initialisation du gestionnaire de la barre d'état système : %v",
		MonitorStartedConsole: "Surveillance démarrée",
		MonitorPausedConsole: "Surveillance mise en pause",

		// Config refresh messages
		ConfigRefreshedReinit: "Configuration actualisée, réinitialisation...",

		// Model testing
		ModelTestFailed: "Test du modèle échoué: %s - %v",
		ModelInitFailed: "Erreur d'initialisation du modèle %s: %v",
		TranslateToChineseOnly: "Veuillez traduire uniquement le texte suivant en chinois, ne répondez pas et n'expliquez pas, montrez seulement la traduction:",
		ModelTestFailedMsg: "Test du modèle %s échoué: %v",
		ModelAvailable: "✅ Modèle %s disponible! Traduction: %s",
		ModelNoResponse: "❌ Modèle %s sans réponse",
		DeleteFailed: "Échec de la suppression: %v",
		SaveFailed: "Échec de l'enregistrement: %v",
		UpdateFailed: "Échec de la mise à jour: %v",
		TestingConnectionMsg: "Test de connexion...",
		TestingMsg: "Test en cours...",
		CreateTranslatorFailedMsg: "❌ Échec de la création du traducteur: %v",
		TranslationFailedMsg: "❌ Échec de la traduction: %v",
		TranslationResultMsg: "✅ Résultat de la traduction:\nOriginal: %s\nTraduction: %s\nModèle: %s\nPrompt: %s",
		PreviewColon: "Aperçu:",
		VersionFormat: "Version: %s",
		MonitorStatusActiveMsg: "État de la surveillance: Actif",
		MonitorStatusPausedMsg: "État de la surveillance: En pause",
		TranslationCountMsg: "Nombre de traductions: %d",
		StatusActive: "Actif",
		StatusPaused: "En pause",
		ModelLabel: "Modèle: ",
		APILabel: "API: ",
		TryAgainMsg: " (Appuyez sur Entrée pour réessayer)",
		StatsFormat: " | Entrée: %d | Sortie: %d | Total: %d",

		// API config messages
		ConnectionFailed: "Échec de la connexion",
		TestingConnection: "Test en cours...",
		NoModelsFound: "Aucun modèle trouvé",
		CurrentSuffix: " (Actuel)",
		UnrecognizedAPIKey: "Impossible de reconnaître la clé API: %v",
		ConnectionSuccessNoModels: "Connexion réussie (%s) - Impossible d'obtenir la liste des modèles: %v",
		ConnectionSuccessWithModels: "Connexion réussie (%s) - %d modèles",
		TestingInProgress: "Test en cours...",

		// Tray and logs
		ExportLogs: "Exporter les journaux",
		GetProgramPathFailed: "Échec de l'obtention du chemin du programme",
		WriteLogFileFailed: "Échec de l'écriture du fichier journal",

		// Main interface
		Title:           "Configuration xiaoniao",
		ConfigTitle:     "xiaoniao - Paramètres",
		APIKey:          "Clé API",
		APIConfig:       "Configuration API",
		TranslateStyle:  "Style de traduction",
		TestConnection:  "Test de traduction",
		SaveAndExit:     "Enregistrer et quitter",
		Language:        "Langue de l'interface",
		ManagePrompts:   "Gérer les prompts",
		Theme:           "Thème de l'interface",
		Hotkeys:         "Raccourcis clavier",
		AutoPaste:       "Collage automatique",
		
		// Status messages
		Provider:        "Fournisseur",
		Model:           "Modèle",
		NotSet:          "Non configuré",
		Testing:         "Test de connexion...",
		TestSuccess:     "✅ Connexion réussie!",
		TestFailed:      "❌ Échec de connexion",
		APIKeySet:       "Clé API configurée",
		APIKeyNotSet:    "Clé API non configurée",
		ChangeModel:     "Changer de modèle",
		Enabled:         "Activé",
		Disabled:        "Désactivé",
		
		// Help information
		HelpMove:        "↑↓ Déplacer",
		HelpSelect:      "Entrée Sélectionner",
		HelpBack:        "Échap Retour",
		HelpQuit:        "Ctrl+C Quitter",
		HelpTab:         "Tab Basculer",
		HelpEdit:        "e Éditer",
		HelpDelete:      "d Supprimer",
		HelpAdd:         "+ Ajouter",
		
		// Prompt management
		PromptManager:   "Gestionnaire de prompts",
		AddPrompt:       "Ajouter un prompt",
		EditPrompt:      "Éditer le prompt",
		DeletePrompt:    "Supprimer le prompt",
		PromptName:      "Nom",
		PromptContent:   "Contenu",
		ConfirmDelete:   "Confirmer la suppression?",
		
		// Running interface
		Running:         "En cours",
		Monitoring:      "Surveillance du presse-papiers...",
		CopyToTranslate: "Copiez du texte pour traduire automatiquement",
		ExitTip:         "Appuyez sur Ctrl+C pour quitter",
		Translating:     "Traduction...",
		Complete:        "Terminé",
		Failed:          "Échoué",
		Original:        "Original",
		Translation:     "Traduction",
		TotalCount:      "Total traduit",
		Goodbye:         "Au revoir! 👋",
		TranslateCount:  "fois",
		
		// Help documentation
		HelpTitle:       "xiaoniao",
		HelpDesc:        "Outil de traduction du presse-papiers alimenté par IA",
		Commands:        "Description des commandes",
		RunCommand:      "xiaoniao run",
		RunDesc:         "Démarrer la surveillance du presse-papiers et traduire automatiquement le contenu copié",
		TrayCommand:     "xiaoniao tray",
		TrayDesc:        "Démarrer le mode barre d'état système",
		ConfigCommand:   "xiaoniao config",
		ConfigDesc:      "Ouvrir l'interface de configuration interactive",
		HelpCommand:     "xiaoniao help",
		HelpDesc2:       "Afficher cette aide",
		VersionCommand:  "xiaoniao version",
		VersionDesc:     "Afficher les informations de version",
		HowItWorks:      "Comment ça marche",
		Step1:           "Exécutez xiaoniao config pour configurer l'API",
		Step2:           "Exécutez xiaoniao run pour démarrer la surveillance",
		Step3:           "Copiez n'importe quel texte (Ctrl+C)",
		Step4:           "Traduction automatique et remplacement du presse-papiers",
		Step5:           "Quand vous entendez le son, collez directement (Ctrl+V)",
		Warning:         "Attention: La traduction écrasera le contenu original du presse-papiers!",
		
		// Error messages
		NoAPIKey:        "❌ Clé API non configurée",
		RunConfigFirst:  "Veuillez d'abord exécuter: xiaoniao config",
		AlreadyRunning:  "❌ xiaoniao est déjà en cours d'exécution",
		InitFailed:      "Échec de l'initialisation",
		ConfigNotFound:  "Fichier de configuration introuvable",
		InvalidAPIKey:   "Clé API invalide",
		NetworkError:    "Erreur de connexion réseau",
		TranslateFailed: "Échec de la traduction",
		
		// API Config
		EnterAPIKey:     "Veuillez entrer la clé API",
		EnterNewAPIKey:  "Entrer une nouvelle clé API",
		ChangeAPIKey:    "Changer la clé API",
		SelectMainModel: "Sélectionner le modèle principal",
		SupportedProviders: "Fournisseurs pris en charge",
		SearchModel:     "Rechercher un modèle...",
		MainModel:       "Modèle principal",
		NoPromptAvailable: "(Aucun prompt disponible)",
		
		// Usage messages
		Usage:           "Utilisation",
		UnknownCommand:  "Commande inconnue",
		OpeningConfig:   "Ouverture de l'interface de configuration...",
		
		// Tray menu
		TrayShow:        "Afficher la fenêtre",
		TrayHide:        "Masquer la fenêtre",
		TraySettings:    "Paramètres",
		TrayQuit:        "Quitter",
		TrayToggle:      "Basculer la surveillance",
		TrayRefresh:     "Actualiser la configuration",
		TrayAbout:       "À propos",
		
		// Theme related
		SelectTheme:      "Sélectionner le thème de l'interface",
		DefaultTheme:     "Par défaut",
		ClassicBlue:      "Thème bleu classique",
		DarkTheme:        "Thème sombre",
		
		// Hotkey related
		HotkeySettings:   "Paramètres des raccourcis",
		ToggleMonitor:    "Basculer la surveillance",
		SwitchPromptKey:  "Changer de prompt",
		PressEnterToSet:  "Appuyez sur Entrée pour définir le raccourci",
		PressDeleteToClear: "Appuyez sur Suppr pour effacer le raccourci",
		NotConfigured:    "(Non configuré)",
		
		// Test translation
		TestTranslation:  "Test de traduction",
		CurrentConfig:    "Configuration actuelle",
		EnterTextToTranslate: "Entrez le texte à traduire",
		TranslationResult: "Résultat de la traduction",
		
		// About page
		About:            "À propos de xiaoniao",
		Author:           "Auteur",
		License:          "Licence open source",
		ProjectUrl:       "URL du projet",
		SupportAuthor:    "💝 Soutenir l'auteur",
		PriceNote:        "Le prix du produit est de 1$, mais vous pouvez l'utiliser gratuitement.",
		ShareNote:        "Si cela vous a vraiment aidé, offrez-moi un café\nou partagez-le avec plus de personnes! :)",
		ThanksForUsing:   "Merci d'utiliser xiaoniao!",
		BackToMainMenu:   "[Échap] Retour au menu principal",
		ComingSoon:       "(Bientôt open source)",
		
		// Model selection
		TotalModels:      "Total %d modèles",
		SearchModels:     "Rechercher",
		SelectToConfirm:  "Sélectionner",
		TestModel:        "Tester",
		SearchSlash:      "Rechercher",
		
		// Debug info
		DebugInfo:        "Informations de débogage",
		CursorPosition:   "Curseur",
		InputFocus:       "Focus d'entrée",
		KeyPressed:       "Touche pressée",
		
		// Additional messages
		MonitorStarted:  "✅ Surveillance démarrée",
		MonitorStopped:  "⏸️ Surveillance arrêtée",
		StopMonitor:     "Arrêter la surveillance",
		StartMonitor:    "Démarrer la surveillance",
		ConfigUpdated:   "✅ Configuration mise à jour",
		RefreshFailed:   "❌ Échec de la mise à jour de la configuration",
		SwitchPrompt:    "Basculé vers",
		PrewarmModel:    "Préchauffage du modèle...",
		PrewarmSuccess:  "✅",
		PrewarmFailed:   "⚠️ (ignorable: %v)",
		
		// Additional UI text
		WaitingForKeys:  "En attente de touches...",
		DetectedKeys:    "Détecté",
		HotkeyTip:       "Conseil",
		HoldModifier:    "Maintenez Ctrl/Alt/Shift + autre touche",
		DetectedAutoSave: "Sauvegarde auto après détection de combinaison",
		PressEscCancel:  "Appuyez sur ESC pour annuler",
		DefaultName:     "Par défaut",
		MinimalTheme:    "Minimaliste",
		
		// Model selection
		ConnectionSuccess: "Connexion réussie",
		ModelsCount:      "%d modèles",
		SelectModel:      "Sélectionner",
		TestingModel:     "Test du modèle %s...",
		SearchModels2:    "Rechercher",
		TotalModelsCount: "Total %d modèles",
		
		// Hotkey messages
		HotkeyAvailable:  "✅ Disponible, appuyez sur Entrée pour confirmer",
		PressEnterConfirm: "Appuyez sur Entrée pour confirmer",
		
		// Help text additions
		HelpEnterConfirm: "Entrée Confirmer",
		HelpTabSwitch:    "Tab Basculer",
		HelpEscReturn:    "Échap Retour",
		HelpUpDownSelect: "↑↓ Sélectionner",
		HelpTTest:        "T Tester",
		HelpSearchSlash:  "/ Rechercher",
		HelpTranslate:    "Entrée: Traduire",
		HelpCtrlSSaveExit: "Ctrl+S Enregistrer et quitter",
		HelpCtrlSSave:    "Ctrl+S Enregistrer",
		
		// Theme descriptions
		DarkThemeTokyoNight: "Thème sombre inspiré du paysage nocturne de Tokyo",
		ChocolateTheme:      "Thème chocolat noir",
		LatteTheme:          "Thème latte lumineux",
		DraculaTheme:        "Thème Dracula sombre",
		GruvboxDarkTheme:    "Thème rétro sombre",
		GruvboxLightTheme:   "Thème rétro clair",
		NordTheme:           "Style minimaliste nordique",
		SolarizedDarkTheme:  "Thème sombre qui protège les yeux",
		SolarizedLightTheme: "Thème clair qui protège les yeux",
		MinimalBWTheme:      "Thème noir et blanc simple",
		
		// Prompt management additions
		HelpNewPrompt:    "n Nouveau",
		HelpEditPrompt:   "e Éditer",
		HelpDeletePrompt: "d Supprimer",
		ConfirmDeleteKey: "Appuyez sur d pour confirmer la suppression",
		CancelDelete:     "Appuyez sur une autre touche pour annuler",
		
		// Status messages
		DetectingProvider: "Détection réussie",
		
		// About page additions
		ProjectAuthor: "Auteur",
		OpenSourceLicense: "Licence open source",
		AuthorName: "梨梨果",
		
		// Key bindings help
		KeyUp: "Haut",
		KeyDown: "Bas",
		KeySelect: "Sélectionner",
		KeyReturn: "Retour",
		KeyQuit: "Quitter",
		KeySwitch: "Basculer",
		KeyEdit: "Éditer",
		KeyDelete: "Supprimer",
		KeyNew: "Nouveau",
		KeyTest: "Tester",
		
		// Prompt test UI
		TestPromptTitle: "Test du Prompt",
		CurrentPrompt: "Prompt Actuel",
		PromptContentLabel: "Contenu",
		TestText: "Texte de test",
		TestingAI: "Appel de la traduction IA",
		TranslationResultLabel: "Résultat de traduction",
		InputTestText: "Entrez le texte à tester...",
		ResultWillShowHere: "Le résultat de la traduction apparaîtra ici...",
		TranslatingText: "Traduction...",
		TabSwitchFocus: "Tab pour changer le focus",
		CtrlEnterTest: "Ctrl+Entrée pour tester",
		EscReturn: "Échap pour retour",
		EditingPrompt: "Édition",
		NewPrompt: "Nouveau Prompt",
		NameLabel: "Nom",
		ContentLabel: "Contenu",
		SaveKey: "[Entrée] Enregistrer",
		TestKey: "[T] Tester",
		CancelKey: "[Échap] Annuler",
		TabSwitchInput: "Tab pour changer l'entrée",
		TestPrompt: "T pour tester le prompt",
		UnnamedPrompt: "Prompt sans nom",
		TranslateToChineseDefault: "Traduis le contenu suivant en chinois:",
		EmptyInput: "Texte d'entrée vide",
		NoAPIKeyConfigured: "Clé API non configurée",
		CreateTranslatorFailed: "Échec de création du traducteur: %v",
		TestSentenceAI: "L'intelligence artificielle change notre mode de vie.",
		UsingModel: "Utilisation",
		APINotConfigured: "API non configurée",
		
		// Status messages additional
		ConfigRefreshed: "✅ Configuration actualisée, le traducteur sera réinitialisé",
		TranslateOnlyPrompt: "Traduis uniquement le contenu suivant en français, sans réponse ni explication, seulement la traduction:",
		CustomSuffix: " (personnalisé)",
		PreviewLabel: "Aperçu:",
		SaveButton: "Entrée Enregistrer",
		NotConfiguredBrackets: "(non configuré)",
		UnknownProvider: "Inconnu",
		RecordingHotkey: "🔴 Enregistrement du raccourci",
		SetMonitorHotkey: "Définir le raccourci de surveillance",
		SetSwitchPromptHotkey: "Définir le raccourci de changement de prompt",
		PressDesiredHotkey: "Appuyez sur la combinaison de touches souhaitée",
		
		// Console messages
		MonitorStartedTray: "✅ Surveillance démarrée depuis la barre d'état",
		MonitorStoppedTray: "⏸️ Surveillance arrêtée depuis la barre d'état",
		AutoPasteEnabled: "✅ Collage automatique activé",
		AutoPasteDisabled: "❌ Collage automatique désactivé",
		HotkeysLabel: "Raccourcis:",
		MonitorToggleKey: "Basculer surveillance: %s",
		SwitchStyleKey: "Changer style: %s",
		MonitorPausedByHotkey: "⏸ Surveillance mise en pause (raccourci)",
		MonitorResumedByHotkey: "▶ Surveillance reprise (raccourci)",
		StartingTray: "Démarrage de la barre d'état système...",
		ControlFromTray: "Contrôlez xiaoniao depuis la barre d'état système",
		GoodbyeEmoji: "Au revoir! 👋",
		DirectTranslation: "Traduction directe",
		TranslateToChineseColon: "Traduis le contenu suivant en chinois:",
		
		// API config messages
		
		// System hotkey
		SystemHotkeyFormat: "Raccourci système: %s",
		SystemHotkeyLabel: "Raccourci système",
		XiaoniaoToggleMonitor: "xiaoniao basculer surveillance",
		XiaoniaoSwitchStyle: "xiaoniao changer style",
		
		// Translator error detection
		CannotProceed: "Impossible de procéder",
		AIReturnedMultiline: "L'IA a retourné plusieurs lignes (longueur: %d)",
		UsingFirstLine: "Utilisation de la première ligne seulement: %s",
		CannotTranslate: "Impossible de traduire",
		UnableToTranslate: "Traduction impossible",
		Sorry: "Désolé",
		
		// Theme names and descriptions
		DefaultThemeName: "Par défaut",
		DefaultThemeDesc: "Thème bleu classique",
		TokyoNightDesc: "Thème sombre inspiré du paysage nocturne de Tokyo",
		SoftPastelDesc: "Thème aux couleurs pastel douces",
		MinimalThemeName: "Minimaliste",
		MinimalThemeDesc: "Thème noir et blanc simple",
		
		// Tray messages
		StatusTranslated: "Statut: %d traductions",
		DefaultPrompt: "Par défaut",
		TrayMonitoring: "xiaoniao - Surveillance | Style: %s",
		TrayStopped: "xiaoniao - Arrêté | Style: %s",
		StyleLabel: "Style",

		// New i18n fields for v1.0
		FormatError: "Erreur de format : Utilisez le format 'Modificateur+Touche', comme 'Ctrl+Q'",
		InvalidModifier: "Modificateur invalide : %s",
		ProviderLabel: "Fournisseur : ",
		CommonExamples: "Exemples courants",
		InputFormat: "Format d'entrée",
		ModifierPlusKey: "Modificateur+Touche",
		SwitchFunction: "Changer de fonction",
		Save: "Enregistrer",
		InvalidMainKey: "Touche principale invalide : %s",
		SingleModifier: "Modificateur unique",
		SingleKey: "Touche unique",
		Edit: "Modifier",
		Back: "Retour",

		// Additional missing fields
		APIConfigCompleted: "Configuration API terminée, réinitialisation du service de traduction...",
		AuthorLabel: "Auteur:",
		AutoPasteEnabledMsg: "Collage automatique activé",
		ClassicBlueFallback: "Thème bleu classique",
		CleanBWFallback: "Thème noir et blanc épuré",
		ConfigRefreshedAndReinit: "Configuration actualisée et réinitialisée",
		ConfigRefreshedDetail: "Configuration actualisée, réinitialisation...",
		ConnectionFailedFormat: "Échec de la connexion: %v",
		CreatePromptsJsonFailed: "Échec de la création de prompts.json:",
		DarkThemeTokyoNightFallback: "Thème sombre inspiré de Tokyo Night",
		DefaultThemeNameFallback: "Par défaut",
		DeleteBuiltinPromptError: "Erreur lors de la suppression du prompt intégré:",
		DetectedProvider: "Fournisseur détecté:",
		EnterYourAPIKey: "Entrez votre clé API",
		ExportLogsFailed: "Échec de l'exportation des journaux",
		HotkeysColon: "Raccourcis:",
		HotkeySettingsTitle: "Paramètres des raccourcis clavier",
		HotkeysSaved: "✅ Raccourcis sauvegardés",
		LicenseLabel: "Licence:",
		LoadUserPromptsFailed: "Échec du chargement des prompts utilisateur:",
		LogsExportedTo: "Journaux exportés vers:",
		MainModelChanged: "Modèle principal changé: %s -> %s",
		MinimalThemeNameFallback: "Minimal",
		ModelAvailableTranslation: "✅ %s disponible! Traduction: %s",
		ModelUnavailable: "❌ %s indisponible: %v",
		MonitorPausedMsg: "Surveillance mise en pause",
		MonitorPausedViaHotkey: "Surveillance mise en pause via raccourci",
		MonitorResumedMsg: "Surveillance reprise",
		MonitorResumedViaHotkey: "Surveillance reprise via raccourci",
		MonitorStartedCopyToTranslate: "Surveillance démarrée - Copiez du texte pour traduire",
		MonitorToggleHotkey: "Basculer la surveillance",
		MonitorToggleLabel: "Basculer la surveillance",
		NotConfiguredStatus: "Non configuré",
		OriginalText: "Texte original",
		PleaseConfigureAPIFirst: "Veuillez d'abord configurer l'API",
		PleaseSelectModel: "Veuillez sélectionner un modèle",
		PrewarmingModel: "Préchauffage du modèle...",
		PrewarmSkip: "Passer le préchauffage",
		ProjectUrlLabel: "URL du projet:",
		RefreshConfigFailed: "Échec de l'actualisation de la configuration",
		SelectAIModel: "Sélectionner le modèle IA",
		SelectedBrackets: "[Sélectionné]",
		SoftPastelFallback: "Thème pastel doux",
		StartTranslating: "Commencer la traduction",
		StatusTranslatedCount: "Statut: Traduit %d fois",
		Success: "Succès!",
		SwitchedTo: "Basculé vers:",
		SwitchPromptMsg: "Prompt changé",
		SwitchPromptViaHotkey: "Prompt changé via raccourci",
		SwitchStyleHotkey: "Changer de style",
		SwitchStyleLabel: "Changer de style",
		SystemTrayStartFailed: "Échec du démarrage de la barre d'état système",
		TestingConnectionDots: "Test de connexion...",
		TestingModelFormat: "Test de %s...",
		TestingModelMsg: "Test du modèle...",
		TranslatedText: "Texte traduit",
		TranslateToChineseProvider: "Traduire en chinois",
		TranslationComplete: "Traduction terminée",
		TranslationFailedError: "Erreur de traduction",
		TranslationStyle: "Style de traduction",
		TranslatorRefreshed: "Traducteur actualisé",
		TranslatorRefreshFailed: "Échec de l'actualisation du traducteur",
		Tutorial: "Tutoriel",
		TutorialContent: `Guide de démarrage rapide :

1. Configurer la clé API
   • Sélectionnez "Configuration API" dans le menu principal
   • Entrez votre clé API (OpenAI, Anthropic, etc.)
   • Le système détectera automatiquement le fournisseur

2. Sélectionner le modèle
   • Après la configuration API, sélectionnez "Sélectionner le modèle"
   • Choisissez un modèle IA dans la liste

3. Configurer les raccourcis clavier (Optionnel)
   • Sélectionnez "Paramètres des raccourcis" dans le menu principal
   • Configurez les raccourcis pour basculer la surveillance et changer de prompt

4. Commencer à utiliser
   • Ctrl+X pour couper ou Ctrl+C pour copier du texte déclenche la traduction
   • Le programme remplace automatiquement le contenu du presse-papiers
   • Ctrl+V pour coller le résultat traduit
   • Certaines applications peuvent nécessiter un collage manuel

Tutoriels vidéo :
• Bilibili : (Bientôt disponible)
• YouTube : (Bientôt disponible)`,
		UnknownProviderDefault: "Fournisseur inconnu (par défaut: OpenAI)",
		UnsupportedOS: "Système d'exploitation non pris en charge: %s",
		UsingPrompt: "Utilisation du prompt:",
		XiaoniaoMonitoring: "xiaoniao - Surveillance | Style: %s",
		XiaoniaoStopped: "xiaoniao - Arrêté | Style: %s",
	}
}
