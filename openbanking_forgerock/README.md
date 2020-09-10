# `openbanking_forgerock`

## Builds

[![CircleCI](https://circleci.com/gh/banaio/openbanking_forgerock.svg?style=svg)](https://circleci.com/gh/banaio/openbanking_forgerock)

## Upcoming

I am rewriting the code to re-use these libraries instead of rewriting everything:

* <https://github.com/coreos/go-oidc>
* <https://godoc.org/golang.org/x/oauth2>

Please see this branch <https://github.com/banaio/golang/openbanking_forgerock/tree/go-oid_oauth2_refactor>.

## Features

Connect to ForgeRock's directory:

1. Registers on ForgeRock's directory.
2. Tests the MATLS setup.
3. Get an access token to represent you as a TPP using the Client credential flow.
4. Create an account request.
5. Consume the accounts API.

See [ForgeRock Opens Up Open Banking](https://www.forgerock.com/about-us/press-releases/forgerock-opens-open-banking).

## Demo

![](./demos/demo.gif)

## TODO

### Configuration

* Document configuration.

### Run

```sh
$ go get -u github.com/banaio/golang/openbanking_forgerock/...
$ cd $GOPATH/src/github.com/banaio/golang/openbanking_forgerock/
$ make run
...
INFO MTLSTest                                      StatusCode=200 result="{\"issuerId\":\"...\",\"authorities\":[{\"authority\":\"AISP\"},{\"authority\":\"PISP\"}]}"
INFO GetAccessToken                                StatusCode=200 accessToken="{Access_token:... Scope:openid payments accounts Id_token:... Token_type:Bearer Expires_in:86399}"
```

## Reading

* [TPP developer user guide](https://backstage.forgerock.com/knowledge/openbanking/book/b77473305)
* [Test Environments Model Banks](https://openbanking.atlassian.net/wiki/spaces/DZ/pages/22512362/Test+Environments+Model+Banks)
* [Integrating a TPP with ForgeRock Model Bank on Directory Sandbox](https://openbanking.atlassian.net/wiki/spaces/DZ/pages/187793608/Integrating+a+TPP+with+ForgeRock+Model+Bank+on+Directory+Sandbox)
* [Integrating a TPP with Ozone Model Banks Using Postman on Directory Sandbox](https://openbanking.atlassian.net/wiki/spaces/DZ/pages/313918598/Integrating+a+TPP+with+Ozone+Model+Banks+Using+Postman+on+Directory+Sandbox)
