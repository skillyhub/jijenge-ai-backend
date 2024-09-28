.PHONY: integration.test test build lint deps gen server

# Target to generate Go code from protobuf definitions
gen:
	protoc --go_out=/pb/gen --go-grpc_out=pb/gen pb/business.proto

# Target to run the server
server:
	go run cmd/server/main.go

# Target to build the binary
build:
	go build -o bin/business ./cmd/server

# Add other targets as needed, such as test, lint, deps
