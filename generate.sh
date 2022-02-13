#!/usr/bin/env zsh


protoc protos/simple.proto --go_out=plugins=grpc:. --descriptor_set_out=simple.protoset
