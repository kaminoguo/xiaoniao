package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/kaminoguo/xiaoniao-mac/internal/clipboard"
	"github.com/kaminoguo/xiaoniao-mac/internal/config"
	"github.com/kaminoguo/xiaoniao-mac/internal/hotkey"
	"github.com/kaminoguo/xiaoniao-mac/internal/i18n"
	"github.com/kaminoguo/xiaoniao-mac/internal/logbuffer"
	"github.com/kaminoguo/xiaoniao-mac/internal/sound"
	"github.com/kaminoguo/xiaoniao-mac/internal/tray"
	"github.com/kaminoguo/xiaoniao-mac/internal/translator"
	"golang.design/x/hotkey/mainthread"
)

const version = "1.1.0-mac"

type Config struct {
	APIKey        string `json:"api_key"`
	Provider      string `json:"provider"`
	Model         string `json:"model"`
	PromptID      string `json:"prompt_id"`
	Language      string `json:"language,omitempty"`
	Theme         string `json:"theme,omitempty"`
	HotkeyToggle  string `json:"hotkey_toggle,omitempty"`
	HotkeySwitch  string `json:"hotkey_switch,omitempty"`
}

var (
	configPath       string
	cfg              Config
	clipboardMonitor *clipboard.Monitor
	hotkeyManager    *hotkey.Manager
	trayManager      *tray.Manager
	trans            *translator.Translator
	isMonitoring     = false
	logBuffer        = logbuffer.GetInstance()
)

func init() {
	// Setup config directory
	configDir := config.GetConfigDir()
	configPath = filepath.Join(configDir, "config.json")

	// Setup logging
	logPath := config.GetLogPath()
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(logFile)
	}

	// Load config
	loadConfig()

	// Initialize i18n
	if cfg.Language != "" {
		i18n.SetLanguage(i18n.Language(cfg.Language))
	}
}

func main() {
	// macOS requires running on main thread
	mainthread.Init(run)
}

func run() {
	log.Printf("小鸟翻译 Mac版 v%s 启动", version)

	// Initialize translator
	if cfg.APIKey != "" && cfg.Provider != "" && cfg.Model != "" {
		config := &translator.Config{
			APIKey:   cfg.APIKey,
			Provider: cfg.Provider,
			Model:    cfg.Model,
		}
		var err error
		trans, err = translator.NewTranslator(config)
		if err != nil {
			log.Printf("Failed to create translator: %v", err)
		}
	}

	// Initialize tray
	var err error
	trayManager, err = tray.NewManager()
	if err != nil {
		log.Fatal("Failed to create tray manager:", err)
	}

	// Set tray callbacks
	setupTrayCallbacks()

	// Set business logic
	trayManager.SetBusinessLogic(func() {
		// Initialize clipboard monitor
		clipboardMonitor = clipboard.NewMonitor()
		clipboardMonitor.SetOnChange(onClipboardChange)

		// Initialize hotkey manager
		hotkeyManager = hotkey.NewManager()
		registerHotkeys()

		// Wait for tray to be ready
		for !trayManager.IsReady() {
			time.Sleep(100 * time.Millisecond)
		}

		// Start monitoring if configured
		if cfg.APIKey != "" && cfg.Provider != "" && cfg.Model != "" {
			startMonitoring()
		}
	})

	// Run tray (blocks until quit)
	trayManager.Run()
}

func setupTrayCallbacks() {
	trayManager.SetOnQuit(func() {
		stopMonitoring()
		os.Exit(0)
	})

	trayManager.SetOnSettings(func() {
		// Open settings (can be terminal UI or native dialog)
		log.Println("Settings clicked - TODO: implement settings UI")
	})

	trayManager.SetOnToggleMonitor(func(enabled bool) {
		if enabled {
			startMonitoring()
		} else {
			stopMonitoring()
		}
	})

	trayManager.SetOnRefresh(func() {
		loadConfig()
		if trans != nil {
			config := &translator.Config{
			APIKey:   cfg.APIKey,
			Provider: cfg.Provider,
			Model:    cfg.Model,
		}
		var err error
		trans, err = translator.NewTranslator(config)
		if err != nil {
			log.Printf("Failed to create translator: %v", err)
		}
		}
		log.Println("配置已刷新")
	})

	trayManager.SetOnSelectPrompt(func(promptID string) {
		cfg.PromptID = promptID
		saveConfig()
		if trans != nil {
			// trans.SetPromptID(promptID) // TODO: Implement prompt ID support
		}
		trayManager.SetCurrentPrompt(promptID)
	})
}

func registerHotkeys() {
	// Register toggle monitoring hotkey
	if cfg.HotkeyToggle != "" {
		err := hotkeyManager.RegisterFromString("toggle", cfg.HotkeyToggle, func() {
			isMonitoring = !isMonitoring
			trayManager.UpdateMonitorStatus(isMonitoring)
			if isMonitoring {
				startMonitoring()
			} else {
				stopMonitoring()
			}
		})
		if err != nil {
			log.Printf("Failed to register toggle hotkey: %v", err)
		}
	}

	// Register switch prompt hotkey
	if cfg.HotkeySwitch != "" {
		err := hotkeyManager.RegisterFromString("switch", cfg.HotkeySwitch, func() {
			// TODO: Implement prompt switching
			log.Println("Switch prompt hotkey pressed")
		})
		if err != nil {
			log.Printf("Failed to register switch hotkey: %v", err)
		}
	}
}

func onClipboardChange(text string) {
	if !isMonitoring || trans == nil {
		return
	}

	log.Printf("检测到剪贴板变化: %d 字符", len(text))
	trayManager.SetStatus(tray.StatusTranslating)

	// Translate
	result, err := trans.Translate(text, "")
	if err != nil {
		log.Printf("翻译失败: %v", err)
		trayManager.SetStatus(tray.StatusError)
		sound.PlayError()
		return
	}

	// Set translated text back to clipboard
	clipboardMonitor.SetLastTranslation(result.Translation)
	if err := clipboard.SetClipboard(result.Translation); err != nil {
		log.Printf("设置剪贴板失败: %v", err)
		trayManager.SetStatus(tray.StatusError)
		return
	}

	log.Printf("翻译成功: %d 字符", len(result.Translation))
	trayManager.SetStatus(tray.StatusIdle)
	trayManager.IncrementTranslationCount()
	sound.PlaySuccess()

	// Auto paste (simulate Cmd+V)
	// TODO: Implement auto paste for macOS
}

func startMonitoring() {
	if clipboardMonitor != nil {
		clipboardMonitor.Start()
		isMonitoring = true
		trayManager.UpdateMonitorStatus(true)
		log.Println("开始监控剪贴板")
	}
}

func stopMonitoring() {
	if clipboardMonitor != nil {
		clipboardMonitor.Stop()
		isMonitoring = false
		trayManager.UpdateMonitorStatus(false)
		log.Println("停止监控剪贴板")
	}
}

func loadConfig() {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return
	}
	json.Unmarshal(data, &cfg)
}

func saveConfig() {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(configPath, data, 0644)
}

// Handle interrupt signal
func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		stopMonitoring()
		os.Exit(0)
	}()
}