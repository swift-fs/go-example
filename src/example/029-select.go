package example

import (
	"fmt"
	"time"
)

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在一定时间后接收一个值， 通过这种方式来模拟并行的协程执行时造成的阻塞（耗时）。
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		c2 <- "two"
	}()

	// 我们使用 select 关键字来同时等待这两个值， 并打印各自接收到的值。
	// 注意，程序总共仅运行了4s左右。因为 2 秒 和 4 秒的 Sleeps 是并发执行的，
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}