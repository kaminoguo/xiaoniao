#!/bin/bash

# xiaoniao Advanced Windows Build Script
# 支持多种资源生成方法的高级构建脚本

set -e

VERSION="1.6.4"
BUILD_DATE=$(date +%Y%m%d)

echo "========================================="
echo "xiaoniao 高级 Windows 构建 v$VERSION"
echo "  支持终极隐藏控制台模式"
echo "========================================="

# 清理旧构建
echo "→ 清理旧构建..."
rm -rf dist/
mkdir -p dist

# 检查所需工具
echo "→ 检查构建工具..."
TOOLS_OK=true

if [ ! -f ~/go/bin/rsrc ]; then
    echo "  ❌ rsrc 工具未找到，正在安装..."
    go install github.com/akavel/rsrc@latest
    if [ $? -eq 0 ]; then
        echo "  ✓ rsrc 工具安装成功"
    else
        echo "  ❌ rsrc 工具安装失败"
        TOOLS_OK=false
    fi
else
    echo "  ✓ rsrc 工具已找到"
fi

if [ ! -f ~/go/bin/goversioninfo ]; then
    echo "  ❌ goversioninfo 工具未找到，正在安装..."
    go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
    if [ $? -eq 0 ]; then
        echo "  ✓ goversioninfo 工具安装成功"
    else
        echo "  ❌ goversioninfo 工具安装失败"
        TOOLS_OK=false
    fi
else
    echo "  ✓ goversioninfo 工具已找到"
fi

if [ "$TOOLS_OK" != true ]; then
    echo "❌ 工具检查失败，退出构建"
    exit 1
fi

# 更新版本信息
echo "→ 更新版本信息..."
sed -i "s/\"Major\": [0-9]*,/\"Major\": 1,/g" versioninfo.json
sed -i "s/\"Minor\": [0-9]*,/\"Minor\": 6,/g" versioninfo.json
sed -i "s/\"Patch\": [0-9]*,/\"Patch\": 4,/g" versioninfo.json
sed -i "s/\"FileVersion\": \"[^\"]*\"/\"FileVersion\": \"$VERSION.0\"/g" versioninfo.json
sed -i "s/\"ProductVersion\": \"[^\"]*\"/\"ProductVersion\": \"$VERSION\"/g" versioninfo.json

# 验证图标文件
echo "→ 验证图标文件..."
if [ ! -f assets/icon.ico ]; then
    echo "❌ 图标文件 assets/icon.ico 不存在"
    exit 1
fi

ICON_SIZE=$(ls -l assets/icon.ico | awk '{print $5}')
echo "  ✓ 图标文件大小: $((ICON_SIZE / 1024))KB"

# 方法1: 使用 rsrc 生成图标资源（推荐）
echo "→ 方法1: 使用 rsrc 生成图标资源..."
~/go/bin/rsrc -ico assets/icon.ico -o cmd/xiaoniao/resource.syso

if [ -f cmd/xiaoniao/resource.syso ]; then
    RSRC_SIZE=$(ls -l cmd/xiaoniao/resource.syso | awk '{print $5}')
    echo "  ✓ 资源文件已生成，大小: $((RSRC_SIZE / 1024))KB"
    
    # 构建可执行文件（控制台程序）
    echo "→ 构建 Windows 可执行文件（控制台程序）..."
    GOOS=windows GOARCH=amd64 go build \
        -tags cross_compile \
        -ldflags="-s -w -X main.version=$VERSION" \
        -o dist/xiaoniao.exe \
        ./cmd/xiaoniao
    
    BUILD_SUCCESS=true
else
    echo "  ❌ rsrc 方法失败"
    BUILD_SUCCESS=false
fi

# 如果方法1失败，尝试方法2
if [ "$BUILD_SUCCESS" != true ]; then
    echo "→ 方法2: 使用 goversioninfo 生成资源..."
    rm -f cmd/xiaoniao/resource.syso
    
    ~/go/bin/goversioninfo -64 -icon=assets/icon.ico
    
    if [ -f resource.syso ]; then
        mv resource.syso cmd/xiaoniao/
        echo "  ✓ goversioninfo 资源文件已移动到 cmd/xiaoniao/"
        
        # 构建可执行文件（控制台程序）
        echo "→ 构建 Windows 可执行文件（控制台程序）..."
        GOOS=windows GOARCH=amd64 go build \
            -tags cross_compile \
            -ldflags="-s -w -X main.version=$VERSION" \
            -o dist/xiaoniao.exe \
            ./cmd/xiaoniao
        
        BUILD_SUCCESS=true
    else
        echo "  ❌ goversioninfo 方法也失败"
        BUILD_SUCCESS=false
    fi
fi

# 保留资源文件以便后续手动构建
echo "→ 保留资源文件 cmd/xiaoniao/resource.syso 以便后续使用"

# 检查构建结果
if [ "$BUILD_SUCCESS" = true ] && [ -f dist/xiaoniao.exe ]; then
    SIZE=$(ls -lh dist/xiaoniao.exe | awk '{print $5}')
    echo "✓ 构建成功! 大小: $SIZE"
    
    # 验证资源段
    echo "→ 验证资源段..."
    if objdump -h dist/xiaoniao.exe | grep -q -i rsrc; then
        RSRC_HEX=$(objdump -h dist/xiaoniao.exe | grep -i rsrc | awk '{print $3}')
        RSRC_SIZE=$((0x$RSRC_HEX))
        echo "  ✓ 找到 .rsrc 段，大小: $((RSRC_SIZE / 1024))KB"
    else
        echo "  ⚠ 未找到 .rsrc 段"
    fi
    
    # 创建发布包
    echo "→ 创建发布包..."
    cd dist
    zip -q -9 "xiaoniao-windows-v${VERSION}.zip" xiaoniao.exe
    cd ..
    
    echo ""
    echo "========================================="
    echo "✅ 构建完成!"
    echo "========================================="
    echo "📦 输出文件:"
    echo "  - dist/xiaoniao.exe"
    echo "  - dist/xiaoniao-windows-v${VERSION}.zip"
    echo ""
    echo "🔧 构建特性:"
    echo "  ✅ 控制台程序模式（默认显示控制台窗口）"
    echo "  ✅ 支持调试控制台显示/隐藏功能"
    echo "  ✅ 系统托盘运行"
    echo ""
    echo "📋 下一步:"
    echo "  1. 运行 ./verify-icon.sh 验证图标"
    echo "  2. 复制到 Windows 系统测试"
    echo "  3. 确认程序以控制台模式运行（可显示/隐藏控制台窗口）"
    echo ""
else
    echo ""
    echo "========================================="
    echo "❌ 构建失败!"
    echo "========================================="
    echo "请检查错误信息并重试"
    exit 1
fi