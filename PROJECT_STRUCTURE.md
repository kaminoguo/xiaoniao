# Project Structure

## Directory Layout

```
xiaoniao/
├── cmd/xiaoniao/               # Application entry point
│   ├── main.go                 # Main program
│   ├── config_ui.go            # Configuration UI
│   ├── api_config_ui.go        # API configuration
│   ├── prompt_test_ui.go       # Prompt testing interface
│   ├── prompts.go              # Prompt management
│   ├── signals_unix.go         # Unix signal handling
│   └── signals_windows.go      # Windows signal handling
│
├── internal/                   # Internal packages
│   ├── translator/             # Translation engine
│   │   ├── translator.go       # Core translation logic
│   │   ├── provider.go         # Provider interface
│   │   ├── provider_registry.go # Provider registry
│   │   ├── providers_2025.go   # Provider configurations
│   │   ├── openai_compatible.go # OpenAI-compatible providers
│   │   ├── openrouter.go       # OpenRouter implementation
│   │   ├── groq_provider.go    # Groq provider
│   │   ├── together_provider.go # Together AI provider
│   │   ├── base_prompt.go      # Base prompt template
│   │   └── user_prompts.go     # User prompt management
│   │
│   ├── i18n/                   # Internationalization
│   │   ├── i18n.go             # Language management
│   │   ├── lang_zh_cn.go       # Simplified Chinese
│   │   ├── lang_en.go          # English
│   │   └── lang_others.go      # Other languages
│   │
│   ├── clipboard/              # Clipboard management
│   │   ├── monitor.go          # Clipboard monitor
│   │   ├── clipboard_linux.go  # Linux implementation
│   │   ├── clipboard_windows.go # Windows implementation
│   │   └── clipboard_darwin.go # macOS implementation
│   │
│   ├── hotkey/                 # Global hotkeys
│   │   ├── hotkey.go           # Linux hotkeys
│   │   ├── hotkey_windows.go   # Windows hotkeys
│   │   └── hotkey_darwin.go    # macOS hotkeys
│   │
│   ├── tray/                   # System tray
│   │   ├── tray.go             # Common tray implementation
│   │   ├── tray_windows.go     # Windows-specific tray
│   │   └── tray_darwin.go      # macOS-specific tray
│   │
│   ├── sound/                  # Sound notifications
│   │   ├── sound.go            # Linux sound
│   │   ├── sound_windows.go    # Windows sound
│   │   ├── sound_darwin.go     # macOS sound
│   │   └── assets/             # Sound files
│   │
│   └── config/                 # Configuration
│       ├── themes.go           # UI themes
│       ├── config_linux.go     # Linux config paths
│       ├── config_windows.go   # Windows config paths
│       └── config_darwin.go    # macOS config paths
│
├── assets/                     # Application resources
│   └── icon.png               # Application icon
│
├── build.sh                    # Build script
├── linux-install.sh           # Linux installer
├── linux-uninstall.sh         # Linux uninstaller
├── xiaoniao.bat               # Windows launcher
├── start.command              # macOS launcher
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── LICENSE                    # GPL-3.0 license
├── README.md                  # Project documentation
└── PROJECT_STRUCTURE.md       # This file
```

## Module Descriptions

### Command Line Interface (cmd/xiaoniao)

The main application entry point containing:
- Terminal UI configuration interface with theme support
- API configuration and model selection
- Prompt management system
- Platform-specific signal handling

### Translation Engine (internal/translator)

Core translation functionality:
- Support for 20+ AI providers
- Dynamic model listing
- Unified prompt system
- Provider auto-detection based on API key format

### Internationalization (internal/i18n)

Multi-language support:
- Simplified Chinese
- Traditional Chinese
- English
- Japanese
- Korean
- Spanish
- French

### Clipboard Management (internal/clipboard)

Platform-specific clipboard monitoring:
- Linux: X11/Wayland support using xclip/xsel/wl-clipboard
- Windows: Windows Clipboard API
- macOS: pbcopy/pbpaste integration

### Configuration (internal/config)

Platform-specific configuration paths:
- Linux: `~/.config/xiaoniao/`
- Windows: `%APPDATA%\xiaoniao\`
- macOS: `~/Library/Application Support/xiaoniao/`

## Build Instructions

### Quick Build

```bash
./build.sh
```

Creates distribution packages in `dist/` directory:
- `xiaoniao-linux-amd64`
- `xiaoniao-windows.zip`
- `xiaoniao-darwin-amd64.zip` (Intel)
- `xiaoniao-darwin-arm64.zip` (Apple Silicon)

### Manual Build

```bash
# Current platform
go build -ldflags="-s -w" -o xiaoniao ./cmd/xiaoniao

# Cross-platform
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao-linux cmd/xiaoniao/*.go
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao-darwin cmd/xiaoniao/*.go
```

## Testing

### Unit Tests

```bash
go test ./...
```

### Integration Tests

```bash
# Configuration interface
xiaoniao config

# Clipboard monitoring
xiaoniao run

# API connectivity
xiaoniao test-api

# Prompt testing
xiaoniao test-prompt
```

### Platform-Specific Testing

#### Linux
```bash
echo "test" | xclip -selection clipboard
xiaoniao run
```

#### Windows
```powershell
Set-Clipboard "test"
.\xiaoniao.exe run
```

#### macOS
```bash
echo "test" | pbcopy
./xiaoniao run
```

## Performance Metrics

- Binary size: ~12MB
- Memory usage: <50MB idle
- CPU usage: <1% monitoring
- Translation latency: 1-3 seconds
- Supported models: 300+ via OpenRouter
- Supported providers: 20+

## Version History

### v1.6.1 (2025-09-08)
- Fixed Windows/macOS startup behavior
- System tray now appears even without API configuration
- Auto-opens configuration UI when no API key is set
- Fixed nil pointer crash in tray initialization
- Removed unnecessary .vbs launcher for Windows
- Simplified Windows distribution package

### v1.6.0 (2025-09-07)
- Full Windows and macOS platform support
- Improved installation scripts
- Language auto-detection

### v1.5.0 (2025-09-07)
- Windows platform support
- Cross-platform build system
- Project renamed from pixel-translator

### v1.4.1 (2025-09-06)
- Complete internationalization
- Support for 7 languages
- Binary location management

## Development Guidelines

### Code Style
- Follow Go standard formatting
- Use meaningful variable names
- Keep functions focused and small
- Handle errors explicitly

### Commit Convention
- feat: New features
- fix: Bug fixes
- docs: Documentation updates
- refactor: Code refactoring
- test: Test additions/changes
- chore: Build/tooling changes

### Platform-Specific Code
- Use build tags for platform separation
- Keep platform-specific code in separate files
- Share common logic where possible
- Test on all supported platforms

## License

GPL-3.0 License

## Author

Lyrica

---

Last updated: 2025-09-08 | Version: 1.6.1