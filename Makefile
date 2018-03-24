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

vendor_clean:
	rm -rf ./_vendor/src

# We have to set GOPATH to just the _vendor
# directory to ensure that `go get` doesn't
# update packages in our primary GOPATH instead.
# This will happen if you already have the package
# installed in GOPATH since `go get` will use
# that existing location as the destination.
vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor go get -d -u -v \
		github.com/kylelemons/godebug/pretty \

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet

vet:
	go vet ./...
