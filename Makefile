.PHONY: build
build:
		go build -v ./cmd/go_rest_api_template

.DEFAULT_GOAL := build