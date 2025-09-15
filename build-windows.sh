#!/bin/bash

# xiaoniao Windows Build Script
# Builds Windows executable from Linux

set -e

VERSION="1.0.0"
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
sed -i "s/\"FileVersion\": \"[0-9.]*\"/\"FileVersion\": \"$VERSION.0\"/g" versioninfo.json
sed -i "s/\"ProductVersion\": \"[0-9.]*\"/\"ProductVersion\": \"$VERSION\"/g" versioninfo.json

# Generate Windows resource file with icon
echo "→ Generating Windows resource file..."

# Use rsrc to generate icon resource (more reliable for icons)
echo "  → Using rsrc to generate icon resource..."
~/go/bin/rsrc -ico assets/icon.ico -o cmd/xiaoniao/resource.syso

# Check if resource was created
if [ -f cmd/xiaoniao/resource.syso ]; then
    echo "  ✓ Resource file with icon created in cmd/xiaoniao/"
else
    echo "  ❌ Failed to generate resource file"
    exit 1
fi

# Build Windows 64-bit executable (Console mode)
echo "→ Building Windows amd64 executable..."
# Removed cross_compile tag to use real Windows API functions
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
zip -q -9 "xiaoniao-v${VERSION}.zip" xiaoniao.exe
cd ..

echo ""
echo "========================================="
echo "✅ Build complete!"
echo "========================================="
echo "📦 Output: dist/xiaoniao-v${VERSION}.zip"
echo ""