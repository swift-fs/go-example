package example

import "fmt"

func MapKeys[Key comparable, V any](m map[Key]V) []Key {
	r := make([]Key, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func Generic() {
	// 一开始写程序时不必考虑泛型，当后续使用时，逻辑一样，但是数据类型不同时可以重构为泛型
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(MapKeys(m))

	_ = MapKeys(m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
