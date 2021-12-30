package collections

import "fmt"

type nothing struct{}

type set[T comparable] map[T]nothing

type Set[T comparable] interface {
	Collection[T]

	Insert(items ...T)
	Delete(item T)
	Has(item T) bool
}

func NewSet[T comparable](items ...T) Set[T] {
	s := set[T]{}
	s.Insert(items...)
	return s
}

func (s set[T]) Len() int {
	return len(s)
}

func (s set[T]) String() string {
	items := make([]T, len(s))
	s.Each(func(index int, item T) {
		items[index] = item
	})
	return fmt.Sprint(items)
}

func (s set[T]) Insert(items ...T) {
	for _, item := range items {
		s[item] = nothing{}
	}
}

func (s set[T]) Delete(item T) {
	delete(s, item)
}

func (s set[T]) Has(item T) bool {
	_, exists := s[item]
	return exists
}

func (s set[T]) Each(do func(index int, item T)) {
	i := 0
	for item := range s {
		do(i, item)
		i++
	}
}

func (s set[T]) EachUntil(do func(index int, item T), stop func(index int, item T) bool) {
	i := 0
	for item := range s {
		do(i, item)
		i++
		if stop(i, item) {
			return
		}
	}
}

func Union[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	result := NewSet[T]()

	s1.Each(func(_ int, item T) {
		result.Insert(item)
	})

	s2.Each(func(_ int, item T) {
		result.Insert(item)
	})

	return result
}

func Difference[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	result := NewSet[T]()

	s1.Each(func(_ int, item T) {
		if !s2.Has(item) {
			result.Insert(item)
		}
	})

	return result
}

func Intersection[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	result := NewSet[T]()

	a, b := s1, s2
	if s1.Len() > s2.Len() {
		a, b = b, a
	}

	a.Each(func(_ int, item T) {
		if b.Has(item) {
			result.Insert(item)
		}
	})

	return result
}

func IsSubset[T comparable](s1 Set[T], s2 Set[T]) bool {
	if s1.Len() > s2.Len() {
		return false
	}

	possibleSubset := true
	s1.EachUntil(func(_ int, item T) {
		if !s2.Has(item) {
			possibleSubset = false
		}
	}, func(index int, item T) bool {
		return !possibleSubset
	})

	return possibleSubset
}

func IsProperSubset[T comparable](s1 Set[T], s2 Set[T]) bool {
	return IsSubset(s1, s2) && s1.Len() < s2.Len()
}
