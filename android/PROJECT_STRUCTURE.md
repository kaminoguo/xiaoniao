# Project Structure

## Core Features
- Clipboard translation tool for Android
- Support for 20+ AI providers and 300+ models
- Native GUI using Fyne framework
- Text selection menu integration (ACTION_PROCESS_TEXT)
- Accessibility service for clipboard monitoring
- Material Design UI with 4-tab navigation

## Directory Layout

```
xiaoniao_android/
├── cmd/xiaoniao/               # Main application entry
│   └── main.go                 # Android app main entry point
│
├── internal/                   # Internal packages
│   ├── translator/             # Translation engine (shared with desktop)
│   │   ├── translator.go       # Core translation logic
│   │   ├── provider.go         # Provider interface
│   │   ├── provider_registry.go # Provider registry
│   │   ├── providers_2025.go   # Provider configurations
│   │   ├── openai_compatible.go # OpenAI compatible providers
│   │   ├── openrouter.go       # OpenRouter implementation
│   │   ├── groq_provider.go    # Groq provider
│   │   ├── together_provider.go # Together AI provider
│   │   ├── base_prompt.go      # Base prompt templates
│   │   ├── prompts.go          # Prompt management
│   │   ├── user_prompts.go     # User prompt management
│   │   ├── cache.go            # Translation cache
│   │   ├── stream.go           # Stream processing
│   │   └── http_client.go      # HTTP client configuration
│   │
│   ├── provider/               # AI provider management (shared)
│   │   ├── provider.go         # Provider interface
│   │   ├── detector.go         # Provider auto-detection
│   │   └── registry.go         # Provider registry
│   │
│   ├── models/                 # Model management (shared)
│   │   ├── models.go           # Model definitions
│   │   └── detector.go         # Model detection
│   │
│   ├── i18n/                   # Internationalization (shared)
│   │   ├── i18n.go             # Language management
│   │   ├── strings.go          # String keys
│   │   └── lang/               # Language files
│   │       ├── en_US.json      # English
│   │       ├── zh_CN.json      # Simplified Chinese
│   │       └── ja_JP.json      # Japanese
│   │
│   ├── logbuffer/              # Log buffer (shared)
│   │   └── logbuffer.go        # Circular log buffer
│   │
│   ├── config/                 # Configuration (adapted for Android)
│   │   ├── config.go           # Config structures (shared)
│   │   └── paths.go            # Android-specific paths
│   │
│   ├── ui/                     # Fyne UI layer (Android specific)
│   │   ├── app.go              # Main app structure with bottom navigation
│   │   ├── screens/            # Screen implementations
│   │   │   ├── translate/      # Translation screen (first tab)
│   │   │   │   └── main.go     # Google Translate-like interface
│   │   │   ├── status/         # Status screen (second tab)
│   │   │   │   └── main.go     # Monitor mode selection and status
│   │   │   ├── settings/       # Settings screen (third tab)
│   │   │   │   ├── main.go     # Settings main menu
│   │   │   │   ├── api.go      # API configuration
│   │   │   │   ├── model.go    # Model selection with test
│   │   │   │   ├── prompts.go  # Prompt management
│   │   │   │   └── permissions.go # Permission guide
│   │   │   └── about/          # About screen (fourth tab)
│   │   │       └── main.go     # About page with links
│   │   ├── components/         # Reusable UI components
│   │   │   └── bottom_nav.go   # Bottom navigation bar (4 tabs)
│   │   └── theme/              # Theme customization
│   │       └── material.go     # Material Design theme
│   │
│   ├── accessibility/          # Accessibility service (Android specific)
│   │   ├── service.go          # Service implementation
│   │   └── clipboard.go        # Clipboard monitoring
│   │
│   ├── processtext/            # ACTION_PROCESS_TEXT handler
│   │   ├── activity.go         # Text processing activity
│   │   └── handler.go          # Translation handler
│   │
│   └── jni/                    # JNI bridge
│       ├── bridge.go           # Go-Android communication
│       └── native.c            # Native C code
│
├── android/                    # Android-specific resources
│   ├── AndroidManifest.xml     # Android manifest
│   ├── res/                    # Android resources
│   │   ├── drawable/           # Icon resources
│   │   └── values/             # String resources
│   └── java/                   # Java code for Android integration
│       └── com/liliguo/xiaoniao/
│           ├── AccessibilityService.java
│           └── ProcessTextActivity.java
│
├── assets/                     # Static assets
│   └── icon.png               # App icon (shared with desktop)
│
├── build/                      # Build scripts
│   ├── build.sh               # Build script
│   └── sign.sh                # Signing script
│
├── go.mod                     # Go module definition
├── go.sum                     # Go dependency lock
├── README.md                  # Project readme
├── ANDROID_PORTING.md         # Android porting documentation
└── PROJECT_STRUCTURE.md       # This file
```

