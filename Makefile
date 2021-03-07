.PHONY: build
build:
	@rm -rf bin
	@mkdir bin
	@go build -ldflags "-s -w" -o bin/server server/main.go
	@go build -ldflags "-s -w" -o bin/client client/main.go
