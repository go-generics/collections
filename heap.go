package collections

import (
	"container/heap"
	"fmt"
)

type theap[T ordered] []T

func (h theap[T]) Len() int {
	return len(h)
}

func (h theap[T]) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h theap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *theap[T]) Push(a interface{}) {
	*h = append(*h, a.(T))
}

func (h *theap[T]) Pop() interface{} {
	c := *h
	*h = c[:len(c)-1]
	return c[len(c)-1]
}

type Heap[T ordered] interface {
	fmt.Stringer

	Collection
	Fronter[T]
}

func NewHeap[T ordered](values ...T) Heap[T] {
	h := theap[T](values)
	heap.Init(&h)
	return &h
}

func (h theap[T]) Front() T {
	return h[0]
}

func (h theap[T]) String() string {
	return fmt.Sprint(h)
}

func (h *theap[T]) PopFront() T {
	return heap.Pop(h).(T)
}

func (h *theap[T]) PushFront(item T) {
	heap.Push(h, item)
}
