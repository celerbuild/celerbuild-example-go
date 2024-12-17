#!/bin/bash
# start.sh - Script to start the Go web server in the background

echo "Starting the Go web server..."

# Start the server and save PID
nohup ./server > server.log 2>&1 & echo $! > server.pid

# Wait for server to start
sleep 2

# Check if process is running
if ps -p $(cat server.pid) > /dev/null; then
    echo "Server started successfully! (PID: $(cat server.pid))"
    echo "Recent logs:"
    tail -n 5 server.log
else
    echo "Failed to start server!"
    echo "Error logs:"
    cat server.log
    exit 1
fi