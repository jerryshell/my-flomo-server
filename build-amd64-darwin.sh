#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o my-flomo-server-amd64-darwin
