package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// OpenAICompatibleProvider 通用的OpenAI兼容API Provider
type OpenAICompatibleProvider struct {
	*BaseProvider
	BaseURL      string
	ProviderName string
	Headers      map[string]string // 额外的请求头
}

// NewOpenAICompatibleProvider 创建OpenAI兼容的Provider
func NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model string) *OpenAICompatibleProvider {
	// 如果没有指定baseURL，从注册表获取
	if baseURL == "" {
		if config, exists := GetProviderConfig(providerName); exists {
			baseURL = config.BaseURL
		}
	}
	
	// 如果没有指定model，使用默认值
	if model == "" {
		model = getDefaultModel(providerName)
	}
	
	headers := make(map[string]string)
	
	// 特殊处理某些provider的请求头
	switch providerName {
	case "OpenRouter":
		headers["HTTP-Referer"] = "https://xiaoniao-translator.com"
		headers["X-Title"] = "Xiaoniao Translator"
	case "Perplexity":
		headers["Accept"] = "application/json"
	}
	
	return &OpenAICompatibleProvider{
		BaseProvider: NewBaseProvider(apiKey, model),
		BaseURL:      baseURL,
		ProviderName: providerName,
		Headers:      headers,
	}
}

// Name 返回Provider名称
func (p *OpenAICompatibleProvider) Name() string {
	return p.ProviderName
}

// ListModels 获取模型列表
func (p *OpenAICompatibleProvider) ListModels() ([]string, error) {
	url := p.BaseURL + "/models"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置认证头
	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	
	// 设置额外的请求头
	for key, value := range p.Headers {
		req.Header.Set(key, value)
	}

	client := GetSharedHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list models: %s", string(body))
	}

	// 尝试解析响应
	var result struct {
		Data []struct {
			ID   string `json:"id"`
			Name string `json:"name"` // 某些API使用name而不是id
		} `json:"data"`
		Models []string `json:"models"` // 某些API直接返回models数组
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse models: %v", err)
	}

	models := []string{}
	
	// 处理不同的响应格式
	if len(result.Data) > 0 {
		for _, model := range result.Data {
			if model.ID != "" {
				models = append(models, model.ID)
			} else if model.Name != "" {
				models = append(models, model.Name)
			}
		}
	} else if len(result.Models) > 0 {
		models = result.Models
	}

	// 如果没有获取到模型，返回一些默认模型
	if len(models) == 0 && p.ProviderName != "" {
		models = getDefaultModels(p.ProviderName)
	}

	return models, nil
}

// Translate 翻译文本
func (p *OpenAICompatibleProvider) Translate(text, prompt string) (string, error) {
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
	}

	// 某些provider需要特殊处理
	if p.ProviderName == "OpenRouter" {
		// OpenRouter支持provider参数来指定使用哪个底层provider
		payload["provider"] = map[string]interface{}{
			"order": []string{"Together", "DeepInfra", "Hyperbolic"},
		}
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	req.Header.Set("Content-Type", "application/json")
	
	// 设置额外的请求头
	for key, value := range p.Headers {
		req.Header.Set(key, value)
	}

	client := GetSharedHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed (%d): %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	if result.Error.Message != "" {
		return "", fmt.Errorf(result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	translation := strings.TrimSpace(result.Choices[0].Message.Content)
	return translation, nil
}

// TestConnection 测试连接
func (p *OpenAICompatibleProvider) TestConnection() error {
	// 尝试获取模型列表来测试连接
	models, err := p.ListModels()
	if err != nil {
		// 如果获取模型列表失败，尝试一个简单的翻译请求
		_, err = p.Translate("Hello", "翻译成中文")
		if err != nil {
			return fmt.Errorf("connection test failed: %v", err)
		}
	} else if len(models) == 0 {
		// 即使没有返回模型，连接可能是成功的
		return nil
	}
	return nil
}

// getDefaultModel 获取provider的默认模型
func getDefaultModel(provider string) string {
	defaults := map[string]string{
		"OpenRouter":  "openai/gpt-4o-mini",
		"Groq":        "llama-3.1-8b-instant", 
		"Together":    "meta-llama/Llama-3-8b-chat-hf",
		"Perplexity":  "llama-3.1-sonar-small-128k-online",
		"DeepSeek":    "deepseek-chat",
		"Mistral":     "mistral-small-latest",
		"Cohere":      "command-r",
		"HuggingFace": "meta-llama/Llama-2-7b-chat-hf",
		"Replicate":   "meta/llama-2-7b-chat",
	}
	
	if model, exists := defaults[provider]; exists {
		return model
	}
	return ""
}

// getDefaultModels 获取provider的默认模型列表（当API调用失败时使用）
func getDefaultModels(provider string) []string {
	// 这里只返回一些基础模型，实际模型应该通过API获取
	defaults := map[string][]string{
		"OpenRouter": {
			"openai/gpt-4o", "openai/gpt-4o-mini", "anthropic/claude-3.5-sonnet",
			"google/gemini-pro-1.5", "meta-llama/llama-3.1-405b", "deepseek/deepseek-v3",
		},
		"Groq": {
			"llama-3.1-405b-reasoning", "llama-3.1-70b-versatile", "llama-3.1-8b-instant",
			"mixtral-8x7b-32768", "gemma-7b-it",
		},
		"Together": {
			"meta-llama/Meta-Llama-3.1-405B-Instruct-Turbo",
			"meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo",
			"meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo",
			"mistralai/Mixtral-8x7B-Instruct-v0.1",
		},
	}
	
	if models, exists := defaults[provider]; exists {
		return models
	}
	return []string{}
}