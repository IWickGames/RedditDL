@echo off

echo Building Windows64
set GOOS=windows
set GOARCH=amd64
go build -o "bin/redditdl-win64.exe"

echo Building Windows32
set GOARCH=386
go build -o "bin/redditdl-win32.exe"

echo Building Linux64
set GOOS=linux
set GOARCH=amd64
go build -o "bin/redditdl-linux64"

echo Building Linux32
set GOARCH=386
go build -o "bin/redditdl-linux32"

echo Building LinuxARM
set GOARCH=arm
go build -o "bin/redditdl-linuxARM"

echo Building LinuxARM64
set GOARCH=arm64
go build -o "bin/redditdl-linuxARM64"

echo Building Mac64
set GOOS=darwin
set GOARCH=amd64
go build -o "bin/redditdl-mac64"

echo Building Mac32
set GOARCH=amd64
go build -o "bin/redditdl-mac32"

echo Building MacARM64
set GOARCH=arm64
go build -o "bin/redditdl-macARM64"