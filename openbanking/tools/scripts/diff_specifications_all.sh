#!/usr/bin/env bash
# Tested against Bash 4.4.23(1)-release

set -eu -o pipefail

set -o noclobber    # Avoid overlay files (echo "hi" > foo)
set -o errexit      # Used to exit upon error, avoiding cascading errors
set -o pipefail     # Unveils hidden failures
set -o nounset      # Exposes unset variables

# set -o nullglob     # Non-matching globs are removed  ('*.foo' => '')
shopt -s nullglob   # Non-matching globs are removed  ('*.foo' => '')
# set -o failglob     # Non-matching globs throw errors
shopt -s failglob   # Non-matching globs throw errors
# set -o nocaseglob   # Case insensitive globs
shopt -s nocaseglob # Case insensitive globs
# set -o globstar     # Allow ** for recursive matches ('lib/**/*.rb' => 'lib/a/b/c.rb')
shopt -s globstar # Allow ** for recursive matches ('lib/**/*.rb' => 'lib/a/b/c.rb')

function check_swagger_command() {
    # Possibly update `swagger` if `swagger diff` commands fails as this subcommand is only present in newer releases?
    # $ go get -u github.com/go-swagger/go-swagger/...

    # if ! type "swagger" > /dev/null; then
    if ! [[ -x "$(command -v swagger)" ]]; then
        printf "%b" "\033[93m" "swagger (https://github.com/go-swagger/go-swagger): installing ..." "\033[0m" "\n"

        download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest |
            jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
        curl -o /usr/local/bin/swagger -L'#' "$download_url"
        chmod +x /usr/local/bin/swagger

        # # https://goswagger.io/install.html#installing-from-source
        # TEMP_DIR=$(mktemp -d)
        # git clone git@github.com:go-swagger/go-swagger.git "${TEMP_DIR}"
        # cd "${TEMP_DIR}"
        # go install ./cmd/swagger
    fi

    printf "%b" "\033[92m" "swagger (https://github.com/go-swagger/go-swagger): path=$(command -v swagger), version=$(swagger version)" "\033[0m" "\n"
}

function diff_swagger_files() {
    # Pair-wise combinations of file to run `swagger diff` against
    for i in specifications/read-write/**/account-info-swagger.yaml; do
        # E.g., specifications/read-write/v3.1/account-info-swagger.yaml
        SWAGGER_FILE="${i}"
        # E.g., account-info-swagger.yaml
        FILE_NAME=$(basename "${SWAGGER_FILE}")

        # only diff against other `account-info-swagger.yaml` specifications
        for j in specifications/read-write/**/"${FILE_NAME}"; do
            if [ "${i}" \< "${j}" ]; then
                # E.g., `specifications/read-write/v3.1.1/account-info-swagger.yaml` becomes `v3.1.1`.
                SPEC1_VERSION=$(basename "$(dirname "${i}")")
                # E.g., SPEC1=specifications/read-write/v3.1.1/account-info-swagger.yaml
                SPEC1="${i}"
                # E.g., `specifications/read-write/v3.1.1/account-info-swagger.yaml` becomes `account-info-swagger.yaml`.
                SPEC1_FILE_NAME=$(basename "${SPEC1}")

                SPEC2_VERSION=$(basename "$(dirname "${j}")")
                SPEC2="${j}"
                SPEC2_FILE_NAME=$(basename "${SPEC2}")

                # E.g., DESTINATION=generated/diff/v3.1.1_account-info_vs_v3.1.2_account-info.{log,json}
                DESTINATION="$(pwd)/generated/diff/${SPEC1_VERSION}_${SPEC1_FILE_NAME/-swagger.yaml/}_vs_${SPEC2_VERSION}_${SPEC2_FILE_NAME/-swagger.yaml/}"

                printf "%b" "\033[92m" "swagger diff\\n\\tSPEC1=${SPEC1}\\n\\tSPEC2=${SPEC2}\\n\\tDESTINATION=${DESTINATION}.{log,json}" "\033[0m" "\n"
                swagger diff --log-output "${DESTINATION}.log" --format='json' "${SPEC1}" "${SPEC2}" >"${DESTINATION}.json" || true
            fi
        done
    done
}

check_swagger_command
diff_swagger_files
