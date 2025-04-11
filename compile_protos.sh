#!/bin/bash

# This script generates source codes from .proto files
# using Protobuf compiler protoc.

protoc --go_out=./gen/golang \
  --go-grpc_out=./gen/golang \
  proto/**/*.proto
sleep 2