//go:build windows

package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	user32                    = syscall.NewLazyDLL("user32.dll")
	kernel32                  = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleWindow      = kernel32.NewProc("GetConsoleWindow")
	procShowWindow            = user32.NewProc("ShowWindow")
	procAllocConsole          = kernel32.NewProc("AllocConsole")
	procFreeConsole           = kernel32.NewProc("FreeConsole")
	procSetConsoleCtrlHandler = kernel32.NewProc("SetConsoleCtrlHandler")
	procMessageBox            = user32.NewProc("MessageBoxW")
	procSendInput             = user32.NewProc("SendInput")
	procEnableWindow          = user32.NewProc("EnableWindow")
	procGetStdHandle          = kernel32.NewProc("GetStdHandle")
	procSetStdHandle          = kernel32.NewProc("SetStdHandle")
	procCreateFile            = kernel32.NewProc("CreateFileW")
	procSetForegroundWindow   = user32.NewProc("SetForegroundWindow")
	procGetForegroundWindow   = user32.NewProc("GetForegroundWindow")
	procSetConsoleTitle       = kernel32.NewProc("SetConsoleTitleW")
	
	// Track console window visibility state
	mainConsoleHandle      uintptr  // 保存主控制台的句柄，用于显示/隐藏切换
	isConsoleVisible       = true
	hasDebugConsole       = false
	originalStdOut        uintptr
	originalStdErr        uintptr
)

const (
	SW_HIDE            = 0
	SW_SHOWNORMAL      = 1
	
	// Standard handles
	STD_INPUT_HANDLE  = ^uintptr(9)  // -10
	STD_OUTPUT_HANDLE = ^uintptr(10) // -11
	STD_ERROR_HANDLE  = ^uintptr(11) // -12
	
	// File access modes
	GENERIC_READ    = 0x80000000
	GENERIC_WRITE   = 0x40000000
	FILE_SHARE_READ = 0x00000001
	FILE_SHARE_WRITE = 0x00000002
	OPEN_EXISTING   = 3
	
	// Virtual Key Codes
	VK_CONTROL         = 0x11
	VK_V               = 0x56
	
	// Input Types
	INPUT_KEYBOARD     = 1
	
	// Key Event Flags
	KEYEVENTF_KEYUP    = 0x0002
	
	// Console control handler events
	CTRL_CLOSE_EVENT = 2
)

// consoleCtrlHandler handles console close events to prevent program exit
func consoleCtrlHandler(ctrlType uintptr) uintptr {
	if ctrlType == CTRL_CLOSE_EVENT {
		// Don't exit the program, just hide the debug console
		hideDebugConsole()
		return 1 // TRUE - handled
	}
	return 0 // FALSE - not handled
}

// mainConsoleCtrlHandler handles main console close events to prevent program exit
func mainConsoleCtrlHandler(ctrlType uintptr) uintptr {
	if ctrlType == CTRL_CLOSE_EVENT {
		// Don't exit the program, just hide the main console
		hideConsoleWindow()
		return 1 // TRUE - handled
	}
	return 0 // FALSE - not handled
}


// showConsoleWindow ensures console window is visible on Windows
func showConsoleWindow() {
	// Try to show existing console first
	if mainConsoleHandle != 0 {
		// Try to show the saved console
		ret, _, _ := procShowWindow.Call(mainConsoleHandle, SW_SHOWNORMAL)
		if ret != 0 {
			// Successfully shown
			procSetForegroundWindow.Call(mainConsoleHandle)
			isConsoleVisible = true
			return
		}
	}
	
	// If can't show existing console, allocate a new one
	ret, _, _ := procAllocConsole.Call()
	if ret != 0 {
		// Get the new console window handle
		consoleWindow, _, _ := procGetConsoleWindow.Call()
		if consoleWindow != 0 {
			mainConsoleHandle = consoleWindow
			
			// Set console title
			titlePtr, _ := syscall.UTF16PtrFromString("xiaoniao - 主控制台")
			procSetConsoleTitle.Call(uintptr(unsafe.Pointer(titlePtr)))
			
			// Reopen standard handles to connect to new console
			reopenStdHandles()
			
			// Show the console
			procShowWindow.Call(consoleWindow, SW_SHOWNORMAL)
			procSetForegroundWindow.Call(consoleWindow)
			
			// Print welcome message
			fmt.Println("=== xiaoniao 主控制台 ===")
			fmt.Println("控制台已重新创建")
			fmt.Println("等待翻译活动...")
			fmt.Println("")
			
			isConsoleVisible = true
		}
	}
}

// hideConsoleWindow hides console window (works with Windows Terminal)
func hideConsoleWindow() {
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	
	if consoleWindow != 0 {
		// Save main console handle on first call (from main process)
		if mainConsoleHandle == 0 {
			mainConsoleHandle = consoleWindow
		}
		
		// Windows Terminal fix: Set as foreground window first
		// This is required for Windows Terminal to properly hide
		procSetForegroundWindow.Call(consoleWindow)
		
		// Get the updated handle after setting foreground
		// Windows Terminal needs this to properly hide instead of just minimizing
		handle, _, _ := procGetForegroundWindow.Call()
		
		// Now SW_HIDE properly hides the window completely (not just minimize to taskbar)
		procShowWindow.Call(handle, SW_HIDE)
		isConsoleVisible = false
	}
}


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

