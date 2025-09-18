//go:build darwin
// +build darwin

package sound

import (
	"os/exec"
)

// PlaySuccess plays a success sound on macOS
func PlaySuccess() error {
	// Use the system Glass sound
	return exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Run()
}

// PlayError plays an error sound on macOS
func PlayError() error {
	// Use the system Funk sound for errors
	return exec.Command("afplay", "/System/Library/Sounds/Funk.aiff").Run()
}

// PlayFile plays a custom sound file on macOS
func PlayFile(filepath string) error {
	return exec.Command("afplay", filepath).Run()
}

// Beep produces a simple beep sound
func Beep() {
	// Try to play a system sound, fallback to terminal beep
	if err := exec.Command("afplay", "/System/Library/Sounds/Tink.aiff").Run(); err != nil {
		// Fallback to terminal beep
		print("\a")
	}
}