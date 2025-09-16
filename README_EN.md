# xiaoniao

[中文](README.md) | [日本語](README_JP.md) | [한국어](README_KR.md)

Windows Clipboard Translation Tool

## Quick Start

### 1. Configure API Key
- Select "API Configuration" from main menu
- Enter your API key (OpenAI, Anthropic, etc.)
- System will auto-detect the provider

### 2. Select Model
- After setting API, select "Choose Model"
- Pick an AI model from the list

### 3. Set Hotkeys (Optional)
- Select "Hotkey Settings" from main menu
- Configure hotkeys for monitoring toggle and prompt switching

### 4. Start Using
- Ctrl+X to cut or Ctrl+C to copy text triggers translation
- Program auto-replaces clipboard content
- Ctrl+V to paste translated result

## Download

[xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) - Windows 10/11 (64-bit)

## How to Update

1. Delete old xiaoniao.exe
2. Download new xiaoniao.exe
3. Configuration files are saved automatically, won't be lost


## Build

```bash
# With icon
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

Requirements: Go 1.20+, Windows

### Video Tutorials

- Bilibili: [https://www.bilibili.com/video/BV13zpUzhEeK/](https://www.bilibili.com/video/BV13zpUzhEeK/)
- YouTube: [https://www.youtube.com/watch?v=iPye0tYkBaY](https://www.youtube.com/watch?v=iPye0tYkBaY)

## License

MIT

## Support

- Ko-fi: [ko-fi.com/gogogod](https://ko-fi.com/gogogod)
- WeChat: [View QR Code](assets/wechat-pay.jpg)