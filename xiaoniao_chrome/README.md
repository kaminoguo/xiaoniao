# Xiaoniao for Chrome

**Type Translation Assistant for Real-Time Chat**

Xiaoniao automatically translates your clipboard content while you're chatting on Discord, Instagram, WhatsApp Web, or any website. Just copy text, wait a moment, and paste the translation.

---

## ✨ Features

- **🚀 Automatic Translation**: Copy text → Translation happens in background → Paste translated text
- **⚡ Fast & Private**: Uses Chrome's built-in AI (Gemini Nano) by default - no data leaves your browser
- **🎨 Visual Status**: Extension icon changes color to show translation status:
  - **Blue** - Ready (waiting for copy)
  - **Red** - Translating... (please wait)
  - **Green** - Translation ready! (paste now)
- **🎯 Hybrid AI**: Choose between Chrome Built-in AI (fast, private) or Gemini API (higher quality)
- **💬 Multiple Styles**: Auto-detect, casual, formal, or create custom translation styles
- **🌐 Works Everywhere**: Discord, Instagram, WhatsApp Web, Twitter, and all websites

---

## 🎬 How It Works

1. **Copy** text you want to translate (Ctrl+C / Cmd+C)
2. **Wait** 1-2 seconds while icon turns red → green
3. **Paste** (Ctrl+V / Cmd+V) and get the translation!

That's it! No buttons to click, no menus to open.

---

## 📦 Installation

### Method 1: Load Unpacked (for hackathon/testing)

1. Open Chrome and go to `chrome://extensions/`
2. Enable **Developer mode** (toggle in top right)
3. Click **Load unpacked**
4. Select the `xiaoniao_chrome` folder
5. Done! The extension is now active

### Method 2: Chrome Web Store (coming soon)

Will be available after hackathon review.

---

## ⚙️ Setup

### Default Mode: Chrome Built-in AI (Recommended)

1. **Enable Chrome AI** (one-time setup):
   - Go to `chrome://flags/#optimization-guide-on-device-model`
   - Set to **Enabled**
   - Go to `chrome://flags/#prompt-api-for-gemini-nano`
   - Set to **Enabled**
   - Restart Chrome

2. **That's it!** The extension now uses local AI (Gemini Nano)

### Optional: Use Gemini API (Higher Quality)

1. Click extension icon
2. Switch to **Gemini API** mode
3. Get a free API key from [ai.google.dev](https://ai.google.dev)
4. Enter your API key and click **Test**
5. Done!

---

## 🎯 Translation Styles

### Built-in Styles

- **Auto Detect**: Automatically detects language and translates (Chinese ↔ English)
- **English to Chinese**: Translate any text to Chinese
- **Chinese to English**: Translate any text to English
- **Casual**: Friendly, informal tone
- **Formal**: Professional, business tone

### Custom Styles

Click **Add Custom Style** in the popup to create your own:

**Examples**:
- "Translate to Spanish in a friendly tone"
- "Translate to Japanese, keep emoji and formatting"
- "Translate technical documentation to French"

---

## 🎨 Icon Status

| Color | Meaning |
|-------|---------|
| 🔵 **Blue** | Ready - extension is waiting for you to copy text |
| 🔴 **Red** | Translating - AI is working, please wait a moment |
| 🟢 **Green** | Ready to paste - translation is in your clipboard! |

---

## ⚡ Performance

- **Chrome Built-in AI**: < 1 second translation
- **Gemini API**: 1-2 seconds translation
- **Memory usage**: < 50MB
- **Privacy**: Built-in mode keeps all data local (nothing sent to servers)

---

## 🏆 For Chrome Built-in AI Challenge 2025

This extension is built for the [Google Chrome Built-in AI Challenge 2025](https://googlechromeai2025.devpost.com/).

### Key Technologies

- **Chrome Built-in AI APIs**:
  - `window.ai.languageModel.create()` - Gemini Nano for translation
  - `navigator.clipboard` API - Clipboard read/write
- **Hybrid AI Strategy**:
  - Default: On-device Gemini Nano (fast, private, offline-capable)
  - Optional: Cloud Gemini 2.5 Flash API (higher quality)
- **Content Scripts**: Intercept copy events on all websites
- **Service Worker**: Background translation processing

---

## 🛠️ Development

### Project Structure

```
xiaoniao_chrome/
├── manifest.json           # Extension configuration
├── popup/
│   ├── popup.html         # Settings UI
│   ├── popup.css          # Gemini-style design
│   └── popup.js           # UI logic
├── content/
│   └── content.js         # Copy event interceptor
├── background/
│   └── service-worker.js  # Translation engine & icon updates
├── lib/
│   ├── prompts.js         # Translation prompt templates
│   └── translator.js      # Hybrid AI translation logic
└── icons/
    ├── icon_blue.png      # Idle state
    ├── icon_red.png       # Translating state
    └── icon_green.png     # Ready state
```

### Building

No build process required! Load directly in Chrome for development.

### Testing

1. Load extension in Chrome
2. Open any website (e.g., Discord web)
3. Select and copy text
4. Watch icon change: Blue → Red → Green
5. Paste to get translation

---

## 📝 License

MIT License - see [LICENSE](../LICENSE) file

---

## 🤝 Contributing

This is part of the Xiaoniao project family:
- **Windows version**: Full-featured desktop app
- **Chrome version**: Browser extension (this)
- **Android version**: Mobile app (coming soon)

---

## 💬 Feedback

Found a bug? Have a suggestion? Please [open an issue](https://github.com/kaminoguo/xiaoniao/issues)!

---

**Made with ❤️ for seamless multilingual communication**
