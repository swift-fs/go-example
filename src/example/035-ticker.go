package example

import (
	"fmt"
	"time"
)

func Ticker() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	t := make(chan int)
	close(t)
	// 从一个已关闭的通道接收值会立即返回该通道类型的零值，并且不会阻塞。
	fmt.Println("t is closed:", <-t)
}
