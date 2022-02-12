#!/usr/bin/env zsh


protoc gen-test/simple.proto --go_out=plugins=grpc:. --descriptor_set_out=simple.protoset
