#!/bin/bash

# xiaoniao Icon Verification Script
# 验证Windows exe文件是否包含图标资源

echo "========================================="
echo "图标资源验证"
echo "========================================="

EXE_FILE="dist/xiaoniao.exe"

if [ ! -f "$EXE_FILE" ]; then
    echo "❌ 未找到 $EXE_FILE"
    exit 1
fi

echo "→ 检查exe文件..."
file "$EXE_FILE"

echo ""
echo "→ 检查资源段..."
objdump -h "$EXE_FILE" | grep -i rsrc
if [ $? -eq 0 ]; then
    echo "✓ 找到 .rsrc 资源段"
else
    echo "❌ 未找到 .rsrc 资源段"
fi

echo ""
echo "→ 检查资源段大小..."
RSRC_SIZE=$(objdump -h "$EXE_FILE" | grep -i rsrc | awk '{print $3}')
if [ ! -z "$RSRC_SIZE" ]; then
    SIZE_DECIMAL=$((0x$RSRC_SIZE))
    SIZE_KB=$((SIZE_DECIMAL / 1024))
    echo "✓ 资源段大小: ${SIZE_KB}KB (${SIZE_DECIMAL} bytes)"
    if [ $SIZE_DECIMAL -gt 50000 ]; then
        echo "✓ 资源段足够大，可能包含图标"
    else
        echo "⚠ 资源段较小，可能缺少图标"
    fi
else
    echo "❌ 无法获取资源段大小"
fi

echo ""
echo "→ 检查原始图标文件..."
if [ -f "assets/icon.ico" ]; then
    ICON_SIZE=$(ls -l assets/icon.ico | awk '{print $5}')
    echo "✓ 原始图标文件大小: $((ICON_SIZE / 1024))KB (${ICON_SIZE} bytes)"
    file assets/icon.ico
else
    echo "❌ 未找到原始图标文件 assets/icon.ico"
fi

echo ""
echo "========================================="
echo "验证完成"
echo "========================================="
echo "📋 建议："
echo "1. 将 $EXE_FILE 复制到Windows系统"
echo "2. 在文件管理器中查看是否显示图标"
echo "3. 如果仍无图标，可能需要："
echo "   - 检查Windows系统缓存"
echo "   - 重启Windows资源管理器"
echo "   - 确保图标格式兼容"