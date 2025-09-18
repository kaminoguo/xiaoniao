# Project Structure - macOS版 v1.1.0

## 核心特性
- macOS专用剪贴板翻译工具
- 支持20+AI提供商和300+模型
- 系统菜单栏常驻后台运行
- 全局快捷键控制（Cmd+Alt组合）
- 自动粘贴翻译结果
- 多语言界面支持
- 原生macOS应用体验

## 目录结构

```
xiaoniao_mac/
├── cmd/xiaoniao/               # 应用程序入口点
│   └── main.go                 # 主程序（macOS原生支持）
│
├── internal/                   # 内部包
│   ├── translator/             # 翻译引擎（共享代码）
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
│   ├── i18n/                   # 国际化（共享代码）
│   │   ├── i18n.go             # 语言管理
│   │   ├── lang_zh_cn.go       # 简体中文
│   │   ├── lang_en.go          # 英语
│   │   └── lang_others.go      # 其他语言
│   │
│   ├── clipboard/              # 剪贴板管理（Mac特定）
│   │   └── clipboard.go        # macOS剪贴板实现（pbcopy/pbpaste）
│   │
│   ├── hotkey/                 # 全局热键（Mac特定）
│   │   └── hotkey.go           # macOS热键实现（golang.design/x/hotkey）
│   │
│   ├── tray/                   # 系统托盘（Mac特定）
│   │   ├── tray.go             # 菜单栏管理器
│   │   └── icon.go             # 图标资源
│   │
│   ├── sound/                  # 声音通知（Mac特定）
│   │   └── sound.go            # macOS声音实现（afplay）
│   │
│   ├── config/                 # 配置（Mac特定）
│   │   └── config.go           # macOS配置路径
│   │
│   └── logbuffer/              # 日志缓冲（共享代码）
│       └── logbuffer.go        # 循环日志缓冲区实现
│
├── build.sh                    # macOS构建脚本
├── go.mod                      # Go模块定义
├── go.sum                      # Go模块校验和
└── PROJECT_STRUCTURE.md        # 本文件
```

## 模块描述

### 主应用程序 (cmd/xiaoniao)

macOS原生应用程序，包含：
- 使用mainthread.Init()确保UI操作在主线程
- 菜单栏图标和菜单管理
- 全局快捷键监听（Cmd+Alt+C等）
- 剪贴板监控与翻译
- 配置管理
- 日志记录

### 剪贴板管理 (internal/clipboard)

macOS剪贴板监控：
- 使用pbcopy/pbpaste命令行工具
- 定时轮询剪贴板变化
- 自动翻译检测到的文本
- 防循环翻译机制
- 自动粘贴功能（AppleScript实现）

### 系统托盘 (internal/tray)

macOS菜单栏集成：
- 使用getlantern/systray库
- 动态状态图标（空闲/翻译中/错误）
- 右键菜单功能
- 翻译计数显示
- 提示词切换

### 全局热键 (internal/hotkey)

macOS热键支持：
- 使用golang.design/x/hotkey库
- 支持Cmd、Alt、Ctrl、Shift组合键
- 切换监控状态热键
- 切换提示词热键
- 需要辅助功能权限

### 声音通知 (internal/sound)

macOS系统声音：
- 使用afplay播放系统声音
- 成功音效：Glass.aiff
- 错误音效：Funk.aiff
- 支持自定义音频文件

### 配置 (internal/config)

macOS配置路径：`~/Library/Application Support/xiaoniao/`

配置文件：
- `config.json`：主要应用设置
- `prompts.json`：自定义翻译提示词
- `xiaoniao.log`：应用日志

## macOS特性支持

### 系统集成
- 菜单栏应用（LSUIElement）
- 辅助功能权限（全局快捷键）
- Apple Events权限（自动粘贴）
- 原生.app包格式
- Retina显示支持

### 平台差异

与Windows版本的主要差异：

| 功能 | Windows | macOS |
|-----|---------|-------|
| 剪贴板API | Windows Clipboard API | pbcopy/pbpaste |
| 系统托盘 | 系统托盘（右下角） | 菜单栏（右上角） |
| 快捷键 | Win+Alt组合 | Cmd+Alt组合 |
| 声音 | Windows WinMM | macOS afplay |
| 配置路径 | %APPDATA%\xiaoniao | ~/Library/Application Support/xiaoniao |
| 打包格式 | .exe | .app bundle |
| UI线程 | 默认处理 | 需要mainthread.Init() |

## 构建说明

### 在macOS上构建

```bash
# 使用构建脚本
./build.sh v1.1.0

# 或手动编译
CGO_ENABLED=1 go build -ldflags="-s -w -X main.version=1.1.0" -o xiaoniao ./cmd/xiaoniao
```

### 创建应用包

构建脚本会自动创建.app包，包含：
- Info.plist配置
- 执行权限设置
- 系统权限声明
- 版本信息

### 分发格式

- `xiaoniao` - 命令行可执行文件
- `xiaoniao.app` - macOS应用包
- `xiaoniao-mac-{arch}.zip` - 分发压缩包

## 依赖项

### Go模块依赖
- `github.com/getlantern/systray` - 系统托盘
- `golang.design/x/hotkey` - 全局热键
- `golang.design/x/hotkey/mainthread` - 主线程支持
- 其他共享依赖（HTTP客户端等）

### 系统要求
- macOS 10.12 (Sierra) 或更高版本
- 辅助功能权限（用于全局快捷键）
- 64位处理器（Intel或Apple Silicon）

## 安装说明

1. 下载对应架构的压缩包
2. 解压得到xiaoniao.app
3. 拖动到Applications文件夹
4. 首次运行右键选择"打开"
5. 在系统设置中授予辅助功能权限

## 性能指标

- 应用大小：~15MB
- 内存使用：<60MB（空闲时）
- CPU使用：<1%（监控时）
- 翻译延迟：1-3秒
- 支持模型：300+（通过OpenRouter）
- 支持提供商：20+

## 开发指南

### macOS特定注意事项

1. **主线程要求**：所有UI操作必须在主线程执行
2. **权限处理**：需要正确处理系统权限请求
3. **应用签名**：分发版本需要代码签名
4. **沙盒限制**：某些功能可能受沙盒限制

### 调试技巧

```bash
# 查看应用日志
tail -f ~/Library/Application\ Support/xiaoniao/xiaoniao.log

# 测试剪贴板
echo "test" | pbcopy
pbpaste

# 检查权限
tccutil reset Accessibility com.lyrica.xiaoniao
```

## 版本历史

### v1.1.0-mac (2025-01-18)
- 首个macOS正式版本
- 完整功能移植自Windows版
- 原生macOS体验优化
- 支持Intel和Apple Silicon

## 许可证

MIT许可证

## 作者

Lyrica

---

最后更新：2025-01-18 | 版本：1.1.0-mac