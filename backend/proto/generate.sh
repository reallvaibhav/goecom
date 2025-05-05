#!/bin/bash

# Generate Go code for statistics service
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    statistics.proto 