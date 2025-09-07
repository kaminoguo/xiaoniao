#!/bin/bash
# xiaoniao Linux 一键安装脚本

set -e

echo "🐦 xiaoniao 一键安装"
echo "===================="

# 检查是否已安装
if command -v xiaoniao &> /dev/null; then
    echo "⚠️  xiaoniao 已安装，版本: $(xiaoniao --version)"
    read -p "是否重新安装？(y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 0
    fi
fi

# 下载并安装
echo "📥 正在下载..."
wget -q --show-progress -O /tmp/xiaoniao https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-linux-amd64

# 检查下载是否成功
if [ ! -f /tmp/xiaoniao ]; then
    echo "❌ 下载失败"
    exit 1
fi

echo "📦 正在安装..."
sudo mv /tmp/xiaoniao /usr/local/bin/xiaoniao
sudo chmod +x /usr/local/bin/xiaoniao

# 创建桌面快捷方式
echo "🖥️ 创建桌面快捷方式..."
mkdir -p ~/.local/share/applications

cat > ~/.local/share/applications/xiaoniao.desktop << 'EOF'
[Desktop Entry]
Version=1.0
Type=Application
Name=xiaoniao
Comment=智能剪贴板翻译
Exec=xiaoniao run
Terminal=false
Categories=Utility;
StartupNotify=false
EOF

# 复制到桌面
cp ~/.local/share/applications/xiaoniao.desktop ~/Desktop/ 2>/dev/null || true
chmod +x ~/Desktop/xiaoniao.desktop 2>/dev/null || true

echo ""
echo "✅ 安装完成！"
echo ""
echo "使用方法："
echo "  1. 点击桌面的【xiaoniao】图标"
echo "  2. 或在终端运行: xiaoniao config"
echo ""