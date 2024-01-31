package example

import (
	"fmt"
	"math"
)

// 接口是方法签名的命名集合。
type geometry interface {
	area() float64
	perim() float64
}

type react02 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要在Go中实现一个接口，我们只需要实现接口中的所有方法
func (r react02) area() float64 {
	return r.width * r.height
}
func (r react02) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量有一个接口类型，那么我们可以调用命名接口中的方法
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interface() {
	r := react02{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