## Module Descriptions

### Application Entry (cmd/xiaoniao/)
- Initialize Fyne application
- Set Android-specific configurations
- Launch main UI with 4-tab navigation

### UI Layer (internal/ui/)
- Built with Fyne framework
- Material Design theme
- 4 bottom navigation tabs:
  - Translate: Google Translate-like interface
  - Status: Monitor mode selection and statistics
  - Settings: API, model, prompts, permissions
  - About: Author info and links

### Android Integration

#### Accessibility Service (internal/accessibility/)
- Clipboard monitoring
- System-level text access
- Requires user permission

#### Process Text (internal/processtext/)
- ACTION_PROCESS_TEXT handler
- Text selection menu integration
- Popup translation window

#### JNI Bridge (internal/jni/)
- Go-Android communication
- Native method calls
- Clipboard and UI interaction

### Shared Core Modules (80% code reuse)

#### Translator (internal/translator/)
- Translation engine
- Caching system
- Stream processing
- Prompt management
- Provider integration
- HTTP client configuration

#### Provider (internal/provider/)
- 20+ AI provider support
- Auto-detection based on API key
- Unified interface

#### Models (internal/models/)
- 300+ model support
- Model detection
- Capability checking

#### Log Buffer (internal/logbuffer/)
- Circular buffer
- Memory efficient
- Shared with desktop

#### Internationalization (internal/i18n/)
- Multiple language support
- Dynamic loading
- JSON-based translations

#### Configuration (internal/config/)
- Config structures (mostly shared)
- Android-specific storage paths

## Key Differences from Desktop Version

### Android-Specific Features
1. Text selection menu integration (ACTION_PROCESS_TEXT)
2. Accessibility service for clipboard
3. Mobile-optimized UI with bottom navigation
4. Android permission management
5. Foreground service with notification

### Shared with Desktop (80%)
- Translation engine and logic
- Provider and model management
- Configuration system
- Logging system
- Internationalization

### Platform Comparison

| Feature | Desktop (Win/Mac) | Android |
|---------|------------------|---------|
| UI Framework | System tray + TUI | Fyne GUI with tabs |
| Clipboard Access | Direct API | Accessibility service |
| Text Selection | N/A | ACTION_PROCESS_TEXT |
| Hotkeys | Global hotkeys | N/A |
| Auto-paste | Keyboard simulation | Accessibility service |
| Storage | File system | Android internal storage |
| Background | System service | Foreground service |

## Build Process

### Prerequisites
- Go 1.21 or higher
- Android SDK (API 21+)
- Android NDK
- Fyne CLI tool

### Build Commands

```bash
# Install Fyne CLI
go install fyne.io/fyne/v2/cmd/fyne@latest

# Build APK
cd xiaoniao_android
fyne package -os android -appID com.liliguo.xiaoniao -name xiaoniao -release

# Sign for release
./build/sign.sh xiaoniao.apk
```

## Technical Highlights

### Code Reuse
- **100% Shared**: Translation engine, provider management, model detection
- **90% Shared**: Configuration system, logging, internationalization
- **Android-Specific**: Fyne UI, accessibility service, ACTION_PROCESS_TEXT

### Performance Metrics
- APK size: ~15-20MB
- Memory usage: <80MB (idle)
- CPU usage: <2% (monitoring)
- Translation latency: 1-3 seconds
- Supported models: 300+
- Supported providers: 20+

## Development Guidelines

### Android-Specific Considerations
1. **Permissions**: Properly guide users to enable accessibility service
2. **Battery**: Use foreground service wisely
3. **UI**: Adapt to different screen sizes
4. **Theme**: Keep Material Design consistency

### Debugging

```bash
# View app logs
adb logcat | grep xiaoniao

# Install test build
adb install xiaoniao.apk

# Check permissions
adb shell dumpsys package com.liliguo.xiaoniao
```

## License

MIT License

## Author

Liliguo