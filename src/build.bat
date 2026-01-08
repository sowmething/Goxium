@echo off
cd /d "%~dp0"

mkdir bin
go clean

echo make 64 bit
set GOARCH=amd64
go build -ldflags="-s -w" -o bin/goxiumx64.exe
copy goxium_ca.crt bin
copy goxium_ca.key bin
timeout 1 > nul

echo make 32 bit
set GOARCH=386
go build -ldflags="-s -w" -o bin/goxiumx86.exe
copy goxium_ca.crt bin
copy goxium_ca.key bin
timeout 1 > nul

echo make arm64
set GOARCH =arm64
go build -ldflags="-s -w" -o bin/goxiumarm64.exe
copy goxium_ca.crt bin
copy goxium_ca.key bin
timeout 1 > nul

echo done
