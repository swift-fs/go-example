package example

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name] += 1
}

func Mutex() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	c := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}
	go func() {
		for i := 0; i < 10000; i++ {
			c.inc("a")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			c.inc("a")
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("counters:", c.counters)

}
