package collections

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type doublyLinkedNode[T ordered] struct {
	value T
	next  *doublyLinkedNode[T]
	prev  *doublyLinkedNode[T]
}
