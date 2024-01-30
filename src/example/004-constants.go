package example

import "fmt"

const Pi = 3.14

func Constants() {
	fmt.Println("Pi:", Pi)
	const World = "世界"
	fmt.Println("World:", World)

	const d = 5000
	const e = 3e20 / d
	fmt.Println("e:", e)
	const f = d + 10
	fmt.Println("f:", f)

}
