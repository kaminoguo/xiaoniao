# xiaoniao

[English](README_EN.md) | [日本語](README_JP.md) | [한국어](README_KR.md)

Windows 剪贴板翻译工具

## 功能

- 监控剪贴板，自动翻译
- 支持多语言界面（中/英/日/韩/法/西/德/俄/阿拉伯语）
- 系统托盘运行

## 安装

下载 [xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) 并运行

## 使用

1. 运行 xiaoniao.exe
2. 复制文本 (Ctrl+C)
3. 自动翻译并替换剪贴板
4. 粘贴 (Ctrl+V) 得到译文

托盘图标状态：蓝色-监控中 / 绿色-翻译中 / 红色-已停止

## 配置

```cmd
xiaoniao.exe config
```

支持 OpenAI、Anthropic、Google、DeepSeek 等 API

## 构建

```bash
# 带图标
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

要求：Go 1.20+, Windows

## License

MIT