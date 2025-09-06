package translator

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TranslationResult contains the result of a translation
type TranslationResult struct {
	Original    string    `json:"original"`
	Translation string    `json:"translation"`
	Provider    string    `json:"provider"`
	Model       string    `json:"model"`
	Prompt      string    `json:"prompt"`
	Time        int64     `json:"time"`     // milliseconds
	Tokens      int       `json:"tokens"`    // estimated tokens
	Success     bool      `json:"success"`
	Error       string    `json:"error,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

// TranslationHistory manages translation history
type TranslationHistory struct {
	mu       sync.RWMutex
	items    []TranslationResult
	maxItems int
}

// Translator manages translations using different providers
type Translator struct {
	provider   Provider
	config     *Config
	history    *TranslationHistory
	mu         sync.RWMutex
}

// Config holds translator configuration
type Config struct {
	APIKey        string `json:"apiKey"`
	Provider      string `json:"provider"`
	Model         string `json:"model"`
	FallbackModel string `json:"fallbackModel,omitempty"` // 副模型
	MaxRetries    int    `json:"maxRetries"`
	Timeout       int    `json:"timeout"` // seconds
	HistorySize   int    `json:"historySize"`
}


// NewTranslator creates a new translator instance
func NewTranslator(config *Config) (*Translator, error) {
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}
	if config.Timeout == 0 {
		config.Timeout = 30
	}
	if config.HistorySize == 0 {
		config.HistorySize = 100
	}

	translator := &Translator{
		config: config,
		history: &TranslationHistory{
			items:    make([]TranslationResult, 0, config.HistorySize),
			maxItems: config.HistorySize,
		},
	}

	// Initialize provider based on configuration
	if config.APIKey != "" && config.Provider != "" {
		if err := translator.SetProvider(config.Provider, config.APIKey, config.Model); err != nil {
			return nil, err
		}
	}

	return translator, nil
}

// SetProvider sets the active provider
func (t *Translator) SetProvider(providerName, apiKey, model string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	var provider Provider

	switch providerName {
	case "OpenAI":
		provider = NewOpenAIProvider(apiKey, model)
	case "Anthropic":
		provider = NewAnthropicProvider(apiKey, model)
	case "DeepSeek":
		provider = NewDeepSeekProvider(apiKey, model)
	case "Moonshot":
		provider = NewMoonshotProvider(apiKey, model)
	case "OpenRouter":
		p := NewOpenRouterProvider(apiKey, model)
		if t.config.FallbackModel != "" {
			p.SetFallbackModel(t.config.FallbackModel)
		}
		provider = p
	case "Groq":
		provider = NewGroqProvider(apiKey, model)
	case "Together", "TogetherAI":
		provider = NewTogetherProvider(apiKey, model)
	default:
		// 尝试作为OpenAI兼容的provider
		provider = NewOpenAICompatibleProvider(providerName, apiKey, "", model)
	}

	// Test connection - 跳过初始化时的测试，避免启动失败
	// 实际翻译时如果有问题会报错
	// if err := provider.TestConnection(); err != nil {
	//	return fmt.Errorf("failed to connect to %s: %w", providerName, err)
	// }

	t.provider = provider
	t.config.Provider = providerName
	t.config.APIKey = apiKey
	t.config.Model = model

	return nil
}

// Translate performs a translation using the current provider
func (t *Translator) Translate(text, prompt string) (*TranslationResult, error) {
	if t.provider == nil {
		return nil, fmt.Errorf("no provider configured")
	}


	startTime := time.Now()

	// Perform translation with retries
	var translation string
	var lastErr error
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t.config.Timeout)*time.Second)
	defer cancel()

	retryCount := 0
	for retryCount < t.config.MaxRetries {
		// Check context
		select {
		case <-ctx.Done():
			lastErr = fmt.Errorf("translation timeout")
			break
		default:
		}

		translation, lastErr = t.provider.Translate(text, prompt)
		if lastErr == nil {
			// Check if the model returned "1" indicating it cannot translate
			if translation == "1" {
				lastErr = fmt.Errorf("model unable to perform translation")
			} else {
				break
			}
		}

		retryCount++
		if retryCount < t.config.MaxRetries {
			// Exponential backoff
			time.Sleep(time.Duration(retryCount) * time.Second)
		}
	}

	elapsed := time.Since(startTime).Milliseconds()

	result := &TranslationResult{
		Original:    text,
		Translation: translation,
		Provider:    t.config.Provider,
		Model:       t.config.Model,
		Prompt:      prompt,
		Time:        elapsed,
		Tokens:      estimateTokens(text + translation),
		Success:     lastErr == nil,
		Timestamp:   time.Now(),
	}

	if lastErr != nil {
		result.Error = lastErr.Error()
	}

	// Add to history
	t.history.Add(*result)

	return result, lastErr
}

// GetHistory returns translation history
func (t *Translator) GetHistory(limit int) []TranslationResult {
	return t.history.GetRecent(limit)
}

// ClearHistory clears translation history
func (t *Translator) ClearHistory() {
	t.history.Clear()
}

// GetProviderInfo returns information about the current provider
func (t *Translator) GetProviderInfo() (map[string]interface{}, error) {
	if t.provider == nil {
		return nil, fmt.Errorf("no provider configured")
	}

	models, err := t.provider.ListModels()
	if err != nil {
		models = []string{t.config.Model}
	}

	return map[string]interface{}{
		"name":   t.provider.Name(),
		"model":  t.config.Model,
		"models": models,
	}, nil
}

// TranslationHistory methods

func (h *TranslationHistory) Add(result TranslationResult) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.items = append(h.items, result)
	
	// Keep only maxItems
	if len(h.items) > h.maxItems {
		h.items = h.items[len(h.items)-h.maxItems:]
	}
}

func (h *TranslationHistory) GetRecent(limit int) []TranslationResult {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if limit <= 0 || limit > len(h.items) {
		limit = len(h.items)
	}

	// Return most recent items first
	result := make([]TranslationResult, limit)
	for i := 0; i < limit; i++ {
		result[i] = h.items[len(h.items)-1-i]
	}

	return result
}

func (h *TranslationHistory) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.items = h.items[:0]
}

// estimateTokens provides a rough estimate of token count
func estimateTokens(text string) int {
	// Rough estimate: ~4 characters per token for English, ~2 for Chinese
	return len(text) / 3
}