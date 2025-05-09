#!/bin/bash
echo ""
echo "------ Server stopping at $(date) ------"

# Get the PID of the process using port 8080 (change if your server uses another port)
PID=$(lsof -ti :8080)

if [ -n "$PID" ]; then
    kill $PID
    echo "Server running on port 8080 stopped (PID: $PID)"
else
    echo "o server running on port 8080 found"
fi
