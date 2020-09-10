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

export VERBOSE													:=
export DIR_FULL_PATH										:= $(shell pwd)
export DIR_NAME													:= $(shell basename $(shell pwd))

export PATH															:= $(shell echo ${GOPATH}/bin:${PATH})
export DIR_NAME													:= $(shell basename $(shell pwd))
export MODULE_NAME											:= github.com/banaio/golang

# 1 = print_separator_prefix: prevent warnings from turning on `--warn-undefined-variables` `Makefile` flag.
1 :=
define print_separator
	@print_separator_prefix=$(if $1,$1,$@); \
		prefix="$${print_separator_prefix}:  "; \
		line_terminator=$$'\n'; \
		printf "%b" "${GREEN}" "$${prefix}" `printf -- '-%.0s' $$(seq 1 $$(expr $$(tput cols) - $${#prefix} - $${#line_terminator}))` "${RESET}" "$${line_terminator}"
endef

.PHONY: all
all: pre_commit

# This allows us to accept extra arguments (by doing nothing when we get a job that doesn't match, rather than throwing an error).
%:
	@:

.PHONY: run
run:
	$(print_separator)
	$(call print_separator,"started $@")
	@echo "running:" $(filter-out $@,$(MAKECMDGOALS))
	$(call print_separator,"completed $@")

.PHONY: pre_commit
pre_commit:
	$(print_separator)
	$(call print_separator,"started $@")
	$(MAKE) debug_env
	$(MAKE) run
	$(call print_separator,"completed $@")

.PHONY: debug_env
debug_env: VARS_BUILD:=DIR_NAME DIR_FULL_PATH VERBOSE
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

.PHONY: install_tools
install_tools: lint
	$(print_separator)
	cd ${GOPATH}; \
		go get mvdan.cc/gofumpt; \
		go get mvdan.cc/gofumpt/gofumports; \
		go get -u github.com/rakyll/gotest; \
		printf -- '%s\n' \
			"gofumpt=$$(command -v gofumpt 2> /dev/null)" \
			"gofumports=$$(command -v gofumports 2> /dev/null)" \
			"gotest=$$(command -v gotest 2> /dev/null)"

# https://github.com/mvdan/gofumpt

# .PHONY: install_tools
# install_tools: lint
# 	$(print_separator)
# 	@cd ${GOPATH}; \
# 		if [[ ! -x "$$(command -v gotest 2> /dev/null)" ]]; then \
# 			printf "${YELLOW}WARN:${RESET} %s\n" "not installed - github.com/rakyll/gotest"; \
# 			go get -u github.com/rakyll/gotest; \
# 			printf "${GREEN}INFO:${RESET} %s - %s\n" \
# 				"github.com/rakyll/gotest installed" \
# 				"GOPATH/bin=$$(ls -1 --format=commas $$(go env GOPATH)/bin), gotest=$$(command -v gotest 2> /dev/null)"; \
# 		else \
# 			printf "${GREEN}INFO:${RESET} %s\n" "installed github.com/rakyll/gotest - gotest=$$(command -v gotest 2> /dev/null)"; \
# 		fi \
# 		GO111MODULE=on go get mvdan.cc/gofumpt; \
# 		GO111MODULE=on go get mvdan.cc/gofumpt/gofumports

.PHONY: git_compress
git_compress:
	$(print_separator)
	$(call print_separator,"before $@")
	git count-objects -Hv
	git gc --aggressive --prune=now
	git gc --prune=now
	git repack -Ad
	git prune
	git reflog expire --all --expire=now
	git gc --aggressive --prune=now
	git gc --aggressive
	git prune
	git gc --aggressive --prune=now
	git gc --prune=now
	git repack -Ad
	git prune
	git reflog
	$(call print_separator,"after $@")
	git count-objects -Hv

# have_term := $(shell echo $$TERM)
# ifdef have_term
# define my_color =
#     @tput setaf $2
#     @tput bold
#     @echo $1
#     @tput sgr0
# endef
# else
# my_color = @echo $1
# endif
