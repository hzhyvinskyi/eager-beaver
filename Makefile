.PHONY: build
build:
	go build -v ./cmd/eager-beaver

.DEFAULT_GOAL := build
