#!/bin/bash
# xiaoniao Linux 一键安装脚本

set -e

echo "xiaoniao Installation Script"
echo "============================"

# 检测系统语言
detect_system_language() {
    local lang="${LANG:-en_US}"
    
    # 提取语言代码 - 仅支持实际实现的7种语言
    case "${lang:0:2}" in
        zh)
            if [[ "$lang" == *"TW"* ]] || [[ "$lang" == *"HK"* ]]; then
                echo "zh_TW"  # 繁体中文
            else
                echo "zh_CN"  # 简体中文
            fi
            ;;
        en) echo "en" ;;    # English
        ja) echo "ja" ;;    # 日本語
        ko) echo "ko" ;;    # 한국어
        es) echo "es" ;;    # Español
        fr) echo "fr" ;;    # Français
        *) echo "en" ;;     # 默认英文（不支持的语言）
    esac
}

# 检测桌面环境
detect_desktop_environment() {
    # 首先检查是否在图形环境中
    if [ -z "$DISPLAY" ] && [ -z "$WAYLAND_DISPLAY" ]; then
        echo "none"
        return
    fi
    
    # 检测具体的桌面环境
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

# 检测终端类型
detect_terminal() {
    # 检查常见的终端环境变量
    if [ "$TERM_PROGRAM" ]; then
        echo "$TERM_PROGRAM"
    elif [ "$TERMINAL_EMULATOR" ]; then
        echo "$TERMINAL_EMULATOR"
    elif [ "$VTE_VERSION" ]; then
        echo "vte-based"  # GNOME Terminal, Tilix等
    elif [ "$KONSOLE_VERSION" ]; then
        echo "konsole"
    elif [ "$ALACRITTY_SOCKET" ]; then
        echo "alacritty"
    elif [ "$KITTY_WINDOW_ID" ]; then
        echo "kitty"
    elif [ "$WEZTERM_EXECUTABLE" ]; then
        echo "wezterm"
    else
        # 尝试从进程树中检测
        local parent_pid=$PPID
        local parent_cmd=$(ps -p $parent_pid -o comm= 2>/dev/null)
        case "$parent_cmd" in
            gnome-terminal*) echo "gnome-terminal" ;;
            konsole*) echo "konsole" ;;
            xfce4-terminal*) echo "xfce4-terminal" ;;
            terminator*) echo "terminator" ;;
            alacritty*) echo "alacritty" ;;
            kitty*) echo "kitty" ;;
            *) echo "$TERM" ;;
        esac
    fi
}

# 显示检测信息
SYSTEM_LANG=$(detect_system_language)
DESKTOP_ENV=$(detect_desktop_environment)
TERMINAL_TYPE=$(detect_terminal)

echo ""
echo "System Detection:"
echo "  Language: $SYSTEM_LANG"
echo "  Desktop: $DESKTOP_ENV"
echo "  Terminal: $TERMINAL_TYPE"
echo ""

# 检查是否已安装
if command -v xiaoniao &> /dev/null; then
    echo "xiaoniao is already installed, version: $(xiaoniao --version)"
    read -p "Reinstall? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 0
    fi
fi

# 下载并安装
echo "Downloading xiaoniao..."
wget -q --show-progress -O /tmp/xiaoniao https://github.com/kaminoguo/xiaoniao/releases/latest/download/xiaoniao-linux-amd64

# 检查下载是否成功
if [ ! -f /tmp/xiaoniao ]; then
    echo "Download failed"
    exit 1
fi

echo "Installing..."
sudo mv /tmp/xiaoniao /usr/local/bin/xiaoniao
sudo chmod +x /usr/local/bin/xiaoniao

# 创建初始配置文件（设置系统语言）
echo "Creating initial configuration..."
mkdir -p ~/.config/xiaoniao

# 如果配置文件不存在，创建默认配置
if [ ! -f ~/.config/xiaoniao/config.json ]; then
    # 将系统语言映射到配置语言 - 仅支持7种语言
    CONFIG_LANG="en"  # 默认英文
    case "$SYSTEM_LANG" in
        zh_CN) CONFIG_LANG="zh-CN" ;;
        zh_TW) CONFIG_LANG="zh-TW" ;;
        en) CONFIG_LANG="en" ;;
        ja) CONFIG_LANG="ja" ;;
        ko) CONFIG_LANG="ko" ;;
        es) CONFIG_LANG="es" ;;
        fr) CONFIG_LANG="fr" ;;
        *) CONFIG_LANG="en" ;;  # 其他语言默认英文
    esac
    
    cat > ~/.config/xiaoniao/config.json << EOF
{
    "language": "$CONFIG_LANG",
    "theme": "tokyo-night",
    "auto_paste": true,
    "prompt_id": "direct"
}
EOF
    echo "  Interface language set to: $CONFIG_LANG"
fi

# 根据桌面环境创建快捷方式
if [ "$DESKTOP_ENV" != "none" ]; then
    echo "Creating desktop shortcut..."
    mkdir -p ~/.local/share/applications
    
    # 创建 .desktop 文件
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
    
    # 复制到桌面（如果桌面存在）
    if [ -d ~/Desktop ]; then
        cp ~/.local/share/applications/xiaoniao.desktop ~/Desktop/ 2>/dev/null || true
        chmod +x ~/Desktop/xiaoniao.desktop 2>/dev/null || true
    elif [ -d ~/桌面 ]; then
        cp ~/.local/share/applications/xiaoniao.desktop ~/桌面/ 2>/dev/null || true
        chmod +x ~/桌面/xiaoniao.desktop 2>/dev/null || true
    fi
    
    # 特定桌面环境的额外配置
    case "$DESKTOP_ENV" in
        gnome)
            echo "  Detected GNOME desktop"
            # 刷新 GNOME 应用列表
            update-desktop-database ~/.local/share/applications 2>/dev/null || true
            ;;
        kde)
            echo "  Detected KDE Plasma desktop"
            # KDE 特定配置
            kbuildsycoca5 2>/dev/null || true
            ;;
        xfce)
            echo "  Detected XFCE desktop"
            ;;
        hyprland|sway|i3)
            echo "  Detected tiling window manager: $DESKTOP_ENV"
            echo "  Recommend using hotkeys to launch: xiaoniao run"
            ;;
    esac
fi

# 检测并配置终端特性
echo ""
echo "Terminal Configuration:"
case "$TERMINAL_TYPE" in
    gnome-terminal|vte-based)
        echo "  Full TUI support"
        echo "  256 colors and Unicode support"
        ;;
    konsole)
        echo "  KDE Konsole detected"
        echo "  Full TUI support"
        ;;
    alacritty|kitty|wezterm)
        echo "  Modern terminal detected: $TERMINAL_TYPE"
        echo "  Excellent performance and rendering"
        ;;
    xterm*)
        echo "  Basic terminal, may need TERM=xterm-256color"
        ;;
    *)
        echo "  Terminal type: $TERMINAL_TYPE"
        ;;
esac

echo ""
echo "Installation complete!"
echo ""
echo "Usage:"

if [ "$DESKTOP_ENV" != "none" ]; then
    echo "  1. Click the xiaoniao icon on desktop or application menu"
    echo "  2. Or run in terminal: xiaoniao run"
else
    echo "  Configure API: xiaoniao config"
    echo "  Start monitoring: xiaoniao run"
fi

echo ""
echo "First time setup:"
echo "  Run 'xiaoniao config' to configure API Key"
echo "  Interface language automatically set to: $SYSTEM_LANG"
echo ""

# 询问是否立即配置
read -p "Open configuration now? (Y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Nn]$ ]]; then
    xiaoniao config
fi