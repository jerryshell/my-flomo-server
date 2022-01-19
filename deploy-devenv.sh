#!/usr/bin/env bash
make amd64-linux
rsync -aP my-flomo-server-amd64-linux Dockerfile restart.sh root@devenv.d8s.fun:/root/my-flomo-server/
rsync -aP config.prod.json root@devenv.d8s.fun:/root/my-flomo-server/config.json
ssh root@devenv.d8s.fun 'cd /root/my-flomo-server && ./restart.sh'
