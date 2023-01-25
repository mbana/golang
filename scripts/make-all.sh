#!/usr/bin/env bash
set -euf \
  -o nounset \
  -o errexit \
  -o noclobber \
  -o pipefail

set -x

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

source ./scripts/functions.sh

# https://stackoverflow.com/questions/23356779/how-can-i-store-the-find-command-results-as-an-array-in-bash/54561526
# https://unix.stackexchange.com/questions/111949/get-list-of-subdirectories-which-contain-a-file-whose-name-contains-a-string
FOUND=$(find "$(pwd)" -type f -name 'go.mod' -printf '%h\n' | grep -v 'samples/')
TOTAL=$(echo "${FOUND}" | wc -l)
export FOUND
export TOTAL

print_separator_v3
VARS_SCRIPT=("FOUND" "TOTAL")
PRINT_VARS_SCRIPT=$(printf '%s\n' "${VARS_SCRIPT[@]}" | xargs -n1 -IV bash -c 'echo -en "${INDENT}${GREEN}V${RESET}=$V\n"')
echo "${PRINT_VARS_SCRIPT}"

###################################################################################################

print_separator

for MODULE_DIR in ${FOUND[*]}; do
  (
    # Run `$ go tool vet help` then look at `Registered analyzers` section for all the available ones.
    export GO_TEST_VET='-vet asmdecl,assign,atomic,bools,buildtag,cgocall,composites,copylocks,errorsas,httpresponse,loopclosure,lostcancel,nilfunc,printf,shift,stdmethods,structtag,tests,unmarshal,unreachable,unsafeptr,unusedresult'

    cd "${MODULE_DIR}"
    echo "  ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR}, ${GREEN}pwd${RESET}=$(pwd)"
    go build ./... || printf "%b" "${RED_STRING}" "Failed on go go build ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR} ... continuing" $'\n'
    go test -race -v -count=1 ${GO_TEST_VET} ./... || printf "%b" "${RED_STRING}" "Failed on go test  ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR} ... continuing" $'\n'
  ) || printf "%b" "${RED_STRING}" "Failed to on ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR} ... continuing" $'\n'
done

print_separator

printf "%b" \
  "${GREEN_STRING}" \
  "built" \
  "${FOUND[@]}" \
  $'\n'

exit 0
