// Package cli - cli options. This uses the CLI parsing library: https://pkg.go.dev/github.com/jessevdk/go-flags@v1.4.0.
package cli

// Options - contains the flag options.
// See:
// * https://pkg.go.dev/github.com/jessevdk/go-flags@v1.4.0.
// * https://pkg.go.dev/github.com/jessevdk/go-flags@v1.4.0?tab=doc#hdr-Basic_usage.
//
// nolint:lll
type Options struct {
	// Verbose: This specifies one option with a short name -v and a long name --verbose.
	// When specifying -v is NOT specified: WARN level log messages are logged only.
	// When specifying -v: INFO level log messages are logged only.
	// When specifying -vv: DEBUG and INFO level log messages are logged.
	// When specifying -vvv: TRACE, DEBUG and INFO level log messages are logged
	Verbose []bool `required:"false" long:"verbose" short:"v" description:"This specifies one option with a short name -v and a long name --verbose. When specifying -v: WARN level log messages are logged only. When specifying -vv: INFO level log messages are logged. When specifying -vvv: TRACE level log messages are logged."`
}
