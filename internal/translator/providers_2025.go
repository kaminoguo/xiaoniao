package translator

import (
	"strings"
)

// Provider模型定义 - 2025年8月最新
var ProviderModels = map[string][]string{
	"OpenAI": {
		// GPT-5 系列 (2025年8月发布)
		"gpt-5",
		"gpt-5-mini",
		"gpt-5-nano",
		
		// GPT-4.1 系列
		"gpt-4.1",
		"gpt-4.1-preview",
		
		// O系列推理模型
		"o1-preview",
		"o1-mini",
		"o3",
		"o3-mini",
		
		// GPT-4o 系列
		"gpt-4o",
		"gpt-4o-mini",
		"gpt-4o-search-preview",
		"gpt-4o-mini-search-preview",
		"gpt-4o-2024-08-06",
		"gpt-4o-2024-05-13",
		
		// GPT-4 系列
		"gpt-4-turbo",
		"gpt-4-turbo-preview",
		"gpt-4-turbo-2024-04-09",
		"gpt-4-1106-preview",
		"gpt-4-0125-preview",
		"gpt-4",
		"gpt-4-32k",
		
		// GPT-3.5 系列
		"gpt-3.5-turbo",
		"gpt-3.5-turbo-16k",
		"gpt-3.5-turbo-1106",
		"gpt-3.5-turbo-0125",
	},
	
	"Anthropic": {
		// Claude 3.5 系列
		"claude-3.5-opus",
		"claude-3.5-sonnet-latest",
		"claude-3.5-sonnet-20241022",
		"claude-3.5-haiku",
		
		// Claude 3 系列
		"claude-3-opus-20240229",
		"claude-3-opus-latest",
		"claude-3-sonnet-20240229",
		"claude-3-haiku-20240307",
		
		// Claude 2 系列
		"claude-2.1",
		"claude-2.0",
		"claude-instant-1.2",
	},
	
	"Google": {
		// Gemini 2.5 系列
		"gemini-2.5-pro",
		"gemini-2.5-flash",
		
		// Gemini 2.0 系列
		"gemini-2.0-flash-thinking-exp-1219",
		"gemini-2.0-flash-exp",
		"gemini-exp-1206",
		
		// Gemini 1.5 系列
		"gemini-1.5-pro",
		"gemini-1.5-pro-latest",
		"gemini-1.5-pro-002",
		"gemini-1.5-flash",
		"gemini-1.5-flash-latest",
		"gemini-1.5-flash-002",
		
		// Gemini 1.0 系列
		"gemini-1.0-pro",
		"gemini-pro",
		"gemini-pro-vision",
	},
	
	"DeepSeek": {
		// R1 推理模型
		"deepseek-r1",
		"deepseek-r1-distill-qwen-32b",
		"deepseek-r1-distill-llama-70b",
		
		// V3 系列
		"deepseek-v3",
		
		// Chat 系列
		"deepseek-chat",
		"deepseek-coder",
	},
	
	"Moonshot": {
		"moonshot-v1-8k",
		"moonshot-v1-16k",
		"moonshot-v1-32k",
		"moonshot-v1-128k",
	},
	
	"Alibaba": {
		// 通义千问系列
		"qwen-turbo",
		"qwen-plus",
		"qwen-max",
		"qwen-max-1201",
		"qwen-max-longcontext",
		
		// Qwen2.5系列
		"qwen2.5-72b-instruct",
		"qwen2.5-32b-instruct",
		"qwen2.5-14b-instruct",
		"qwen2.5-7b-instruct",
		
		// 多模态模型
		"qwen-vl-plus",
		"qwen-vl-max",
		"qwen-audio-turbo",
	},
	
	"Baidu": {
		// 文心一言系列
		"ERNIE-4.0-8K",
		"ERNIE-4.0-Turbo-8K",
		"ERNIE-3.5-8K",
		"ERNIE-3.5-128K",
		"ERNIE-Speed-8K",
		"ERNIE-Speed-128K",
		"ERNIE-Lite-8K",
		"ERNIE-Tiny-8K",
	},
	
	"ByteDance": {
		// 豆包系列
		"Doubao-lite-4k",
		"Doubao-lite-32k",
		"Doubao-lite-128k",
		"Doubao-pro-4k",
		"Doubao-pro-32k",
		"Doubao-pro-128k",
	},
	
	"Zhipu": {
		// 智谱GLM系列
		"glm-4-plus",
		"glm-4-0520",
		"glm-4-air",
		"glm-4-airx",
		"glm-4-flash",
		"glm-4v-plus",
		"glm-4v",
		"chatglm-turbo",
	},
	
	"01AI": {
		// 零一万物
		"yi-lightning",
		"yi-large",
		"yi-large-turbo",
		"yi-medium",
		"yi-medium-200k",
		"yi-spark",
		"yi-large-fc",
		"yi-large-rag",
		"yi-vision",
	},
	
	"Mistral": {
		"mistral-large-latest",
		"mistral-large-2407",
		"mistral-medium-latest",
		"mistral-small-latest",
		"codestral-latest",
		"open-mistral-nemo",
		"open-codestral-mamba",
	},
	
	"Cohere": {
		"command-r-plus",
		"command-r",
		"command",
		"command-light",
		"c4ai-aya-expanse-8b",
		"c4ai-aya-expanse-32b",
	},
	
	"Perplexity": {
		"llama-3.1-sonar-small-128k-online",
		"llama-3.1-sonar-large-128k-online",
		"llama-3.1-sonar-huge-128k-online",
	},
	
	"xAI": {
		"grok-2",
		"grok-2-mini",
		"grok-beta",
	},
	
	"Meta": {
		// Llama 3.2 系列
		"llama-3.2-90b-text-preview",
		"llama-3.2-11b-text-preview",
		"llama-3.2-3b-preview",
		"llama-3.2-1b-preview",
		
		// Llama 3.1 系列
		"llama-3.1-405b-instruct",
		"llama-3.1-70b-instruct",
		"llama-3.1-8b-instruct",
	},
	
	"OpenRouter": {
		// OpenRouter 聚合了多个provider的模型
		"auto", // 自动选择最佳模型
		"openai/gpt-4o", "openai/gpt-4o-mini", "openai/gpt-4-turbo",
		"anthropic/claude-3.5-sonnet", "anthropic/claude-3-opus", 
		"google/gemini-pro-1.5", "google/gemini-flash-1.5",
		"meta-llama/llama-3.1-405b", "meta-llama/llama-3.1-70b",
		"mistralai/mistral-large", "mistralai/mixtral-8x7b",
		"deepseek/deepseek-v3", "deepseek/deepseek-r1",
		"perplexity/llama-3.1-sonar-large",
		"cohere/command-r-plus",
	},
	
	"Groq": {
		// Groq 专注于高速推理
		"llama-3.1-405b-reasoning", "llama-3.1-70b-versatile", "llama-3.1-8b-instant",
		"llama3-groq-70b-8192-tool-use-preview", "llama3-groq-8b-8192-tool-use-preview",
		"mixtral-8x7b-32768", "gemma-7b-it", "gemma2-9b-it",
	},
	
	"Together": {
		// Together AI 提供多种开源模型
		"meta-llama/Meta-Llama-3.1-405B-Instruct-Turbo",
		"meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo",
		"meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo",
		"mistralai/Mixtral-8x7B-Instruct-v0.1",
		"mistralai/Mistral-7B-Instruct-v0.3",
		"Qwen/Qwen2.5-72B-Instruct-Turbo",
		"deepseek-ai/deepseek-llm-67b-chat",
	},
	
	"Replicate": {
		// Replicate 提供多种模型API
		"meta/llama-2-70b-chat", "meta/llama-2-13b-chat", "meta/llama-2-7b-chat",
		"mistralai/mistral-7b-instruct-v0.2",
		"stability-ai/sdxl", // 图像生成
	},
	
	"HuggingFace": {
		// HuggingFace Inference API
		"meta-llama/Llama-2-70b-chat-hf", "meta-llama/Llama-2-13b-chat-hf",
		"mistralai/Mistral-7B-Instruct-v0.2",
		"google/flan-t5-xxl", "google/flan-ul2",
		"bigscience/bloom", "bigscience/bloomz",
	},
	
	"AWS": {
		// Amazon Bedrock
		"anthropic.claude-3-sonnet", "anthropic.claude-3-haiku",
		"amazon.titan-text-express-v1", "amazon.titan-text-lite-v1",
		"ai21.j2-ultra-v1", "ai21.j2-mid-v1",
		"cohere.command-text-v14", "cohere.command-light-text-v14",
		"meta.llama3-70b-instruct-v1", "meta.llama3-8b-instruct-v1",
		"mistral.mistral-large-2407-v1", "mistral.mixtral-8x7b-instruct-v0",
	},
	
	"Azure": {
		// Azure OpenAI Service
		"gpt-4", "gpt-4-turbo", "gpt-35-turbo",
		"text-embedding-ada-002", "text-embedding-3-small", "text-embedding-3-large",
	},
}

