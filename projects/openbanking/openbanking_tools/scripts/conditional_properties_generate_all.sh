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

function generate_conditional_properties() {
    for SWAGGER_FILE_FULL_RELATIVE_PATH in specifications/read-write/**/*-swagger.yaml; do
        # E.g., `specifications/read-write/v3.1.1/account-info-swagger.yaml` becomes `account-info-swagger.yaml`.
        FILE_NAME=$(basename "${SWAGGER_FILE_FULL_RELATIVE_PATH}")
        # E.g., `specifications/read-write/v3.1.1/account-info-swagger.yaml` becomes `v3.1.1`.
        SPECIFICATION_VERSION=$(basename "$(dirname "${SWAGGER_FILE_FULL_RELATIVE_PATH}")")
        # Remove `-swagger.yaml` suffix from filename, e.g., `account-info-swagger.yaml` becomes `account-info`.
        FILE_NAME_NO_SUFFIX=${FILE_NAME/-swagger.yaml/}
        OUTPUT_FILE="generated/conditional_properties/${SPECIFICATION_VERSION}/${FILE_NAME_NO_SUFFIX}-conditional_properties.json"

        printf "%b" "\033[92m" "Generating conditional_properties\\n\\tSWAGGER_FILE=${SWAGGER_FILE_FULL_RELATIVE_PATH}\\n\\tOUTPUT_FILE=${OUTPUT_FILE}" "\033[0m" "\n"
        # E.g.,
        # go run cmd/openbanking_tools/main.go conditional_properties --swagger_file 'specifications/read-write/v3.1/account-info-swagger.yaml' --output_file 'generated/v3.1/account-info-swagger.json'
        go run cmd/openbanking_tools/main.go conditional_properties --swagger_file="${SWAGGER_FILE_FULL_RELATIVE_PATH}" --output_file="${OUTPUT_FILE}"
    done
}

generate_conditional_properties
