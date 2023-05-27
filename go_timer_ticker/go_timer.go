package main

import (
	"fmt"
	"time"
)

func output(str string) {
	fmt.Println(str)
}

func main() {
	timer := time.NewTimer(time.Second * 5) //creating timer
	//tk := time.Tick(10 * time.Second)
	for i := 0; i < 10; {
		//timer.Stop()   //stops the timer
		select {

		case <-timer.C:
			panic("time out")
		default:
			output("TImer is not over")
			fmt.Println(i)
			time.Sleep(time.Second)
			i++
		}
	}
}

/*Timers represent a single event in the future.
You tell the timer how long you want to wait,
and it provides a channel(C) that will be notified at that time.
This timer will wait 5 seconds.
The <-timer1.C blocks on the timer’s channel C until it
sends a value indicating that the timer fired.*/
