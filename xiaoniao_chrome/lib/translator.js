// Translation engine with Hybrid AI support
// Default: Chrome Built-in AI (Gemini Nano)
// Optional: Gemini API (user provides key)

import { buildSystemPrompt, buildSimpleSystemPrompt, getCurrentPrompt } from './prompts.js';

/**
 * Translation modes
 */
export const TranslationMode = {
  BUILTIN: 'builtin',     // Chrome Built-in AI (default)
  GEMINI: 'gemini',       // Gemini 2.5 Flash API
  OPENROUTER: 'openrouter', // OpenRouter API
  FREETRY: 'freetry'      // Free Try with DeepSeek V3.1
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
    // Check if LanguageModel API is available (Chrome 141+)
    if (typeof LanguageModel === 'undefined') {
      throw new Error('Chrome Built-in AI not available. Please enable it in chrome://flags/#prompt-api-for-gemini-nano and restart Chrome');
    }

    // Check availability
    console.log('[Translator] Checking LanguageModel availability...');
    const availability = await LanguageModel.availability();
    console.log('[Translator] Availability status:', availability);

    if (availability === 'no') {
      throw new Error('LanguageModel not available. Please enable it in chrome://flags/#prompt-api-for-gemini-nano and restart Chrome');
    }

    // Status can be "readily" (ready) or "available" (will download on first use)
    if (availability !== 'readily' && availability !== 'available') {
      throw new Error(`LanguageModel status: ${availability}`);
    }

    console.log('[Translator] Creating language model session...');
    // Create language model session (NEW API)
    // If status is "available", calling create() will trigger download
    const session = await LanguageModel.create({
      systemPrompt: systemPrompt,
      temperature: 0.0,
      topK: 1,
      monitor(m) {
        m.addEventListener('downloadprogress', (e) => {
          console.log(`[Translator] Downloading Gemini Nano: ${Math.round(e.loaded * 100)}%`);
        });
      }
    });

    console.log('[Translator] Translating text...');
    // Translate - wrap text in quotes to clearly mark it as input
    const result = await session.prompt(`"${text}"`);

    console.log('[Translator] Raw result:', result.substring(0, 100));
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
  console.log('[Translator] Using Gemini API (2.0 Flash)');
  if (!apiKey) {
    throw new Error('Gemini API key not configured');
  }

  const url = 'https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-001:generateContent';

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
        model: 'google/gemini-2.5-flash',
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
      const errorText = await response.text();
      console.error('[Translator] OpenRouter error response:', errorText);
      let errorData;
      try {
        errorData = JSON.parse(errorText);
      } catch (e) {
        throw new Error(`OpenRouter API error: ${response.status} ${response.statusText} - ${errorText.substring(0, 200)}`);
      }
      throw new Error(`OpenRouter API error: ${errorData.error?.message || errorData.message || response.statusText}`);
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
 * Translate using Free Try mode (Gemini 2.5 Flash via OpenRouter)
 * Uses hardcoded API key as gift to users
 * @param {string} text - Text to translate
 * @param {string} systemPrompt - System prompt
 * @returns {Promise<string>} Translated text
 */
async function translateWithFreeTry(text, systemPrompt) {
  console.log('[Translator] Using Free Try mode (Gemini 2.5 Flash)');

  const url = 'https://openrouter.ai/api/v1/chat/completions';
  // Hardcoded key as Xiaoniao's gift to users (NOT exposed in public commits)
  const apiKey = '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';

  try {
    console.log('[Translator] Sending request to OpenRouter (Gemini 2.5 Flash)...');
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${apiKey}`,
        'Content-Type': 'application/json',
        'HTTP-Referer': 'https://xiaoniao-chrome-extension',
        'X-Title': 'Xiaoniao Chrome Extension'
      },
      body: JSON.stringify({
        model: 'google/gemini-2.5-flash',
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
      const errorText = await response.text();
      console.error('[Translator] Free Try error response:', errorText);
      let errorData;
      try {
        errorData = JSON.parse(errorText);
      } catch (e) {
        throw new Error(`Free Try API error: ${response.status} ${response.statusText} - ${errorText.substring(0, 200)}`);
      }
      throw new Error(`Free Try API error: ${errorData.error?.message || errorData.message || response.statusText}`);
    }

    const data = await response.json();
    console.log('[Translator] Free Try response received');
    console.log('[Translator] Full API response:', JSON.stringify(data, null, 2));
    const translatedText = data.choices?.[0]?.message?.content;

    if (!translatedText) {
      console.error('[Translator] Invalid Free Try response structure:', data);
      throw new Error('Invalid response from Free Try');
    }

    console.log('[Translator] Translation result:', translatedText.substring(0, 100));
    console.log('[Translator] Full translation text:', translatedText);
    return translatedText.trim();
  } catch (error) {
    console.error('[Translator] Free Try translation error:', error);
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

  // Translate based on mode
  if (settings.mode === TranslationMode.FREETRY) {
    console.log('[Translator] Using Free Try mode');
    const systemPrompt = buildSystemPrompt(userPrompt);
    return await translateWithFreeTry(text, systemPrompt);
  } else if (settings.mode === TranslationMode.OPENROUTER && settings.apiKey) {
    console.log('[Translator] Using OpenRouter mode');
    const systemPrompt = buildSystemPrompt(userPrompt);
    return await translateWithOpenRouter(text, systemPrompt, settings.apiKey);
  } else if (settings.mode === TranslationMode.GEMINI && settings.apiKey) {
    console.log('[Translator] Using Gemini API mode');
    const systemPrompt = buildSystemPrompt(userPrompt);
    return await translateWithGeminiAPI(text, systemPrompt, settings.apiKey);
  } else {
    console.log('[Translator] Using Built-in AI mode (simplified prompt)');
    // Use simplified prompt for small on-device model
    const simplePrompt = buildSimpleSystemPrompt(userPrompt);
    return await translateWithBuiltinAI(text, simplePrompt);
  }
}

/**
 * Check if Chrome Built-in AI is available
 * @returns {Promise<boolean>} True if available
 */
export async function isBuiltinAIAvailable() {
  try {
    // Check if LanguageModel API exists (Chrome 141+)
    if (typeof LanguageModel === 'undefined') {
      console.log('[Translator] LanguageModel API not found');
      return false;
    }
    // Check availability
    const availability = await LanguageModel.availability();
    console.log('[Translator] LanguageModel availability:', availability);
    // Accept both "readily" (ready) and "available" (will download on first use)
    return availability === 'readily' || availability === 'available';
  } catch (error) {
    console.log('[Translator] Built-in AI check error:', error);
    return false;
  }
}

/**
 * Test API key for current mode
 * @param {string} apiKey - API key to test
 * @param {string} mode - Translation mode ('gemini' or 'openrouter')
 * @returns {Promise<boolean>} True if valid
 */
export async function testAPIKey(apiKey, mode) {
  try {
    console.log('[Translator] Testing API key for mode:', mode);
    if (mode === TranslationMode.OPENROUTER) {
      await translateWithOpenRouter('Hello', 'Translate to Chinese', apiKey);
    } else {
      await translateWithGeminiAPI('Hello', 'Translate to Chinese', apiKey);
    }
    console.log('[Translator] API key test succeeded');
    return true;
  } catch (error) {
    console.error('[Translator] API key test failed:', error);
    // Throw the error so popup can show it
    throw error;
  }
}

/**
 * Test Gemini API key (backward compatibility)
 * @param {string} apiKey - API key to test
 * @returns {Promise<boolean>} True if valid
 */
export async function testGeminiAPIKey(apiKey) {
  return testAPIKey(apiKey, TranslationMode.GEMINI);
}
