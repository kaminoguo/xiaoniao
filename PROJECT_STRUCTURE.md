# 项目结构说明 (Project Structure)

最后更新: 2025-09-07 | 版本: v1.6.0 | GitHub: https://github.com/kaminoguo/xiaoniao

## 目录树 (Directory Tree)

```
xiaoniao/                        # 项目根目录 (跨平台CLI)
├── README.md                    # 项目主文档
├── PROJECT_STRUCTURE.md        # 项目结构说明（本文件）
├── LICENSE                      # GPL-3.0 开源协议
├── go.mod                       # Go模块定义
├── go.sum                       # Go依赖校验和
├── build.sh                     # 跨平台构建脚本
├── linux-install.sh            # Linux一键安装脚本
├── linux-uninstall.sh          # Linux一键卸载脚本
├── xiaoniao.bat                # Windows启动脚本
├── start.command               # macOS启动脚本
│
├── cmd/                         # 应用程序入口
│   └── xiaoniao/               # xiaoniao CLI程序
│       ├── main.go             # 程序主入口
│       ├── config_ui.go        # TUI配置界面
│       ├── api_config_ui.go    # API配置界面
│       ├── prompt_test_ui.go   # Prompt测试界面
│       ├── prompts.go          # Prompt桥接层
│       ├── signals_unix.go     # Unix信号处理（Linux/macOS）
│       └── signals_windows.go  # Windows信号处理
│
├── internal/                    # 内部包（不对外暴露）[核心]
│   ├── translator/             # 翻译核心模块 [核心]
│   │   ├── translator.go      # 翻译器主逻辑（无缓存）[核心]
│   │   ├── provider.go         # Provider接口定义和基础实现 [核心]
│   │   ├── provider_registry.go # Provider注册表（20+服务商）[核心]
│   │   ├── providers_2025.go  # 2025年Provider配置 [核心]
│   │   ├── openai_compatible.go # OpenAI兼容Provider [核心]
│   │   ├── openrouter.go      # OpenRouter实现（含n=1参数优化）[核心]
│   │   ├── groq_provider.go   # Groq高速推理 [核心]
│   │   ├── together_provider.go # Together AI [核心]
│   │   ├── base_prompt.go     # 统一的底层系统prompt模板（无语言限制）[核心]
│   │   └── user_prompts.go    # 用户prompt管理 [核心]
│   │
│   ├── i18n/                   # 国际化模块 [核心]
│   │   ├── i18n.go            # 国际化主逻辑（语言检测、切换）[核心]
│   │   ├── lang_zh_cn.go      # 简体中文翻译（350+字段）[核心]
│   │   ├── lang_en.go         # 英文翻译（350+字段）[核心]
│   │   └── lang_others.go     # 其他语言（日/韩/西/法/德/俄/阿拉伯/繁体，完整翻译）[核心]
│   │
│   ├── clipboard/              # 剪贴板管理模块 [核心]
│   │   ├── monitor.go         # 剪贴板监控器（含循环防护）[核心]
│   │   ├── clipboard_linux.go # Linux特定实现（X11/Wayland）[核心]
│   │   ├── clipboard_windows.go # Windows特定实现（Windows API）[核心]
│   │   └── clipboard_darwin.go # macOS特定实现（pbcopy/pbpaste）[核心]
│   │
│   ├── hotkey/                 # 全局快捷键 [核心]
│   │   ├── hotkey.go          # Linux快捷键实现（golang.design/x/hotkey）[核心]
│   │   ├── hotkey_windows.go  # Windows快捷键实现 [核心]
│   │   └── hotkey_darwin.go   # macOS快捷键实现 [核心]
│   │
│   ├── tray/                   # 系统托盘 [核心]
│   │   ├── tray.go            # 托盘通用实现（使用getlantern/systray）[核心]
│   │   ├── tray_windows.go   # Windows托盘特定功能 [核心]
│   │   └── tray_darwin.go    # macOS托盘特定功能 [核心]
│   │
│   ├── sound/                  # 声音提示 [核心]
│   │   ├── sound.go           # Linux音效播放 [核心]
│   │   ├── sound_windows.go   # Windows音效播放 [核心]
│   │   ├── sound_darwin.go    # macOS音效播放 [核心]
│   │   └── assets/            # 音效资源文件 [核心]
│   │
│   └── config/                 # 配置管理模块 [核心]
│       ├── themes.go          # 主题配置 [核心]
│       ├── config_linux.go    # Linux配置路径 (~/.config/xiaoniao) [核心]
│       ├── config_windows.go  # Windows配置路径 (%APPDATA%\xiaoniao) [核心]
│       └── config_darwin.go   # macOS配置路径 (~/Library/Application Support/xiaoniao) [核心]
│
└── assets/                      # 资源文件
    └── icon.png               # 应用图标
```

