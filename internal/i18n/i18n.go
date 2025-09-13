package i18n

import (
	"encoding/json"
	"os"
	"strings"
)

// Language represents a supported language
type Language string

const (
	LangZhCN Language = "zh-CN" // 简体中文
	LangZhTW Language = "zh-TW" // 繁体中文
	LangEN   Language = "en"    // English
	LangJA   Language = "ja"    // 日本語
	LangKO   Language = "ko"    // 한국어
	LangES   Language = "es"    // Español
	LangFR   Language = "fr"    // Français
)

// Translations holds all text strings for the UI
type Translations struct {
	// Main interface
	Title           string `json:"title"`
	ConfigTitle     string `json:"config_title"`
	APIKey          string `json:"api_key"`
	APIConfig       string `json:"api_config"`
	TranslateStyle  string `json:"translate_style"`
	TestConnection  string `json:"test_connection"`
	SaveAndExit     string `json:"save_and_exit"`
	Language        string `json:"language"`
	ManagePrompts   string `json:"manage_prompts"`
	Theme           string `json:"theme"`
	Hotkeys         string `json:"hotkeys"`
	AutoPaste       string `json:"auto_paste"`
	
	// Status messages
	Provider        string `json:"provider"`
	Model           string `json:"model"`
	NotSet          string `json:"not_set"`
	Testing         string `json:"testing"`
	TestSuccess     string `json:"test_success"`
	TestFailed      string `json:"test_failed"`
	APIKeySet       string `json:"api_key_set"`
	APIKeyNotSet    string `json:"api_key_not_set"`
	ChangeModel     string `json:"change_model"`
	Enabled         string `json:"enabled"`
	Disabled        string `json:"disabled"`
	
	// Help information
	HelpMove        string `json:"help_move"`
	HelpSelect      string `json:"help_select"`
	HelpBack        string `json:"help_back"`
	HelpQuit        string `json:"help_quit"`
	HelpTab         string `json:"help_tab"`
	HelpEdit        string `json:"help_edit"`
	HelpDelete      string `json:"help_delete"`
	HelpAdd         string `json:"help_add"`
	
	// Prompt management
	PromptManager   string `json:"prompt_manager"`
	AddPrompt       string `json:"add_prompt"`
	EditPrompt      string `json:"edit_prompt"`
	DeletePrompt    string `json:"delete_prompt"`
	PromptName      string `json:"prompt_name"`
	PromptContent   string `json:"prompt_content"`
	ConfirmDelete   string `json:"confirm_delete"`
	
	// Running interface
	Running         string `json:"running"`
	Monitoring      string `json:"monitoring"`
	CopyToTranslate string `json:"copy_to_translate"`
	ExitTip         string `json:"exit_tip"`
	Translating     string `json:"translating"`
	Complete        string `json:"complete"`
	Failed          string `json:"failed"`
	Original        string `json:"original"`
	Translation     string `json:"translation"`
	TotalCount      string `json:"total_count"`
	Goodbye         string `json:"goodbye"`
	TranslateCount  string `json:"translate_count"`

	// Tutorial
	Tutorial        string `json:"tutorial"`
	TutorialContent string `json:"tutorial_content"`

	// Help documentation
	HelpTitle       string `json:"help_title"`
	HelpDesc        string `json:"help_desc"`
	Commands        string `json:"commands"`
	RunCommand      string `json:"run_command"`
	RunDesc         string `json:"run_desc"`
	TrayCommand     string `json:"tray_command"`
	TrayDesc        string `json:"tray_desc"`
	ConfigCommand   string `json:"config_command"`
	ConfigDesc      string `json:"config_desc"`
	HelpCommand     string `json:"help_command"`
	HelpDesc2       string `json:"help_desc2"`
	VersionCommand  string `json:"version_command"`
	VersionDesc     string `json:"version_desc"`
	HowItWorks      string `json:"how_it_works"`
	Step1           string `json:"step1"`
	Step2           string `json:"step2"`
	Step3           string `json:"step3"`
	Step4           string `json:"step4"`
	Step5           string `json:"step5"`
	Warning         string `json:"warning"`
	
	// Error messages
	NoAPIKey        string `json:"no_api_key"`
	RunConfigFirst  string `json:"run_config_first"`
	InitFailed      string `json:"init_failed"`
	ConfigNotFound  string `json:"config_not_found"`
	InvalidAPIKey   string `json:"invalid_api_key"`
	NetworkError    string `json:"network_error"`
	TranslateFailed string `json:"translate_failed"`
	AlreadyRunning  string `json:"already_running"`
	
	// API Config
	EnterAPIKey     string `json:"enter_api_key"`
	EnterNewAPIKey  string `json:"enter_new_api_key"`
	ChangeAPIKey    string `json:"change_api_key"`
	SelectMainModel string `json:"select_main_model"`
	SupportedProviders string `json:"supported_providers"`
	SearchModel     string `json:"search_model"`
	MainModel       string `json:"main_model"`
	NoPromptAvailable string `json:"no_prompt_available"`
	
	// Usage messages
	Usage           string `json:"usage"`
	UnknownCommand  string `json:"unknown_command"`
	OpeningConfig   string `json:"opening_config"`
	
	// Tray menu
	TrayShow        string `json:"tray_show"`
	TrayHide        string `json:"tray_hide"`
	TraySettings    string `json:"tray_settings"`
	TrayQuit        string `json:"tray_quit"`
	TrayToggle      string `json:"tray_toggle"`
	TrayRefresh     string `json:"tray_refresh"`
	TrayAbout       string `json:"tray_about"`
	
	// Theme related
	SelectTheme      string `json:"select_theme"`
	DefaultTheme     string `json:"default_theme"`
	ClassicBlue      string `json:"classic_blue"`
	DarkTheme        string `json:"dark_theme"`
	
	// Hotkey related
	HotkeySettings   string `json:"hotkey_settings"`
	ToggleMonitor    string `json:"toggle_monitor"`
	SwitchPromptKey  string `json:"switch_prompt_key"`
	PressEnterToSet  string `json:"press_enter_to_set"`
	PressDeleteToClear string `json:"press_delete_to_clear"`
	NotConfigured    string `json:"not_configured"`
	
	// Test translation
	TestTranslation  string `json:"test_translation"`
	CurrentConfig    string `json:"current_config"`
	EnterTextToTranslate string `json:"enter_text_to_translate"`
	TranslationResult string `json:"translation_result"`
	
	// About page
	About            string `json:"about"`
	Author           string `json:"author"`
	License          string `json:"license"`
	ProjectUrl       string `json:"project_url"`
	SupportAuthor    string `json:"support_author"`
	PriceNote        string `json:"price_note"`
	ShareNote        string `json:"share_note"`
	ThanksForUsing   string `json:"thanks_for_using"`
	BackToMainMenu   string `json:"back_to_main_menu"`
	ComingSoon       string `json:"coming_soon"`
	
	// Model selection
	TotalModels      string `json:"total_models"`
	SearchModels     string `json:"search_models"`
	SelectToConfirm  string `json:"select_to_confirm"`
	TestModel        string `json:"test_model"`
	SearchSlash      string `json:"search_slash"`
	
	// Debug info
	DebugInfo        string `json:"debug_info"`
	CursorPosition   string `json:"cursor_position"`
	InputFocus       string `json:"input_focus"`
	KeyPressed       string `json:"key_pressed"`
	
	// Additional messages
	MonitorStarted  string `json:"monitor_started"`
	MonitorStopped  string `json:"monitor_stopped"`
	StopMonitor     string `json:"stop_monitor"`
	StartMonitor    string `json:"start_monitor"`
	ConfigUpdated   string `json:"config_updated"`
	RefreshFailed   string `json:"refresh_failed"`
	SwitchPrompt    string `json:"switch_prompt"`
	PrewarmModel    string `json:"prewarm_model"`
	PrewarmSuccess  string `json:"prewarm_success"`
	PrewarmFailed   string `json:"prewarm_failed"`
	
	// Additional UI text
	WaitingForKeys  string `json:"waiting_for_keys"`
	DetectedKeys    string `json:"detected_keys"`
	HotkeyTip       string `json:"hotkey_tip"`
	HoldModifier    string `json:"hold_modifier"`
	DetectedAutoSave string `json:"detected_auto_save"`
	PressEscCancel  string `json:"press_esc_cancel"`
	DefaultName     string `json:"default_name"`
	MinimalTheme    string `json:"minimal_theme"`
	
	// Model selection
	ConnectionSuccess string `json:"connection_success"`
	ModelsCount      string `json:"models_count"`
	SelectModel      string `json:"select_model"`
	TestingModel     string `json:"testing_model"`
	ModelTestFailed  string `json:"model_test_failed"`
	SearchModels2    string `json:"search_models2"`
	TotalModelsCount string `json:"total_models_count"`
	
	// Hotkey messages
	HotkeyAvailable  string `json:"hotkey_available"`
	PressEnterConfirm string `json:"press_enter_confirm"`
	
	// Help text additions
	HelpEnterConfirm string `json:"help_enter_confirm"`
	HelpTabSwitch    string `json:"help_tab_switch"`
	HelpEscReturn    string `json:"help_esc_return"`
	HelpUpDownSelect string `json:"help_up_down_select"`
	HelpTTest        string `json:"help_t_test"`
	HelpSearchSlash  string `json:"help_search_slash"`
	HelpTranslate    string `json:"help_translate"`
	
	// Theme descriptions
	DarkThemeTokyoNight string `json:"dark_theme_tokyo_night"`
	ChocolateTheme      string `json:"chocolate_theme"`
	LatteTheme          string `json:"latte_theme"`
	DraculaTheme        string `json:"dracula_theme"`
	GruvboxDarkTheme    string `json:"gruvbox_dark_theme"`
	GruvboxLightTheme   string `json:"gruvbox_light_theme"`
	NordTheme           string `json:"nord_theme"`
	SolarizedDarkTheme  string `json:"solarized_dark_theme"`
	SolarizedLightTheme string `json:"solarized_light_theme"`
	MinimalBWTheme      string `json:"minimal_bw_theme"`
	
	// Prompt management additions
	HelpNewPrompt    string `json:"help_new_prompt"`
	HelpEditPrompt   string `json:"help_edit_prompt"`
	HelpDeletePrompt string `json:"help_delete_prompt"`
	ConfirmDeleteKey string `json:"confirm_delete_key"`
	CancelDelete     string `json:"cancel_delete"`
	
	// Status messages
	TestingConnection string `json:"testing_connection"`
	DetectingProvider string `json:"detecting_provider"`
	
	// About page additions
	ProjectAuthor string `json:"project_author"`
	OpenSourceLicense string `json:"open_source_license"`
	AuthorName string `json:"author_name"`
	
	// Key bindings help
	KeyUp string `json:"key_up"`
	KeyDown string `json:"key_down"`
	KeySelect string `json:"key_select"`
	KeyReturn string `json:"key_return"`
	KeyQuit string `json:"key_quit"`
	KeySwitch string `json:"key_switch"`
	KeyEdit string `json:"key_edit"`
	KeyDelete string `json:"key_delete"`
	KeyNew string `json:"key_new"`
	KeyTest string `json:"key_test"`
	
	// Prompt test UI
	TestPromptTitle string `json:"test_prompt_title"`
	CurrentPrompt string `json:"current_prompt"`
	PromptContentLabel string `json:"prompt_content_label"`
	TestText string `json:"test_text"`
	TestingAI string `json:"testing_ai"`
	TranslationResultLabel string `json:"translation_result_label"`
	InputTestText string `json:"input_test_text"`
	ResultWillShowHere string `json:"result_will_show_here"`
	TranslatingText string `json:"translating_text"`
	TabSwitchFocus string `json:"tab_switch_focus"`
	CtrlEnterTest string `json:"ctrl_enter_test"`
	EscReturn string `json:"esc_return"`
	EditingPrompt string `json:"editing_prompt"`
	NewPrompt string `json:"new_prompt"`
	NameLabel string `json:"name_label"`
	ContentLabel string `json:"content_label"`
	SaveKey string `json:"save_key"`
	TestKey string `json:"test_key"`
	CancelKey string `json:"cancel_key"`
	TabSwitchInput string `json:"tab_switch_input"`
	TestPrompt string `json:"test_prompt"`
	UnnamedPrompt string `json:"unnamed_prompt"`
	TranslateToChineseDefault string `json:"translate_to_chinese_default"`
	EmptyInput string `json:"empty_input"`
	NoAPIKeyConfigured string `json:"no_api_key_configured"`
	CreateTranslatorFailed string `json:"create_translator_failed"`
	TestSentenceAI string `json:"test_sentence_ai"`
	UsingModel string `json:"using_model"`
	APINotConfigured string `json:"api_not_configured"`
	
	// Status messages additional
	ConfigRefreshed string `json:"config_refreshed"`
	TranslateOnlyPrompt string `json:"translate_only_prompt"`
	CustomSuffix string `json:"custom_suffix"`
	PreviewLabel string `json:"preview_label"`
	SaveButton string `json:"save_button"`
	NotConfiguredBrackets string `json:"not_configured_brackets"`
	UnknownProvider string `json:"unknown_provider"`
	RecordingHotkey string `json:"recording_hotkey"`
	SetMonitorHotkey string `json:"set_monitor_hotkey"`
	SetSwitchPromptHotkey string `json:"set_switch_prompt_hotkey"`
	PressDesiredHotkey string `json:"press_desired_hotkey"`
	
	// Console messages
	MonitorStartedTray string `json:"monitor_started_tray"`
	MonitorStoppedTray string `json:"monitor_stopped_tray"`
	AutoPasteEnabled string `json:"auto_paste_enabled"`
	AutoPasteDisabled string `json:"auto_paste_disabled"`
	HotkeysLabel string `json:"hotkeys_label"`
	MonitorToggleKey string `json:"monitor_toggle_key"`
	SwitchStyleKey string `json:"switch_style_key"`
	MonitorPausedByHotkey string `json:"monitor_paused_by_hotkey"`
	MonitorResumedByHotkey string `json:"monitor_resumed_by_hotkey"`
	StartingTray string `json:"starting_tray"`
	ControlFromTray string `json:"control_from_tray"`
	GoodbyeEmoji string `json:"goodbye_emoji"`
	DirectTranslation string `json:"direct_translation"`
	TranslateToChineseColon string `json:"translate_to_chinese_colon"`
	
	// API config messages
	NoModelsFound string `json:"no_models_found"`
	CurrentSuffix string `json:"current_suffix"`
	UnrecognizedAPIKey string `json:"unrecognized_api_key"`
	ConnectionFailed string `json:"connection_failed"`
	ConnectionSuccessNoModels string `json:"connection_success_no_models"`
	ConnectionSuccessWithModels string `json:"connection_success_with_models"`
	TestingInProgress string `json:"testing_in_progress"`
	
	// System hotkey
	SystemHotkeyFormat string `json:"system_hotkey_format"`
	SystemHotkeyLabel string `json:"system_hotkey_label"`
	XiaoniaoToggleMonitor string `json:"xiaoniao_toggle_monitor"`
	XiaoniaoSwitchStyle string `json:"xiaoniao_switch_style"`
	
	// Translator error detection
	CannotProceed string `json:"cannot_proceed"`
	AIReturnedMultiline string `json:"ai_returned_multiline"`
	UsingFirstLine string `json:"using_first_line"`
	CannotTranslate string `json:"cannot_translate"`
	UnableToTranslate string `json:"unable_to_translate"`
	Sorry string `json:"sorry"`
	
	// Theme names and descriptions
	DefaultThemeName string `json:"default_theme_name"`
	DefaultThemeDesc string `json:"default_theme_desc"`
	TokyoNightDesc string `json:"tokyo_night_desc"`
	SoftPastelDesc string `json:"soft_pastel_desc"`
	MinimalThemeName string `json:"minimal_theme_name"`
	MinimalThemeDesc string `json:"minimal_theme_desc"`
	
	// Tray messages
	StatusTranslated string `json:"status_translated"`
	DefaultPrompt string `json:"default_prompt"`
	TrayMonitoring string `json:"tray_monitoring"`
	TrayStopped string `json:"tray_stopped"`
	StyleLabel string `json:"style_label"`

	// New fields for missing translations
	ProgramAlreadyRunning string `json:"program_already_running"`
	TrayManagerInitFailed string `json:"tray_manager_init_failed"`
	SystemTrayStartFailed string `json:"system_tray_start_failed"`
	NotConfiguredStatus string `json:"not_configured_status"`
	PleaseConfigureAPIFirst string `json:"please_configure_api_first"`
	APIConfigCompleted string `json:"api_config_completed"`
	MonitorStartedConsole string `json:"monitor_started_console"`
	MonitorPausedConsole string `json:"monitor_paused_console"`
	ExportLogsFailed string `json:"export_logs_failed"`
	LogsExportedTo string `json:"logs_exported_to"`
	ConfigRefreshedDetail string `json:"config_refreshed_detail"`
	RefreshConfigFailed string `json:"refresh_config_failed"`
	SwitchedTo string `json:"switched_to"`
	ConfigRefreshedAndReinit string `json:"config_refreshed_and_reinit"`
	MonitorPausedMsg string `json:"monitor_paused_msg"`
	MonitorResumedMsg string `json:"monitor_resumed_msg"`
	SwitchPromptMsg string `json:"switch_prompt_msg"`
	TranslationStyle string `json:"translation_style"`
	AutoPasteEnabledMsg string `json:"auto_paste_enabled_msg"`
	HotkeysColon string `json:"hotkeys_colon"`
	MonitorToggleLabel string `json:"monitor_toggle_label"`
	SwitchStyleLabel string `json:"switch_style_label"`
	MonitorStartedCopyToTranslate string `json:"monitor_started_copy_to_translate"`
	StartTranslating string `json:"start_translating"`
	UsingPrompt string `json:"using_prompt"`
	TranslationFailedError string `json:"translation_failed_error"`
	TranslationComplete string `json:"translation_complete"`
	OriginalText string `json:"original_text"`
	TranslatedText string `json:"translated_text"`
	MonitorPausedViaHotkey string `json:"monitor_paused_via_hotkey"`
	MonitorResumedViaHotkey string `json:"monitor_resumed_via_hotkey"`
	SwitchPromptViaHotkey string `json:"switch_prompt_via_hotkey"`
	PrewarmingModel string `json:"prewarming_model"`
	PrewarmSuccess2 string `json:"prewarm_success2"`
	PrewarmSkip string `json:"prewarm_skip"`
	TranslatorRefreshed string `json:"translator_refreshed"`
	TranslatorRefreshFailed string `json:"translator_refresh_failed"`
	ConfigRefreshedReinit string `json:"config_refreshed_reinit"`
	MainModelChanged string `json:"main_model_changed"`
	TestingModelMsg string `json:"testing_model_msg"`
	ModelInitFailed string `json:"model_init_failed"`
	TranslateToChineseOnly string `json:"translate_to_chinese_only"`
	ModelTestFailedMsg string `json:"model_test_failed_msg"`
	ModelAvailable string `json:"model_available"`
	ModelNoResponse string `json:"model_no_response"`
	DeleteFailed string `json:"delete_failed"`
	SaveFailed string `json:"save_failed"`
	UpdateFailed string `json:"update_failed"`
	TestingConnectionMsg string `json:"testing_connection_msg"`
	TestingMsg string `json:"testing_msg"`
	CreateTranslatorFailedMsg string `json:"create_translator_failed_msg"`
	TranslationFailedMsg string `json:"translation_failed_msg"`
	TranslationResultMsg string `json:"translation_result_msg"`
	PreviewColon string `json:"preview_colon"`
	HotkeySettingsTitle string `json:"hotkey_settings_title"`
	MonitorToggleHotkey string `json:"monitor_toggle_hotkey"`
	SwitchStyleHotkey string `json:"switch_style_hotkey"`
	AuthorLabel string `json:"author_label"`
	LicenseLabel string `json:"license_label"`
	ProjectUrlLabel string `json:"project_url_label"`
	HotkeysSaved string `json:"hotkeys_saved"`
	EnterYourAPIKey string `json:"enter_your_api_key"`
	DetectedProvider string `json:"detected_provider"`
	UnknownProviderDefault string `json:"unknown_provider_default"`
	Success string `json:"success"`
	SelectAIModel string `json:"select_ai_model"`
	SelectedBrackets string `json:"selected_brackets"`
	PleaseSelectModel string `json:"please_select_model"`
	TestingModelFormat string `json:"testing_model_format"`
	ModelAvailableTranslation string `json:"model_available_translation"`
	ModelUnavailable string `json:"model_unavailable"`
	TestingConnectionDots string `json:"testing_connection_dots"`
	ConnectionFailedFormat string `json:"connection_failed_format"`
	LoadUserPromptsFailed string `json:"load_user_prompts_failed"`
	CreatePromptsJsonFailed string `json:"create_prompts_json_failed"`
	DeleteBuiltinPromptError string `json:"delete_builtin_prompt_error"`
	ExportLogs string `json:"export_logs"`
	StatusTranslatedCount string `json:"status_translated_count"`
	XiaoniaoMonitoring string `json:"xiaoniao_monitoring"`
	XiaoniaoStopped string `json:"xiaoniao_stopped"`
	GetProgramPathFailed string `json:"get_program_path_failed"`
	WriteLogFileFailed string `json:"write_log_file_failed"`
	UnsupportedOS string `json:"unsupported_os"`
	VersionFormat string `json:"version_format"`
	MonitorStatusActiveMsg string `json:"monitor_status_active_msg"`
	MonitorStatusPausedMsg string `json:"monitor_status_paused_msg"`
	TranslationCountMsg string `json:"translation_count_msg"`
	StatusActive string `json:"status_active"`
	StatusPaused string `json:"status_paused"`
	ModelLabel string `json:"model_label"`
	APILabel string `json:"api_label"`
	TryAgainMsg string `json:"try_again_msg"`
	StatsFormat string `json:"stats_format"`
	MainMenu string `json:"main_menu"`
	APIConfiguration string `json:"api_configuration"`
	ThemeSwitcher string `json:"theme_switcher"`
	ViewVersion string `json:"view_version"`
	ExitProgram string `json:"exit_program"`
	ThemesUppercase string `json:"themes_uppercase"`
	AuthorUppercase string `json:"author_uppercase"`
	VersionUppercase string `json:"version_uppercase"`
	PromptLabel string `json:"prompt_label"`
	ModelConfiguration string `json:"model_configuration"`
	ViewModelDetails string `json:"view_model_details"`
	SelectBackupModel string `json:"select_backup_model"`
	ReturnToMainMenu string `json:"return_to_main_menu"`

	// Hotkey screen
	CommonExamples string `json:"common_examples"`
	InputFormat string `json:"input_format"`
	ModifierPlusKey string `json:"modifier_plus_key"`
	SingleModifier string `json:"single_modifier"`
	SingleKey string `json:"single_key"`
	SwitchFunction string `json:"switch_function"`
	Edit string `json:"edit"`
	Save string `json:"save"`
	Back string `json:"back"`

	// Hotkey validation
	FormatError string `json:"format_error"`
	InvalidModifier string `json:"invalid_modifier"`
	InvalidMainKey string `json:"invalid_main_key"`

	// Provider label
	ProviderLabel string `json:"provider_label"`

	DefaultThemeNameFallback string `json:"default_theme_name_fallback"`
	ClassicBlueFallback string `json:"classic_blue_fallback"`
	DarkThemeTokyoNightFallback string `json:"dark_theme_tokyo_night_fallback"`
	SoftPastelFallback string `json:"soft_pastel_fallback"`
	MinimalThemeNameFallback string `json:"minimal_theme_name_fallback"`
	CleanBWFallback string `json:"clean_bw_fallback"`
	TranslateToChineseProvider string `json:"translate_to_chinese_provider"`
}

