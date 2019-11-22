#!/usr/bin/env bash

go build
#GOOS=windows GOARCH=386 go build
#GOOS=windows GOARCH=386 go build -o spring_block.exe main.go
GOOS=windows GOARCH=386 go build -o hello.exe main.go
./spring_block --verb True &
cd frontend
npm start &
cd ..