// 检测Provider类型的特征
var ProviderKeyPatterns = map[string][]string{
	"OpenAI": {
		"sk-",           // 标准OpenAI key
		"sk-proj-",      // 项目key
		"sess-",         // 会话key
	},
	"Anthropic": {
		"sk-ant-",       // Anthropic key
	},
	"Google": {
		"AIza",          // Google API key前缀
	},
	"Azure": {
		// Azure使用endpoint URL而不是key pattern
	},
	"DeepSeek": {
		"sk-",           // DeepSeek使用类似OpenAI的格式
	},
	"Moonshot": {
		"sk-",           // Moonshot使用类似OpenAI的格式
	},
	"Alibaba": {
		"sk-",           // 阿里云使用类似格式
	},
	"OpenRouter": {
		"sk-or-",        // OpenRouter特定前缀
	},
	"Groq": {
		"gsk_",          // Groq API key前缀
	},
	"Together": {
		"sk-",           // Together AI使用类似格式
		"together_",     // 另一种格式
	},
	"Replicate": {
		"r8_",           // Replicate API token前缀
	},
	"HuggingFace": {
		"hf_",           // HuggingFace API token前缀
	},
	"Cohere": {
		"sk-",           // Cohere使用类似格式
	},
	"Perplexity": {
		"pplx-",         // Perplexity API key前缀
	},
	"Mistral": {
		"sk-",           // Mistral使用类似格式
	},
	"AWS": {
		"AKIA",          // AWS Access Key ID前缀
	},
}

// 根据API Key自动检测Provider
func DetectProviderByKey(apiKey string) string {
	// 使用新的provider_registry.go中的函数
	return DetectProviderByAPIKey(apiKey)
}

// 搜索模型（支持模糊搜索）
func SearchModels(provider string, query string) []string {
	models, exists := ProviderModels[provider]
	if !exists {
		return []string{}
	}
	
	if query == "" {
		return models
	}
	
	query = strings.ToLower(query)
	var results []string
	
	for _, model := range models {
		modelLower := strings.ToLower(model)
		if strings.Contains(modelLower, query) {
			results = append(results, model)
		}
	}
	
	return results
}

// 获取所有支持的Provider
func GetSupportedProviders() []string {
	var providers []string
	for provider := range ProviderModels {
		providers = append(providers, provider)
	}
	return providers
}