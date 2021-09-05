# Tooling
lint:
	golint ./...

vet:
	go vet ./...

ci:
	golangci-lint

pre-cmt:
	go fmt;
	goimports -l -w .; 
	make lint && make vet

# CMD's
run:
	go run main.go

build:
	go build -o hello_world main.go