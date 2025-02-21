# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: install clean build test 
.PHONY: mockgen-install lint-prepare lint critic-prepare critic gosec-prepare gosec deadcode-prepare deadcode

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run
GOINSTALL = env GO111MODULE=on go install -v
GOTEST = env GO111MODULE=on go test -v
# Name of the Go binary output
BINARY_NAME=./bin/dkg

build:
	@echo "Building solana programs..."
	anchor build

build-all:
	@echo "Building solana all programs and clients..."
	cargo build

clean:
	@echo "Cleaning solana programs..."
	anchor clean

# Recipe to run tests
test:
	@echo "running tests"
	anchor test --skip-local-validator
