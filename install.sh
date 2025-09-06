#!/bin/bash

# 小鸟翻译 (xiaoniao) 安装脚本

echo "================================"
echo "  小鸟翻译 (xiaoniao) 安装脚本"
echo "================================"

# 构建程序
echo "正在构建..."
go build -o xiaoniao cmd/xiaoniao/main.go

if [ $? -ne 0 ]; then
    echo "构建失败"
    exit 1
fi

# 安装到用户目录
INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

echo "正在安装到 $INSTALL_DIR..."
cp xiaoniao "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/xiaoniao"

# 检查PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo "⚠️  请将以下内容添加到你的 ~/.bashrc 或 ~/.zshrc:"
    echo "    export PATH=\"\$HOME/.local/bin:\$PATH\""
fi

echo ""
echo "✅ 安装完成！"
echo ""
echo "使用方法："
echo "  xiaoniao --help           查看帮助"
echo "  xiaoniao --list           查看所有Prompt"
echo "  xiaoniao --key=YOUR_KEY --daemon  启动监控"
echo ""
echo "或设置环境变量："
echo "  export XIAONIAO_API_KEY=YOUR_API_KEY"
echo "  xiaoniao --daemon"
echo ""