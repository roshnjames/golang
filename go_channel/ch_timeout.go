package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 111
		time.Sleep(time.Second * 2)
	}()

	for i := 0; i < 2; i++ {
		select {
		case ans1 := <-ch1:
			fmt.Println(ans1)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 111")

		}
	}

	go func() {
		ch2 <- 222
		time.Sleep(time.Second * 4)
	}()
	for i := 0; i < 2; i++ {
		select {
		case ans2 := <-ch2:
			fmt.Println(ans2)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 222")
		}
	}
	defer close(ch1) //closing the channel means --no more send or recieve operations are posiible with that channel
	defer close(ch2)

}
