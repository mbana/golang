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

source ./scripts/functions.sh

# https://stackoverflow.com/questions/23356779/how-can-i-store-the-find-command-results-as-an-array-in-bash/54561526
# https://unix.stackexchange.com/questions/111949/get-list-of-subdirectories-which-contain-a-file-whose-name-contains-a-string
FOUND=$(find "$(pwd)" -type f -name 'go.mod' -printf '%h\n')
TOTAL=$(echo "${FOUND}" | wc -l)
export FOUND
export TOTAL

print_separator_v3
VARS_SCRIPT=("FOUND" "TOTAL")
PRINT_VARS_SCRIPT=$(printf '%s\n' "${VARS_SCRIPT[@]}" | xargs -n1 -IV bash -c 'echo -en "${INDENT}${GREEN}V${RESET}=$V\n"')
echo "${PRINT_VARS_SCRIPT}"

###################################################################################################
#  go env
## --------------------------------------------------------------------------------
## GO111MODULE=""
## GOARCH="amd64"
## GOBIN=""
## GOCACHE="/home/runner/.cache/go-build"
## GOENV="/home/runner/.config/go/env"
## GOEXE=""
## GOFLAGS=""
## GOHOSTARCH="amd64"
## GOHOSTOS="linux"
## GONOPROXY=""
## GONOSUMDB=""
## GOOS="linux"
## GOPATH="/home/runner/go"
## GOPRIVATE=""
## GOPROXY="https://proxy.golang.org,direct"
## GOROOT="/opt/hostedtoolcache/go/1.13.15/x64"
## GOSUMDB="sum.golang.org"
## GOTMPDIR=""
## GOTOOLDIR="/opt/hostedtoolcache/go/1.13.15/x64/pkg/tool/linux_amd64"
## GCCGO="gccgo"
## AR="ar"
## CC="gcc"
## CXX="g++"
## CGO_ENABLED="1"
## GOMOD=""
## CGO_CFLAGS="-g -O2"
## CGO_CPPFLAGS=""
## CGO_CXXFLAGS="-g -O2"
## CGO_FFLAGS="-g -O2"
## CGO_LDFLAGS="-g -O2"
## PKG_CONFIG="pkg-config"
## GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build380702071=/tmp/go-build -gno-record-gcc-switches"
## --------------------------------------------------------------------------------
#print_separator

# GO111MODULE=""
# GOPATH="/home/runner/go"
# GOROOT=/opt/hostedtoolcache/go/1.13.15/x64
GOPATH=$(go env GOPATH)
GOROOT=$(go env GOROOT)
PATH="${GOPATH}/bin:${GOROOT:-/usr/local/go/bin}:${GOROOT:-/usr/local/go}:${GOROOT:-/usr/local/bin}:${PATH}"
export GOPATH
export GOROOT
export PATH

printf "${GREEN}GOPATH=${RESET}%s\n" "${GOPATH}"
printf "${GREEN}GOROOT=${RESET}%s\n" "${GOROOT}"
printf "${GREEN}PATH=${RESET}%s\n" "${PATH}"

mkdir -p "${GOROOT}"
mkdir -p "${GOPATH}"
cd "${GOPATH}"

ls -lah .
print_separator
ls -lah ./bin || true
print_separator

if [[ ! -x "$(command -v gotest 2>/dev/null)" ]]; then
  printf "${YELLOW}WARN:${RESET} %s\n" "not installed - github.com/rakyll/gotest"
  go get -u -x github.com/rakyll/gotest
  printf "${GREEN}INFO:${RESET} %s - %s\n" \
    "github.com/rakyll/gotest installed" \
    "gotest=$(command -v gotest 2>/dev/null)"
else
  printf "${GREEN}INFO:${RESET} %s\n" "found github.com/rakyll/gotest - gotest=$(command -v gotest 2>/dev/null)"
fi

###################################################################################################

print_separator

for MODULE_DIR in ${FOUND[*]}; do
  (
    export GOMAXPROCS=32
    # export VERBOSE_FLAGS='-v -x'
    export VERBOSE='-v'
    export PEDANTIC_FLAGS='-race'
    export GO_TEST_PARALLEL='-parallel 32'
    export GO_TEST_NO_CACHE='-count=1'
    # Run `$ go tool vet help` then look at `Registered analyzers` section for all the available ones.
    export GO_TEST_VET='-vet asmdecl,assign,atomic,bools,buildtag,cgocall,composites,copylocks,errorsas,httpresponse,loopclosure,lostcancel,nilfunc,printf,shift,stdmethods,structtag,tests,unmarshal,unreachable,unsafeptr,unusedresult'

    cd "${MODULE_DIR}"
    echo "  ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR}, ${GREEN}pwd${RESET}=$(pwd)"
    go build ./... || printf "%b" "${RED_STRING}" "Failed on go go build ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR} ... continuing" $'\n'
    gotest ${PEDANTIC_FLAGS} ${VERBOSE} ${GO_TEST_NO_CACHE} ${GO_TEST_PARALLEL} ${GO_TEST_VET} ./... || printf "%b" "${RED_STRING}" "Failed on gotest  ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR} ... continuing" $'\n'
    go run cmd/main.go || make || printf "%b" "${RED_STRING}" "Failed on go run cmd/main.go  ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR} ... continuing" $'\n'
  ) || printf "%b" "${RED_STRING}" "Failed to on ${GREEN}MODULE_DIR${RESET}=${MODULE_DIR} ... continuing" $'\n'
done

print_separator

printf "%b" \
  "${GREEN_STRING}" \
  "built" \
  "${FOUND[@]}" \
  $'\n'

exit 0