var (
	currentLang  Language
	translations map[Language]*Translations
)

// Initialize loads all translations
func Initialize(configLang string) {
	// Load all translation files
	loadTranslations()
	
	// Set language based on config or system
	if configLang != "" {
		SetLanguage(Language(configLang))
	} else {
		DetectAndSetLanguage()
	}
}

// loadTranslations loads all translation files
func loadTranslations() {
	translations = make(map[Language]*Translations)
	
	// Load built-in translations
	translations[LangZhCN] = getChineseSimplified()
	translations[LangZhTW] = getChineseTraditional()
	translations[LangEN] = getEnglish()
	translations[LangJA] = getJapanese()
	translations[LangKO] = getKorean()
	translations[LangES] = getSpanish()
	translations[LangFR] = getFrench()
}

// DetectAndSetLanguage detects system language and sets it
func DetectAndSetLanguage() {
	// Check LANG environment variable
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LANGUAGE")
	}
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	
	// Parse language code
	lang = strings.ToLower(lang)
	
	// Match language
	switch {
	case strings.HasPrefix(lang, "zh_cn"), strings.HasPrefix(lang, "zh-cn"):
		currentLang = LangZhCN
	case strings.HasPrefix(lang, "zh_tw"), strings.HasPrefix(lang, "zh-tw"), 
	     strings.HasPrefix(lang, "zh_hk"), strings.HasPrefix(lang, "zh-hk"):
		currentLang = LangZhTW
	case strings.HasPrefix(lang, "en"):
		currentLang = LangEN
	case strings.HasPrefix(lang, "ja"):
		currentLang = LangJA
	case strings.HasPrefix(lang, "ko"):
		currentLang = LangKO
	case strings.HasPrefix(lang, "es"):
		currentLang = LangES
	case strings.HasPrefix(lang, "fr"):
		currentLang = LangFR
	default:
		// Default to English if no match
		currentLang = LangEN
	}
}

