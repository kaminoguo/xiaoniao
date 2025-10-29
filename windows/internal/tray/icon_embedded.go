package tray

import (
	_ "embed"
)

// Embed the ICO icons for Windows

//go:embed icon_blue.ico
var iconBlueICO []byte

//go:embed icon_green.ico
var iconGreenICO []byte

//go:embed icon_red.ico
var iconRedICO []byte

//go:embed icon_default.ico
var defaultIconICO []byte

// GetDefaultIcon returns the embedded default icon
func GetDefaultIcon() []byte {
	return GetIconForStatus("blue")
}

// GetIconForStatus returns the embedded icon for specific status
func GetIconForStatus(status string) []byte {
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
		return iconBlueICO // fallback to blue icon
	}
}