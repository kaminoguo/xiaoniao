//go:build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32                 = syscall.NewLazyDLL("user32.dll")
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleWindow   = kernel32.NewProc("GetConsoleWindow")
	procShowWindow         = user32.NewProc("ShowWindow")
	procShowWindowAsync    = user32.NewProc("ShowWindowAsync")
	procAllocConsole       = kernel32.NewProc("AllocConsole")
	procFreeConsole        = kernel32.NewProc("FreeConsole")
	procMessageBox         = user32.NewProc("MessageBoxW")
	procSendInput          = user32.NewProc("SendInput")
	procGetWindowLongPtr   = user32.NewProc("GetWindowLongPtrW")
	procSetWindowLongPtr   = user32.NewProc("SetWindowLongPtrW")
	procSetParent          = user32.NewProc("SetParent")
	procGetDesktopWindow   = user32.NewProc("GetDesktopWindow")
	procSetWindowPos       = user32.NewProc("SetWindowPos")
	procEnableWindow       = user32.NewProc("EnableWindow")
	
	// Track console window visibility state
	isConsoleVisible = true
	originalParent   uintptr = 0
	
	// Window style constants that need to be variables for negative values
	GWL_EXSTYLE = uintptr(0xFFFFFFFFFFFFFFEC) // -20 as unsigned value on 64-bit
	GWL_STYLE   = uintptr(0xFFFFFFFFFFFFFFF0) // -16 as unsigned value on 64-bit
)

const (
	SW_HIDE            = 0
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_FORCEMINIMIZE   = 11
	
	// Window style constants  
	WS_EX_APPWINDOW    = 0x00040000
	WS_EX_TOOLWINDOW   = 0x00000080
	WS_EX_NOACTIVATE   = 0x08000000
	WS_VISIBLE         = 0x10000000
	WS_MINIMIZE        = 0x20000000
	
	// SetWindowPos flags
	SWP_NOSIZE         = 0x0001
	SWP_NOMOVE         = 0x0002
	SWP_NOZORDER       = 0x0004
	SWP_NOREDRAW       = 0x0008
	SWP_NOACTIVATE     = 0x0010
	SWP_FRAMECHANGED   = 0x0020
	SWP_SHOWWINDOW     = 0x0040
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOOWNERZORDER  = 0x0200
	
	HWND_BOTTOM        = 1
	
	// Virtual Key Codes
	VK_CONTROL         = 0x11
	VK_V               = 0x56
	
	// Input Types
	INPUT_KEYBOARD     = 1
	
	// Key Event Flags
	KEYEVENTF_KEYUP    = 0x0002
)

// hideFromTaskbar removes window from taskbar by modifying window styles
func hideFromTaskbar(hwnd uintptr) {
	if hwnd == 0 {
		return
	}
	
	// Get current extended window styles
	exStyle, _, _ := procGetWindowLongPtr.Call(hwnd, GWL_EXSTYLE)
	
	// Remove WS_EX_APPWINDOW and add WS_EX_TOOLWINDOW to hide from taskbar
	newExStyle := (exStyle &^ uintptr(WS_EX_APPWINDOW)) | uintptr(WS_EX_TOOLWINDOW)
	
	// Set the new extended window styles
	procSetWindowLongPtr.Call(hwnd, GWL_EXSTYLE, newExStyle)
}

// hideFromTaskbarCompletely completely removes window from taskbar and Alt+Tab
func hideFromTaskbarCompletely(hwnd uintptr) {
	if hwnd == 0 {
		return
	}
	
	// Get current extended window styles
	exStyle, _, _ := procGetWindowLongPtr.Call(hwnd, GWL_EXSTYLE)
	
	// 彻底隐藏：移除所有任务栏和Alt+Tab相关样式
	// 添加 WS_EX_TOOLWINDOW：工具窗口，不在任务栏和Alt+Tab中显示
	// 移除 WS_EX_APPWINDOW：应用程序窗口标志
	newExStyle := (exStyle | uintptr(WS_EX_TOOLWINDOW)) &^ uintptr(WS_EX_APPWINDOW)
	
	// Set the new extended window styles
	procSetWindowLongPtr.Call(hwnd, GWL_EXSTYLE, newExStyle)
}

