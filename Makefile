.PHONY: all

all: format lint test

format:
	@echo "Formatting code"
	@go fmt ./...

lint:
	@echo "Linting code"
	@golangci-lint run

build:
	@echo "Building binary"
	@go build -o demo_api cmd/main.go

test: test_unit test_integration

test_unit:
	@echo "Running all unit tests"
	@go test -run="Unit" -v ./...

test_integration:
	@echo "Running all integration tests"
	@go test -run="Integration" -v ./...