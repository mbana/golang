// +build tools

package tools

// This file exists to cause `go mod` and `go get` to believe these packages
// are dependencies, even though they are not runtime dependencies.  This means
// they will appear in our `go.mod` file, but will not be a part of the build
// unless the "tools" build tag is specified.
//
// References:
// * https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// * https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
// * https://github.com/golang/go/issues/25922#issuecomment-412992431
//
// Examples:
// * https://github.com/manifoldco/healthz/blob/master/tools.go
// * https://github.com/manifoldco/grafton/blob/master/tools.go
// * https://github.com/kubernetes-csi/drivers/blob/master/vendor/google.golang.org/grpc/test/tools/tools.go
// * https://gitlab.cncf.ci/grpc/grpc-go/commit/ce4f3c8a89229d9db3e0c30d28a9f905435ad365#3387d70491f41f4047ef4566403ce47b46147f30


import (
	_ "github.com/fzipp/gocyclo"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/vektra/mockery/cmd/mockery"
	_ "golang.org/x/lint/golint"
)
