package i18n

// getEnglish returns English translations
func getEnglish() *Translations {
	return &Translations{
		// Main interface
		Title:           "xiaoniao Config",
		ConfigTitle:     "Xiaoniao Translator - Settings",
		APIKey:          "API Key",
		APIConfig:       "API Configuration",
		TranslateStyle:  "Translation Style",
		TestConnection:  "Test Connection",
		SaveAndExit:     "Save & Exit",
		Language:        "Interface Language",
		ManagePrompts:   "Manage Prompts",
		Theme:           "Theme",
		Hotkeys:         "Hotkeys",
		AutoPaste:       "Auto Paste",
		
		// Status messages
		Provider:        "Provider",
		Model:           "Model",
		NotSet:          "Not Set",
		Testing:         "Testing connection...",
		TestSuccess:     "✅ Connection successful!",
		TestFailed:      "❌ Connection failed",
		APIKeySet:       "API key is set",
		APIKeyNotSet:    "API key not set",
		ChangeModel:     "Change Model",
		Enabled:         "Enabled",
		Disabled:        "Disabled",
		
		// Help information
		HelpMove:        "↑↓ Move",
		HelpSelect:      "Enter Select",
		HelpBack:        "Esc Back",
		HelpQuit:        "Ctrl+C Quit",
		HelpTab:         "Tab Switch",
		HelpEdit:        "e Edit",
		HelpDelete:      "d Delete",
		HelpAdd:         "+ Add",
		
		// Prompt management
		PromptManager:   "Prompt Manager",
		AddPrompt:       "Add Prompt",
		EditPrompt:      "Edit Prompt",
		DeletePrompt:    "Delete Prompt",
		PromptName:      "Name",
		PromptContent:   "Content",
		ConfirmDelete:   "Confirm delete?",
		
		// Running interface
		Running:         "Running",
		Monitoring:      "Monitoring clipboard...",
		CopyToTranslate: "Copy any text to auto-translate",
		ExitTip:         "Press Ctrl+C to exit",
		Translating:     "Translating...",
		Complete:        "Complete",
		Failed:          "Failed",
		Original:        "Original",
		Translation:     "Translation",
		TotalCount:      "Total translations",
		Goodbye:         "Goodbye! 👋",
		TranslateCount:  "times",
		
		// Help documentation
		HelpTitle:       "xiaoniao",
		HelpDesc:        "AI-powered clipboard translator",
		Commands:        "Commands",
		RunCommand:      "xiaoniao run",
		RunDesc:         "Start clipboard monitoring for auto-translation",
		TrayCommand:     "xiaoniao tray",
		TrayDesc:        "Start system tray mode",
		ConfigCommand:   "xiaoniao config",
		ConfigDesc:      "Open interactive configuration",
		HelpCommand:     "xiaoniao help",
		HelpDesc2:       "Show this help",
		VersionCommand:  "xiaoniao version",
		VersionDesc:     "Show version information",
		HowItWorks:      "How it works",
		Step1:           "Run 'xiaoniao config' to configure API",
		Step2:           "Run 'xiaoniao run' to start monitoring",
		Step3:           "Copy any text (Ctrl+C)",
		Step4:           "Auto-translate and replace clipboard",
		Step5:           "Paste after hearing notification (Ctrl+V)",
		Warning:         "Note: Translation will overwrite clipboard content!",

		// Tutorial
		Tutorial:        "Tutorial",
		TutorialContent: `Quick Start Guide:

1. Configure API Key
   • Select "API Configuration" from main menu
   • Enter your API key (OpenAI, Anthropic, etc.)
   • System will auto-detect the provider

2. Select Model
   • After API setup, select "Select Model"
   • Choose an AI model from the list

3. Set Hotkeys (Optional)
   • Select "Hotkey Settings" from main menu
   • Configure toggle and prompt switch hotkeys

4. Start Using
   • Ctrl+X to cut or Ctrl+C to copy text triggers translation
   • Program auto-replaces clipboard content
   • Ctrl+V to paste translated result
   • Some apps may require manual paste

Video Tutorials:
• Bilibili: https://www.bilibili.com/video/BV13zpUzhEeK/
• YouTube: https://www.youtube.com/watch?v=iPye0tYkBaY`,

		// Error messages
		NoAPIKey:        "❌ No API key configured",
		RunConfigFirst:  "Please run: xiaoniao config",
		AlreadyRunning:  "❌ xiaoniao is already running",
		InitFailed:      "Initialization failed",
		ConfigNotFound:  "Configuration file not found",
		InvalidAPIKey:   "Invalid API key",
		NetworkError:    "Network connection error",
		TranslateFailed: "Translation failed",
		
		// API Config
		EnterAPIKey:     "Please enter API Key",
		EnterNewAPIKey:  "Enter new API Key",
		ChangeAPIKey:    "Change API Key",
		SelectMainModel: "Select Model",
		SupportedProviders: "Supported Providers",
		SearchModel:     "Search models...",
		MainModel:       "Model",
		NoPromptAvailable: "(No prompts available)",
		
		// Usage messages
		Usage:           "Usage",
		UnknownCommand:  "Unknown command",
		OpeningConfig:   "Opening configuration...",
		
		// Tray menu
		TrayShow:        "Show Window",
		TrayHide:        "Hide Window",
		TraySettings:    "Settings",
		TrayQuit:        "Quit",
		TrayToggle:      "Toggle Monitor",
		TrayRefresh:     "Refresh Config",
		TrayAbout:       "About",
		
		// Theme related
		SelectTheme:      "Select Theme",
		DefaultTheme:     "Default",
		ClassicBlue:      "Classic blue theme",
		DarkTheme:        "Dark theme",
		
		// Hotkey related
		HotkeySettings:   "Hotkey Settings",
		ToggleMonitor:    "Toggle Monitor",
		SwitchPromptKey:  "Switch Prompt",
		PressEnterToSet:  "Press Enter to set hotkey",
		PressDeleteToClear: "Press Delete to clear hotkey",
		NotConfigured:    "(Not configured)",
		
		// Test translation
		TestTranslation:  "Test Translation",
		CurrentConfig:    "Current Configuration",
		EnterTextToTranslate: "Enter text to translate",
		TranslationResult: "Translation Result",
		
		// About page
		About:            "About xiaoniao",
		Author:           "Author",
		License:          "License",
		ProjectUrl:       "Project URL",
		SupportAuthor:    "💝 Support Author",
		PriceNote:        "Price: $1, but free to use.",
		ShareNote:        "When it really helps you, buy me a coffee,\nor share it with more people! :)",
		ThanksForUsing:   "Thanks for using xiaoniao!",
		BackToMainMenu:   "[Esc] Back to main menu",
		ComingSoon:       "(Coming soon)",
		
		// Model selection
		TotalModels:      "Total %d models",
		SearchModels:     "Search",
		SelectToConfirm:  "Select",
		TestModel:        "Test",
		SearchSlash:      "Search",
		
		// Debug info
		DebugInfo:        "Debug info",
		CursorPosition:   "Cursor",
		InputFocus:       "Input focus",
		KeyPressed:       "Key",
		
		// Additional messages
		MonitorStarted:  "✅ Monitor started",
		MonitorStopped:  "⏸️ Monitor stopped",
		StopMonitor:     "Stop Monitor",
		StartMonitor:    "Start Monitor",
		ConfigUpdated:   "✅ Configuration updated",
		RefreshFailed:   "❌ Failed to refresh configuration",
		SwitchPrompt:    "Switched to",
		PrewarmModel:    "Prewarming model...",
		PrewarmSuccess:  "✅",
		PrewarmFailed:   "⚠️ (Can be ignored: %v)",
		
		// Additional UI text
		WaitingForKeys:  "Waiting for keys...",
		DetectedKeys:    "Detected",
		HotkeyTip:       "Tips",
		HoldModifier:    "Hold Ctrl/Alt/Shift + other keys",
		DetectedAutoSave: "Auto-save when key combination detected",
		PressEscCancel:  "Press ESC to cancel recording",
		DefaultName:     "Default",
		MinimalTheme:    "Minimal",
		
		// Model selection
		ConnectionSuccess: "Connection successful",
		ModelsCount:      "%d models",
		SelectModel:      "Select",
		TestingModel:     "Testing model %s...",
		ModelTestFailed:  "Model %s test failed: %v",
		SearchModels2:    "Search",
		TotalModelsCount: "Total %d models",
		
		// Hotkey messages
		HotkeyAvailable:  "✅ Available, press Enter to confirm",
		PressEnterConfirm: "Press Enter to confirm",
		
		// Help text additions
		HelpEnterConfirm: "Enter Confirm",
		HelpTabSwitch:    "Tab Switch",
		HelpEscReturn:    "Esc Return",
		HelpUpDownSelect: "↑↓ Select",
		HelpTTest:        "T Test",
		HelpSearchSlash:  "/ Search",
		HelpTranslate:    "Enter: Translate",
		HelpCtrlSSaveExit: "Ctrl+S Save & Exit",
		HelpCtrlSSave:    "Ctrl+S Save",
		
		// Theme descriptions
		DarkThemeTokyoNight: "Dark theme, inspired by Tokyo night",
		ChocolateTheme:      "Deep chocolate theme",
		LatteTheme:          "Bright latte theme",
		DraculaTheme:        "Vampire dark theme",
		GruvboxDarkTheme:    "Retro dark theme",
		GruvboxLightTheme:   "Retro light theme",
		NordTheme:           "Nordic minimalist style",
		SolarizedDarkTheme:  "Eye-friendly dark theme",
		SolarizedLightTheme: "Eye-friendly light theme",
		MinimalBWTheme:      "Clean black and white theme",
		
		// Prompt management additions
		HelpNewPrompt:    "n New",
		HelpEditPrompt:   "e Edit",
		HelpDeletePrompt: "d Delete",
		ConfirmDeleteKey: "Press d to confirm deletion",
		CancelDelete:     "Press any other key to cancel",
		
		// Status messages
		TestingConnection: "Testing...",
		DetectingProvider: "Detection successful",
		
		// About page additions
		ProjectAuthor: "Author",
		OpenSourceLicense: "License",
		AuthorName: "LiLiGuo",
		
		// Key bindings help
		KeyUp: "Up",
		KeyDown: "Down",
		KeySelect: "Select",
		KeyReturn: "Return",
		KeyQuit: "Quit",
		KeySwitch: "Switch",
		KeyEdit: "Edit",
		KeyDelete: "Delete",
		KeyNew: "New",
		KeyTest: "Test",
		
		// Prompt test UI
		TestPromptTitle: "Test Prompt",
		CurrentPrompt: "Current Prompt",
		PromptContentLabel: "Content",
		TestText: "Test Text",
		TestingAI: "Calling AI for translation",
		TranslationResultLabel: "Translation Result",
		InputTestText: "Enter text to test...",
		ResultWillShowHere: "Translation result will be displayed here...",
		TranslatingText: "Translating...",
		TabSwitchFocus: "Tab Switch Focus",
		CtrlEnterTest: "Ctrl+Enter Test",
		EscReturn: "Esc Return",
		EditingPrompt: "Editing",
		NewPrompt: "New Prompt",
		NameLabel: "Name",
		ContentLabel: "Content",
		SaveKey: "[Enter] Save",
		TestKey: "[T] Test",
		CancelKey: "[Esc] Cancel",
		TabSwitchInput: "Tab Switch Input",
		TestPrompt: "T Test prompt",
		UnnamedPrompt: "Unnamed Prompt",
		TranslateToChineseDefault: "Translate the following to Chinese:",
		EmptyInput: "Input text is empty",
		NoAPIKeyConfigured: "No API Key configured",
		CreateTranslatorFailed: "Failed to create translator: %v",
		TestSentenceAI: "Artificial intelligence is changing the way we live.",
		UsingModel: "Using",
		APINotConfigured: "API not configured",
		
		// Status messages additional
		ConfigRefreshed: "✅ Configuration refreshed, translator will be reinitialized",
		TranslateOnlyPrompt: "Please translate the following to Chinese only, do not answer or explain, just output the translation:",
		CustomSuffix: " (Custom)",
		PreviewLabel: "Preview:",
		SaveButton: "Enter Save",
		NotConfiguredBrackets: "(Not Configured)",
		UnknownProvider: "Unknown",
		RecordingHotkey: "🔴 Recording Hotkey",
		SetMonitorHotkey: "Set Monitor Toggle Hotkey",
		SetSwitchPromptHotkey: "Set Switch Prompt Hotkey",
		PressDesiredHotkey: "Press your desired key combination",
		
		// Console messages
		MonitorStartedTray: "✅ Monitor started via tray",
		MonitorStoppedTray: "⏸️ Monitor stopped via tray",
		AutoPasteEnabled: "✅ Auto-paste enabled",
		AutoPasteDisabled: "❌ Auto-paste disabled",
		HotkeysLabel: "Hotkeys:",
		MonitorToggleKey: "Monitor toggle: %s",
		SwitchStyleKey: "Switch style: %s",
		MonitorPausedByHotkey: "⏸ Monitor paused (via hotkey)",
		MonitorResumedByHotkey: "▶ Monitor resumed (via hotkey)",
		StartingTray: "Starting system tray...",
		ControlFromTray: "Please control xiaoniao from system tray",
		GoodbyeEmoji: "Goodbye! 👋",
		DirectTranslation: "Direct Translation",
		TranslateToChineseColon: "Translate the following to Chinese:",
		
		// API config messages
		NoModelsFound: "No models found",
		CurrentSuffix: " (Current)",
		UnrecognizedAPIKey: "Cannot recognize API Key: %v",
		ConnectionFailed: "Connection failed (%s): %v",
		ConnectionSuccessNoModels: "Connection successful (%s) - Unable to get model list: %v",
		ConnectionSuccessWithModels: "Connection successful (%s) - %d models",
		TestingInProgress: "Testing...",
		
		// System hotkey
		SystemHotkeyFormat: "System hotkey: %s",
		SystemHotkeyLabel: "System Hotkey",
		XiaoniaoToggleMonitor: "xiaoniao Toggle Monitor",
		XiaoniaoSwitchStyle: "xiaoniao Switch Style",
		
		// Translator error detection
		CannotProceed: "cannot proceed",
		AIReturnedMultiline: "AI returned multiple lines (length: %d)",
		UsingFirstLine: "Using first line only: %s",
		CannotTranslate: "cannot translate",
		UnableToTranslate: "unable to translate",
		Sorry: "sorry",
		
		// Theme names and descriptions
		DefaultThemeName: "Default",
		DefaultThemeDesc: "Classic blue theme",
		TokyoNightDesc: "Dark theme inspired by Tokyo night",
		SoftPastelDesc: "Soft pastel theme",
		MinimalThemeName: "Minimal",
		MinimalThemeDesc: "Clean black and white theme",
		
		// Tray messages
		StatusTranslated: "Status: Translated %d times",
		DefaultPrompt: "Default",
		TrayMonitoring: "xiaoniao - Monitoring | Style: %s",
		TrayStopped: "xiaoniao - Stopped | Style: %s",
		StyleLabel: "Style",

		// Missing translations from main.go
		ProgramAlreadyRunning: "Program is already running. Please check the system tray icon.\nIf you don't see the tray icon, try ending all xiaoniao processes and restart.",
		TrayManagerInitFailed: "Tray manager initialization failed: %v\n\nPlease check if your system supports system tray functionality.",
		SystemTrayStartFailed: "System tray startup failed: %v\n\nPossible reasons:\n1. System tray feature is disabled\n2. Insufficient permissions\n3. Insufficient system resources\n\nPlease check system settings and try again.",
		NotConfiguredStatus: "Not Configured",
		PleaseConfigureAPIFirst: "Please configure API first",
		APIConfigCompleted: "API configuration completed, reinitializing translation service...",
		MonitorStartedConsole: "Monitor started",
		MonitorPausedConsole: "Monitor paused",
		ExportLogsFailed: "Export logs failed: %v",
		LogsExportedTo: "Logs exported to: %s",
		ConfigRefreshedDetail: "Config refreshed: %s | %s | %s",
		RefreshConfigFailed: "Refresh config failed: %v",
		SwitchedTo: "Switched to: %s",
		ConfigRefreshedAndReinit: "Config refreshed, translator will be reinitialized",
		MonitorPausedMsg: "Monitor paused",
		MonitorResumedMsg: "Monitor resumed",
		SwitchPromptMsg: "🔄 Switch Prompt: %s",
		TranslationStyle: "Translation style: %s",
		AutoPasteEnabledMsg: "Auto-paste: Enabled",
		HotkeysColon: "Hotkeys:",
		MonitorToggleLabel: "  Monitor toggle: %s",
		SwitchStyleLabel: "  Switch style: %s",
		MonitorStartedCopyToTranslate: "Monitor started - Copy text to translate",
		StartTranslating: "Start translating: %s",
		UsingPrompt: "Using Prompt: %s (content length: %d)",
		TranslationFailedError: " Failed\n  Error: %v",
		TranslationComplete: " Complete (#%d)",
		OriginalText: "  Original: %s",
		TranslatedText: "  Translation: %s",
		MonitorPausedViaHotkey: "Monitor paused (via hotkey)",
		MonitorResumedViaHotkey: "Monitor resumed (via hotkey)",
		SwitchPromptViaHotkey: "🔄 Switch Prompt: %s (via hotkey)",
		PrewarmingModel: "Prewarming model...",
		PrewarmSuccess2: " Success",
		PrewarmSkip: " Skip (can be ignored: %v)",
		TranslatorRefreshed: "Translator refreshed: %s | %s",
		TranslatorRefreshFailed: "Translator refresh failed: %v",

		// Missing translations from config_ui.go
		ConfigRefreshedReinit: "✅ Config refreshed, translator will be reinitialized",
		MainModelChanged: "✅ Main model changed to: %s",
		TestingModelMsg: "🔄 Testing model...",
		ModelInitFailed: "Model %s initialization failed: %v",
		TranslateToChineseOnly: "Please translate the following to Chinese only, do not answer or explain, just output the translation:",
		ModelTestFailedMsg: "Model %s test failed: %v",
		ModelAvailable: "✅ Model %s available! Translation: %s",
		ModelNoResponse: "❌ Model %s no response",
		DeleteFailed: "Delete failed: %v",
		SaveFailed: "Save failed: %v",
		UpdateFailed: "Update failed: %v",
		TestingConnectionMsg: "Testing connection...",
		TestingMsg: "Testing...",
		CreateTranslatorFailedMsg: "❌ Failed to create translator: %v",
		TranslationFailedMsg: "❌ Translation failed: %v",
		TranslationResultMsg: "✅ Translation result:\nOriginal: %s\nTranslation: %s\nModel: %s\nPrompt: %s",
		PreviewColon: "Preview:",
		HotkeySettingsTitle: "Hotkey Settings",
		MonitorToggleHotkey: "Monitor Toggle",
		SwitchStyleHotkey: "Switch Style",
		AuthorLabel: "Author:",
		LicenseLabel: "License:",
		ProjectUrlLabel: "Project URL:",
		HotkeysSaved: "✅ Hotkeys saved",

		// Missing translations from api_config_ui.go
		EnterYourAPIKey: "Enter your API key",
		DetectedProvider: "Detected provider:",
		UnknownProviderDefault: "Unknown provider (defaulting to OpenAI)",
		Success: "Success!",
		SelectAIModel: "Select AI Model",
		SelectedBrackets: "[Selected]",
		PleaseSelectModel: "Please select a model",
		TestingModelFormat: "Testing %s...",
		ModelAvailableTranslation: "✅ %s available! Translation: %s",
		ModelUnavailable: "❌ %s unavailable: %v",
		TestingConnectionDots: "Testing connection...",
		ConnectionFailedFormat: "Connection failed: %v",

		// Missing translations from prompts.go
		LoadUserPromptsFailed: "Failed to load user prompts:",
		CreatePromptsJsonFailed: "Failed to create prompts.json:",
		DeleteBuiltinPromptError: "Error deleting built-in prompt:",

		// Missing translations from tray.go
		ExportLogs: "Export Logs",
		StatusTranslatedCount: "Status: Translated %d times",
		XiaoniaoMonitoring: "xiaoniao - Monitoring | Style: %s",
		XiaoniaoStopped: "xiaoniao - Stopped | Style: %s",

		// Missing translations from logbuffer.go
		GetProgramPathFailed: "Failed to get program path: %v",
		WriteLogFileFailed: "Failed to write log file: %v",
		UnsupportedOS: "Unsupported operating system: %s",

		// Missing translations from themes.go (keep in English as fallback)
		DefaultThemeNameFallback: "Default",
		ClassicBlueFallback: "Classic blue theme",
		DarkThemeTokyoNightFallback: "Dark theme inspired by Tokyo night",
		SoftPastelFallback: "Soft pastel theme",
		MinimalThemeNameFallback: "Minimal",
		CleanBWFallback: "Clean black and white theme",

		// Missing from together_provider.go
		TranslateToChineseProvider: "Translate to Chinese",

		// Missing from config_ui.go hotkey screen
		CommonExamples: "Common Examples",
		InputFormat: "Input Format",
		ModifierPlusKey: "Modifier+Key",
		SingleModifier: "Single Modifier",
		SingleKey: "Single Key",
		SwitchFunction: "Switch Function",
		Edit: "Edit",
		Save: "Save",
		Back: "Back",

		// Missing from hotkey_input.go
		FormatError: "Format error: Please use 'Modifier+Key' format, like 'Ctrl+Q'",
		InvalidModifier: "Invalid modifier: %s",
		InvalidMainKey: "Invalid main key: %s",

		// Missing from main.go
		ProviderLabel: "Provider: ",
	}
}