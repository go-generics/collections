package collections_test

import (
	"fmt"

	"github.com/go-generics/collections"
)

func ExampleHeap() {
	h := collections.NewHeap(2, 1, 5)
	h.PushFront(3)
	fmt.Printf("minimum: %d\n", h.Front())
	for h.Len() > 0 {
		fmt.Print(h.PopFront(), " ")
	}

	// Output:
	// minimum: 1
	// 1 2 3 5
}
