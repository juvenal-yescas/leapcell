#!/bin/sh
nohup ./app &

cloudflared tunnel --url http://localhost:8080
