package example

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Worker 1 done")

	}()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Worker 2 done")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All workers done")
}
