// Translation engine with Hybrid AI support
// Default: Chrome Built-in AI (Gemini Nano)
// Optional: Gemini API (user provides key)

import { buildSystemPrompt, getCurrentPrompt } from './prompts.js';

/**
 * Translation modes
 */
export const TranslationMode = {
  BUILTIN: 'builtin',     // Chrome Built-in AI (default)
  GEMINI: 'gemini',       // Gemini 2.5 Flash API
  OPENROUTER: 'openrouter' // OpenRouter API
};

/**
 * Get translation settings from storage
 * @returns {Promise<Object>} Settings object
 */
async function getSettings() {
  const result = await chrome.storage.sync.get(['translationMode', 'geminiApiKey', 'extensionEnabled']);
  return {
    mode: result.translationMode || TranslationMode.BUILTIN,
    apiKey: result.geminiApiKey || '',
    enabled: result.extensionEnabled !== false // default true
  };
}

/**
 * Translate using Chrome Built-in AI (Gemini Nano)
 * @param {string} text - Text to translate
 * @param {string} systemPrompt - System prompt
 * @returns {Promise<string>} Translated text
 */
async function translateWithBuiltinAI(text, systemPrompt) {
  console.log('[Translator] Using Chrome Built-in AI');
  try {
    // Check if Chrome Built-in AI is available
    if (!self.ai || !self.ai.languageModel) {
      throw new Error('Chrome Built-in AI not available. Please enable it in chrome://flags/#optimization-guide-on-device-model and chrome://flags/#prompt-api-for-gemini-nano');
    }

    console.log('[Translator] Creating language model session...');
    // Create language model session
    const session = await self.ai.languageModel.create({
      systemPrompt: systemPrompt,
      temperature: 0.3,
      topK: 3
    });

    console.log('[Translator] Translating text...');
    // Translate
    const result = await session.prompt(text);

    console.log('[Translator] Translation result:', result.substring(0, 100));
    // Clean up session
    session.destroy();

    return result.trim();
  } catch (error) {
    console.error('[Translator] Built-in AI translation error:', error);
    throw error;
  }
}

/**
 * Translate using Gemini API
 * @param {string} text - Text to translate
 * @param {string} systemPrompt - System prompt
 * @param {string} apiKey - Gemini API key
 * @returns {Promise<string>} Translated text
 */
async function translateWithGeminiAPI(text, systemPrompt, apiKey) {
  console.log('[Translator] Using Gemini API');
  if (!apiKey) {
    throw new Error('Gemini API key not configured');
  }

  const url = 'https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-exp:generateContent';

  try {
    console.log('[Translator] Sending request to Gemini API...');
    const response = await fetch(`${url}?key=${apiKey}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        contents: [{
          parts: [{
            text: text
          }]
        }],
        systemInstruction: {
          parts: [{
            text: systemPrompt
          }]
        },
        generationConfig: {
          temperature: 0.3,
          topK: 3,
          topP: 0.95,
          maxOutputTokens: 2048
        }
      })
    });

    if (!response.ok) {
      const errorData = await response.json();
      console.error('[Translator] Gemini API error response:', errorData);
      throw new Error(`Gemini API error: ${errorData.error?.message || response.statusText}`);
    }

    const data = await response.json();
    console.log('[Translator] Gemini API response received');
    const translatedText = data.candidates?.[0]?.content?.parts?.[0]?.text;

    if (!translatedText) {
      console.error('[Translator] Invalid Gemini API response structure:', data);
      throw new Error('Invalid response from Gemini API');
    }

    console.log('[Translator] Translation result:', translatedText.substring(0, 100));
    return translatedText.trim();
  } catch (error) {
    console.error('[Translator] Gemini API translation error:', error);
    throw error;
  }
}

/**
 * Translate using OpenRouter API
 * @param {string} text - Text to translate
 * @param {string} systemPrompt - System prompt
 * @param {string} apiKey - OpenRouter API key
 * @returns {Promise<string>} Translated text
 */
async function translateWithOpenRouter(text, systemPrompt, apiKey) {
  console.log('[Translator] Using OpenRouter API');
  if (!apiKey) {
    throw new Error('OpenRouter API key not configured');
  }

  const url = 'https://openrouter.ai/api/v1/chat/completions';

  try {
    console.log('[Translator] Sending request to OpenRouter...');
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${apiKey}`,
        'Content-Type': 'application/json',
        'HTTP-Referer': 'https://xiaoniao-chrome-extension',
        'X-Title': 'Xiaoniao Chrome Extension'
      },
      body: JSON.stringify({
        model: 'google/gemini-2.0-flash-exp:free',
        messages: [
          {
            role: 'system',
            content: systemPrompt
          },
          {
            role: 'user',
            content: text
          }
        ],
        temperature: 0.3,
        max_tokens: 2048
      })
    });

    if (!response.ok) {
      const errorData = await response.json();
      console.error('[Translator] OpenRouter error response:', errorData);
      throw new Error(`OpenRouter API error: ${errorData.error?.message || response.statusText}`);
    }

    const data = await response.json();
    console.log('[Translator] OpenRouter response received');
    const translatedText = data.choices?.[0]?.message?.content;

    if (!translatedText) {
      console.error('[Translator] Invalid OpenRouter response structure:', data);
      throw new Error('Invalid response from OpenRouter');
    }

    console.log('[Translator] Translation result:', translatedText.substring(0, 100));
    return translatedText.trim();
  } catch (error) {
    console.error('[Translator] OpenRouter translation error:', error);
    throw error;
  }
}

/**
 * Main translation function
 * @param {string} text - Text to translate
 * @returns {Promise<string>} Translated text
 */
export async function translate(text) {
  console.log('[Translator] translate() called with text:', text.substring(0, 50));

  if (!text || text.trim().length === 0) {
    throw new Error('Empty text provided');
  }

  // Get settings
  const settings = await getSettings();
  console.log('[Translator] Settings:', settings);

  if (!settings.enabled) {
    throw new Error('Extension is disabled');
  }

  // Get current prompt
  const userPrompt = await getCurrentPrompt();
  console.log('[Translator] User prompt:', userPrompt);
  const systemPrompt = buildSystemPrompt(userPrompt);

  // Translate based on mode
  if (settings.mode === TranslationMode.OPENROUTER && settings.apiKey) {
    console.log('[Translator] Using OpenRouter mode');
    return await translateWithOpenRouter(text, systemPrompt, settings.apiKey);
  } else if (settings.mode === TranslationMode.GEMINI && settings.apiKey) {
    console.log('[Translator] Using Gemini API mode');
    return await translateWithGeminiAPI(text, systemPrompt, settings.apiKey);
  } else {
    console.log('[Translator] Using Built-in AI mode');
    // Default to built-in AI
    return await translateWithBuiltinAI(text, systemPrompt);
  }
}

/**
 * Check if Chrome Built-in AI is available
 * @returns {Promise<boolean>} True if available
 */
export async function isBuiltinAIAvailable() {
  try {
    return !!(window.ai && window.ai.languageModel);
  } catch (error) {
    return false;
  }
}

/**
 * Test Gemini API key
 * @param {string} apiKey - API key to test
 * @returns {Promise<boolean>} True if valid
 */
export async function testGeminiAPIKey(apiKey) {
  try {
    await translateWithGeminiAPI('Hello', 'Translate to Chinese', apiKey);
    return true;
  } catch (error) {
    console.error('API key test failed:', error);
    return false;
  }
}
