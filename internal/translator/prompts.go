package translator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Prompt represents a translation prompt
type Prompt struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	IsBuiltin bool   `json:"isBuiltin"`
}

// PromptManager manages translation prompts
type PromptManager struct {
	mu       sync.RWMutex
	prompts  []Prompt
	custom   []Prompt
	dataPath string
}

// NewPromptManager creates a new prompt manager
func NewPromptManager(dataPath string) *PromptManager {
	pm := &PromptManager{
		dataPath: dataPath,
		prompts:  make([]Prompt, 0),
		custom:   make([]Prompt, 0),
	}

	// Initialize with built-in prompts
	pm.loadBuiltinPrompts()
	
	// Load custom prompts from file
	pm.loadCustomPrompts()

	return pm
}

// loadBuiltinPrompts loads all built-in prompts
func (pm *PromptManager) loadBuiltinPrompts() {
	builtinPrompts := []Prompt{
		// Basic translations
		{ID: "direct", Name: "直译", Content: "将以下文本翻译成中文：{{text}}", IsBuiltin: true},
		{ID: "to_english", Name: "译成英文", Content: "Translate the following text to English: {{text}}", IsBuiltin: true},
		{ID: "to_japanese", Name: "译成日文", Content: "次のテキストを日本語に翻訳してください：{{text}}", IsBuiltin: true},
		
		// Style-specific translations
		{ID: "5ch", Name: "5ch论坛风格", Content: "将以下内容翻译成5ch论坛风格的中文（使用网络用语、颜文字、2ch黑话等）：{{text}}", IsBuiltin: true},
		{ID: "business", Name: "商务英语", Content: "Translate to formal business English suitable for professional communication: {{text}}", IsBuiltin: true},
		{ID: "academic", Name: "学术风格", Content: "翻译为学术论文风格的中文，使用专业术语和正式表达：{{text}}", IsBuiltin: true},
		
		// Creative translations
		{ID: "literary", Name: "文学翻译", Content: "以优美的文学语言翻译，保持原文的艺术性和感染力：{{text}}", IsBuiltin: true},
		{ID: "poetry", Name: "诗意翻译", Content: "用诗意的语言翻译，注重韵律和意境：{{text}}", IsBuiltin: true},
		{ID: "humorous", Name: "幽默风格", Content: "用幽默诙谐的方式翻译，适当加入笑点：{{text}}", IsBuiltin: true},
		
		// Technical translations
		{ID: "programming", Name: "编程术语", Content: "翻译编程相关内容，保留技术术语的准确性：{{text}}", IsBuiltin: true},
		{ID: "medical", Name: "医学翻译", Content: "使用专业医学术语准确翻译：{{text}}", IsBuiltin: true},
		{ID: "legal", Name: "法律翻译", Content: "使用法律专业术语进行精确翻译：{{text}}", IsBuiltin: true},
		
		// Internet culture
		{ID: "meme", Name: "梗图文化", Content: "用网络流行梗和表情包文化的方式翻译：{{text}}", IsBuiltin: true},
		{ID: "twitch", Name: "直播弹幕", Content: "翻译成直播间弹幕风格（KEKW、PogChamp等）：{{text}}", IsBuiltin: true},
		{ID: "discord", Name: "Discord风格", Content: "翻译成Discord聊天风格，包含表情符号：{{text}}", IsBuiltin: true},
		
		// Regional dialects
		{ID: "beijing", Name: "北京话", Content: "翻译成地道的北京话，带儿化音：{{text}}", IsBuiltin: true},
		{ID: "cantonese", Name: "粤语", Content: "翻译成广东话/粤语：{{text}}", IsBuiltin: true},
		{ID: "taiwan", Name: "台湾腔", Content: "翻译成台湾繁体中文用语：{{text}}", IsBuiltin: true},
		
		// Age-specific
		{ID: "genz", Name: "Z世代用语", Content: "用Z世代流行语翻译（yyds、绝绝子等）：{{text}}", IsBuiltin: true},
		{ID: "classical", Name: "文言文", Content: "翻译成古典文言文：{{text}}", IsBuiltin: true},
		{ID: "kids", Name: "儿童语言", Content: "用简单易懂的儿童语言翻译：{{text}}", IsBuiltin: true},
		
		// Gaming
		{ID: "game_guide", Name: "游戏攻略", Content: "翻译游戏攻略，保留游戏术语：{{text}}", IsBuiltin: true},
		{ID: "esports", Name: "电竞解说", Content: "用电竞解说风格翻译，激情澎湃：{{text}}", IsBuiltin: true},
		{ID: "rpg", Name: "RPG风格", Content: "用RPG游戏的叙事风格翻译：{{text}}", IsBuiltin: true},
		
		// Social media
		{ID: "twitter", Name: "推特风格", Content: "翻译成推特风格，简洁有力，可加标签：{{text}}", IsBuiltin: true},
		{ID: "xiaohongshu", Name: "小红书风格", Content: "翻译成小红书风格，加emoji和话题标签：{{text}}", IsBuiltin: true},
		{ID: "weibo", Name: "微博体", Content: "翻译成微博风格，可加超话标签：{{text}}", IsBuiltin: true},
		
		// Special purposes
		{ID: "subtitle", Name: "字幕翻译", Content: "翻译成适合字幕的简洁文本，注意长度：{{text}}", IsBuiltin: true},
		{ID: "summary", Name: "摘要翻译", Content: "翻译并总结要点，突出关键信息：{{text}}", IsBuiltin: true},
		{ID: "explain", Name: "解释翻译", Content: "翻译并解释难懂的概念和背景：{{text}}", IsBuiltin: true},
		
		// Tone adjustments  
		{ID: "polite", Name: "礼貌用语", Content: "用非常礼貌和正式的语言翻译：{{text}}", IsBuiltin: true},
		{ID: "casual", Name: "口语化", Content: "用轻松的口语翻译，像朋友聊天：{{text}}", IsBuiltin: true},
		{ID: "aggressive", Name: "战斗民族", Content: "用强硬激进的语气翻译（战斗民族风格）：{{text}}", IsBuiltin: true},
		
		// Code and tech
		{ID: "git_commit", Name: "Git提交", Content: "翻译成规范的Git commit message格式：{{text}}", IsBuiltin: true},
		{ID: "api_doc", Name: "API文档", Content: "翻译成专业的API文档风格：{{text}}", IsBuiltin: true},
		{ID: "code_comment", Name: "代码注释", Content: "翻译成适合代码注释的简洁说明：{{text}}", IsBuiltin: true},
	}

	pm.prompts = builtinPrompts
}

