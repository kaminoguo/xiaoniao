# 项目结构说明 (Project Structure)

最后更新: 2025-09-06 | 版本: v1.4.1

## 目录树 (Directory Tree)

```
pixel-translator/                # 项目根目录 (CLI-Only)
├── README.md                    # 项目主文档 [核心]
├── PROJECT_STRUCTURE.md        # 项目结构说明（本文件）[核心]
├── PROMPT_SYSTEM.md            # Prompt系统说明 [核心]
├── LICENSE                      # GPL-3.0 开源协议 [核心]
├── go.mod                       # Go模块定义 [核心]
├── go.sum                       # Go依赖校验和 [核心]
├── install.sh                   # 安装脚本 [核心]
├── xiaoniao-launcher.sh        # 桌面启动脚本（后台运行）[核心]
├── update_xiaoniao.sh          # 二进制文件同步脚本 [工具]
│
├── cmd/                         # 应用程序入口 [核心]
│   └── xiaoniao/               # 小鸟翻译CLI程序 [核心]
│       ├── main.go             # 程序主入口（v1.4 - 统一run模式，集成快捷键）[核心]
│       ├── config_ui.go        # TUI配置界面（v1.4 - 修复prompt管理）[核心]
│       ├── system_hotkey.go    # 系统快捷键集成（v1.3新增）[核心]
│       ├── api_config_ui.go    # API配置界面（支持副模型选择）[核心]
│       ├── prompt_test_ui.go   # Prompt测试界面 [核心]
│       └── prompts.go          # Prompt桥接层（调用translator单）[核心]
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
│   │   └── clipboard_linux.go # Linux特定实现（X11/Wayland）[核心]
│   │
│   ├── hotkey/                 # 全局快捷键 [核心]
│   │   └── hotkey.go          # 快捷键实现（golang.design/x/hotkey）[核心]
│   │
│   ├── tray/                   # 系统托盘 [核心]
│   │   └── tray.go            # 托盘实现（使用getlantern/systray）[核心]
│   │
│   ├── sound/                  # 声音提示 [核心]
│   │   ├── sound.go           # 跨平台音效播放 [核心]
│   │   └── assets/            # 音效资源文件 [核心]
│   │
│   └── config/                 # 配置管理模块 [核心]
│       └── themes.go          # 主题配置 [核心]
│
├── assets/                      # 资源文件 [核心]
│   └── icon.png               # 应用图标 [核心]
│
└── docs/                        # 文档目录 [可选]
```

## 核心模块说明

### 1. 命令行界面 (`cmd/xiaoniao/`)
- **config_ui.go**: TUI配置界面主文件
  - 11种精美界面主题
  - 智能快捷键配置系统
  - 自定义Prompt的创建、编辑、删除
  - 配置持久化保存
- **system_hotkey.go**: 系统快捷键集成 (v1.3新增)
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
- 支持X11和Wayland环境
- 智能内容检测

### 4. 配置管理
- 配置文件位置：`~/.config/xiaoniao/`
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

### CLI版本
```bash
# 构建CLI工具（优化大小）
go build -ldflags="-s -w" -o xiaoniao ./cmd/xiaoniao

# 安装到用户目录（推荐）
cp xiaoniao ~/.local/bin/

# 注意：只保留一个版本，避免混淆
# 如果之前安装到/usr/local/bin/，请删除：
# sudo rm -f /usr/local/bin/xiaoniao
```

### 跨平台构建
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o xiaoniao-linux-amd64 cmd/xiaoniao/*.go

# Windows
GOOS=windows GOARCH=amd64 go build -o xiaoniao.exe cmd/xiaoniao/*.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o xiaoniao-macos cmd/xiaoniao/*.go
```

## 性能指标

- **源码大小**: ~400KB
- **二进制大小**: ~11MB (含Bubble Tea TUI框架、系统托盘、国际化)
  - Bubble Tea及Unicode表: ~7MB
  - Systray CGO绑定: ~0.5MB
  - 应用代码: ~0.3MB
- **启动时间**: < 1秒
- **内存占用**: < 50MB (空闲时)
- **CPU使用**: < 1% (监控时)
- **翻译延迟**: 1-3秒 (API响应)
- **连接复用**: HTTP/2长连接
- **支持模型数**: 300+ (通过OpenRouter)
- **支持Provider**: 20+
- **TUI响应**: < 50ms

## 版本历史

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
  - 移除所有"小鸟翻译"、"Translator"等变体

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

- **版本**: v1.4
- **源码大小**: ~370KB
- **支持平台**: Linux (X11/Wayland)
- **依赖管理**: Go Modules
- **最小Go版本**: 1.21+

---

更新日期: 2025-09-05
作者: Lyrica