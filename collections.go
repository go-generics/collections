package collections

type Collection interface {
	Len() int
}

type Haser[T comparable] interface {
	Has(item T) bool
}

type Eacher[T any] interface {
	Each(action func(index int, item T))
	EachUntil(action func(index int, item T), stop func(index int, item T) bool)
}

type Inserter[T any] interface {
	Insert(item T)
}

type Deleter[T any] interface {
	Delete(item T)
}

type Backer[T any] interface {
	Back() T
	PopBack() T
	PushBack(item T)
}

type Fronter[T any] interface {
	Front() T
	PopFront() T
	PushFront(item T)
}
