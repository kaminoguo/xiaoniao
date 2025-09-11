#!/bin/bash

# xiaoniao Windows Build Script
# Builds Windows executable from Linux

set -e

VERSION="1.6.4"
BUILD_DATE=$(date +%Y%m%d)

echo "========================================="
echo "Building xiaoniao for Windows v$VERSION"
echo "========================================="

# Clean old builds
echo "→ Cleaning old builds..."
rm -rf dist/
mkdir -p dist

# Build Windows 64-bit executable (console app for TUI)
echo "→ Building Windows amd64 executable..."
GOOS=windows GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=$VERSION" \
    -o dist/xiaoniao.exe \
    ./cmd/xiaoniao

# Check if build succeeded
if [ ! -f dist/xiaoniao.exe ]; then
    echo "❌ Build failed!"
    exit 1
fi

# Get file size
SIZE=$(ls -lh dist/xiaoniao.exe | awk '{print $5}')
echo "✓ Build successful! Size: $SIZE"

# Create release package
echo "→ Creating release package..."
cd dist
zip -q -9 "xiaoniao-windows-v${VERSION}.zip" xiaoniao.exe
cd ..

echo ""
echo "========================================="
echo "✅ Build complete!"
echo "========================================="
echo "📦 Output: dist/xiaoniao-windows-v${VERSION}.zip"
echo ""