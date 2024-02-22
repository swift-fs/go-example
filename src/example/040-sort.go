package example

import (
	"fmt"
	"slices"
	"sort"
)

type Person struct {
	Name string
	Age  uint
}

type People []Person

// People 实现 sort 接口实现排序
func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Age > p[j].Age
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort() {
	names := []string{"c", "a", "b"}
	sort.Strings(names)
	fmt.Println("names:", names)

	names = []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	slices.Sort[[]string](names)
	fmt.Println("names:", names)

	persons := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	slices.SortFunc[[]Person](persons, func(a, b Person) int {
		return int(a.Age) - int(b.Age)
	})
	fmt.Println("persons:", persons)

	var people People = persons
	sort.Sort(people)
	fmt.Println("people:", people)
}
