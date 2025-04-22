#!/bin/bash
echo ""
echo "------ Server stopping at $(date) ------"

# Get the PID of the process using port 8080 (change if your server uses another port)
PID=$(lsof -ti :9090)

if [ -n "$PID" ]; then
    kill $PID
    echo "Server running on port 9090 stopped (PID: $PID)"
else
    echo "no server running on port 9090 found"
fi
