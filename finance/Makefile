
.PHONY: integration.test test build lint deps 
gen:
	protoc --go_out=/pb/gen --go-grpc_out=pb/gen pb/finance.proto
server:
	go run cmd/server/main.go