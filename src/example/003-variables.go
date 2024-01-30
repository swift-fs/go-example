package example

import "fmt"

func Variables() {
	var a string = "initial"
	fmt.Println("a:", a)
	var b, c int = 100, 200
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	var d = true
	fmt.Println("d:", d)
	var e int
	fmt.Println("e:", e)
	f := "orange"
	fmt.Println("f:", f)
}
