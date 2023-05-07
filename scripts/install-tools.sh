#!/bin/sh

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
