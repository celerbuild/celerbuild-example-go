#!/bin/bash
# start.sh - Script to start the Go web server in the background

echo "Starting the Go web server..."

nohup ./server > server.log 2>&1 & echo $! > server.pid