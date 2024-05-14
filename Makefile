# Change these variables as necessary.
MAIN_PACKAGE_PATH := ./
BINARY_NAME ?= monkey

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit: tidy
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest ./...
	go test -race -buildvcs -vet=off ./...


## vulncheck: check for known vulnerabilities
.PHONY: vulncheck
vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the binary
.PHONY: build
build:
	go build -o=./tmp/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## build/prod: build the stripped binary, to build for different platforms, use `env GOOS=linux GOARCH=amd64 make build/prod`
.PHONY: build/prod
build/prod:
	go build -tags netgo -ldflags '-s -w' -o=./${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  binary
.PHONY: run
run: build
	./tmp/${BINARY_NAME}
