.PHONY: build
build:
	@go build -ldflags='-s -w' -o prae .
