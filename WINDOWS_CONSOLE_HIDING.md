# Windows 控制台终极隐藏方案

## 问题描述
xiaoniao Windows版本在运行时，控制台窗口仍会在任务栏显示，影响用户体验。

## 当前解决方案 - 隐形父窗口方案（v1.6.5+）

### 实现原理
利用Windows系统规则：子窗口不会在任务栏显示。通过创建一个不可见的父窗口，将控制台窗口设置为其子窗口，从而实现任务栏隐藏。

### 技术实现
- 文件：`/home/lyrica/xiaoniao/cmd/xiaoniao/windows.go`
- 核心函数：
  - `createInvisibleParentWindow()` - 创建隐形父窗口
  - `initializeHiddenConsole()` - 初始化隐藏控制台
  - `showConsoleWindow()` - 显示时临时解除父子关系
  - `hideConsoleWindow()` - 隐藏时重新建立父子关系

### 实现步骤
1. **注册窗口类**：使用 `RegisterClassEx` 注册自定义窗口类
2. **创建隐形父窗口**：`CreateWindowEx` 创建 `WS_POPUP` 样式窗口（不在任务栏）
3. **设置父子关系**：`SetParent` 将控制台设为隐形窗口的子窗口
4. **动态切换**：
   - 显示控制台时：临时解除父子关系，允许正常显示
   - 隐藏控制台时：重新建立父子关系，从任务栏消失

### 3. 启动时立即隐藏
```go
func runDaemonWithHotkey() {
    // 在程序启动的最早阶段就初始化隐藏的控制台
    initializeHiddenConsole()
    // ...
}
```

## 构建方法

使用标准构建脚本：
```bash
./build-windows.sh
# 或
./build-windows-advanced.sh
```

**注意**：不再使用 `-H windowsgui` 标志，保持为控制台程序以确保终端功能正常。

构建特性：
- ✅ 隐形父窗口实现任务栏隐藏
- ✅ 保持控制台功能完整性
- ✅ 任务栏和Alt+Tab完全隐藏
- ✅ 托盘菜单控制显示/隐藏

## 测试验证

1. **任务栏检查**：确认xiaoniao不在任务栏显示
2. **Alt+Tab检查**：确认xiaoniao不在Alt+Tab列表中
3. **托盘功能检查**：确认只能通过托盘图标访问
4. **终端切换检查**：确认托盘菜单的"显示/隐藏终端"功能正常

## 恢复机制

当需要显示控制台时（通过托盘菜单），`showConsoleWindow()`会：
1. 临时解除父子关系（`SetParent(console, 0)`）
2. 显示控制台窗口（`ShowWindow(SW_SHOW)`）
3. 控制台可正常使用和交互
4. 隐藏时重新设置父子关系

## 跨平台兼容性

非Windows平台通过`notwindows.go`提供空实现：
```go
func initializeHiddenConsole() {
    // No-op on non-Windows platforms
}
```

## 技术细节

### 新增 Windows API 调用（v1.6.5+）
- `RegisterClassEx()` - 注册窗口类
- `CreateWindowEx()` - 创建隐形父窗口
- `GetModuleHandle()` - 获取模块句柄
- `DefWindowProc()` - 默认窗口过程
- `SetParent()` - 设置父子窗口关系
- `GetConsoleWindow()` - 获取控制台窗口句柄
- `ShowWindow()` - 显示/隐藏窗口

### 关键常量定义
```go
const (
    WS_POPUP         = 0x80000000  // 弹出窗口（不在任务栏）
    CS_HREDRAW       = 0x0002      // 水平重绘
    CS_VREDRAW       = 0x0001      // 垂直重绘
    COLOR_WINDOW     = 5           // 窗口背景色
    SW_HIDE          = 0           // 隐藏窗口
    SW_SHOW          = 5           // 显示窗口
)
```

## 历史方案记录

### 旧方案问题
- **-H windowsgui**：导致控制台功能完全丢失，出现黑屏
- **WS_EX_TOOLWINDOW**：只能从Alt+Tab隐藏，无法从任务栏隐藏
- **窗口样式修改**：效果有限，无法完全隐藏任务栏图标

### 当前方案优势
- ✅ 利用Windows系统规则，可靠性高
- ✅ 保持控制台功能完整
- ✅ 完全从任务栏和Alt+Tab隐藏
- ✅ 代码实现简洁，易于维护