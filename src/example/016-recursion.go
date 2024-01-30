package example

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func Recursion() {

	// 递归函数就是函数调用自身，形成循环调用，必须有递归终止条件
	// 一般情况下递归都可以使用 for 循环替代，除非特殊情况，一般不建议使用递归
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
