#!/bin/bash

protoc --proto_path=./ --go_out=plugins=grpc:./ ./test.proto
