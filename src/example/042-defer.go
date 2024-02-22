package example

import "fmt"

func Defer() {
	defer func() {
		fmt.Println("world")
	}()
	fmt.Println("hello")
}
