#!/bin/bash
# xiaoniao macOS 启动脚本

# 切换到脚本所在目录
cd "$(dirname "$0")"

# 检查是否有配置文件
CONFIG_DIR="$HOME/Library/Application Support/xiaoniao"
CONFIG_FILE="$CONFIG_DIR/config.json"

# 确保配置目录存在
mkdir -p "$CONFIG_DIR"

# 如果没有配置文件，先打开配置界面
if [ ! -f "$CONFIG_FILE" ]; then
    echo "🐦 欢迎使用 xiaoniao"
    echo "===================="
    echo ""
    echo "首次运行，需要配置 API Key"
    echo ""
    ./xiaoniao config
    echo ""
    echo "配置完成后，将自动启动..."
    sleep 2
fi

# 启动主程序
echo "🐦 启动 xiaoniao..."
./xiaoniao run