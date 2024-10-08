#!/bin/bash
# stop.sh - Script to gracefully stop the Go web server

PID_FILE="./server.pid"

if [ -f $PID_FILE ]; then
    PID=$(cat $PID_FILE)
    echo "Sending SIGTERM to the Go web server with PID $PID..."
    kill -SIGTERM $PID

    # Loop to check if the process still exists
    while kill -0 $PID 2>/dev/null; do
        echo "Waiting for process $PID to terminate..."
        sleep 1
    done

    echo "Process $PID has terminated."
    rm -f $PID_FILE
else
    echo "Server PID file not found. Is the server running?"
fi