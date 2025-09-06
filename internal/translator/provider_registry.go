package translator

import "strings"

// ProviderConfig 定义每个Provider的配置
type ProviderConfig struct {
	Name        string
	BaseURL     string
	KeyPrefix   []string
	NeedsAuth   bool
	ModelPrefix string // 某些provider需要添加模型前缀，如OpenRouter的"openai/"
}

// ProviderRegistry 所有支持的Provider配置
var ProviderRegistry = map[string]ProviderConfig{
	"OpenRouter": {
		Name:      "OpenRouter",
		BaseURL:   "https://openrouter.ai/api/v1",
		KeyPrefix: []string{"sk-or-"},
		NeedsAuth: true,
	},
	"OpenAI": {
		Name:      "OpenAI",
		BaseURL:   "https://api.openai.com/v1",
		KeyPrefix: []string{"sk-", "sk-proj-"},
		NeedsAuth: true,
	},
	"Anthropic": {
		Name:      "Anthropic",
		BaseURL:   "https://api.anthropic.com/v1",
		KeyPrefix: []string{"sk-ant-"},
		NeedsAuth: true,
	},
	"Groq": {
		Name:      "Groq",
		BaseURL:   "https://api.groq.com/openai/v1",
		KeyPrefix: []string{"gsk_"},
		NeedsAuth: true,
	},
	"Together": {
		Name:      "Together",
		BaseURL:   "https://api.together.xyz/v1",
		KeyPrefix: []string{"sk-", "together_"},
		NeedsAuth: true,
	},
	"DeepSeek": {
		Name:      "DeepSeek",
		BaseURL:   "https://api.deepseek.com/v1",
		KeyPrefix: []string{"sk-"},
		NeedsAuth: true,
	},
	"Perplexity": {
		Name:      "Perplexity",
		BaseURL:   "https://api.perplexity.ai",
		KeyPrefix: []string{"pplx-"},
		NeedsAuth: true,
	},
	"Mistral": {
		Name:      "Mistral",
		BaseURL:   "https://api.mistral.ai/v1",
		KeyPrefix: []string{"sk-"},
		NeedsAuth: true,
	},
	"Cohere": {
		Name:      "Cohere",
		BaseURL:   "https://api.cohere.ai/v1",
		KeyPrefix: []string{"sk-"},
		NeedsAuth: true,
	},
	"Google": {
		Name:      "Google",
		BaseURL:   "https://generativelanguage.googleapis.com/v1beta",
		KeyPrefix: []string{"AIza"},
		NeedsAuth: true,
	},
	"HuggingFace": {
		Name:      "HuggingFace",
		BaseURL:   "https://api-inference.huggingface.co/models",
		KeyPrefix: []string{"hf_"},
		NeedsAuth: true,
	},
	"Replicate": {
		Name:      "Replicate",
		BaseURL:   "https://api.replicate.com/v1",
		KeyPrefix: []string{"r8_"},
		NeedsAuth: true,
	},
	"Moonshot": {
		Name:      "Moonshot",
		BaseURL:   "https://api.moonshot.cn/v1",
		KeyPrefix: []string{"sk-"},
		NeedsAuth: true,
	},
	"Zhipu": {
		Name:      "Zhipu",
		BaseURL:   "https://open.bigmodel.cn/api/paas/v4",
		KeyPrefix: []string{""},
		NeedsAuth: true,
	},
	"Baidu": {
		Name:      "Baidu",
		BaseURL:   "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop",
		KeyPrefix: []string{""},
		NeedsAuth: true,
	},
	"Alibaba": {
		Name:      "Alibaba",
		BaseURL:   "https://dashscope.aliyuncs.com/api/v1",
		KeyPrefix: []string{"sk-"},
		NeedsAuth: true,
	},
	"Azure": {
		Name:      "Azure",
		BaseURL:   "", // Azure需要自定义endpoint
		KeyPrefix: []string{""},
		NeedsAuth: true,
	},
	"AWS": {
		Name:      "AWS",
		BaseURL:   "", // AWS Bedrock使用不同的认证机制
		KeyPrefix: []string{"AKIA"},
		NeedsAuth: true,
	},
}

// GetProviderConfig 获取Provider配置
func GetProviderConfig(provider string) (ProviderConfig, bool) {
	config, exists := ProviderRegistry[provider]
	return config, exists
}

// DetectProviderByAPIKey 根据API Key检测Provider
func DetectProviderByAPIKey(apiKey string) string {
	// 优先检查特定前缀
	for name, config := range ProviderRegistry {
		for _, prefix := range config.KeyPrefix {
			if prefix != "" && strings.HasPrefix(apiKey, prefix) {
				// 特殊处理sk-前缀
				if prefix == "sk-" {
					// 进一步判断
					if strings.HasPrefix(apiKey, "sk-ant-") {
						return "Anthropic"
					}
					if strings.HasPrefix(apiKey, "sk-or-") {
						return "OpenRouter"
					}
					if strings.HasPrefix(apiKey, "sk-proj-") {
						return "OpenAI"
					}
					// 根据长度判断
					if len(apiKey) == 32 {
						return "DeepSeek"
					}
					if len(apiKey) > 50 {
						return "OpenAI"
					}
				}
				return name
			}
		}
	}
	return ""
}