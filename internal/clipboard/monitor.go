package clipboard

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Monitor watches clipboard for changes
type Monitor struct {
	ctx            context.Context
	cancel         context.CancelFunc
	mu             sync.RWMutex
	enabled        bool
	lastContent    string
	lastTranslation string  // 记录最后的译文，避免重复翻译
	checkInterval  time.Duration
	onChange       func(string)
	autoTranslate  bool
	wg             sync.WaitGroup
}

// NewMonitor creates a new clipboard monitor
func NewMonitor() *Monitor {
	ctx, cancel := context.WithCancel(context.Background())
	return &Monitor{
		ctx:           ctx,
		cancel:        cancel,
		enabled:       false,
		checkInterval: 500 * time.Millisecond,
		autoTranslate: true,
	}
}

// Start starts monitoring the clipboard
func (m *Monitor) Start() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.enabled {
		return fmt.Errorf("monitor already running")
	}

	// 获取当前剪贴板内容作为初始值，避免启动时立即翻译
	initialContent, _ := GetClipboard()
	m.lastContent = strings.TrimSpace(initialContent)

	m.enabled = true
	m.wg.Add(1)
	
	go m.monitor()
	
	return nil
}

// Stop stops monitoring the clipboard
func (m *Monitor) Stop() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.enabled {
		return fmt.Errorf("monitor not running")
	}

	m.enabled = false
	m.cancel()
	
	// Wait for monitor goroutine to finish
	m.wg.Wait()
	
	// Create new context for potential restart
	m.ctx, m.cancel = context.WithCancel(context.Background())
	
	return nil
}

// SetEnabled enables or disables the monitor
func (m *Monitor) SetEnabled(enabled bool) error {
	if enabled {
		return m.Start()
	}
	return m.Stop()
}

// IsEnabled returns whether the monitor is enabled
func (m *Monitor) IsEnabled() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.enabled
}

// SetAutoTranslate sets whether to auto-translate clipboard content
func (m *Monitor) SetAutoTranslate(enabled bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.autoTranslate = enabled
}

// SetOnChange sets the callback for clipboard changes
func (m *Monitor) SetOnChange(callback func(string)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.onChange = callback
}

// monitor is the main monitoring loop
func (m *Monitor) monitor() {
	defer m.wg.Done()
	
	ticker := time.NewTicker(m.checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-ticker.C:
			m.checkClipboard()
		}
	}
}

// checkClipboard checks for clipboard changes
func (m *Monitor) checkClipboard() {
	content, err := GetClipboard()
	if err != nil {
		// Silently ignore errors
		return
	}

	// Trim whitespace for comparison
	content = strings.TrimSpace(content)
	
	m.mu.RLock()
	lastContent := m.lastContent
	lastTranslation := m.lastTranslation
	onChange := m.onChange
	autoTranslate := m.autoTranslate
	m.mu.RUnlock()

	// Check if content changed and is not the last translation
	if content != "" && content != lastContent && content != lastTranslation {
		m.mu.Lock()
		m.lastContent = content
		m.mu.Unlock()

		// Call onChange callback if set and auto-translate is enabled
		if onChange != nil && autoTranslate {
			onChange(content)
		}
	}
}

// SetLastTranslation 设置最后的译文，避免重复翻译
func (m *Monitor) SetLastTranslation(translation string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lastTranslation = strings.TrimSpace(translation)
}

// GetLastContent returns the last clipboard content
func (m *Monitor) GetLastContent() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.lastContent
}

// ClearLastContent clears the last content
func (m *Monitor) ClearLastContent() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lastContent = ""
}