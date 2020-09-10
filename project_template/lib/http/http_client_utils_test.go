package lib

import "strings"

const (
	TestURL = "https://bana.io"
)

// NewBodyEmpty - return new reader. Might remove later as I think we should just pass `nil` as the `body` argument.
func NewBodyEmpty() *strings.Reader { return strings.NewReader("") }

// NewBodyNil - return nil `body` argument for request. Might remove later.
// func NewBodyNil() interface{} { return nil }
// I think this is the correct implementation.
