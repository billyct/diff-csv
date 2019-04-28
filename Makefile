PWD := $(shell pwd)

build:
	@echo "Building diff-csv binary to './diff-csv'"
	@go build ./cmd/main.go
	@mv $(PWD)/main $(PWD)/diff-csv

build-win:
	@echo "Building diff-csv binary to './diff-csv.exe'"
	@env GOOS=windows GOARCH=amd64 go build ./cmd/main.go
	@mv $(PWD)/main.exe $(PWD)/diff-csv.exe