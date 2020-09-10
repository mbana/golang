SHELL:=/bin/sh
GIT_REV:=$(shell git rev-parse --short HEAD)
ENVS:=$(shell grep -v '^\#' .env)
GOMAXPROCS:=12
PARALLEL:=${GOMAXPROCS}

.PHONY: all
all: help

.PHONY: devtools
devtools:
	@echo "\033[92m  ---> Installing golangci-lint (https://github.com/golangci/golangci-lint) ... \033[0m"
	curl -sfL "https://install.goreleaser.com/github.com/golangci/golangci-lint.sh" | sh -s -- -b $(shell go env GOPATH)/bin v1.16.0

.PHONY: lint
lint:
	@echo "\033[92m  ---> Linting ... \033[0m"
	golangci-lint run --config ./.golangci.yml ./...

.PHONY: fomrat
format:
	@echo "\033[92m  ---> Formatting ... \033[0m"
	@for DIR in $(shell go list -f {{.Dir}} ./...); do \
		echo "\033[92m  ---> formatting dir=$${DIR} ... \033[0m"; \
		goimports -e -w -local github.com/banaio/golang/openbanking_forgerock $${DIR}/*.go; \
	done

.PHONY: init
init: ## initialise
	@echo -e "\033[92m  ---> initialising ... \033[0m"
	go mod download

.PHONY: test
test:  ## run the test suite
	@echo -e "\033[92m ---> test: running \033[0m"
	@GOMAXPROCS=${GOMAXPROCS} go test \
		-v \
		-race \
		-cover \
		-parallel ${PARALLEL} \
		./...

.PHONY: build
build: ## build the binary directly
	@echo -e "\033[92m  ---> building ... \033[0m"
	go build

.PHONY: build_image
build_image: ## build the docker image
	@echo -e "\033[92m  ---> building image ... \033[0m"
	docker build -t "gcr.io/www-bana-io/openbanking_forgerock:latest" .

.PHONY: run
run: ## run without race detector
	@$(ENVS) go run main.go

.PHONY: run_race
run_race: ## run with race detector
	@$(ENVS) go run -race main.go

.PHONY: ci_validate
ci_validate: ## validate the cloudbuild.yaml file
	@echo -e "\033[92m ---> ci: validating cloudbuild.yaml \033[0m"
	cloud-build-local --config=cloudbuild.yaml --dryrun=true .
	# cloud-build-local --config=cloudbuild.yaml --dryrun=true -substitutions SHORT_SHA=$(GIT_REV) .
	@# cloud-build-local --config=cloudbuild.yaml --dryrun=true .

.PHONY: ci
ci: ## run ci build locally
	@echo -e "\033[92m ---> ci: build locally \033[0m"
	cloud-build-local --config=cloudbuild.yaml --dryrun=false .
	# cloud-build-local --config=cloudbuild.yaml --dryrun=false -substitutions SHORT_SHA=$(GIT_REV) .
	@# cloud-build-local --config=cloudbuild.yaml --dryrun=false .

# TODO: remove as forgerock already provides this.
.PHONY: gen
gen: ## generate certificates
	@echo -e "\033[92m ---> generating certs \033[0m"

	rm -fr certs/signing/
	mkdir -p certs/signing/
	d certs/signing/ && $(ENVS) openssl req -new -newkey rsa:2048  -sha256 -nodes -out $${SSID}.csr -keyout $${SSID}.key -subj "/C=GB/O=mbana/OU=$${ORGANISATION_ID}/CN=$${SSID}"

	rm -fr certs/transport/
	mkdir -p certs/transport/
	cd certs/transport/ && $(ENVS) openssl req -new -newkey rsa:2048  -sha256 -nodes -out $${SSID}.csr -keyout $${SSID}.key -subj "/C=GB/O=mbana/OU=$${ORGANISATION_ID}/CN=$${SSID}"

.PHONY: help
help: ## print this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
