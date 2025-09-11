package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
	
	"github.com/kaminoguo/xiaoniao/internal/clipboard"
	"github.com/kaminoguo/xiaoniao/internal/hotkey"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
	"github.com/kaminoguo/xiaoniao/internal/tray"
	"github.com/kaminoguo/xiaoniao/internal/translator"
	"golang.design/x/hotkey/mainthread"
)

const version = "1.6.3"

type Config struct {
	APIKey        string `json:"api_key"`
	Provider      string `json:"provider"`
	Model         string `json:"model"`
	FallbackModel string `json:"fallback_model,omitempty"` // 副模型
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
	// Windows下确保控制台窗口可见（如果需要）
	if len(os.Args) >= 2 && (os.Args[1] == "config" || os.Args[1] == "about" || os.Args[1] == "help" || os.Args[1] == "version") {
		// 对于需要显示输出的命令，确保控制台窗口可见
		showConsoleWindow()
	}
	
	// 无参数时默认执行 run
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "run")
	}
	
	command := os.Args[1]
	
	switch command {
	case "run":
		// Acquire lock for run mode
		if ok, cleanup := acquireLock(); !ok {
			// 显示错误消息（用户可见）
			showErrorMessage("xiaoniao", "程序已在运行中。请检查系统托盘图标。\n如果没有看到托盘图标，请尝试结束所有xiaoniao进程后重新启动。")
			os.Exit(1)
		} else {
			defer cleanup()
		}
		// 需要使用mainthread来支持快捷键
		mainthread.Init(func() {
			runDaemonWithHotkey()
		})
	case "config":
		showConfigUI()
	case "about":
		// 设置环境变量后显示配置界面
		os.Setenv("SHOW_ABOUT", "1")
		showConfigUI()
	case "version", "--version", "-v":
		fmt.Printf("xiaoniao version %s\n", version)
	case "help", "--help", "-h":
		showHelp()
	default:
		t := i18n.T()
		fmt.Printf("%s: %s\n", t.UnknownCommand, command)
		showUsage()
	}
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
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Printf("║       %s v%s       ║\n", t.HelpTitle, version)
	fmt.Printf("║         %s         ║\n", t.HelpDesc)
	fmt.Println("╚════════════════════════════════════════╝")
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

