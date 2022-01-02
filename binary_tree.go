package collections

import (
	"fmt"
	"strings"
)

type binaryTree[T ordered] struct {
	root *doublyLinkedNode[T]
}

func NewBinaryTree[T ordered]() BinaryTree[T] {
	return &binaryTree[T]{}
}

type BinaryTree[T ordered] interface {
	fmt.Stringer

	Collection
	Deleter[T]
	Eacher[T]
	Inserter[T]

	ToDeque() Deque[T]
	Levels()
}

func (node *doublyLinkedNode[T]) btLen() int {
	if node == nil {
		return 0
	}
	return node.size
}

func (tree *binaryTree[T]) Len() int {
	return tree.root.btLen()
}

func (node *doublyLinkedNode[T]) btString() string {
	sb := &strings.Builder{}
	if node != nil {
		if node.prev != nil {
			sb.WriteString(node.prev.btString())
			sb.WriteRune(' ')
		}
		fmt.Fprint(sb, node.value)
		if node.next != nil {
			sb.WriteRune(' ')
			sb.WriteString(node.next.btString())
		}
	}
	return sb.String()
}

func (tree *binaryTree[T]) String() string {
	sb := &strings.Builder{}
	sb.WriteRune('[')
	sb.WriteString(tree.root.btString())
	sb.WriteRune(']')
	return sb.String()
}

func (node *doublyLinkedNode[T]) btEach(i *int, do func(index int, item T)) {
	if node == nil {
		return
	}

	node.prev.btEach(i, do)

	do(*i, node.value)
	*i++

	node.next.btEach(i, do)
}

func (tree *binaryTree[T]) Each(do func(index int, item T)) {
	i := 0
	tree.root.btEach(&i, do)
}

func (node *doublyLinkedNode[T]) btEachUntil(i *int, do func(index int, item T), stop func(index int, item T) bool) {
	if node == nil {
		return
	}

	node.prev.btEachUntil(i, do, stop)

	do(*i, node.value)

	if stop(*i, node.value) {
		return
	}

	*i++

	node.next.btEachUntil(i, do, stop)
}

func (tree *binaryTree[T]) EachUntil(do func(index int, item T), stop func(index int, item T) bool) {
	i := 0
	tree.root.btEachUntil(&i, do, stop)
}

func (node *doublyLinkedNode[T]) btInsert(item T) *doublyLinkedNode[T] {
	if node == nil {
		node = &doublyLinkedNode[T]{value: item}
	} else if item < node.value {
		node.prev = node.prev.btInsert(item)
	} else if item > node.value {
		node.next = node.next.btInsert(item)
	}
	node.size++
	return node
}

func (tree *binaryTree[T]) Insert(item T) {
	tree.root = tree.root.btInsert(item)
}

func (node *doublyLinkedNode[T]) btMerge(next *doublyLinkedNode[T]) *doublyLinkedNode[T] {
	if node == nil {
		return next
	} else if next == nil {
		return node
	} else {
		node.size += next.size
		return node.next.btMerge(next)
	}
}

func (node *doublyLinkedNode[T]) btDelete(item T) *doublyLinkedNode[T] {
	if node == nil {
		return nil
	} else if item == node.value {
		return node.prev.btMerge(node.next)
	} else if item < node.value {
		node.prev = node.prev.btDelete(item)
	} else {
		node.next = node.next.btDelete(item)
	}
	node.size--
	return node
}

func (tree *binaryTree[T]) Delete(item T) {
	tree.root = tree.root.btDelete(item)
}

func (node *doublyLinkedNode[T]) btToDeque() *deque[T] {
	if node == nil {
		return &deque[T]{}
	}

	dequeLeft := node.prev.btToDeque()
	dequeRight := node.next.btToDeque()

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
	d := tree.root.btToDeque()
	tree.root = nil
	return d
}

func (tree *binaryTree[T]) Levels() {
	current_level := []*doublyLinkedNode[T]{tree.root}
	next_level := []*doublyLinkedNode[T]{}

	for len(current_level) > 0 {
		node := current_level[0]
		if node != nil {
			fmt.Printf(" %v", node.value)
			next_level = append(next_level, node.prev, node.next)
		}
		current_level = current_level[1:]
		if len(current_level) == 0 {
			fmt.Println()
			next_level, current_level = current_level, next_level
		}
	}
}
