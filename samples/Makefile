SHELL 														:=/bin/bash
RESET															:= $(shell tput sgr0)
BOLD   														:= $(shell tput bold)
RED																:= $(shell tput bold; tput setaf 1)
GREEN															:= $(shell tput bold; tput setaf 2)
YELLOW														:= $(shell tput bold; tput setaf 3)

.PHONY: devtools
devtools:
	@echo "${GREEN}  ---> Installing golangci-lint (https://github.com/golangci/golangci-lint) ... ${RESET}"
	curl -sfL "https://install.goreleaser.com/github.com/golangci/golangci-lint.sh" | sh -s -- -b $(shell go env GOPATH)/bin v1.12.5

.PHONY: lint
lint:
	@for dir in */; do \
		( \
			CONFIG_PATH="$(shell pwd)/.golangci.yml"; \
			cd "$${dir}"; \
			echo "${GREEN} ---> linting ${RESET}"; \
			echo "${GREEN}            dir=$${dir} ${RESET}"; \
			echo "${GREEN}            module=$$(go list -m) ${RESET}"; \
			golangci-lint run --config "$${CONFIG_PATH}" ./...; \
		) \
	done

.PHONY: build
build:
	@for dir in */; do \
		( \
			cd "$${dir}"; \
			echo "${GREEN} ---> build ${RESET}"; \
			echo "${GREEN}            dir=$${dir} ${RESET}"; \
			echo "${GREEN}            module=$$(go list -m) ${RESET}"; \
			go build ./...; \
		) \
	done

.PHONY: test
test:
	@for dir in */; do \
		( \
			cd "$${dir}"; \
			echo "${GREEN} ---> test ${RESET}"; \
			echo "${GREEN}            dir=$${dir} ${RESET}"; \
			echo "${GREEN}            module=$$(go list -m) ${RESET}"; \
			go test -v -count=1 ./...; \
			go test -v -count=1 -parallel=32 ./...; \
			go test -v -count=1 -race ./...; \
			go test -v -count=1 -race -parallel=32 ./...; \
		) \
	done

.PHONY: run
run:
	@for dir in */; do \
		( \
			cd "$${dir}"; \
			echo "${GREEN} ---> build ${RESET}"; \
			echo "${GREEN}            dir=$${dir} ${RESET}"; \
			echo "${GREEN}            module=$$(go list -m) ${RESET}"; \
			MAIN=$$(find . -type f -name 'main.go'); \
			echo "main="$${MAIN}; \
			if [[ -f "$${MAIN}"  ]]; then \
				timeout --verbose --preserve-status --kill-after=2s 2s go run "$${MAIN}"; \
			else \
				timeout --verbose --preserve-status --kill-after=2s 2s go run main.go; \
			fi \
		) \
	done

# timeout --verbose --preserve-status --kill-after=2s 2s go run main.go; \
