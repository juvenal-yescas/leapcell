#!/bin/sh

apk add curl

curl -L https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64 -o /usr/bin/cloudflared
chmod +x /usr/bin/cloudflared

