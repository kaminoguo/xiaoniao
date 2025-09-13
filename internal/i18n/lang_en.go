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
		SelectMainModel: "Select Main Model",
		SupportedProviders: "Supported Providers",
		SearchModel:     "Search models...",
		MainModel:       "Main Model",
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
	}
}