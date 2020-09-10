// Package lib - contains documentation about Sitemap etc. I didn't want to clutter up the
// documentation on the structs themselves so I've put them here instead.
//
// Sitemap - Sitemap 0.9 - https://www.sitemaps.org/protocol.html
//
// Additionally see:
// * https://support.google.com/webmasters/answer/183668#sitemapformat
// * https://support.google.com/webmasters/answer/183668
// * https://support.google.com/webmasters/answer/156184
// * https://www.sitemaps.org/index.html
// * https://en.wikipedia.org/wiki/Sitemaps
// * https://www.sitemaps.org/protocol.html#xmlTagDefinitions
//
// The individual source files contains the code.
package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"golang.org/x/mod/modfile"
)

// Some errors code to indicate exit status.
const (
	// From: /usr/local/go/src/os/proc.go
	// Exit causes the current program to exit with the given status code.
	// Conventionally, code zero indicates success, non-zero an error.
	// The program terminates immediately; deferred functions are not run.
	//
	// For portability, the status code should be in the range [0, 125].
	ExitCodeOK   = 0
	ExitCodeErr1 = 1
)

// Some colours to make logging statements more obvious.
const (
	ColourRed     = "\033[91m"
	ColourGreen   = "\033[92m"
	ColourReset   = "\033[0m"
	ColourBlue    = "\033[34m"
	ColourMagenta = "\033[95m"
)

const (
	GoModFileName = "go.mod"
)

var ErrFailedToFindGoMod = errors.New(`couldn't find "go.mod"`)

// GetAppName - get the module name as defined in `go.mod`.
// Function 'GetAppName' is too long (63 > 60) (funlen)go-lint
// nolint:funlen
func GetAppName(loggerParent *zerolog.Logger) (string, error) {
	logger := loggerParent.With().Str("func", "GetAppName").Str("GoModFileName", GoModFileName).Logger()

	pwd, err := os.Getwd()
	if err != nil {
		// Possibly better to panic in here ...
		errWrapped := fmt.Errorf("failed to os.Getwd in GetAppName, error=%v, parent_error=%w", ErrFailedToFindGoMod, err)
		logger.Error().Err(errWrapped).Msg("failure in GetAppName")

		return "", err
	}

	parentsDirsToTry := 4
	locations := []string{}
	// Generates something like below which goes upto 4 parents directories ...
	// "pwd/"
	// "pwd/../"
	// "pwd/../../"
	// "pwd/../../../"
	for level := 0; level < parentsDirsToTry; level++ {
		pathToParent := strings.Repeat("../", level)
		locations = append(locations, fmt.Sprintf("%s/%s", pwd, pathToParent))
	}
	for pathIndex, path := range locations {
		filename := fmt.Sprintf("%s%s", path, GoModFileName)
		loggerLoop := logger.With().
			Interface("path", path).Interface("locations", locations).Interface("pathIndex", pathIndex).Interface("filename", filename).
			Logger()

		// Possibly check the type of `err`???
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			loggerLoop.Trace().
				Err(err).
				Interface("contents", string(contents)).
				Msg("failed to find file in current path, attempting next path ...")

			continue
		}
		if len(contents) == 0 {
			loggerLoop.Trace().
				Err(err).
				Interface("contents", string(contents)).
				Msg("failed because file in current path has no content, attempting next path ...")

			continue
		}

		moduleName := modfile.ModulePath(contents)
		if moduleName == "" {
			loggerLoop.Trace().
				Interface("contents", string(contents)).
				Interface("moduleName", moduleName).
				Msg("failed because modfile.ModulePath(contents) return an empty string, attempting next path ...")

			continue
		}

		loggerLoop.Trace().
			Interface("contents", string(contents)).
			Interface("moduleName", moduleName).
			Msg("found file - search completed ...")

		return moduleName, nil
	}

	// Possibly better to panic in here ...
	// logger.Panic().Err(err).Msg("failure in GetAppName")
	return "", errors.Wrapf(ErrFailedToFindGoMod, "failed to find file in GetAppName - file=%q, pwd=%q, locations=%v",
		GoModFileName,
		pwd,
		locations,
	)
}
