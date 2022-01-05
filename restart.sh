#!/usr/bin/env bash
docker build . -t my-flomo-server
docker stop my-flomo-server
docker rm my-flomo-server
docker run -d --name my-flomo-server -p 8060:8060 my-flomo-server
