package example

import "fmt"

// 这是一个函数，它将接受任意数量的 int 作为参数。
func sum(nums ...int) {

	// 在函数内， nums 的类型相当于 []int 。我们可以调用 len(nums) ，使用 range 迭代它，等等。
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func VarFunc() {
	// 可以使用任意数量的尾随参数来调用可变参数函数。
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	// 如果切片中已经有多个参数，请使用 func(slice...) 将它们应用到可变参数函数
	sum(nums...)
}
