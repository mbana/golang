# `openbanking/tools`

Open Banking tools written in Go (golang).

## Generate Conditional Endpoints

```sh
$ ./scripts/conditional_properties_generate_all.sh
SWAGGER_FILE=specifications/read-write/v3.1.1/account-info-swagger.yaml
OUTPUT_FILE=generated/v3.1.1/account-info-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.1/account-info-conditional_properties.json swagger_file=specifications/read-write/v3.1.1/account-info-swagger.yaml
INFO Finished generating conditional properties    total=86
INFO Finished writing to file ...                  output_file=generated/v3.1.1/account-info-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.1/callback-urls-swagger.yaml
OUTPUT_FILE=generated/v3.1.1/callback-urls-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.1/callback-urls-conditional_properties.json swagger_file=specifications/read-write/v3.1.1/callback-urls-swagger.yaml
INFO Finished generating conditional properties    total=11
INFO Finished writing to file ...                  output_file=generated/v3.1.1/callback-urls-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.1/client-registration-swagger.yaml
OUTPUT_FILE=generated/v3.1.1/client-registration-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.1/client-registration-conditional_properties.json swagger_file=specifications/read-write/v3.1.1/client-registration-swagger.yaml
INFO Finished generating conditional properties    total=5
INFO Finished writing to file ...                  output_file=generated/v3.1.1/client-registration-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.1/confirmation-funds-swagger.yaml
OUTPUT_FILE=generated/v3.1.1/confirmation-funds-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.1/confirmation-funds-conditional_properties.json swagger_file=specifications/read-write/v3.1.1/confirmation-funds-swagger.yaml
INFO Finished generating conditional properties    total=11
INFO Finished writing to file ...                  output_file=generated/v3.1.1/confirmation-funds-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.1/event-notifications-swagger.yaml
OUTPUT_FILE=generated/v3.1.1/event-notifications-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.1/event-notifications-conditional_properties.json swagger_file=specifications/read-write/v3.1.1/event-notifications-swagger.yaml
INFO Finished generating conditional properties    total=0
INFO Finished writing to file ...                  output_file=generated/v3.1.1/event-notifications-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.1/payment-initiation-swagger.yaml
OUTPUT_FILE=generated/v3.1.1/payment-initiation-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.1/payment-initiation-conditional_properties.json swagger_file=specifications/read-write/v3.1.1/payment-initiation-swagger.yaml
INFO Finished generating conditional properties    total=101
INFO Finished writing to file ...                  output_file=generated/v3.1.1/payment-initiation-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.2/account-info-swagger.yaml
OUTPUT_FILE=generated/v3.1.2/account-info-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.2/account-info-conditional_properties.json swagger_file=specifications/read-write/v3.1.2/account-info-swagger.yaml
INFO Finished generating conditional properties    total=115
INFO Finished writing to file ...                  output_file=generated/v3.1.2/account-info-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.2/aggregated-polling-swagger.yaml
OUTPUT_FILE=generated/v3.1.2/aggregated-polling-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.2/aggregated-polling-conditional_properties.json swagger_file=specifications/read-write/v3.1.2/aggregated-polling-swagger.yaml
INFO Finished generating conditional properties    total=4
INFO Finished writing to file ...                  output_file=generated/v3.1.2/aggregated-polling-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.2/callback-urls-swagger.yaml
OUTPUT_FILE=generated/v3.1.2/callback-urls-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.2/callback-urls-conditional_properties.json swagger_file=specifications/read-write/v3.1.2/callback-urls-swagger.yaml
INFO Finished generating conditional properties    total=11
INFO Finished writing to file ...                  output_file=generated/v3.1.2/callback-urls-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.2/confirmation-funds-swagger.yaml
OUTPUT_FILE=generated/v3.1.2/confirmation-funds-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.2/confirmation-funds-conditional_properties.json swagger_file=specifications/read-write/v3.1.2/confirmation-funds-swagger.yaml
INFO Finished generating conditional properties    total=15
INFO Finished writing to file ...                  output_file=generated/v3.1.2/confirmation-funds-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.2/event-notifications-swagger.yaml
OUTPUT_FILE=generated/v3.1.2/event-notifications-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.2/event-notifications-conditional_properties.json swagger_file=specifications/read-write/v3.1.2/event-notifications-swagger.yaml
INFO Finished generating conditional properties    total=0
INFO Finished writing to file ...                  output_file=generated/v3.1.2/event-notifications-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.2/event-subscriptions-swagger.yaml
OUTPUT_FILE=generated/v3.1.2/event-subscriptions-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.2/event-subscriptions-conditional_properties.json swagger_file=specifications/read-write/v3.1.2/event-subscriptions-swagger.yaml
INFO Finished generating conditional properties    total=15
INFO Finished writing to file ...                  output_file=generated/v3.1.2/event-subscriptions-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.2/payment-initiation-swagger.yaml
OUTPUT_FILE=generated/v3.1.2/payment-initiation-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.2/payment-initiation-conditional_properties.json swagger_file=specifications/read-write/v3.1.2/payment-initiation-swagger.yaml
INFO Finished generating conditional properties    total=163
INFO Finished writing to file ...                  output_file=generated/v3.1.2/payment-initiation-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.3/account-info-swagger.yaml
OUTPUT_FILE=generated/v3.1.3/account-info-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.3/account-info-conditional_properties.json swagger_file=specifications/read-write/v3.1.3/account-info-swagger.yaml
INFO Finished generating conditional properties    total=115
INFO Finished writing to file ...                  output_file=generated/v3.1.3/account-info-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.3/aggregated-polling-swagger.yaml
OUTPUT_FILE=generated/v3.1.3/aggregated-polling-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.3/aggregated-polling-conditional_properties.json swagger_file=specifications/read-write/v3.1.3/aggregated-polling-swagger.yaml
INFO Finished generating conditional properties    total=4
INFO Finished writing to file ...                  output_file=generated/v3.1.3/aggregated-polling-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.3/callback-urls-swagger.yaml
OUTPUT_FILE=generated/v3.1.3/callback-urls-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.3/callback-urls-conditional_properties.json swagger_file=specifications/read-write/v3.1.3/callback-urls-swagger.yaml
INFO Finished generating conditional properties    total=11
INFO Finished writing to file ...                  output_file=generated/v3.1.3/callback-urls-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.3/confirmation-funds-swagger.yaml
OUTPUT_FILE=generated/v3.1.3/confirmation-funds-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.3/confirmation-funds-conditional_properties.json swagger_file=specifications/read-write/v3.1.3/confirmation-funds-swagger.yaml
INFO Finished generating conditional properties    total=15
INFO Finished writing to file ...                  output_file=generated/v3.1.3/confirmation-funds-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.3/event-notifications-swagger.yaml
OUTPUT_FILE=generated/v3.1.3/event-notifications-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.3/event-notifications-conditional_properties.json swagger_file=specifications/read-write/v3.1.3/event-notifications-swagger.yaml
INFO Finished generating conditional properties    total=0
INFO Finished writing to file ...                  output_file=generated/v3.1.3/event-notifications-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.3/event-subscriptions-swagger.yaml
OUTPUT_FILE=generated/v3.1.3/event-subscriptions-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.3/event-subscriptions-conditional_properties.json swagger_file=specifications/read-write/v3.1.3/event-subscriptions-swagger.yaml
INFO Finished generating conditional properties    total=15
INFO Finished writing to file ...                  output_file=generated/v3.1.3/event-subscriptions-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1.3/payment-initiation-swagger.yaml
OUTPUT_FILE=generated/v3.1.3/payment-initiation-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1.3/payment-initiation-conditional_properties.json swagger_file=specifications/read-write/v3.1.3/payment-initiation-swagger.yaml
INFO Finished generating conditional properties    total=163
INFO Finished writing to file ...                  output_file=generated/v3.1.3/payment-initiation-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1/account-info-swagger.yaml
OUTPUT_FILE=generated/v3.1/account-info-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1/account-info-conditional_properties.json swagger_file=specifications/read-write/v3.1/account-info-swagger.yaml
INFO Finished generating conditional properties    total=83
INFO Finished writing to file ...                  output_file=generated/v3.1/account-info-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1/confirmation-funds-swagger.yaml
OUTPUT_FILE=generated/v3.1/confirmation-funds-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1/confirmation-funds-conditional_properties.json swagger_file=specifications/read-write/v3.1/confirmation-funds-swagger.yaml
INFO Finished generating conditional properties    total=11
INFO Finished writing to file ...                  output_file=generated/v3.1/confirmation-funds-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1/event-notifications-swagger.yaml
OUTPUT_FILE=generated/v3.1/event-notifications-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1/event-notifications-conditional_properties.json swagger_file=specifications/read-write/v3.1/event-notifications-swagger.yaml
INFO Finished generating conditional properties    total=0
INFO Finished writing to file ...                  output_file=generated/v3.1/event-notifications-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v3.1/payment-initiation-swagger.yaml
OUTPUT_FILE=generated/v3.1/payment-initiation-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v3.1/payment-initiation-conditional_properties.json swagger_file=specifications/read-write/v3.1/payment-initiation-swagger.yaml
INFO Finished generating conditional properties    total=101
INFO Finished writing to file ...                  output_file=generated/v3.1/payment-initiation-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v4.0/account-info-swagger.yaml
OUTPUT_FILE=generated/v4.0/account-info-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v4.0/account-info-conditional_properties.json swagger_file=specifications/read-write/v4.0/account-info-swagger.yaml
INFO Finished generating conditional properties    total=115
INFO Finished writing to file ...                  output_file=generated/v4.0/account-info-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v4.0/aggregated-polling-swagger.yaml
OUTPUT_FILE=generated/v4.0/aggregated-polling-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v4.0/aggregated-polling-conditional_properties.json swagger_file=specifications/read-write/v4.0/aggregated-polling-swagger.yaml
INFO Finished generating conditional properties    total=4
INFO Finished writing to file ...                  output_file=generated/v4.0/aggregated-polling-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v4.0/client-registration-swagger.yaml
OUTPUT_FILE=generated/v4.0/client-registration-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v4.0/client-registration-conditional_properties.json swagger_file=specifications/read-write/v4.0/client-registration-swagger.yaml
INFO Finished generating conditional properties    total=5
INFO Finished writing to file ...                  output_file=generated/v4.0/client-registration-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v4.0/confirmation-funds-swagger.yaml
OUTPUT_FILE=generated/v4.0/confirmation-funds-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v4.0/confirmation-funds-conditional_properties.json swagger_file=specifications/read-write/v4.0/confirmation-funds-swagger.yaml
INFO Finished generating conditional properties    total=15
INFO Finished writing to file ...                  output_file=generated/v4.0/confirmation-funds-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v4.0/event-notifications-swagger.yaml
OUTPUT_FILE=generated/v4.0/event-notifications-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v4.0/event-notifications-conditional_properties.json swagger_file=specifications/read-write/v4.0/event-notifications-swagger.yaml
INFO Finished generating conditional properties    total=0
INFO Finished writing to file ...                  output_file=generated/v4.0/event-notifications-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v4.0/event-subscriptions-swagger.yaml
OUTPUT_FILE=generated/v4.0/event-subscriptions-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v4.0/event-subscriptions-conditional_properties.json swagger_file=specifications/read-write/v4.0/event-subscriptions-swagger.yaml
INFO Finished generating conditional properties    total=15
INFO Finished writing to file ...                  output_file=generated/v4.0/event-subscriptions-conditional_properties.json
SWAGGER_FILE=specifications/read-write/v4.0/payment-initiation-swagger.yaml
OUTPUT_FILE=generated/v4.0/payment-initiation-conditional_properties.json
INFO flags                                         log_level=INFO
INFO Parsing started ...                           output_file=generated/v4.0/payment-initiation-conditional_properties.json swagger_file=specifications/read-write/v4.0/payment-initiation-swagger.yaml
INFO Finished generating conditional properties    total=163
INFO Finished writing to file ...                  output_file=generated/v4.0/payment-initiation-conditional_properties.json
➜  openbanking_tools git:(master) ✗
```

## Run

**TODO/WIP**.

```sh
$ make build_image
...
Successfully tagged gcr.io/io-bana/openbanking_tools:latest
...
$ docker run --rm -it gcr.io/io-bana/openbanking_tools:latest
TODO

Usage:
  openbanking_tools [command]

Available Commands:
  conditional_properties TODO
  help                   Help about any command

Flags:
  -h, --help               help for openbanking_tools
      --log_level string   Log level (default "INFO")

Use "openbanking_tools [command] --help" for more information about a command.
```
