#!/usr/bin/env bash
set -e
if [ -f "config.yml" ]; then
    cp config.yml config
fi
if [ -f "config.dev.yml" ]; then 
    cp config.dev.yml config
fi
version=$(git describe --tags)
buildTime=$(date '+%Y-%m-%d %H:%M:%S')
echo version:$version
echo buildTime:$buildTime
argsVersion="main.Version=$version"
argsBuildTime="main.BuildTime=$buildTime"
argsRunMode="talk.miniapp.mgtv.com/config.RunMode=product"

env GOOS=linux GOARCH=amd64 go build -o app -mod=vendor -ldflags="-X '$argsVersion' \
    -X '$argsBuildTime' \
    -X '$argsRunMode'"
targetServer="/home/project/jx3box-latest"
echo "copy to $targetServer"
scp app git.jx3box:$targetServer
rm app