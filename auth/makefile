.PHONY: build
build:
	go build -v -o a ./cmd/auth

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

.PHONY: start
start:
	go run ./cmd/auth

.DEFAULT_GOAL := build