// runDaemonWithHotkey 在主线程运行，支持全局快捷键
func runDaemonWithHotkey() {
	// 初始化托盘管理器
	trayManager, err := tray.NewManager()
	if err != nil {
		showErrorMessage("xiaoniao 启动失败", fmt.Sprintf("托盘管理器初始化失败：%v\n\n请检查系统是否支持系统托盘功能。", err))
		return
	}
	
	// Windows需要在主线程中运行systray
	// 设置业务逻辑回调到托盘管理器的onReady中
	trayManager.SetBusinessLogic(func() {
		runDaemonBusinessLogic(trayManager)
	})
	
	// 直接在主线程中启动托盘（这是阻塞调用）
	if err := trayManager.Initialize(); err != nil {
		showErrorMessage("xiaoniao 启动失败", fmt.Sprintf("系统托盘启动失败：%v\n\n可能的原因：\n1. 系统托盘功能被禁用\n2. 权限不足\n3. 系统资源不足\n\n请检查系统设置并重试。", err))
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
		trayManager.SetCurrentPrompt("未配置 / Not Configured")
		
		// 设置托盘回调 - 只允许打开设置
		trayManager.SetOnSettings(func() {
			openConfigInTerminal()
			go watchConfig()
		})
		
		trayManager.SetOnToggleMonitor(func(enabled bool) {
			fmt.Println("请先配置 API / Please configure API first")
		})
		
		trayManager.SetOnQuit(func() {
			os.Exit(0)
		})
		
		// 自动打开配置界面
		go func() {
			time.Sleep(500 * time.Millisecond)
			showConsoleWindow() // 确保控制台窗口可见
			showConfigUI()
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
					fmt.Println("\n✅ API配置已完成，重新启动翻译服务...")
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
			FallbackModel: config.FallbackModel,
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
				fmt.Println("\n✅ 监控已通过托盘启动")
			} else {
				monitor.Stop()
				fmt.Println("\n⏸️ 监控已通过托盘停止")
			}
		})
	
	
		trayManager.SetOnSettings(func() {
			// 打开配置界面
			showConsoleWindow() // 确保控制台窗口可见
			showConfigUI()
			// 启动配置文件监控
			go watchConfig()
		})
	
		trayManager.SetOnToggleTerminal(func() {
			// 切换终端窗口显示/隐藏
			toggleTerminalVisibility()
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
				FallbackModel: config.FallbackModel,
				MaxRetries:    3,
				Timeout:      60,
			}
		
			newTrans, err := translator.NewTranslator(translatorConfig)
			if err == nil {
				trans = newTrans
				fmt.Printf("\n✅ 配置已刷新: %s | %s | %s\n", 
					config.Provider, config.Model, getPromptName(config.PromptID))
				
				// 如果切换了模型或Provider，进行预热
				if config.Model != oldModel || config.Provider != oldProvider {
					go prewarmModel(trans)
				}
			} else {
				fmt.Printf("\n❌ 刷新配置失败: %v\n", err)
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
			fmt.Printf("\n切换到: %s\n", promptName)
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
				fmt.Println("\n⏸ 监控已暂停")
				monitoring = false
			} else {
				monitor.Start()
				trayManager.UpdateMonitorStatus(true)
				fmt.Println("\n▶ 监控已恢复")
				monitoring = true
			}
		})
		if err != nil {
			fmt.Printf("⚠️ 无法注册快捷键 %s: %v\n", config.HotkeyToggle, err)
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
			fmt.Printf("\n🔄 切换Prompt: %s\n", promptName)
			trayManager.SetCurrentPrompt(promptName)
			// 不弹窗通知
		})
		if err != nil {
			fmt.Printf("⚠️ 无法注册快捷键 %s: %v\n", config.HotkeySwitch, err)
		}
	}
	
	clearScreen()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Printf("║            xiaoniao - %s           ║\n", t.Running)
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("%s: %s | %s: %s\n", t.Provider, config.Provider, t.Model, config.Model)
	fmt.Printf("%s: %s\n", t.TranslateStyle, getPromptName(config.PromptID))
	fmt.Printf("%s: ✅ %s\n", t.AutoPaste, t.Enabled)
	
	// 显示快捷键信息
	if config.HotkeyToggle != "" || config.HotkeySwitch != "" {
		fmt.Printf("%s\n", t.HotkeysLabel)
		if config.HotkeyToggle != "" {
			fmt.Printf("  %s\n", fmt.Sprintf(t.MonitorToggleKey, config.HotkeyToggle))
		}
		if config.HotkeySwitch != "" {
			fmt.Printf("  %s\n", fmt.Sprintf(t.SwitchStyleKey, config.HotkeySwitch))
		}
	}
	
	fmt.Println("─────────────────────────────────────────")
	fmt.Println(t.Monitoring)
	fmt.Println(t.CopyToTranslate)
	fmt.Println(t.Step5)
	fmt.Println(t.ExitTip)
	fmt.Println("─────────────────────────────────────────")
	
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
		
		fmt.Printf("\n[%s] %s", time.Now().Format("15:04:05"), t.Translating)
		trayManager.SetStatus(tray.StatusTranslating)
		
		// 每次翻译前重新获取prompt（以防配置文件被修改）
		currentPrompt := getPromptContent(config.PromptID)
		fmt.Printf("\n开始翻译: %s\n", text)
		fmt.Printf("使用Prompt: %s (内容长度: %d)\n", config.PromptID, len(currentPrompt))
		
		// 执行翻译
		result, err := trans.Translate(text, currentPrompt)
		if err != nil {
			fmt.Printf(" ❌ %s: %v\n", t.Failed, err)
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
			
			fmt.Printf(" ✅ %s #%d\n", t.Complete, translationCount)
			trayManager.IncrementTranslationCount()
			trayManager.SetStatus(tray.StatusIdle)
			fmt.Printf("   %s: %s\n", t.Original, truncate(text, 50))
			fmt.Printf("   %s: %s\n", t.Translation, truncate(result.Translation, 50))
			
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
					fmt.Println("\n⏸ 监控已暂停 (通过快捷键)")
					monitoring = false
				} else {
					monitor.Start()
					trayManager.UpdateMonitorStatus(true)
					fmt.Println("\n▶ 监控已恢复 (通过快捷键)")
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
					fmt.Printf("\n🔄 切换Prompt: %s (通过快捷键)\n", promptName)
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
	fmt.Print("\033[H\033[2J")
}

