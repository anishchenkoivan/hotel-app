all: api/swagger.json api/swagger.yaml

setup:
	go install github.com/swaggo/swag/cmd/swag@latest

api/swagger.%: api/api.go internal/app/handlers/handler.go
	swag init -g api/api.go -d . -o api

.PHONY: all setup
