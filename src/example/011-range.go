package example

import "fmt"

func Range() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 字符串上的 range 迭代 Unicode 代码点。第一个值是 rune 的起始字节索引，第二个值是 rune 本身。有关更多详细信息，请参阅字符串和符文。
	for i, c := range "大河之水天上来，Go!" {
		fmt.Println(i, string(c))
	}
}
