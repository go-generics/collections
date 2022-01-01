package collections

import "fmt"

type deque[T any] struct {
	size  int
	front *dequeNode[T]
	back  *dequeNode[T]
}

type dequeNode[T any] struct {
	value T
	next  *dequeNode[T]
	prev  *dequeNode[T]
}

func NewDeque[T any]() Deque[T] {
	return &deque[T]{}
}

type Deque[T any] interface {
	Collection[T]

	PushBack(item T)
	PushFront(item T)
	PopBack() T
	PopFront() T
	Back() T
	Front() T
}

func (d *deque[T]) Len() int {
	return d.size
}

func (d *deque[T]) String() string {
	items := make([]T, d.size)
	d.Each(func(index int, item T) {
		items[index] = item
	})
	return fmt.Sprint(items)
}

func (d *deque[T]) Each(do func(index int, item T)) {
	for node, i := d.front, 0; node != nil; node, i = node.next, i+1 {
		do(i, node.value)
	}
}

func (d *deque[T]) EachUntil(do func(index int, item T), stop func(index int, item T) bool) {
	for node, i := d.front, 0; node != nil; node, i = node.next, i+1 {
		do(i, node.value)
		if stop(i, node.value) {
			return
		}
	}
}

func (d *deque[T]) PushBack(item T) {
	newNode := dequeNode[T]{
		value: item,
		prev:  d.back,
	}
	if d.back != nil {
		d.back.next = &newNode
	} else {
		d.front = &newNode
	}
	d.back = &newNode
	d.size++
}

func (d *deque[T]) PushFront(item T) {
	newNode := dequeNode[T]{
		value: item,
		next:  d.front,
	}
	if d.front != nil {
		d.front.prev = &newNode
	} else {
		d.back = &newNode
	}
	d.front = &newNode
	d.size++
}

func (d *deque[T]) PopBack() T {
	deletedNode := d.back
	d.back = d.back.prev
	if d.back != nil {
		d.back.next = nil
	} else {
		d.front = nil
	}
	d.size--
	return deletedNode.value
}

func (d *deque[T]) PopFront() T {
	deletedNode := d.front
	d.front = d.front.next
	if d.front != nil {
		d.front.prev = nil
	} else {
		d.back = nil
	}
	d.size--
	return deletedNode.value
}

func (d *deque[T]) Back() T {
	return d.back.value
}

func (d *deque[T]) Front() T {
	return d.front.value
}
