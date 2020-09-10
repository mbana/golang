package lib

import "fmt"

type ValueConstraint interface {
	[]byte | string
}

type Set[T ValueConstraint] interface {
	fmt.Stringer

	Add(value T) error
	Count() uint32
}

func Algorithms() map[string]bool {
	return map[string]bool{
		"pcsa": true,
		"PCSA": true,
		"cms":  true,
		"CMS":  true,
	}
}
