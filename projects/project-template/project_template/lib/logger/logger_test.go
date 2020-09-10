package logger

import (
	"testing"

	"github.com/rs/zerolog"

	"github.com/banaio/golang/project_template/lib/utils"
)

func Test_NewLogger(t *testing.T) {
	assert, require := utils.MakeAssertAndRequire(t)
	_, _ = assert, require

	// logLevel := len(verbose) <0 == PanicLevel
	// logLevel := len(verbose) 0  == warn
	// logLevel := len(verbose) 1  == info
	// logLevel := len(verbose) 2  == debug
	// logLevel := len(verbose) 3  == trace
	// logLevel := len(verbose) >3 == PanicLevel
	type TestCase struct {
		verbose       []bool
		expectedLevel zerolog.Level
	}
	testCases := []*TestCase{
		{
			verbose:       make([]bool, 0),
			expectedLevel: zerolog.WarnLevel,
		},
		{
			verbose:       make([]bool, 1),
			expectedLevel: zerolog.InfoLevel,
		},
		{
			verbose:       make([]bool, 2),
			expectedLevel: zerolog.DebugLevel,
		},
		{
			verbose:       make([]bool, 3), // make([]bool, zerolog.WarnLevel),
			expectedLevel: zerolog.TraceLevel,
		},
		{
			verbose:       make([]bool, 4),
			expectedLevel: zerolog.DebugLevel,
		},
		{
			verbose:       make([]bool, 8),
			expectedLevel: zerolog.DebugLevel,
		},
		{
			verbose:       make([]bool, 16),
			expectedLevel: zerolog.DebugLevel,
		},
	}
	for index, testCase := range testCases {
		verbose := testCase.verbose
		logWithTimestamps := false
		logger := NewLogger(verbose, logWithTimestamps)
		assert.NotNilf(logger, "testCase=%#v, index=%v", testCase, index)

		expected := testCase.expectedLevel
		actual := logger.GetLevel()

		require.Equalf(expected.String(), actual.String(), "testCase=%#v, index=%v", testCase, index)
	}
}
