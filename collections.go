package collections

import "fmt"

type Collection[T comparable] interface {
	fmt.Stringer

	Len() int
	Each(action func(index int, item T))
	EachUntil(action func(index int, item T), stop func(index int, item T) bool)
}
