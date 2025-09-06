package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"pixel-translator/internal/clipboard"
	"pixel-translator/internal/hotkey"
	"pixel-translator/internal/i18n"
	"pixel-translator/internal/tray"
	"pixel-translator/internal/translator"
	"runtime"
	"strconv"
	"syscall"
	"time"
	
	"golang.design/x/hotkey/mainthread"
)

const version = "1.4.1"

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
	if len(os.Args) < 2 {
		showUsage()
		return
	}
	
	command := os.Args[1]
	
	switch command {
	case "run":
		// Acquire lock for run mode
		if ok, cleanup := acquireLock(); !ok {
			fmt.Println(i18n.T().AlreadyRunning)
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
	// 先执行原有的初始化逻辑
	runDaemonCore()
	
	// 保持主线程运行（mainthread需要）
	select {}
}

func runDaemonCore() {
	// 原runDaemon的全部逻辑，但不包含最后的阻塞
	runDaemonInternal()
}

// runDaemon 保留用于兼容（不使用快捷键时调用）
func runDaemon() {
	runDaemonInternal()
	// 阻塞等待
	select {}
}

func runDaemonInternal() {
	// 检查配置
	t := i18n.T()
	if config.APIKey == "" {
		fmt.Println(t.NoAPIKey)
		fmt.Println(t.OpeningConfig)
		
		// 使用和托盘图标相同的方法打开配置
		openConfigInTerminal()
		
		// 等待一下，避免程序立即退出
		time.Sleep(2 * time.Second)
		return
	}
	
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
	
	trans, err := translator.NewTranslator(translatorConfig)
	if err != nil {
		fmt.Printf("%s: %v\n", t.InitFailed, err)
		return
	}
	
	// 预热模型（异步执行，不阻塞启动）
	go prewarmModel(trans)
	
	// 启动刷新信号监控
	go monitorRefreshSignal(&trans)
	
	// 初始化剪贴板监控（提前创建，供托盘使用）
	monitor := clipboard.NewMonitor()
	translationCount := 0
	
	// 创建托盘图标
	trayManager := tray.NewManager()
	
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
		// 在新终端窗口中打开配置界面
		openConfigInTerminal()
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
	
	// 在后台启动托盘
	trayStarted := make(chan bool)
	go func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
			trayStarted <- true
		}()
		trayManager.Initialize()
	}()
	
	// 等待托盘初始化
	<-trayStarted
	
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
			prompts := loadAllPrompts()
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
	
	// 更新托盘状态
	trayManager.UpdateMonitorStatus(true)
	
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
		signal.Notify(sigChan, 
			syscall.SIGINT,   // Ctrl+C
			syscall.SIGTERM,  // 终止信号
			syscall.SIGUSR1,  // 切换监控
			syscall.SIGUSR2,  // 切换Prompt
		)
		
		for sig := range sigChan {
			switch sig {
			case syscall.SIGUSR1:
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
				
			case syscall.SIGUSR2:
				// 切换到下一个Prompt
				prompts := loadAllPrompts()
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
				
			case syscall.SIGINT, syscall.SIGTERM:
				// 退出程序
				monitor.Stop()
				trayManager.Quit()
				fmt.Printf("\n\n%s %d %s\n", t.TotalCount, translationCount, t.TranslateCount)
				fmt.Println(t.Goodbye)
				os.Exit(0)
			}
		}
	}()
}


// 辅助函数

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// simulatePaste 模拟粘贴操作
func simulatePaste() {
	switch runtime.GOOS {
	case "linux":
		// 尝试使用xdotool
		if err := exec.Command("xdotool", "key", "ctrl+v").Run(); err != nil {
			// 如果xdotool不可用，尝试ydotool（Wayland）
			exec.Command("ydotool", "key", "29:1", "47:1", "47:0", "29:0").Run()
		}
	case "darwin":
		// macOS使用osascript
		exec.Command("osascript", "-e", `tell application "System Events" to keystroke "v" using command down`).Run()
	case "windows":
		// Windows暂不支持自动粘贴
		// 需要使用Windows API或AutoHotkey
	}
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
	// 直接从新系统获取prompt
	prompts := GetAllPrompts()
	
	for _, p := range prompts {
		if p.ID == id {
			// 调试：打印实际的内容长度
			fmt.Printf("\n[DEBUG] Found prompt %s, actual content length: %d\n", id, len(p.Content))
			if len(p.Content) < 100 {
				fmt.Printf("[DEBUG] Content: %s\n", p.Content)
			} else {
				fmt.Printf("[DEBUG] Content first 100 chars: %.100s...\n", p.Content)
			}
			return p.Content
		}
	}
	return "Translate the following to Chinese:"
}

var terminalVisible = false  // Start as false when running in background
var terminalPID = 0  // PID of the log viewer terminal

