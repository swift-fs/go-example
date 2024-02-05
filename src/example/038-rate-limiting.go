package example

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func RateLimiting() {

	// 速率限制
	// 速率限制 是控制服务资源利用和质量的重要机制。

	// 基本速率限制
	// 每秒钟允许 1 个请求

	// requests := make(chan int, 5)
	// go func() {
	// 	i := 1
	// 	for {
	// 		requests <- i
	// 		i++
	// 	}

	// }()

	// go func() {
	// 	limiter := time.Tick(1 * time.Second)

	// 	for request := range requests {
	// 		<-limiter
	// 		fmt.Println("request:", request)
	// 	}
	// }()

	// 漏桶算法速率限制
	// 漏桶算法则会以固定的速率从桶中漏水，当请求到来时，如果桶中有足够的空间来装下该请求，就允许处理该请求，否则拒绝或者等待，直到桶中有足够的空间来装下请求
	ratelimit01 := ratelimit.New(10)
	prev := time.Now()
	go func() {
		i := 0
		for {
			i++
			ratelimit01.Take()
			now := time.Now()
			fmt.Println(i, now.Sub(prev))
			prev = now

		}
	}()

	// 令牌桶算法会以固定的速率往桶里放入令牌，每个令牌代表一个可用的请求。如果桶满了，多余的令牌会被丢弃。当一个请求到来时，如果有可用的令牌，则允许处理该请求，否则拒绝该请求。

	select {}

}
