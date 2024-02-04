package example

import (
	"fmt"
	"time"
)

func task(jobs chan int, results chan int) {
	for job := range jobs {
		time.Sleep(time.Second)
		results <- job * 2
	}
}
func Worker() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	go task(jobs, results)
	go task(jobs, results)
	go task(jobs, results)

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		result := <-results
		fmt.Println("result:", result)
	}

}
