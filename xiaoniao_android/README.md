# xiaoniao Android

A clipboard translation tool for Android with support for 20+ AI providers and 300+ models.

## Features

- Text selection menu integration (ACTION_PROCESS_TEXT)
- Clipboard monitoring with user control
- Support for multiple AI providers (OpenAI, Anthropic, Google, etc.)
- Customizable translation prompts
- Minimal battery usage
- Material Design UI

## Requirements

- Android 5.0 (API 21) or higher
- Network connection for API calls
- API key from supported provider

## Installation

Download the latest APK from the releases page and install on your Android device.

## Usage

### Text Menu Mode (Recommended)

1. Select any text in any app
2. Choose "xiaoniao" from the text selection menu
3. View translation in popup window
4. Copy result or retry with different prompt

### Clipboard Mode

1. Open xiaoniao app
2. Switch to clipboard monitoring mode
3. Grant necessary permissions
4. Copy any text to translate automatically

## Configuration

### First Time Setup

1. Enter your API key
2. Select AI model
3. Choose default prompt
4. Grant permissions if using clipboard mode

### Supported Providers

The app automatically detects provider based on API key format:
- OpenAI (sk-...)
- Anthropic (sk-ant-...)
- Google (AIza...)
- Groq (gsk_...)
- And 15+ more providers

## Permissions

### Required
- Internet access for API calls

### Optional (for clipboard mode)
- Accessibility service for clipboard monitoring
- Notification permission for background service
- Battery optimization exemption for reliable operation

## Building from Source

### Prerequisites

- Go 1.21 or higher
- Android SDK
- Android NDK
- Fyne CLI tool

### Build Steps

```bash
# Install Fyne CLI
go install fyne.io/fyne/v2/cmd/fyne@latest

# Clone repository
git clone https://github.com/username/xiaoniao
cd xiaoniao/xiaoniao_android

# Build APK
fyne package -os android -appID com.liliguo.xiaoniao -name xiaoniao -release

# Sign APK for release
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 \
  -keystore xiaoniao.keystore xiaoniao.apk xiaoniao
```

## Project Structure

See [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) for detailed project structure.

See [ANDROID_PORTING.md](ANDROID_PORTING.md) for Android porting technical details.

## Technical Details

### Architecture

The app uses a hybrid approach combining native Android features with Go-based translation logic:

1. UI Layer: Fyne framework for cross-platform GUI
2. Translation Layer: Shared Go code with desktop version
3. Platform Layer: JNI bridge for Android-specific features
4. Service Layer: Optional background service for clipboard monitoring

### Code Reuse

Approximately 80% of the code is shared with the desktop version:
- Translation engine
- Provider management
- Model detection
- Configuration system
- Internationalization

Android-specific code handles:
- Text selection menu integration
- Accessibility service
- Android permissions
- Mobile UI adaptations

## Troubleshooting

### Text menu not appearing
- Ensure xiaoniao is installed properly
- Some apps may block custom text actions (WeChat, Telegram)
- Try restarting the device

### Clipboard monitoring not working
- Check accessibility service is enabled in system settings
- Grant all required permissions
- Disable battery optimization for xiaoniao

### Translation errors
- Verify API key is correct
- Check network connection
- Ensure selected model is available for your API key

## License

MIT License

## Author

Liliguo

## Acknowledgments

Thanks to all users of xiaoniao. Your feedback helps improve the app.

Price: $1 suggested donation, but you can use it for free.