package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/banaio/golang/project_template/lib/utils"
)

const (
	// TimeFormatDefault - date/time format used in logging out.
	TimeFormatDefault = time.RFC3339
	// TimestampEmpty - don't print timestamps.
	TimestampEmpty = ""
)

// Parameter to use to get a logger configured with trace level logging.
func NewLoggerParameterValueForTraceLevelLogger() []bool {
	return make([]bool, zerolog.WarnLevel)
}

func NewLoggerWarnAndAboveLevels(logWithTimestamps bool) *zerolog.Logger {
	return NewLogger(make([]bool, zerolog.WarnLevel), logWithTimestamps)
}

// NewLogger - configure logging level. If `verbose` array is empty it will only issue logs of a WARN level.
// verbose := []bool{}
// Logger(verbose).GetLevel() == zerolog.WarnLevel
//
// `logWithTimestamps`: Determines if date/time should be on the log line.
// See, https://pkg.go.dev/github.com/jessevdk/go-flags@v1.4.0?tab=doc#hdr-Basic_usage.
func NewLogger(verbose []bool, logWithTimestamps bool) *zerolog.Logger {
	var level zerolog.Level

	verboseFlagsCount := len(verbose)
	// mnd: Magic number: 2, in <condition> detected (gomnd)
	// nolint:mnd
	if verboseFlagsCount == 0 {
		level = zerolog.WarnLevel
	} else if verboseFlagsCount == 1 {
		level = zerolog.InfoLevel
	} else if verboseFlagsCount == 2 {
		level = zerolog.DebugLevel
	} else if verboseFlagsCount == 3 {
		level = zerolog.TraceLevel
	} else {
		level = zerolog.DebugLevel
	}

	timeFormat := TimeFormatDefault
	if !logWithTimestamps {
		timeFormat = TimestampEmpty
	}

	writer := &zerolog.ConsoleWriter{
		NoColor:    false,
		Out:        os.Stderr,
		TimeFormat: timeFormat,
	}
	writer.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s%s%s", utils.ColourBlue, i, utils.ColourReset)
	}
	writer.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("\n\t%s%s%s=", utils.ColourMagenta, i, utils.ColourReset)
	}
	writer.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("'%s'", i)
	}
	if !logWithTimestamps {
		writer.FormatTimestamp = func(i interface{}) string {
			return TimestampEmpty
		}
		logger := makeLogger(writer, level)

		return &logger
	}

	logger := makeLogger(writer, level).With().Timestamp().Logger()

	return &logger
}

func makeLogger(writer io.Writer, level zerolog.Level) zerolog.Logger {
	return log.Output(writer).
		Level(level).
		With().
		Stack(). // Stack enables stack trace printing for the error passed to Err().
		Logger().
		With().
		Caller(). // Caller adds the file:line of the caller with the zerolog.CallerFieldName key.
		Logger()
}
