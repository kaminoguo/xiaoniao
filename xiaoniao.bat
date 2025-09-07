@echo off
chcp 65001 >nul
title xiaoniao

REM 检查配置文件是否存在
set CONFIG_DIR=%APPDATA%\xiaoniao
set CONFIG_FILE=%CONFIG_DIR%\config.json

REM 确保配置目录存在
if not exist "%CONFIG_DIR%" mkdir "%CONFIG_DIR%"

REM 如果没有配置文件，先打开配置界面
if not exist "%CONFIG_FILE%" (
    echo xiaoniao - Clipboard Translation Tool
    echo ======================================
    echo.
    echo First run detected. Opening configuration...
    echo.
    xiaoniao.exe config
    echo.
    echo Starting xiaoniao...
    timeout /t 2 >nul
)

REM 启动主程序
xiaoniao.exe run