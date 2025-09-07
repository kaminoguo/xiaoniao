@echo off
chcp 65001 >nul
title xiaoniao

if not exist "%APPDATA%\xiaoniao\config.json" (
    echo First time setup, opening configuration...
    xiaoniao.exe config
    pause
)

xiaoniao.exe run