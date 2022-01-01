package collections_test

import (
	"fmt"

	"github.com/go-generics/collections"
)

func ExampleDeque_PushBack() {
	deque := collections.NewDeque[int]()

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)

	fmt.Print(deque)

	// Output:
	// [1 2 3]
}

func ExampleDeque_PushFront() {
	deque := collections.NewDeque[int]()

	deque.PushFront(6)
	deque.PushFront(5)
	deque.PushFront(4)

	fmt.Print(deque)

	// Output:
	// [4 5 6]
}

func ExampleDeque_PopBack() {
	deque := collections.NewDeque[int]()

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)
	deque.PushBack(4)

	fmt.Println(deque.PopBack())
	fmt.Println(deque)

	// Output:
	// 4
	// [1 2 3]
}

func ExampleDeque_PopFront() {
	deque := collections.NewDeque[int]()

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)
	deque.PushBack(4)

	fmt.Println(deque.PopFront())
	fmt.Println(deque)

	// Output:
	// 1
	// [2 3 4]
}

func ExampleDeque_Back() {
	deque := collections.NewDeque[int]()

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)
	deque.PushBack(4)

	fmt.Println(deque.Back())

	// Output:
	// 4
}

func ExampleDeque_Front() {
	deque := collections.NewDeque[int]()

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)
	deque.PushBack(4)

	fmt.Println(deque.Front())

	// Output:
	// 1
}

func ExampleDeque_Len() {
	deque := collections.NewDeque[int]()
	for i := 0; i < 1000; i++ {
		deque.PushBack(i)
	}

	fmt.Println(deque.Len())
	fmt.Println(deque.Front())
	fmt.Println(deque.Back())

	// Output:
	// 1000
	// 0
	// 999
}

func ExampleDeque_Each() {
	deque := collections.NewDeque[int]()
	for i := 0; i < 1000; i++ {
		deque.PushBack(i)
	}

	sum := 0
	deque.Each(func(_, item int) {
		sum += item
	})

	fmt.Println(sum)

	// Output:
	// 499500
}

func ExampleDeque_EachUntil() {
	deque := collections.NewDeque[int]()
	for i := 0; i < 1000; i++ {
		deque.PushBack(i)
	}

	sum := 0
	deque.EachUntil(func(_, item int) {
		sum += item
	}, func(_ int, item int) bool {
		return item == 3
	})

	fmt.Println(sum)

	// Output:
	// 6
}
