package example

import "fmt"

func Channel() {

	// 默认情况下，发送和接收会阻塞，直到发送者和接收者都准备好为止。
	// 此特性允许我们在程序结束时等待 "ping" 消息，而无需使用任何其他同步。
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println("msg:", msg)

}
