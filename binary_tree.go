package collections

import (
	"fmt"
	"strings"
)

type binaryTree[T ordered] struct {
	*doublyLinkedNode[T]
}

func NewBinaryTree[T ordered]() BinaryTree[T] {
	return &binaryTree[T]{}
}

type BinaryTree[T ordered] interface {
	Collection[T]

	Insert(item T)
	Delete(item T)
	ToDeque() Deque[T]
}

func (node *doublyLinkedNode[T]) Len() int {
	if node == nil {
		return 0
	}
	return 1 + node.prev.Len() + node.next.Len()
}

func (node *doublyLinkedNode[T]) string() string {
	sb := &strings.Builder{}
	if node != nil {
		if node.prev != nil {
			sb.WriteString(node.prev.string())
			sb.WriteRune(' ')
		}
		fmt.Fprint(sb, node.value)
		if node.next != nil {
			sb.WriteRune(' ')
			sb.WriteString(node.next.string())
		}
	}
	return sb.String()
}

func (node *doublyLinkedNode[T]) String() string {
	sb := &strings.Builder{}
	sb.WriteRune('[')
	sb.WriteString(node.string())
	sb.WriteRune(']')
	return sb.String()
}

func (node *doublyLinkedNode[T]) each(i *int, do func(index int, item T)) {
	if node == nil {
		return
	}

	node.prev.each(i, do)

	do(*i, node.value)
	*i++

	node.next.each(i, do)
}

func (node *doublyLinkedNode[T]) Each(do func(index int, item T)) {
	i := 0
	node.each(&i, do)
}

func (node *doublyLinkedNode[T]) eachUntil(i *int, do func(index int, item T), stop func(index int, item T) bool) {
	if node == nil {
		return
	}

	node.prev.eachUntil(i, do, stop)

	do(*i, node.value)

	if stop(*i, node.value) {
		return
	}

	*i++

	node.next.eachUntil(i, do, stop)
}

func (node *doublyLinkedNode[T]) EachUntil(do func(index int, item T), stop func(index int, item T) bool) {
	i := 0
	node.eachUntil(&i, do, stop)
}

func (node *doublyLinkedNode[T]) insert(item T) *doublyLinkedNode[T] {
	if node == nil {
		node = &doublyLinkedNode[T]{value: item}
	} else if item < node.value {
		node.prev = node.prev.insert(item)
	} else {
		node.next = node.next.insert(item)
	}
	return node
}

func (tree *binaryTree[T]) Insert(item T) {
	tree.doublyLinkedNode = tree.doublyLinkedNode.insert(item)
}

func (node *doublyLinkedNode[T]) merge(next *doublyLinkedNode[T]) *doublyLinkedNode[T] {
	if node == nil {
		return next
	} else if next == nil {
		return node
	} else {
		return node.next.merge(next)
	}
}

func (node *doublyLinkedNode[T]) delete(item T) *doublyLinkedNode[T] {
	if node == nil {
		return nil
	} else if item == node.value {
		return node.prev.merge(node.next)
	} else if item < node.value {
		node.prev = node.prev.delete(item)
	} else {
		node.next = node.next.delete(item)
	}
	return node
}

func (tree *binaryTree[T]) Delete(item T) {
	tree.doublyLinkedNode = tree.doublyLinkedNode.delete(item)
}

func (node *doublyLinkedNode[T]) toDeque() *deque[T] {
	if node == nil {
		return &deque[T]{}
	}

	dequeLeft := node.prev.toDeque()
	dequeRight := node.next.toDeque()

	node.prev = dequeLeft.back
	node.next = dequeRight.front

	front, back := node, node

	if node.prev != nil {
		node.prev.next = node
		front = dequeLeft.front
	}
	if node.next != nil {
		node.next.prev = node
		back = dequeRight.back
	}

	return &deque[T]{
		size:  dequeLeft.size + 1 + dequeRight.size,
		front: front,
		back:  back,
	}
}

func (tree *binaryTree[T]) ToDeque() Deque[T] {
	d := tree.toDeque()
	tree.doublyLinkedNode = nil
	return d
}
