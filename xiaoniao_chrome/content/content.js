// Content script - Intercepts copy events on all web pages
// Sends copied text to background for translation
// Auto-inserts translated text into active input element

console.log('[Xiaoniao] Content script loaded');

// Store the active element when copy happens
let lastActiveElement = null;

/**
 * Insert text into the currently focused element
 * Supports: input, textarea, contenteditable
 * @param {string} text - Text to insert
 * @returns {Promise<boolean>} - Success status
 */
async function insertTextIntoActiveElement(text) {
  const element = lastActiveElement || document.activeElement;

  console.log('[Xiaoniao] 🎯 Attempting to insert text');
  console.log('[Xiaoniao] Target element:', element?.tagName, element?.type, element?.className);

  if (!element) {
    console.error('[Xiaoniao] ❌ No active element found');
    return false;
  }

  try {
    // Handle regular input and textarea elements
    if (element.tagName === 'INPUT' || element.tagName === 'TEXTAREA') {
      console.log('[Xiaoniao] 📝 Processing INPUT/TEXTAREA element');
      const start = element.selectionStart || 0;
      const end = element.selectionEnd || 0;

      // Use setRangeText for modern browsers (better than direct value assignment)
      if (element.setRangeText) {
        element.setRangeText(text, start, end, 'end');
      } else {
        // Fallback for older browsers
        const value = element.value;
        element.value = value.substring(0, start) + text + value.substring(end);
        element.selectionStart = element.selectionEnd = start + text.length;
      }

      // Trigger events to notify frameworks (React, Vue, etc.)
      element.dispatchEvent(new KeyboardEvent('keydown', { bubbles: true }));
      element.dispatchEvent(new Event('input', { bubbles: true }));
      element.dispatchEvent(new KeyboardEvent('keyup', { bubbles: true }));
      element.dispatchEvent(new Event('change', { bubbles: true }));

      // Ensure focus
      element.focus();

      console.log('[Xiaoniao] Text inserted into', element.tagName);
      return true;
    }

    // Handle contenteditable elements (rich text editors like Discord/Slack)
    if (element.isContentEditable) {
      console.log('[Xiaoniao] 📝 Processing contenteditable element');
      element.focus();

      // Try Method 1: beforeinput + input events (preferred for Slate.js)
      try {
        console.log('[Xiaoniao] Trying beforeinput event...');

        const beforeInputEvent = new InputEvent('beforeinput', {
          inputType: 'insertFromPaste',
          data: text,
          bubbles: true,
          cancelable: true
        });

        const inputEvent = new InputEvent('input', {
          inputType: 'insertFromPaste',
          data: text,
          bubbles: true
        });

        element.dispatchEvent(beforeInputEvent);
        element.dispatchEvent(inputEvent);

        // Wait 100ms and check if insertion succeeded
        await new Promise(resolve => setTimeout(resolve, 100));

        if (element.textContent.includes(text)) {
          console.log('[Xiaoniao] ✅ beforeinput method succeeded');
          return true;
        }

        console.log('[Xiaoniao] beforeinput did not work, trying IME flow...');
      } catch (e) {
        console.warn('[Xiaoniao] beforeinput failed:', e);
      }

      // Try Method 2: IME composition events (fallback)
      try {
        console.log('[Xiaoniao] Trying IME composition...');

        element.dispatchEvent(new CompositionEvent('compositionstart', { bubbles: true }));
        element.dispatchEvent(new CompositionEvent('compositionupdate', { data: text, bubbles: true }));

        const beforeInputEvent = new InputEvent('beforeinput', {
          inputType: 'insertCompositionText',
          data: text,
          bubbles: true,
          cancelable: true
        });
        element.dispatchEvent(beforeInputEvent);

        element.dispatchEvent(new CompositionEvent('compositionend', { data: text, bubbles: true }));

        // Wait 100ms and check
        await new Promise(resolve => setTimeout(resolve, 100));

        if (element.textContent.includes(text)) {
          console.log('[Xiaoniao] ✅ IME composition succeeded');
          return true;
        }

        console.log('[Xiaoniao] IME composition did not work');
      } catch (e) {
        console.warn('[Xiaoniao] IME composition failed:', e);
      }

      // Both methods failed - notify user to paste manually
      console.log('[Xiaoniao] ❌ Auto-paste failed for contenteditable');
      console.log('[Xiaoniao] 💡 Translation is in clipboard - Press Ctrl+V to paste');
      return false;
    }

    console.log('[Xiaoniao] Element is not editable:', element.tagName);
    return false;

  } catch (error) {
    console.error('[Xiaoniao] Error inserting text:', error);
    return false;
  }
}

/**
 * Handle copy event
 */
document.addEventListener('copy', async (event) => {
  try {
    // Get selected text
    const selectedText = window.getSelection().toString().trim();

    if (!selectedText || selectedText.length === 0) {
      return; // Nothing selected, let default copy happen
    }

    // Store the currently focused element for later auto-paste
    lastActiveElement = document.activeElement;

    console.log('[Xiaoniao] Copy detected:', selectedText.substring(0, 50) + '...');
    console.log('[Xiaoniao] Active element:', lastActiveElement?.tagName);

    // Send to background for translation
    chrome.runtime.sendMessage({
      type: 'COPY_EVENT',
      text: selectedText
    }, (response) => {
      if (chrome.runtime.lastError) {
        console.error('[Xiaoniao] Error sending message:', chrome.runtime.lastError);
        return;
      }

      if (response && response.success) {
        console.log('[Xiaoniao] Translation started in background');
      }
    });

    // Let the default copy happen (copy original text to clipboard)
    // Background will replace it with translation when ready
  } catch (error) {
    console.error('[Xiaoniao] Error in copy handler:', error);
  }
});

/**
 * Listen for messages from background
 */
chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  console.log('[Xiaoniao] Message received:', message.type);

  if (message.type === 'TRANSLATION_STATUS') {
    console.log('[Xiaoniao] Translation status:', message.status);
    sendResponse({ received: true });
    return false;
  }

  // Auto-insert translated text
  if (message.type === 'AUTO_INSERT') {
    console.log('[Xiaoniao] 🔥 Auto-insert request received');
    console.log('[Xiaoniao] Text to insert:', message.text?.substring(0, 50) + '...');
    console.log('[Xiaoniao] Last active element:', lastActiveElement?.tagName);
    console.log('[Xiaoniao] Current active element:', document.activeElement?.tagName);

    if (!message.text) {
      console.error('[Xiaoniao] ❌ No text provided for auto-insert');
      sendResponse({ success: false, error: 'No text provided' });
      return false;
    }

    // Handle async function
    insertTextIntoActiveElement(message.text).then(success => {
      console.log('[Xiaoniao] Auto-insert result:', success ? '✅ Success' : '❌ Failed');
      sendResponse({ success });
    });

    return true; // Async response
  }

  sendResponse({ received: true });
  return false;
});
