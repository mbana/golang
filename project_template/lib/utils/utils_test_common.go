package utils

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// isParallel - see `init` function further down.
// nolint:gochecknoglobals
var isParallel bool = false

func MakeAssertAndRequire(t *testing.T) (asserter *assert.Assertions, requirer *require.Assertions) {
	if isParallel {
		t.Parallel()
	}
	// Mark as a helper so it doesn't show up in test logs.
	t.Helper()

	// Use assert if you want the tests to continue expecting despite assertions failures.
	asserter = assert.New(t)
	// Use require if you want the tests to stop upon an assertion failure
	requirer = require.New(t)

	return asserter, requirer
}

func MakeTestCaseDependencies(t *testing.T) (asserter *assert.Assertions, requirer *require.Assertions, logger *zerolog.Logger) {
	if isParallel {
		t.Parallel()
	}
	// Mark as a helper so it doesn't show up in test logs.
	t.Helper()

	// Use assert if you want the tests to continue expecting despite assertions failures.
	asserter = assert.New(t)
	// Use require if you want the tests to stop upon an assertion failure
	requirer = require.New(t)
	// Logger that dicards input
	loggerNop := zerolog.Nop()
	logger = &loggerNop

	return asserter, requirer, logger
}

// init - set to true if `-parallel` flag is set in the call to run the test, e.g., `$ go test -parallel=4 ./...`, will set it to `true`.
// Sees like this defaults to `${GOMAXPROCS}` which is a bit odd.
//
// To see why this is needed try the following piece of code:
//
// 	fmt.Fprintln(os.Stderr, ColorGreen, "init ... ", os.Args, ColorReset)
// 	// explicitly init testing module so flags are registered before call to flag.Parse
// 	testing.Init()
// 	fmt.Fprintln(os.Stderr, ColorGreen, "init ... ", os.Args, ColorReset)
// 	flag.Parse()
// 	fmt.Fprintln(os.Stderr, ColorGreen, "init ... ", os.Args, ColorReset)
// nolint:gochecknoinits
func init() {
	colouredString := func(val interface{}) string {
		return fmt.Sprintf("%s%v%s", ColourGreen, val, ColourReset)
	}

	flagSet := flag.CommandLine
	fmt.Fprintf(os.Stderr, "--> lib/utils_test_common.init: args=%v, flagSet=%v", colouredString(os.Args), colouredString(flagSet))

	// HACK: only call `testing.Init()` if these functions return true.
	//
	// 1. Short reports whether the -test.short flag is set.
	//
	// 2. CoverMode reports what the test coverage mode is set to. The
	// values are "set", "count", or "atomic". The return value will be
	// empty if test coverage is not enabled.
	//
	// 3. Verbose reports whether the -test.v flag is set.
	//
	// NB: testing.Short() and testing.Verbose(), so the only way to detect if we are running in test mode is to call testing.CoverMode().
	shouldInit := testing.CoverMode() != "" // || testing.Short() || testing.Verbose()
	if !shouldInit {
		return
	}

	// explicitly init testing module so flags are registered before call to flag.Parse
	testing.Init()
	flag.Parse()

	parallelFlag := flag.Lookup("test.parallel")
	if parallelFlag != nil {
		// This can only ever be a positive int but do the check if it greater than zero just to be sure, e.g.,:
		//
		// 	$ go test -parallel aasda ./...
		// 	invalid value "aasda" for flag -test.parallel: parse error
		//
		// 	$ go test -parallel -1 ./...
		// 	testing: -parallel can only be given a positive integer
		if n, err := (parallelFlag.Value.(flag.Getter).Get()).(int); n > 0 {
			isParallel = true
		} else {
			fmt.Fprintf(os.Stderr, "--> lib/test_utils.go.init: args=%v, error=%s%v%s", os.Args, ColourRed, err, ColourReset)
			isParallel = false
		}
	}
}
