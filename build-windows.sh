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

# Update versioninfo.json version
echo "→ Updating version info..."
sed -i "s/\"Major\": 1,/\"Major\": 1,/g; s/\"Minor\": 6,/\"Minor\": 6,/g; s/\"Patch\": 1,/\"Patch\": 4,/g" versioninfo.json
sed -i "s/\"FileVersion\": \"1.6.1.0\"/\"FileVersion\": \"$VERSION.0\"/g" versioninfo.json
sed -i "s/\"ProductVersion\": \"1.6.1\"/\"ProductVersion\": \"$VERSION\"/g" versioninfo.json

# Generate version info resource file (includes icon)
echo "→ Generating Windows resource file..."
~/go/bin/goversioninfo -64

# Move resource file to correct location
if [ -f resource.syso ]; then
    mv resource.syso cmd/xiaoniao/
    echo "  ✓ Resource file moved to cmd/xiaoniao/"
fi

# Build Windows 64-bit executable (console app for TUI)
# NOTE: NOT using -H windowsgui flag to keep console window for TUI
echo "→ Building Windows amd64 executable..."
GOOS=windows GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=$VERSION" \
    -o dist/xiaoniao.exe \
    ./cmd/xiaoniao

# Clean up resource file
rm -f cmd/xiaoniao/resource.syso

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