.PHONY: build
build:
	go build -v -o a ./cmd/authO2

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

.PHONY: start
start:
	go run ./cmd/authO2

.DEFAULT_GOAL := build