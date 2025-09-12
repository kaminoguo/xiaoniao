# Windows 控制台终极隐藏方案

## 问题描述
xiaoniao Windows版本在运行时，控制台窗口仍会在任务栏显示，影响用户体验。

## 解决方案

### 1. 编译时隐藏 (-H windowsgui)
```bash
go build -ldflags="-H windowsgui" ./cmd/xiaoniao
```
这个标志告诉Go编译器生成GUI应用程序而非控制台应用程序。

### 2. 运行时多层次隐藏

#### 方案实现位置
- 文件：`/home/lyrica/xiaoniao/cmd/xiaoniao/windows.go`
- 函数：`initializeHiddenConsole()`, `hideConsoleWindow()`, `hideFromTaskbarUltimate()`

#### 隐藏策略
1. **异步最小化**：`ShowWindowAsync(SW_FORCEMINIMIZE)` - 避免窗口闪烁
2. **窗口样式修改**：
   - 添加 `WS_EX_TOOLWINDOW` - 工具窗口，不在任务栏显示
   - 添加 `WS_EX_NOACTIVATE` - 不激活窗口
   - 移除 `WS_EX_APPWINDOW` - 移除应用窗口标志
   - 移除 `WS_VISIBLE` - 窗口不可见
3. **底层位置控制**：`SetWindowPos(HWND_BOTTOM, SWP_HIDEWINDOW)`
4. **完全隐藏**：`ShowWindow(SW_HIDE)`
5. **禁用交互**：`EnableWindow(FALSE)`

### 3. 启动时立即隐藏
```go
func runDaemonWithHotkey() {
    // 在程序启动的最早阶段就初始化隐藏的控制台
    initializeHiddenConsole()
    // ...
}
```

## 构建方法

使用高级构建脚本：
```bash
./build-windows-advanced.sh
```

构建特性：
- ✅ 终极控制台隐藏模式 (-H windowsgui)
- ✅ 多层次窗口样式隐藏
- ✅ 任务栏和Alt+Tab完全隐藏

## 测试验证

1. **任务栏检查**：确认xiaoniao不在任务栏显示
2. **Alt+Tab检查**：确认xiaoniao不在Alt+Tab列表中
3. **托盘功能检查**：确认只能通过托盘图标访问
4. **终端切换检查**：确认托盘菜单的"显示/隐藏终端"功能正常

## 恢复机制

当需要显示控制台时（如config命令），`showConsoleWindow()`会：
1. 重新启用窗口交互
2. 恢复正常窗口样式
3. 恢复到任务栏
4. 使用SetWindowPos显示窗口
5. 正常显示控制台

## 跨平台兼容性

非Windows平台通过`notwindows.go`提供空实现：
```go
func initializeHiddenConsole() {
    // No-op on non-Windows platforms
}
```

## 技术细节

### Windows API 调用
- `GetConsoleWindow()` - 获取控制台窗口句柄
- `ShowWindowAsync()` - 异步显示/隐藏窗口
- `SetWindowLongPtr()` - 修改窗口样式
- `SetWindowPos()` - 设置窗口位置和状态
- `EnableWindow()` - 启用/禁用窗口

### 常量定义
```go
const (
    WS_EX_TOOLWINDOW   = 0x00000080  // 工具窗口
    WS_EX_NOACTIVATE   = 0x08000000  // 不激活
    WS_EX_APPWINDOW    = 0x00040000  // 应用窗口
    WS_VISIBLE         = 0x10000000  // 可见
    SW_FORCEMINIMIZE   = 11          // 强制最小化
    SWP_HIDEWINDOW     = 0x0080      // 隐藏窗口
)
```

这个方案应该能够完全解决控制台窗口在任务栏显示的问题。