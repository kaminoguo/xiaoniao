//go:build darwin
// +build darwin

package sound

import (
	"os/exec"
)

func PlaySuccess() {
	exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Run()
}

func PlayError() {
	exec.Command("afplay", "/System/Library/Sounds/Basso.aiff").Run()
}

func PlayInfo() {
	exec.Command("afplay", "/System/Library/Sounds/Pop.aiff").Run()
}