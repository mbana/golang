# `TESTS`

## Run Tests

```sh
make init
export PEDANTIC_FLAGS="-race" VERBOSE_FLAGS="-v -x" VERBOSE="-v" DOCKER_NO_CACHE="--no-cache=false"
make test_coloured # make test
...
```

For details on debug flags see Run one of these shell commands then read [`On section in DEBUG.md`](./DEBUG.md#on) section:
