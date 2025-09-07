#!/bin/bash

echo "Building xiaoniao..."

# Build for Linux
echo "Building Linux version..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao-linux cmd/xiaoniao/*.go

# Build for Windows
echo "Building Windows version..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o xiaoniao-windows.exe cmd/xiaoniao/*.go

echo "Build complete!"
echo ""
echo "Linux binary: xiaoniao-linux"
echo "Windows binary: xiaoniao-windows.exe"