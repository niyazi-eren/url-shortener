#!/bin/bash

echo "Build script"

go build -tags netgo -ldflags '-s -w' -o app

go run .

