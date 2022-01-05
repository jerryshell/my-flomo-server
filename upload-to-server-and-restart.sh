#!/usr/bin/env bash
./build-linux-amd64.sh
rsync -aP my-flomo-server-linux-amd64 Dockerfile restart.sh root@devenv.d8s.fun:/root/my-flomo-server/
rsync -aP config.prod.json root@devenv.d8s.fun:/root/my-flomo-server/config.json
ssh root@devenv.d8s.fun 'cd /root/my-flomo-server && ./restart.sh'
