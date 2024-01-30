package example

import "fmt"

// Go 支持匿名函数，它可以形成闭包。当您想要内联定义函数而无需命名它时，匿名函数非常有用。

// 此函数 intSeq 返回另一个函数，我们在 intSeq 主体中匿名定义该函数。返回的函数引用了变量 i 以形成闭包。
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func Closure() {

	nextInt := intSeq()
	// nextInt拥有上下文中的 i
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 闭包是隔离的
	newInts := intSeq()
	fmt.Println(newInts())
}
