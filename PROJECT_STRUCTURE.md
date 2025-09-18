# Project Structure - v1.1.0

## 核心特性
- 跨平台剪贴板翻译工具（Windows/macOS）
- 支持20+AI提供商和300+模型
- 系统托盘常驻后台运行
- 全局快捷键控制
- 自动粘贴翻译结果
- 多语言界面支持

## 目录结构

```
xiaoniao/
├── cmd/xiaoniao/               # 应用程序入口点
│   ├── main.go                 # 主程序（GUI模式）
│   ├── config_ui.go            # 配置界面（终端UI）
│   ├── api_config_ui.go        # API配置（3列布局）
│   ├── prompts.go              # 提示词管理
│   ├── system_hotkey.go        # 系统热键处理
│   ├── windows.go              # Windows特定功能（动态调试窗口）
│   └── hotkey_input.go         # 快捷键输入处理
│
├── internal/                   # 内部包
│   ├── translator/             # 翻译引擎
│   │   ├── translator.go       # 核心翻译逻辑
│   │   ├── provider.go         # 提供商接口
│   │   ├── provider_registry.go # 提供商注册表
│   │   ├── providers_2025.go   # 提供商配置
│   │   ├── openai_compatible.go # OpenAI兼容提供商
│   │   ├── openrouter.go       # OpenRouter实现
│   │   ├── groq_provider.go    # Groq提供商
│   │   ├── together_provider.go # Together AI提供商
│   │   ├── base_prompt.go      # 基础提示词模板
│   │   ├── prompts.go          # 提示词管理
│   │   ├── user_prompts.go     # 用户提示词管理
│   │   └── http_client.go      # HTTP客户端配置
│   │
│   ├── i18n/                   # 国际化
│   │   ├── i18n.go             # 语言管理
│   │   ├── lang_zh_cn.go       # 简体中文
│   │   ├── lang_en.go          # 英语
│   │   └── lang_others.go      # 其他语言
│   │
│   ├── clipboard/              # 剪贴板管理
│   │   ├── clipboard.go        # 剪贴板监控器（通用stub）
│   │   ├── clipboard_windows.go # Windows剪贴板实现
│   │   └── clipboard_darwin.go  # macOS剪贴板实现
│   │
│   ├── hotkey/                 # 全局热键
│   │   ├── hotkey.go           # 热键管理（通用stub）
│   │   ├── hotkey_windows.go   # Windows热键实现
│   │   └── hotkey_darwin.go    # macOS热键实现
│   │
│   ├── tray/                   # 系统托盘
│   │   ├── tray.go             # 托盘管理器
│   │   ├── tray_linux.go       # Linux开发用stub
│   │   ├── tray_windows.go     # Windows托盘实现
│   │   ├── tray_darwin.go      # macOS托盘实现
│   │   ├── icon_embedded.go    # 嵌入式图标资源（ICO格式）
│   │   └── icon_*.ico          # 状态图标（蓝/绿/红）
│   │
│   ├── sound/                  # 声音通知
│   │   ├── sound.go            # 声音接口
│   │   ├── sound_windows.go    # Windows声音实现（WinMM API）
│   │   └── sound_darwin.go     # macOS声音实现（afplay）
│   │
│   ├── config/                 # 配置
│   │   ├── themes.go           # UI主题
│   │   ├── config_windows.go   # Windows配置路径
│   │   └── config_darwin.go    # macOS配置路径
│   │
│   └── logbuffer/              # 日志缓冲
│       └── logbuffer.go        # 循环日志缓冲区实现
│
├── assets/                     # 应用程序资源
│   └── icon.ico               # Windows应用图标
│
├── .github/                   # GitHub配置
│   └── workflows/
│       └── build-release.yml   # 多平台自动构建和发布
│
├── build-windows.sh           # Linux上构建Windows版本的脚本
├── build-mac.sh               # macOS构建脚本
├── install.ps1                # Windows PowerShell安装脚本
├── versioninfo.json           # Windows版本信息
├── go.mod                     # Go模块定义
├── go.sum                     # Go模块校验和
├── LICENSE                    # MIT许可证
├── README.md                  # 项目文档（中文）
├── README_EN.md               # 英文文档
├── README_JP.md               # 日文文档
├── README_KR.md               # 韩文文档
├── PROJECT_STRUCTURE.md       # 本文件
├── PROMPT_SYSTEM.md           # 提示词系统文档
├── FILE_DESCRIPTIONS.md       # 文件功能描述
└── RELEASE_v*.md              # 版本发布说明
```

