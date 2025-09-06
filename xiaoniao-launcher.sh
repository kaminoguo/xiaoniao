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
