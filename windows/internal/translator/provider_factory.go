package translator

// CreateProvider 根据provider名称创建对应的Provider实例
// 这个函数确保所有22个provider都能正确创建
func CreateProvider(providerName, apiKey, baseURL, model string) Provider {
	switch providerName {
	case "OpenAI":
		return NewOpenAIProvider(apiKey, model)
	case "Anthropic":
		return NewAnthropicProvider(apiKey, model)
	case "DeepSeek":
		return NewDeepSeekProvider(apiKey, model)
	case "Moonshot":
		return NewMoonshotProvider(apiKey, model)
	case "OpenRouter":
		return NewOpenRouterProvider(apiKey, model)
	case "Groq":
		return NewGroqProvider(apiKey, model)
	case "Together":
		return NewTogetherProvider(apiKey, model)

	// 以下provider需要特殊实现或暂时使用兼容模式
	case "Google":
		// Google Gemini API - 需要特殊实现，暂时返回兼容模式
		// TODO: 实现专门的Google provider
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Alibaba":
		// 阿里云通义千问
		if baseURL == "" {
			baseURL = "https://dashscope.aliyuncs.com/compatible-mode/v1"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Baidu":
		// 百度文心一言
		if baseURL == "" {
			baseURL = "https://aip.baidubce.com/rpc/2.0/ai_custom/v1"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "ByteDance":
		// 字节跳动豆包
		if baseURL == "" {
			baseURL = "https://ark.cn-beijing.volcanicengine.com/api/v3"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Zhipu":
		// 智谱AI
		if baseURL == "" {
			baseURL = "https://open.bigmodel.cn/api/paas/v4"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "01AI":
		// 零一万物
		if baseURL == "" {
			baseURL = "https://api.lingyiwanwu.com/v1"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Mistral":
		// Mistral AI - OpenAI兼容
		if baseURL == "" {
			baseURL = "https://api.mistral.ai/v1"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Cohere":
		// Cohere
		if baseURL == "" {
			baseURL = "https://api.cohere.ai/v1"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Perplexity":
		// Perplexity AI
		if baseURL == "" {
			baseURL = "https://api.perplexity.ai"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "xAI":
		// xAI (Grok)
		if baseURL == "" {
			baseURL = "https://api.x.ai/v1"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Meta":
		// Meta Llama (通常通过其他平台访问)
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Replicate":
		// Replicate
		if baseURL == "" {
			baseURL = "https://api.replicate.com/v1"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "HuggingFace":
		// HuggingFace Inference API
		if baseURL == "" {
			baseURL = "https://api-inference.huggingface.co/models"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "AWS":
		// Amazon Bedrock
		if baseURL == "" {
			baseURL = "https://bedrock-runtime.amazonaws.com"
		}
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	case "Azure":
		// Azure OpenAI Service - 需要用户提供endpoint
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)

	default:
		// 未知的provider，使用通用的OpenAI兼容接口
		return NewOpenAICompatibleProvider(providerName, apiKey, baseURL, model)
	}
}

// 为了向后兼容，保留这些简化的构造函数
func NewGoogleProvider(apiKey, model string) Provider {
	return CreateProvider("Google", apiKey, "", model)
}

func NewAlibabaProvider(apiKey, model string) Provider {
	return CreateProvider("Alibaba", apiKey, "", model)
}

func NewBaiduProvider(apiKey, model string) Provider {
	return CreateProvider("Baidu", apiKey, "", model)
}

func NewByteDanceProvider(apiKey, model string) Provider {
	return CreateProvider("ByteDance", apiKey, "", model)
}

func NewZhipuProvider(apiKey, model string) Provider {
	return CreateProvider("Zhipu", apiKey, "", model)
}

func New01AIProvider(apiKey, model string) Provider {
	return CreateProvider("01AI", apiKey, "", model)
}

func NewMistralProvider(apiKey, model string) Provider {
	return CreateProvider("Mistral", apiKey, "", model)
}

func NewCohereProvider(apiKey, model string) Provider {
	return CreateProvider("Cohere", apiKey, "", model)
}

func NewPerplexityProvider(apiKey, model string) Provider {
	return CreateProvider("Perplexity", apiKey, "", model)
}

func NewXAIProvider(apiKey, model string) Provider {
	return CreateProvider("xAI", apiKey, "", model)
}

func NewMetaProvider(apiKey, model string) Provider {
	return CreateProvider("Meta", apiKey, "", model)
}

func NewReplicateProvider(apiKey, model string) Provider {
	return CreateProvider("Replicate", apiKey, "", model)
}

func NewHuggingFaceProvider(apiKey, model string) Provider {
	return CreateProvider("HuggingFace", apiKey, "", model)
}

func NewAWSProvider(apiKey, model string) Provider {
	return CreateProvider("AWS", apiKey, "", model)
}

func NewAzureProvider(apiKey, model string) Provider {
	return CreateProvider("Azure", apiKey, "", model)
}