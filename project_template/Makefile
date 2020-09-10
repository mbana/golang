# https://gist.github.com/rsperl/d2dfe88a520968fbc1f49db0a29345b9
# http://www.lunderberg.com/2015/08/25/cpp-makefile-pretty-output/
# http://agdr.org/2020/05/14/Polyglot-Makefiles.html
# https://tech.davis-hansson.com/p/make/

export BASHOPTS													:= extglob:globstar:nullglob:failglob:gnu_errfmt:localvar_unset:dotglob:xpg_echo:functrace:verbose
export SHELLOPTS												:= allexport:braceexpand:emacs:errexit:errtrace:hashall:ignoreeof:interactive-comments:keyword:monitor:noclobber:noglob:nolog:notify:nounset:onecmd:physical:pipefail:posix:vi
export TERM															?= xterm-256color
# export SHELLCHECK_OPTS								?= --shell=bash --check-sourced --external-sources

.DEFAULT_GOAL														:= all
# .DELETE_ON_ERROR:
MAKEFLAGS 															+= --environment-overrides --warn-undefined-variables #--no-builtin-rules --no-builtin-variables #--print-directory

# .ONESHELL:
SHELL																		:= bash
# IFS																		= $'\n\t'

export SCREEN_RESET											:= $(shell tput reset)
export SCREEN_CLEAR											:= $(shell tput clear)
export INDENT														:= $(shell tput ht)
export TAB															:= $(shell printf '\011')

export RESET														:= $(shell tput sgr0)
export BOLD   													:= $(shell tput bold)
export RED															:= $(shell tput bold; tput setaf 1)
export GREEN														:= $(shell tput bold; tput setaf 2)
export YELLOW														:= $(shell tput bold; tput setaf 3)

export GIT_REV										:= $(shell (git rev-parse --short HEAD 2>/dev/null || echo "ERROR: failed on 'git rev-parse --short HEAD'"))
export DIR_NAME										:= $(shell basename $(shell pwd))
export PATH_FULL									:= $(shell pwd)
export PATH												:= $(shell echo ${GOPATH}/bin:${PATH})
# project_template
export DIR_NAME										:= $(shell basename $(shell pwd))
# github.com/banaio/golang/project_template
export MODULE_NAME								:= github.com/banaio/golang/${DIR_NAME}
export CMD_NAME										:= ${PATH_FULL}/cmd/main
export CMD_MAIN										:= ./cmd/main.go
export CMD_BUILD_OUTPUT						:= ./project_template.exe
export GOMAXPROCS 								?= 32
export PARALLEL 									?= ${GOMAXPROCS}
export GOLANGCI_LINT_VERSION 			:= v1.30.0
export GOLANGCI_LINT_BINARY				:= $(shell command -v golangci-lint 2> /dev/null)
export VERBOSE										?= -v
# export VERBOSE_FLAGS							?= -v -x
export VERBOSE_FLAGS							?=
export PEDANTIC_FLAGS							?= -race
# export DOCKER_NO_CACHE						?= --no-cache=true
export DOCKER_NO_CACHE						?= --no-cache=false
export GO_TEST_PARALLEL						?= -parallel 32
# When 'go test' runs in package list mode, 'go test' caches successful
# package test results to avoid unnecessary repeated running of tests. To
# disable test caching, use any test flag or argument other than the
# cacheable flags. The idiomatic way to disable test caching explicitly
# is to use -count=1.
export GO_TEST_NO_CACHE						?= -count=1
# Run `$ go tool vet help` then look at `Registered analyzers` section for all the available ones.
export GO_TEST_VET								?= -vet "asmdecl,assign,atomic,bools,buildtag,cgocall,composites,copylocks,errorsas,httpresponse,loopclosure,lostcancel,nilfunc,printf,shift,stdmethods,structtag,tests,unmarshal,unreachable,unsafeptr,unusedresult"

# 1 = print_separator_prefix: prevent warnings from turning on `--warn-undefined-variables` `Makefile` flag.
1 :=
define print_separator
	@print_separator_prefix=$(if $1,$1,$@); \
		prefix="$${print_separator_prefix}:  "; \
		line_terminator=$$'\n'; \
		printf "%b" "${GREEN}" "$${prefix}" `printf -- '-%.0s' $$(seq 1 $$(expr $$(tput cols) - $${#prefix} - $${#line_terminator}))` "${RESET}" "$${line_terminator}"
