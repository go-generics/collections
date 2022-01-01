package collections

import (
	"fmt"
	"strings"
)

type binaryTree[T ordered] struct {
	*binaryTreeNode[T]
}

type binaryTreeNode[T ordered] struct {
	value T
	left  *binaryTreeNode[T]
	right *binaryTreeNode[T]
}

func NewBinaryTree[T ordered]() BinaryTree[T] {
	return &binaryTree[T]{}
}

type BinaryTree[T ordered] interface {
	Collection[T]

	Insert(item T)
	Delete(item T)
}

func (node *binaryTreeNode[T]) Len() int {
	if node == nil {
		return 0
	}
	return 1 + node.left.Len() + node.right.Len()
}

func (node *binaryTreeNode[T]) string() string {
	sb := &strings.Builder{}
	if node != nil {
		if node.left != nil {
			sb.WriteString(node.left.string())
			sb.WriteRune(' ')
		}
		fmt.Fprint(sb, node.value)
		if node.right != nil {
			sb.WriteRune(' ')
			sb.WriteString(node.right.string())
		}
	}
	return sb.String()
}

func (node *binaryTreeNode[T]) String() string {
	sb := &strings.Builder{}
	sb.WriteRune('[')
	sb.WriteString(node.string())
	sb.WriteRune(']')
	return sb.String()
}

func (node *binaryTreeNode[T]) each(i *int, do func(index int, item T)) {
	if node == nil {
		return
	}

	node.left.each(i, do)

	do(*i, node.value)
	*i++

	node.right.each(i, do)
}

func (node *binaryTreeNode[T]) Each(do func(index int, item T)) {
	i := 0
	node.each(&i, do)
}

func (node *binaryTreeNode[T]) eachUntil(i *int, do func(index int, item T), stop func(index int, item T) bool) {
	if node == nil {
		return
	}

	node.left.eachUntil(i, do, stop)

	do(*i, node.value)

	if stop(*i, node.value) {
		return
	}

	*i++

	node.right.eachUntil(i, do, stop)
}

func (node *binaryTreeNode[T]) EachUntil(do func(index int, item T), stop func(index int, item T) bool) {
	i := 0
	node.eachUntil(&i, do, stop)
}

func (node *binaryTreeNode[T]) insert(item T) *binaryTreeNode[T] {
	if node == nil {
		node = &binaryTreeNode[T]{value: item}
	} else if item < node.value {
		node.left = node.left.insert(item)
	} else {
		node.right = node.right.insert(item)
	}
	return node
}

func (tree *binaryTree[T]) Insert(item T) {
	tree.binaryTreeNode = tree.binaryTreeNode.insert(item)
}

func (node *binaryTreeNode[T]) merge(right *binaryTreeNode[T]) *binaryTreeNode[T] {
	if node == nil {
		return right
	} else if right == nil {
		return node
	} else {
		return node.right.merge(right)
	}
}

func (node *binaryTreeNode[T]) delete(item T) *binaryTreeNode[T] {
	if node == nil {
		return nil
	} else if item == node.value {
		return node.left.merge(node.right)
	} else if item < node.value {
		node.left = node.left.delete(item)
	} else {
		node.right = node.right.delete(item)
	}
	return node
}

func (tree *binaryTree[T]) Delete(item T) {
	tree.binaryTreeNode = tree.binaryTreeNode.delete(item)
}
