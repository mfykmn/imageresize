HAVE_GOLINT:=$(shell which golint)
HAVE_DEP:=$(shell which dep)
HAVE_ENUMER:=$(shell which enumer)

## Go
.PHONY: setup all-check lint vet test build
setup: dep enumer
	@echo "go setup"
	@dep ensure
	@go generate ./...

all-check: lint vet test
	@echo "all check"

lint: setup golint
	@echo "go lint"
	@golint $(shell go list ./...|grep -v vendor)

vet: setup
	@echo "go vet"
	@go vet ./...

test: setup
	@echo "go test"
	@go test -v $(shell go list ./... | grep -v /vendor/)

build: setup
	@echo "go build"
	@go build -ldflags "-X main.version=$(shell git describe --tags)" -o ./bin/imageresize ./cmd/imageresize

## Go package
.PHONY: dep golint enumer
dep:
ifndef HAVE_DEP
	@echo "Installing dep"
	@go get -u github.com/golang/dep/cmd/dep
endif

golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif

enumer:
ifndef HAVE_ENUMER
	@echo "Installing enumer"
	@go get -u github.com/alvaroloes/enumer
endif