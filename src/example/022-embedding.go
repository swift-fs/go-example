package example

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

func Embed() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	// 嵌入结构可以用于将接口实现赋予其他结构。这里我们看到，一个 container 现在实现了 describer 接口，因为它嵌入了 base 。
	var d describer = co
	fmt.Println("describer:", d.describe())
}
