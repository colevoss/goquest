.PHONY: test

default: all

all: test

test:
	go test -v ./...

test-cov:
	go test -coverprofile=coverage.out ./...

cov:
	go tool cover -html=coverage.out
