// Background service worker - Main translation logic
// Handles copy events, translates text, updates clipboard and icon

import { translate } from '../lib/translator.js';

console.log('[Xiaoniao] Background service worker started');

/**
 * Icon states
 */
const IconState = {
  IDLE: 'blue',      // Waiting for copy event
  TRANSLATING: 'red', // Translation in progress
  READY: 'green'     // Translation ready, user can paste
};

/**
 * Current state
 */
let currentState = IconState.IDLE;
let translationTimeout = null;

/**
 * Update extension icon
 * @param {string} state - Icon state (blue/red/green)
 * @param {number} tabId - Tab ID (optional, applies to all tabs if not specified)
 */
async function updateIcon(state, tabId) {
  currentState = state;
  const iconPath = `icons/icon_${state}.png`;

  try {
    const iconConfig = {
      path: {
        "16": iconPath,
        "48": iconPath,
        "128": iconPath
      }
    };

    if (tabId) {
      iconConfig.tabId = tabId;
    }

    await chrome.action.setIcon(iconConfig);
    console.log(`[Xiaoniao] Icon updated to ${state}`);
  } catch (error) {
    // Icon errors are non-critical, just log quietly
    console.log(`[Xiaoniao] Icon update skipped (${state})`);
  }
}

/**
 * Write text to clipboard
 * @param {string} text - Text to write
 * @param {number} tabId - Tab ID to execute in
 */
async function writeToClipboard(text, tabId) {
  try {
    // In Manifest V3, service workers don't have navigator.clipboard
    // We need to inject script into the page to write to clipboard
    await chrome.scripting.executeScript({
      target: { tabId: tabId },
      func: (textToWrite) => {
        navigator.clipboard.writeText(textToWrite);
      },
      args: [text]
    });
    console.log('[Xiaoniao] Clipboard updated with translation');
    return true;
  } catch (error) {
    console.error('[Xiaoniao] Error writing to clipboard:', error);
    return false;
  }
}

/**
 * Update translation statistics
 */
async function updateStatistics() {
  try {
    const stats = await chrome.storage.sync.get(['translationCount', 'firstUseDate']);

    // Increment translation count
    const newCount = (stats.translationCount || 0) + 1;

    // Set first use date if not set
    const firstUseDate = stats.firstUseDate || Date.now();

    await chrome.storage.sync.set({
      translationCount: newCount,
      firstUseDate: firstUseDate
    });

    console.log(`[Xiaoniao] Statistics updated: ${newCount} translations`);
  } catch (error) {
    console.error('[Xiaoniao] Error updating statistics:', error);
  }
}

/**
 * Auto-insert translated text into active element
 * @param {string} text - Text to insert
 * @param {number} tabId - Tab ID to insert into
 */
async function autoInsertText(text, tabId) {
  try {
    console.log('[Xiaoniao] Sending AUTO_INSERT message to tab:', tabId);

    // Send message to content script to insert text
    const response = await chrome.tabs.sendMessage(tabId, {
      type: 'AUTO_INSERT',
      text: text
    });

    console.log('[Xiaoniao] AUTO_INSERT response:', response);

    if (response && response.success) {
      console.log('[Xiaoniao] ✅ Text auto-inserted successfully');
      return true;
    } else {
      console.log('[Xiaoniao] ❌ Auto-insert failed, text remains in clipboard');
      return false;
    }
  } catch (error) {
    console.error('[Xiaoniao] ❌ Error auto-inserting text:', error);
    console.error('[Xiaoniao] Error details:', error.message);
    return false;
  }
}

/**
 * Handle copy event from content script
 * @param {string} text - Copied text
 * @param {number} tabId - Tab ID where copy occurred
 */
async function handleCopyEvent(text, tabId) {
  try {
    // Check if extension is enabled
    const settings = await chrome.storage.sync.get([
      'extensionEnabled',
      'autoPasteEnabled'
    ]);

    if (settings.extensionEnabled === false) {
      console.log('[Xiaoniao] Extension is disabled');
      return;
    }

    const autoPasteEnabled = settings.autoPasteEnabled !== false; // Default: true

    console.log('[Xiaoniao] Starting translation for:', text.substring(0, 50) + '...');
    console.log('[Xiaoniao] Auto-paste enabled:', autoPasteEnabled);

    // Set icon to TRANSLATING (red)
    await updateIcon(IconState.TRANSLATING, tabId);

    // Clear any existing timeout
    if (translationTimeout) {
      clearTimeout(translationTimeout);
    }

    // Translate
    const startTime = Date.now();
    const translatedText = await translate(text);
    const duration = Date.now() - startTime;

    console.log(`[Xiaoniao] Translation completed in ${duration}ms`);
    console.log('[Xiaoniao] Result:', translatedText.substring(0, 50) + '...');

    // Update statistics
    await updateStatistics();

    // Always write to clipboard first (fallback)
    await writeToClipboard(translatedText, tabId);

    // Auto-insert if enabled
    if (autoPasteEnabled) {
      console.log('[Xiaoniao] Attempting auto-insert...');
      await autoInsertText(translatedText, tabId);
    } else {
      console.log('[Xiaoniao] Auto-paste disabled, user needs to press Ctrl+V');
    }

    // Set icon to READY (green)
    await updateIcon(IconState.READY, tabId);

    // Reset to IDLE after 3 seconds
    translationTimeout = setTimeout(async () => {
      await updateIcon(IconState.IDLE, tabId);
    }, 3000);

  } catch (error) {
    console.error('[Xiaoniao] Translation error:', error);

    // Show error by keeping red icon for 2 seconds, then back to idle
    setTimeout(async () => {
      await updateIcon(IconState.IDLE, tabId);
    }, 2000);
  }
}

/**
 * Listen for messages from content scripts
 */
chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  console.log('[Xiaoniao] Message received:', message.type, 'from tab:', sender.tab?.id);

  if (message.type === 'COPY_EVENT') {
    const tabId = sender.tab?.id;

    // Handle async
    handleCopyEvent(message.text, tabId)
      .then(() => {
        sendResponse({ success: true });
      })
      .catch((error) => {
        console.error('[Xiaoniao] Error in handleCopyEvent:', error);
        sendResponse({ success: false, error: error.message });
      });

    // Return true to indicate async response
    return true;
  }
});

/**
 * Initialize on install/update
 */
chrome.runtime.onInstalled.addListener(async (details) => {
  console.log('[Xiaoniao] Extension installed/updated:', details.reason);

  // Set default settings
  const defaults = {
    extensionEnabled: true,
    translationMode: 'gemini',
    activePrompt: 'CN_EN',
    autoPasteEnabled: true
  };

  // Only set if not already set
  const existing = await chrome.storage.sync.get(Object.keys(defaults));
  const toSet = {};

  for (const [key, value] of Object.entries(defaults)) {
    if (existing[key] === undefined) {
      toSet[key] = value;
    }
  }

  if (Object.keys(toSet).length > 0) {
    await chrome.storage.sync.set(toSet);
    console.log('[Xiaoniao] Default settings initialized:', toSet);
  }

  // Don't set icon during install - can cause errors
  console.log('[Xiaoniao] ✅ Initialization complete');

  // NOTE: Content scripts are registered via manifest.json
  // No need for programmatic registration
});

/**
 * Handle extension startup
 */
chrome.runtime.onStartup.addListener(async () => {
  console.log('[Xiaoniao] Extension started');
});
