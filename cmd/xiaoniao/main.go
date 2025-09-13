package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	
	"github.com/kaminoguo/xiaoniao/internal/clipboard"
	"github.com/kaminoguo/xiaoniao/internal/hotkey"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
	"github.com/kaminoguo/xiaoniao/internal/logbuffer"
	"github.com/kaminoguo/xiaoniao/internal/tray"
	"github.com/kaminoguo/xiaoniao/internal/translator"
	"golang.design/x/hotkey/mainthread"
)

const version = "1.0"

type Config struct {
	APIKey        string `json:"api_key"`
	Provider      string `json:"provider"`
	Model         string `json:"model"`
	PromptID      string `json:"prompt_id"`
	Language      string `json:"language,omitempty"`
	Theme         string `json:"theme,omitempty"`      // UI主题
	HotkeyToggle  string `json:"hotkey_toggle,omitempty"`  // 监控开关快捷键
	HotkeySwitch  string `json:"hotkey_switch,omitempty"`  // 切换prompt快捷键
}

var (
	configPath string
	config     Config
)

func init() {
	// 获取配置目录
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}
	appDir := filepath.Join(configDir, "xiaoniao")
	os.MkdirAll(appDir, 0755)
	configPath = filepath.Join(appDir, "config.json")
	
	// 加载配置
	loadConfig()
	
	// 初始化i18n
	i18n.Initialize(config.Language)
}

// acquireLock creates a lock file to prevent multiple instances
func acquireLock() (bool, func()) {
	configDir, _ := os.UserConfigDir()
	appDir := filepath.Join(configDir, "xiaoniao")
	lockFile := filepath.Join(appDir, "xiaoniao.lock")
	
	// Check if lock file exists and if the process is still running
	if data, err := os.ReadFile(lockFile); err == nil {
		if pid, err := strconv.Atoi(string(data)); err == nil {
			// Check if process is still running
			if process, err := os.FindProcess(pid); err == nil {
				if err := process.Signal(syscall.Signal(0)); err == nil {
					// Process is still running
					return false, nil
				}
			}
		}
		// Process is not running, remove stale lock file
		os.Remove(lockFile)
	}
	
	// Create lock file with current PID
	pid := os.Getpid()
	if err := os.WriteFile(lockFile, []byte(strconv.Itoa(pid)), 0644); err != nil {
		return false, nil
	}
	
	// Return cleanup function
	cleanup := func() {
		os.Remove(lockFile)
	}
	
	return true, cleanup
}


func main() {
	// 开始捕获所有控制台输出到日志缓冲区
	logbuffer.CaptureStdout()
	
	// 只有在没有参数（主进程）时才隐藏控制台
	// config、about 等子命令不应该隐藏控制台
	if len(os.Args) == 1 {
		// 延迟一小段时间让Windows Terminal完全初始化
		// 然后隐藏控制台窗口
		go func() {
			time.Sleep(100 * time.Millisecond)
			hideConsoleWindow()
		}()
	}
	
	// Handle special commands
	if len(os.Args) >= 2 && (os.Args[1] == "config" || os.Args[1] == "about" || os.Args[1] == "tutorial" || os.Args[1] == "help" || os.Args[1] == "version") {

		command := os.Args[1]
		switch command {
		case "config":
			showConfigUI()
		case "about":
			os.Setenv("SHOW_ABOUT", "1")
			showConfigUI()
		case "tutorial":
			os.Setenv("SHOW_TUTORIAL", "1")
			showConfigUI()
		case "version", "--version", "-v":
			fmt.Printf("xiaoniao version %s\n", version)
		case "help", "--help", "-h":
			showHelp()
		}
		
		// Wait for user input before closing console
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln()
		return
	}
	
	// GUI mode - hide console window on startup
	hideConsoleWindow()
	// Set up console control handler to prevent program exit on X click (Windows only)
	setupMainConsoleHandler()
	// Acquire lock for run mode
	if ok, cleanup := acquireLock(); !ok {
		// Show error message using Windows message box
		t := i18n.T()
		showErrorMessage("xiaoniao", t.ProgramAlreadyRunning)
		os.Exit(1)
	} else {
		defer cleanup()
	}
	
	
	// Run GUI mode - this is blocking and keeps the app running
	mainthread.Init(func() {
		runDaemonWithHotkey()
	})
}

