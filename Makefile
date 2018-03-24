.PHONY: build doc fmt lint run test vendor_clean vendor_get vendor_update vet
# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
GOPATH := ${PWD}/_vendor:${GOPATH}
PROJECT_NAME := npuzzle
export GOPATH

default: build
build:
	go build -v -o ./bin/$(PROJECT_NAME) ./src/*.go
fmt:
	go fmt ./
lint:
	golint .

vet:
	go vet ./...
