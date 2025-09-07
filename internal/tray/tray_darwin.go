//go:build darwin
// +build darwin

package tray

import (
	"fmt"
	"os/exec"
)

func ShowNotification(title, message string) {
	script := fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)
	exec.Command("osascript", "-e", script).Run()
}

func OpenConfigUI() {
	exec.Command("open", "-a", "Terminal", "--args", "xiaoniao", "config").Run()
}