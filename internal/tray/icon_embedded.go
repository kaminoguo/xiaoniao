package tray

import (
	_ "embed"
	"runtime"
)

// Embed the default icons for different platforms and states

// PNG icons for Linux/macOS
//go:embed icon_blue.png
var iconBluePNG []byte

//go:embed icon_green.png
var iconGreenPNG []byte

//go:embed icon_red.png
var iconRedPNG []byte

// ICO icons for Windows
//go:embed icon_blue.ico
var iconBlueICO []byte

//go:embed icon_green.ico
var iconGreenICO []byte

//go:embed icon_red.ico
var iconRedICO []byte

// Fallback icon
//go:embed icon_default.png
var defaultIconPNG []byte

//go:embed icon_default.ico
var defaultIconICO []byte

// GetDefaultIcon returns the embedded default icon appropriate for the platform
func GetDefaultIcon() []byte {
	return GetIconForStatus("blue")
}

// GetIconForStatus returns the embedded icon for specific status
func GetIconForStatus(status string) []byte {
	if runtime.GOOS == "windows" {
		// Windows uses ICO format
		switch status {
		case "green":
			return iconGreenICO
		case "red":
			return iconRedICO
		case "blue":
			return iconBlueICO
		default:
			if len(defaultIconICO) > 0 {
				return defaultIconICO
			}
		}
	}
	
	// Linux/macOS use PNG format
	switch status {
	case "green":
		return iconGreenPNG
	case "red":
		return iconRedPNG
	case "blue":
		return iconBluePNG
	default:
		return defaultIconPNG
	}
}