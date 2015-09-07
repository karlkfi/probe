GOPATH=$(shell cd ../../../.. && pwd)

default: all

all: restoredeps test build

restoredeps:
	@echo "--> Restoring build dependencies"
	@godep restore

savedeps:
	@echo "--> Saving build dependencies"
	@godep save

updatedeps:
	@echo "--> Updating build dependencies"
	@godep update ${ARGS}

format:
	@echo "--> Running go fmt"
	@godep go fmt ./...

vet:
	@echo "--> Running go vet"
	@godep go vet ./...

build:
	@echo "--> Building probe"
	@godep go build -o probe

test_banner:
	@echo "--> Testing probe"

test: test_banner
	@godep go test ./...

test.v: test_banner
	@godep go test -test.v ./...

testrace:
	@godep go test -race ./...

clean:
	@echo "--> Cleaning probe"
	@godep go clean

env:
	@godep go env
