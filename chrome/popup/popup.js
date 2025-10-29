// Popup UI logic
import { getAllPrompts, getCurrentPrompt, setActivePrompt, saveCustomPrompt, deleteCustomPrompt } from '../lib/prompts.js';
import { testAPIKey, testGeminiAPIKey, isBuiltinAIAvailable } from '../lib/translator.js';

console.log('[Xiaoniao Popup] Loaded');

// DOM elements
const extensionToggle = document.getElementById('extensionToggle');
const modeBtns = document.querySelectorAll('.mode-btn');
const apiKeySection = document.getElementById('apiKeySection');
const apiKeyInput = document.getElementById('apiKeyInput');
const testApiKeyBtn = document.getElementById('testApiKey');
const promptList = document.getElementById('promptList');
const addCustomPromptBtn = document.getElementById('addCustomPrompt');

// Modal elements
const promptModal = document.getElementById('promptModal');
const modalTitle = document.getElementById('modalTitle');
const promptNameInput = document.getElementById('promptNameInput');
const promptContentInput = document.getElementById('promptContentInput');
const modalClose = document.getElementById('modalClose');
const modalCancel = document.getElementById('modalCancel');
const modalSave = document.getElementById('modalSave');

let currentEditingPrompt = null; // Track which prompt is being edited

// Unlock mechanism for Gift mode
const GIFT_UNLOCK_KEY = 'xiaoniaoGiftUnlocked';

/**
 * Check if Gift mode is unlocked
 */
function isGiftUnlocked() {
  return localStorage.getItem(GIFT_UNLOCK_KEY) === 'true';
}

/**
 * Unlock Gift mode
 */
function unlockGift() {
  localStorage.setItem(GIFT_UNLOCK_KEY, 'true');
  console.log('[Xiaoniao Popup] Gift mode unlocked');

  // Remove locked class from Gift button
  const giftBtn = document.querySelector('[data-mode="freetry"]');
  if (giftBtn) {
    giftBtn.classList.remove('locked');
  }
}

/**
 * Shake footer icons to hint unlock
 */
function shakeFooterIcons() {
  const footerLinks = document.querySelectorAll('.footer-link');
  footerLinks.forEach(link => {
    link.classList.add('shake');
    setTimeout(() => {
      link.classList.remove('shake');
    }, 500);
  });
}

/**
 * Load and display settings
 */
async function loadSettings() {
  try {
    const settings = await chrome.storage.sync.get([
      'extensionEnabled',
      'translationMode',
      'geminiApiKey',
      'activePrompt'
    ]);

    // Extension toggle
    extensionToggle.checked = settings.extensionEnabled !== false;

    // Translation mode
    const mode = settings.translationMode || 'gemini';
    modeBtns.forEach(btn => {
      if (btn.dataset.mode === mode) {
        btn.classList.add('active');
      } else {
        btn.classList.remove('active');
      }

      // Check if Gift mode is locked
      if (btn.dataset.mode === 'freetry' && !isGiftUnlocked()) {
        btn.classList.add('locked');
      }
    });

    // API key section visibility and labels (only update when expanded)
    const isExpanded = !document.getElementById('translationEngineHeader').classList.contains('collapsed');

    if (mode === 'openrouter') {
      document.getElementById('apiKeyTitle').textContent = 'OpenRouter API Key';
      document.getElementById('apiKeyHint').innerHTML = 'Get your free API key from <a href="https://openrouter.ai/keys" target="_blank" id="apiKeyLink">openrouter.ai/keys</a>';
      if (isExpanded) {
        apiKeySection.style.display = 'flex';
        apiKeySection.classList.remove('hidden');
      }
    } else if (mode === 'gemini') {
      document.getElementById('apiKeyTitle').textContent = 'Gemini API Key';
      document.getElementById('apiKeyHint').innerHTML = 'Get your free API key from <a href="https://ai.google.dev" target="_blank" id="apiKeyLink">ai.google.dev</a>';
      if (isExpanded) {
        apiKeySection.style.display = 'flex';
        apiKeySection.classList.remove('hidden');
      }
    } else {
      apiKeySection.classList.add('hidden');
      setTimeout(() => {
        apiKeySection.style.display = 'none';
      }, 300);
    }

    if (settings.geminiApiKey) {
      apiKeyInput.value = settings.geminiApiKey;
    }

    // Load prompts
    await loadPrompts(settings.activePrompt || 'CN_EN');

    // Load statistics
    await loadStatistics();

  } catch (error) {
    console.error('[Xiaoniao Popup] Error loading settings:', error);
  }
}

/**
 * Load and display prompts
 */
