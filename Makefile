# Variables
GOHOSTOS:=$(shell go env GOHOSTOS)
GOHOSTARCH:=$(shell go env GOHOSTARCH)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BUILD_TIME=$(shell date +%Y-%m-%dT%H:%M:%S)
GIT_COMMIT=$(shell git rev-parse HEAD)

.PHONY: wire
wire:
	cd internal/di && wire

.PHONY: debug
debug:
	go run main.go -c configs/debug-config.yaml
