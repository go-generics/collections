package collections_test

import (
	"fmt"

	"github.com/go-generics/collections"
)

func ExampleImplicitTreap() {
	t := collections.NewImplicitTreap(3, 1, 4, 1, 5, 4, 9, 2)
	fmt.Println(t)

	// Output:
	// 3 1 4 1 5 4 9 2
}
