//go:build windows
// +build windows

package sound

import (
	"syscall"
	"unsafe"
)

var (
	winmm         = syscall.NewLazyDLL("winmm.dll")
	playSound     = winmm.NewProc("PlaySoundW")
	mciSendString = winmm.NewProc("mciSendStringW")
)

const (
	sndAsync    = 0x0001
	sndFilename = 0x00020000
)

// PlaySuccess plays a success sound on Windows
func PlaySuccess() {
	// Use Windows system sound
	text, _ := syscall.UTF16PtrFromString("SystemAsterisk")
	playSound.Call(uintptr(unsafe.Pointer(text)), 0, sndAsync)
}

// PlayError plays an error sound on Windows
func PlayError() {
	// Use Windows system sound
	text, _ := syscall.UTF16PtrFromString("SystemExclamation")
	playSound.Call(uintptr(unsafe.Pointer(text)), 0, sndAsync)
}

// PlayFile plays a sound file on Windows
func PlayFile(filepath string) error {
	text, err := syscall.UTF16PtrFromString(filepath)
	if err != nil {
		return err
	}
	
	ret, _, _ := playSound.Call(
		uintptr(unsafe.Pointer(text)),
		0,
		sndFilename|sndAsync,
	)
	
	if ret == 0 {
		return syscall.GetLastError()
	}
	return nil
}