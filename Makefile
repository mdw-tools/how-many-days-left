#!/usr/bin/make -f

VERSION := $(shell git describe)

test:
	go fmt ./...
	go test -cover -count=1 -timeout=1s -race -v ./...

install:
	go install -ldflags="-X 'main.Version=$(VERSION)'"