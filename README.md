# xiaoniao

<div align="center">
  <h1>🐦 xiaoniao</h1>
  <p>基于AI的跨平台剪贴板翻译工具 - 支持20+主流AI Provider</p>
  
  [![License: GPL-3.0](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](LICENSE)
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](go.mod)
  ![Version](https://img.shields.io/badge/Version-v1.5.0-purple)
  ![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20Windows%20%7C%20macOS-blue)
  ![Providers](https://img.shields.io/badge/Providers-20+-green)
  ![Models](https://img.shields.io/badge/Models-300+-orange)
  
  [English](README_EN.md) | 简体中文
</div>

---

## ✨ 特性

- 📋 **智能剪贴板监控** - 自动检测并翻译剪贴板内容，实时替换
- 🤖 **20+ AI Provider支持** - 业界最全的AI服务商支持
- 🔄 **真正的动态模型获取** - 实时从API获取可用模型列表
- 🎨 **灵活的Prompt系统** - 双层架构，底层系统prompt + 用户自定义prompt
- ✏️ **动态Prompt管理** - 实时增删改查，自动持久化到文件
- 🎮 **网络英语翻译** - 游戏聊天、日常网络用语、竞技对话等多种风格
- 🧪 **Prompt测试功能** - 实时测试翻译效果，快速优化prompt
- 🖥️ **优雅的TUI界面** - 终端中的GUI体验，键盘操作优先
- 🌐 **多语言界面** - 支持9种语言界面切换（中/英/日/韩/西/法/德/俄/阿拉伯）
- ⌨️ **智能快捷键系统** - 自动检测桌面环境，智能配置系统快捷键
- 🚀 **轻量运行** - 纯Go实现，内存占用<50MB，二进制约12MB
- 🖥️ **跨平台支持** - 支持Linux、Windows 10/11和macOS 10.15+
- 🔒 **单实例保护** - 防止多个实例同时运行
- 🔧 **智能Provider检测** - 根据API Key自动识别服务商
- 💾 **配置持久化** - 所有设置自动保存，下次启动自动加载
- 🎯 **统一底层Prompt** - 所有Provider共享系统prompt模板
- 🏗️ **XML结构化Prompt** - 使用XML标签提高AI解析准确性

## 🎯 支持的AI Provider

### 顶级Provider（专门优化）
- **OpenAI** - GPT-4o, GPT-4, GPT-3.5等全系列
- **Anthropic** - Claude 3.5 Opus/Sonnet/Haiku全系列  
- **OpenRouter** - 300+模型聚合，智能路由
- **Google** - Gemini Pro/Flash系列
- **DeepSeek** - DeepSeek V3/R1系列
- **Groq** - 超高速推理（18x faster）
- **Together AI** - 200+开源模型

### 其他支持的Provider
- **Perplexity** - AI搜索引擎模型
- **Mistral AI** - 欧洲领先AI
- **Cohere** - 企业级AI
- **Replicate** - 模型托管平台
- **HuggingFace** - 开源模型社区
- **Moonshot** - 月之暗面（中国）
- **Zhipu/GLM** - 智谱AI（中国）
- **Baidu/Qianfan** - 百度文心一言
- **Alibaba/Qwen** - 阿里通义千问
- **Azure OpenAI** - 微软Azure
- **AWS Bedrock** - 亚马逊云
- **以及任何OpenAI兼容API**

## 📦 安装与卸载

### Linux 用户

#### 一键安装：
```bash
# 下载并运行安装脚本
curl -sSL https://github.com/kaminoguo/xiaoniao/releases/latest/download/linux-install.sh | bash
```

#### 一键卸载：
```bash
# 下载并运行卸载脚本
curl -sSL https://github.com/kaminoguo/xiaoniao/releases/latest/download/linux-uninstall.sh | bash
```

### Windows 用户

#### 安装（傻瓜式）：
1. 下载：[xiaoniao-windows.zip](https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-windows.zip)
2. 解压到任意位置
3. 双击 `xiaoniao.exe` 即可使用

#### 卸载：
直接删除解压出来的文件夹即可

### macOS 用户

#### 安装（便携版）：
1. 下载对应架构版本：
   - Intel芯片：[xiaoniao-darwin-amd64.zip](https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-darwin-amd64.zip)
   - Apple Silicon (M1/M2/M3)：[xiaoniao-darwin-arm64.zip](https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-darwin-arm64.zip)
2. 解压到 Applications 或任意位置
3. 首次运行需要右键点击选择「打开」（绕过Gatekeeper）
4. 或在终端运行：`chmod +x xiaoniao && ./xiaoniao config`

#### 卸载：
1. 删除解压的程序文件夹
2. 删除配置文件（可选）：`rm -rf ~/Library/Application\ Support/xiaoniao`

## 🚀 快速开始

### 1. 配置API
```bash
xiaoniao config  # 打开配置界面
```

配置界面功能：
- 🔑 **API配置** - 输入API Key，自动检测Provider
- 🎨 **翻译风格** - 选择或自定义翻译风格（含日语对话）
- 🌐 **界面语言** - 切换中英文界面
- 🎯 **界面主题** - 11种主题可选（Tokyo Night、Catppuccin、Nord等）
- ⌨️ **快捷键设置** - 智能快捷键配置（实时按键检测、冲突检查、自动系统集成）
- 🧪 **连接测试** - 测试API连接状态
- 📝 **模型选择** - 主模型和副模型独立选择

### 2. 启动方式

#### 桌面图标启动（推荐）
- 点击桌面图标自动后台运行
- 系统托盘显示运行状态
- 无终端窗口干扰

#### 命令行启动
```bash
xiaoniao run   # 启动剪贴板监控（带托盘和快捷键）
```

### 3. 托盘菜单功能
- ✅ **监控开关** - 开始/停止剪贴板监控
- 📝 **当前风格** - 显示当前翻译风格
- ✅ **自动粘贴** - 翻译后自动粘贴译文
- 🔄 **刷新配置** - 重新加载配置文件
- ⚙️ **配置** - 打开配置界面
- 🔀 **切换prompt** - 快速切换翻译风格
- 📺 **显示/隐藏终端** - 查看运行日志
- ℹ️ **关于** - 查看版本信息
- ❌ **退出** - 关闭程序

使用方法：
1. 复制任何文本（Ctrl+C）
2. 听到提示音后，翻译已完成
3. 直接粘贴（Ctrl+V）即可得到译文

### 3. 翻译风格示例

| 风格 | 说明 | 适用场景 |
|------|------|----------|
| Gaming Chat | 游戏聊天英语 | 游戏交流 |
| Casual Chat | 休闲网络英语 | 日常聊天 |
| Toxic/Competitive | 竞技垃圾话 | 游戏对抗 |
| 日语对话（朋友） | 休闲日语 | 朋友聊天 |
| 直译 | 准确传达原意 | 通用翻译 |
| 意译 | 符合目标语言习惯 | 日常交流 |
| 文学诗意 | 优美的文学表达 | 文学作品 |
| 学术论文 | 严谨的学术用语 | 学术文献 |
| 技术文档 | 专业技术术语 | 技术文档 |
| 商务正式 | 正式商务用语 | 商务邮件 |

## 🛠️ 高级功能

### 自定义Prompt
```bash
# 在配置界面中
1. 选择"翻译风格"
2. 按 'n' 新建自定义prompt
3. 输入名称和翻译规则
4. 按 't' 测试效果
```

### 快捷键配置

程序支持智能快捷键配置：

1. **自动检测** - 检测桌面环境（GNOME/KDE/XFCE等）
2. **实时检测** - 直接检测用户按下的组合键
3. **冲突检查** - 自动检查系统快捷键冲突
4. **智能配置** - 根据桌面环境自动配置系统快捷键

支持的桌面环境：
- GNOME/GNOME Classic - 完全自动化
- XFCE - 自动配置
- Cinnamon - 自动配置  
- MATE - 自动配置
- KDE Plasma - 需手动配置
- 通用X11 - 通过xbindkeys

### API Key格式

| Provider | Key前缀 | 示例 |
|----------|---------|------|
| OpenAI | sk- | sk-proj-xxxxx |
| Anthropic | sk-ant- | sk-ant-xxxxx |
| OpenRouter | sk-or- | sk-or-xxxxx |
| Groq | gsk_ | gsk_xxxxx |
| Together | sk-/together_ | sk-xxxxx |
| Perplexity | pplx- | pplx-xxxxx |
| Replicate | r8_ | r8_xxxxx |
| HuggingFace | hf_ | hf_xxxxx |

## 📁 项目结构

```
xiaoniao/
├── README.md                    # 项目文档
├── LICENSE                      # GPL-3.0 开源协议
├── go.mod / go.sum             # Go依赖管理
├── build.sh                    # 跨平台构建脚本
├── linux-install.sh            # Linux一键安装
├── linux-uninstall.sh          # Linux一键卸载
├── xiaoniao.bat               # Windows启动脚本
├── start.command              # macOS启动脚本
│
├── cmd/xiaoniao/               # 命令行程序
│   ├── main.go                # 程序入口
│   ├── config_ui.go           # TUI配置界面
│   ├── api_config_ui.go       # API配置界面
│   ├── prompt_test_ui.go      # Prompt测试
│   └── signals_*.go           # 平台信号处理
│
├── internal/                   # 内部包
│   ├── translator/            # 翻译核心（20+ Provider）
│   ├── clipboard/             # 剪贴板监控
│   ├── hotkey/                # 全局快捷键
│   ├── tray/                  # 系统托盘
│   ├── i18n/                  # 国际化（9种语言）
│   ├── sound/                 # 音效系统
│   └── config/                # 配置管理
│
└── assets/                    # 资源文件
    └── icon.png              # 应用图标
```

## 🔧 配置说明

### 主配置
位置：
- Linux: `~/.config/xiaoniao/config.json`
- Windows: `%APPDATA%\xiaoniao\config.json`
- macOS: `~/Library/Application Support/xiaoniao/config.json`

```json
{
  "api_key": "your-api-key",
  "provider": "OpenRouter",
  "model": "openai/gpt-4o-mini",
  "fallback_model": "openai/gpt-3.5-turbo",
  "prompt_id": "direct",
  "language": "cn",
  "auto_paste": true,
  "theme": "tokyo-night",
  "hotkey_toggle": "Ctrl+Alt+X",
  "hotkey_switch": "Ctrl+Alt+P"
}
```

### Prompt配置
位置：
- Linux: `~/.config/xiaoniao/prompts.json`
- Windows: `%APPDATA%\xiaoniao\prompts.json`
- macOS: `~/Library/Application Support/xiaoniao/prompts.json`

```json
[
  {
    "id": "direct",
    "name": "直译",
    "content": "Translate to Chinese directly and accurately."
  }
]
```

> 💡 所有Prompt修改都会实时保存到文件，配置界面和文件始终保持同步

## 📊 性能指标

- **二进制大小**: ~12MB（含TUI框架）
- **内存占用**: < 50MB
- **CPU使用**: < 1%（空闲时）
- **响应时间**: < 100ms（本地处理）
- **翻译延迟**: 1-3秒（API响应）
- **连接复用**: HTTP/2长连接
- **启动速度**: <1秒

## 🐛 故障排除

### 常见问题

1. **API连接失败**
   - 检查API Key是否正确
   - 检查网络连接
   - 某些Provider可能需要代理

2. **模型列表为空**
   - 部分Provider需要特定权限
   - 尝试手动输入模型名称

3. **剪贴板监控不工作**
   - Linux: 确保有剪贴板访问权限，某些Wayland环境可能需要额外配置
   - Windows: 确保程序有剪贴板访问权限，杀毒软件可能会阻止
   - macOS: 首次运行需要在「系统偏好设置 > 安全性与隐私 > 辅助功能」中授权

4. **Windows终端显示问题**
   - 推荐使用 Windows Terminal 或支持 ANSI 的终端
   - PowerShell 7+ 或 Windows Terminal 能提供最佳体验
   - 避免使用传统的 cmd.exe

## 🤝 贡献

欢迎贡献代码、报告问题或提出建议！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 📝 更新日志

### v1.5.0 (2025-09-07) - 最新版本
- 🚀 **跨平台支持**
  - Windows平台完整支持（Windows 10/11）
  - macOS平台完整支持（10.15+，Intel/Apple Silicon）
  - Windows/macOS剪贴板API实现
  - 跨平台系统托盘和热键
  - 配置路径自动适配（Linux: ~/.config，Windows: %APPDATA%，macOS: ~/Library/Application Support）
  
- 🔧 **项目重构**
  - 项目重命名：pixel-translator → xiaoniao
  - 模块路径更新为 github.com/kaminoguo/xiaoniao
  - 使用构建标签(build tags)分离平台特定代码
  - 单一代码库支持多平台，约80%代码共享

### v1.4.1 (2025-09-06)
- 🌍 **完整国际化支持** - 支持9种语言界面（中/英/日/韩/西/法/德/俄/阿拉伯）
- 🔧 **架构优化** - 统一到 `xiaoniao run` 支持所有功能

### v1.4.0 (2025-09-05)
- 🔧 **架构优化** - 统一运行模式，快捷键集成到run模式
- 🎯 **功能改进** - Prompt管理修复，终端日志查看器
- 🐛 **问题修复** - 修复热键录制和托盘图标显示

[查看完整更新历史](https://github.com/kaminoguo/xiaoniao/releases)


## 📄 开源协议

本项目采用 [GPL-3.0](LICENSE) 协议开源

---

更新日期: 2025-09-07

<div align="center">
  <p>如果觉得好用，请给个 ⭐ Star 支持一下！</p>
  <p>Made with ❤️ by Lyrica</p>
</div>