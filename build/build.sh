#! /bin/bash
rm -rf release
mkdir release
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./release/fzuconnect_linux_arm -ldflags="-w -s"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./release/fzuconnect_win_amd64.exe -ldflags="-w -s"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./release/fzuconnect_linux_amd64 -ldflags="-w -s"