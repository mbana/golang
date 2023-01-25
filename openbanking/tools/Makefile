SHELL 		:= bash
GIT_REV 	:= $(shell git rev-parse --short HEAD)
GOMAXPROCS 	:= 24
PARALLEL 	:= ${GOMAXPROCS}
MODULE_NAME := github.com/banaio/golang/openbanking_tools

.PHONY: all
all: format lint test

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	@printf "%b" "\033[92m" "  ---> Linting ... " "\033[0m" "\n"
	golangci-lint run --fix --config ./.golangci.yml ./...

.PHONY: format
format:
	@printf "%b" "\033[92m" "  ---> Formatting ... " "\033[0m" "\n"
	go mod tidy
	@for DIR in $(shell go list -f {{.Dir}} ./...); do \
		echo "\033[92m  ---> formatting dir=$${DIR} ... \033[0m"; \
		goimports -v -e -w -local ${MODULE_NAME} $${DIR}/*.go; \
	done
	@# find cmd -name "*.go" -exec goimports -w {} \;
	@# find generator -name "*.go" -exec goimports -w {} \;
	@# find scan -name "*.go" -exec goimports -w {} \;

.PHONY: test
test:
	@printf "%b" "\033[92m" "  ---> Testing ... " "\033[0m" "\n"
	GOMAXPROCS=${GOMAXPROCS} go test -parallel ${PARALLEL} -count=1 -race -v ./...

.PHONY: build
build:
	@printf "%b" "\033[92m" "  ---> Building ... " "\033[0m" "\n"
	go build -ldflags='-w -s' -installsuffix 'static' -o ./openbanking_tools ./cmd/openbanking_tools/main.go

.PHONY: build_image
build_image:
	@printf "%b" "\033[92m" "  ---> Building Image ... " "\033[0m" "\n"
	@printf "%b" "\033[93m" "WIP" "\033[0m" "\n"
	@# Doesn't work as I have no setup repository just yet.
	docker build -t "gcr.io/io-bana/openbanking_tools:latest" .

.PHONY: devtools
devtools:
	@printf "%b" "\033[92m" "  ---> Installing golangci-lint (https://github.com/golangci/golangci-lint) ... " "\033[0m" "\n"
	curl -sfL "https://install.goreleaser.com/github.com/golangci/golangci-lint.sh" | sh -s -- -b $(shell go env GOPATH)/bin v1.16.0

.PHONY: clean
clean:
	@printf "%b" "\033[92m" "  ---> Cleaning ... " "\033[0m" "\n"
	go clean -i -r -cache -testcache -modcache

.PHONY: init
init: ## initialise
	@printf "%b" "\033[92m" "  ---> Initialising ... " "\033[0m" "\n"
	go mod download
