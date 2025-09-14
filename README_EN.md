# xiaoniao

[中文](README.md) | [日本語](README_JP.md) | [한국어](README_KR.md)

Windows Clipboard Translation Tool

## Features

- Monitor clipboard and auto-translate
- Multi-language UI support (CN/EN/JP/KR/FR/ES/DE/RU/AR)
- System tray integration

## Installation

Download [xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) and run

## Usage

1. Run xiaoniao.exe
2. Copy text (Ctrl+C)
3. Auto-translate and replace clipboard
4. Paste (Ctrl+V) to get translation

Tray icon status: Blue-Monitoring / Green-Translating / Red-Stopped

## Configuration

```cmd
xiaoniao.exe config
```

Supports OpenAI, Anthropic, Google, DeepSeek and other APIs

## Build

```bash
# With icon
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

Requirements: Go 1.20+, Windows

## License

MIT