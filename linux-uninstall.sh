#!/bin/bash
# xiaoniao Linux 一键卸载脚本

set -e

echo "🐦 xiaoniao 一键卸载"
echo "===================="

# 检测桌面环境
detect_desktop_environment() {
    if [ -z "$DISPLAY" ] && [ -z "$WAYLAND_DISPLAY" ]; then
        echo "none"
        return
    fi
    
    if [ "$XDG_CURRENT_DESKTOP" ]; then
        case "$XDG_CURRENT_DESKTOP" in
            *GNOME*) echo "gnome" ;;
            *KDE*|*Plasma*) echo "kde" ;;
            *XFCE*) echo "xfce" ;;
            *Cinnamon*) echo "cinnamon" ;;
            *MATE*) echo "mate" ;;
            *LXDE*|*LXQT*) echo "lxde" ;;
            *Hyprland*) echo "hyprland" ;;
            *sway*) echo "sway" ;;
            *i3*) echo "i3" ;;
            *) echo "unknown" ;;
        esac
    elif [ "$DESKTOP_SESSION" ]; then
        case "$DESKTOP_SESSION" in
            gnome*) echo "gnome" ;;
            kde*|plasma*) echo "kde" ;;
            xfce*) echo "xfce" ;;
            cinnamon*) echo "cinnamon" ;;
            mate*) echo "mate" ;;
            lxde*|lxqt*) echo "lxde" ;;
            hyprland*) echo "hyprland" ;;
            sway*) echo "sway" ;;
            i3*) echo "i3" ;;
            *) echo "unknown" ;;
        esac
    else
        echo "unknown"
    fi
}

# 检测系统语言
detect_system_language() {
    local lang="${LANG:-en_US}"
    case "${lang:0:2}" in
        zh) echo "中文" ;;
        en) echo "English" ;;
        ja) echo "日本語" ;;
        ko) echo "한국어" ;;
        es) echo "Español" ;;
        fr) echo "Français" ;;
        de) echo "Deutsch" ;;
        ru) echo "Русский" ;;
        ar) echo "العربية" ;;
        *) echo "English" ;;
    esac
}

DESKTOP_ENV=$(detect_desktop_environment)
SYSTEM_LANG=$(detect_system_language)

echo ""
echo "📊 系统信息："
echo "  • 系统语言: $SYSTEM_LANG"
echo "  • 桌面环境: $DESKTOP_ENV"
echo ""

# 检查是否安装
if ! command -v xiaoniao &> /dev/null; then
    echo "⚠️  xiaoniao 未安装"
    exit 0
fi

# 显示当前版本
echo "当前版本: $(xiaoniao --version 2>/dev/null || echo '未知')"
echo ""

# 确认卸载
echo "⚠️  即将卸载以下内容:"
echo "  • 程序文件: /usr/local/bin/xiaoniao"
if [ -d ~/.config/xiaoniao ]; then
    echo "  • 配置文件: ~/.config/xiaoniao/"
fi
if [ -f ~/.local/share/applications/xiaoniao.desktop ]; then
    echo "  • 应用快捷方式: ~/.local/share/applications/xiaoniao.desktop"
fi
if [ -f ~/Desktop/xiaoniao.desktop ]; then
    echo "  • 桌面快捷方式: ~/Desktop/xiaoniao.desktop"
fi
if [ -f ~/桌面/xiaoniao.desktop ]; then
    echo "  • 桌面快捷方式: ~/桌面/xiaoniao.desktop"
fi

echo ""
read -p "确认卸载？(y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "已取消"
    exit 0
fi

# 停止运行中的进程
echo ""
echo "🛑 停止运行中的进程..."
if pgrep -x "xiaoniao" > /dev/null; then
    pkill -x "xiaoniao" || true
    echo "  ✓ 已停止 xiaoniao 进程"
    sleep 1
else
    echo "  • 没有运行中的进程"
fi

# 删除程序文件
echo "🗑️ 删除程序文件..."
if [ -f /usr/local/bin/xiaoniao ]; then
    sudo rm -f /usr/local/bin/xiaoniao
    echo "  ✓ 已删除程序文件"
fi

# 删除快捷方式
echo "🗑️ 删除快捷方式..."
rm -f ~/.local/share/applications/xiaoniao.desktop 2>/dev/null || true
rm -f ~/Desktop/xiaoniao.desktop 2>/dev/null || true
rm -f ~/桌面/xiaoniao.desktop 2>/dev/null || true
echo "  ✓ 已删除快捷方式"

# 根据桌面环境执行额外清理
case "$DESKTOP_ENV" in
    gnome)
        # 刷新 GNOME 应用列表
        update-desktop-database ~/.local/share/applications 2>/dev/null || true
        ;;
    kde)
        # 刷新 KDE 应用缓存
        kbuildsycoca5 2>/dev/null || true
        ;;
esac

# 询问是否删除配置文件
if [ -d ~/.config/xiaoniao ]; then
    echo ""
    read -p "是否删除配置文件？(y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -rf ~/.config/xiaoniao
        echo "  ✓ 已删除配置文件"
    else
        echo "  • 保留配置文件: ~/.config/xiaoniao"
    fi
fi

echo ""
echo "✅ 卸载完成！"
echo ""
echo "感谢使用 xiaoniao"
echo "如需重新安装，请访问: https://github.com/kaminoguo/xiaoniao"
echo ""