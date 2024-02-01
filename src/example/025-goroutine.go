package example

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 20; i++ {
		fmt.Println("from", from, ":", i)
	}
}

func Goroutine() {
	f("direct")

	go f("goroutine")
	// 匿名函数调用启动一个goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 当我们运行这个程序时，我们首先看到阻塞调用的输出，然后是两个goroutine的输出。goroutine的输出可能是交错的，因为goroutine是由Go运行时并发运行的。
	time.Sleep(time.Second * 1)
	fmt.Println("done")

}
