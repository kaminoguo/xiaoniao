package main

import (
	"pixel-translator/internal/translator"
)

// Prompt 本地Prompt结构（用于UI显示）
type Prompt struct {
	ID      string
	Name    string
	Content string
}

// GetAllPrompts 获取所有可用的prompts（从文件读取）
func GetAllPrompts() []Prompt {
	userPrompts := translator.GetUserPrompts()
	prompts := make([]Prompt, len(userPrompts))
	
	for i, up := range userPrompts {
		prompts[i] = Prompt{
			ID:      up.ID,
			Name:    up.Name,
			Content: up.Content,
		}
	}
	
	return prompts
}

// GetPromptByID 根据ID获取prompt
func GetPromptByID(id string) *Prompt {
	if up := translator.GetPromptByID(id); up != nil {
		return &Prompt{
			ID:      up.ID,
			Name:    up.Name,
			Content: up.Content,
		}
	}
	return nil
}

// AddPrompt 添加新prompt（写入文件）
func AddPrompt(id, name, content string) error {
	return translator.AddUserPrompt(translator.UserPrompt{
		ID:      id,
		Name:    name,
		Content: content,
	})
}

// UpdatePrompt 更新prompt（写入文件）
func UpdatePrompt(id, name, content string) error {
	return translator.UpdateUserPrompt(id, translator.UserPrompt{
		ID:      id,
		Name:    name,
		Content: content,
	})
}

// DeletePrompt 删除prompt（写入文件）
func DeletePrompt(id string) error {
	return translator.DeleteUserPrompt(id)
}

// ReloadPrompts 重新加载prompts
func ReloadPrompts() error {
	return translator.LoadUserPrompts()
}