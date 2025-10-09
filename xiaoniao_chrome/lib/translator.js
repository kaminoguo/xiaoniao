// Translation engine with Hybrid AI support
// Default: Chrome Built-in AI (Gemini Nano)
// Optional: Gemini API (user provides key)

import { buildSystemPrompt, getCurrentPrompt } from './prompts.js';

/**
 * Translation modes
 */
export const TranslationMode = {
  BUILTIN: 'builtin',  // Chrome Built-in AI (default)
  GEMINI: 'gemini'     // Gemini 2.5 Flash API
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
  try {
    // Check if Chrome Built-in AI is available
    if (!window.ai || !window.ai.languageModel) {
      throw new Error('Chrome Built-in AI not available. Please enable it in chrome://flags/#optimization-guide-on-device-model');
    }

    // Create language model session
    const session = await window.ai.languageModel.create({
      systemPrompt: systemPrompt,
      temperature: 0.3,
      topK: 3
    });

    // Translate
    const result = await session.prompt(text);

    // Clean up session
    session.destroy();

    return result.trim();
  } catch (error) {
    console.error('Built-in AI translation error:', error);
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
  if (!apiKey) {
    throw new Error('Gemini API key not configured');
  }

  const url = 'https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-exp:generateContent';

  try {
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
      throw new Error(`Gemini API error: ${errorData.error?.message || response.statusText}`);
    }

    const data = await response.json();
    const translatedText = data.candidates?.[0]?.content?.parts?.[0]?.text;

    if (!translatedText) {
      throw new Error('Invalid response from Gemini API');
    }

    return translatedText.trim();
  } catch (error) {
    console.error('Gemini API translation error:', error);
    throw error;
  }
}

/**
 * Main translation function
 * @param {string} text - Text to translate
 * @returns {Promise<string>} Translated text
 */
export async function translate(text) {
  if (!text || text.trim().length === 0) {
    throw new Error('Empty text provided');
  }

  // Get settings
  const settings = await getSettings();

  if (!settings.enabled) {
    throw new Error('Extension is disabled');
  }

  // Get current prompt
  const userPrompt = await getCurrentPrompt();
  const systemPrompt = buildSystemPrompt(userPrompt);

  // Translate based on mode
  if (settings.mode === TranslationMode.GEMINI && settings.apiKey) {
    return await translateWithGeminiAPI(text, systemPrompt, settings.apiKey);
  } else {
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