// simulatePaste 模拟粘贴操作 (Windows实现)
func simulatePaste() {
	// Windows实现自动粘贴功能
	// 可以使用PowerShell或Windows API实现
	// 目前留空，可以在后续添加Windows特定的实现
	// 例如：使用PowerShell发送Ctrl+V
	// exec.Command("powershell", "-Command", "Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.SendKeys]::SendWait('^v')").Run()
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
	// 简单的默认prompt内容
	switch id {
	case "direct":
		return "Please translate the following text to Chinese:"
	default:
		return "Translate the following to Chinese:"
	}
}

var terminalVisible = false  // Start as false when running in background
var terminalPID = 0  // PID of the log viewer terminal

// hideTerminal 隐藏日志查看终端窗口 (Windows实现)
func hideTerminal() {
	if !terminalVisible || terminalPID == 0 {
		return
	}
	
	// Windows: Kill the log viewer terminal
	if terminalPID > 0 {
		exec.Command("taskkill", "/PID", strconv.Itoa(terminalPID)).Run()
		terminalPID = 0
	}
	terminalVisible = false
}

// showTerminal 显示日志查看终端窗口 (Windows实现)
func showTerminal() {
	if terminalVisible {
		return
	}
	
	// Windows: Open Command Prompt with tail equivalent
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "powershell Get-Content /tmp/xiaoniao.log -Wait")
	if err := cmd.Start(); err == nil {
		terminalPID = cmd.Process.Pid
		terminalVisible = true
	}
}

// toggleTerminalVisibility 切换日志查看终端的显示/隐藏状态
func toggleTerminalVisibility() {
	// 切换显示/隐藏日志查看终端
	if terminalVisible {
		hideTerminal()
	} else {
		showTerminal()
	}
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
				
				// 如果模型或提供商变了，打印提示
				if config.Model != oldModel || config.Provider != oldProvider {
					fmt.Printf("\n🔄 配置已更新: %s | %s\n", config.Provider, config.Model)
				}
			}
		}
	}
}

// prewarmModel 预热模型
func prewarmModel(trans *translator.Translator) {
	fmt.Print("预热模型中...")
	err := translator.PrewarmConnection(trans)
	if err == nil {
		fmt.Println(" ✅")
	} else {
		// 预热失败不影响使用，只是警告
		fmt.Printf(" ⚠️ (可忽略: %v)\n", err)
	}
}

// monitorRefreshSignal 监控刷新信号文件
func monitorRefreshSignal(trans **translator.Translator) {
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
				FallbackModel: config.FallbackModel,
				MaxRetries:    3,
				Timeout:      60,
			}
			
			newTrans, err := translator.NewTranslator(translatorConfig)
			if err == nil {
				*trans = newTrans
				fmt.Printf("\n✅ 翻译器已刷新: %s | %s\n", config.Provider, config.Model)
				
				// 检查是否切换了模型或Provider，如果是则预热
				if config.Model != lastModel || config.Provider != lastProvider {
					go prewarmModel(newTrans)
					lastModel = config.Model
					lastProvider = config.Provider
				}
			} else {
				fmt.Printf("\n❌ 翻译器刷新失败: %v\n", err)
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
	// 简单启动配置界面
	showConfigUI()
}



// 删除重复定义的函数，使用prompts.go和config_ui.go中的实现