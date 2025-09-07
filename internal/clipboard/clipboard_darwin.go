//go:build darwin
// +build darwin

package clipboard

import (
	"bytes"
	"os/exec"
	"strings"
)

type Monitor struct {
	enabled       bool
	onChange      func(string)
	lastContent   string
	autoTranslate bool
}

func NewMonitor() *Monitor {
	return &Monitor{
		enabled: true,
	}
}

func (m *Monitor) SetEnabled(enabled bool) {
	m.enabled = enabled
}

func (m *Monitor) IsEnabled() bool {
	return m.enabled
}

func (m *Monitor) SetOnChange(fn func(string)) {
	m.onChange = fn
}

func (m *Monitor) SetAutoTranslate(enabled bool) {
	m.autoTranslate = enabled
}

func (m *Monitor) IsAutoTranslate() bool {
	return m.autoTranslate
}

func (m *Monitor) GetLastContent() string {
	return m.lastContent
}

func Read() (string, error) {
	cmd := exec.Command("pbpaste")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func Write(text string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func Clear() error {
	return Write("")
}