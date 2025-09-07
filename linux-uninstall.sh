#!/bin/bash
# xiaoniao Linux 一键卸载脚本

echo "🗑️ xiaoniao 一键卸载"
echo "===================="

# 删除程序
echo "删除程序文件..."
sudo rm -f /usr/local/bin/xiaoniao

# 删除配置
read -p "是否删除配置文件？(y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    rm -rf ~/.config/xiaoniao
    echo "✅ 配置文件已删除"
fi

# 删除桌面快捷方式
rm -f ~/.local/share/applications/xiaoniao.desktop
rm -f ~/Desktop/xiaoniao.desktop

echo ""
echo "✅ 卸载完成！"
echo ""