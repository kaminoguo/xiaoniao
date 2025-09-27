//go:build windows
// +build windows

// Package clipboard provides clipboard monitoring functionality

package clipboard

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	// Use atomic operations for thread-safe access
	running         int32 // 0 = stopped, 1 = running
	onChange        func(string)
	lastContent     string
	lastTranslation string
	autoTranslate   bool
	
	// Synchronization
	mu              sync.RWMutex
	stopCh          chan struct{}
	doneCh          chan struct{}
	wg              sync.WaitGroup
}

// NewMonitor creates a new clipboard monitor (Windows implementation)
func NewMonitor() *Monitor {
	return &Monitor{
		running:       0,
		autoTranslate: false,
		stopCh:        make(chan struct{}),
		doneCh:        make(chan struct{}),
	}
}

// Start starts the clipboard monitoring
func (m *Monitor) Start() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// Check if already running
	if atomic.LoadInt32(&m.running) == 1 {
		return
	}
	
	// Stop any existing goroutine first (in case of race condition)
	m.stopInternal()
	
	// Create new channels for this session
	m.stopCh = make(chan struct{})
	m.doneCh = make(chan struct{})
	
	// Get current clipboard content and set it as lastContent to prevent initial translation
	if currentText, err := GetText(); err == nil {
		m.lastContent = currentText
	}
	
	// Set running flag
	atomic.StoreInt32(&m.running, 1)
	
	// Start monitoring goroutine
	m.wg.Add(1)
	go func() {
		defer func() {
			m.wg.Done()
			close(m.doneCh)
		}()
		
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-m.stopCh:
				// Stop signal received
				return
			case <-ticker.C:
				// Check if still running (double check)
				if atomic.LoadInt32(&m.running) == 0 {
					return
				}
				
				// Monitor clipboard
				text, err := GetText()
				if err == nil && text != "" {
					m.mu.RLock()
					lastContent := m.lastContent
					lastTranslation := m.lastTranslation
					onChange := m.onChange
					m.mu.RUnlock()
					
					if text != lastContent {
						// Check if the new content is the same as our last translation
						// If so, skip it to avoid circular translation
						if text == lastTranslation {
							m.mu.Lock()
							m.lastContent = text
							m.mu.Unlock()
							continue
						}
						
						m.mu.Lock()
						m.lastContent = text
						m.mu.Unlock()
						
						if onChange != nil {
							onChange(text)
						}
					}
				}
			}
		}
	}()
}

// stopInternal stops monitoring without acquiring lock (internal use)
func (m *Monitor) stopInternal() {
	if atomic.LoadInt32(&m.running) == 0 {
		return
	}
	
	// Set running flag to false
	atomic.StoreInt32(&m.running, 0)
	
	// Signal stop and wait for goroutine to finish
	select {
	case <-m.stopCh:
		// Channel already closed
	default:
		close(m.stopCh)
	}
	
	// Wait for goroutine to finish (with timeout to avoid deadlock)
	done := make(chan struct{})
	go func() {
		m.wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		// Goroutine finished cleanly
	case <-time.After(2 * time.Second):
		// Timeout - goroutine is taking too long, proceed anyway
		// This shouldn't happen in normal circumstances
	}
}

// Stop stops the clipboard monitoring
func (m *Monitor) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.stopInternal()
}

// Enable enables the clipboard monitoring (deprecated - use Start instead)
func (m *Monitor) Enable() {
	m.Start()
}

// Disable disables the clipboard monitoring (deprecated - use Stop instead)
func (m *Monitor) Disable() {
	m.Stop()
}

// IsEnabled returns whether monitoring is enabled
func (m *Monitor) IsEnabled() bool {
	return atomic.LoadInt32(&m.running) == 1
}

// SetAutoTranslate sets auto-translate mode
func (m *Monitor) SetAutoTranslate(enabled bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.autoTranslate = enabled
}

// IsAutoTranslate returns whether auto-translate is enabled
func (m *Monitor) IsAutoTranslate() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.autoTranslate
}

// SetOnChange sets the callback function
func (m *Monitor) SetOnChange(fn func(string)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.onChange = fn
}

// GetLastContent returns the last clipboard content
func (m *Monitor) GetLastContent() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.lastContent
}

// SetLastTranslation sets the last translation (for avoiding duplicates)
func (m *Monitor) SetLastTranslation(translation string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lastTranslation = translation
}
