package example

import "fmt"

func For() {
	i := 1
	for i <= 3 {
		fmt.Println("i:", i)
		i = i + 1
	}

	for j := 0; j < 7; j++ {
		fmt.Println("j:", j)
	}
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n:", n)
	}
}
