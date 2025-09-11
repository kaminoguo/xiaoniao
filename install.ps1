# Windows PowerShell 安装脚本
# xiaoniao Windows Installation Script

param(
    [string]$InstallPath = "$env:ProgramFiles\xiaoniao",
    [switch]$CreateShortcut = $true,
    [switch]$AddToStartup = $false,
    [switch]$AddToPath = $false
)

# 颜色输出函数
function Write-ColorOutput {
    param([string]$Message, [string]$Color = "White")
    $colorMap = @{
        "Red" = [ConsoleColor]::Red
        "Green" = [ConsoleColor]::Green  
        "Yellow" = [ConsoleColor]::Yellow
        "Blue" = [ConsoleColor]::Blue
        "Cyan" = [ConsoleColor]::Cyan
        "White" = [ConsoleColor]::White
    }
    Write-Host $Message -ForegroundColor $colorMap[$Color]
}

Write-ColorOutput "🚀 xiaoniao Windows 安装程序" "Blue"
Write-ColorOutput "============================================" "Blue"

# 检查管理员权限
$currentUser = [Security.Principal.WindowsIdentity]::GetCurrent()
$principal = New-Object Security.Principal.WindowsPrincipal($currentUser)
$isAdmin = $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

if (-not $isAdmin -and $InstallPath.StartsWith($env:ProgramFiles)) {
    Write-ColorOutput "⚠️ 需要管理员权限安装到 Program Files" "Yellow"
    Write-ColorOutput "请使用管理员权限运行 PowerShell，或选择用户目录安装:" "Yellow"
    Write-ColorOutput ".\install.ps1 -InstallPath `$env:LOCALAPPDATA\xiaoniao" "Cyan"
    exit 1
}

# 检查xiaoniao.exe是否存在
if (-not (Test-Path "xiaoniao.exe")) {
    Write-ColorOutput "❌ xiaoniao.exe 未找到" "Red"
    Write-ColorOutput "请确保在包含 xiaoniao.exe 的目录中运行此脚本" "Red"
    exit 1
}

