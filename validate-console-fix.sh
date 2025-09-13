#!/bin/bash
echo "=== Console Hiding Fix Validation ==="
echo ""

# Check if executables exist
if [ -f "xiaoniao.exe" ] && [ -f "xiaoniao-debug.exe" ]; then
    echo "✅ Both executables built successfully"
    echo "   - xiaoniao.exe (production)"
    echo "   - xiaoniao-debug.exe (with debug output)"
else
    echo "❌ Missing executables"
    exit 1
fi

# Check executable format
echo ""
echo "🔍 Executable verification:"
file xiaoniao.exe | grep -q "PE32+ executable.*GUI" && echo "✅ Production version: GUI application" || echo "❌ Production version: Wrong format"
file xiaoniao-debug.exe | grep -q "PE32+ executable.*GUI" && echo "✅ Debug version: GUI application" || echo "❌ Debug version: Wrong format"

# Check for required functions in source
echo ""
echo "🔍 Source code verification:"
grep -q "initializeHiddenConsole" cmd/xiaoniao/main.go && echo "✅ initializeHiddenConsole() is called from main" || echo "❌ Missing initializeHiddenConsole() call"
grep -q "createInvisibleParentWindow" cmd/xiaoniao/windows.go && echo "✅ Invisible parent window creation implemented" || echo "❌ Missing invisible parent window"
grep -q "WS_EX_TOOLWINDOW" cmd/xiaoniao/windows.go && echo "✅ Proper window styles for taskbar hiding" || echo "❌ Missing taskbar hiding styles"
grep -q "SetParent" cmd/xiaoniao/windows.go && echo "✅ Console parenting implemented" || echo "❌ Missing console parenting"

echo ""
echo "=== Testing Instructions ==="
echo ""
echo "To test on Windows machine:"
echo "1. Copy xiaoniao.exe to Windows"
echo "2. Run: xiaoniao.exe run"
echo "3. Verify console does NOT appear in taskbar"
echo "4. Check system tray for xiaoniao icon"
echo ""
echo "To debug issues:"
echo "1. Copy xiaoniao-debug.exe to Windows" 
echo "2. Run: set DEBUG_CONSOLE=1 && xiaoniao-debug.exe run"
echo "3. Check console output for Windows API call results"
echo "4. Look for handle values and success/failure indicators"
echo ""
echo "Expected behavior:"
echo "✅ No console window in taskbar"
echo "✅ System tray icon appears"
echo "✅ Can toggle console via tray menu"
echo "✅ Debug output shows successful API calls"
echo ""
