#!/bin/sh

nohup cloudflared tunnel --url http://localhost:8080 &

./app