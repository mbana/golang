# `DEBUG`

## On

```sh
export PEDANTIC_FLAGS="-race" VERBOSE_FLAGS="-v -x" VERBOSE="-v" DOCKER_NO_CACHE="--no-cache=false"
make init
# make test_coloured # make test
make run
make image_build
make image_run
```

## Off

```sh
export PEDANTIC_FLAGS="" VERBOSE_FLAGS="" VERBOSE="" DOCKER_NO_CACHE="--no-cache=false"
make init
# make test_coloured # make test
make run
make image_build
make image_run
```