try {
    # 创建安装目录
    Write-ColorOutput "📁 创建安装目录: $InstallPath" "Blue"
    if (-not (Test-Path $InstallPath)) {
        New-Item -ItemType Directory -Force -Path $InstallPath | Out-Null
    }

    # 复制可执行文件
    Write-ColorOutput "📋 复制 xiaoniao.exe 到安装目录..." "Blue"
    Copy-Item "xiaoniao.exe" -Destination $InstallPath -Force

    # 验证安装
    $installedExe = Join-Path $InstallPath "xiaoniao.exe"
    if (Test-Path $installedExe) {
        Write-ColorOutput "✅ xiaoniao.exe 安装成功" "Green"
    } else {
        throw "安装失败：无法复制文件"
    }

    # 创建开始菜单快捷方式
    if ($CreateShortcut) {
        Write-ColorOutput "🔗 创建开始菜单快捷方式..." "Blue"
        
        $startMenuPath = "$env:APPDATA\Microsoft\Windows\Start Menu\Programs"
        $shortcutPath = Join-Path $startMenuPath "xiaoniao.lnk"
        
        $WshShell = New-Object -comObject WScript.Shell
        $Shortcut = $WshShell.CreateShortcut($shortcutPath)
        $Shortcut.TargetPath = $installedExe
        $Shortcut.WorkingDirectory = $InstallPath
        $Shortcut.Description = "xiaoniao - AI 翻译助手"
        $Shortcut.WindowStyle = 7  # Minimized
        $Shortcut.Save()
        
        Write-ColorOutput "✅ 快捷方式已创建" "Green"
    }

    # 添加到系统启动项
    if ($AddToStartup) {
        Write-ColorOutput "🔄 添加到系统启动项..." "Blue"
        
        $startupPath = "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup"
        $startupShortcut = Join-Path $startupPath "xiaoniao.lnk"
        
        $WshShell = New-Object -comObject WScript.Shell
        $Shortcut = $WshShell.CreateShortcut($startupShortcut)
        $Shortcut.TargetPath = $installedExe
        $Shortcut.Arguments = "run"
        $Shortcut.WorkingDirectory = $InstallPath
        $Shortcut.Description = "xiaoniao - AI 翻译助手 (启动时运行)"
        $Shortcut.WindowStyle = 7  # Minimized
        $Shortcut.Save()
        
        Write-ColorOutput "✅ 已添加到启动项" "Green"
    }

    # 添加到系统PATH
    if ($AddToPath) {
        Write-ColorOutput "🛣️ 添加到系统 PATH..." "Blue"
        
        $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
        if ($currentPath -notlike "*$InstallPath*") {
            $newPath = $currentPath + ";" + $InstallPath
            [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
            Write-ColorOutput "✅ 已添加到用户 PATH" "Green"
            Write-ColorOutput "重启终端后可直接使用 'xiaoniao' 命令" "Yellow"
        } else {
            Write-ColorOutput "ℹ️ 安装路径已在 PATH 中" "Cyan"
        }
    }

    # 创建卸载脚本
    Write-ColorOutput "📝 创建卸载脚本..." "Blue"
    $uninstallScript = @"
# xiaoniao 卸载脚本
Write-Host "🗑️ 卸载 xiaoniao..." -ForegroundColor Blue

# 停止运行的进程
Get-Process xiaoniao -ErrorAction SilentlyContinue | Stop-Process -Force

# 删除安装文件
if (Test-Path "$InstallPath") {
    Remove-Item "$InstallPath" -Recurse -Force
    Write-Host "✅ 安装文件已删除" -ForegroundColor Green
}

# 删除快捷方式
`$shortcuts = @(
    "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\xiaoniao.lnk",
    "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup\xiaoniao.lnk"
)

foreach (`$shortcut in `$shortcuts) {
    if (Test-Path `$shortcut) {
        Remove-Item `$shortcut -Force
    }
}

# 从PATH中移除
`$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (`$currentPath -like "*$InstallPath*") {
    `$newPath = `$currentPath -replace [regex]::Escape("$InstallPath" + ";"), ""
    `$newPath = `$newPath -replace [regex]::Escape(";" + "$InstallPath"), ""
    `$newPath = `$newPath -replace [regex]::Escape("$InstallPath"), ""
    [Environment]::SetEnvironmentVariable("Path", `$newPath, "User")
}

Write-Host "✅ xiaoniao 卸载完成" -ForegroundColor Green
Write-Host "配置文件保留在: `$env:APPDATA\xiaoniao" -ForegroundColor Yellow
"@

    $uninstallPath = Join-Path $InstallPath "uninstall.ps1"
    $uninstallScript | Out-File -FilePath $uninstallPath -Encoding UTF8
    Write-ColorOutput "✅ 卸载脚本已创建: $uninstallPath" "Green"

    # 获取版本信息
    $versionOutput = & $installedExe "version" 2>$null
    if ($versionOutput) {
        Write-ColorOutput "📋 已安装版本: $versionOutput" "Cyan"
    }

    Write-ColorOutput "" "White"
    Write-ColorOutput "🎉 xiaoniao 安装完成！" "Green"
    Write-ColorOutput "============================================" "Green"
    Write-ColorOutput "安装位置: $InstallPath" "Cyan"
    
    if ($CreateShortcut) {
        Write-ColorOutput "开始菜单: 搜索 'xiaoniao'" "Cyan"
    }
    
    Write-ColorOutput "" "White"
    Write-ColorOutput "🚀 使用方法:" "Blue"
    Write-ColorOutput "  1. 配置: xiaoniao config" "White"
    Write-ColorOutput "  2. 运行: xiaoniao run" "White"
    Write-ColorOutput "  3. 帮助: xiaoniao help" "White"
    Write-ColorOutput "" "White"
    Write-ColorOutput "💡 首次使用请运行配置命令设置 API 密钥" "Yellow"
    
    # 询问是否立即运行配置
    Write-ColorOutput "❓ 现在打开配置界面吗? (y/N): " "Yellow" -NoNewline
    $response = Read-Host
    if ($response -eq "y" -or $response -eq "Y" -or $response -eq "yes") {
        Write-ColorOutput "🔧 启动配置界面..." "Blue"
        & $installedExe "config"
    }

} catch {
    Write-ColorOutput "❌ 安装失败: $($_.Exception.Message)" "Red"
    exit 1
}

Write-ColorOutput "" "White"
Write-ColorOutput "📖 更多信息请访问: https://github.com/kaminoguo/xiaoniao" "Cyan"