// Popup UI logic
import { getAllPrompts, getCurrentPrompt, setActivePrompt, saveCustomPrompt, deleteCustomPrompt } from '../lib/prompts.js';
import { testGeminiAPIKey, isBuiltinAIAvailable } from '../lib/translator.js';

console.log('[Xiaoniao Popup] Loaded');

// DOM elements
const extensionToggle = document.getElementById('extensionToggle');
const modeBtns = document.querySelectorAll('.mode-btn');
const apiKeySection = document.getElementById('apiKeySection');
const apiKeyInput = document.getElementById('apiKeyInput');
const testApiKeyBtn = document.getElementById('testApiKey');
const promptList = document.getElementById('promptList');
const addCustomPromptBtn = document.getElementById('addCustomPrompt');
const statusDot = document.getElementById('statusDot');
const statusText = document.getElementById('statusText');

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
    const mode = settings.translationMode || 'builtin';
    modeBtns.forEach(btn => {
      if (btn.dataset.mode === mode) {
        btn.classList.add('active');
      } else {
        btn.classList.remove('active');
      }
    });

    // API key section visibility
    apiKeySection.style.display = mode === 'gemini' ? 'block' : 'none';
    if (settings.geminiApiKey) {
      apiKeyInput.value = settings.geminiApiKey;
    }

    // Load prompts
    await loadPrompts(settings.activePrompt || 'Auto Detect');

    // Update status
    updateStatus();

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

      // Delete button (only for custom prompts)
      const isDefault = ['English to Chinese', 'Chinese to English', 'Auto Detect', 'Casual', 'Formal'].includes(name);
      if (!isDefault) {
        const deleteBtn = document.createElement('button');
        deleteBtn.className = 'icon-btn';
        deleteBtn.innerHTML = '🗑️';
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
 * Update status indicator
 */
function updateStatus() {
  const enabled = extensionToggle.checked;

  if (!enabled) {
    statusDot.className = 'status-dot idle';
    statusText.textContent = 'Disabled';
  } else {
    statusDot.className = 'status-dot idle';
    statusText.textContent = 'Ready';
  }
}

/**
 * Event listeners
 */

// Extension toggle
extensionToggle.addEventListener('change', async () => {
  const enabled = extensionToggle.checked;
  await chrome.storage.sync.set({ extensionEnabled: enabled });
  updateStatus();
  console.log('[Xiaoniao Popup] Extension', enabled ? 'enabled' : 'disabled');
});

// Mode selection
modeBtns.forEach(btn => {
  btn.addEventListener('click', async () => {
    const mode = btn.dataset.mode;

    // Update UI
    modeBtns.forEach(b => b.classList.remove('active'));
    btn.classList.add('active');

    // Save to storage
    await chrome.storage.sync.set({ translationMode: mode });

    // Show/hide API key section
    apiKeySection.style.display = mode === 'gemini' ? 'block' : 'none';

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

  testApiKeyBtn.textContent = 'Testing...';
  testApiKeyBtn.disabled = true;

  try {
    const valid = await testGeminiAPIKey(apiKey);

    if (valid) {
      // Save API key
      await chrome.storage.sync.set({ geminiApiKey: apiKey });
      alert('✅ API key is valid and saved!');
    } else {
      alert('❌ API key is invalid. Please check and try again.');
    }
  } catch (error) {
    console.error('[Xiaoniao Popup] API key test error:', error);
    alert(`❌ Error: ${error.message}`);
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
addCustomPromptBtn.addEventListener('click', async () => {
  const name = prompt('Enter prompt name:');
  if (!name) return;

  const content = prompt('Enter prompt content (e.g., "Translate to Spanish in casual tone"):');
  if (!content) return;

  try {
    await saveCustomPrompt(name, content);
    await setActivePrompt(name);
    await loadSettings();
  } catch (error) {
    console.error('[Xiaoniao Popup] Error saving custom prompt:', error);
    alert('Error saving prompt');
  }
});

// Initialize
loadSettings();

// Check Built-in AI availability
(async () => {
  const available = await isBuiltinAIAvailable();
  if (!available) {
    console.warn('[Xiaoniao Popup] Chrome Built-in AI not available');
    // Could show a warning in the UI
  }
})();
