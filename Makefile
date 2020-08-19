SHELL := /bin/bash

compile: vet
	# darwin binaries
	GOOS=darwin GOARCH=amd64 go build -o build/grpc-client-darwin-amd64 ./cmd/client
	GOOS=darwin GOARCH=amd64 go build -o build/grpc-server-darwin-amd64 ./cmd/server

	# linux binaries
	GOOS=linux GOARCH=amd64 go build -o build/grpc-client-linux-amd64 ./cmd/client
	GOOS=linux GOARCH=amd64 go build -o build/grpc-server-linux-amd64 ./cmd/server
proto:
	protoc -I internal/proto/ internal/proto/stream_example.proto --go_out=plugins=grpc:internal/proto
vet: proto fmt
	go vet ./...
fmt:
	go fmt ./...
