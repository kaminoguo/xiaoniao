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
		SelectMainModel: "選擇主模型",
		SelectFallback:  "選擇副模型",
		SupportedProviders: "支持的服務商",
		SearchModel:     "搜索模型...",
		MainModel:       "主模型",
		FallbackModel:   "副模型",
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
		SelectFallback:  "フォールバックモデルを選択",
		SupportedProviders: "サポートされているプロバイダー",
		SearchModel:     "モデルを検索...",
		MainModel:       "メインモデル",
		FallbackModel:   "フォールバックモデル",
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
		SelectFallback:  "대체 모델 선택",
		SupportedProviders: "지원되는 공급자",
		SearchModel:     "모델 검색...",
		MainModel:       "메인 모델",
		FallbackModel:   "대체 모델",
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
		SelectFallback:  "Seleccionar modelo de respaldo",
		SupportedProviders: "Proveedores soportados",
		SearchModel:     "Buscar modelo...",
		MainModel:       "Modelo principal",
		FallbackModel:   "Modelo de respaldo",
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
	}
}
// getFrench returns French translations
func getFrench() *Translations {
	return &Translations{
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
		SelectFallback:  "Sélectionner le modèle de secours",
		SupportedProviders: "Fournisseurs pris en charge",
		SearchModel:     "Rechercher un modèle...",
		MainModel:       "Modèle principal",
		FallbackModel:   "Modèle de secours",
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
		ModelTestFailed:  "Échec du test du modèle %s: %v",
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
		TestingConnection: "Test en cours...",
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
		NoModelsFound: "Aucun modèle trouvé",
		CurrentSuffix: " (actuel)",
		UnrecognizedAPIKey: "Impossible de reconnaître la clé API: %v",
		ConnectionFailed: "Échec de connexion (%s): %v",
		ConnectionSuccessNoModels: "Connexion réussie (%s) - Impossible d'obtenir la liste des modèles: %v",
		ConnectionSuccessWithModels: "Connexion réussie (%s) - %d modèles",
		TestingInProgress: "Test en cours...",
		
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
	}
}
