#!/bin/sh

cat /etc/os-release

apk add curl

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        CLOUDFLARED_ARCH="amd64"
        ;;
    aarch64|arm64)
        CLOUDFLARED_ARCH="arm64"
        ;;
    armv7l|armhf)
        CLOUDFLARED_ARCH="arm"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

echo "Detected architecture: $ARCH, downloading cloudflared-linux-$CLOUDFLARED_ARCH"
curl -L https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-$CLOUDFLARED_ARCH -o /usr/bin/cloudflared
chmod +x /usr/bin/cloudflared

