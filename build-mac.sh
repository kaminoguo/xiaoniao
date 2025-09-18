#!/bin/bash

# Build script for macOS versions of xiaoniao

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

    GOOS=darwin GOARCH=$arch CGO_ENABLED=1 go build \
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

# Build for Intel Macs
build_mac "amd64" "xiaoniao-darwin-amd64"

# Build for Apple Silicon Macs
build_mac "arm64" "xiaoniao-darwin-arm64"

# Create universal binary (optional)
echo -e "${YELLOW}Creating universal binary...${NC}"
lipo -create dist/xiaoniao-darwin-amd64 dist/xiaoniao-darwin-arm64 \
     -output dist/xiaoniao-universal

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Created universal binary successfully${NC}"
fi

# Create app bundle
create_app_bundle() {
    local binary=$1
    local app_name=$2

    echo -e "${YELLOW}Creating app bundle for $app_name...${NC}"

    APP_DIR="dist/$app_name.app"
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

    # Create icon (placeholder - should be replaced with actual .icns file)
    if [ -f "assets/icon.icns" ]; then
        cp assets/icon.icns "$APP_DIR/Contents/Resources/xiaoniao.icns"
        # Add icon reference to Info.plist
        sed -i '' '/<key>NSHighResolutionCapable<\/key>/a\
    <key>CFBundleIconFile</key>\
    <string>xiaoniao</string>' "$APP_DIR/Contents/Info.plist"
    fi

    echo -e "${GREEN}✓ Created $app_name.app${NC}"
}

# Create app bundles
create_app_bundle "xiaoniao-universal" "xiaoniao"
create_app_bundle "xiaoniao-darwin-amd64" "xiaoniao-intel"
create_app_bundle "xiaoniao-darwin-arm64" "xiaoniao-arm"

# Create DMG (optional, requires create-dmg tool)
if command -v create-dmg &> /dev/null; then
    echo -e "${YELLOW}Creating DMG installer...${NC}"
    create-dmg \
        --volname "小鸟翻译 $VERSION" \
        --window-pos 200 120 \
        --window-size 600 400 \
        --icon-size 100 \
        --app-drop-link 450 185 \
        "dist/xiaoniao-$VERSION.dmg" \
        "dist/xiaoniao.app"

    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Created DMG successfully${NC}"
    fi
else
    echo -e "${YELLOW}Note: Install create-dmg for DMG creation${NC}"
    echo "  brew install create-dmg"
fi

# Create ZIP archives
echo -e "${YELLOW}Creating ZIP archives...${NC}"
cd dist

for app in xiaoniao.app xiaoniao-intel.app xiaoniao-arm.app; do
    if [ -d "$app" ]; then
        zip -r "${app%.app}.zip" "$app"
        echo -e "${GREEN}✓ Created ${app%.app}.zip${NC}"
    fi
done

cd ..

echo -e "${GREEN}Build complete! Files in dist/:${NC}"
ls -la dist/

echo -e "${YELLOW}Installation instructions:${NC}"
echo "1. Unzip the appropriate version for your Mac"
echo "2. Drag xiaoniao.app to Applications folder"
echo "3. Right-click and select 'Open' for first launch"
echo "4. Grant accessibility permissions in System Settings"

# Make script executable
chmod +x build-mac.sh