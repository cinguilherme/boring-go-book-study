.DEFAULT_GOAL := build

fmt:
	go fmt ./...

.PHONY: fmt

lint: fmt
	golint ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

ci: vet
	golangci-lint run
.PHONY: ci

build: ci
	go build -o hello_world main.go
.PHONY: build