# xiaoniao v1.0

Windows clipboard translation tool. Monitor clipboard, auto-translate, auto-paste.

🎉 **v1.0 Major Update**
- ✅ Complete internationalization support (7 languages)
- ✅ Fixed garbled text issues on non-Chinese Windows systems
- ✅ Supports English, Simplified Chinese, Traditional Chinese, Japanese, Korean, Spanish, French
- ✅ All UI text and prompts fully translated

## Installation

Download [xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) and run.

Windows SmartScreen warning: Click "More info" → "Run anyway"

First run requires API key configuration.

## Usage

1. Run xiaoniao.exe
2. Copy text (Ctrl+C)
3. Auto-translate and replace clipboard
4. Paste (Ctrl+V) to get translation

System tray icon:
- Blue: Monitoring
- Green: Translating
- Red: Stopped

## Configuration

```cmd
xiaoniao.exe config
```

Config file: `%APPDATA%\xiaoniao\`

Supports OpenAI, Anthropic, Google, DeepSeek and other APIs.

## Build

```bash
# Basic build
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go

# With icon
go generate
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
```

Requirements: Go 1.21+, Windows

## License

MIT License