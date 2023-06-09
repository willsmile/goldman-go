export GO111MODULE=on
BIN := goldman-go
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: build
build: $(BIN)

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

$(BIN): $(SOURCES)
	go build -o $@
