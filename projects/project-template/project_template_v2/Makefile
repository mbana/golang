# Control verbosity, example call:
# $ V=1 make lint
V 						?= 0
VERBOSE_FLAG 	=
QUIET 				= @
ifneq ($V,0)
	QUIET =
	VERBOSE_FLAG = -v
endif

.DEFAULT_GOAL						:= all
# .DELETE_ON_ERROR:
MAKEFLAGS 							+= --environment-overrides --warn-undefined-variables #--print-directory --no-builtin-rules --no-builtin-variables

# .ONESHELL:
SHELL										:= bash

GOLANGCI_LINT_VERSION		= v1.40.1
GOLANGCI_LINT_BINARY		= $(shell command -v golangci-lint 2> /dev/null)
GIT_SHORT_COMMIT_SHA		= $(shell git rev-parse --short HEAD)
IMAGE_NAME							= project_template_v2
PWD											= $(shell pwd)

export GO111MODULE			=	on

.PHONY: all
all: test lint

.PHONY: test
test:
	go test -race ${VERBOSE_FLAG} ./...

.PHONY: build
build:
	go build -ldflags="-w -s" -o main

.PHONY: build-binary-docker
build-binary-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint: lint-go

.PHONY: lint-go
lint-go:
	$(QUIET) printf "%s\n" "running linter='go mod tidy' ..."
	go mod tidy
	$(QUIET) if [[ ! -x "${GOLANGCI_LINT_BINARY}" ]]; then \
		printf "%s\n" "installing linter='${GOLANGCI_LINT_BINARY}@${GOLANGCI_LINT_VERSION}' ..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin ${GOLANGCI_LINT_VERSION}; \
		golangci-lint --version; \
	fi
	$(QUIET) printf "%s\n" "cleaning cache linter='${GOLANGCI_LINT_BINARY}' ..."
	golangci-lint cache clean
	$(QUIET) printf "%s\n" "running linter='${GOLANGCI_LINT_BINARY}' ..."
	golangci-lint run ${VERBOSE_FLAG} --config .golangci.yml ./...

.PHONY: build-docker-image
build-docker-image: build-binary-docker
	$(QUIET) printf "%s\n" "building image '${IMAGE_NAME}' ..."
	docker build \
		--tag ${IMAGE_NAME}:latest \
		--file ./Dockerfile \
		.
	$(QUIET) printf "%s\n" "running image '${IMAGE_NAME}' ..."
	docker run --rm -it --tty ${IMAGE_NAME}:latest

.PHONY: lint-go-fix
lint-go-fix:
	golangci-lint run ${VERBOSE_FLAG} --config .golangci.yml --fix ./...

.PHONY: pre-commit
pre-commit: lint-go-fix test