async function loadPrompts(activePromptName) {
  try {
    const prompts = await getAllPrompts();
    promptList.innerHTML = '';

    for (const [name, content] of Object.entries(prompts)) {
      const item = document.createElement('div');
      item.className = `prompt-item ${name === activePromptName ? 'active' : ''}`;

      const nameSpan = document.createElement('span');
      nameSpan.className = 'prompt-name';
      nameSpan.textContent = name;

      const actions = document.createElement('div');
      actions.className = 'prompt-actions';

      // Edit button (for all prompts)
      const editBtn = document.createElement('button');
      editBtn.className = 'icon-btn';
      editBtn.innerHTML = '<svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M3 17.25V21h3.75L17.81 9.94l-3.75-3.75L3 17.25zM20.71 7.04c.39-.39.39-1.02 0-1.41l-2.34-2.34c-.39-.39-1.02-.39-1.41 0l-1.83 1.83 3.75 3.75 1.83-1.83z"/></svg>';
      editBtn.title = 'Edit';
      editBtn.onclick = async (e) => {
        e.stopPropagation();
        openEditModal(name, content);
      };
      actions.appendChild(editBtn);

      // Delete button (only for custom prompts)
      const isDefault = ['CN_EN'].includes(name);
      if (!isDefault) {
        const deleteBtn = document.createElement('button');
        deleteBtn.className = 'icon-btn';
        deleteBtn.innerHTML = '<svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"/></svg>';
        deleteBtn.title = 'Delete';
        deleteBtn.onclick = async (e) => {
          e.stopPropagation();
          if (confirm(`Delete "${name}"?`)) {
            await deleteCustomPrompt(name);
            await loadSettings();
          }
        };
        actions.appendChild(deleteBtn);
      }

      item.appendChild(nameSpan);
      item.appendChild(actions);

      // Click to activate
      item.onclick = async () => {
        await setActivePrompt(name);
        await loadSettings();
      };

      promptList.appendChild(item);
    }
  } catch (error) {
    console.error('[Xiaoniao Popup] Error loading prompts:', error);
  }
}

/**
 * Load and display statistics
 */
async function loadStatistics() {
  try {
    const stats = await chrome.storage.sync.get(['translationCount', 'firstUseDate']);

    // Update translation count
    const count = stats.translationCount || 0;
    document.getElementById('translationCount').textContent = count.toLocaleString();

    // Calculate usage days
    const firstUse = stats.firstUseDate || Date.now();
    const daysSinceFirstUse = Math.floor((Date.now() - firstUse) / (1000 * 60 * 60 * 24)) + 1;
    document.getElementById('usageDays').textContent = daysSinceFirstUse.toLocaleString();

  } catch (error) {
    console.error('[Xiaoniao Popup] Error loading statistics:', error);
  }
}

/**
 * Event listeners
 */

// Extension toggle
extensionToggle.addEventListener('change', async () => {
  const enabled = extensionToggle.checked;
  await chrome.storage.sync.set({ extensionEnabled: enabled });
  console.log('[Xiaoniao Popup] Extension', enabled ? 'enabled' : 'disabled');
});

// Mode selection
modeBtns.forEach(btn => {
  btn.addEventListener('click', async () => {
    const mode = btn.dataset.mode;

    // Check if Gift mode is locked
    if (mode === 'freetry' && !isGiftUnlocked()) {
      shakeFooterIcons();
      console.log('[Xiaoniao Popup] Gift mode is locked, please click footer icons to unlock');
      return;
    }

    // Update UI
    modeBtns.forEach(b => b.classList.remove('active'));
    btn.classList.add('active');

    // Save to storage
    await chrome.storage.sync.set({ translationMode: mode });

    // Show/hide and update API key section (only when expanded)
    const isExpanded = !document.getElementById('translationEngineHeader').classList.contains('collapsed');

    if (mode === 'openrouter') {
      document.getElementById('apiKeyTitle').textContent = 'OpenRouter API Key';
      document.getElementById('apiKeyHint').innerHTML = 'Get your free API key from <a href="https://openrouter.ai/keys" target="_blank">openrouter.ai/keys</a>';
      if (isExpanded) {
        apiKeySection.style.display = 'flex';
        apiKeySection.classList.remove('hidden');
      }
    } else if (mode === 'gemini') {
      document.getElementById('apiKeyTitle').textContent = 'Gemini API Key';
      document.getElementById('apiKeyHint').innerHTML = 'Get your free API key from <a href="https://ai.google.dev" target="_blank">ai.google.dev</a>';
      if (isExpanded) {
        apiKeySection.style.display = 'flex';
        apiKeySection.classList.remove('hidden');
      }
    } else {
      apiKeySection.classList.add('hidden');
      setTimeout(() => {
        apiKeySection.style.display = 'none';
      }, 300);
    }

    console.log('[Xiaoniao Popup] Mode changed to', mode);
  });
});