// hideFromTaskbarUltimate 终极任务栏隐藏方案
func hideFromTaskbarUltimate(hwnd uintptr) {
	if hwnd == 0 {
		return
	}
	
	// Get current extended window styles
	exStyle, _, _ := procGetWindowLongPtr.Call(hwnd, GWL_EXSTYLE)
	
	// 终极隐藏组合：
	// 1. 添加 WS_EX_TOOLWINDOW - 工具窗口，不显示在任务栏
	// 2. 添加 WS_EX_NOACTIVATE - 不激活窗口，减少可见性
	// 3. 移除 WS_EX_APPWINDOW - 移除应用程序窗口标志
	newExStyle := (exStyle | uintptr(WS_EX_TOOLWINDOW|WS_EX_NOACTIVATE)) &^ uintptr(WS_EX_APPWINDOW)
	
	// Set the new extended window styles
	procSetWindowLongPtr.Call(hwnd, GWL_EXSTYLE, newExStyle)
}

// modifyWindowStyleForHiding 修改基本窗口样式以实现隐藏
func modifyWindowStyleForHiding(hwnd uintptr) {
	if hwnd == 0 {
		return
	}
	
	// Get current window styles
	style, _, _ := procGetWindowLongPtr.Call(hwnd, GWL_STYLE)
	
	// 移除窗口可见性和最小化状态
	// 移除 WS_VISIBLE - 窗口可见标志
	// 移除 WS_MINIMIZE - 最小化标志
	newStyle := style &^ uintptr(WS_VISIBLE|WS_MINIMIZE)
	
	// Set the new window styles
	procSetWindowLongPtr.Call(hwnd, GWL_STYLE, newStyle)
}

// restoreToTaskbar restores window to taskbar by modifying window styles
func restoreToTaskbar(hwnd uintptr) {
	if hwnd == 0 {
		return
	}
	
	// Get current extended window styles
	exStyle, _, _ := procGetWindowLongPtr.Call(hwnd, GWL_EXSTYLE)
	
	// Add WS_EX_APPWINDOW and remove WS_EX_TOOLWINDOW to show in taskbar
	// Also remove WS_EX_NOACTIVATE to allow normal activation
	newExStyle := (exStyle | uintptr(WS_EX_APPWINDOW)) &^ uintptr(WS_EX_TOOLWINDOW|WS_EX_NOACTIVATE)
	
	// Set the new extended window styles
	procSetWindowLongPtr.Call(hwnd, GWL_EXSTYLE, newExStyle)
}

// restoreWindowStyleFromHiding 从隐藏状态恢复正常的窗口样式
func restoreWindowStyleFromHiding(hwnd uintptr) {
	if hwnd == 0 {
		return
	}
	
	// Get current window styles
	style, _, _ := procGetWindowLongPtr.Call(hwnd, GWL_STYLE)
	
	// 恢复窗口可见性标志
	// 添加 WS_VISIBLE - 窗口可见标志
	newStyle := style | uintptr(WS_VISIBLE)
	
	// Set the new window styles
	procSetWindowLongPtr.Call(hwnd, GWL_STYLE, newStyle)
}

// showConsoleWindow ensures console window is visible on Windows
func showConsoleWindow() {
	// Get console window handle
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	
	// Only work with existing console window, don't allocate new one
	if consoleWindow == 0 {
		return // No console window available
	}
	
	if consoleWindow != 0 {
		// 恢复窗口的完整过程
		
		// Step 1: 重新启用窗口交互
		procEnableWindow.Call(consoleWindow, 1) // 1 = TRUE
		
		// Step 2: 恢复正常的窗口样式
		restoreWindowStyleFromHiding(consoleWindow)
		
		// Step 3: 恢复到任务栏 (恢复正常样式)
		restoreToTaskbar(consoleWindow)
		
		// Step 4: 使用SetWindowPos恢复窗口
		procSetWindowPos.Call(
			consoleWindow,
			0,              // HWND_TOP
			0, 0, 0, 0,     // 位置和大小参数（被忽略）
			SWP_SHOWWINDOW|SWP_NOSIZE|SWP_NOMOVE|SWP_FRAMECHANGED,
		)
		
		// Step 5: 显示控制台窗口 (正常状态)
		procShowWindow.Call(consoleWindow, SW_SHOWNORMAL)
		
		isConsoleVisible = true
	}
}

