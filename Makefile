GOPATH=$(shell cd ../../../.. && pwd)

VERSION=$(shell git describe --long --tags --dirty --always)

# Apply -dirty version suffix if there are staged or unstaged changes in ./builder
BUILDER_DIRTY=$(shell git diff-files --quiet -- "builder" && git diff-index --quiet --cached HEAD -- "builder" || echo "-dirty")
# Version builder by short commit sha of the builder dir, not the last probe version tag
BUILDER_VERSION=$(shell git rev-list -1 HEAD -- "builder" | cut -c1-7)${BUILDER_DIRTY}

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

build-cross:
	@echo "--> Building probe"
	gox -osarch="darwin/amd64" -osarch="linux/amd64" -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

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

.PHONY: builder
builder:
	@echo "--> Building builder: karlkfi/probe-builder:${BUILDER_VERSION}"
	@docker build -t karlkfi/probe-builder:${BUILDER_VERSION} ./builder

build-docker:
	@echo "--> Building probe (in karlkfi/probe-builder:${BUILDER_VERSION})"
	@docker run -v "$(shell pwd):/go/src/github.com/karlkfi/probe" karlkfi/probe-builder:${BUILDER_VERSION}

build-docker-cross:
	@echo "--> Building probe (in karlkfi/probe-builder:${BUILDER_VERSION}) for all platforms"
	@docker run -v "$(shell pwd):/go/src/github.com/karlkfi/probe" karlkfi/probe-builder:${BUILDER_VERSION} make restoredeps test build-cross