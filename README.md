# xiaoniao v1.0

Windows剪贴板翻译工具。监控剪贴板，自动翻译，自动粘贴。

🎉 **v1.0 重大更新**
- ✅ 完整国际化支持（7种语言）
- ✅ 修复非中文Windows系统乱码问题
- ✅ 支持英语、简体中文、繁体中文、日语、韩语、西班牙语、法语
- ✅ 所有界面文字和提示信息完整翻译

## 安装

下载 [xiaoniao.exe](https://github.com/kaminoguo/xiaoniao/releases/latest) 并运行。

Windows SmartScreen警告：点击"更多信息" → "仍要运行"

首次运行需配置API密钥。

## 使用

1. 运行 xiaoniao.exe
2. 复制文本（Ctrl+C）
3. 自动翻译并替换剪贴板
4. 粘贴（Ctrl+V）得到译文

系统托盘图标：
- 蓝色：监控中
- 绿色：翻译中
- 红色：已停止

## 配置

```cmd
xiaoniao.exe config
```

配置文件：`%APPDATA%\xiaoniao\`

支持OpenAI、Anthropic、Google、DeepSeek等API。

## 构建

```bash
# 基础构建
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go

# 带图标
go generate
go build -ldflags="-s -w" -o xiaoniao.exe cmd/xiaoniao/*.go
```

要求：Go 1.21+, Windows

## License

MIT

## Support

如果觉得有用：

[![Ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/xiaoniao)

<details>
<summary>微信赞赏</summary>

[微信赞赏码占位 - 请添加图片]

</details>