endef

all: init test run clean lint format
.PHONY: init test run clean lint format

# This allows us to accept extra arguments (by doing nothing when we get a job that doesn't match, rather than throwing an error).
%:
	@:

# Ensure we aren't caching anything before running to ensure everything works.
.PHONY: pre_commit
pre_commit: export VERBOSE=
pre_commit: export PEDANTIC_FLAGS=
pre_commit: export DOCKER_NO_CACHE=--no-cache=true
pre_commit: export GO_TEST_NO_CACHE=-count=1
pre_commit: export GO_TEST_VET=-vet "asmdecl,assign,atomic,bools,buildtag,cgocall,composites,copylocks,errorsas,httpresponse,loopclosure,lostcancel,nilfunc,printf,shift,stdmethods,structtag,tests,unmarshal,unreachable,unsafeptr,unusedresult"
pre_commit:
	$(print_separator)
	# $(MAKE) clean
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test_coloured
	$(MAKE) build
	$(MAKE) run_built_binary
	$(MAKE) image_build
	$(MAKE) image_run

# -count=1: cleans go test cache so tests are always re-run from fresh
.PHONY: test
test:
	$(print_separator)
	GOMAXPROCS=${GOMAXPROCS} go test ${PEDANTIC_FLAGS} ${VERBOSE} ${GO_TEST_NO_CACHE} ${GO_TEST_PARALLEL} ${GO_TEST_VET} ./...

# -count=1: cleans go test cache so tests are always re-run from fresh
.PHONY: test_coloured
test_coloured:
	$(print_separator)
	GOMAXPROCS=${GOMAXPROCS} gotest ${PEDANTIC_FLAGS} ${VERBOSE} ${GO_TEST_NO_CACHE} ${GO_TEST_PARALLEL} ${GO_TEST_VET} ./...

.PHONY: run_built_binary
run_built_binary: build
	$(print_separator)
	${CMD_BUILD_OUTPUT} -vvv

# https://stackoverflow.com/questions/6273608/how-to-pass-argument-to-makefile-from-command-line/6273809
# https://stackoverflow.com/questions/2214575/passing-arguments-to-make-run
#
# The `filter-out` call allows you to do something like:
# tput reset && make run -- -vvv --flag_name "<flag_value>"
.PHONY: run
run:
	$(print_separator)
	go run ${VERBOSE_FLAGS} ${PEDANTIC_FLAGS} "${CMD_MAIN}" $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
build:
	$(print_separator)
	go build -o ${CMD_BUILD_OUTPUT} -trimpath -buildmode=exe -mod readonly -ldflags "-w -s" -installsuffix "static" ${VERBOSE_FLAGS} ${CMD_MAIN}

.PHONY: image_build
image_build:
	$(print_separator)
	docker build ${DOCKER_NO_CACHE} --file=./Dockerfile --tag="${MODULE_NAME}:latest" .

.PHONY: image_run
image_run:
	$(print_separator)
	docker run --interactive --tty --rm "${MODULE_NAME}:latest" $(filter-out $@,$(MAKECMDGOALS)) -vvv

.PHONY: init
init: install_tools
	$(print_separator)
	go mod download

# When cleaning we are only interested in errors hence the pipe to `/dev/null` for stdout.
.PHONY: clean
clean:
	$(print_separator)
	$(call print_separator,"dry run started on go clean")
	@go clean -n -i -r -cache -testcache -modcache ./... 1>/dev/null || (printf "${YELLOW}WARN:${RESET} %s\n" "failed on go clean dry run, exit_code=$$?"; exit 0)
	$(call print_separator,"proper go clean")
	@go clean ${VERBOSE_FLAGS} -i -r -cache -testcache -modcache ./... 1>/dev/null || (printf "${YELLOW}YELLOW:${RESET} %s\n" "failed on go clean, exit_code=$$?"; exit 0)
	@printf "${GREEN}INFO:${RESET} %s\n" "go clean completed"

