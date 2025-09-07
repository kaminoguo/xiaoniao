#!/bin/bash
# xiaoniao Linux 一键卸载脚本

echo "🗑️ xiaoniao 一键卸载"
echo "===================="

# 检查是否安装
if ! command -v xiaoniao &> /dev/null; then
    echo "⚠️  xiaoniao 未安装"
    exit 0
fi

# 确认卸载
echo "即将卸载 xiaoniao $(xiaoniao --version 2>/dev/null || echo '')"
read -p "确定要卸载吗？(y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "已取消"
    exit 0
fi

# 停止运行中的进程
echo "停止运行中的进程..."
pkill -f "xiaoniao run" 2>/dev/null || true

# 删除程序
echo "删除程序文件..."
sudo rm -f /usr/local/bin/xiaoniao
sudo rm -f ~/.local/bin/xiaoniao 2>/dev/null || true

# 删除配置
read -p "是否删除配置文件？(y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    rm -rf ~/.config/xiaoniao
    echo "✅ 配置文件已删除"
else
    echo "保留配置文件"
fi

# 删除桌面快捷方式
echo "删除快捷方式..."
rm -f ~/.local/share/applications/xiaoniao.desktop
rm -f ~/Desktop/xiaoniao.desktop
rm -f ~/桌面/xiaoniao.desktop 2>/dev/null || true

echo ""
echo "✅ 卸载完成！"
echo ""