// showDebugConsole creates and shows a debug console window
func showDebugConsole() {
	if hasDebugConsole {
		// Console already exists, just show it
		consoleWindow, _, _ := procGetConsoleWindow.Call()
		if consoleWindow != 0 {
			procShowWindow.Call(consoleWindow, SW_SHOWNORMAL)
		}
		return
	}
	
	// Store original handles
	originalStdOut, _, _ = procGetStdHandle.Call(STD_OUTPUT_HANDLE)
	originalStdErr, _, _ = procGetStdHandle.Call(STD_ERROR_HANDLE)
	
	// Allocate a new console
	result, _, err := procAllocConsole.Call()
	if err != nil && err.Error() != "The operation completed successfully." {
		return
	}
	
	if result == 0 {
		return // Failed to allocate console
	}
	
	hasDebugConsole = true
	
	// Set console control handler to prevent program exit on console close
	handlerPtr := syscall.NewCallback(consoleCtrlHandler)
	procSetConsoleCtrlHandler.Call(handlerPtr, 1) // TRUE - add handler
	
	// Redirect stdout/stderr to the new console
	// Note: In console mode, stdout/stderr are already available
	
	// Set console window title
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	if consoleWindow != 0 {
		titlePtr, _ := syscall.UTF16PtrFromString("xiaoniao - 调试控制台")
		procSetConsoleTitle := kernel32.NewProc("SetConsoleTitleW")
		procSetConsoleTitle.Call(uintptr(unsafe.Pointer(titlePtr)))
	}
	
	// Show welcome message
	fmt.Println("=== xiaoniao 调试控制台 ===")
	fmt.Println("调试信息将在此处显示")
	fmt.Println("关闭此窗口不会退出程序，只会隐藏调试窗口")
	fmt.Println("可通过系统托盘菜单重新显示")
	fmt.Println("================================")
	fmt.Println("控制台已启动，等待翻译活动...")
	fmt.Println("")
}

// hideDebugConsole hides and frees the debug console
func hideDebugConsole() {
	if !hasDebugConsole {
		return
	}
	
	// Restore original handles
	if originalStdOut != 0 {
		procSetStdHandle.Call(STD_OUTPUT_HANDLE, originalStdOut)
	}
	if originalStdErr != 0 {
		procSetStdHandle.Call(STD_ERROR_HANDLE, originalStdErr)
	}
	
	// Remove console control handler
	handlerPtr := syscall.NewCallback(consoleCtrlHandler)
	procSetConsoleCtrlHandler.Call(handlerPtr, 0) // FALSE - remove handler
	
	// Free the console
	procFreeConsole.Call()
	
	hasDebugConsole = false
}

// isDebugConsoleVisible returns whether debug console is currently visible
func isDebugConsoleVisible() bool {
	return hasDebugConsole
}

// setupMainConsoleHandler sets up the main console control handler to prevent program exit on X click
func setupMainConsoleHandler() {
	// Set console control handler to prevent program exit on console close
	handlerPtr := syscall.NewCallback(mainConsoleCtrlHandler)
	procSetConsoleCtrlHandler.Call(handlerPtr, 1) // TRUE - add handler
}

// isMainConsoleVisible returns whether main console is currently visible
func isMainConsoleVisible() bool {
	return isConsoleVisible
}

// reopenStdHandles reconnects standard I/O to the new console
func reopenStdHandles() {
	// Reopen stdout
	stdout, err := syscall.UTF16PtrFromString("CONOUT$")
	if err == nil {
		handle, _, _ := procCreateFile.Call(
			uintptr(unsafe.Pointer(stdout)),
			GENERIC_WRITE,
			FILE_SHARE_READ|FILE_SHARE_WRITE,
			0,
			OPEN_EXISTING,
			0,
			0,
		)
		if handle != 0 && handle != ^uintptr(0) {
			procSetStdHandle.Call(STD_OUTPUT_HANDLE, handle)
			// Reconnect Go's stdout
			os.Stdout = os.NewFile(uintptr(handle), "stdout")
		}
	}
	
	// Reopen stderr
	stderr, err := syscall.UTF16PtrFromString("CONOUT$")
	if err == nil {
		handle, _, _ := procCreateFile.Call(
			uintptr(unsafe.Pointer(stderr)),
			GENERIC_WRITE,
			FILE_SHARE_READ|FILE_SHARE_WRITE,
			0,
			OPEN_EXISTING,
			0,
			0,
		)
		if handle != 0 && handle != ^uintptr(0) {
			procSetStdHandle.Call(STD_ERROR_HANDLE, handle)
			// Reconnect Go's stderr
			os.Stderr = os.NewFile(uintptr(handle), "stderr")
		}
	}
	
	// Reopen stdin
	stdin, err := syscall.UTF16PtrFromString("CONIN$")
	if err == nil {
		handle, _, _ := procCreateFile.Call(
			uintptr(unsafe.Pointer(stdin)),
			GENERIC_READ|GENERIC_WRITE,
			FILE_SHARE_READ|FILE_SHARE_WRITE,
			0,
			OPEN_EXISTING,
			0,
			0,
		)
		if handle != 0 && handle != ^uintptr(0) {
			procSetStdHandle.Call(STD_INPUT_HANDLE, handle)
			// Reconnect Go's stdin
			os.Stdin = os.NewFile(uintptr(handle), "stdin")
		}
	}
}