# Run golint instead???
.PHONY: lint
lint:
	$(print_separator)
	@if [[ ! -x "${GOLANGCI_LINT_BINARY}" ]]; then \
		printf "${YELLOW}WARN:${RESET} %s\n" "linter not installed - golangci-lint=${GOLANGCI_LINT_BINARY}"; \
		printf "${YELLOW}WARN:${RESET} %s\n" "installing golangci-lint@${GOLANGCI_LINT_VERSION} (https://github.com/golangci/golangci-lint) ..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin ${GOLANGCI_LINT_VERSION}; \
		printf "${GREEN}INFO:${RESET} %s - %s\n" \
			"linter installed" \
			"GOPATH/bin=$$(ls -1 --format=commas $$(go env GOPATH)/bin), golangci-lint --version=$$(golangci-lint --version)"; \
	else \
		printf "${GREEN}INFO:${RESET} %s\n" "linter already installed - golangci-lint --version=$$(golangci-lint --version)"; \
	fi
	$(call print_separator,"cache status of golangci-lint") 1>&2
	@golangci-lint ${VERBOSE} cache status 1>&2
	$(call print_separator,"clean cache of golangci-lint") 1>&2
	@golangci-lint ${VERBOSE} cache clean 1>&2
	$(call print_separator,"run fix on golangci-lint")
	golangci-lint run --fix --config .golangci.yml ./...
	$(call print_separator,"cache status of golangci-lint") 1>&2
	golangci-lint ${VERBOSE} cache status 1>&2
	$(call print_separator,"run golangci-lint")
	golangci-lint run --config .golangci.yml ./...
	$(call print_separator,"cache status of golangci-lint") 1>&2
	@golangci-lint ${VERBOSE} cache status 1>&2

# golangci-lint run --fix --config .golangci.yml ./... || printf "\n${YELLOW}WARN:${RESET} %s\n\n" "linting returned an error, exit_code=$$?"
# golangci-lint run --config .golangci.yml ./... || printf "\n${YELLOW}WARN:${RESET} %s\n" "linting returned an error, exit_code=$$?"

.PHONY: install_tools
install_tools: lint
	$(print_separator)
	@if [[ ! -x "$$(command -v gotest 2> /dev/null)" ]]; then \
		printf "${YELLOW}WARN:${RESET} %s\n" "not installed - github.com/rakyll/gotest"; \
		cd ${GOPATH}; \
		go get -u github.com/rakyll/gotest; \
		printf "${GREEN}INFO:${RESET} %s - %s\n" \
			"github.com/rakyll/gotest installed" \
			"GOPATH/bin=$$(ls -1 --format=commas $$(go env GOPATH)/bin), gotest=$$(command -v gotest 2> /dev/null)"; \
	else \
		printf "${GREEN}INFO:${RESET} %s\n" "installed github.com/rakyll/gotest - gotest=$$(command -v gotest 2> /dev/null)"; \
	fi

.PHONY: format
format:
	$(print_separator)
	@printf "${GREEN}INFO:${RESET} %s\n" "cleaning up modules, running go mod tidy ..."
	go mod tidy -v
	@for DIR in $(shell go list -f {{.Dir}} ./...); do \
		printf "%s: ${GREEN}%s${RESET}=%s\n" "formatting" "dir" $$(realpath --relative-to=. "$${DIR}"); \
		goimports -e -w -local ${MODULE_NAME} $${DIR}/; \
	done

.PHONY: debug_env
debug_env: VARS_BUILD:=GIT_REV DIR_NAME PATH_FULL PATH DIR_NAME MODULE_NAME CMD_NAME GOMAXPROCS PARALLEL GOLANGCI_LINT_VERSION GOLANGCI_LINT_BINARY VERBOSE VERBOSE_FLAGS PEDANTIC_FLAGS CMD_MAIN DOCKER_NO_CACHE GO_TEST_NO_CACHE GO_TEST_VET
debug_env:
	$(print_separator)
	$(call print_separator,"VARS_BUILD $@")
	@VARS_BUILD_PRINT=$$( \
		printf '%s\n' ${VARS_BUILD} | \
		xargs -n1 -IV bash -c 'printf "$${INDENT}$${GREEN}%s$${RESET}=%s\n" 'V' "$$(eval "echo $${V}")"' | \
		sort -i \
	); \
	echo "$${VARS_BUILD_PRINT}"
	$(call print_separator,"VARIABLES $@")
	@printf -- '%s\n' $(foreach V, $(filter-out  SCREEN_RESET SCREEN_CLEAR .VARIABLES, $(sort $(.VARIABLES))), \
		$(if $(filter file, $(origin $(V))), \
			'${INDENT}${GREEN}$V${RESET}=$($V) ($(value $V))' \
		) \
	)
