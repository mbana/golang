package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/pkg/errors"

	"github.com/banaio/golang/project_template/cmd/cli"
	"github.com/banaio/golang/project_template/lib"
	"github.com/banaio/golang/project_template/lib/logger"
)

// Some errors that could be thrown.
var (
	ErrFlagsNewParserNil   = errors.New("flags.NewParser parser == nil")
	ErrFailedToGetAppName  = errors.New("failed on GetAppName")
	ErrFailedToParserFlags = errors.New("failed on flags.NewParser(&options, flags.Default)")
)

func main() {
	// Use a logger set with the trace level - the log level passed on the command line will be used.
	logWithTimestamps := false
	appName, err := lib.GetAppName(logger.NewLogger(logger.NewLoggerParameterValueForTraceLevelLogger(), logWithTimestamps))
	if err != nil {
		exitf(func() {}, lib.ExitCodeErr1, "%v", fmt.Errorf("error=%v - parent_error=%w", ErrFailedToGetAppName, err))
	}

	cliOpts := cli.Options{}
	parser := flags.NewParser(&cliOpts, (flags.Default &^ flags.PrintErrors)) // Don't print errors as we will do this ourselves
	if parser == nil {
		exitf(
			func() {},
			lib.ExitCodeErr1,
			"%v",
			fmt.Errorf("error=%v - cliOpts=%+v, parent_error=%w", ErrFailedToParserFlags, cliOpts, ErrFlagsNewParserNil),
		)
	} else {
		parser.Command.Name = appName
	}

	ParseCLIOrExit(parser)

	// Possibly use appName by doing:
	// 	lib.NewLogger(cliOpts.Verbose).With().Str("appName", appName).Logger()
	logger := logger.NewLogger(cliOpts.Verbose, logWithTimestamps)
	logger.Info().
		Interface("cliOpts", cliOpts).
		Str("logger.GetLevel()", logger.GetLevel().String()).Str("func", "main").Interface("cliOpts", cliOpts).
		Msg("configured logger settings")

	logger.Info().Str("func", "main").Msg("project_template finished")
}

// ParseCLIOrExit - parses CLI options or terminates depending on options and flags
// passed. Ideally, we should just pass the error back to `main`, but `main` is getting too large.
func ParseCLIOrExit(parser *flags.Parser) {
	_, err := parser.Parse()
	if err != nil {
		// Check errort type so don't print errors/warnings in red.
		// `--help`: should be in plain colour.
		flagsErrorType := err.(*flags.Error).Type
		if flagsErrorType == flags.ErrHelp || flagsErrorType == flags.ErrCommandRequired {
			parser.WriteHelp(os.Stderr)
			os.Exit(lib.ExitCodeOK)
		}

		// Print in red if it's not a call to `--help`, then the usage menu.
		exitf(func() {
			parser.WriteHelp(os.Stderr)
		}, lib.ExitCodeErr1, "\n%v\n", err)
	}
}

// exitf - `beforeExitFunc` is called before printing the message to stderr.
func exitf(beforeExitFunc func(), code int, format string, args ...interface{}) {
	if beforeExitFunc != nil {
		beforeExitFunc()
	}

	formatStr := fmt.Sprintf("%s%s%s", lib.ColourRed, format, lib.ColourReset)
	fmt.Fprintf(os.Stderr, formatStr, args...)
	os.Exit(code)
}
