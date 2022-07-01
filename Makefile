SHELL := /bin/bash

all: build

.PHONY: build
build: build-go

.PHONY: build-go
build-go:
	mkdir -p ./bin
	go build -o bin
	cp -r ./bin/* /tmp

.PHONY: fmt
fmt: fmt-go

.PHONY: fmt-go
fmt-go:
	go fmt ./...

.PHONY: init
init:
	go mod tidy
	go mod vendor
