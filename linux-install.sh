#!/bin/bash
# xiaoniao Linux 一键安装脚本

set -e

echo "🐦 xiaoniao 一键安装"
echo "===================="

# 检测系统语言
detect_system_language() {
    local lang="${LANG:-en_US}"
    
    # 提取语言代码
    case "${lang:0:2}" in
        zh)
            if [[ "$lang" == *"TW"* ]] || [[ "$lang" == *"HK"* ]]; then
                echo "zh_TW"  # 繁体中文
            else
                echo "zh_CN"  # 简体中文
            fi
            ;;
        en) echo "en" ;;
        ja) echo "ja" ;;
        ko) echo "ko" ;;
        es) echo "es" ;;
        fr) echo "fr" ;;
        de) echo "de" ;;
        ru) echo "ru" ;;
        ar) echo "ar" ;;
        *) echo "en" ;;  # 默认英文
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
echo "📊 系统检测："
echo "  • 系统语言: $SYSTEM_LANG"
echo "  • 桌面环境: $DESKTOP_ENV"
echo "  • 终端类型: $TERMINAL_TYPE"
echo ""

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

# 创建初始配置文件（设置系统语言）
echo "⚙️ 创建初始配置..."
mkdir -p ~/.config/xiaoniao

# 如果配置文件不存在，创建默认配置
if [ ! -f ~/.config/xiaoniao/config.json ]; then
    # 将系统语言映射到配置语言
    CONFIG_LANG="cn"  # 默认中文
    case "$SYSTEM_LANG" in
        en) CONFIG_LANG="en" ;;
        zh_CN) CONFIG_LANG="cn" ;;
        zh_TW) CONFIG_LANG="tw" ;;
        ja) CONFIG_LANG="jp" ;;
        ko) CONFIG_LANG="kr" ;;
        es) CONFIG_LANG="es" ;;
        fr) CONFIG_LANG="fr" ;;
        de) CONFIG_LANG="de" ;;
        ru) CONFIG_LANG="ru" ;;
        ar) CONFIG_LANG="ar" ;;
    esac
    
    cat > ~/.config/xiaoniao/config.json << EOF
{
    "language": "$CONFIG_LANG",
    "theme": "tokyo-night",
    "auto_paste": true,
    "prompt_id": "direct"
}
EOF
    echo "  ✓ 已设置界面语言为: $CONFIG_LANG"
fi

# 根据桌面环境创建快捷方式
if [ "$DESKTOP_ENV" != "none" ]; then
    echo "🖥️ 创建桌面快捷方式..."
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
            echo "  • 检测到 GNOME 桌面"
            # 刷新 GNOME 应用列表
            update-desktop-database ~/.local/share/applications 2>/dev/null || true
            ;;
        kde)
            echo "  • 检测到 KDE Plasma 桌面"
            # KDE 特定配置
            kbuildsycoca5 2>/dev/null || true
            ;;
        xfce)
            echo "  • 检测到 XFCE 桌面"
            ;;
        hyprland|sway|i3)
            echo "  • 检测到平铺式窗口管理器: $DESKTOP_ENV"
            echo "  • 建议使用快捷键启动: xiaoniao run"
            ;;
    esac
fi

# 检测并配置终端特性
echo ""
echo "🖥️ 终端配置："
case "$TERMINAL_TYPE" in
    gnome-terminal|vte-based)
        echo "  • 支持完整的 TUI 界面"
        echo "  • 支持 256 色和 Unicode"
        ;;
    konsole)
        echo "  • KDE Konsole 检测到"
        echo "  • 完美支持 TUI 界面"
        ;;
    alacritty|kitty|wezterm)
        echo "  • 现代终端检测到: $TERMINAL_TYPE"
        echo "  • 优秀的性能和渲染"
        ;;
    xterm*)
        echo "  • 基础终端，可能需要配置 TERM=xterm-256color"
        ;;
    *)
        echo "  • 终端类型: $TERMINAL_TYPE"
        ;;
esac

echo ""
echo "✅ 安装完成！"
echo ""
echo "📌 使用方法："

if [ "$DESKTOP_ENV" != "none" ]; then
    echo "  1. 点击桌面或应用菜单中的【xiaoniao】图标"
    echo "  2. 或在终端运行: xiaoniao run"
else
    echo "  • 配置 API: xiaoniao config"
    echo "  • 启动监控: xiaoniao run"
fi

echo ""
echo "💡 首次使用："
echo "  运行 'xiaoniao config' 配置 API Key"
echo "  界面语言已自动设置为: $SYSTEM_LANG"
echo ""

# 询问是否立即配置
read -p "是否立即打开配置界面？(Y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Nn]$ ]]; then
    xiaoniao config
fi