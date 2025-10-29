package translator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// UserPrompt 用户prompt定义
type UserPrompt struct {
	ID      string `json:"id"`      // 唯一标识符
	Name    string `json:"name"`    // 显示名称
	Content string `json:"content"` // Prompt内容
}

var (
	userPrompts []UserPrompt
	promptMutex sync.RWMutex
	promptsFile string
)

func init() {
	// 设置prompt文件路径
	homeDir, _ := os.UserHomeDir()
	promptsFile = filepath.Join(homeDir, ".config", "xiaoniao", "prompts.json")
	
	// 加载prompts
	LoadUserPrompts()
}

// LoadUserPrompts 从文件加载用户prompts
func LoadUserPrompts() error {
	promptMutex.Lock()
	defer promptMutex.Unlock()
	
	// 如果文件不存在，创建默认的
	if _, err := os.Stat(promptsFile); os.IsNotExist(err) {
		userPrompts = getDefaultPrompts()
		return SaveUserPromptsLocked()
	}
	
	// 读取文件
	data, err := os.ReadFile(promptsFile)
	if err != nil {
		return fmt.Errorf("failed to read prompts file: %w", err)
	}
	
	// 解析JSON
	if err := json.Unmarshal(data, &userPrompts); err != nil {
		return fmt.Errorf("failed to parse prompts: %w", err)
	}
	
	return nil
}

// SaveUserPrompts 保存用户prompts到文件
func SaveUserPrompts() error {
	promptMutex.Lock()
	defer promptMutex.Unlock()
	return SaveUserPromptsLocked()
}

// SaveUserPromptsLocked 保存prompts（已加锁）
func SaveUserPromptsLocked() error {
	// 确保目录存在
	dir := filepath.Dir(promptsFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	
	// 序列化为JSON
	data, err := json.MarshalIndent(userPrompts, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal prompts: %w", err)
	}
	
	// 写入文件
	if err := os.WriteFile(promptsFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write prompts file: %w", err)
	}
	
	return nil
}

// GetUserPrompts 获取所有用户prompts
func GetUserPrompts() []UserPrompt {
	promptMutex.RLock()
	defer promptMutex.RUnlock()
	
	// 返回副本，避免外部修改
	result := make([]UserPrompt, len(userPrompts))
	copy(result, userPrompts)
	return result
}

// GetPromptByID 根据ID获取prompt
func GetPromptByID(id string) *UserPrompt {
	promptMutex.RLock()
	defer promptMutex.RUnlock()
	
	for _, p := range userPrompts {
		if p.ID == id {
			// 返回副本
			result := p
			return &result
		}
	}
	return nil
}

// AddUserPrompt 添加新的用户prompt
func AddUserPrompt(prompt UserPrompt) error {
	promptMutex.Lock()
	defer promptMutex.Unlock()
	
	// 检查ID是否已存在
	for _, p := range userPrompts {
		if p.ID == prompt.ID {
			return fmt.Errorf("prompt with ID %s already exists", prompt.ID)
		}
	}
	
	userPrompts = append(userPrompts, prompt)
	return SaveUserPromptsLocked()
}

// UpdateUserPrompt 更新用户prompt
func UpdateUserPrompt(id string, prompt UserPrompt) error {
	promptMutex.Lock()
	defer promptMutex.Unlock()
	
	for i, p := range userPrompts {
		if p.ID == id {
			userPrompts[i] = prompt
			return SaveUserPromptsLocked()
		}
	}
	
	return fmt.Errorf("prompt with ID %s not found", id)
}

// DeleteUserPrompt 删除用户prompt
func DeleteUserPrompt(id string) error {
	promptMutex.Lock()
	defer promptMutex.Unlock()
	
	for i, p := range userPrompts {
		if p.ID == id {
			// 删除元素
			userPrompts = append(userPrompts[:i], userPrompts[i+1:]...)
			return SaveUserPromptsLocked()
		}
	}
	
	return fmt.Errorf("prompt with ID %s not found", id)
}

// getDefaultPrompts 返回默认的prompts（空列表，让用户自己添加）
func getDefaultPrompts() []UserPrompt {
	return []UserPrompt{
		// 只保留一个基础的直译prompt
		{
			ID:      "direct",
			Name:    "直译",
			Content: "Translate to Chinese directly and accurately.",
		},
	}
}