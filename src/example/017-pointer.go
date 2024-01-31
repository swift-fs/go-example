package example

import "fmt"

func val(v int) {
	v = 10
}
func ptr(v *int) {
	*v = 10
}
func Pointer() {
	age := 0
	val(age)
	fmt.Println("age:", age)
	ptr(&age)
	fmt.Println("age:", age)
}
