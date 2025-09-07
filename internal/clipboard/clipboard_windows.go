//go:build windows
// +build windows

package clipboard

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32                     = syscall.NewLazyDLL("user32.dll")
	kernel32                   = syscall.NewLazyDLL("kernel32.dll")
	openClipboard              = user32.NewProc("OpenClipboard")
	closeClipboard             = user32.NewProc("CloseClipboard")
	getClipboardData           = user32.NewProc("GetClipboardData")
	setClipboardData           = user32.NewProc("SetClipboardData")
	emptyClipboard             = user32.NewProc("EmptyClipboard")
	isClipboardFormatAvailable = user32.NewProc("IsClipboardFormatAvailable")
	globalLock                 = kernel32.NewProc("GlobalLock")
	globalUnlock               = kernel32.NewProc("GlobalUnlock")
	globalAlloc                = kernel32.NewProc("GlobalAlloc")
	globalSize                 = kernel32.NewProc("GlobalSize")
)

const (
	cfUnicodeText = 13
	gmemMoveable  = 0x0002
)

// GetText retrieves text from the Windows clipboard
func GetText() (string, error) {
	// Open clipboard
	ret, _, err := openClipboard.Call(0)
	if ret == 0 {
		return "", fmt.Errorf("failed to open clipboard: %v", err)
	}
	defer closeClipboard.Call()

	// Check if text is available
	ret, _, _ = isClipboardFormatAvailable.Call(cfUnicodeText)
	if ret == 0 {
		return "", nil // No text available
	}

	// Get clipboard data
	handle, _, err := getClipboardData.Call(cfUnicodeText)
	if handle == 0 {
		return "", fmt.Errorf("failed to get clipboard data: %v", err)
	}

	// Lock memory
	ptr, _, err := globalLock.Call(handle)
	if ptr == 0 {
		return "", fmt.Errorf("failed to lock memory: %v", err)
	}
	defer globalUnlock.Call(handle)

	// Get size and convert to string
	size, _, _ := globalSize.Call(handle)
	if size == 0 {
		return "", nil
	}

	// Convert UTF-16 to string
	text := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(ptr))[:size/2])
	return text, nil
}

// SetText sets text to the Windows clipboard
func SetText(text string) error {
	// Convert string to UTF-16
	utf16, err := syscall.UTF16FromString(text)
	if err != nil {
		return fmt.Errorf("failed to convert text: %v", err)
	}

	// Calculate size
	size := len(utf16) * 2

	// Open clipboard
	ret, _, err := openClipboard.Call(0)
	if ret == 0 {
		return fmt.Errorf("failed to open clipboard: %v", err)
	}
	defer closeClipboard.Call()

	// Empty clipboard
	ret, _, err = emptyClipboard.Call()
	if ret == 0 {
		return fmt.Errorf("failed to empty clipboard: %v", err)
	}

	// Allocate global memory
	handle, _, err := globalAlloc.Call(gmemMoveable, uintptr(size))
	if handle == 0 {
		return fmt.Errorf("failed to allocate memory: %v", err)
	}

	// Lock memory
	ptr, _, err := globalLock.Call(handle)
	if ptr == 0 {
		return fmt.Errorf("failed to lock memory: %v", err)
	}
	defer globalUnlock.Call(handle)

	// Copy data
	for i, v := range utf16 {
		*(*uint16)(unsafe.Pointer(ptr + uintptr(i*2))) = v
	}

	// Set clipboard data
	ret, _, err = setClipboardData.Call(cfUnicodeText, handle)
	if ret == 0 {
		return fmt.Errorf("failed to set clipboard data: %v", err)
	}

	return nil
}

// SetClipboard is an alias for SetText (Windows compatibility)
func SetClipboard(text string) error {
	return SetText(text)
}

// Monitor structure for Windows implementation
type Monitor struct {
	enabled       bool
	onChange      func(string)
	lastContent   string
	autoTranslate bool
}

// NewMonitor creates a new clipboard monitor (Windows implementation)
func NewMonitor() *Monitor {
	return &Monitor{
		enabled:       false,
		autoTranslate: false,
	}
}

// Start starts the clipboard monitoring
func (m *Monitor) Start() {
	if m.enabled {
		return
	}
	m.enabled = true
	
	go func() {
		for m.enabled {
			text, err := GetText()
			if err == nil && text != "" && text != m.lastContent {
				m.lastContent = text
				if m.onChange != nil {
					m.onChange(text)
				}
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

// Stop stops the clipboard monitoring
func (m *Monitor) Stop() {
	m.enabled = false
}

// Enable enables the clipboard monitoring
func (m *Monitor) Enable() {
	m.enabled = true
}

// Disable disables the clipboard monitoring
func (m *Monitor) Disable() {
	m.enabled = false
}

// IsEnabled returns whether monitoring is enabled
func (m *Monitor) IsEnabled() bool {
	return m.enabled
}

// SetAutoTranslate sets auto-translate mode
func (m *Monitor) SetAutoTranslate(enabled bool) {
	m.autoTranslate = enabled
}

// IsAutoTranslate returns whether auto-translate is enabled
func (m *Monitor) IsAutoTranslate() bool {
	return m.autoTranslate
}

// SetOnChange sets the callback function
func (m *Monitor) SetOnChange(fn func(string)) {
	m.onChange = fn
}

// GetLastContent returns the last clipboard content
func (m *Monitor) GetLastContent() string {
	return m.lastContent
}

// SetLastTranslation sets the last translation (for avoiding duplicates)
func (m *Monitor) SetLastTranslation(translation string) {
	// Not used in Windows implementation
}