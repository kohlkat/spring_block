#!/usr/bin/env bash

env GOOS="linux" GOARCH="amd64" go build -v
chmod +x spring_block

scp -i "key.pem" spring_block peduzzi_gaspard@35.234.77.98:/home/peduzzi_gaspard/bot/
