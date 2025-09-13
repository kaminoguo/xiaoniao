package logbuffer

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	maxBufferSize = 10 * 1024 * 1024 // 10MB 缓冲区大小限制
)

// LogBuffer 日志缓冲区
type LogBuffer struct {
	buffer bytes.Buffer
	mutex  sync.RWMutex
}

// globalBuffer 全局日志缓冲区实例
var globalBuffer *LogBuffer
var once sync.Once
var originalStdout *os.File

// GetInstance 获取全局日志缓冲区实例
func GetInstance() *LogBuffer {
	once.Do(func() {
		globalBuffer = &LogBuffer{}
		originalStdout = os.Stdout
	})
	return globalBuffer
}

// Write 实现 io.Writer 接口，捕获所有输出
func (lb *LogBuffer) Write(p []byte) (n int, err error) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	
	// 写入缓冲区
	n, err = lb.buffer.Write(p)
	
	// 如果缓冲区太大，保留最后的部分
	if lb.buffer.Len() > maxBufferSize {
		content := lb.buffer.Bytes()
		// 保留最后的 80% 内容
		keepSize := maxBufferSize * 8 / 10
		if len(content) > keepSize {
			lb.buffer.Reset()
			lb.buffer.Write(content[len(content)-keepSize:])
		}
	}
	
	// 同时输出到原始控制台
	if originalStdout != nil {
		originalStdout.Write(p)
	}
	
	return n, err
}

// ExportToFile 导出日志到文件并自动打开
func (lb *LogBuffer) ExportToFile() (string, error) {
	lb.mutex.RLock()
	content := lb.buffer.String()
	lb.mutex.RUnlock()
	
	// 生成文件名，使用导出时间
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("xiaoniao_log_%s.txt", timestamp)
	
	// 获取exe所在目录
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("获取程序路径失败: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	
	// 完整文件路径
	filePath := filepath.Join(exeDir, filename)
	
	// 创建文件并写入内容
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("写入日志文件失败: %v", err)
	}
	
	// 自动打开文件
	if err := OpenFileWithDefaultEditor(filePath); err != nil {
		// 打开失败不影响导出成功
		return filePath, nil
	}
	
	return filePath, nil
}

// OpenFileWithDefaultEditor 使用系统默认编辑器打开文件
func OpenFileWithDefaultEditor(filePath string) error {
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
	case "windows":
		// Windows: 使用默认程序打开
		cmd = exec.Command("cmd", "/c", "start", "", filePath)
	case "darwin":
		// macOS
		cmd = exec.Command("open", filePath)
	case "linux":
		// Linux
		cmd = exec.Command("xdg-open", filePath)
	default:
		return fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}
	
	return cmd.Start()
}

// Clear 清空日志缓冲区
func (lb *LogBuffer) Clear() {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	lb.buffer.Reset()
}

// CaptureStdout 开始捕获标准输出
func CaptureStdout() {
	lb := GetInstance()
	
	// 创建管道
	r, w, _ := os.Pipe()
	
	// 保存原始stdout
	originalStdout = os.Stdout
	
	// 重定向stdout到管道写入端
	os.Stdout = w
	
	// 启动goroutine读取管道并写入缓冲区
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := r.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				lb.Write(buf[:n])
			}
		}
	}()
}

// ExportLogs 导出全局缓冲区的日志
func ExportLogs() (string, error) {
	return GetInstance().ExportToFile()
}