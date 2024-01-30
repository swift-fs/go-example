package example

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func MultReturn() {
	// Go 内置了对多个返回值的支持。此功能在 Go 中经常使用，例如从函数返回结果值和错误值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果您只需要返回值的子集，请使用空白标识符 _ 。
	_, c := vals()
	fmt.Println(c)
}
