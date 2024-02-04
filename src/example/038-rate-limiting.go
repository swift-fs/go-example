package example

import (
	"fmt"
	"time"
)

func RateLimiting() {

	// 速率限制
	// 速率限制 是控制服务资源利用和质量的重要机制。

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(1 * time.Second)

	for request := range requests {
		<-limiter
		fmt.Println("request:", request)
	}

}
