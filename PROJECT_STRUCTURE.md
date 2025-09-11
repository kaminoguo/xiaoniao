# Project Structure - Windows专用版

## 目录结构

```
xiaoniao/
├── cmd/xiaoniao/               # 应用程序入口点
│   ├── main.go                 # 主程序
│   ├── config_ui.go            # 配置界面
│   ├── api_config_ui.go        # API配置
│   ├── prompt_test_ui.go       # 提示词测试界面
│   ├── prompts.go              # 提示词管理
│   └── system_hotkey.go        # 系统热键处理
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
│   │   └── user_prompts.go     # 用户提示词管理
│   │
│   ├── i18n/                   # 国际化
│   │   ├── i18n.go             # 语言管理
│   │   ├── lang_zh_cn.go       # 简体中文
│   │   ├── lang_en.go          # 英语
│   │   └── lang_others.go      # 其他语言
│   │
│   ├── clipboard/              # 剪贴板管理
│   │   ├── clipboard.go        # 剪贴板监控器（Linux开发用stub）
│   │   └── clipboard_windows.go # Windows剪贴板完整实现
│   │
│   ├── hotkey/                 # 全局热键
│   │   ├── hotkey.go           # 热键管理（Linux开发用stub）
│   │   └── hotkey_windows.go   # Windows热键完整实现
│   │
│   ├── tray/                   # 系统托盘
│   │   ├── tray.go             # 托盘管理器
│   │   ├── tray_linux.go       # Linux开发用stub
│   │   ├── tray_windows.go     # Windows托盘实现
│   │   ├── icon_embedded.go    # 嵌入式图标资源（ICO格式）
│   │   └── icon_*.ico          # Windows状态图标（蓝/绿/红）
│   │
│   ├── sound/                  # 声音通知
│   │   ├── sound.go            # 声音接口
│   │   └── sound_windows.go    # Windows声音实现（WinMM API）
│   │
│   └── config/                 # 配置
│       ├── themes.go           # UI主题
│       └── config_windows.go   # Windows配置路径
│
├── assets/                     # 应用程序资源
│   └── icon.ico               # Windows应用图标
│
├── build-windows.sh           # Linux上构建Windows版本的脚本
├── install.ps1                # Windows PowerShell安装脚本
├── versioninfo.json           # Windows版本信息
├── go.mod                     # Go模块定义
├── go.sum                     # Go模块校验和
├── LICENSE                    # GPL-3.0许可证
├── README.md                  # 项目文档
└── PROJECT_STRUCTURE.md       # 本文件
```

## 模块描述

### 命令行界面 (cmd/xiaoniao)

主应用程序入口，包含：
- 带主题支持的终端UI配置界面
- API配置和模型选择
- 提示词管理系统
- Windows特定信号处理

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

### 配置 (internal/config)

Windows配置路径：`%APPDATA%\xiaoniao\`

配置文件：
- `config.json`：主要应用设置
- `prompts.json`：自定义翻译提示词

## 构建说明

### 在Linux上交叉编译Windows版本

```bash
# 使用构建脚本
./build-windows.sh

# 或手动编译
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
```

### 在Windows上构建

```cmd
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
```

### 带图标构建

使用goversioninfo嵌入应用图标：

```cmd
go generate
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
```

构建标志说明：
- `-s -w`：减小二进制文件大小（strip符号表）
- **重要**：不使用`-H=windowsgui`，因为配置界面需要控制台窗口（TUI）

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

### v1.6.4 (2025-09-11)
- 完全移除Linux/macOS支持，专注Windows平台
- 简化代码结构，移除平台判断逻辑
- 优化Windows专属功能
- 支持在Linux上交叉编译

### v1.6.3 (2025-09-09)
- 使用goversioninfo修复Windows可执行文件图标嵌入
- 实现多色系统托盘图标（蓝/绿/红）用于状态指示
- 添加嵌入式图标资源（ICO格式）
- 修复Windows系统托盘显示问题
- 增强Windows终端处理配置UI
- 所有图标现在嵌入二进制文件，无需外部文件

### v1.6.2 (2025-09-08)
- 修复Windows系统托盘初始化问题
- 实现Windows特定守护进程初始化
- 系统托盘现在在主线程中正确运行
- 清理不必要的启动脚本

### v1.6.1 (2025-09-08)
- 修复Windows启动行为
- 即使没有API配置也会显示系统托盘
- 未设置API密钥时自动打开配置UI
- 修复托盘初始化中的空指针崩溃
- 简化Windows分发包

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

GPL-3.0许可证

## 作者

Lyrica

---

最后更新：2025-09-09 | 版本：1.6.3