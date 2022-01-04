#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o my-flomo-server-windows-amd64