// hideTerminal 隐藏日志查看终端窗口
func hideTerminal() {
	if !terminalVisible || terminalPID == 0 {
		return
	}
	
	switch runtime.GOOS {
	case "linux":
		// Kill the log viewer terminal
		if terminalPID > 0 {
			exec.Command("kill", strconv.Itoa(terminalPID)).Run()
			terminalPID = 0
		}
		terminalVisible = false
		
	case "darwin":
		// macOS: Kill the log viewer terminal
		if terminalPID > 0 {
			exec.Command("kill", strconv.Itoa(terminalPID)).Run()
			terminalPID = 0
		}
		terminalVisible = false
		
	case "windows":
		// Windows: Kill the log viewer terminal
		if terminalPID > 0 {
			exec.Command("taskkill", "/PID", strconv.Itoa(terminalPID)).Run()
			terminalPID = 0
		}
		terminalVisible = false
	}
}

// showTerminal 显示日志查看终端窗口
func showTerminal() {
	if terminalVisible {
		return
	}
	
	switch runtime.GOOS {
	case "linux":
		// Open a new terminal to tail the log file
		var cmd *exec.Cmd
		
		// Try different terminal emulators
		if _, err := exec.LookPath("ptyxis"); err == nil {
			cmd = exec.Command("ptyxis", "--title", "xiaoniao 日志", "--", "tail", "-f", "/tmp/xiaoniao.log")
		} else if _, err := exec.LookPath("gnome-terminal"); err == nil {
			cmd = exec.Command("gnome-terminal", "--title=xiaoniao 日志", "--", "tail", "-f", "/tmp/xiaoniao.log")
		} else if _, err := exec.LookPath("konsole"); err == nil {
			cmd = exec.Command("konsole", "-caption", "xiaoniao 日志", "-e", "tail", "-f", "/tmp/xiaoniao.log")
		} else if _, err := exec.LookPath("xterm"); err == nil {
			cmd = exec.Command("xterm", "-title", "xiaoniao 日志", "-e", "tail", "-f", "/tmp/xiaoniao.log")
		}
		
		if cmd != nil {
			if err := cmd.Start(); err == nil {
				terminalPID = cmd.Process.Pid
				terminalVisible = true
			}
		}
		
	case "darwin":
		// macOS: Open Terminal with tail command
		cmd := exec.Command("osascript", "-e", `tell application "Terminal" to do script "tail -f /tmp/xiaoniao.log"`)
		if err := cmd.Start(); err == nil {
			terminalPID = cmd.Process.Pid
			terminalVisible = true
		}
		
	case "windows":
		// Windows: Open Command Prompt with tail equivalent
		cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "powershell Get-Content /tmp/xiaoniao.log -Wait")
		if err := cmd.Start(); err == nil {
			terminalPID = cmd.Process.Pid
			terminalVisible = true
		}
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


func openConfigInTerminal() {
	// 防止创建desktop文件的终极方案
	
	// 首先检查并删除任何自动生成的配置desktop文件
	configDesktopPath := filepath.Join(os.Getenv("HOME"), ".local/share/applications/xiaoniao-config.desktop")
	os.RemoveAll(configDesktopPath) // 删除文件或目录
	
	// 输出调试信息
	// fmt.Println("Opening configuration...")
	
	// 尝试多种方式打开终端
	// 1. 使用 ptyxis (Fedora 的新默认终端)
	cmd := exec.Command("ptyxis", "--", "xiaoniao", "config")
	
	if err := cmd.Start(); err != nil {
		
		// 2. 尝试 gnome-terminal (通用)
		cmd = exec.Command("gnome-terminal", "--", "xiaoniao", "config")
		if err := cmd.Start(); err != nil {
			fmt.Printf("gnome-terminal 失败: %v\n", err)
			
			// 3. 尝试 kgx (GNOME Console)
			cmd = exec.Command("kgx", "--", "xiaoniao", "config")
			if err := cmd.Start(); err != nil {
				fmt.Printf("kgx 失败: %v\n", err)
				
				// 4. 尝试 xterm 作为最后备用
				cmd = exec.Command("xterm", "-hold", "-e", "xiaoniao", "config")
				if err := cmd.Start(); err != nil {
					fmt.Printf("xterm 也失败: %v\n", err)
					
					// 5. 尝试 konsole (KDE)
					cmd = exec.Command("konsole", "-e", "xiaoniao", "config")
					if err := cmd.Start(); err != nil {
						fmt.Printf("所有终端都无法打开\n")
						// 最后的备用：通知用户手动运行
						// 不显示通知，直接输出到终端
						fmt.Println("请手动运行: xiaoniao config")
					}
				}
			}
		}
	}
	
	// 等待一下让终端有时间启动
	time.Sleep(1 * time.Second)
	
	// 延迟再次清理（防止延迟创建）
	go func() {
		time.Sleep(500 * time.Millisecond)
		os.RemoveAll(configDesktopPath)
		// 创建一个同名目录阻止文件创建
		os.MkdirAll(configDesktopPath, 0755)
	}()
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