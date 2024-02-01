package example

import (
	"fmt"
	"time"
)

// 我们将要在协程中运行这个函数。 done 通道将被用于通知其他协程这个函数已经完成工作。
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们已经完工啦。
	done <- true
}
func ChannelSync() {
	// 我们可以使用通道来同步各个 goroutine 的执行。下面是使用阻塞接收来等待 goroutine 完成的示例。当等待多个 goroutine 完成时，您可能更喜欢使用 WaitGroup。
	done := make(chan bool)
	// 运行一个 worker 协程，并给予用于通知的通道。
	go worker(done)
	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知
	// 如果你把 <- done 这行代码从程序中移除， 程序甚至可能在 worker 开始运行前就结束
	<-done
	fmt.Println("任务完毕")
}
