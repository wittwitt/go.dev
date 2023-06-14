#!/bin/bash

rm -fr helloworld_grpc.pb.go

rm -fr helloworld.pb.go

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto