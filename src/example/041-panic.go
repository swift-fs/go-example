package example

import "fmt"

func Panic() {
	panic("Panic!")
	fmt.Println("此处无法执行")
}
