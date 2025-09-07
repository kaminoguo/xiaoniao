#!/bin/bash

# 更新xiaoniao桌面快捷方式

echo "更新 xiaoniao 桌面快捷方式..."

# 确保使用最新编译的二进制文件
if [ -f "./xiaoniao" ]; then
    echo "找到编译的二进制文件"
else
    echo "未找到二进制文件，开始编译..."
    go build -o xiaoniao ./cmd/xiaoniao
fi

# 更新desktop文件
cat > ~/.local/share/applications/xiaoniao.desktop << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=xiaoniao
Name[en]=xiaoniao
Name[zh_CN]=xiaoniao
Comment=智能剪贴板翻译工具
Comment[en]=Smart clipboard translator
Comment[zh_CN]=智能剪贴板翻译工具
GenericName=Translation Tool
GenericName[zh_CN]=翻译工具
Exec=/home/lyrica/xiaoniao/xiaoniao-launcher.sh
Icon=/home/lyrica/xiaoniao/assets/icon.png
Terminal=false
Categories=Utility;Translation;Office;
Keywords=translate;translation;clipboard;xiaoniao;
StartupNotify=false
StartupWMClass=xiaoniao
Path=/home/lyrica/xiaoniao
EOF

# 更新启动脚本，确保使用最新的二进制文件
cat > /home/lyrica/xiaoniao/xiaoniao-launcher.sh << 'EOF'
#!/bin/bash
# Launch xiaoniao run in background without terminal

# Get the directory of this script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Change to the xiaoniao directory
cd "$DIR"

# 确保使用最新版本
if [ -f "./xiaoniao" ]; then
    # 检查是否需要重新编译
    SOURCE_NEWER=$(find ./cmd ./internal -name "*.go" -newer ./xiaoniao 2>/dev/null | head -1)
    if [ ! -z "$SOURCE_NEWER" ]; then
        echo "检测到源代码更新，重新编译..." > /tmp/xiaoniao.log
        go build -o xiaoniao ./cmd/xiaoniao 2>> /tmp/xiaoniao.log
    fi
fi

# Run xiaoniao in background with output redirected
nohup ./xiaoniao run >> /tmp/xiaoniao.log 2>&1 &

# The xiaoniao process will create its own tray icon
# User can control terminal visibility through the tray menu
EOF

chmod +x /home/lyrica/xiaoniao/xiaoniao-launcher.sh

# 刷新桌面数据库
update-desktop-database ~/.local/share/applications/ 2>/dev/null

echo "更新完成！"
echo ""
echo "你可以："
echo "1. 从应用程序菜单中启动 xiaoniao"
echo "2. 或运行: /home/lyrica/xiaoniao/xiaoniao-launcher.sh"
echo "3. 或直接运行: ./xiaoniao run"