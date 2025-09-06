# xiaoniao

<div align="center">
  <h1>🐦 xiaoniao</h1>
  <p>基于AI的Linux剪贴板翻译工具 - 支持20+主流AI Provider</p>
  
  [![License: GPL-3.0](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](LICENSE)
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](go.mod)
  ![Version](https://img.shields.io/badge/Version-v1.4.1-purple)
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
- 🚀 **轻量运行** - 纯Go实现，内存占用<50MB，二进制约8MB
- 🔒 **单实例保护** - 防止多个实例同时运行
- 🔧 **智能Provider检测** - 根据API Key自动识别服务商
- 💾 **配置持久化** - 所有设置自动保存，下次启动自动加载
- 🎯 **统一底层Prompt** - 所有Provider共享系统prompt模板（v2.5.0优化）
- 🏗️ **XML结构化Prompt** - 使用XML标签提高AI解析准确性（v2.5.0新增）

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

## 📦 安装

### 快速安装（推荐）
```bash
# 下载并安装
curl -sSL https://github.com/yourusername/pixel-translator/releases/latest/download/install.sh | bash

# 或者手动下载二进制文件
wget https://github.com/yourusername/pixel-translator/releases/latest/download/xiaoniao-linux-amd64
chmod +x xiaoniao-linux-amd64
sudo mv xiaoniao-linux-amd64 /usr/local/bin/xiaoniao
```

### 从源码构建
```bash
# 克隆仓库
git clone https://github.com/yourusername/pixel-translator.git
cd pixel-translator

# 构建小鸟翻译CLI（优化大小）
go build -ldflags="-s -w" -o xiaoniao ./cmd/xiaoniao
sudo mv xiaoniao /usr/local/bin/

# 构建GUI版本（需要Wails）
wails build
```

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
pixel-translator/
├── README.md                    # 项目文档
├── LICENSE                      # GPL-3.0 开源协议
├── go.mod                       # Go依赖管理
├── go.sum                       # Go依赖校验
│
├── cmd/                         # 命令行程序
│   └── xiaoniao/               # xiaoniao CLI
│       ├── main.go             # 程序入口（v1.4 - 统一run模式，集成快捷键）
│       ├── config_ui.go        # TUI配置界面（v1.4 - 修复prompt管理）
│       ├── system_hotkey.go    # 系统快捷键集成（v1.3新增）
│       ├── api_config_ui.go    # API配置界面（支持副模型选择）
│       ├── prompt_test_ui.go   # Prompt测试界面
│       ├── prompts.go          # Prompt桥接层
│       └── lang.go             # i18n多语言支持
│
├── internal/                    # 内部包
│   ├── translator/             # 翻译核心
│   │   ├── translator.go      # 翻译器主逻辑（无缓存）
│   │   ├── provider.go         # Provider接口定义
│   │   ├── provider_registry.go # Provider注册表
│   │   ├── providers_2025.go  # 2025年Provider配置
│   │   ├── openai_compatible.go # OpenAI兼容Provider
│   │   ├── openrouter.go      # OpenRouter专用实现（含n=1参数优化）
│   │   ├── groq_provider.go   # Groq专用实现
│   │   ├── together_provider.go # Together专用实现
│   │   ├── base_prompt.go     # 底层系统prompt模板（无语言限制）
│   │   └── user_prompts.go    # 用户prompt管理
│   │
│   ├── clipboard/              # 剪贴板功能
│   │   ├── monitor.go         # 剪贴板监控
│   │   └── clipboard.go       # 剪贴板操作
│   │
│   ├── tray/                   # 系统托盘（v1.1 优化）
│   │   └── tray.go            # 托盘管理器（简化菜单+红鸟图标）
│   │
│   ├── sound/                  # 音效系统
│   │   └── sound.go           # 提示音播放
│   │
│   └── config/                 # 配置管理
│       └── config.go          # 配置持久化
│
├── frontend/                    # Web UI (Wails版本)
│   ├── index.html              # 主页面
│   ├── src/
│   │   ├── main.js            # 前端入口
│   │   ├── style.css          # 基础样式
│   │   └── pixel-ui.css       # UI组件样式
│   └── package.json           # 前端依赖
│
├── build/                       # 构建相关
│   ├── linux/                  # Linux构建脚本
│   ├── windows/                # Windows构建脚本
│   └── appimage/               # AppImage打包
│
├── docs/                        # 文档
│   ├── INSTALL.md              # 安装指南
│   ├── BUILD.md                # 构建指南
│   ├── PROVIDERS.md            # Provider详细说明
│   ├── PROMPTS.md              # Prompt使用指南
│   └── API.md                  # API文档
│
└── .github/                     # GitHub相关
    └── workflows/              # CI/CD工作流
        └── release.yml         # 自动发布
```

## 🔧 配置说明

### 主配置
位置：`~/.config/xiaoniao/config.json`

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
位置：`~/.config/xiaoniao/prompts.json`

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

- **二进制大小**: ~7.8MB（含TUI框架）
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
   - 确保有剪贴板访问权限
   - 某些Wayland环境可能需要额外配置

## 🤝 贡献

欢迎贡献代码、报告问题或提出建议！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 📝 更新日志

### v1.4 (2025-09-05)
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

### v1.3.1 (2025-09-04)
- 🔧 **代码优化**
  - 移除鼠标点击支持，专注键盘操作
  - 优化二进制大小（使用-ldflags="-s -w"）
  - 清理关于页面的多余信息
  - 修复ptyxis终端打开配置界面问题
  
- 🐛 **问题修复**
  - 修复API配置页面第4个选项无法选择
  - 修复从托盘打开配置界面闪退
  - 移除终端窗口大小设置
  
### v1.3 (2025-09-03)
- ⌨️ **智能快捷键系统**
  - 实时按键检测，无需手动输入
  - 自动检测桌面环境（GNOME/KDE/XFCE等）
  - 智能冲突检查，避免系统快捷键冲突
  - 一键配置系统快捷键，自动创建控制脚本
  
- 🎨 **UI优化**
  - 托盘菜单显示当前翻译风格
  - 优化通知系统，减少不必要的弹窗
  - 改进状态显示方式
  
- 🔧 **代码优化**
  - 清理测试文件和临时文件
  - 改进系统集成方式
  - 支持更多桌面环境

### v1.4.1 (2025-09-06)
- 🌍 **完整国际化支持**
  - 支持9种语言界面（中/英/日/韩/西/法/德/俄/阿拉伯）
  - 修复所有硬编码文本
  - 统一二进制文件管理
  - 完善所有语言的翻译字段（350+字段）
  
- 🔧 **架构优化**
  - 移除冗余的 `xiaoniao tray` 命令
  - 统一到 `xiaoniao run` 支持所有功能
  - 清理过时的二进制文件
  - 优化项目结构

### v1.2 (2025-09-02)
- 🔑 **全局快捷键支持**
  - 使用 golang.design/x/hotkey 实现原生快捷键
  - 支持快捷键实时检测和录制
  - 监控开关快捷键（可自定义）
  - 切换Prompt快捷键（可自定义）
  
- 🛠️ **功能优化**
  - 改进快捷键设置界面（自动检测按键）
  - 优化启动脚本代理检测
  - 修复托盘菜单崩溃问题
  - 改进版本号管理

### v1.1 (2025-09-02)
- 🎆 **新增功能**
  - 11种界面主题（Tokyo Night、Catppuccin、Nord、Dracula等）
  - 快捷键设置功能（监控开关、切换Prompt）
  - 关于页面（通过托盘菜单或配置界面访问）
  - 桌面图标后台启动模式（无终端窗口）
  - 托盘菜单显示/隐藏终端功能
  - 托盘图标状态显示（蓝鸟正常/红鸟停止）
  - Prompt快速切换（托盘菜单循环切换）
  - 日语翻译prompt（朋友对话风格）
  
- 🔧 **UI优化**
  - 统一项目名称为"xiaoniao"
  - 简化托盘菜单（移除多余符号和emoji）
  - 使用checkbox样式的开关选项
  - 主题颜色完整覆盖（包括边框）
  - 1秒通知显示时长（之前3秒太长）
  
- 🐛 **修复问题**
  - 修复日语prompt翻译方向错误
  - 修复prompt内容长度显示错误
  - 移除底层prompt语言限制
  - 优化API参数防止多版本生成（n=1, top_p=0.9）
  - 修复终端自动隐藏功能

### v1.0.0 (2025-08-31)
- 🎉 **正式发布版**
  - 支持20+ AI Provider
  - 300+模型通过OpenRouter
  - 36种预设翻译风格
  - 自定义Prompt支持


## 📄 开源协议

本项目采用 [GPL-3.0](LICENSE) 协议开源

---

更新日期: 2025-09-05

<div align="center">
  <p>如果觉得好用，请给个 ⭐ Star 支持一下！</p>
  <p>Made with ❤️ by Lyrica</p>
</div>