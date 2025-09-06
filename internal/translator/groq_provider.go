package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GroqProvider Groq高速推理provider
type GroqProvider struct {
	*BaseProvider
	BaseURL string
}

// NewGroqProvider 创建Groq provider
func NewGroqProvider(apiKey, model string) *GroqProvider {
	if model == "" {
		model = "llama-3.1-8b-instant"
	}
	return &GroqProvider{
		BaseProvider: NewBaseProvider(apiKey, model),
		BaseURL:      "https://api.groq.com/openai/v1",
	}
}

func (p *GroqProvider) Name() string {
	return "Groq"
}

// ListModels 获取Groq的模型列表
func (p *GroqProvider) ListModels() ([]string, error) {
	url := p.BaseURL + "/models"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	req.Header.Set("Content-Type", "application/json")

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

	var result struct {
		Data []struct {
			ID     string `json:"id"`
			Object string `json:"object"`
			Created int   `json:"created"`
			OwnedBy string `json:"owned_by"`
		} `json:"data"`
		Object string `json:"object"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse models: %v", err)
	}

	models := []string{}
	for _, model := range result.Data {
		if model.Object == "model" {
			models = append(models, model.ID)
		}
	}

	// 如果API调用失败，返回已知的Groq模型
	if len(models) == 0 {
		models = []string{
			"llama-3.3-70b-versatile",
			"llama-3.1-405b-reasoning",
			"llama-3.1-70b-versatile", 
			"llama-3.1-8b-instant",
			"llama3-groq-70b-8192-tool-use-preview",
			"llama3-groq-8b-8192-tool-use-preview",
			"llama-3.2-1b-preview",
			"llama-3.2-3b-preview",
			"llama-3.2-11b-vision-preview",
			"llama-3.2-90b-vision-preview",
			"mixtral-8x7b-32768",
			"gemma-7b-it",
			"gemma2-9b-it",
		}
	}

	return models, nil
}

// Translate 使用Groq进行翻译
func (p *GroqProvider) Translate(text, prompt string) (string, error) {
	url := p.BaseURL + "/chat/completions"
	
	// 使用统一的底层系统prompt
	systemPrompt := BuildSystemPrompt(prompt)
	
	payload := map[string]interface{}{
		"model": p.Model,
		"messages": []map[string]interface{}{
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
		"stream":     false,
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
			Type    string `json:"type"`
			Code    string `json:"code"`
		} `json:"error"`
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	if result.Error.Message != "" {
		return "", fmt.Errorf("%s: %s", result.Error.Type, result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from Groq API")
	}

	translation := strings.TrimSpace(result.Choices[0].Message.Content)
	return translation, nil
}

// TestConnection 测试Groq连接
func (p *GroqProvider) TestConnection() error {
	// 尝试获取模型列表来测试连接
	models, err := p.ListModels()
	if err != nil {
		// 如果获取模型列表失败，尝试一个简单的翻译请求
		_, err = p.Translate("Hello", "翻译成中文")
		if err != nil {
			return fmt.Errorf("Groq connection test failed: %v", err)
		}
	} else if len(models) == 0 {
		return fmt.Errorf("no Groq models available")
	}
	return nil
}