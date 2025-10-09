// Content script - Intercepts copy events on all web pages
// Sends copied text to background for translation

console.log('[Xiaoniao Content] Script loaded on:', window.location.href);

/**
 * Handle copy event
 */
document.addEventListener('copy', async (event) => {
  try {
    console.log('[Xiaoniao Content] ===== COPY EVENT FIRED =====');

    // Get selected text
    const selectedText = window.getSelection().toString().trim();
    console.log('[Xiaoniao Content] Selection length:', selectedText.length);

    if (!selectedText || selectedText.length === 0) {
      console.log('[Xiaoniao Content] No text selected, ignoring');
      return; // Nothing selected, let default copy happen
    }

    console.log('[Xiaoniao Content] Text copied:', selectedText.substring(0, 100) + (selectedText.length > 100 ? '...' : ''));

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
 * (for future features like showing translation status)
 */
chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message.type === 'TRANSLATION_STATUS') {
    console.log('[Xiaoniao] Translation status:', message.status);
    // Could show visual indicator here in the future
  }

  sendResponse({ received: true });
});
