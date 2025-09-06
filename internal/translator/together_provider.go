package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// TogetherProvider Together AI provider
type TogetherProvider struct {
	*BaseProvider
	BaseURL string
}

// NewTogetherProvider 创建Together provider
func NewTogetherProvider(apiKey, model string) *TogetherProvider {
	if model == "" {
		model = "meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
	}
	return &TogetherProvider{
		BaseProvider: NewBaseProvider(apiKey, model),
		BaseURL:      "https://api.together.xyz/v1",
	}
}

func (p *TogetherProvider) Name() string {
	return "Together"
}

// ListModels 获取Together的模型列表
func (p *TogetherProvider) ListModels() ([]string, error) {
	url := p.BaseURL + "/models"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+p.APIKey)

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

	// Together返回的格式可能不同
	var result []struct {
		ID          string `json:"id"`
		Object      string `json:"object"`
		Created     int    `json:"created"`
		Type        string `json:"type"`
		DisplayName string `json:"display_name"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		// 尝试OpenAI格式
		var openAIFormat struct {
			Data []struct {
				ID string `json:"id"`
			} `json:"data"`
		}
		if err2 := json.Unmarshal(body, &openAIFormat); err2 == nil {
			models := []string{}
			for _, model := range openAIFormat.Data {
				models = append(models, model.ID)
			}
			return models, nil
		}
		return nil, fmt.Errorf("failed to parse models: %v", err)
	}

	models := []string{}
	for _, model := range result {
		// 只返回聊天模型
		if model.Type == "chat" || model.Type == "" {
			models = append(models, model.ID)
		}
	}

	// 如果API调用失败，返回已知的Together模型
	if len(models) == 0 {
		models = []string{
			"meta-llama/Meta-Llama-3.1-405B-Instruct-Turbo",
			"meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo",
			"meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo",
			"meta-llama/Llama-3.2-11B-Vision-Instruct-Turbo",
			"meta-llama/Llama-3.2-3B-Instruct-Turbo",
			"meta-llama/Llama-3.2-90B-Vision-Instruct-Turbo",
			"mistralai/Mixtral-8x7B-Instruct-v0.1",
			"mistralai/Mistral-7B-Instruct-v0.3",
			"Qwen/Qwen2.5-72B-Instruct-Turbo",
			"Qwen/Qwen2.5-7B-Instruct-Turbo",
			"deepseek-ai/deepseek-llm-67b-chat",
			"google/gemma-2b-it",
			"google/gemma-7b-it",
		}
	}

	return models, nil
}

// Translate 使用Together进行翻译
func (p *TogetherProvider) Translate(text, prompt string) (string, error) {
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
		} `json:"error"`
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	if result.Error.Message != "" {
		return "", fmt.Errorf(result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from Together API")
	}

	translation := strings.TrimSpace(result.Choices[0].Message.Content)
	return translation, nil
}

// TestConnection 测试Together连接
func (p *TogetherProvider) TestConnection() error {
	// 尝试获取模型列表来测试连接
	models, err := p.ListModels()
	if err != nil {
		// 如果获取模型列表失败，尝试一个简单的翻译请求
		_, err = p.Translate("Hello", "翻译成中文")
		if err != nil {
			return fmt.Errorf("Together connection test failed: %v", err)
		}
	} else if len(models) == 0 {
		return fmt.Errorf("no Together models available")
	}
	return nil
}