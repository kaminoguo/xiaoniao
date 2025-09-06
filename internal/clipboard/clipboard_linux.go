//go:build linux

package clipboard

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

// GetClipboard 获取剪贴板内容
func GetClipboard() (string, error) {
	// 尝试使用 xclip
	cmd := exec.Command("xclip", "-selection", "clipboard", "-o")
	var out bytes.Buffer
	cmd.Stdout = &out
	
	err := cmd.Run()
	if err != nil {
		// 如果 xclip 失败，尝试 xsel
		cmd = exec.Command("xsel", "-b", "-o")
		out.Reset()
		cmd.Stdout = &out
		
		err = cmd.Run()
		if err != nil {
			// 如果都失败，尝试 wl-paste (Wayland)
			cmd = exec.Command("wl-paste")
			out.Reset()
			cmd.Stdout = &out
			err = cmd.Run()
			if err != nil {
				return "", err
			}
		}
	}
	
	return strings.TrimSpace(out.String()), nil
}

// SetClipboard 设置剪贴板内容
func SetClipboard(text string) error {
	// 尝试使用 xclip
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = strings.NewReader(text)
	
	err := cmd.Run()
	if err != nil {
		// 如果 xclip 失败，尝试 xsel
		cmd = exec.Command("xsel", "-b", "-i")
		cmd.Stdin = strings.NewReader(text)
		
		err = cmd.Run()
		if err != nil {
			// 如果都失败，尝试 wl-copy (Wayland)
			cmd = exec.Command("wl-copy")
			cmd.Stdin = strings.NewReader(text)
			err = cmd.Run()
		}
	}
	
	// 小延迟确保剪贴板更新
	if err == nil {
		time.Sleep(50 * time.Millisecond)
	}
	
	return err
}

// ClearClipboard 清空剪贴板
func ClearClipboard() error {
	return SetClipboard("")
}