package example

import (
	"fmt"
	"time"
)

func TimerExample() {

	timer := time.NewTimer(time.Second * 2)

	<-timer.C

	fmt.Println("Time's up!")

	timer2 := time.NewTimer(time.Second * 5)
	go func() {

		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {

		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(time.Second * 10)

}
