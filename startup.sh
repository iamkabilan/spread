#!/bin/bash

# Start the server in the background
go run cmd/server/main.go &
SERVER_PID=$!

# Spin up the storage nodes in the background
go run cmd/node/main.go -port=3000 &
NODE1_PID=$!

go run cmd/node/main.go -port=3001 &
NODE2_PID=$!

# Wait for all processes to finish
wait $SERVER_PID
wait $NODE1_PID
wait $NODE2_PID