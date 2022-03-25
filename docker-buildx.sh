#!/usr/bin/env bash
docker buildx build --platform linux/arm,linux/amd64 . -t jerryshell/my-flomo-server:latest --push
