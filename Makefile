SHELL := /bin/bash

build: generate tidy

generate:
	@echo "generate code"
	@go generate ./...
	@go fmt ./...
	@echo "ok"

test: generate
	go test ./... -v

tidy:
	@go mod tidy && go mod verify