#!/usr/bin/env bash

echo "build macos x64 cvemonitor..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./cvemonitor_macos_x64 ./main.go
echo "build windows x64 cvemonitor..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./cvemonitor_windows_x64.exe ./main.go
echo "build linux x64 cvemonitor..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./cvemonitor_linux_x64 ./cvemonitor.go
echo "build macos arm64 cvemonitor..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -trimpath -o ./cvemonitor_macos_arm64 ./cvemonitor
