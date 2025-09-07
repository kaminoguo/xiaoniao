#!/bin/bash
# 自动编译和更新 xiaoniao 到所有位置

cd /home/lyrica/xiaoniao

# 编译
echo "正在编译..."
go build -o xiaoniao_new ./cmd/xiaoniao || exit 1

# 停止运行中的进程
echo "停止旧进程..."
pkill xiaoniao 2>/dev/null
sleep 1

# 更新到所有位置
echo "更新到 ~/.local/bin/..."
cp xiaoniao_new ~/.local/bin/xiaoniao

echo "更新到 /usr/local/bin/... (需要sudo权限)"
sudo cp xiaoniao_new /usr/local/bin/xiaoniao 2>/dev/null || echo "跳过 /usr/local/bin (需要sudo权限)"

# 清理
rm xiaoniao_new

echo "✅ 小鸟翻译已更新到所有位置！"
echo "   ~/.local/bin/xiaoniao"
echo "   /usr/local/bin/xiaoniao"