// Test API key
testApiKeyBtn.addEventListener('click', async () => {
  const apiKey = apiKeyInput.value.trim();

  if (!apiKey) {
    alert('Please enter an API key');
    return;
  }

  // Get current mode
  const settings = await chrome.storage.sync.get(['translationMode']);
  const mode = settings.translationMode || 'builtin';

  testApiKeyBtn.textContent = 'Testing...';
  testApiKeyBtn.disabled = true;

  try {
    await testAPIKey(apiKey, mode);

    // If we get here, test succeeded
    await chrome.storage.sync.set({ geminiApiKey: apiKey });
    alert('✅ API key is valid and saved!');
  } catch (error) {
    console.error('[Xiaoniao Popup] API key test error:', error);
    alert(`❌ API key test failed:\n\n${error.message}`);
  } finally {
    testApiKeyBtn.textContent = 'Test';
    testApiKeyBtn.disabled = false;
  }
});

// Save API key on blur
apiKeyInput.addEventListener('blur', async () => {
  const apiKey = apiKeyInput.value.trim();
  if (apiKey) {
    await chrome.storage.sync.set({ geminiApiKey: apiKey });
  }
});

// Add custom prompt
addCustomPromptBtn.addEventListener('click', () => {
  openEditModal(null, null);
});

// Collapsible Translation Engine section
const translationEngineHeader = document.getElementById('translationEngineHeader');
translationEngineHeader.addEventListener('click', () => {
  const modeSelector = document.querySelector('.mode-selector');
  const apiKeySection = document.getElementById('apiKeySection');
  const isCollapsed = translationEngineHeader.classList.contains('collapsed');

  if (isCollapsed) {
    // Expand
    translationEngineHeader.classList.remove('collapsed');
    modeSelector.classList.remove('hidden');
    // Show API key section if needed
    const settings = chrome.storage.sync.get(['translationMode'], (result) => {
      const mode = result.translationMode || 'gemini';
      if (mode === 'openrouter' || mode === 'gemini') {
        apiKeySection.style.display = 'flex';
        setTimeout(() => {
          apiKeySection.classList.remove('hidden');
        }, 10);
      }
    });
  } else {
    // Collapse
    translationEngineHeader.classList.add('collapsed');
    modeSelector.classList.add('hidden');
    apiKeySection.classList.add('hidden');
    setTimeout(() => {
      apiKeySection.style.display = 'none';
    }, 300);
  }
});

// Initialize
loadSettings();

// Start with collapsed state
translationEngineHeader.classList.add('collapsed');
document.querySelector('.mode-selector').classList.add('hidden');

// Check Built-in AI availability
(async () => {
  const available = await isBuiltinAIAvailable();
  const warningSection = document.getElementById('builtinAIWarning');

  if (!available) {
    console.warn('[Xiaoniao Popup] Chrome Built-in AI not available');
    // Show warning in UI
    warningSection.style.display = 'block';
  } else {
    warningSection.style.display = 'none';
  }
})();

// Footer links unlock Gift mode
document.querySelectorAll('.footer-link').forEach(link => {
  link.addEventListener('click', () => {
    if (!isGiftUnlocked()) {
      unlockGift();
      console.log('[Xiaoniao Popup] Gift mode unlocked via footer link');
    }
  });
});

/**
 * Modal Functions
 */

// Open modal for editing or creating prompt
function openEditModal(name, content) {
  currentEditingPrompt = name;

  if (name) {
    // Editing existing prompt
    modalTitle.textContent = 'Edit Translation Style';
    promptNameInput.value = name;
    promptContentInput.value = content;

    // Disable name input for default prompts
    const isDefault = ['CN_EN'].includes(name);
    promptNameInput.disabled = isDefault;
  } else {
    // Creating new prompt
    modalTitle.textContent = 'Add Translation Style';
    promptNameInput.value = '';
    promptContentInput.value = '';
    promptNameInput.disabled = false;
  }

  promptModal.style.display = 'flex';
}

// Close modal
function closeModal() {
  promptModal.style.display = 'none';
  currentEditingPrompt = null;
  promptNameInput.value = '';
  promptContentInput.value = '';
}

// Modal close handlers
modalClose.addEventListener('click', closeModal);
modalCancel.addEventListener('click', closeModal);

// Click outside to close
promptModal.addEventListener('click', (e) => {
  if (e.target === promptModal) {
    closeModal();
  }
});

// Save prompt
modalSave.addEventListener('click', async () => {
  const name = promptNameInput.value.trim();
  const content = promptContentInput.value.trim();

  if (!name) {
    alert('Please enter a name');
    return;
  }

  if (!content) {
    alert('Please enter an instruction');
    return;
  }

  try {
    // If editing an existing custom prompt and name changed, delete old one
    if (currentEditingPrompt && currentEditingPrompt !== name) {
      const isDefault = ['CN_EN'].includes(currentEditingPrompt);
      if (!isDefault) {
        await deleteCustomPrompt(currentEditingPrompt);
      }
    }

    // Save the prompt
    await saveCustomPrompt(name, content);
    await setActivePrompt(name);
    await loadSettings();
    closeModal();
  } catch (error) {
    console.error('[Xiaoniao Popup] Error saving prompt:', error);
    alert('Error saving prompt');
  }
});
