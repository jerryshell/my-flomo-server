#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my-flomo-server-linux-amd64
