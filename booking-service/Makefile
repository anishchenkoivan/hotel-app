all: api/swagger.json api/swagger.yaml api/api_v1pb/api_v1_grpc.pb.go  api/api_v1pb/api_v1.pb.go

setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/swaggo/swag/cmd/swag@latest

api/swagger.%: api/api.go internal/app/handlers/handler.go
	swag init -g api/api.go -d . -o api

api/api_v1pb/%: api/api_v1.proto
	protoc --go_out=. --go-grpc_out=. api/api_v1.proto

.PHONY: all setup
