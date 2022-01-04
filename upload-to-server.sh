#!/usr/bin/env bash
./build-linux-amd64.sh
rsync -aP my-flomo-server-linux-amd64 root@devenv.d8s.fun:/root/my-flomo-server/my-flomo-server-linux-amd64