// GetAll returns all prompts (builtin + custom)
func (pm *PromptManager) GetAll() []Prompt {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	allPrompts := make([]Prompt, 0, len(pm.prompts)+len(pm.custom))
	allPrompts = append(allPrompts, pm.prompts...)
	allPrompts = append(allPrompts, pm.custom...)
	
	return allPrompts
}

// GetByID returns a prompt by ID
func (pm *PromptManager) GetByID(id string) (*Prompt, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	// Check builtin prompts
	for _, p := range pm.prompts {
		if p.ID == id {
			return &p, nil
		}
	}

	// Check custom prompts
	for _, p := range pm.custom {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("prompt not found: %s", id)
}

// AddCustom adds a custom prompt
func (pm *PromptManager) AddCustom(name, content string) (*Prompt, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Generate unique ID
	id := fmt.Sprintf("custom_%d", len(pm.custom)+1)
	
	// Check for duplicate ID
	for _, p := range pm.custom {
		if p.ID == id {
			id = fmt.Sprintf("custom_%d_%d", len(pm.custom)+1, os.Getpid())
		}
	}

	prompt := Prompt{
		ID:        id,
		Name:      name,
		Content:   content,
		IsBuiltin: false,
	}

	pm.custom = append(pm.custom, prompt)
	
	// Save to file
	if err := pm.saveCustomPrompts(); err != nil {
		return nil, err
	}

	return &prompt, nil
}

// UpdateCustom updates a custom prompt
func (pm *PromptManager) UpdateCustom(id, name, content string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for i, p := range pm.custom {
		if p.ID == id {
			pm.custom[i].Name = name
			pm.custom[i].Content = content
			return pm.saveCustomPrompts()
		}
	}

	return fmt.Errorf("custom prompt not found: %s", id)
}

// DeleteCustom deletes a custom prompt
func (pm *PromptManager) DeleteCustom(id string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for i, p := range pm.custom {
		if p.ID == id {
			// Remove from slice
			pm.custom = append(pm.custom[:i], pm.custom[i+1:]...)
			return pm.saveCustomPrompts()
		}
	}

	return fmt.Errorf("custom prompt not found: %s", id)
}

// loadCustomPrompts loads custom prompts from file
func (pm *PromptManager) loadCustomPrompts() error {
	if pm.dataPath == "" {
		return nil
	}

	filePath := filepath.Join(pm.dataPath, "custom_prompts.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, that's ok
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &pm.custom)
}

// saveCustomPrompts saves custom prompts to file
func (pm *PromptManager) saveCustomPrompts() error {
	if pm.dataPath == "" {
		return nil
	}

	// Ensure directory exists
	if err := os.MkdirAll(pm.dataPath, 0755); err != nil {
		return err
	}

	filePath := filepath.Join(pm.dataPath, "custom_prompts.json")
	
	data, err := json.MarshalIndent(pm.custom, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// ExportPrompts exports all prompts to JSON
func (pm *PromptManager) ExportPrompts() ([]byte, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	allPrompts := pm.GetAll()
	return json.MarshalIndent(allPrompts, "", "  ")
}

// ImportPrompts imports prompts from JSON
func (pm *PromptManager) ImportPrompts(data []byte) error {
	var prompts []Prompt
	if err := json.Unmarshal(data, &prompts); err != nil {
		return err
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Add non-builtin prompts as custom
	for _, p := range prompts {
		if !p.IsBuiltin {
			// Check if already exists
			exists := false
			for _, existing := range pm.custom {
				if existing.ID == p.ID {
					exists = true
					break
				}
			}
			
			if !exists {
				pm.custom = append(pm.custom, p)
			}
		}
	}

	return pm.saveCustomPrompts()
}