package example

import "fmt"

func CloseChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是工作协程。使用 j, more := <- jobs 循环的从 jobs 接收数据。 根据接收的第二个值，如果 jobs 已经关闭了， 并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。 当我们完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
	fmt.Println("done")
}
