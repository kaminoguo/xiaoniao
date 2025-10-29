package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Provider represents an AI translation provider
type Provider interface {
	Name() string
	Translate(text, prompt string) (string, error)
	ListModels() ([]string, error)
	TestConnection() error
}

// BaseProvider contains common fields for all providers
type BaseProvider struct {
	APIKey string
	Model  string
	Client *http.Client
}

// NewBaseProvider creates a new base provider
func NewBaseProvider(apiKey, model string) *BaseProvider {
	return &BaseProvider{
		APIKey: apiKey,
		Model:  model,
		Client: GetSharedHTTPClient(),
	}
}

// DetectProvider detects the provider type from API key
func DetectProvider(apiKey string) (string, []string, error) {
	// Try different providers
	providers := []struct {
		name     string
		detector func(string) ([]string, error)
	}{
		{"OpenAI", detectOpenAI},
		{"Anthropic", detectAnthropic},
		{"DeepSeek", detectDeepSeek},
		{"Moonshot", detectMoonshot},
	}

	for _, p := range providers {
		if models, err := p.detector(apiKey); err == nil {
			return p.name, models, nil
		}
	}

	return "", nil, fmt.Errorf("unable to detect provider for the given API key")
}

// OpenAI Provider
type OpenAIProvider struct {
	*BaseProvider
	BaseURL string
}

func NewOpenAIProvider(apiKey, model string) *OpenAIProvider {
	if model == "" {
		model = "gpt-4o-mini"
	}
	return &OpenAIProvider{
		BaseProvider: NewBaseProvider(apiKey, model),
		BaseURL:      "https://api.openai.com/v1",
	}
}

func (p *OpenAIProvider) Name() string {
	return "OpenAI"
}

func (p *OpenAIProvider) Translate(text, prompt string) (string, error) {
	url := p.BaseURL + "/chat/completions"
	
	// 使用统一的底层系统prompt
	systemPrompt := BuildSystemPrompt(prompt)
	
	payload := map[string]interface{}{
		"model": p.Model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt,
			},
			{
				"role":    "user",
				"content": text,
			},
		},
		"temperature": 0.3,
		"max_tokens":  2000,
		"n": 1,           // 只生成1个响应
		"top_p": 0.9,     // 限制采样范围，提高一致性
	}

	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.APIKey)

	resp, err := p.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", string(body))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no translation returned")
}

func (p *OpenAIProvider) ListModels() ([]string, error) {
	url := p.BaseURL + "/models"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+p.APIKey)

	resp, err := p.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list models")
	}

	var result struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	models := []string{}
	for _, model := range result.Data {
		// 返回所有模型，不过滤
		models = append(models, model.ID)
	}

	return models, nil
}

func (p *OpenAIProvider) TestConnection() error {
	_, err := p.ListModels()
	return err
}

func detectOpenAI(apiKey string) ([]string, error) {
	provider := NewOpenAIProvider(apiKey, "")
	return provider.ListModels()
}

// Anthropic Provider
type AnthropicProvider struct {
	*BaseProvider
	BaseURL string
}

func NewAnthropicProvider(apiKey, model string) *AnthropicProvider {
	if model == "" {
		model = "claude-3-haiku-20240307"
	}
	return &AnthropicProvider{
		BaseProvider: NewBaseProvider(apiKey, model),
		BaseURL:      "https://api.anthropic.com/v1",
	}
}

func (p *AnthropicProvider) Name() string {
	return "Anthropic"
}

func (p *AnthropicProvider) Translate(text, prompt string) (string, error) {
	url := p.BaseURL + "/messages"
	
	// Anthropic需要特殊处理 - 将text包含在prompt中
	systemPrompt := BaseSystemPromptForAnthropic(prompt, text)
	
	payload := map[string]interface{}{
		"model": p.Model,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": systemPrompt,
			},
		},
		"max_tokens": 2000,
		"system":     "You are a translation tool. Only output the translation, no explanations.",
	}

	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := p.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s", string(body))
	}

	var result struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.Content) > 0 {
		return result.Content[0].Text, nil
	}

	return "", fmt.Errorf("no translation returned")
}

func (p *AnthropicProvider) ListModels() ([]string, error) {
	// Anthropic doesn't have a list models endpoint, return known models
	return []string{
		"claude-3-opus-20240229",
		"claude-3-sonnet-20240229",
		"claude-3-haiku-20240307",
		"claude-2.1",
		"claude-2.0",
	}, nil
}

func (p *AnthropicProvider) TestConnection() error {
	// Try a simple request to test the API key
	_, err := p.Translate("test", "Translate to English: {{text}}")
	return err
}

func detectAnthropic(apiKey string) ([]string, error) {
	provider := NewAnthropicProvider(apiKey, "")
	err := provider.TestConnection()
	if err != nil {
		return nil, err
	}
	return provider.ListModels()
}

// DeepSeek Provider (similar to OpenAI)
type DeepSeekProvider struct {
	*OpenAIProvider
}

func NewDeepSeekProvider(apiKey, model string) *DeepSeekProvider {
	if model == "" {
		model = "deepseek-chat"
	}
	provider := NewOpenAIProvider(apiKey, model)
	provider.BaseURL = "https://api.deepseek.com/v1"
	return &DeepSeekProvider{
		OpenAIProvider: provider,
	}
}

func (p *DeepSeekProvider) Name() string {
	return "DeepSeek"
}

func detectDeepSeek(apiKey string) ([]string, error) {
	provider := NewDeepSeekProvider(apiKey, "")
	return provider.ListModels()
}

// Moonshot Provider (similar to OpenAI)
type MoonshotProvider struct {
	*OpenAIProvider
}

func NewMoonshotProvider(apiKey, model string) *MoonshotProvider {
	if model == "" {
		model = "moonshot-v1-8k"
	}
	provider := NewOpenAIProvider(apiKey, model)
	provider.BaseURL = "https://api.moonshot.cn/v1"
	return &MoonshotProvider{
		OpenAIProvider: provider,
	}
}

func (p *MoonshotProvider) Name() string {
	return "Moonshot"
}

func detectMoonshot(apiKey string) ([]string, error) {
	provider := NewMoonshotProvider(apiKey, "")
	return provider.ListModels()
}