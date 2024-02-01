package example

import (
	"fmt"
	"time"
)

func Timeout() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "result 2"
	}()

	// 修改为 2,3,4 查看运行结果
	d := 4

	// select 默认处理第一个已准备好的接收/发送操作
	select {
	case res := <-c1:
		fmt.Println(res)
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * time.Duration(d)):
		fmt.Println("timeout")
	}
	fmt.Println("done")
}
