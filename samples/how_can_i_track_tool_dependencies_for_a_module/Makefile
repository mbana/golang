SHELL:=/bin/bash

GO_PKGS=$(shell go list ./...)
GO_PKGS_FOLDERS=$(shell go list -f '{{.Dir}}/' ./...)

.PHONY: all
all: init devtools lint cyclomatic linter

.PHONY: init
init:
	@echo -e "\033[92m  ---> Initialising ... \033[0m"
	go mod download

.PHONY: devtools
devtools:
	@# 1. Get all imports from the `tools/tools.go` file.
	@# 2. Install all of them.
	@TOOLS_PKGS="$(shell go list -f '{{ range $$index, $$import := .Imports }} \
		{{ if $$index }} \
			{{ "\n" }} \
		{{ end }} \
		{{ $$import }} \
		{{ end }}' tools/tools.go)"; \
		for TOOL_PKG in $${TOOLS_PKGS}; do \
			echo -e "\033[92m  ---> Installing $${TOOL_PKG} ... \033[0m"; \
			go install $${TOOL_PKG}; \
		done

.PHONY: lint
lint:
	@echo -e "\033[92m  ---> Vetting ... \033[0m"
	go vet ${GO_PKGS}
	@echo -e "\033[92m  ---> Linting ... \033[0m"
	golint -min_confidence 1.0 -set_exit_status ${GO_PKGS}
	@echo -e "\033[92m  ---> Formatting ... \033[0m"
	@for GO_PKG_DIR in ${GO_PKGS_FOLDERS}; do \
		echo -e "\033[92m  ---> Formatting $${GO_PKG_DIR}*.go ... \033[0m"; \
		gofmt -e -s -w $${GO_PKG_DIR}*.go; \
	done

.PHONY: cyclomatic
cyclomatic:
	@echo -e "\033[92m  ---> Checking cyclomatic complexity ... \033[0m"
	gocyclo -over 12 ${GO_PKGS_FOLDERS}

.PHONY: linter
linter:
	@echo -e "\033[92m  ---> Checking other qa tools ... \033[0m"
	golangci-lint run --no-config --issues-exit-code=0 --deadline=1m \
  		--disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
  		--enable=structcheck --enable=maligned --enable=errcheck --enable=dupl --enable=ineffassign \
  		--enable=interfacer --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck ./...
