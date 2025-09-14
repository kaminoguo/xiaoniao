# Changelog

## [1.0.0] - 2025-09-14

### 🎉 Major Release - Complete Internationalization

### Added
- 🌍 完整的国际化（i18n）支持 - 7种语言
  - English (英语) - 372 翻译条目
  - 简体中文 (Simplified Chinese) - 363 翻译条目
  - 繁體中文 (Traditional Chinese) - 382 翻译条目
  - 日本語 (Japanese) - 382 翻译条目
  - 한국어 (Korean) - 382 翻译条目
  - Español (Spanish) - 379 翻译条目
  - Français (French) - 372 翻译条目
- 每种语言都包含完整的教程内容
- 所有界面元素完整翻译，无遗漏

### Fixed
- ✅ 彻底修复了非中文 Windows 系统的乱码问题
- ✅ 修复了所有硬编码的中文字符串（400+ 个字符串）
- ✅ 修复了配置界面、API 配置、快捷键设置等所有界面的国际化
- ✅ 修复了系统托盘菜单的国际化
- ✅ 修复了日志和错误消息的国际化
- ✅ 修复了控制台输出的国际化
- ✅ 修复了教程页面在非中英文语言下显示"请参考主页面"的问题

### Changed
- 所有用户界面文本现在通过 i18n 系统管理
- 程序版本号更新为 v1.0.0，标志着第一个稳定的国际化版本
- 编译后程序大小保持在 8.5MB

### Technical Details
- 系统化检查并补齐了所有语言缺失的字段
- 更新了 `lang_en.go`, `lang_zh_cn.go` 和 `lang_others.go`
- 修复了 `main.go`, `config_ui.go`, `api_config_ui.go`, `tray.go`, `logbuffer.go` 等文件中的硬编码字符串
- 确保所有字符串都通过 `i18n.T()` 获取
- 每种语言都经过完整性检查，确保没有遗漏的翻译条目

## [1.6.7] - 2025-09-13

### Changed
- 重构快捷键配置系统，改用文本输入方式替代键盘钩子录制
- 修复 prompt 内容读取问题，确保使用用户自定义的 prompt 内容
- 优化程序体积，从 12MB 减少到 8.5MB（减少约 29%）
- 源代码从 240KB 减少到 196KB（减少约 18%）

### Fixed
- 修复了 prompt 编辑后不生效的问题
- 解决了 Linux/WSL2 交叉编译时的 CGO 依赖问题
- 移除了无法正常工作的键盘钩子录制功能

### Technical Details
- 移除了所有 gohook 和 keyboard hook 相关代码
- 实现了纯 Go 的快捷键验证和标准化功能
- 支持的快捷键格式：Ctrl+C, Alt+Shift+V, Ctrl+Alt+X 等

## [1.6.6] - 2025-09-13

### Added
- 导出日志功能
- 多彩系统托盘图标状态指示

### Known Issues
- 快捷键录制功能在交叉编译环境下无法正常工作

## [1.6.5] - 2025-09-12

### Added
- 调试控制台功能
- 改进的错误处理和日志记录

## [1.6.4] - 2025-09-11

### Added
- Windows 专用版本发布
- 系统托盘集成
- 全局热键支持

## [1.6.3] - 2025-09-10

### Added
- 多语言界面支持（中文、英语、日语、韩语、西班牙语、法语）
- 自定义翻译提示词功能

## Earlier Versions

详见 GitHub Releases 页面