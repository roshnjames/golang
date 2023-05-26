package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "roshin"
		time.Sleep(time.Millisecond * 2)
	}()

	go func() {
		ch2 <- "james"
		time.Sleep(time.Millisecond * 20)
	}()

	for i := 0; i < 2; i++ {
		select {
		case ans1 := <-ch1:
			fmt.Println(ans1)
		case ans2 := <-ch2:
			fmt.Println(ans2)

		}
	}
	time.Sleep(time.Second)

}

//'select' keyword allows the selection of concurrently executing goroutines
