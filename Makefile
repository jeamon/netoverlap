default: help

SHELL:=/usr/bin/env bash

GOLANGCI_LINT_VERSION:=1.52.0
# Download from https://github.com/golangci/golangci-lint/releases/tag/v1.52.2

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: linter
linter: ## Install golangci-lint executable via curl.
	## manual download from https://github.com/golangci/golangci-lint/releases/
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v${GOLANGCI_LINT_VERSION}

.PHONY: lint
lint: linter ## Run and fix linters warnings
	golangci-lint -v --fix --tests=false --timeout=3m run

.PHONY: test-cover
test-cover: ## Run unit tests and output coverage on console
	go test -v -race -count=1 -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

.PHONY: coverc
coverc: ## Obtain codebase testing coverage and view stats in console.
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

.PHONY: coverh
coverh: ## Obtain codebase testing coverage and view stats in browser.
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

.PHONY: coverage
coverage: test-cover-console test-cover-html ## Coverage and view stats on console and in browser.

.PHONY: test
build: test ## Run test and then build the program
	go build -o bin/netoverlap -a -ldflags "-extldflags '-static' -X 'main.GitCommit=$(git rev-parse --short HEAD)' -X 'main.GitTag=$(git describe --tags --abbrev=0)' -X 'main.BuildTime=$(date -u '+%Y-%m-%d %I:%M:%S %p GMT')'" main.go

.PHONY: run
run: ## Run lint and test-unit commands
	go run -ldflags "-X 'main.GitCommit=$(git rev-parse --short HEAD)' -X 'main.GitTag=$(git describe --tags --abbrev=0)' -X 'main.BuildTime=$(date -u '+%Y-%m-%d %I:%M:%S %p GMT')'" main.go

.PHONY: format
format: ## Format codebase
	gofumpt -l -w .

.PHONY: vuln
vuln: ## Run go vulnerability scanner.
	which govulncheck || go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...