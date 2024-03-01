#!/bin/bash

go run script/create_db/main.go

sleep 2

fresh

# Keep the container running
tail -f /dev/null