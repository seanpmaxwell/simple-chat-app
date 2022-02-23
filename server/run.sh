#!/bin/bash


## Set the environment ##
env=${2:-development}

## Run go ##
if [ $1 = "start" ]; then
    go run ./src "$env"
elif [ $1 = "hot" ]; then
    gin --path ./src/ --port 9080 "$env"
elif [ $1 = "build" ]; then
    go build -o dist ./src
elif [ $1 = "startb" ]; then
    ./dist "$env"
elif [ $1 = "full" ]; then
    go build -o dist ./src && ./dist "$env"
elif [ $1 = "clean" ]; then
    go mod tidy
elif [ $1 = "wire" ]; then
    wire ./src/
else
    echo "Command not found"
fi
