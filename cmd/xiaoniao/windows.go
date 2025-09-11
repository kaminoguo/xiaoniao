//go:build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleWindow = kernel32.NewProc("GetConsoleWindow")
	procShowWindow       = user32.NewProc("ShowWindow")
	procAllocConsole     = kernel32.NewProc("AllocConsole")
	procMessageBox       = user32.NewProc("MessageBoxW")
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
)

// showConsoleWindow ensures console window is visible on Windows
func showConsoleWindow() {
	// Get console window handle
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	
	if consoleWindow == 0 {
		// No console window, allocate one
		procAllocConsole.Call()
		consoleWindow, _, _ = procGetConsoleWindow.Call()
	}
	
	if consoleWindow != 0 {
		// Show the console window (normal state)
		procShowWindow.Call(consoleWindow, SW_SHOWNORMAL)
	}
}

// hideConsoleWindow hides console window on Windows
func hideConsoleWindow() {
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	if consoleWindow != 0 {
		procShowWindow.Call(consoleWindow, SW_HIDE)
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