## 核心模块说明

### 1. 命令行界面 (`cmd/xiaoniao/`)
- **config_ui.go**: TUI配置界面主文件
  - 11种精美界面主题
  - 智能快捷键配置系统
  - 自定义Prompt的创建、编辑、删除
  - 配置持久化保存
- **signals_unix.go/signals_windows.go**: 平台特定信号处理
  - 自动检测桌面环境
  - 智能冲突检查
  - 一键配置系统快捷键
  - 支持GNOME/KDE/XFCE等环境
- **api_config_ui.go**: API配置专用界面
  - 自动检测Provider（通过API Key前缀）
  - 动态获取模型列表（不是硬编码）
  - 连接测试功能
  - 模型搜索和选择
- **prompt_test_ui.go**: Prompt测试界面
  - 实时测试翻译效果
  - 快速优化prompt

### 2. 翻译核心 (`internal/translator/`)

#### Provider系统
- **支持20+ AI Provider**:
  - OpenAI、Anthropic、Google、DeepSeek
  - OpenRouter（300+模型聚合）
  - Groq（超高速推理）
  - Together AI（200+开源模型）
  - 更多Provider...

### 3. 剪贴板管理 (`internal/clipboard/`)
- 实时监控剪贴板变化
- 自动替换翻译结果
- 跨平台支持：
  - Linux: X11/Wayland (xclip/xsel/wl-clipboard)
  - Windows: Windows Clipboard API
  - macOS: pbcopy/pbpaste
- 智能内容检测和循环防护

### 4. 配置管理
- 配置文件位置：
  - Linux: `~/.config/xiaoniao/`
  - Windows: `%APPDATA%\xiaoniao\`
  - macOS: `~/Library/Application Support/xiaoniao/`
- 支持热重载
- 多配置文件：
  - `config.json`: 主配置
  - `prompts.json`: 用户自定义Prompt

## 关键特性实现

### 动态模型获取
```go
// 不是硬编码模型列表，而是从API动态获取
func (p *OpenRouterProvider) ListModels() ([]string, error) {
    // 调用API获取实时模型列表
    // 返回300+可用模型
}
```

### Provider自动检测
```go
// 根据API Key前缀自动识别Provider
func DetectProviderByKey(apiKey string) string {
    switch {
    case strings.HasPrefix(apiKey, "sk-ant-"):
        return "Anthropic"
    case strings.HasPrefix(apiKey, "sk-or-"):
        return "OpenRouter"
    case strings.HasPrefix(apiKey, "gsk_"):
        return "Groq"
    // ... 更多Provider
    }
}
```

## Prompt系统架构 (v1.1更新)

### 双层Prompt设计
1. **底层系统Prompt** (base_prompt.go)
   - 移除了中英文翻译限制，支持任意语言
   - 统一的任务框架，所有Provider共享
   - 添加n=1, top_p=0.9参数防止多版本生成
   
2. **用户自定义Prompt** (user_prompts.go)
   - 动态读写 `~/.config/xiaoniao/prompts.json`
   - 支持实时增删改查
   - 配置界面与文件始终同步
   - 新增日语朋友对话prompt

## 技术栈

### 后端
- **语言**: Go 1.21+
- **TUI框架**: Bubbletea + Lipgloss
- **HTTP客户端**: 标准库 net/http
- **JSON处理**: encoding/json

### 依赖管理
- Go Modules (go.mod)
- 最小化外部依赖
- 无需额外运行时

## 构建说明

### 快速构建
```bash
# 使用自动构建脚本（推荐）
chmod +x build.sh
./build.sh

# 构建产物在 dist/ 目录：
# - xiaoniao-linux-amd64
# - xiaoniao-windows.zip
# - xiaoniao-darwin-amd64.zip (Intel Mac)
# - xiaoniao-darwin-arm64.zip (Apple Silicon)
```

### 本地开发构建
```bash
# 构建当前平台版本
go build -ldflags="-s -w" -o xiaoniao ./cmd/xiaoniao

# 安装到用户目录（Linux/macOS）
cp xiaoniao ~/.local/bin/