func showUsage() {
	t := i18n.T()
	fmt.Printf("%s: xiaoniao <%s>\n", t.Usage, t.Commands)
	fmt.Println()
	fmt.Printf("%s:\n", t.Commands)
	fmt.Printf("  run     - %s\n", t.RunDesc)
	fmt.Printf("  config  - %s\n", t.ConfigDesc)
	fmt.Printf("  version - %s\n", t.VersionDesc)
	fmt.Printf("  help    - %s\n", t.HelpDesc2)
}

func showHelp() {
	t := i18n.T()
	fmt.Printf("\n%s v%s\n", t.HelpTitle, version)
	fmt.Printf("%s\n", t.HelpDesc)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println()
	fmt.Printf("%s:\n", t.Commands)
	fmt.Printf("  %s\n", t.RunCommand)
	fmt.Printf("    %s\n", t.RunDesc)
	fmt.Println("    ")
	fmt.Printf("  %s\n", t.ConfigCommand)
	fmt.Printf("    %s\n", t.ConfigDesc)
	fmt.Println("    ")
	fmt.Printf("  %s\n", t.HelpCommand)
	fmt.Printf("    %s\n", t.HelpDesc2)
	fmt.Println()
	fmt.Printf("%s:\n", t.HowItWorks)
	fmt.Printf("  1. %s\n", t.Step1)
	fmt.Printf("  2. %s\n", t.Step2)
	fmt.Printf("  3. %s\n", t.Step3)
	fmt.Printf("  4. %s\n", t.Step4)
	fmt.Printf("  5. %s\n", t.Step5)
	fmt.Println()
	fmt.Println(t.Warning)
}

// runDaemonWithHotkey 在主线程运行，支持全局快捷键（控制台模式）
func runDaemonWithHotkey() {
	// 初始化托盘管理器
	trayManager, err := tray.NewManager()
	if err != nil {
		t := i18n.T()
		showErrorMessage("xiaoniao", fmt.Sprintf(t.TrayManagerInitFailed, err))
		return
	}
	
	// Windows需要在主线程中运行systray
	// 设置业务逻辑回调到托盘管理器的onReady中
	trayManager.SetBusinessLogic(func() {
		runDaemonBusinessLogic(trayManager)
	})
	
	// 直接在主线程中启动托盘（这是阻塞调用）
	if err := trayManager.Initialize(); err != nil {
		t := i18n.T()
		showErrorMessage("xiaoniao", fmt.Sprintf(t.SystemTrayStartFailed, err))
		return
	}
}