// hideConsoleWindow hides console window completely from taskbar and Alt+Tab on Windows
func hideConsoleWindow() {
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	if consoleWindow != 0 {
		// 终极隐藏方案：多层次隐藏策略
		
		// Step 1: 先异步最小化窗口（避免闪烁）
		procShowWindowAsync.Call(consoleWindow, SW_FORCEMINIMIZE)
		
		// Step 2: 修改窗口样式 - 完全从任务栏和Alt+Tab中移除
		hideFromTaskbarUltimate(consoleWindow)
		
		// Step 3: 修改基本窗口样式，移除可见性
		modifyWindowStyleForHiding(consoleWindow)
		
		// Step 4: 使用SetWindowPos进行更底层的隐藏
		procSetWindowPos.Call(
			consoleWindow,
			HWND_BOTTOM,    // 置于最底层
			0, 0, 0, 0,     // 位置和大小参数（被忽略）
			SWP_HIDEWINDOW|SWP_NOSIZE|SWP_NOMOVE|SWP_NOACTIVATE|SWP_NOZORDER,
		)
		
		// Step 5: 最后隐藏窗口
		procShowWindow.Call(consoleWindow, SW_HIDE)
		
		// Step 6: 禁用窗口交互
		procEnableWindow.Call(consoleWindow, 0) // 0 = FALSE
		
		isConsoleVisible = false
	}
}

// initializeHiddenConsole - DEPRECATED: This was for GUI mode, not needed for console application
// func initializeHiddenConsole() {
//	// 获取当前控制台窗口
//	consoleWindow, _, _ := procGetConsoleWindow.Call()
//	
//	if consoleWindow == 0 {
//		// 如果没有控制台，分配一个
//		procAllocConsole.Call()
//		consoleWindow, _, _ = procGetConsoleWindow.Call()
//	}
//	
//	if consoleWindow != 0 {
//		// 立即彻底隐藏
//		hideConsoleWindow()
//	}
// }

// toggleConsoleWindow toggles console window visibility on Windows
func toggleConsoleWindow() {
	if isConsoleVisible {
		hideConsoleWindow()
	} else {
		showConsoleWindow()
	}
}

// showErrorMessage shows error dialog on Windows
func showErrorMessage(title, message string) {
	// First try to show console window
	showConsoleWindow()
	
	// Print to console as fallback
	fmt.Printf("Error - %s: %s\n", title, message)
	
	// Also show Windows message box
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	messagePtr, _ := syscall.UTF16PtrFromString(message)
	
	// MB_OK = 0, MB_ICONERROR = 16
	procMessageBox.Call(0, uintptr(unsafe.Pointer(messagePtr)), uintptr(unsafe.Pointer(titlePtr)), 0+16)
}

// INPUT structure for SendInput API
type INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
	_    [8]byte // padding for union
}

// KEYBDINPUT structure
type KEYBDINPUT struct {
	VirtualKey uint16
	ScanCode   uint16
	Flags      uint32
	Time       uint32
	ExtraInfo  uintptr
}

// simulatePaste simulates Ctrl+V key combination using Windows SendInput API
func simulatePaste() {
	// Create INPUT structures for key events
	inputs := make([]INPUT, 4) // Press Ctrl, Press V, Release V, Release Ctrl
	
	// Press Ctrl
	inputs[0] = INPUT{
		Type: INPUT_KEYBOARD,
		Ki: KEYBDINPUT{
			VirtualKey: VK_CONTROL,
			ScanCode:   0,
			Flags:      0, // Key down
			Time:       0,
			ExtraInfo:  0,
		},
	}
	
	// Press V
	inputs[1] = INPUT{
		Type: INPUT_KEYBOARD,
		Ki: KEYBDINPUT{
			VirtualKey: VK_V,
			ScanCode:   0,
			Flags:      0, // Key down
			Time:       0,
			ExtraInfo:  0,
		},
	}
	
	// Release V
	inputs[2] = INPUT{
		Type: INPUT_KEYBOARD,
		Ki: KEYBDINPUT{
			VirtualKey: VK_V,
			ScanCode:   0,
			Flags:      KEYEVENTF_KEYUP, // Key up
			Time:       0,
			ExtraInfo:  0,
		},
	}
	
	// Release Ctrl
	inputs[3] = INPUT{
		Type: INPUT_KEYBOARD,
		Ki: KEYBDINPUT{
			VirtualKey: VK_CONTROL,
			ScanCode:   0,
			Flags:      KEYEVENTF_KEYUP, // Key up
			Time:       0,
			ExtraInfo:  0,
		},
	}
	
	// Send the input events
	procSendInput.Call(
		uintptr(len(inputs)),                    // nInputs
		uintptr(unsafe.Pointer(&inputs[0])),     // pInputs
		uintptr(unsafe.Sizeof(inputs[0])),       // cbSize
	)
}