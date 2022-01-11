#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o my-flomo-server-arm-linux
