# xiaoniao

Windows专用剪贴板AI翻译工具 v1.6.7，支持20+语言模型。

## 概述

xiaoniao是一款专为Windows设计的剪贴板监控翻译工具，可自动将复制的文本通过AI进行翻译并自动粘贴。它作为后台程序运行，通过系统托盘图标提供控制，实时监控剪贴板变化并自动粘贴翻译后的内容。

## 功能特性

- 实时剪贴板监控与自动翻译
- 自动粘贴翻译结果
- 导出日志功能（v1.6.6新增）
- 支持20+AI提供商，包括OpenAI、Anthropic、Google和OpenRouter
- 可自定义翻译提示词
- 基于终端的配置界面（TUI）
- 系统托盘集成，支持多彩图标状态
- 全局热键支持（文本输入方式配置，v1.6.7改进）
- 多语言界面（中文、英语、日语、韩语、西班牙语、法语）
- 防循环翻译机制

## 系统要求

- Windows 10 或 Windows 11
- 64位系统架构

## 安装

### 下载安装

1. 下载 [xiaoniao-windows.zip](https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-windows.zip)
2. 解压到任意目录
3. 双击运行 xiaoniao.exe
4. 如遇到安全警告，请点击"更多信息" → "仍要运行"

### 快速开始

首次运行时会自动打开配置界面，按照提示设置API密钥即可开始使用。

## 配置

### 初始设置

```cmd
xiaoniao.exe config
```

配置界面允许您：
- 设置所选提供商的API密钥
- 选择翻译模型
- 选择或创建自定义翻译提示词
- 配置界面语言和主题
- 设置全局热键（直接输入快捷键组合，如 Ctrl+C）

### 配置文件位置

配置文件保存在：`%APPDATA%\xiaoniao\`

主要配置文件：
- `config.json`：主要应用程序设置
- `prompts.json`：自定义翻译提示词

## 使用方法

### 启动应用

双击 xiaoniao.exe 启动应用程序。应用程序会在系统托盘显示图标。如果未配置API，配置窗口将自动打开。

### 系统托盘功能

- **切换监控**：开始/停止剪贴板监控
- **切换翻译风格**：选择不同的翻译风格
- **显示调试窗口**：显示/隐藏调试控制台（v1.7.0）
- **设置**：配置API密钥和模型
- **刷新**：重新加载配置
- **退出**：退出应用程序

#### 系统托盘图标状态
- 🔵 **蓝色小鸟**：监控中（空闲）
- 🟢 **绿色小鸟**：正在翻译
- 🔴 **红色小鸟**：监控已停止或发生错误

### 工作原理

1. 复制任意文本到剪贴板
2. xiaoniao自动检测并翻译
3. 原始剪贴板内容被翻译结果替换
4. 粘贴（Ctrl+V）插入翻译后的文本

## 支持的AI提供商

### 主要提供商

| 提供商 | 模型 | API密钥格式 |
|--------|------|-------------|
| OpenAI | GPT-4, GPT-3.5 | sk-... |
| Anthropic | Claude 3.5 | sk-ant-... |
| Google | Gemini Pro/Flash | ... |
| OpenRouter | 300+模型 | sk-or-... |

### 其他提供商

- DeepSeek
- Groq（高速推理）
- Together AI
- Perplexity
- Mistral AI
- Cohere
- Azure OpenAI
- AWS Bedrock
- 任何兼容OpenAI的API

## 从源代码构建

### 前置要求

- Go 1.21+
- Git
- Windows开发环境

### 构建步骤

```cmd
git clone https://github.com/kaminoguo/xiaoniao.git
cd xiaoniao
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
```

构建输出将生成在当前目录。

### 带图标构建（推荐）

```cmd
go generate
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
```

这将生成带有应用图标的可执行文件。

## 开发

### 项目结构

```
xiaoniao/
├── cmd/xiaoniao/       # 主应用程序入口
├── internal/           # 核心模块
│   ├── translator/     # 翻译引擎
│   ├── clipboard/      # 剪贴板监控
│   ├── config/         # 配置管理
│   ├── hotkey/         # 全局热键
│   ├── tray/          # 系统托盘
│   └── i18n/          # 国际化
└── assets/            # 资源文件
```

### 测试

运行测试：
```cmd
go test ./...
```

测试特定功能：
```cmd
# 测试配置界面
xiaoniao.exe config

# 测试剪贴板监控
xiaoniao.exe run

# 测试API连接
xiaoniao.exe test-api
```

## 许可证

GPL-3.0许可证。详见[LICENSE](LICENSE)文件。

## 贡献

1. Fork存储库
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建Pull Request

## 支持

报告问题：https://github.com/kaminoguo/xiaoniao/issues

---

版本 1.6.7 | 更新日期：2025-09-13