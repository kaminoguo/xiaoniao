#!/bin/bash

# Test script to run xiaoniao with debug output and check console hiding
echo "Testing xiaoniao console hiding functionality..."

# Set up minimal config for testing
mkdir -p "$HOME/.config/xiaoniao" 2>/dev/null
cat > "$HOME/.config/xiaoniao/config.json" << INNER_EOF
{
  "api_key": "test-key-for-debug",
  "provider": "OpenAI",
  "model": "gpt-4o-mini",
  "prompt_id": "direct"
}
INNER_EOF

echo "Config created. Now testing Windows executable..."

# Run the Windows executable through Wine (if available) or just show the file info
if command -v wine &> /dev/null; then
    echo "Running through Wine to test Windows functionality:"
    timeout 10s wine xiaoniao.exe run 2>&1 || echo "Wine test completed or timed out"
else
    echo "Wine not available. Executable built successfully:"
    file xiaoniao.exe
    echo ""
    echo "To test on Windows, run: xiaoniao.exe run"
    echo "Expected debug output should show:"
    echo "  - initializeHiddenConsole() called"
    echo "  - Creating invisible parent window with handles"
    echo "  - Console window detection and SetParent calls"
    echo "  - Console hiding steps with results"
fi
