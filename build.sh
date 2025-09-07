#!/bin/bash

echo "🐦 Building xiaoniao for all platforms..."
echo "========================================"
echo ""

# 创建输出目录
mkdir -p dist

# Build for Linux (AMD64)
echo "📦 Building Linux (AMD64) version..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/xiaoniao-linux-amd64 cmd/xiaoniao/*.go
if [ $? -eq 0 ]; then
    echo "  ✓ Linux AMD64 build complete"
else
    echo "  ✗ Linux AMD64 build failed"
fi

# Build for Windows (AMD64)
echo "📦 Building Windows (AMD64) version..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/xiaoniao.exe cmd/xiaoniao/*.go
if [ $? -eq 0 ]; then
    echo "  ✓ Windows AMD64 build complete"
else
    echo "  ✗ Windows AMD64 build failed"
fi

# Build for macOS (Intel)
echo "📦 Building macOS (Intel) version..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/xiaoniao-darwin-amd64 cmd/xiaoniao/*.go
if [ $? -eq 0 ]; then
    echo "  ✓ macOS Intel build complete"
else
    echo "  ✗ macOS Intel build failed"
fi

# Build for macOS (Apple Silicon)
echo "📦 Building macOS (Apple Silicon) version..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o dist/xiaoniao-darwin-arm64 cmd/xiaoniao/*.go
if [ $? -eq 0 ]; then
    echo "  ✓ macOS Apple Silicon build complete"
else
    echo "  ✗ macOS Apple Silicon build failed"
fi

echo ""
echo "📋 Creating distribution packages..."

# Create Linux package
if [ -f dist/xiaoniao-linux-amd64 ]; then
    echo "  • Creating Linux package..."
    cp linux-install.sh dist/
    cp linux-uninstall.sh dist/
    echo "    ✓ Linux package ready"
fi

# Create Windows package
if [ -f dist/xiaoniao.exe ]; then
    echo "  • Creating Windows ZIP..."
    cd dist
    cp ../xiaoniao.bat .
    zip -q xiaoniao-windows.zip xiaoniao.exe xiaoniao.bat
    rm xiaoniao.bat
    cd ..
    echo "    ✓ Windows ZIP created: dist/xiaoniao-windows.zip"
fi

# Create macOS Intel package
if [ -f dist/xiaoniao-darwin-amd64 ]; then
    echo "  • Creating macOS Intel ZIP..."
    cd dist
    mkdir -p xiaoniao-mac-intel
    cp xiaoniao-darwin-amd64 xiaoniao-mac-intel/xiaoniao
    cp ../start.command xiaoniao-mac-intel/
    chmod +x xiaoniao-mac-intel/xiaoniao
    chmod +x xiaoniao-mac-intel/start.command
    zip -q -r xiaoniao-darwin-amd64.zip xiaoniao-mac-intel
    rm -rf xiaoniao-mac-intel
    cd ..
    echo "    ✓ macOS Intel ZIP created: dist/xiaoniao-darwin-amd64.zip"
fi

# Create macOS Apple Silicon package
if [ -f dist/xiaoniao-darwin-arm64 ]; then
    echo "  • Creating macOS Apple Silicon ZIP..."
    cd dist
    mkdir -p xiaoniao-mac-arm64
    cp xiaoniao-darwin-arm64 xiaoniao-mac-arm64/xiaoniao
    cp ../start.command xiaoniao-mac-arm64/
    chmod +x xiaoniao-mac-arm64/xiaoniao
    chmod +x xiaoniao-mac-arm64/start.command
    zip -q -r xiaoniao-darwin-arm64.zip xiaoniao-mac-arm64
    rm -rf xiaoniao-mac-arm64
    cd ..
    echo "    ✓ macOS Apple Silicon ZIP created: dist/xiaoniao-darwin-arm64.zip"
fi

echo ""
echo "✅ Build complete!"
echo ""
echo "📦 Distribution files:"
echo "  • Linux: dist/xiaoniao-linux-amd64"
echo "  • Windows: dist/xiaoniao-windows.zip"
echo "  • macOS Intel: dist/xiaoniao-darwin-amd64.zip"
echo "  • macOS Apple Silicon: dist/xiaoniao-darwin-arm64.zip"
echo ""
echo "📝 Installation scripts:"
echo "  • Linux: dist/linux-install.sh"
echo "  • Linux: dist/linux-uninstall.sh"
echo ""