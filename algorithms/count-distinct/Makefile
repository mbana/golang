SHELL:=/bin/sh
GIT_REV:=$(shell git rev-parse --short HEAD)
GOMAXPROCS:=8
PARALLEL:=${GOMAXPROCS}

.PHONY: test
test:
	@echo "\033[92m ---> test \033[0m"
	@GOMAXPROCS=${GOMAXPROCS} go test -count=1 -parallel=${GOMAXPROCS} ./...

.PHONY: devtools
devtools:
	@echo "\033[92m ---> installing golangci-lint (https://github.com/golangci/golangci-lint) ... \033[0m"
	curl -sfL "https://install.goreleaser.com/github.com/golangci/golangci-lint.sh" | sh -s -- -b $(shell go env GOPATH)/bin v1.12.5

.PHONY: lint
lint:
	@echo "\033[92m ---> lint \033[0m"
	golangci-lint run --config ./.golangci.yml ./...

.PHONY: generate_random
generate_random:
	@echo "\033[92m ---> generating random data to $(shell pwd)/countdistinct-elements.txt \033[0m"
	@for x in `seq 10`; do \
		echo $$(shuf -r -n10  /usr/share/dict/words); \
	 done > $(shell pwd)/countdistinct-elements.txt


# shuf -n1  /usr/share/dict/words

# .PHONY: generate_random
# generate_random:
# 	@echo "\033[92m ---> generating random data to $(shell pwd)/countdistinct-elements.txt \033[0m"
# 	@for x in `seq 10`; do \
# 		for y in `seq 64`; do \
# 			printf "%d" $$(( $$RANDOM % 2 )); \
# 		done; \
# 		echo; \
# 	 done > $(shell pwd)/countdistinct-elements.txt

# .PHONY: generate_random
# generate_random:
# 	@echo "\033[92m ---> generating random data to $(shell pwd)/countdistinct-elements.txt \033[0m"
# 	@for x in `seq 10`; do \
# 		echo "obase=16; ibase=2; $$(for y in `seq 32`; do \
# 			printf "%d" $$(( $$RANDOM % 2 )); \
# 		done)" | bc; \
# 		echo; \
# 	 done
