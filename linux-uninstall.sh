#!/bin/bash
# xiaoniao Linux 一键卸载脚本

set -e

echo "xiaoniao Uninstall Script"
echo "========================"

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
        zh)
            if [[ "$lang" == *"TW"* ]] || [[ "$lang" == *"HK"* ]]; then
                echo "繁體中文"
            else
                echo "简体中文"
            fi
            ;;
        en) echo "English" ;;
        ja) echo "日本語" ;;
        ko) echo "한국어" ;;
        es) echo "Español" ;;
        fr) echo "Français" ;;
        *) echo "English" ;;  # Default to English for unsupported languages
    esac
}

DESKTOP_ENV=$(detect_desktop_environment)
SYSTEM_LANG=$(detect_system_language)

echo ""
echo "System Information:"
echo "  Language: $SYSTEM_LANG"
echo "  Desktop: $DESKTOP_ENV"
echo ""

# 检查是否安装
if ! command -v xiaoniao &> /dev/null; then
    echo "xiaoniao is not installed"
    exit 0
fi

# 显示当前版本
echo "Current version: $(xiaoniao --version 2>/dev/null || echo 'unknown')"
echo ""

# 确认卸载
echo "Will uninstall the following:"
echo "  Program file: /usr/local/bin/xiaoniao"
if [ -d ~/.config/xiaoniao ]; then
    echo "  Configuration: ~/.config/xiaoniao/"
fi
if [ -f ~/.local/share/applications/xiaoniao.desktop ]; then
    echo "  Application shortcut: ~/.local/share/applications/xiaoniao.desktop"
fi
if [ -f ~/Desktop/xiaoniao.desktop ]; then
    echo "  Desktop shortcut: ~/Desktop/xiaoniao.desktop"
fi
if [ -f ~/桌面/xiaoniao.desktop ]; then
    echo "  Desktop shortcut: ~/桌面/xiaoniao.desktop"
fi

echo ""
read -p "Confirm uninstall? (y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Cancelled"
    exit 0
fi

# 停止运行中的进程
echo ""
echo "Stopping running processes..."
if pgrep -x "xiaoniao" > /dev/null; then
    pkill -x "xiaoniao" || true
    echo "  Stopped xiaoniao process"
    sleep 1
else
    echo "  No running processes"
fi

# 删除程序文件
echo "Removing program files..."
if [ -f /usr/local/bin/xiaoniao ]; then
    sudo rm -f /usr/local/bin/xiaoniao
    echo "  Program file removed"
fi

# 删除快捷方式
echo "Removing shortcuts..."
rm -f ~/.local/share/applications/xiaoniao.desktop 2>/dev/null || true
rm -f ~/Desktop/xiaoniao.desktop 2>/dev/null || true
rm -f ~/桌面/xiaoniao.desktop 2>/dev/null || true
echo "  Shortcuts removed"

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
    read -p "Remove configuration files? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -rf ~/.config/xiaoniao
        echo "  Configuration files removed"
    else
        echo "  Configuration files kept: ~/.config/xiaoniao"
    fi
fi

echo ""
echo "Uninstall complete!"
echo ""
echo "Thank you for using xiaoniao"
echo "To reinstall, visit: https://github.com/kaminoguo/xiaoniao"
echo ""