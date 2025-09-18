#!/bin/bash

# Build script for macOS version of xiaoniao

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Version
VERSION=${1:-"v1.1.0"}

echo -e "${GREEN}Building xiaoniao for macOS - Version $VERSION${NC}"

# Create dist directory
mkdir -p dist

# Function to build for a specific architecture
build_mac() {
    local arch=$1
    local output=$2

    echo -e "${YELLOW}Building for macOS $arch...${NC}"

    CGO_ENABLED=1 GOOS=darwin GOARCH=$arch go build \
        -ldflags="-s -w -X main.version=$VERSION" \
        -o "dist/$output" \
        ./cmd/xiaoniao

    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Built $output successfully${NC}"
    else
        echo -e "${RED}✗ Failed to build $output${NC}"
        exit 1
    fi
}

# Detect current architecture
ARCH=$(uname -m)
if [ "$ARCH" = "arm64" ]; then
    echo -e "${YELLOW}Detected Apple Silicon Mac${NC}"
    build_mac "arm64" "xiaoniao"
elif [ "$ARCH" = "x86_64" ]; then
    echo -e "${YELLOW}Detected Intel Mac${NC}"
    build_mac "amd64" "xiaoniao"
else
    echo -e "${RED}Unknown architecture: $ARCH${NC}"
    exit 1
fi

# Create app bundle
create_app_bundle() {
    local binary=$1

    echo -e "${YELLOW}Creating app bundle...${NC}"

    APP_DIR="dist/xiaoniao.app"
    rm -rf "$APP_DIR"
    mkdir -p "$APP_DIR/Contents/MacOS"
    mkdir -p "$APP_DIR/Contents/Resources"

    # Copy binary
    cp "dist/$binary" "$APP_DIR/Contents/MacOS/xiaoniao"
    chmod +x "$APP_DIR/Contents/MacOS/xiaoniao"

    # Create Info.plist
    cat > "$APP_DIR/Contents/Info.plist" << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>xiaoniao</string>
    <key>CFBundleIdentifier</key>
    <string>com.lyrica.xiaoniao</string>
    <key>CFBundleName</key>
    <string>小鸟翻译</string>
    <key>CFBundleDisplayName</key>
    <string>小鸟翻译</string>
    <key>CFBundleVersion</key>
    <string>$VERSION</string>
    <key>CFBundleShortVersionString</key>
    <string>$VERSION</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>LSMinimumSystemVersion</key>
    <string>10.12</string>
    <key>LSUIElement</key>
    <true/>
    <key>NSHighResolutionCapable</key>
    <true/>
    <key>NSAppleEventsUsageDescription</key>
    <string>小鸟翻译需要控制其他应用程序以实现自动粘贴功能。</string>
    <key>NSAccessibilityUsageDescription</key>
    <string>小鸟翻译需要辅助功能权限以监听全局快捷键。</string>
</dict>
</plist>
EOF

    echo -e "${GREEN}✓ Created xiaoniao.app${NC}"
}

# Create app bundle
create_app_bundle "xiaoniao"

# Create ZIP archive
echo -e "${YELLOW}Creating ZIP archive...${NC}"
cd dist
zip -r "xiaoniao-mac-$ARCH.zip" "xiaoniao.app"
cd ..

echo -e "${GREEN}Build complete!${NC}"
echo -e "${GREEN}Files created:${NC}"
echo "  - dist/xiaoniao (executable)"
echo "  - dist/xiaoniao.app (app bundle)"
echo "  - dist/xiaoniao-mac-$ARCH.zip (distribution)"

echo -e "${YELLOW}Installation instructions:${NC}"
echo "1. Unzip xiaoniao-mac-$ARCH.zip"
echo "2. Drag xiaoniao.app to Applications folder"
echo "3. Right-click and select 'Open' for first launch"
echo "4. Grant accessibility permissions in System Settings"
echo ""
echo -e "${YELLOW}Note: First run 'go mod tidy' to download dependencies${NC}"