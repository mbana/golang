# `openbankingforgerock`

## Builds

[![CircleCI](https://circleci.com/gh/banaio/openbankingforgerock.svg?style=svg)](https://circleci.com/gh/banaio/openbankingforgerock)

## Features

Connect's ForgeRock's directory and provides the following features:

1. Command to do OpenID Connect Dynamic Client Registration on ForgeRock's directory.
2. Command to test MATLS setup is correct.
3. Command to consume Accounts (`/open-banking/v3.1/accounts/*`) endpoints:
    * Get an access token to represent you as a TPP using the Client credential flow.
    * Create an Accounts (`/open-banking/v3.1/account-access-consents`) request.
    * Consume the Accounts (`/open-banking/v3.1/accounts`) API using the Hybrid Flow.

See [ForgeRock Opens Up Open Banking](https://www.forgerock.com/about-us/press-releases/forgerock-opens-open-banking).

### Limitations

* For the `token_endpoint_auth_method` we use `private_key_jwt` *only*, no other method is supported.
* Only consumes the `GET` `/open-banking/v3.1/accounts` endpoint for the time being.

## Demo

![./assets/demo/demo.gif](./assets/demo/demo.gif)

## Guide

Jump to [Register](#Register) if you know what you are doing.

### Reading

* [TPP developer user guide](https://backstage.forgerock.com/knowledge/openbanking/book/b77473305)
* [Test Environments Model Banks](https://openbanking.atlassian.net/wiki/spaces/DZ/pages/22512362/Test+Environments+Model+Banks)
* [Integrating a TPP with ForgeRock Model Bank on Directory Sandbox](https://openbanking.atlassian.net/wiki/spaces/DZ/pages/187793608/Integrating+a+TPP+with+ForgeRock+Model+Bank+on+Directory+Sandbox)
* [Integrating a TPP with Ozone Model Banks Using Postman on Directory Sandbox](https://openbanking.atlassian.net/wiki/spaces/DZ/pages/313918598/Integrating+a+TPP+with+Ozone+Model+Banks+Using+Postman+on+Directory+Sandbox)

## TODO

### Configuration

* Document configuration.

### Register

```sh
$ git clone git@github.com:banaio/openbankingforgerock.git
$ cd openbankingforgerock
$ go run cmd/openbankingforgerock/main.go register -r ./config/register-response.json
...
```

### Accounts

```sh
$ git clone git@github.com:banaio/openbankingforgerock.git
$ cd openbankingforgerock
$ go run cmd/openbankingforgerock/main.go accounts -r ./config/register-response.json
...
```

## Upcoming

Rewrite code to re-use these libraries instead of rewriting everything:

* <https://github.com/coreos/go-oidc>
* <https://godoc.org/golang.org/x/oauth2>

Please see this branch <https://github.com/banaio/openbankingforgerock/tree/go-oid_oauth2_refactor>.
