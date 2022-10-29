#!/bin/bash
protoc -I ./ --go_out=./ --go-grpc_out=./ ./blogpb/blog.proto
