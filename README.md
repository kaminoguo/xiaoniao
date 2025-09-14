# xiaoniao

[English](README_EN.md) | [日本語](README_JP.md) | [한국어](README_KR.md)

Windows 剪贴板翻译工具

## 快速上手

### 1. 配置API密钥
- 在主菜单选择"API配置"
- 输入你的API密钥（如OpenAI、Anthropic等）
- 系统会自动识别提供商

### 2. 选择模型
- 设置API后，选择"选择模型"
- 从列表中选择合适的AI模型

### 3. 设置快捷键（可选）
- 在主菜单选择"快捷键设置"
- 设置监控开关和切换prompt的快捷键

### 4. 开始使用
- Ctrl+C 复制文本触发翻译
- 程序会自动替换剪贴板内容
- Ctrl+V 粘贴翻译结果

## 下载

[xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) - Windows 10/11 (64-bit)

## 构建

```bash
# 带图标
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
cd cmd/xiaoniao && goversioninfo -manifest=../../xiaoniao.exe.manifest -icon=../../assets/icon.ico ../../versioninfo.json
cd ../.. && go build -ldflags="-s -w" -o xiaoniao.exe ./cmd/xiaoniao
```

要求：Go 1.20+, Windows

### 视频教程

- Bilibili: （即将发布）
- YouTube: （即将发布）

## License

MIT

## 支持作者

- Ko-fi: （即将发布）
- 微信赞赏: （即将发布）