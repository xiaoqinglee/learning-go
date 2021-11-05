package supplement

import (
	"fmt"
	"time"
)

func Timer() {
	//Timer的特殊之处在于它可以在目标时间来临前提供取消继续计时的机会
	var timer1 *time.Timer = time.NewTimer(4 * time.Second)
	var timeChan <-chan time.Time = timer1.C
	<-timeChan
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(4 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	/*
		Stop prevents the Timer from firing.
		It returns true if the call stops the timer, false if the timer has already expired or been stopped.
		Stop does not close the channel, to prevent a read from the channel succeeding incorrectly.
		To ensure the channel is empty after a call to Stop, check the return value and drain the channel.
		For example, assuming the program has not received from t.C already:
		 if !t.Stop() {
		  <-t.C
		 }
	*/
	//因为有receiver goroutine在接受chan上的元素, 所以此处不额外进行drain chan动作.

	//time.Sleep(10 * time.Second)

	ok := timer2.Stop()
	if ok {
		fmt.Println("succeeded to stop timer2")
	}

	time.Sleep(10 * time.Second)
}
