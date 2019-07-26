OUT_DIR ?= build
GOLANGCI_LINT_VERSION ?= v1.17.1

all: deps test lint build

init: githooks
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

githooks:
	git config core.hooksPath .githooks

out_dir:
	mkdir -p $(OUT_DIR)

test: out_dir
	go test ./... -tags 'release netgo osusergo' -v -coverprofile=$(OUT_DIR)/coverage.out
	go tool cover -func=$(OUT_DIR)/coverage.out

lint: out_dir
	golangci-lint run --exclude-use-default=false --fix -v

format:
	go fmt ./...

deps:
	go mod download
	go mod tidy
	go mod verify