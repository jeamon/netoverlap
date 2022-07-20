# Perform below commands
all: static-check test run

static-check:
	golangci-lint --tests=false run

## Clean package and fix linters warnings
lint:
	go mod tidy
	golangci-lint --fix --tests=false --timeout=2m30s run

## Run unit tests
test:
	go test -v ./... -count=1


## Obtain codebase testing coverage and view stats in console.
test-cover-console:
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

## Obtain codebase testing coverage and view stats in browser.
test-cover-html:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## Obtain codebase testing coverage and view stats on console and in browser.
test-cover: test-cover-console test-cover-html

## Run test and then build the program
build: test
	go build -o bin/netoverlap -a -ldflags "-extldflags '-static' -X 'main.GitCommit=$(git rev-parse --short HEAD)' -X 'main.GitTag=$(git describe --tags --abbrev=0)' -X 'main.BuildTime=$(date -u '+%Y-%m-%d %I:%M:%S %p GMT')'" main.go

## Run lint and test-unit commands
run:
	go run -ldflags "-X 'main.GitCommit=$(git rev-parse --short HEAD)' -X 'main.GitTag=$(git describe --tags --abbrev=0)' -X 'main.BuildTime=$(date -u '+%Y-%m-%d %I:%M:%S %p GMT')'" main.go

# Format codebase
format:
	gofumpt -l -w .