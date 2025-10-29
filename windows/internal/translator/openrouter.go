package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// OpenRouterProvider OpenRouter API提供商
type OpenRouterProvider struct {
	APIKey        string
	Model         string
	BaseURL       string
}

// NewOpenRouterProvider 创建OpenRouter提供商
func NewOpenRouterProvider(apiKey, model string) *OpenRouterProvider {
	return &OpenRouterProvider{
		APIKey:  apiKey,
		Model:   model,
		BaseURL: "https://openrouter.ai/api/v1",
	}
}


// GetAvailableModels 获取可用的模型列表
func (p *OpenRouterProvider) GetAvailableModels() ([]string, error) {
	url := p.BaseURL + "/models"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	req.Header.Set("HTTP-Referer", "https://xiaoniao-translator.com")
	req.Header.Set("X-Title", "Xiaoniao Translator")
	
	// 使用共享的HTTP客户端
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
		return nil, fmt.Errorf("failed to get models: %s", string(body))
	}
	
	// 解析响应
	var result struct {
		Data []struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			Pricing       struct {
				Prompt     string `json:"prompt"`
				Completion string `json:"completion"`
				Image      string `json:"image"`
				Request    string `json:"request"`
			} `json:"pricing"`
			ContextLength int    `json:"context_length"`
			Architecture  struct {
				Modality      string `json:"modality"`
				Tokenizer     string `json:"tokenizer"`
				InstructType  string `json:"instruct_type"`
			} `json:"architecture"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse models: %v", err)
	}

	models := []string{}
	for _, model := range result.Data {
		// 只返回文本相关模型，排除纯图像生成模型
		// OpenRouter的modality格式：text->text, text+image->text, image->image等
		modality := model.Architecture.Modality
		if modality == "" || 
		   strings.Contains(modality, "text") && !strings.HasPrefix(modality, "image->") {
			models = append(models, model.ID)
		}
	}

	return models, nil
}

// Translate 使用OpenRouter进行翻译
func (p *OpenRouterProvider) Translate(text, prompt string) (string, error) {
	// 直接使用主模型翻译
	return p.translateWithModel(text, prompt, p.Model)
}

// translateWithModel 使用指定模型翻译
func (p *OpenRouterProvider) translateWithModel(text, prompt, model string) (string, error) {
	url := p.BaseURL + "/chat/completions"
	
	// 使用统一的底层系统prompt
	systemPrompt := BuildSystemPrompt(prompt)
	
	payload := map[string]interface{}{
		"model": model,
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
		"frequency_penalty": 0,  // 不惩罚重复
		"presence_penalty": 0,   // 不强制新颖性
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
	req.Header.Set("HTTP-Referer", "https://xiaoniao-translator.com")
	req.Header.Set("X-Title", "Xiaoniao Translator")

	// 使用共享的HTTP客户端
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
		return "", fmt.Errorf("API request failed: %s", string(respBody))
	}

	var result struct {
		Model   string `json:"model"`
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
		return "", err
	}

	if result.Error.Message != "" {
		return "", fmt.Errorf(result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	translation := strings.TrimSpace(result.Choices[0].Message.Content)
	
	// 调试：如果翻译内容包含换行符或多个版本，记录日志
	if strings.Contains(translation, "\n") {
		fmt.Printf("[DEBUG] AI返回了多行内容 (长度: %d):\n%s\n", len(translation), translation)
		// 只取第一行作为翻译结果
		lines := strings.Split(translation, "\n")
		translation = strings.TrimSpace(lines[0])
		fmt.Printf("[DEBUG] 只使用第一行: %s\n", translation)
	}
	
	// 检查翻译是否是拒绝回复
	if strings.Contains(translation, "I cannot") ||
	   strings.Contains(translation, "I can't") ||
	   strings.Contains(translation, "不能翻译") ||
	   strings.Contains(translation, "无法翻译") ||
	   strings.Contains(translation, "抱歉") {
		return "", fmt.Errorf("model refused to translate")
	}
	
	return translation, nil
}

// Name 返回提供商名称
func (p *OpenRouterProvider) Name() string {
	return "OpenRouter"
}

// ListModels 返回可用的模型列表
func (p *OpenRouterProvider) ListModels() ([]string, error) {
	return p.GetAvailableModels()
}

// TestConnection 测试OpenRouter连接
func (p *OpenRouterProvider) TestConnection() error {
	// OpenRouter提供了一个免费的/auth/key端点来验证API key
	// 这个端点不会消耗任何credits
	url := "https://openrouter.ai/api/v1/auth/key"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	
	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	
	// 使用共享的HTTP客户端
	client := GetSharedHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == 401 {
		return fmt.Errorf("invalid API key")
	}
	
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API test failed: %s", string(body))
	}
	
	return nil
}