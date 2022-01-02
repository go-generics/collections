package collections_test

import (
	"fmt"

	"github.com/go-generics/collections"
)

func ExampleNewSet() {
	set := collections.NewSet(0, 1)

	fmt.Println(set.Has(0))
	fmt.Println(set.Has(1))
	fmt.Println(!set.Has(2))

	// Output:
	// true
	// true
	// true
}

func ExampleSet_Insert() {
	set := collections.NewSet[int]()
	set.Insert(0)
	set.Insert(1)

	fmt.Println(set.Has(0))
	fmt.Println(set.Has(1))
	fmt.Println(!set.Has(2))

	// Output:
	// true
	// true
	// true
}

func ExampleSet_Delete() {
	set := collections.NewSet(0, 1)
	set.Delete(1)

	fmt.Println(set.Has(0))
	fmt.Println(!set.Has(1))

	// Output:
	// true
	// true
}

func ExampleSet_Each() {
	set := collections.NewSet(3, 4)
	sum := 0
	set.Each(func(_, item int) {
		sum += item
	})

	fmt.Println(sum)

	// Output:
	// 7
}

func ExampleSet_EachUntil() {
	set := collections.NewSet(3, 4)
	sum := 0
	set.EachUntil(func(_, item int) {
		sum += item
	}, func(_, item int) bool {
		return true
	})

	fmt.Println(set.Has(sum))

	// Output:
	// true
}

func ExampleUnion() {
	ab := collections.NewSet("A", "B")
	bc := collections.NewSet("B", "C")
	union := collections.Union(ab, bc)

	fmt.Println(union.Has("A"))
	fmt.Println(union.Has("B"))
	fmt.Println(union.Has("C"))
	fmt.Println(union.Len())

	// Output:
	// true
	// true
	// true
	// 3
}

func ExampleDifference() {
	ab := collections.NewSet("A", "B")
	bc := collections.NewSet("B", "C")

	fmt.Println(collections.Difference(ab, bc))
	fmt.Println(collections.Difference(bc, ab))

	// Output:
	// [A]
	// [C]
}

func ExampleIntersection() {
	ab := collections.NewSet("A", "B")
	bc := collections.NewSet("B", "C")
	intersection := collections.Intersection(ab, bc)

	fmt.Println(intersection.String())

	// Output
	// [B]
}

func ExampleIsSubset() {
	a := collections.NewSet(1.2)
	b := collections.NewSet(1.2, 3.4)

	fmt.Println(collections.IsSubset(a, a))
	fmt.Println(collections.IsSubset(a, b))
	fmt.Println(!collections.IsSubset(b, a))

	// Output:
	// true
	// true
	// true
}

func ExampleIsProperSubset() {
	a := collections.NewSet(1.2)
	b := collections.NewSet(1.2, 3.4)

	fmt.Println(!collections.IsProperSubset(a, a))
	fmt.Println(collections.IsProperSubset(a, b))
	fmt.Println(!collections.IsProperSubset(b, a))

	// Output:
	// true
	// true
	// true
}