# 注意：只保留一个版本，避免混淆
# 如果之前安装到/usr/local/bin/，请删除：
# sudo rm -f /usr/local/bin/xiaoniao
```

### 跨平台手动构建
```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao-linux-amd64 cmd/xiaoniao/*.go

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao-darwin-amd64 cmd/xiaoniao/*.go

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o xiaoniao-darwin-arm64 cmd/xiaoniao/*.go
```

## 测试说明

### 功能测试
```bash
# 1. 测试配置界面
xiaoniao config

# 2. 测试剪贴板监控
xiaoniao run
# 复制任意文本测试翻译功能

# 3. 测试API连接
xiaoniao test-api

# 4. 测试Prompt
xiaoniao test-prompt
```

### 平台特定测试

#### Linux测试
```bash
# 测试剪贴板（X11）
echo "test" | xclip -selection clipboard
xiaoniao run

# 测试剪贴板（Wayland）
echo "test" | wl-copy
xiaoniao run

# 测试快捷键
xiaoniao config  # 配置快捷键后测试
```

#### Windows测试
```powershell
# 在 PowerShell 中测试
Set-Clipboard "test text"
.\xiaoniao.exe run

# 测试托盘图标
.\xiaoniao.exe run  # 检查系统托盘
```

#### macOS测试
```bash
# 测试剪贴板
echo "test" | pbcopy
./xiaoniao run

# 测试通知
./xiaoniao run  # 翻译时应显示系统通知

# 测试快捷键（需要辅助功能权限）
# 系统偏好设置 > 安全性与隐私 > 辅助功能
```

### 多语言测试
```bash
# 测试不同系统语言
LANG=zh_CN.UTF-8 xiaoniao config  # 中文界面
LANG=en_US.UTF-8 xiaoniao config  # 英文界面
LANG=ja_JP.UTF-8 xiaoniao config  # 日文界面
```

## 性能指标

- **源码大小**: ~400KB
- **二进制大小**: ~12MB (含TUI框架、系统托盘、国际化)
- **启动时间**: < 1秒
- **内存占用**: < 50MB (空闲时)
- **CPU使用**: < 1% (监控时)
- **翻译延迟**: 1-3秒 (API响应)
- **连接复用**: HTTP/2长连接
- **支持模型数**: 300+ (通过OpenRouter)
- **支持Provider**: 20+
- **TUI响应**: < 50ms

## 版本历史

### v1.6.0 (2025-09-07) - macOS支持与智能安装
- 🍎 **macOS完整支持**
  - 添加macOS剪贴板实现（pbcopy/pbpaste）
  - macOS系统托盘和通知中心集成
  - macOS热键支持（Cmd/Option/Control）
  - 配置路径：~/Library/Application Support/xiaoniao
  - 支持Intel和Apple Silicon架构
- 🔧 **智能安装脚本**
  - Linux安装脚本自动检测系统语言（9种语言）
  - 自动检测桌面环境（GNOME/KDE/XFCE/Hyprland等）
  - 自动检测终端类型优化TUI显示
  - 首次运行自动设置界面语言与系统一致
- 📦 **便携版发布**
  - Windows/macOS提供ZIP便携版
  - 无需安装，解压即用
  - 自动检测配置，首次运行打开配置界面

### v1.5.0 (2025-09-07) - Windows支持
- 🚀 **Windows平台支持**
  - 添加Windows剪贴板API实现
  - Windows系统托盘支持
  - Windows热键注册
  - 配置路径自动适配（%APPDATA%）
- 🔧 **项目重构**
  - 项目重命名：pixel-translator → xiaoniao
  - 模块路径更新为 github.com/kaminoguo/xiaoniao
  - 使用构建标签分离平台特定代码

### v1.4.1 (2025-09-06) - 完整国际化
- 🌍 **国际化完善**
  - 支持9种语言界面（350+翻译字段）
  - 修复所有硬编码文本
  - 统一二进制文件位置管理
  - 移除冗余的 `xiaoniao tray` 命令
- 🔧 **代码清理**
  - 删除过时的测试二进制文件
  - 优化项目结构
  - 提升代码质量

### v1.4 (2025-09-05) - 架构优化
- 🔧 **架构优化**
  - 统一运行模式：移除冗余的tray命令，run命令集成所有功能
  - 快捷键支持集成到run模式，无需单独的tray模式
  - 后台运行优化：使用nohup实现真正的后台运行
- 🎯 **功能改进**
  - 终端日志查看器：显示/隐藏终端改为打开日志查看器
  - Prompt管理修复：修复无法添加第二个prompt的bug
  - ID生成优化：自动递增生成唯一的prompt ID  
  - 国际化优化：作者名固定显示"梨梨果"
- 🐛 **问题修复**
  - 修复Alt键检测不灵敏问题
  - 修复Ctrl+Alt+X热键无法录制（清理系统残留注册）
  - 修复彩色托盘图标显示（重新生成红绿图标）
  - 修复版本号显示（统一为v1.4）
  - 修复桌面启动器路径问题
  - 移除"直译"默认prompt名称

### v1.3.2 (2025-09-04) - 国际化更新
- 🌐 **完整国际化支持**
  - 支持9种语言界面（简体/繁体/英/日/韩/西/法/德/俄/阿拉伯）
  - 基于系统语言自动检测
  - 所有UI元素全面国际化
- 🔒 **单实例保护**
  - PID锁文件机制
  - 防止多个实例同时运行
  - 自动清理守护锁文件
- 🎯 **品牌统一**  
  - 所有语言下统一使用"xiaoniao"名称
  - 统一使用"xiaoniao"名称

### v1.3.1 (2025-09-04) - 优化更新
- 🔧 **代码优化**
  - 移除鼠标点击支持，专注键盘操作
  - 优化二进制大小（使用-ldflags="-s -w"）
  - 清理关于页面的多余信息
- 🐛 **问题修复**
  - 修复API配置页面选项无法选择问题
  - 修复从托盘打开配置界面闪退
  - 修复ptyxis终端兼容性

### v1.3 (2025-09-03) - 智能快捷键系统
- ⌨️ **全新快捷键系统**
  - 实时按键检测，无需手动输入
  - 自动检测桌面环境（GNOME/KDE/XFCE等）
  - 智能冲突检查，避免系统快捷键冲突
  - 一键配置系统快捷键，无需手动设置
- 🎨 **UI优化**
  - 托盘菜单显示当前翻译风格
  - 优化通知系统，减少不必要的弹窗
- 🔧 **代码优化**
  - 清理测试文件和临时文件
  - 支持更多桌面环境

### v1.2 (2025-09-02) - 全局快捷键
- 🔑 **全局快捷键支持**
  - 使用 golang.design/x/hotkey 实现原生快捷键
  - 支持快捷键实时检测和录制
- 🔧 **功能优化**
  - 改进快捷键设置界面
  - 优化启动脚本代理检测
  - 修复托盘菜单崩溃问题

### v1.1 (2025-09-01) - 功能增强
- 🎆 **新增功能**
  - 日语翻译prompt（朋友对话风格）
  - 副模型选择功能（与主模型一样的UI）
  - 桌面图标无API时自动打开配置
- 🔧 **修复与优化**
  - 修复prompt编辑区太小的问题（现在可显示12行）
  - 修复prompt保存同步问题
  - 移除底层prompt中英文翻译限制
  - 优化API请求参数，防止生成多个翻译版本
  - 降低Token使用量（添加n=1, top_p=0.9参数）

### v1.0.0 (2025-08-31) - 正式发布
- 🎉 **完整功能发布**
  - 20+ AI Provider支持
  - 300+ 模型通过OpenRouter
  - 36种预设翻译风格
  - 自定义Prompt系统
- 🚀 **性能优化**
  - HTTP连接池100%复用率（测试验证）
  - 预热机制首次请求提升47%
  - 响应时间优化到2-3秒
  - 全局共享HTTP客户端
- ✨ **核心功能**
  - 智能剪贴板监控
  - 系统托盘模式
  - TUI配置界面
  - 动态模型获取
  - 实时Prompt测试
- 🌍 **多语言支持**
  - 中英文界面
  - 多种翻译风格
  - 网络英语模式
- 🔧 **技术实现**
  - 纯Go实现，无依赖
  - 内存占用<50MB
  - 项目大小324KB

## 开发规范

### 代码风格
- 遵循Go官方代码规范
- 使用gofmt格式化
- 函数注释使用中文
- 错误处理优先返回error

### 提交规范
- feat: 新功能
- fix: Bug修复
- docs: 文档更新
- style: 代码风格
- refactor: 重构
- test: 测试
- chore: 构建/工具

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'feat: Add AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 许可证

GPL-3.0 License - 详见 [LICENSE](LICENSE) 文件

## 当前状态

- **版本**: v1.6.0
- **源码大小**: ~450KB
- **支持平台**: 
  - Linux (X11/Wayland) - 所有主流发行版
  - Windows 10/11
  - macOS 10.15+ (Intel & Apple Silicon)
- **依赖管理**: Go Modules
- **最小Go版本**: 1.21+
- **跨平台**: 使用构建标签(build tags)分离平台代码
- **代码共享率**: ~80%（平台特定代码仅20%）

---

更新日期: 2025-09-07
作者: Lyrica