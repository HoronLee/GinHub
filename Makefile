# Variables
GOHOSTOS:=$(shell go env GOHOSTOS)
GOHOSTARCH:=$(shell go env GOHOSTARCH)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BUILD_TIME=$(shell date +%Y-%m-%dT%H:%M:%S)
GIT_COMMIT=$(shell git rev-parse HEAD)

.PHONY: deps
deps:
	go mod tidy

.PHONY: wire
wire:
	cd internal/di && wire

.PHONY: swagger
swagger:
	swag init -g internal/server/http.go -o internal/swagger --parseDependency --parseInternal --exclude example

.PHONY: swagger-install
swagger-install:
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: build
build: swagger wire
	go build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)" -o bin/ginhub main.go

.PHONY: dev
dev: swagger
	go run main.go serve -c configs/debug-config.yaml

.PHONY: prod
prod: swagger
	go run main.go serve -c configs/production-config.yaml

.PHONY: build-prod
build-prod: swagger wire
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)" -o bin/ginhub-prod main.go

.PHONY: clean
clean:
	rm -rf bin/
	rm -rf internal/swagger/docs.go
	rm -rf internal/swagger/swagger.json
	rm -rf internal/swagger/swagger.yaml
