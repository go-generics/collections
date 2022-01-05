package collections

import (
	"fmt"
	"math/rand"
)

type ImplicitTreap[T comparable] struct {
	Root *ImplicitTreapNode[T]
}

type ImplicitTreapNode[T comparable] struct {
	Left     *ImplicitTreapNode[T]
	Right    *ImplicitTreapNode[T]
	Value    T
	Size     int
	Priority float64
}

func NewImplicitTreapNode[T comparable](value T) *ImplicitTreapNode[T] {
	return &ImplicitTreapNode[T]{
		Value:    value,
		Size:     1,
		Priority: rand.Float64(),
	}
}

func NewImplicitTreap[T comparable](values ...T) *ImplicitTreap[T] {
	return &ImplicitTreap[T]{NewImplicitTreapRoot(values...)}
}

func NewImplicitTreapRoot[T comparable](values ...T) *ImplicitTreapNode[T] {
	if len(values) == 0 {
		return nil
	}

	mid := len(values) / 2
	newNode := NewImplicitTreapNode(values[mid])
	newNode.Left = NewImplicitTreapRoot(values[:mid]...)
	newNode.Right = NewImplicitTreapRoot(values[mid+1:]...)

	newNode.updatePriorities()
	newNode.updateSize()

	return newNode
}

func (n *ImplicitTreap[T]) String() string {
	return n.Root.String()
}

func (n *ImplicitTreapNode[T]) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v%v %v", n.Left, n.Value, n.Right)
}

func (n *ImplicitTreapNode[T]) effectiveSize() int {
	if n == nil {
		return 0
	}
	return n.Size
}

func (n *ImplicitTreapNode[T]) updateSize() {
	if n != nil {
		n.Size = n.Left.effectiveSize() + 1 + n.Right.effectiveSize()
	}
}

func (n *ImplicitTreapNode[T]) updatePriorities() {
	if n == nil {
		return
	}
	max := n
	if n.Left != nil && n.Left.Priority > max.Priority {
		max = n.Left
	}
	if n.Right != nil && n.Right.Priority > max.Priority {
		max = n.Right
	}
	if max != n {
		max.Priority, n.Priority = n.Priority, max.Priority
		max.updatePriorities()
	}
}

func (it *ImplicitTreap[T]) Append(item T) {
	right := NewImplicitTreapNode(item)
	it.Root = it.Root.Merge(right)
}

func (l *ImplicitTreap[T]) Merge(r *ImplicitTreap[T]) {
	l.Root = l.Root.Merge(r.Root)
}

func (l *ImplicitTreapNode[T]) Merge(r *ImplicitTreapNode[T]) (p *ImplicitTreapNode[T]) {
	if l == nil {
		p = r

	} else if r == nil {
		p = l

	} else if l.Priority > r.Priority {
		l.Right = l.Right.Merge(r)
		p = l

	} else {
		r.Left = l.Merge(r.Left)
		p = r
	}

	p.updateSize()
	return p
}

func (it *ImplicitTreap[T]) Split(key int) (l *ImplicitTreap[T], r *ImplicitTreap[T]) {
	ln, rn := it.Root.Split(key, 0)
	l = &ImplicitTreap[T]{ln}
	r = &ImplicitTreap[T]{rn}
	it.Root = nil
	return
}

func (p *ImplicitTreapNode[T]) Split(key, offset int) (l *ImplicitTreapNode[T], r *ImplicitTreapNode[T]) {
	if p == nil {
		return
	}
	position := offset + p.Left.effectiveSize()
	if key <= position {
		l, p.Left = p.Left.Split(key, offset)
		r = p
	} else {
		p.Right, r = p.Right.Split(key, position+1)
		l = p
	}
	p.updateSize()
	return
}

func (it *ImplicitTreap[T]) At(key int) T {
	return it.Root.At(key)
}

func (n *ImplicitTreapNode[T]) At(key int) T {
	leftSize := n.Left.effectiveSize()

	if key == leftSize {
		return n.Value
	} else if key < leftSize {
		return n.Left.At(key)
	} else {
		return n.Right.At(key - leftSize - 1)
	}
}
