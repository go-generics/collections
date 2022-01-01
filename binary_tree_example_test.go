package collections_test

import (
	"fmt"

	"github.com/go-generics/collections"
)

func ExampleBinaryTree_Insert() {
	tree := collections.NewBinaryTree[int]()

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	fmt.Print(tree)

	// Output:
	// [1 2 3]
}

func ExampleBinaryTree_Delete() {
	tree := collections.NewBinaryTree[int]()

	tree.Insert(1)
	tree.Insert(2)
	fmt.Println(tree)

	tree.Delete(1)
	tree.Delete(2)
	fmt.Println(tree)

	tree.Delete(-1)
	fmt.Println(tree)

	// Output:
	// [1 2]
	// []
	// []
}

func ExampleBinaryTree_Len() {
	tree := collections.NewBinaryTree[int]()
	for i := 0; i < 1000; i++ {
		tree.Insert(i)
	}

	fmt.Println(tree.Len())

	// Output:
	// 1000
}

func ExampleBinaryTree_Each() {
	tree := collections.NewBinaryTree[int]()
	for i := 0; i < 1000; i++ {
		tree.Insert(i)
	}

	sum := 0
	tree.Each(func(_, item int) {
		sum += item
	})

	fmt.Println(sum)

	// Output:
	// 499500
}

func ExampleBinaryTree_EachUntil() {
	tree := collections.NewBinaryTree[int]()
	for i := 0; i < 1000; i++ {
		tree.Insert(i)
	}

	sum := 0
	tree.EachUntil(func(_, item int) {
		sum += item
	}, func(_ int, item int) bool {
		return item == 3
	})

	fmt.Println(sum)

	// Output:
	// 6
}