// runDaemonBusinessLogic 运行守护进程的业务逻辑
// trayManager 必须已经初始化
func runDaemonBusinessLogic(trayManager *tray.Manager) {
	// 检查配置
	t := i18n.T()
	
	// 初始化变量
	var trans *translator.Translator
	var monitor *clipboard.Monitor
	translationCount := 0
	
	// 如果没有 API 配置
	if config.APIKey == "" {
		fmt.Println(t.NoAPIKey)
		fmt.Println(t.OpeningConfig)
		
		// 设置托盘为未配置状态
		trayManager.SetCurrentPrompt(t.NotConfiguredStatus)
		
		// 设置托盘回调 - 只允许打开设置
		trayManager.SetOnSettings(func() {
			openConfigInTerminal()
			go watchConfig()
		})
		
		trayManager.SetOnToggleMonitor(func(enabled bool) {
			fmt.Println(t.PleaseConfigureAPIFirst)
		})
		
		trayManager.SetOnQuit(func() {
			os.Exit(0)
		})
		
		// 自动在新窗口中打开配置界面
		go func() {
			time.Sleep(500 * time.Millisecond)
			openConfigInTerminal()
		}()
		
		// 设置等待状态，让托盘保持运行
		go func() {
			// 持续监控配置文件变化
			for {
				time.Sleep(2 * time.Second)
				oldAPIKey := config.APIKey
				loadConfig()
				if config.APIKey != "" && config.APIKey != oldAPIKey {
					// API配置完成，重新初始化业务逻辑
					fmt.Println("\n" + t.APIConfigCompleted)
					go runDaemonBusinessLogic(trayManager)
					return
				}
			}
		}()
		
		return // 返回但保持托盘运行
	} else {
		// 有 API 配置，执行正常的初始化
		
		// 确保加载最新的用户prompts
		ReloadPrompts()
		
		// 初始化翻译器
		translatorConfig := &translator.Config{
			APIKey:        config.APIKey,
			Provider:      config.Provider,
			Model:         config.Model,
			MaxRetries:    3,
			Timeout:      60,  // 增加到60秒
		}
	
		var err error
		trans, err = translator.NewTranslator(translatorConfig)
		if err != nil {
			fmt.Printf("%s: %v\n", t.InitFailed, err)
			return
		}
	
		// 预热模型（异步执行，不阻塞启动）
		go prewarmModel(trans)
	
		// 启动刷新信号监控
		go monitorRefreshSignal(&trans)
	
		// 初始化剪贴板监控
		monitor = clipboard.NewMonitor()
	
		// 设置当前 prompt 显示
		promptName := getPromptName(config.PromptID)
		trayManager.SetCurrentPrompt(promptName)
	
		// 设置托盘回调
		trayManager.SetOnToggleMonitor(func(enabled bool) {
			if enabled {
				monitor.Start()
				fmt.Println("\n" + t.MonitorStartedConsole)
			} else {
				monitor.Stop()
				fmt.Println("\n" + t.MonitorPausedConsole)
			}
		})
	
	trayManager.SetOnSettings(func() {
		// 在新窗口中打开配置界面
		exePath, err := os.Executable()
		if err != nil {
			exePath = "xiaoniao.exe"
		} else {
			if filepath.Ext(exePath) == "" {
				exePath = exePath + ".exe"
			}
		}
		
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", exePath, "config")
		err = cmd.Start()
		if err != nil {
			// 如果启动失败，尝试使用绝对路径
			if absPath, absErr := filepath.Abs(exePath); absErr == nil {
				cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", absPath, "config")
				cmd.Start()
			}
		}
		// 启动配置文件监控
		go watchConfig()
	})
	trayManager.SetOnToggleDebugConsole(func() {
		// 导出日志到文件
		filePath, err := logbuffer.ExportLogs()
		if err != nil {
			fmt.Printf(t.ExportLogsFailed+"\n", err)
		} else {
			fmt.Printf(t.LogsExportedTo+"\n", filePath)
		}
	})
	
	trayManager.SetOnRefresh(func() {
			oldModel := config.Model
			oldProvider := config.Provider
			oldPrompt := config.PromptID
		
			// 重新加载配置
			loadConfig()
		
			// 更新 prompt 显示
			if config.PromptID != oldPrompt {
				promptName := getPromptName(config.PromptID)
				trayManager.SetCurrentPrompt(promptName)
			}
		
			// 重新创建 translator
			translatorConfig := &translator.Config{
				APIKey:        config.APIKey,
				Provider:      config.Provider,
				Model:         config.Model,
					MaxRetries:    3,
				Timeout:      60,
			}
		
			newTrans, err := translator.NewTranslator(translatorConfig)
			if err == nil {
				trans = newTrans
				fmt.Printf("\n"+t.ConfigRefreshedDetail+"\n", 
					config.Provider, config.Model, getPromptName(config.PromptID))
				
				// 如果切换了模型或Provider，进行预热
				if config.Model != oldModel || config.Provider != oldProvider {
					go prewarmModel(trans)
				}
			} else {
				fmt.Printf("\n"+t.RefreshConfigFailed+"\n", err)
			}
		})
	
		// 设置prompt选择回调
		trayManager.SetOnSelectPrompt(func(promptID string) {
			// 更新配置
			config.PromptID = promptID
			
			// 保存配置
			saveConfig()
		
			// 获取prompt名称
			var promptName string
			for _, p := range GetAllPrompts() {
				if p.ID == promptID {
					promptName = p.Name
					break
				}
			}
		
			// 显示提示
			fmt.Printf("\n"+t.SwitchedTo+"\n", promptName)
			trayManager.SetCurrentPrompt(promptName)
			// 不显示通知，只在终端显示
		})
	
		trayManager.SetOnQuit(func() {
			monitor.Stop()
			fmt.Printf("\n%s %d %s\n", t.TotalCount, translationCount, t.TranslateCount)
			os.Exit(0)
		})
	
	
		// 更新prompt列表到菜单（托盘初始化后）
		prompts := GetAllPrompts()
		promptList := make([]struct{ ID, Name string }, len(prompts))
		for i, p := range prompts {
			promptList[i] = struct{ ID, Name string }{ID: p.ID, Name: p.Name}
		}
		trayManager.UpdatePromptList(promptList)
		
		// 初始化终端菜单状态（程序启动时控制台被隐藏）
		trayManager.UpdateDebugConsoleMenu(false)
	
		// 创建快捷键管理器
		hotkeyManager := hotkey.NewManager()
	
	// 注册快捷键（如果配置了）
	if config.HotkeyToggle != "" {
		monitoring := true // 跟踪监控状态
		err := hotkeyManager.RegisterFromString("toggle", config.HotkeyToggle, func() {
			// 切换监控状态
			if monitoring {
				monitor.Stop()
				trayManager.UpdateMonitorStatus(false)
				fmt.Println("\n" + t.MonitorPausedMsg)
				monitoring = false
			} else {
				monitor.Start()
				trayManager.UpdateMonitorStatus(true)
				fmt.Println("\n" + t.MonitorResumedMsg)
				monitoring = true
			}
		})
		if err != nil {
			// fmt.Printf("警告: 无法注册快捷键 %s: %v\n", config.HotkeyToggle, err)
		}
		}
		
		if config.HotkeySwitch != "" {
		err := hotkeyManager.RegisterFromString("switch", config.HotkeySwitch, func() {
			// 切换到下一个Prompt
			prompts := GetAllPrompts()
			if len(prompts) == 0 {
				return
			}
			
			// 找到当前prompt的索引
			currentIdx := -1
			for i, p := range prompts {
				if p.ID == config.PromptID {
					currentIdx = i
					break
				}
			}
			
			// 切换到下一个
			nextIdx := (currentIdx + 1) % len(prompts)
			config.PromptID = prompts[nextIdx].ID
			saveConfig()
			
			// 显示通知
			promptName := prompts[nextIdx].Name
			fmt.Printf("\n" + t.SwitchPromptMsg + "\n", promptName)
			trayManager.SetCurrentPrompt(promptName)
			// 不弹窗通知
		})
		if err != nil {
			// fmt.Printf("警告: 无法注册快捷键 %s: %v\n", config.HotkeySwitch, err)
		}
	}
	
	// Console mode startup info
	fmt.Printf("\nxiaoniao v%s\n", version)
	fmt.Printf(t.ProviderLabel+"%s | "+t.ModelLabel+"%s\n", config.Provider, config.Model)
	fmt.Printf(t.TranslationStyle+"\n", getPromptName(config.PromptID))
	fmt.Printf(t.AutoPasteEnabledMsg)
	
	// 记录快捷键信息
	if config.HotkeyToggle != "" || config.HotkeySwitch != "" {
		fmt.Printf("\n"+t.HotkeysColon+"\n")
		if config.HotkeyToggle != "" {
			fmt.Printf(t.MonitorToggleLabel+"\n", config.HotkeyToggle)
		}
		if config.HotkeySwitch != "" {
			fmt.Printf(t.SwitchStyleLabel+"\n", config.HotkeySwitch)
		}
	}
	
	fmt.Println("\n" + t.MonitorStartedCopyToTranslate)
	
	// 不播放启动提示音
	// sound.PlayStart()
	
	// 更新托盘状态（只有在有 API 配置时才更新）
	if config.APIKey != "" {
		trayManager.UpdateMonitorStatus(true)
	}
	
	monitor.SetOnChange(func(text string) {
		if text == "" {
			return
		}
		
		fmt.Printf("\n[%s] %s...", time.Now().Format("15:04:05"), t.Translating)
		trayManager.SetStatus(tray.StatusTranslating)
		
		// 每次翻译前重新获取prompt（以防配置文件被修改）
		currentPrompt := getPromptContent(config.PromptID)
		fmt.Printf("\n"+t.StartTranslating+"\n", text)
		fmt.Printf(t.UsingPrompt+"\n", config.PromptID, len(currentPrompt))
		
		// 执行翻译
		result, err := trans.Translate(text, currentPrompt)
		if err != nil {
			fmt.Printf(t.TranslationFailedError+"\n", err)
			// sound.PlayError() // 错误提示音已禁用
			trayManager.SetStatus(tray.StatusError)
			// 3秒后恢复正常状态
			go func() {
				time.Sleep(3 * time.Second)
				trayManager.SetStatus(tray.StatusIdle)
			}()
			return
		}
		
		if result.Success && result.Translation != "" {
			// 记录译文，避免重复翻译
			monitor.SetLastTranslation(result.Translation)
			
			// 替换剪贴板
			clipboard.SetClipboard(result.Translation)
			translationCount++
			
			fmt.Printf(t.TranslationComplete+"\n", translationCount)
			trayManager.IncrementTranslationCount()
			trayManager.SetStatus(tray.StatusIdle)
			fmt.Printf(t.OriginalText+"\n", truncate(text, 60))
			fmt.Printf(t.TranslatedText+"\n", truncate(result.Translation, 60))
			
			// 自动粘贴
			{
				go func() {
					// 稍微延迟，确保剪贴板已更新
					time.Sleep(100 * time.Millisecond)
					simulatePaste()
				}()
			}
			
			// sound.PlaySuccess() // 成功提示音已禁用
			
		}
	})
	
	// 开始监控
	monitor.Start()
	
	// 监控状态
	monitoring := true
	
	// 在goroutine中处理信号
	go func() {
		sigChan := make(chan os.Signal, 1)
		setupSignalHandlers(sigChan)
		
		for sig := range sigChan {
			action := handleSignal(sig)
			switch action {
			case "toggle_monitor":
				// 切换监控状态
				if monitoring {
					monitor.Stop()
					trayManager.UpdateMonitorStatus(false)
					fmt.Println("\n" + t.MonitorPausedViaHotkey)
					monitoring = false
				} else {
					monitor.Start()
					trayManager.UpdateMonitorStatus(true)
					fmt.Println("\n" + t.MonitorResumedViaHotkey)
					monitoring = true
				}
				
			case "toggle_prompt":
				// 切换到下一个Prompt
				prompts := GetAllPrompts()
				if len(prompts) > 0 {
					currentIdx := -1
					for i, p := range prompts {
						if p.ID == config.PromptID {
							currentIdx = i
							break
						}
					}
					
					nextIdx := (currentIdx + 1) % len(prompts)
					config.PromptID = prompts[nextIdx].ID
					saveConfig()
					
					promptName := prompts[nextIdx].Name
					fmt.Printf("\n" + t.SwitchPromptViaHotkey + "\n", promptName)
					trayManager.SetCurrentPrompt(promptName)
					// 只在终端显示，不弹窗
				}
				
			case "exit":
				// 退出程序
				monitor.Stop()
				trayManager.Quit()
				fmt.Printf("\n\n%s %d %s\n", t.TotalCount, translationCount, t.TranslateCount)
				fmt.Println(t.Goodbye)
				os.Exit(0)
			}
		}
	}()
	
	} // else 块结束
	
}


