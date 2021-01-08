#!/bin/bash
echo "generating proto buffer bindings"

GENERATED=$(protoc -I ./protos --go_out=plugins=grpc:pb ./protos/*.proto)


echo "$GENERATED"
