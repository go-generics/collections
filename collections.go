package collections

import "fmt"

type Collection[T comparable] interface {
	fmt.Stringer

	Len() int
	Iterate(action func(index int, item T))
	IterateUntil(action func(index int, item T), stop func(index int, item T) bool)
}
