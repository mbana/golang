SHELL 		:= /bin/sh
GIT_REV 	:= $(shell git rev-parse --short HEAD)
GOMAXPROCS 	:= 12
PARALLEL 	:= ${GOMAXPROCS}

.PHONY: all
all: format lint test

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	@echo "\033[92m  ---> Linting ... \033[0m"
	golangci-lint run --fix --config ./.golangci.yml ./...

.PHONY: format
format:
	@echo "\033[92m  ---> Formatting ... \033[0m"
	@for DIR in $(shell go list -f {{.Dir}} ./...); do \
		echo "\033[92m  ---> formatting dir=$${DIR} ... \033[0m"; \
		goimports -e -w -local github.com/banaio/openbankingforgerock $${DIR}/*.go; \
	done

.PHONY: test
test:
	@echo "\033[92m ---> test: running \033[0m"
	@GOMAXPROCS=${GOMAXPROCS} go test -parallel ${PARALLEL} -count=1 \
		-v \
		-race \
		./...

.PHONY: build
build:
	@echo "\033[92m  ---> building ... \033[0m"
	go build ./...

.PHONY: build_image
build_image:
	@echo "\033[92m  ---> building image ... \033[0m"
	docker build -t "gcr.io/io-bana/openbankingforgerock:latest" .

.PHONY: devtools
devtools:
	@echo "\033[92m  ---> Installing golangci-lint (https://github.com/golangci/golangci-lint) ... \033[0m"
	curl -sfL "https://install.goreleaser.com/github.com/golangci/golangci-lint.sh" | sh -s -- -b $(shell go env GOPATH)/bin v1.16.0

.PHONY: clean
clean:
	@echo "\033[92m  ---> cleaning ... \033[0m"

.PHONY: init
init: ## initialise
	@echo "\033[92m  ---> initialising ... \033[0m"
	go mod download
