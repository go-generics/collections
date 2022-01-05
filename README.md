# go-generics/collections

Go generic collections

`import "github.com/go-generics/collections"`

## Set

Implemented as a map, indexed by T, of string{}

```go
set := collections.NewSet(0, 1, 2)  // set is Set[int]
set.Len()  // 3

set.Has(0)  // true
set.Has(3)  // false

set.Insert(3)
set.Has(3)  // true

set.Delete(3)
set.Has(3)  // false
```

## Deque (Queue / Stack)

Implemented as a doubly linked list.

```go
deque := collections.NewDeque[int]()

deque.PushFront(3)
deque.PushFront(2)
deque.PushFront(1)
deque.PushBack(4)
deque.PushBack(5)
deque.PushBack(6)

fmt.Print(deque)  // [1 2 3 4 5 6]

fmt.Print(deque.Back())     // 6
fmt.Print(deque.PopBack())  // 6
fmt.Print(deque.Back())     // 5

fmt.Print(deque)  // [1 2 3 4 5]

fmt.Print(deque.Front())     // 1
fmt.Print(deque.PopFront())  // 1
fmt.Print(deque.Front())     // 2

fmt.Print(deque)  // [2 3 4 5]
```

## Binary Tree

Implemented with nodes as structs, and references as pointers.

```go
tree := collections.NewBinaryTree[int]()

tree.Insert(1)
tree.Insert(2)
fmt.Println(tree)  // [1 2]

tree.Delete(1)
tree.Delete(2)
fmt.Println(tree)  // []

tree.Delete(-1)
fmt.Println(tree)  // []
```

## Implicit Treap

Implemented with a binary tree, nodes as structs and references as pointers.

Priority is random, stored in each node, and the structure is a binary heap on the priorities.

The key for every value is the position in a sequence.
So the implicit key for every node is the position it has during an in-order traversal.
The structure is a Binary Search Tree on the keys.

https://en.wikipedia.org/wiki/Treap

https://en.wikipedia.org/wiki/Cartesian_tree

https://cp-algorithms.com/data_structures/treap.html

```go
treap := collections.NewImplicitTreap(3, 1, 4, 1, 5, 4, 9, 2)
fmt.Println(treap)  // 3 1 4 1 5 4 9 2
```