// 辅助函数

func clearScreen() {
	// GUI mode: no need to clear screen, function kept for compatibility
}


func loadConfig() {
	data, err := os.ReadFile(configPath)
	if err == nil {
		json.Unmarshal(data, &config)
	}
	
	// 设置默认值
	if config.Provider == "" {
		config.Provider = "OpenAI"
	}
	if config.Model == "" {
		config.Model = "gpt-4o-mini"
	}
	if config.PromptID == "" {
		config.PromptID = "direct"
	}
}

func saveConfig() {
	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(configPath, data, 0644)
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func getPromptName(id string) string {
	prompts := GetAllPrompts()
	for _, p := range prompts {
		if p.ID == id {
			return p.Name
		}
	}
	return "Unknown"
}

func getPromptContent(id string) string {
	// 从文件中获取实际的prompt内容
	if prompt := GetPromptByID(id); prompt != nil {
		return prompt.Content
	}
	// 如果找不到prompt，返回空字符串，不使用默认值
	return ""
}

// toggleTerminalVisibility 切换终端窗口的显示/隐藏状态 (deprecated, replaced by debug console)
func toggleTerminalVisibility() {
	// Legacy function - now replaced by debug console functionality
	// This function is kept for compatibility but does nothing
}



// watchConfig 监控配置文件变化
func watchConfig() {
	lastMod := time.Now()
	for i := 0; i < 60; i++ { // 监控60秒
		time.Sleep(1 * time.Second)
		
		if stat, err := os.Stat(configPath); err == nil {
			if stat.ModTime().After(lastMod) {
				lastMod = stat.ModTime()
				oldModel := config.Model
				oldProvider := config.Provider
				
				loadConfig()
				
				// 如果模型或提供商变了，记录提示
				if config.Model != oldModel || config.Provider != oldProvider {
					fmt.Printf("\n🔄 配置已更新: %s | %s\n", config.Provider, config.Model)
				}
			}
		}
	}
}

// prewarmModel 预热模型
func prewarmModel(trans *translator.Translator) {
	t := i18n.T()
	fmt.Print(t.PrewarmingModel)
	err := translator.PrewarmConnection(trans)
	if err == nil {
		fmt.Println(t.PrewarmSuccess2)
	} else {
		// 预热失败不影响使用，只是警告
		fmt.Printf(t.PrewarmSkip+"\n", err)
	}
}

// monitorRefreshSignal 监控刷新信号文件
func monitorRefreshSignal(trans **translator.Translator) {
	t := i18n.T()
	homeDir, _ := os.UserHomeDir()
	signalPath := filepath.Join(homeDir, ".config", "xiaoniao", ".refresh_signal")

	var lastModel string = config.Model
	var lastProvider string = config.Provider
	
	for {
		time.Sleep(1 * time.Second)
		
		// 检查信号文件是否存在
		if _, err := os.Stat(signalPath); err == nil {
			// 删除信号文件
			os.Remove(signalPath)
			
			// 重新加载配置
			loadConfig()
			
			// 重新创建翻译器
			translatorConfig := &translator.Config{
				APIKey:        config.APIKey,
				Provider:      config.Provider,
				Model:         config.Model,
					MaxRetries:    3,
				Timeout:      60,
			}
			
			newTrans, err := translator.NewTranslator(translatorConfig)
			if err == nil {
				*trans = newTrans
				fmt.Printf("\n"+t.TranslatorRefreshed+"\n", config.Provider, config.Model)
				
				// 检查是否切换了模型或Provider，如果是则预热
				if config.Model != lastModel || config.Provider != lastProvider {
					go prewarmModel(newTrans)
					lastModel = config.Model
					lastProvider = config.Provider
				}
			} else {
				fmt.Printf("\n"+t.TranslatorRefreshFailed+"\n", err)
			}
		}
	}
}

// setupSignalHandlers 设置信号处理器 (跨平台版本)
func setupSignalHandlers(sigChan chan os.Signal) {
	// 跨平台支持的信号
	signal.Notify(sigChan, 
		os.Interrupt,    // Ctrl+C
		syscall.SIGTERM, // 终止信号
	)
}

// handleSignal 处理信号 (跨平台版本)
func handleSignal(sig os.Signal) string {
	switch sig {
	case os.Interrupt, syscall.SIGTERM:
		return "exit"
	default:
		return ""
	}
}

// openConfigInTerminal 在终端中打开配置界面
func openConfigInTerminal() {
	// 在新窗口中打开配置界面
	exePath, err := os.Executable()
	if err != nil {
		exePath = "xiaoniao.exe"
	} else {
		if filepath.Ext(exePath) == "" {
			exePath = exePath + ".exe"
		}
	}
	
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", exePath, "config")
	err = cmd.Start()
	if err != nil {
		// 如果启动失败，尝试使用绝对路径
		if absPath, absErr := filepath.Abs(exePath); absErr == nil {
			cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", absPath, "config")
			cmd.Start()
		}
	}
}



// 删除重复定义的函数，使用prompts.go和config_ui.go中的实现