# PowerShell script to test the console hiding functionality
Write-Host "Testing xiaoniao console hiding functionality..." -ForegroundColor Green

# Set up minimal config
$configDir = "$env:APPDATA\xiaoniao"
$configFile = "$configDir\config.json"

if (!(Test-Path $configDir)) {
    New-Item -ItemType Directory -Path $configDir -Force
}

$config = @{
    api_key = "test-key-for-debug"
    provider = "OpenAI"
    model = "gpt-4o-mini"
    prompt_id = "direct"
} | ConvertTo-Json

Set-Content -Path $configFile -Value $config

Write-Host "Config created at: $configFile" -ForegroundColor Yellow

Write-Host ""
Write-Host "To test console hiding, run:" -ForegroundColor Cyan
Write-Host "  .\xiaoniao.exe run" -ForegroundColor White
Write-Host ""
Write-Host "Expected behavior:" -ForegroundColor Cyan
Write-Host "  1. Console window should not appear in taskbar" -ForegroundColor White
Write-Host "  2. Debug output should show invisible parent creation" -ForegroundColor White
Write-Host "  3. Console should be hidden but functional" -ForegroundColor White
Write-Host "  4. Can toggle console visibility with tray menu" -ForegroundColor White
Write-Host ""
Write-Host "Debug output should include:" -ForegroundColor Cyan
Write-Host "  - initializeHiddenConsole() called" -ForegroundColor White
Write-Host "  - Creating invisible parent window with handles" -ForegroundColor White
Write-Host "  - Console window detection and SetParent calls" -ForegroundColor White
Write-Host "  - Console hiding steps with success indicators" -ForegroundColor White