## 模块描述

### 主应用程序 (cmd/xiaoniao)

GUI模式Windows应用程序，包含：
- 终端UI配置界面（支持多种主题）
- API配置和模型选择（20+提供商）
- 提示词管理系统
- 快捷键配置（手动输入模式）
- 动态调试窗口（AllocConsole/FreeConsole）
- 实时日志查看和导出功能

### 翻译引擎 (internal/translator)

核心翻译功能：
- 支持20+AI提供商
- 动态模型列表
- 统一提示词系统
- 基于API密钥格式的提供商自动检测

### 国际化 (internal/i18n)

多语言支持：
- 简体中文
- 繁体中文
- 英语
- 日语
- 韩语
- 西班牙语
- 法语

### 剪贴板管理 (internal/clipboard)

Windows剪贴板监控：
- 使用Windows Clipboard API
- 实时监控剪贴板变化
- 自动替换翻译内容
- 防循环翻译机制
- 自动粘贴功能（Ctrl+V模拟）

### 配置 (internal/config)

Windows配置路径：`%APPDATA%\xiaoniao\`

配置文件：
- `config.json`：主要应用设置
- `prompts.json`：自定义翻译提示词

### 日志缓冲 (internal/logbuffer)

循环日志缓冲区：
- 固定大小的循环缓冲区（1000行）
- 实时日志查看
- 日志导出功能
- 线程安全操作

## 构建说明

### 在Linux上交叉编译Windows版本

```bash
# 使用构建脚本（自动处理版本信息）
./build-windows.sh

# 或手动编译（GUI模式）
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui -X main.version=1.0.1" -o xiaoniao.exe ./cmd/xiaoniao
```

### 在Windows上构建

```cmd
go build -ldflags="-s -w -H windowsgui -X main.version=1.0.1" -o xiaoniao.exe ./cmd/xiaoniao
```

### 带图标构建

使用rsrc嵌入应用图标：

```bash
~/go/bin/rsrc -ico assets/icon.ico -o cmd/xiaoniao/resource.syso
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui -X main.version=1.0.1" -o xiaoniao.exe ./cmd/xiaoniao
```

构建标志说明：
- `-s -w`：减小二进制文件大小（strip符号表）
- `-H windowsgui`：GUI模式，无控制台窗口
- `-X main.version=1.0.1`：设置版本号

## 测试

### 单元测试

```cmd
go test ./...
```

### 集成测试

```cmd
# 配置界面
xiaoniao.exe config

# 剪贴板监控
xiaoniao.exe run

# API连接测试
xiaoniao.exe test-api

# 提示词测试
xiaoniao.exe test-prompt
```

### Windows特定测试

```powershell
# PowerShell测试剪贴板
Set-Clipboard "test text"
.\xiaoniao.exe run

# 测试系统托盘
.\xiaoniao.exe
```

## 性能指标

- 二进制文件大小：~12MB
- 内存使用：<50MB（空闲时）
- CPU使用：<1%（监控时）
- 翻译延迟：1-3秒
- 支持模型：300+（通过OpenRouter）
- 支持提供商：20+

## 版本历史

### v1.1.0 (2025-09-18)
- 新增macOS支持（Intel和Apple Silicon）
- 跨平台架构重构
- GitHub Actions自动构建
- 优化平台特定功能

### v1.0.1 (2025-09-18)
- 重大更新：全新品牌和版本号体系
- 优化翻译提示词系统
- 改进日志缓冲和导出功能
- 增强配置界面用户体验
- 支持更多AI提供商

### v1.0.0 (2025-09-17)
- 正式发布版本
- 完整的Windows平台支持
- 20+AI提供商集成
- 多语言界面
- 系统托盘常驻

## 开发指南

### 代码风格
- 遵循Go标准格式
- 使用有意义的变量名
- 保持函数专注且小巧
- 显式处理错误

### 提交约定
- feat: 新功能
- fix: 错误修复
- docs: 文档更新
- refactor: 代码重构
- test: 测试添加/更改
- chore: 构建/工具更改

### Windows特定代码
- 使用构建标签进行平台分离
- 将Windows特定代码保存在单独的文件中
- 尽可能共享通用逻辑
- 在Windows上充分测试

## 许可证

MIT许可证

## 作者

Lyrica

---

最后更新：2025-09-18 | 版本：1.1.0