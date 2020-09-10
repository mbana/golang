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

GO_VERSION="go1.13.darwin-amd64.tar.gz"
TMP_DIR="/tmp/go-download/${GO_VERSION}"

mkdir -p "${TMP_DIR}" &&
  cd "${TMP_DIR}" &&
  wget https://golang.org/dl/"${GO_VERSION}" &&
  ls -lah &&
  printf '\n\n' &&
  if [[ -d /usr/local/go ]]; then sudo rm -vr /usr/local/go; fi &&
  printf '\n\n' &&
  sudo tar -C /usr/local -xzf "${GO_VERSION}" &&
  cd /tmp &&
  rm -rv "${TMP_DIR}" &&
  ls -lah &&
  printf '\n\n' &&
  go version
