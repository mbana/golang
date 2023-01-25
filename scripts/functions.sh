#!/usr/bin/env bash
set -euf \
  -o nounset \
  -o errexit \
  -o noclobber \
  -o pipefail

# https://stackoverflow.com/questions/9910966/how-to-get-shell-to-self-detect-using-zsh-or-bash
# https://stackoverflow.com/a/9911082/241993
# https://unix.stackexchange.com/questions/9501/how-to-test-what-shell-i-am-using-in-a-terminal
case "${SHELL}" in
  */zsh)
    # assume Zsh
    ;;
  */bash)
    # assume Bash
    shopt -s \
      extglob \
      globstar \
      nullglob
    ;;
  *) ;;
    # assume something else
esac

alias cp='cp -v'
alias rm='rm -v'
alias mkdir='mkdir -v'

SHELL_CURRENT="${SHELL}"
SHELL_VERSION=$("${SHELL}" --version)

INSTALL_COMMAND=""
DEBIAN_FRONTEND=""
PKG_MANAGER_INSTALL_COMMAND=

alias cp='cp --verbose'
alias rm='rm --verbose'
alias mkdir='mkdir --verbose'

SCRIPT_NAME="$(basename "$0")"
UNAME_S=$(uname -s)
PWD="$(pwd)"

INDENT="  "
RESET="$(tput sgr0)"
BOLD="$(tput bold)"
RED="$(
  tput bold
  tput setaf 1
)"
RED_STRING="${RED}${SCRIPT_NAME} - ERROR: ${RESET}"
YELLOW="$(
  tput bold
  tput setaf 3
)"
YELLOW_STRING="${YELLOW}${SCRIPT_NAME} - WARN: ${RESET}"
GREEN="$(
  tput bold
  tput setaf 2
)"
GREEN_STRING="${GREEN}${SCRIPT_NAME} - INFO: ${RESET}"

function print_separator() {
  print_separator_v2 "" ""
}

function print_separator_v2() {
  # https://unix.stackexchange.com/questions/352866/convert-arg-to-uppercase-to-pass-as-variable
  # | tr [a-z] [A-Z]
  # | tr '[:lower:]' '[:upper:]')
  # newvarname=${3^^}
  # local header="${1^^}" # uppercased
  local header="${1}"
  local padding="${2}"
  local separator_chars=""
  separator_chars=$(printf -- '-%.0s' $(seq 1 $(($(tput cols) - ${#header} - ${#padding}))))
  printf "%b" "$(
    tput bold
    tput setaf 2
  )" "${header}" "${padding}" "${separator_chars}" "$(tput sgr0)" "\n"
}

function print_separator_v3() {
  print_separator_v2 "${SCRIPT_NAME}" ":  "
}

if [[ "${UNAME_S}" = "Linux" ]]; then
  INSTALL_COMMAND="sudo apt-get -qq install -y --yes --assume-yes -qq"
  # INSTALL_COMMAND="sudo apt-get install -y --yes --assume-yes -qq"
  # sudo apt-get install -y --yes --assume-yes -qq openssl-devel
  # sudo apt-get -q install -y --yes --assume-yes openssl-devel
  # sudo apt-get -qq install -y --yes --assume-yes openssl-devel
  # sudo apt-get install -y --yes --assume-yes openssl-devel 1>/dev/null
  export DEBIAN_FRONTEND="noninteractive"
elif [[ "${UNAME_S}" = "Darwin" ]]; then
  INSTALL_COMMAND="brew install"
  DEBIAN_FRONTEND=""
else
  printf "%b" \
    "${RED_STRING}" \
    "INSTALL_COMMAND - unsupported OS, UNAME_S=${UNAME_S}" \
    $'\n'
  exit 1
fi
PKG_MANAGER_INSTALL_COMMAND="${INSTALL_COMMAND}"

export SHELL_CURRENT
export SHELL_VERSION

export INSTALL_COMMAND
export DEBIAN_FRONTEND
export PKG_MANAGER_INSTALL_COMMAND

export SCRIPT_NAME
export UNAME_S

export INDENT
export RESET
export BOLD
export RED
export RED_STRING
export YELLOW
export YELLOW_STRING
export GREEN
export GREEN_STRING

export ERROR_STRING="${RED_STRING}"
export WARN_STRING="${YELLOW_STRING}"
export INFO_STRING="${GREEN_STRING}"

function print_env() {
  print_separator_v3
  local BUILD_VARS=("SHELL_CURRENT" "SHELL_VERSION" "UNAME_S" "PWD" "INSTALL_COMMAND" "DEBIAN_FRONTEND" "PKG_MANAGER_INSTALL_COMMAND")
  print_vars=$(printf '%s\n' "${BUILD_VARS[@]}" | xargs -n1 -IV bash -c 'echo -en "${INDENT}${GREEN}V${RESET}=$V\n"')
  echo "${print_vars}"
}

# print_env
