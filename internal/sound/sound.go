package sound

import (
	"os/exec"
	"runtime"
)

// PlaySuccess 播放成功提示音
func PlaySuccess() {
	switch runtime.GOOS {
	case "linux":
		// 使用paplay播放系统提示音
		exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga").Run()
	case "darwin":
		// macOS使用afplay
		exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Run()
	case "windows":
		// Windows使用PowerShell
		exec.Command("powershell", "-c", "[console]::beep(800,200)").Run()
	}
}

// PlayError 播放错误提示音
func PlayError() {
	switch runtime.GOOS {
	case "linux":
		// 使用paplay播放系统错误音
		exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/dialog-error.oga").Run()
	case "darwin":
		// macOS使用afplay
		exec.Command("afplay", "/System/Library/Sounds/Basso.aiff").Run()
	case "windows":
		// Windows使用PowerShell
		exec.Command("powershell", "-c", "[console]::beep(400,500)").Run()
	}
}

// PlayStart 播放启动提示音
func PlayStart() {
	switch runtime.GOOS {
	case "linux":
		// 使用paplay播放系统启动音
		exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/service-login.oga").Run()
	case "darwin":
		// macOS使用afplay
		exec.Command("afplay", "/System/Library/Sounds/Hero.aiff").Run()
	case "windows":
		// Windows使用PowerShell
		exec.Command("powershell", "-c", "[console]::beep(600,200);[console]::beep(800,200)").Run()
	}
}

// Beep 简单的蜂鸣声（备用）
func Beep() {
	// 终端蜂鸣器作为备用
	print("\a")
}