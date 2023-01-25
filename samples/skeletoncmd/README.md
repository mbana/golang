# `github.com/banaio/golang/samples/skeletoncmd`

## Run

```sh
$ go test ./...
ok  	github.com/banaio/golang/samples/skeletoncmd	(cached)
?   	github.com/banaio/golang/samples/skeletoncmd/cmd/skeletoncmd	[no test files]
$ go run cmd/skeletoncmd/main.go
Error: : unsupported choice
Usage:
  skeletoncmd [flags]

Flags:
  -c, --choice string     Available choices: [stat contents]
  -f, --filename string   file (default "/Users/mbana/dev/banaio/github/go_samples/skeletoncmd/skeletoncmd.txt")
  -h, --help              help for skeletoncmd

: unsupported choice
exit status 1
```
