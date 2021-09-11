echo Building Windows64
env GOOS=windows GOARCH=amd64 go build -o "bin/redditdl-win64.exe"

echo Building Windows32
env GOOS=windows GOARCH=386 go build -o "bin/redditdl-win32.exe"

echo Building Linux64
env GOOS=linux GOARCH=amd64 go build -o "bin/redditdl-linux64"

echo Building Linux32
env GOOS=linux GOARCH=386 go build -o "bin/redditdl-linux32"

echo Building LinuxARM
env GOOS=linux GOARCH=arm go build -o "bin/redditdl-linuxARM"

echo Building LinuxARM64
env GOOS=linux GOARCH=arm64 go build -o "bin/redditdl-linuxARM64"

echo Building Mac64
env GOOS=darwin GOARCH=amd64 go build -o "bin/redditdl-mac64"

echo Building Mac32
env GOOS=darwin GOARCH=amd64 go build -o "bin/redditdl-mac32"

echo Building MacARM64
env GOOS=darwin GOARCH=arm64 go build -o "bin/redditdl-macARM64"