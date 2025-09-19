package config

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/kaminoguo/xiaoniao/internal/translator"
)

// Config holds the application configuration
type Config struct {
	mu sync.RWMutex

	// API Configuration
	APIKey    string `json:"api_key"`
	APIUrl    string `json:"api_url,omitempty"`
	Provider  string `json:"provider"`
	Model     string `json:"model"`

	// Translation Settings
	CurrentPrompt string   `json:"current_prompt"`
	UserPrompts   []string `json:"user_prompts,omitempty"`

	// Monitor Settings
	MonitorMode string `json:"monitor_mode"` // "off", "text_menu", "clipboard"

	// UI Settings
	Language string `json:"language"` // "en_US", "zh_CN", "ja_JP"
}

// NewConfig creates a new configuration with defaults
func NewConfig() *Config {
	return &Config{
		Provider:      "auto",
		Model:         "gpt-4o-mini",
		CurrentPrompt: translator.SimplePrompt,
		MonitorMode:   "text_menu",
		Language:      "en_US",
	}
}

// Load loads configuration from file
func Load() (*Config, error) {
	config := NewConfig()

	configPath := GetConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return default config if file doesn't exist
			return config, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

// Save saves configuration to file
func (c *Config) Save() error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Ensure config directory exists
	configDir := GetConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	// Marshal config to JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	configPath := GetConfigPath()
	return os.WriteFile(configPath, data, 0644)
}

// GetAPIKey returns the API key
func (c *Config) GetAPIKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.APIKey
}

// SetAPIKey sets the API key
func (c *Config) SetAPIKey(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.APIKey = key
}

// GetModel returns the current model
func (c *Config) GetModel() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Model
}

// SetModel sets the model
func (c *Config) SetModel(model string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Model = model
}

// GetCurrentPrompt returns the current prompt
func (c *Config) GetCurrentPrompt() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.CurrentPrompt
}

// SetCurrentPrompt sets the current prompt
func (c *Config) SetCurrentPrompt(prompt string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CurrentPrompt = prompt
}

// GetMonitorMode returns the monitor mode
func (c *Config) GetMonitorMode() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.MonitorMode
}

// SetMonitorMode sets the monitor mode
func (c *Config) SetMonitorMode(mode string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.MonitorMode = mode
}

// GetLanguage returns the UI language
func (c *Config) GetLanguage() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Language
}

// SetLanguage sets the UI language
func (c *Config) SetLanguage(lang string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Language = lang
}