#!/bin/bash
echo "Building Windows debug version with console hiding debug output..."

# Build debug version
GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o xiaoniao-debug.exe ./cmd/xiaoniao

if [ $? -eq 0 ]; then
    echo "✅ Debug build successful: xiaoniao-debug.exe"
    echo ""
    echo "To test with debug output on Windows:"
    echo "  set DEBUG_CONSOLE=1"
    echo "  xiaoniao-debug.exe run"
    echo ""
    echo "Expected debug output will show:"
    echo "  - Invisible parent window creation with handles"
    echo "  - Console window detection and parenting"
    echo "  - All Windows API call results"
else
    echo "❌ Debug build failed"
    exit 1
fi