// SetLanguage sets the current language
func SetLanguage(lang Language) {
	if _, ok := translations[lang]; ok {
		currentLang = lang
	}
}

// GetLanguage returns the current language
func GetLanguage() Language {
	return currentLang
}

// T returns the translations for the current language
func T() *Translations {
	if trans, ok := translations[currentLang]; ok {
		return trans
	}
	// Fallback to English
	return translations[LangEN]
}

// GetLanguageName returns the display name for a language
func GetLanguageName(lang Language) string {
	names := map[Language]string{
		LangZhCN: "简体中文",
		LangZhTW: "繁體中文",
		LangEN:   "English",
		LangJA:   "日本語",
		LangKO:   "한국어",
		LangES:   "Español",
		LangFR:   "Français",
	}
	
	if name, ok := names[lang]; ok {
		return name
	}
	return string(lang)
}

// GetAvailableLanguages returns all available languages
func GetAvailableLanguages() []Language {
	return []Language{
		LangZhCN,
		LangZhTW,
		LangEN,
		LangJA,
		LangKO,
		LangES,
		LangFR,
	}
}

// LoadCustomTranslation loads a custom translation from JSON file
func LoadCustomTranslation(lang Language, filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	
	var trans Translations
	if err := json.Unmarshal(data, &trans); err != nil {
		return err
	}
	
	translations[lang] = &trans
	return nil
}