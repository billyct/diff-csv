PWD := $(shell pwd)

build:
	@echo "Building diff-csv binary to './diff-csv'"
	@go build ./cmd/main.go
	@mv $(PWD)/main $(PWD)/diff-csv