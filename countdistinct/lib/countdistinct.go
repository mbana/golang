package lib

import "fmt"

func Algorithms() map[string]bool {
	return map[string]bool{
		"pcsa": true,
	}
}

type Set interface {
	Add(value []byte) error
	Count() uint32
	fmt.Stringer
}
