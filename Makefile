BINARY := api
PKG    := ./cmd/api

.PHONY: help run build test

help: ## Show this help
	@grep -E '^[a-z-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN{FS=":.*?## "}{printf "  %-8s %s\n", $$1, $$2}'

run: ## Run the API server
	go run $(PKG)

build: ## Build the binary into bin/
	go build -o bin/$(BINARY) $(PKG)

test: ## Run tests with the race detector
	go test -race ./...
