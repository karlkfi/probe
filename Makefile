GOPATH=$(shell cd ../../../.. && pwd)

VERSION=$(shell git describe --long --tags --dirty --always)

# Apply -dirty version suffix if there are staged or unstaged changes in ./builder
BUILDER_DIRTY=$(shell git diff-files --quiet -- "builder" && git diff-index --quiet --cached HEAD -- "builder" || echo "-dirty")
# Version builder by short commit sha of the builder dir, not the last probe version tag
BUILDER_VERSION=$(shell git rev-list -1 HEAD -- "builder" | cut -c1-7)${BUILDER_DIRTY}

default: all

all: test build

format:
	@echo "--> Running go fmt"
	@go fmt ./...

vet:
	@echo "--> Running go vet"
	@go vet ./...

build:
	@echo "--> Building probe"
	@go build -o probe

build-cross:
	@echo "--> Building probe"
	gox -osarch="darwin/amd64" -osarch="linux/amd64" -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

test_banner:
	@echo "--> Testing probe"

test: test_banner
	@go test ./...

test.v: test_banner
	@go test -test.v ./...

testrace:
	@go test -race ./...

clean:
	@echo "--> Cleaning probe"
	@go clean

env:
	@go env

.PHONY: builder
builder:
	@echo "--> Building builder: karlkfi/probe-builder:${BUILDER_VERSION}"
	@docker build -t karlkfi/probe-builder:${BUILDER_VERSION} ./builder

build-docker:
	@echo "--> Building probe (in karlkfi/probe-builder:${BUILDER_VERSION})"
	@docker run -v "$(shell pwd):/go/src/github.com/karlkfi/probe" karlkfi/probe-builder:${BUILDER_VERSION}

build-docker-cross:
	@echo "--> Building probe (in karlkfi/probe-builder:${BUILDER_VERSION}) for all platforms"
	@docker run -v "$(shell pwd):/go/src/github.com/karlkfi/probe" karlkfi/probe-builder:${BUILDER_VERSION} make test build-cross
