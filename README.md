# xiaoniao

AI-powered clipboard translation tool with support for 20+ language models.

## Overview

xiaoniao is a cross-platform clipboard monitor that automatically translates copied text using various AI providers. It runs in the background, monitoring clipboard changes and replacing content with translations.

## Features

- Real-time clipboard monitoring with automatic translation
- Support for 20+ AI providers including OpenAI, Anthropic, Google, and OpenRouter
- Customizable translation prompts
- Terminal-based configuration interface
- System tray integration
- Global hotkey support
- Multi-language interface (Chinese, English, Japanese, Korean, Spanish, French)

## Supported Platforms

- Linux (X11/Wayland)
- Windows 10/11
- macOS 10.15+

## Installation

### Linux

```bash
curl -sSL https://github.com/kaminoguo/xiaoniao/releases/latest/download/linux-install.sh | bash
```

### Windows

1. Download [xiaoniao-windows.zip](https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-windows.zip)
2. Extract to desired location
3. Run xiaoniao.exe

### macOS

Download the appropriate version:
- Intel: [xiaoniao-darwin-amd64.zip](https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-darwin-amd64.zip)
- Apple Silicon: [xiaoniao-darwin-arm64.zip](https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-darwin-arm64.zip)

Extract and run:
```bash
chmod +x xiaoniao
./xiaoniao config
```

## Configuration

### Initial Setup

```bash
xiaoniao config
```

The configuration interface allows you to:
- Set API keys for your chosen provider
- Select translation model
- Choose or create custom translation prompts
- Configure interface language and theme
- Set up global hotkeys

### Configuration Files

Configuration files are stored in platform-specific locations:
- Linux: `~/.config/xiaoniao/`
- Windows: `%APPDATA%\xiaoniao\`
- macOS: `~/Library/Application Support/xiaoniao/`

Two main configuration files:
- `config.json`: Main application settings
- `prompts.json`: Custom translation prompts

## Usage

### Start Monitoring

```bash
xiaoniao run
```

Once running:
1. Copy any text to clipboard
2. Wait for translation (1-3 seconds)
3. Paste to receive translated text

### Tray Menu Options

- Toggle monitoring on/off
- Switch translation prompts
- Open configuration
- View logs
- Exit application

## Supported AI Providers

### Primary Providers

| Provider | Models | API Key Format |
|----------|--------|----------------|
| OpenAI | GPT-4, GPT-3.5 | sk-... |
| Anthropic | Claude 3.5 | sk-ant-... |
| Google | Gemini Pro/Flash | ... |
| OpenRouter | 300+ models | sk-or-... |

### Additional Providers

- DeepSeek
- Groq (high-speed inference)
- Together AI
- Perplexity
- Mistral AI
- Cohere
- Azure OpenAI
- AWS Bedrock
- And any OpenAI-compatible API

## Building from Source

### Prerequisites

- Go 1.21+
- Git

### Build

```bash
git clone https://github.com/kaminoguo/xiaoniao.git
cd xiaoniao
./build.sh
```

Build artifacts will be created in the `dist/` directory.

### Platform-specific builds

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao cmd/xiaoniao/*.go

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go

# macOS
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao cmd/xiaoniao/*.go
```

## Development

### Project Structure

```
xiaoniao/
├── cmd/xiaoniao/       # Main application entry
├── internal/           # Core modules
│   ├── translator/     # Translation engine
│   ├── clipboard/      # Clipboard monitoring
│   ├── config/         # Configuration management
│   ├── hotkey/         # Global hotkeys
│   ├── tray/          # System tray
│   └── i18n/          # Internationalization
└── assets/            # Resources
```

### Testing

Run tests:
```bash
go test ./...
```

Test specific functionality:
```bash
# Test configuration UI
xiaoniao config

# Test clipboard monitoring
xiaoniao run

# Test API connection
xiaoniao test-api
```

## License

GPL-3.0 License. See [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## Support

Report issues at: https://github.com/kaminoguo/xiaoniao/issues

---

Version 1.6.0 | Updated: 2025-09-07